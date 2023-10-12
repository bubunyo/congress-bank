package congressbank

import (
	"math"

	"github.com/bubunyo/congress-bank/store"
)

var (
	accountBase = store.NewStore[Account]()
)

var house = map[CurrencyCode]*Account{}

func init() {
	house[GhanaPesewas] = &Account{
		Num:           genIdWithPrefix("A-"),
		Curr:          GhanaPesewas,
		MoneyOrderIds: store.NewStore[struct{}](),
		Type:          "current",
	}
}

type Account struct {
	Num           string
	Curr          CurrencyCode
	MoneyOrderIds *store.Store[struct{}]
	Type          string // savings/current
}

func (a *Account) Balance() int64 {
	if a.Num == house[a.Curr].Num {
		return math.MaxInt64
	}
	var acc int64
	moneyOrders.Range(func(mid string) {
		m, _ := moneyOrders.Get(mid)
		m.LedgerIds.Range(func(id string) {
			li, ok := ledger.Get(id)
			if !ok {
				return
			}
			if li.AccountId != a.Num {
				return
			}
			if ok {
				if li.TransactionType == Credit {
					acc += int64(li.Amount)
				} else {
					acc -= int64(li.Amount)
				}
			}
		})
	})
	return acc
}

func SimulateAMLChecks(src, des *Account) (bool, string) {
	// return false, "simulated aml check failure"
	return true, ""
}
