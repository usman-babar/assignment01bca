package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// Create a new blockchain
type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	Hash         string
}

// Adding 5 dummy blocks
func NewBlock(transaction string, nonce int, previousHash string) *Block {
	block := &Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
		Hash:         CalculateHash(fmt.Sprintf("%s%d%s", transaction, nonce, previousHash)),
	}
	return block
}

// Create a new blockchain
type Blockchain struct {
	Blocks []*Block
}

// Calculate hash
func CalculateHash(stringToHash string) string {
	hashInBytes := sha256.Sum256([]byte(stringToHash))
	return hex.EncodeToString(hashInBytes[:])
}

// Verify the blockchain
func ChangeBlock(block *Block, newTransaction string) {
	block.Transaction = newTransaction
	block.Hash = CalculateHash(fmt.Sprintf("%s%d%s", block.Transaction, block.Nonce, block.PreviousHash))
}

// Display all blocks
func DisplayBlocks(bc *Blockchain) {
	for _, block := range bc.Blocks {
		fmt.Printf("Transaction: %s, Nonce: %d, Previous Hash: %s, Current Hash: %s\n", block.Transaction, block.Nonce, block.PreviousHash, block.Hash)
	}
}

// Verify the blockchain
func VerifyChain(bc *Blockchain) bool {
	for i := 1; i < len(bc.Blocks); i++ {
		currentBlock := bc.Blocks[i]
		previousBlock := bc.Blocks[i-1]

		if currentBlock.Hash != CalculateHash(fmt.Sprintf("%s%d%s", currentBlock.Transaction, currentBlock.Nonce, currentBlock.PreviousHash)) {
			return false
		}

		if currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}
	}
	return true
}

// main
func main() {
	
	// Create a new blockchain
	blockchain := &Blockchain{
		Blocks: []*Block{
			{
				Transaction:  "genesis",
				Nonce:        0,
				PreviousHash: "",
				Hash:         "",
			},
		},
	}

	// Adding dummy blocks
	for i := 0; i < 6; i++ {
		transaction := fmt.Sprintf("person%d to person%d", i, i+1)
		nonce := i
		previousHash := blockchain.Blocks[len(blockchain.Blocks)-1].Hash
		newBlock := NewBlock(transaction, nonce, previousHash)
		blockchain.Blocks = append(blockchain.Blocks, newBlock)
	}

	// Displaying blocks
	DisplayBlocks(blockchain)

	// Verifying the blockchain
	isValid1 := VerifyChain(blockchain)
	if isValid1 {
		fmt.Println("\nBlockchain is valid.\n")
	} else {
		fmt.Println("\nBlockchain is invalid\n.")
	}

	fmt.Printf("\n\n\n")

	// Change the transaction of the third block
	ChangeBlock(blockchain.Blocks[4], "person4 to person4")

	// Displaying all blocks
	DisplayBlocks(blockchain)

	// Verifying the blockchain
	isValid2 := VerifyChain(blockchain)
	if isValid2 {
		fmt.Println("\nBlockchain is valid.\n")
	} else {
		fmt.Println("\nBlockchain is invalid.\n")
	}
}
