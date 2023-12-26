package model

type Transaction struct {
	Sender    string
	Recipient string
	Amount    float64
	Signature string
}

// Wallet represents a wallet for managing transactions
type Wallet struct {
	PublicKey  string
	PrivateKey string
}

// CreateTransaction creates a new transaction
func (w *Wallet) CreateTransaction(recipient string, amount float64) Transaction {
	transaction := Transaction{
		Sender:    w.PublicKey,
		Recipient: recipient,
		Amount:    amount,
		Signature: "", // Placeholder for digital signature
	}
	// TODO: Sign the transaction with the private key
	return transaction
}
