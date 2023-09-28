package congressbank

import (
	"fmt"

	"github.com/bubunyo/congress-bank/store"
)

type (
	TransactionType string
	OrderStatus     string
)

var (
	ledger      = store.NewStore[*LedgerItem]()
	moneyOrders = store.NewStore[*MoneyOrder]()
)

const (
	Debit  TransactionType = "debit"
	Credit TransactionType = "credit"

	Pending OrderStatus = "pending"
	Success OrderStatus = "success"
	Failed  OrderStatus = "failed"
)

type LedgerItem struct {
	Id              string
	TransactionId   string
	AccountId       string
	TransactionType TransactionType
	Amount          uint64
}

type MoneyOrder struct {
	Id        string
	Reference string
	Amount    uint64
	LedgerIds store.Store[struct{}]
	SrcAcc    string
	DestAcc   string
	Status    OrderStatus
}

func CreateMoneyOrder(src, dest Account, amount uint64, curr CurrencyCode, ref string) error {
	// this should be wrapped inn a trascation

	// trans 1
	// debit src
	// credit house
	// trans 2

	// job for aml checke -- automatic check
	// -- triger from house to destination - manual

	// job
	// aml check goes here
	// debit house
	// credit the dest
	if src.Curr != dest.Curr {
		return fmt.Errorf("Invalid money order currencies")
	}

	m := &MoneyOrder{
		Id:        fmt.Sprintf("l-%d", genId()),
		Reference: ref,
		Amount:    amount,
		SrcAcc:    src.Num,
		DestAcc:   dest.Num,
		Status:    Pending,
	}

	moneyOrders.Insert(m.Id, m)
	src.MoneyOrderIds.Insert(m.Id, struct{}{})
	// dest.MoneyOrders.Insert(m.Id, *m)

	// if src.Balance() < amount {
	// 	m.Status = Failed
	// 	return fmt.Errorf("no enough source balance")
	// }

	// _ := LedgerItem{
	// 	Id: fmt.Sprintf("l-%d", genId()),
	// }

	return nil
}
