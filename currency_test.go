package congressbank

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCurrencyAssertBehaviour(t *testing.T) {
	amount := 100.00
	o, err := GhanaPesewas.ConvertTo(GhanaCedis, amount)
	require.NoError(t, err)
	o, err = GhanaCedis.ConvertTo(GhanaPesewas, o)
	require.NoError(t, err)
	assert.Equal(t, amount, o)
}

func TestCurrencyGhstoGHp(t *testing.T) {
	for i, tt := range []struct{ in, out float64 }{
		{1, 100},
		{45389477363, 4538947736300},
		{453.2, 45320},
		{0.00001, 0.001},
	} {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			out, err := GhanaCedis.ConvertTo(GhanaPesewas, tt.in)
			require.NoError(t, err)
			assert.Equal(t, tt.out, out)
		})
	}
}

func TestCurrencyGhptoGHS(t *testing.T) {
	for i, tt := range []struct{ in, out float64 }{
		// Pesewas -> Cedis
		{100, 1},
		{4532, 45.32},
		{453.2, 4.532},
	} {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			out, err := GhanaPesewas.ConvertTo(GhanaCedis, tt.in)
			require.NoError(t, err)
			assert.Equal(t, tt.out, out)
		})
	}
}

func TestCurrency_SameCurrency(t *testing.T) {
	amount := 100.00
	out, err := GhanaPesewas.ConvertTo(GhanaPesewas, amount)
	require.NoError(t, err)
	assert.Equal(t, amount, out)
}
