package money_test

import (
	"fmt"
	"testing"

	"github.com/gruzovator/money"
	"github.com/stretchr/testify/require"
)

func TestRatesTableConverter_Convert(t *testing.T) {
	converter := money.RatesTableConverter{
		"RUB": 60,
		"USD": 1,
		"GBP": 0.8,
	}
	cases := []struct {
		from money.Money
		to   string
		want money.Money
	}{
		{
			from: money.FromFloat(1, "RUB"),
			to:   "RUB",
			want: money.FromFloat(1, "RUB"),
		},
		{
			from: money.FromFloat(60, "RUB"),
			to:   "USD",
			want: money.FromFloat(1, "USD"),
		},
		{
			from: money.FromFloat(0.1, "USD"),
			to:   "RUB",
			want: money.FromFloat(6, "RUB"),
		},
		{
			from: money.FromFloat(60, "RUB"),
			to:   "GBP",
			want: money.FromFloat(0.8, "GBP"),
		},
		{
			from: money.FromFloat(1, "GBP"),
			to:   "RUB",
			want: money.FromFloat(75, "RUB"),
		},
		{
			from: money.FromFloat(0, "USD"),
			to:   "RUB",
			want: money.FromFloat(0, "RUB"),
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("converting %s to %s", c.from, c.to), func(t *testing.T) {
			m, err := converter.Convert(c.from, c.to)
			require.NoError(t, err)
			require.Equal(t, c.want, m)
		})
	}
}
