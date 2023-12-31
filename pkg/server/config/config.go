package config

import (
	"github.com/ms56bc/blockchain-go/pkg/server/model"
)

type Config struct {
	StaticPeers   []model.Peer
	GossipAddress string
}

func NewConfig() *Config {
	return &Config{}
}
