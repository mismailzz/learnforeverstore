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
	peerNodeList     []string
	connectedPeerMap map[string]p2p.Peer
}

func NewFileServer(opts FileServerOpts, peerList []string) *FileServer {
	return &FileServer{
		FileServerOpts:   opts,
		peerNodeList:     peerList,
		connectedPeerMap: make(map[string]p2p.Peer),
	}
}

func (s *FileServer) Start() error {

	// Start the Server (peer/node)
	if err := s.transport.ListenAndAccept(); err != nil {
		return err
	}

	go s.readLoop()

	s.peerNodeDial()

	return nil
}

func (s *FileServer) readLoop() {

	for {
		rpc := <-s.transport.Consume() // blocking call
		log.Printf("msg recieved %s from %+v\n", rpc.Payload, rpc.From)
	}
}

func (s *FileServer) peerNodeDial() {

	for _, address := range s.peerNodeList {
		if err := s.transport.Dial(address); err != nil {
			log.Printf("error while dialing %s:%+v\n", address, err)
			continue
		}
		log.Printf("Dialed to address: %s\n", address)
	}
}

func (s *FileServer) OnPeer(peer p2p.Peer) error {
	s.connectedPeerMap[peer.RemoteAddress()] = peer
	return nil
}
