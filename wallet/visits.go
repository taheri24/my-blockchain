package wallet

import (
	"crypto/rand"

	"taheri24.ir/blockchain/wallet/ciphers"
)

type VisitFunc func(*Wallet)

func GenerateKeys(keySize KeySize) VisitFunc {
	return func(w *Wallet) {
		w.PrivateKey, w.PublicKey = ciphers.GenerateKeyPair(rand.Reader, int(keySize))

	}
}
func WalletID(id ID) VisitFunc {
	return func(w *Wallet) {
		w.ID = id
	}
}
