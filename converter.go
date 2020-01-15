package money

import (
	"errors"
	"fmt"
	"math"
)

// Converter is a currency converter
type Converter interface {
	Convert(money Money, toCurrency string) (Money, error)
}

type RatesTableConverter map[string]float64

func (c RatesTableConverter) Convert(m Money, toCurrency string) (Money, error) {
	if m.Currency == toCurrency {
		return m, nil
	}
	f1, ok := c[m.Currency]
	if !ok {
		return Money{}, errors.New("currency " + m.Currency + ": no conversion rate")
	}
	if f1 < math.SmallestNonzeroFloat64 {
		return Money{}, fmt.Errorf("currency %s: bad conversion rate: %.2f", m.Currency, f1)
	}
	f2, ok := c[toCurrency]
	if !ok {
		return Money{}, errors.New("currency " + toCurrency + ": no conversion rate")
	}
	if f2 < math.SmallestNonzeroFloat64 {
		return Money{}, fmt.Errorf("currency %s: bad conversion rate: %.2f", toCurrency, f2)
	}
	v := float64(m.Amount) / math.Pow10(ScaleForCurrency(m.Currency))
	v *= f2 / f1
	return FromFloat(v, toCurrency), nil
}
