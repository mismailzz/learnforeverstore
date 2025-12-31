package main

import (
	"log"

	"github.com/mismailzz/learnforeverstore/p2p"
)

func main() {

	t := p2p.NewTCPTransport(":3000")
	if err := t.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

}
