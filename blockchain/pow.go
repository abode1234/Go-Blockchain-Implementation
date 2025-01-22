package blockchain

import "strings"

type ProofOfWork struct {
	Block  *Block
	Config *Config
}

func NewProofOfWork(b *Block, config *Config) *ProofOfWork {
	return &ProofOfWork{b, config}
}

func (pow *ProofOfWork) Validate() bool {
	hash := pow.Block.CalculateHash()
	return strings.HasPrefix(hash, strings.Repeat("0", pow.Config.Difficulty))
}

func (pow *ProofOfWork) Mine() (int, string) {
	var hash string
	var nonce int

	targetPrefix := strings.Repeat("0", pow.Config.Difficulty)

	for nonce = 0; ; nonce++ {
		pow.Block.Nonce = nonce
		hash = pow.Block.CalculateHash()
		if strings.HasPrefix(hash, targetPrefix) {
			break
		}
	}

	return nonce, hash
}
