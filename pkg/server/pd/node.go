package pd

import (
	"context"
	"fmt"
	"github.com/ms56bc/blockchain-go/pkg/server/config"
	"github.com/ms56bc/blockchain-go/pkg/server/model"
	"math/rand"
	"sync"
	"time"
)

type NetworkDiscoveryNode interface {
	DiscoverPeers(ctx context.Context)
	GetNetworkPeers() []model.Peer
	ProcessMessage(message *NodeUpdate)
}

// BlockchainNode represents a node in the blockchain network.
type BlockchainNode struct {
	ID                 string
	Peers              map[string]model.Peer
	peerLock           sync.Mutex
	bootstrap          chan model.Peer
	discovery          chan model.Peer
	disconnect         chan string
	peerConnectionPoll PeerDiscoveryClient
}

// NewBlockchainNode creates a new blockchain node with the given ID.
func NewBlockchainNode(id string, config config.Config) *BlockchainNode {
	b := &BlockchainNode{
		ID:                 id,
		Peers:              make(map[string]model.Peer),
		discovery:          make(chan model.Peer),
		disconnect:         make(chan string),
		peerConnectionPoll: NewDiscoveryClient(),
	}
	self := model.Peer{
		ID: id,
		IP: config.GossipAddress,
	}
	b.broadcastSelfIpToSeedNodes(config.StaticPeers, self)
	return b
}
func (node *BlockchainNode) ProcessMessage(message *NodeUpdate) {
	if message.GetSourceId() == node.ID || message.NodeInfo.GetNodeId() == node.ID {
		return
	}
	fmt.Println("Processing message from", message.GetSourceId())
	switch message.Operation {
	case NodeUpdate_ADD:
		peer := model.Peer{
			ID: message.NodeInfo.GetNodeId(),
			IP: message.NodeInfo.GetIpAddress(),
		}
		node.addPeer(peer)
		fmt.Printf("Add operation: %+v\n", message.NodeInfo)
	case NodeUpdate_REMOVE:
		node.removePeer(message.NodeInfo.GetNodeId())
		fmt.Printf("Remove operation: %+v\n", message.NodeInfo)
	}
}

// Gossip periodically broadcasts the node's information to a random subset of other nodes.
func (node *BlockchainNode) DiscoverPeers(ctx context.Context) {
	defer node.peerConnectionPoll.Close()
	fmt.Println("Gossiping for server", node.ID)
	for {
		// Simulate gossiping by periodically broadcasting the node's information
		time.Sleep(5 * time.Second)

		node.peerLock.Lock()
		peers := make([]model.Peer, 0, len(node.Peers))
		for _, peer := range node.Peers {
			peers = append(peers, peer)
		}
		node.peerLock.Unlock()

		// Randomly select a subset of peers to gossip to
		selectedPeers := node.selectRandomPeers(peers, 2)
		// Broadcast the node's information to the selected peers
		for _, selectedPeer := range selectedPeers {
			node.sendKnownHosts(selectedPeer)
		}
	}
}

// Bootstrap initializes the node with a bootstrap peer.
func (node *BlockchainNode) broadcastSelfIpToSeedNodes(bootstrapPeer []model.Peer, self model.Peer) {
	for _, peer := range bootstrapPeer {
		node.addPeer(peer)
		err := node.peerConnectionPoll.SendDiscoveredPeers(self.ID, peer, []model.Peer{self})
		if err != nil {
			panic(err)
		}
	}
}

func (node *BlockchainNode) addPeer(peer model.Peer) {
	node.peerLock.Lock()
	defer node.peerLock.Unlock()

	node.Peers[peer.ID] = peer
	fmt.Printf("[%s] Discovered new peer: %s\n", node.ID, peer.ID)
}

func (node *BlockchainNode) removePeer(peerID string) {
	node.peerLock.Lock()
	defer node.peerLock.Unlock()

	delete(node.Peers, peerID)
	fmt.Printf("[%s] Disconnected from peer: %s\n", node.ID, peerID)
}

// selectRandomPeers selects a random subset of peers from the given list.
func (node *BlockchainNode) selectRandomPeers(peers []model.Peer, count int) []model.Peer {
	selectedPeers := make([]model.Peer, 0, count)

	node.peerLock.Lock()
	defer node.peerLock.Unlock()

	// Randomly shuffle the list of peers
	shuffledPeers := make([]model.Peer, len(peers))
	copy(shuffledPeers, peers)
	rand.Shuffle(len(shuffledPeers), func(i, j int) {
		shuffledPeers[i], shuffledPeers[j] = shuffledPeers[j], shuffledPeers[i]
	})

	// Select the first 'count' peers
	for i := 0; i < count && i < len(shuffledPeers); i++ {
		selectedPeers = append(selectedPeers, shuffledPeers[i])
	}

	return selectedPeers
}

func (node *BlockchainNode) sendKnownHosts(selectedPeer model.Peer) {
	err := node.peerConnectionPoll.SendDiscoveredPeers(node.ID, selectedPeer, getMapValues(node.Peers))
	if err != nil {
		node.removePeer(selectedPeer.ID)
	}
}
func getMapValues(m map[string]model.Peer) []model.Peer {
	values := make([]model.Peer, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}
