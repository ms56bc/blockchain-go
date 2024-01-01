package pd

import (
	"fmt"
	"github.com/ms56bc/blockchain-go/pkg/server/model"
	"google.golang.org/protobuf/proto"
	"io"
	"net"
)

type PeerDiscoveryClient interface {
	SendDiscoveredPeers(sourceId string, destination model.Peer, discoveredPeers []model.Peer) error
	Close() error
}
type PeerDiscoveryUDPClient struct {
	connections map[string]*net.UDPConn
}

func NewDiscoveryClient() *PeerDiscoveryUDPClient {
	return &PeerDiscoveryUDPClient{connections: make(map[string]*net.UDPConn)}
}

func (pool *PeerDiscoveryUDPClient) SendDiscoveredPeers(sourceId string, destination model.Peer, discoveredPeers []model.Peer) error {
	if pool.connections[destination.ID] == nil {
		conn, err := pool.createUdpClientConnection(destination)
		pool.connections[destination.ID] = conn
		if err != nil {
			return err
		}
	}

	conn := pool.connections[destination.ID]
	pool.sendGossipMessage(conn, discoveredPeers, sourceId)
	return nil
}

func (pool *PeerDiscoveryUDPClient) createUdpClientConnection(peer model.Peer) (*net.UDPConn, error) {
	fmt.Println("Starting UDP client..." + peer.IP)
	// Resolve UDP address
	serverAddr, err := net.ResolveUDPAddr("udp", peer.IP)
	if err != nil {
		return nil, err
	}

	// Create UDP connection
	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Println("Error creating UDP connection:", err)
		return nil, err
	}
	return conn, nil
}

func (pool *PeerDiscoveryUDPClient) sendGossipMessage(conn io.Writer, discoveredPeers []model.Peer, sourceId string) {
	// Marshal to protobuf
	for _, v := range toProtoMessage(discoveredPeers, sourceId) {
		protoData, err := proto.Marshal(v)
		if err != nil {
			fmt.Println("Error marshaling protobuf:", err)
			return
		}

		// Write to the UDP connection
		_, err = conn.Write(protoData)
		if err != nil {
			fmt.Println("Error writing to UDP:", err)
		}
	}
}

func (pool *PeerDiscoveryUDPClient) Close() error {
	for _, conn := range pool.connections {
		err := conn.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func toProtoMessage(discoveredPeers []model.Peer, sourceId string) []*NodeUpdate {
	var msg []*NodeUpdate
	for _, peer := range discoveredPeers {
		msg = append(msg, &NodeUpdate{
			Operation: NodeUpdate_ADD,
			NodeInfo: &NodeInfo{
				NodeId:    peer.ID,
				IpAddress: peer.IP,
			},
			SourceId: sourceId,
		})
	}
	return msg
}
