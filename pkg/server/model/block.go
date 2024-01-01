package model

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	data         []Transaction
	Hash         string
	PreviousHash string
	timestamp    time.Time
	index        int
	Nonce        uint64
}

func (b *Block) CalculateHash() string {
	data, _ := json.Marshal(b.data)
	blockData := b.PreviousHash + string(data) + b.timestamp.String() + strconv.Itoa(b.index) + strconv.FormatUint(b.Nonce, 10)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}
func NewBlock(data []Transaction, previousHash string, index int) *Block {
	block := &Block{
		data:         data,
		PreviousHash: previousHash,
		timestamp:    time.Now(),
		index:        index,
		Nonce:        0,
	}
	return block
}
