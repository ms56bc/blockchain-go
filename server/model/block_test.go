package model

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCalculateHash(t *testing.T) {
	// Create a sample block for testing
	block := NewBlock(
		[]Transaction{{Sender: "Alice", Recipient: "Bob", Amount: 1.0}},
		"previous_hash",
		1)

	// Calculate the expected hash using the same logic
	data, _ := json.Marshal(block.data)
	expectedData := block.PreviousHash + string(data) + block.timestamp.String() + strconv.Itoa(block.index) + strconv.FormatUint(block.Nonce, 10)
	expectedHash := sha256.Sum256([]byte(expectedData))
	expectedHashStr := fmt.Sprintf("%x", expectedHash)

	// Calculate the actual hash using the CalculateHash method
	actualHash := block.CalculateHash()

	// Compare the expected and actual hashes
	if actualHash != expectedHashStr {
		t.Errorf("Expected hash: %s, Actual hash: %s", expectedHashStr, actualHash)
	}
}

func TestBlockInitialization(t *testing.T) {
	// Create a sample block for testing
	block := Block{
		data:         []Transaction{{Sender: "Alice", Recipient: "Bob", Amount: 1.0}},
		Hash:         "previous_hash",
		PreviousHash: "",
		timestamp:    time.Now(),
		index:        1,
		Nonce:        123,
	}

	// Check if the block fields are initialized correctly
	if block.data[0].Sender != "Alice" {
		t.Error("Sender not initialized correctly")
	}
	if block.Hash != "previous_hash" {
		t.Error("Hash not initialized correctly")
	}
	// Add similar checks for other fields as needed
}
