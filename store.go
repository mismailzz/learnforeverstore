package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type StoreOpts struct {
	rootDir string
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

	// Create the Dir (recursively)
	if err := os.MkdirAll(s.rootDir, os.ModePerm); err != nil {
		return err
	}

	fileWithPathname := fmt.Sprintf(s.rootDir + "/" + filename)

	// Create the file
	file, err := os.Create(fileWithPathname)
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

	log.Printf("%+v bytes has been written to disk\n", n)

	return nil
}

func (s *Store) readStream(filename string) error {

	fileWithPathname := fmt.Sprintf("%s", s.rootDir+"/"+filename)
	if !fileExists(fileWithPathname) {
		return fmt.Errorf("file %s doesnt exists", filename)
	}

	// Open file
	file, err := os.Open(fileWithPathname)
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
	log.Printf("read msg: %+v\n", string(data))

	return nil

}

func (s *Store) Delete(filename string) error {
	fileWithPathname := fmt.Sprintf("%s", s.rootDir+"/"+filename)
	if !fileExists(fileWithPathname) {
		return fmt.Errorf("file %s doesnt exists", filename)
	}
	return os.RemoveAll(s.rootDir)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
