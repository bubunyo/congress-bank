package congressbank

import (
	"github.com/bubunyo/congress-bank/store"
)

var (
	accountBase = store.NewStore[Account]()
)

type Account struct {
	Num           string
	Curr          CurrencyCode
	MoneyOrderIds *store.Store[struct{}]
	Type          string // savings/current
}

func (a *Account) Balance() int64 {
	var acc int64
	moneyOrders.Range(func(mid string) {
		m, _ := moneyOrders.Get(mid)
		m.LedgerIds.Range(func(id string) {
			if li, ok := ledger.Get(id); ok {
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
