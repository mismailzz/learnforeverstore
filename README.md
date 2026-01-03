# learnforeverstore

<details>
<summary> PHASE - I - TCP Implementation</summary>

- 1.1: Simple TCP prototype defined 
- 1.2: Handshake functionality addded using callback func
- 1.3: handleNewConnection modifed to read message of its connection (send from telnet)
- 1.3: (Revised) Refactor
- 1.4: Decoder and Refactor (added TCPTransportOpts)
- 1.5: RPC Message implemented (to define payload for comms), to also take the message out from Decoder func to TCPTransport
- 1.6: Peer has been defined with TCPPeer and some refactor (replacing conn with peer)
- 1.7 - OnPeer func added - to take action or can be used to as Notification to take action when the connection (peer) is establishedv

</details>

> **Note:**
> At the point, we are able to create the TCP lib, to which following tasks
> can be done
> - Server(Node) instance start to listen on a sepcified port
> - Upcoming new connections after being accepted, will be handled independently (Handshake, Peer Conversion, Decode) and reading from its conn READ stream to stdout RPC Message
> - Similarly that independent connection will be closed on disconnection from Server (cause EOF) but except this error, it will keep reading unless Server is down

<details>
<summary> PHASE - II - Store Implementation</summary>

- 2.1: Store created with writeStream, to create a file on the disk
- 2.2: readStream and Delete function added
- 2.3: PathTransformFunc defined as it deterministic hash the name of file and derived the pathname from the hash (helpful for discovery), FilePath also added to organize the filename/pathname, refactor the delete/other funcs logic

</details>

> **Note:**
> Create the simple Store lib, having following functions
> - Store the file in the disk, Read the file from the disk
> - Also can delete the file from the disk 


<details>
<summary> PHASE - III - Server lib Implementation </summary>

- 3.1: Setup a simple server, which uses the Transport (TCP) lib to start the server
- 3.2: AcceptLoop (TCPTransport) made go routine, the telnet connection written rpc message, are now being read in the server using TCPTransport read unbuffered channel
- 3.3: Setup a DialUp inside server, to which use by providing the list of peer/nodes to which we can dial - currently we are connecting to server - also eliminate the use of telnet manually
- 3.4: Create a connectedPeerMap on the Server, so we will add newly created peer (conn) - maintain a list which peer are connected to server
</details>

<details>
<summary> PHASE - IV - Refactor </summary>

- 4.1: Refactor the main to push the logic in a func, based on that created two servers (peers) and connect them

</details>


```bash

# Below is our current setup - PDF p2p1.pdf
2026/01/03 19:31:56 Server started listening: :3000
2026/01/03 19:31:56 error while dialing :dial tcp: missing address
2026/01/03 19:31:58 Server started listening: :4001
2026/01/03 19:31:58 Handling the upcoming connection: &{conn:{fd:0x14000120100}}
2026/01/03 19:32:00 Server started listening: :4002
2026/01/03 19:32:00 Handling the upcoming connection: &{conn:{fd:0x14000120280}}
2026/01/03 19:32:00 Handling the upcoming connection: &{conn:{fd:0x1400009a080}}
Server 1 - PeerMap: map[127.0.0.1:49269:0x14000122018 127.0.0.1:49272:0x14000122030]
Server 2 - PeerMap: map[127.0.0.1:49273:0x140000a0018]
Server 3 - PeerMap: map[]
```

<details>
<summary> PHASE - V - (Server lib cont. ) Broadcast Message to Peers </summary>

- 5.1: By this time, we were using TCP lib only in the FileServer but now we defined the store, so that we can read/write RPC message. As for now the RPC message is written to disk using writeStream 

</details>