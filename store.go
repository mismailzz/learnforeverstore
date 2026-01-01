package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type PathTransformFunction func(string) FilePath

type StoreOpts struct {
	rootDir           string
	pathTransformFunc PathTransformFunction
}

type Store struct {
	StoreOpts
}

func NewStore(opts StoreOpts) *Store {
	return &Store{
		opts,
	}
}

func (s *Store) writeStream(filename string, r io.Reader) error {

	filePath := s.pathTransformFunc(filename)
	pathWithRootDir := fmt.Sprintf("%s", s.rootDir+"/"+filePath.pathname)

	// Create the Dir (recursively)
	if err := os.MkdirAll(pathWithRootDir, os.ModePerm); err != nil {
		return err
	}

	fullFilename := fmt.Sprintf("%s", pathWithRootDir+"/"+filePath.filename)

	// Create the file
	file, err := os.Create(fullFilename)
	if err != nil {
		return err
	}
	defer file.Close()

	// io.Copy: gives two flexibility
	// 1. stream data to the dst rather then direct whole copy
	// 2. it uses io.Reader that allows any input type (network, file, bytes, etc.)
	n, err := io.Copy(file, r)
	if err != nil {
		return err
	}

	log.Printf("%+v bytes has been written to disk: %s\n", n, fullFilename)

	return nil
}

func (s *Store) readStream(filename string) error {

	filePath := s.pathTransformFunc(filename)
	pathWithRootDir := fmt.Sprintf("%s", s.rootDir+"/"+filePath.pathname)
	fullFilename := fmt.Sprintf("%s", pathWithRootDir+"/"+filePath.filename)

	if !fileExists(fullFilename) {
		return fmt.Errorf("file %s doesnt exists", filename)
	}

	// Open file
	file, err := os.Open(fullFilename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Stream the file content into buff, as we know its bytes
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, file)
	if err != nil {
		return err
	}

	data := buf.Bytes()
	log.Printf("read msg: %+v from file: %+v \n", string(data), fullFilename)

	return nil

}

func (s *Store) Delete(filename string) error {

	// Delete the file with RootDir

	filePath := s.pathTransformFunc(filename)

	defer func() {
		log.Printf("deleted [%s] from disk", filePath.filename)
	}()

	pathWithRootDir := fmt.Sprintf("%s", s.rootDir+"/"+filePath.pathname)
	fullFilename := fmt.Sprintf("%s", pathWithRootDir+"/"+filePath.filename)

	if !fileExists(fullFilename) {
		return fmt.Errorf("file %s doesnt exists", filename)
	}

	// It can delete only file but not the path
	if err := os.RemoveAll(fullFilename); err != nil {
		return err
	}
	// due to which we delete the parent directory - which is wierd but workaround
	// need double deletion
	parentDir := s.rootDir
	if err := os.RemoveAll(parentDir); err != nil {
		return err
	}

	return nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

type FilePath struct {
	filename string
	pathname string
}

func CASPathTransformFunc(filename string) FilePath {

	// Create determistic hash from same key using SHA1
	hash := sha1.Sum([]byte(filename))
	// Convert the bytes to hex string for hash
	hashStr := hex.EncodeToString(hash[:])

	// Split the hash string into multiple parts for directory structure (depth levels)
	blocksize := 5
	sliceLen := len(hashStr) / blocksize
	paths := make([]string, sliceLen)

	for i := 0; i < sliceLen; i++ {
		from, to := i*blocksize, (i*blocksize)+blocksize
		paths[i] = hashStr[from:to]
	}

	return FilePath{
		filename: hashStr,
		pathname: strings.Join(paths, "/"), // Join the parts with "/" to form the final path

	}
}
