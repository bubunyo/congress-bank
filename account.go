package congressbank

import "github.com/bubunyo/congress-bank/store"

var (
	accountBase = store.NewStore[Account]()
)

type Currency struct {
}

type Transaction struct {
}

type Ledger struct {
}

type Account struct {
	Num          string
	Curr         Currency
	Transactions *store.Store[Transaction]
	Ledger       *store.Store[Ledger]
	Type         string // savings/current
}
