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
		LedgerIds: *store.NewStore[struct{}](),
	}

	moneyOrders.Insert(m.Id, m)
	src.MoneyOrderIds.Insert(m.Id, struct{}{})

	h, ok := house[curr]
	if !ok {
		return fmt.Errorf("house currency does not exist, currency=%s", curr)
	}
	/// job 1

	// check balance of source accoutn
	{
		// job 1
		balance := uint64(src.Balance())
		if amount > balance {
			return fmt.Errorf("order amount is more than account src blance, balance=%d", balance)
		}

		///debit src

		/// this should be in a database transaction
		{
			RunCreditDebit(creditDebit{
				mId:     m.Id,
				srcAcc:  &src,
				destAcc: h,
				amount:  amount,
				curr:    curr,
			})
		}
	}

	{
		// job 2
		// triggered by job 1, triggers job 3
		// do you aml check
		amlCheckSuccess, reason := SimulateAMLChecks(&src, &dest)
		if !amlCheckSuccess {
			RunCreditDebit(creditDebit{
				mId:     m.Id,
				destAcc: &src,
				srcAcc:  h,
				amount:  amount,
				curr:    curr,
			})
			return fmt.Errorf("aml check fails, reason=%s", reason)
		}
	}

	{
		// job 3
		// triggered by job 2
		RunCreditDebit(creditDebit{
			mId:     m.Id,
			srcAcc:  h,
			destAcc: &dest,
			amount:  amount,
			curr:    curr,
		})
	}

	return nil
}

type creditDebit struct {
	mId     string
	srcAcc  *Account
	destAcc *Account
	amount  uint64
	curr    CurrencyCode
}

func RunCreditDebit(c creditDebit) {
	// this is a transaction
	ledgerItem1 := &LedgerItem{
		Id:              genIdWithPrefix("L-"),
		TransactionId:   c.mId,
		AccountId:       c.srcAcc.Num,
		TransactionType: Debit,
		Amount:          c.amount,
	}
	ledger.Insert(ledgerItem1.Id, ledgerItem1)

	ledgerItem2 := &LedgerItem{
		Id:              genIdWithPrefix("L-"),
		TransactionId:   c.mId,
		AccountId:       c.destAcc.Num,
		TransactionType: Credit,
		Amount:          c.amount,
	}
	ledger.Insert(ledgerItem2.Id, ledgerItem2)
	m, _ := moneyOrders.Get(c.mId)

	m.LedgerIds.Insert(ledgerItem1.Id, struct{}{})
	m.LedgerIds.Insert(ledgerItem2.Id, struct{}{})
	moneyOrders.Insert(m.Id, m)
}
