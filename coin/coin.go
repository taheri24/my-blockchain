package coin

import "time"

type Step int

const (
	VerifyOnly Step = iota
	Transfer        // to new owner
	Arrange
)

type MinerID string
type ID int64

type State *struct {
	LastVerify time.Time
}

type Coin struct {
	Key         ID
	Codes       map[MinerID]string
	Sig         string
	OwnerWallet string
	Expire      int64
	Verified    bool
}
