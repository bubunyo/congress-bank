package congressbank

import (
	"fmt"
	"math"
)

const (
	GhanaCedis   CurrencyCode = "GHS"
	GhanaPesewas CurrencyCode = "GHp"
)

var (
	ConversionTable = map[CurrencyCode]map[CurrencyCode]int{
		GhanaCedis:   {GhanaPesewas: 2},
		GhanaPesewas: {GhanaCedis: -2},
	}
)

func (c CurrencyCode) ConvertTo(code CurrencyCode, amount float64) (float64, error) {
	if c == code {
		return amount, nil
	}
	conversion, ok := ConversionTable[c]
	if !ok {
		return 0, fmt.Errorf("invalid conversion from %s to %s", c, code)
	}
	f, ok := conversion[code]
	if !ok {
		return 0, fmt.Errorf("invalid conversion from %s to %s", c, code)
	}

	return amount * math.Pow(10, float64(f)), nil
}

type CurrencyCode string

func (c CurrencyCode) String(amount float64) string {
	return fmt.Sprintf("%s %.2f", c, amount)
}
