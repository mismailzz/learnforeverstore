package p2p

// Transport interface will hold the func(), that would be required for
// the protocols like TCP, UDP, GRPC, etc. The funcs will be added here
// based on the need
type Transport interface{
	ListenAndAccept()
}

// Peer represent any node (connection) part of the p2p network 
// every protcol like TCP, UDP, GRPC etc has to define the Peer
type Peer interface {}


// OnPeer func to take a needed action on the peer
// when the connection is established (if required)
type OnPeer func(Peer) error