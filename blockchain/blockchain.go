package blockchain

import (
	"encoding/json"
	"os"
	"time"
)


type Blockchain struct {

	Chain		[]*Block   `json:"chain"`
	Config		*Config		`json:"-"`
}

func NewBlockchain(c *Config) *Blockchain {
	genesisBlock := &Block{
		Index: 0,
		Timestamp: time.Now().Format(time.RFC3339),
		Transactions: []Transaction{{Sender: "system", Recipient: "genesis", Amount: 0}},
		PrevHash: "0",
		MinerAddress: "system",
	}
	genesisBlock.Hash = genesisBlock.CalculateHash()

	return &Blockchain{
		Chain: []*Block{genesisBlock},
		Config: c,
	}

}

func (bc *Blockchain) AddBlock(newBlock *Block) {
	bc.Chain = append(bc.Chain, newBlock)
}

func (bc *Blockchain) SaveToFile(filename string) error {
	data, err := json.MarshalIndent(bc.Chain, "\n","   ")

	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}


func LoadFromFile(filename string, config *Config) (*Blockchain, error) {


	data, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	var block []*Block

	if err :=json.Unmarshal(data, &block); err != nil {
		return nil, err
	}
	return &Blockchain{
		Chain: block,
		Config: config,
	}, nil

}

