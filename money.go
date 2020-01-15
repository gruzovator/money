package money

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Money is amount of money as integer number of minor units and currency code.
// E.g.: one dollar is Money{100, "USD"} cause "USD" currency has scale 2.
type Money struct {
	Amount   int64
	Currency string
}

// Make creates Money by normalizing amount/scale according to currencies system.
// E.g. Money(100, 0, "USD") gives Money{100_00, "USD"}.
// Rounding is done via math.Round.
func Make(amount int64, scale int, currencyCode string) Money {
	targetScale := ScaleForCurrency(currencyCode)
	if targetScale == scale {
		return Money{
			Amount:   amount,
			Currency: currencyCode,
		}
	}
	amount = int64(math.Round(float64(amount) * math.Pow10(targetScale-scale)))
	return Money{
		Amount:   amount,
		Currency: currencyCode,
	}
}

// FromFloat create Money instance rounding float value according to currency scale.
func FromFloat(val float64, currencyCode string) Money {
	scale := ScaleForCurrency(currencyCode)
	amount := int64(math.Round(val * math.Pow10(scale)))
	return Money{
		Amount:   amount,
		Currency: currencyCode,
	}
}

// String creates string representation, e.g. "100.00 USD" or "100 KRW" (currency with scale 0).
func (m Money) String() string {
	var fmtStr string
	scale := ScaleForCurrency(m.Currency)
	if scale > 0 {
		fmtStr = "%." + strconv.Itoa(scale) + "f"
	} else {
		fmtStr = "%.0f"
	}
	s := fmt.Sprintf(fmtStr, float64(m.Amount)/math.Pow10(scale))
	if m.Currency == "" {
		return s
	}
	return s + " " + m.Currency
}

// FromString parses Money from string that has two parts: float value and currency code.
func FromString(s string) (Money, error) {
	parts := strings.SplitN(s, " ", 2)
	val, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return Money{}, fmt.Errorf("parsing amount: %s", err)
	}
	var currency string
	if len(parts) > 1 {
		currency = parts[1]
	}
	return FromFloat(val, currency), nil
}
