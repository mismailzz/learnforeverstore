package main

import (
	"fmt"
	"time"

	"github.com/mismailzz/learnforeverstore/p2p"
)

func main() {

	// opts := p2p.TCPTransportOpts{
	// 	ListenAddress: ":3000",
	// 	HandshakeFunc: p2p.NoHandshakeFunc,
	// 	Decode:        &p2p.DefaultDecoder{},
	// }
	// t := p2p.NewTCPTransport(opts)
	// if err := t.ListenAndAccept(); err != nil {
	// 	log.Fatal(err)
	// }

	// opts := StoreOpts{
	// 	rootDir:           "db",
	// 	pathTransformFunc: CASPathTransformFunc,
	// }

	// store := NewStore(opts)
	// store.writeStream("example.txt", bytes.NewReader([]byte("hello world")))
	// store.readStream("example.txt")
	// store.Delete("example.txt")

	tcpOpts := p2p.TCPTransportOpts{
		ListenAddress: ":3000",
		HandshakeFunc: p2p.NoHandshakeFunc,
		Decode:        &p2p.DefaultDecoder{},
	}
	tcpTransport := p2p.NewTCPTransport(tcpOpts)

	serverOpts := FileServerOpts{
		transport: tcpTransport,
	}

	peerList := []string{":3000"}
	server1 := NewFileServer(serverOpts, peerList)
	tcpTransport.OnPeer = server1.OnPeer

	server1.Start()

	time.Sleep(2 * time.Second)
	// Check connected peers in server1
	fmt.Println(server1.connectedPeerMap)

	select {} // block
}
