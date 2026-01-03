package main

import (
	"time"

	"github.com/mismailzz/learnforeverstore/p2p"
)

func main() {

	s1 := makeServer(":3000", "")
	s1.Start()

	time.Sleep(2 * time.Second)

	s2 := makeServer(":4000", ":3000") // s2 connect to s1
	s2.Start()

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
