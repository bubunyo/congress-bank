package congressbank

import (
	"fmt"
	"testing"

	"github.com/bubunyo/congress-bank/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccount_BalanceCalculation(t *testing.T) {
	user, ok := CreateNewUser("Bubu")
	assert.True(t, ok)

	acc1, ok := user.NewAccount(GhanaPesewas)
	require.True(t, ok)
	assert.Equal(t, int64(0), acc1.Balance())

	m := &MoneyOrder{
		Id:        fmt.Sprintf("l-%d", genId()),
		Amount:    0,
		Status:    Pending,
		LedgerIds: *store.NewStore[struct{}](),
	}

	createLedgerItem(m, 10, Credit)
	createLedgerItem(m, 5, Debit)
	createLedgerItem(m, 6, Debit)

	moneyOrders.Insert(m.Id, m)
	acc1.MoneyOrderIds.Insert(m.Id, struct{}{})

	assert.Equal(t, int64(-1), acc1.Balance())
}

func createLedgerItem(m *MoneyOrder, amount uint64, t TransactionType) {
	l := &LedgerItem{Id: genIdWithPrefix("l-"), Amount: amount, TransactionType: t}
	m.LedgerIds.Insert(l.Id, struct{}{})
	ledger.Insert(l.Id, l)
}
