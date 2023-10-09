package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	Hash         string
}

type Blockchain struct {
	Blocks []*Block
}

func NewBlock(transaction string, nonce int, previousHash string) *Block {
	block := &Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
		Hash:         CalculateHash(fmt.Sprintf("%s%d%s", transaction, nonce, previousHash)),
	}
	return block
}

func DisplayBlocks(bc *Blockchain) {
	for _, block := range bc.Blocks {
		fmt.Printf("Transaction: %s, Nonce: %d, Previous Hash: %s, Current Hash: %s\n", block.Transaction, block.Nonce, block.PreviousHash, block.Hash)
	}
}

func ChangeBlock(block *Block, newTransaction string) {
	block.Transaction = newTransaction
	block.Hash = CalculateHash(fmt.Sprintf("%s%d%s", block.Transaction, block.Nonce, block.PreviousHash))
}

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

func CalculateHash(stringToHash string) string {
	hashInBytes := sha256.Sum256([]byte(stringToHash))
	return hex.EncodeToString(hashInBytes[:])
}

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

	// Adding 5 dummy blocks
	for i := 0; i < 5; i++ {
		transaction := fmt.Sprintf("user%d to user%d", i, i+1)
		nonce := i
		previousHash := blockchain.Blocks[len(blockchain.Blocks)-1].Hash
		newBlock := NewBlock(transaction, nonce, previousHash)
		blockchain.Blocks = append(blockchain.Blocks, newBlock)
	}

	// Display all blocks
	DisplayBlocks(blockchain)

	// Verify the blockchain
	isValid1 := VerifyChain(blockchain)
	if isValid1 {
		fmt.Println("Blockchain is valid.")
	} else {
		fmt.Println("Blockchain is not valid.")
	}

	fmt.Printf("\n\n\n")

	// Change the transaction of the third block
	ChangeBlock(blockchain.Blocks[3], "user3 to user5")

	// Display all blocks
	DisplayBlocks(blockchain)
	// Verify the blockchain
	isValid2 := VerifyChain(blockchain)
	if isValid2 {
		fmt.Println("Blockchain is valid.")
	} else {
		fmt.Println("Blockchain is not valid.")
	}
}
