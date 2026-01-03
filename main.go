package main

import (
	"fmt"
	"time"

	"github.com/mismailzz/learnforeverstore/p2p"
)

func main() {

	// server-1
	s1 := makeServer(":3000", "")
	s1.Start()

	time.Sleep(2 * time.Second)

	// server-2
	s2 := makeServer(":4001", ":3000") // s2 connect to s1
	s2.Start()

	time.Sleep(2 * time.Second)
	// server-3
	s3 := makeServer(":4002", ":3000", ":4001")
	if err := s3.Start(); err != nil {
		fmt.Println(err)
		panic(err)
	}

	time.Sleep(2 * time.Second)
	// Check Connected Peer Dict:
	fmt.Printf("Server 1 - PeerMap: %+v\n", s1.ConnectedPeerMap)
	fmt.Printf("Server 2 - PeerMap: %+v\n", s2.ConnectedPeerMap)
	fmt.Printf("Server 3 - PeerMap: %+v\n", s3.ConnectedPeerMap)

	select {} // block
}

func makeServer(listenAddr string, peerlist ...string) *FileServer {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddress: listenAddr,
		HandshakeFunc: p2p.NoHandshakeFunc,
		Decode:        &p2p.DefaultDecoder{},
	}

	tcpTransport := p2p.NewTCPTransport(tcpOpts)

	serverOpts := FileServerOpts{
		transport: tcpTransport,
	}

	peerList := peerlist
	server := NewFileServer(serverOpts, peerList)
	tcpTransport.OnPeer = server.OnPeer

	return server
}
