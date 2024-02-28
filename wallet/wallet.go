package wallet

import (
	"crypto/rsa"

	"taheri24.ir/blockchain/wallet/miner"
)

type (
	ID        string
	SecretKey string
	KeySize   int
)

const (
	ks1024, ks2048, ks4096 KeySize = 1024, 2048, 4096
)

func KeySizeFromString(keySize string) KeySize {
	switch keySize {
	case "1024":
		return ks1024
	case "2048":
		return ks2048
	case "4096":
		return ks4096
	default:
		panic("keySize is not valid")
	}
}

func KeySizeFromInt(keySize int) KeySize {
	switch keySize {
	case 1024:
		return ks1024
	case 2048:
		return ks2048
	case 4096:
		return ks4096
	default:
		panic("keySize is not valid")
	}
}

type Wallet struct {
	ID         ID
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
	SecretKeys map[miner.ID]SecretKey
}

func New(visitFns ...VisitFunc) *Wallet {
	w := new(Wallet)
	for _, visitFn := range visitFns {
		visitFn(w)
	}

	return w
}
