package main

import (
	"context"
	"github.com/ms56bc/blockchain-go/pkg/server/config"
	"github.com/ms56bc/blockchain-go/pkg/server/model"
	"github.com/ms56bc/blockchain-go/pkg/server/pd"
)

func main() {
	cnf := config.Config{
		GossipAddress: "127.0.0.1:8091",
		StaticPeers:   []model.Peer{},
	}
	node := pd.NewBlockchainNode("1", cnf)
	discoveryServer := pd.NewPeerDiscoveryServer(node, cnf)
	ctx := context.Background()
	node.DiscoverPeers(ctx)
	discoveryServer.StartGossipUDPServer(ctx)
}
