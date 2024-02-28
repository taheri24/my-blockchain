package daemon

import (
	"log/slog"
	"sync"

	"taheri24.ir/blockchain/coin"
)

type Daemon struct {
	CoinStates struct {
		data     map[coin.ID]coin.State
		syncLock sync.Mutex
	}

	ch chan coin.MinerOp
}

func (d *Daemon) Verify(c *coin.Coin) {
	if c.Verified {
		slog.Error("Verify error", c)
		panic("already verified")
	}

	c.Verified = true
	d.MinerOps.syncLock.Lock()
	defer d.MinerOps.syncLock.Unlock()

	for minerID, code := range c.Codes {
		d.MinerOps.data[minerID] = append(d.MinerOps.data[minerID], code)
	}
}

func (d *Daemon) Transfer(c *coin.Coin) {

}

func (d *Daemon) Process() {

}
