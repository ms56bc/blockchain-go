package server

import (
	"github.com/ms56bc/blockchain-go/server/model"
	"math/rand"
	"strings"
)

// Miner interface represents the mining process
type Miner interface {
	Mine(block *model.Block)
}

// ProofOfWorkMiner implements the Miner interface with a Proof of Work algorithm
type ProofOfWorkMiner struct {
	difficulty int
}

func (m *ProofOfWorkMiner) Mine(block *model.Block) {
	// complete this function
	m.proofOfWork(block)
}

// ProofOfWork performs the proof-of-work algorithm

func (m *ProofOfWorkMiner) proofOfWork(block *model.Block) {
	for {
		block.Nonce = rand.Uint64() // Random nonce for each attempt
		hash := block.CalculateHash()

		// Check if the hash meets the difficulty criteria
		if hash[:m.difficulty] == strings.Repeat("0", m.difficulty) {
			block.Hash = hash
			return
		}
	}
}
