package main

import (
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

func (s *Store) writeStream(filename string, data io.Reader) error {

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
	n, err := io.Copy(file, data)
	if err != nil {
		return err
	}

	log.Printf("%+v bytes has been written to disk\n", n)

	return nil
}
