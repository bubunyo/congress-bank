package congressbank

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMoneyOrders(t *testing.T) {
	bubu, ok := CreateNewUser("Bubu")
	assert.True(t, ok)

	kwame, ok := CreateNewUser("kwame")
	assert.True(t, ok)

	bubuAcc, ok := bubu.NewAccount(GhanaPesewas)
	require.True(t, ok)
	assert.Equal(t, int64(0), bubuAcc.Balance())

	kwameAcc, ok := kwame.NewAccount(GhanaPesewas)
	require.True(t, ok)
	assert.Equal(t, int64(0), kwameAcc.Balance())

	err := CreateMoneyOrder(*house[bubuAcc.Curr], *bubuAcc, 10_000, GhanaPesewas, "Shege reasons")
	assert.NoError(t, err)

	bal := bubuAcc.Balance()
	balStr, err := bubuAcc.Curr.ConvertTo(GhanaCedis, float64(bal))
	require.NoError(t, err)
	fmt.Println("bubu's balance=", balStr)
}
