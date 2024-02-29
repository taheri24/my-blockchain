package wallet

import (
	"crypto/rsa"

	"taheri24.ir/blockchain/wallet/miner"
)

type (
	ID        string
	SecretKey string
)

type Wallet struct {
	ID         ID
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
	SecretKeys map[miner.ID]SecretKey
}

func New(opts ...Option) *Wallet {
	w := new(Wallet)
	applyOptions(w, opts)

	return w
}
