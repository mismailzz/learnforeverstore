package main

import "github.com/mismailzz/learnforeverstore/p2p"

type FileServerOpts struct {
	transport p2p.Transport
}

type FileServer struct {
	FileServerOpts
}

func NewFileServer(opts FileServerOpts) *FileServer {
	return &FileServer{
		opts,
	}
}

func (s *FileServer) Start() error {
	if err := s.transport.ListenAndAccept(); err != nil {
		return err
	}

	return nil
}
