package wallet

import (
	"crypto/rand"

	"taheri24.ir/blockchain/wallet/ciphers"
)

type OptionType int

const (
	optGenerateKeys OptionType = iota
	optID
)

type Option struct {
	optType OptionType
	keySize int
	id      ID
}

func applyOptions(w *Wallet, opts []Option) {
	for _, opt := range opts {
		switch opt.optType {
		case optGenerateKeys:
			w.PrivateKey, w.PublicKey = ciphers.GenerateKeyPair(rand.Reader, int(opt.keySize))
		case optID:
			w.ID = opt.id
		}
	}
}

func GenerateKeys(keySize int) Option {
	return Option{optType: optGenerateKeys, keySize: keySize}
}

func WalletID(id ID) Option {
	return Option{optType: optGenerateKeys, id: id}
}
