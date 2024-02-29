package collection

import "taheri24.ir/blockchain/wallet"

type Collection struct {
	walletMap map[wallet.ID]wallet.Wallet
}

func (c *Collection) BeginCreate(publicKey []byte) wallet.ID {

	return wallet.ID("")
}
