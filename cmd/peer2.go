package main

import (
	"context"
	"github.com/ms56bc/blockchain-go/pkg/server/config"
	"github.com/ms56bc/blockchain-go/pkg/server/model"
	"github.com/ms56bc/blockchain-go/pkg/server/pd"
)

func main() {
	peers := make([]model.Peer, 1)
	peers[0] = model.Peer{ID: "1", IP: "127.0.0.1:8091"}
	config := config.Config{
		GossipAddress: "127.0.0.1:8093",
		StaticPeers:   peers,
	}
	node := pd.NewBlockchainNode("3", config)
	discoveryServer := pd.NewPeerDiscoveryServer(node, config)
	ctx := context.Background()
	go node.DiscoverPeers(ctx)
	discoveryServer.StartGossipUDPServer(ctx)
}
