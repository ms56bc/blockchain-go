package server

import (
	"github.com/ms56bc/blockchain-go/server/model"
)

type BlockchainServer interface {
	AddBlock(transactions []model.Transaction)
	LatestBlock() model.Block
}
type BlockchainProofOfWork struct {
	miner        Miner
	genesisBlock model.Block
	chain        []model.Block
}

func NewBlockchain(miner Miner, genesisBlock model.Block) *BlockchainProofOfWork {
	chain := []model.Block{genesisBlock}
	return &BlockchainProofOfWork{miner, genesisBlock, chain}
}

func (bc *BlockchainProofOfWork) AddBlock(transactions []model.Transaction) {
	var previousHash string
	if len(bc.chain) > 0 {
		previousBlock := bc.chain[len(bc.chain)-1]
		previousHash = previousBlock.Hash
	}
	block := model.NewBlock(transactions, previousHash, len(bc.chain))
	// Proof-of-work (mining)
	bc.miner.Mine(block)

	bc.chain = append(bc.chain, *block)
}

// LatestBlock returns the latest block in the blockchain
func (bc *BlockchainProofOfWork) LatestBlock() model.Block {
	return bc.chain[len(bc.chain)-1]
}

// Validate validates the entire blockchain
func (bc *BlockchainProofOfWork) validate() bool {
	for i := 1; i < len(bc.chain); i++ {
		currentBlock := bc.chain[i]
		previousBlock := bc.chain[i-1]
		if currentBlock.Hash != currentBlock.CalculateHash() {
			return false
		}
		if currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}
	}
	return true
}
