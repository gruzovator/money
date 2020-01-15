package money_test

import (
	"testing"

	"github.com/gruzovator/money"
	"github.com/stretchr/testify/require"
)

func TestForEachCurrency(t *testing.T) {
	scales := make(map[string]int)
	money.ForEachCurrency(func(code string, scale int) {
		scales[code] = scale
	})

	require.Equal(t, 2, scales["USD"])
	require.Equal(t, 2, scales["GBP"])
	require.Equal(t, 2, scales["RUB"])
	require.Greater(t, len(scales), 100)
}
