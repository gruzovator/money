package money

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

type currenciesSystem struct {
	scaleByCode  map[string]int
	defaultScale int
}

var currencies currenciesSystem

func init() {
	_, thisFilePath, _, _ := runtime.Caller(0)
	iso4217DataFile := filepath.Join(filepath.Dir(thisFilePath), "iso-4217.xml")

	f, err := os.Open(iso4217DataFile)
	if err != nil {
		panic(fmt.Sprintf("open currencies scales data file %s: %s", iso4217DataFile, err))
	}
	defer f.Close()

	scales, err := CurrencyScalesFromISO4217(f)
	if err != nil {
		panic(fmt.Sprintf("reading currencies scales data file %s: %s", iso4217DataFile, err))
	}

	currencies = currenciesSystem{
		scaleByCode:  scales,
		defaultScale: 2,
	}
}

// DefaultScale is scale value used for unknown currencies.
func DefaultScale() int {
	return currencies.defaultScale
}

// IsKnownCurrency checks if currency presents in ISO4217 data.
func IsKnownCurrency(curecnyCode string) bool {
	_, ok := currencies.scaleByCode[curecnyCode]
	return ok
}

// ScaleForCurrency gives scale value for currency code. Default scale is used for unknown currency.
func ScaleForCurrency(currencyCode string) int {
	if s, ok := currencies.scaleByCode[currencyCode]; ok {
		return s
	}

	return currencies.defaultScale
}

// ForEachCurrency is for currencies data iteration.
func ForEachCurrency(cb func(code string, scale int)) {
	for c, s := range currencies.scaleByCode {
		cb(c, s)
	}
}

// ReplaceCurrenciesSystem should be used to replace package currencies data.
func ReplaceCurrenciesSystem(scaleByCode map[string]int, defaultScale int) {
	currencies.scaleByCode = scaleByCode
	currencies.defaultScale = defaultScale
}
