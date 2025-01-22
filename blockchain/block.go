package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

type Transaction struct {
	Sender			 string				`json: "sender`
	Recipient	     string			    `json: "recipient`
	Amount		     int				`json: "amount`
}

type Block struct {

	Index        int           `json:"index"`
	Timestamp    string        `json:"timestamp"`
	Transactions []Transaction `json:"transactions"`
	PrevHash     string        `json:"prev_hash"`
	Hash         string        `json:"hash"`
	Nonce        int           `json:"nonce"`
	MinerAddress string        `json:"miner_address"`

}

func (b *Block) CalculateHash() string {

	data, err := json.Marshal(struct {

		Index        int           `json:"index"`
		Timestamp    string        `json:"timestamp"`
		data		 []Transaction `json:"transactions"`
		PrevHash     string        `json:"prev_hash"`
		Nonce        int           `json:"nonce"`
	}{
		b.Index,
		b.Timestamp,
		b.Transactions ,
		b.PrevHash,
		b.Nonce,
	})
	if err != nil {
		return err.Error()
	}
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])

}
