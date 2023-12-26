package server

import (
	"github.com/ms56bc/blockchain-go/server/model"
	"strings"
	"testing"
)

func TestCalculateHash(t *testing.T) {
	// Create a sample block for testing
	miner := ProofOfWorkMiner{difficulty: 1}
	block := model.NewBlock(
		[]model.Transaction{{Sender: "Alice", Recipient: "Bob", Amount: 1.0}},
		"previous_hash",
		1)

	miner.Mine(block)

	// Compare the expected and actual hashes
	if block.Hash[:miner.difficulty] != strings.Repeat("0", miner.difficulty) {
		t.Errorf("Expected hash: %s, Actual hash: %s", strings.Repeat("0", miner.difficulty), block.Hash[:miner.difficulty])
	}
}
