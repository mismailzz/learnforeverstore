package main

import (
	"log"

	"github.com/mismailzz/learnforeverstore/p2p"
)

func main() {

	opts := p2p.TCPTransportOpts{
		ListenAddress: ":3000",
		HandshakeFunc: p2p.NoHandshakeFunc,
		Decode:        &p2p.DefaultDecoder{},
	}
	t := p2p.NewTCPTransport(opts)
	if err := t.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

}
