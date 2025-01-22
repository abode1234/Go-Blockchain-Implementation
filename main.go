package main

import (
	"Blockchain/blockchain"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var bc *blockchain.Blockchain
var config *blockchain.Config

func main() {
	cfg, err := blockchain.LoadConfig("config.json")
	if err != nil {
		panic(err)
	}
	config = cfg

	bc = blockchain.NewBlockchain(config)

	http.HandleFunc("/chain", getChain)
	http.HandleFunc("/mine", mineBlock)
	fmt.Printf("Listening on %s\n", config.Port)
	http.ListenAndServe(config.Port, nil)
}

func getChain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bc.Chain)
}

func mineBlock(w http.ResponseWriter, r *http.Request) {
	lastBlock := bc.Chain[len(bc.Chain)-1]
	newBlock := &blockchain.Block{
		Index:        lastBlock.Index + 1,
		Timestamp:    time.Now().Format(time.RFC3339),
		PrevHash:     lastBlock.Hash,
		MinerAddress: "test_miner",
	}

	pow := blockchain.NewProofOfWork(newBlock, config)
	nonce, hash := pow.Mine()
	newBlock.Nonce = nonce
	newBlock.Hash = hash

	bc.AddBlock(newBlock)
	bc.SaveToFile("blockchain.json")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newBlock)
}
