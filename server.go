package main

import (
	"log"

	"github.com/mismailzz/learnforeverstore/p2p"
)

type FileServerOpts struct {
	transport p2p.Transport
}

type FileServer struct {
	FileServerOpts
}

func NewFileServer(opts FileServerOpts) *FileServer {

	return &FileServer{
		FileServerOpts: opts,
	}
}

func (s *FileServer) Start() error {

	// Start the Server (peer/node)
	if err := s.transport.ListenAndAccept(); err != nil {
		return err
	}

	s.connChanReadLoop()

	return nil
}

func (s *FileServer) connChanReadLoop() {

	for {
		rpc := <-s.transport.Consume()
		log.Printf("msg recieved %s from %+v\n", rpc.Payload, rpc.From)
	}
}
