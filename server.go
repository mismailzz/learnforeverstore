package main

import (
	"io"
	"log"

	"github.com/mismailzz/learnforeverstore/p2p"
)

type FileServerOpts struct {
	transport p2p.Transport
}

type FileServer struct {
	FileServerOpts
	peerNodeList     []string
	ConnectedPeerMap map[string]p2p.Peer
	store            Store
}

func NewFileServer(opts FileServerOpts, storeOpts StoreOpts, peerList []string) *FileServer {
	return &FileServer{
		FileServerOpts:   opts,
		peerNodeList:     peerList,
		ConnectedPeerMap: make(map[string]p2p.Peer),
		store:            *NewStore(storeOpts),
	}
}

// Start() - start the server (node) to
// - listen for peers
// - read on its channel
// - dial to other peers
func (s *FileServer) Start() error {

	// Start the Server (peer/node)
	if err := s.transport.ListenAndAccept(); err != nil {
		return err
	}

	go s.readLoop()

	s.peerNodeDial()

	return nil
}

// read rpc message of server read channel
func (s *FileServer) readLoop() {

	for {
		rpc := <-s.transport.Consume() // blocking call
		log.Printf("msg recieved %s from %+v\n", rpc.Payload, rpc.From)
	}
}

// dial to the peers (other nodes/server)
func (s *FileServer) peerNodeDial() {

	for _, address := range s.peerNodeList {
		if err := s.transport.Dial(address); err != nil {
			log.Printf("error while dialing %s:%+v\n", address, err)
			continue
		}
	}
}

// Action need to be take when connection accepted in handleNewConnection (TCPTransport)
func (s *FileServer) OnPeer(peer p2p.Peer) error {
	s.ConnectedPeerMap[peer.RemoteAddress()] = peer
	return nil
}

func (s *FileServer) StoreData(filename string, r io.Reader) error {
	return s.store.writeStream(filename, r)
}
