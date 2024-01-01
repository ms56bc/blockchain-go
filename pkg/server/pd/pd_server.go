package pd

import (
	"context"
	"fmt"
	"github.com/ms56bc/blockchain-go/pkg/server/config"
	"google.golang.org/protobuf/proto"
	"io"
	"net"
)

type PeerDiscoveryServer interface {
	HandleIncomingMessage(conn io.Reader) error
}
type PeerDiscoveryUDPServer struct {
	node   *BlockchainNode
	config config.Config
}

func NewPeerDiscoveryServer(node *BlockchainNode, config config.Config) *PeerDiscoveryUDPServer {
	return &PeerDiscoveryUDPServer{
		node:   node,
		config: config,
	}
}

func (server *PeerDiscoveryUDPServer) StartGossipUDPServer(ctx context.Context) {
	// Resolve UDP address
	addr, err := net.ResolveUDPAddr("udp", server.config.GossipAddress)
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	// Create UDP listener
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error creating UDP listener:", err)
		return
	}
	defer conn.Close()

	fmt.Println("UDP server started on", addr)

	// Handle incoming messages
	for {
		select {
		case <-ctx.Done():
			return
		default:
			err := server.HandleIncomingMessage(conn)
			if err != nil {
				panic(err)
			}
		}
	}
}
func (server *PeerDiscoveryUDPServer) HandleIncomingMessage(conn io.Reader) error {
	buffer := make([]byte, 1024)

	// Read from the UDP connection
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading from UDP:", err)
		return err
	}

	var receivedMessage NodeUpdate

	// Unmarshal protobuf
	err = proto.Unmarshal(buffer[:n], &receivedMessage)
	if err != nil {
		fmt.Println("Error unmarshaling protobuf:", err)
		return err
	}
	server.node.ProcessMessage(&receivedMessage)
	return nil
}
