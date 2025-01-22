package blockchain_test

import (
	"Blockchain/blockchain"
	"testing"
	"time"
)

func TestNewBlockchain(t *testing.T) {
	cfg := &blockchain.Config{Difficulty: 2}
	bc := blockchain.NewBlockchain(cfg)

	if len(bc.Chain) != 1 {
		t.Errorf("should have one block in the chain")
	}
}

func TestAddBlock(t *testing.T) {
	cfg := &blockchain.Config{Difficulty: 2}
	bc := blockchain.NewBlockchain(cfg)

	newBlock := &blockchain.Block{
		Index:        1,
		Timestamp:    time.Now().Format(time.RFC3339),
		PrevHash:     bc.Chain[0].Hash,
		MinerAddress: "tester",
	}

	pow := blockchain.NewProofOfWork(newBlock, cfg)
	nonce, hash := pow.Mine()
	newBlock.Nonce = nonce
	newBlock.Hash = hash

	bc.AddBlock(newBlock)

	if len(bc.Chain) != 2 {
		t.Errorf("should have two blocks in the chain")
	}
}

func TestPOWValidation(t *testing.T) {
	cfg := &blockchain.Config{Difficulty: 2}
	block := &blockchain.Block{
		Index:        1,
		Timestamp:    time.Now().Format(time.RFC3339),
		PrevHash:     "0000test",
	}

	pow := blockchain.NewProofOfWork(block, cfg)
	nonce, hash := pow.Mine()
	block.Nonce = nonce
	block.Hash = hash

	if !pow.Validate() {
		t.Errorf("block should be valid")
	}
}
