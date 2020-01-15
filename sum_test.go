package money_test

import (
	"fmt"
	"testing"

	"github.com/gruzovator/money"
	"github.com/stretchr/testify/require"
)

//const testMaincurrency := "USD"
//
//type testCurrecnyConverter map[string]float64
//
//func (c testCurrecnyConverter) Convert(money money.Money, toCurrency string) (money.Money, error) {
//	f1, ok := c[m.CurrencyCode]
//	if !ok {
//		return Money{}, errors.New("currency " + m.CurrencyCode + ": no conversion rate")
//	}
//	if f1 < math.SmallestNonzeroFloat64 {
//		return Money{}, fmt.Errorf("currency %s: bad conversion rate: %.2f", m.CurrencyCode, f1)
//	}
//	f2, ok := rates[toCurrencyCode]
//	if !ok {
//		return Money{}, errors.New("currency " + toCurrencyCode + ": no conversion rate")
//	}
//	if f2 < math.SmallestNonzeroFloat64 {
//		return Money{}, fmt.Errorf("currency %s: bad conversion rate: %.2f", toCurrencyCode, f2)
//	}
//	v := float64(m.Amount) / math.Pow10(ScaleFor(m.CurrencyCode))
//	v *= f2 / f1
//	return FromFloat(v, toCurrencyCode), nil
//
//}
//
//func TestSum_Calculate(t *testing.T) {
//	convert
//
//	cases := []struct {
//		items []money.Money
//		want  money.Money
//	}{
//
//	}
//}
func TestSum_Calculate(t *testing.T) {
	converter := money.RatesTableConverter{
		"RUB": 60,
		"USD": 1,
	}

	cases := []struct {
		items []money.Money
		want  money.Money
	}{
		{
			items: nil,
			want: money.Money{
				Amount:   0,
				Currency: "USD",
			},
		},
		{
			items: []money.Money{
				{
					Amount:   60_00,
					Currency: "RUB",
				},
			},
			want: money.Money{
				Amount:   1_00,
				Currency: "USD",
			},
		},
		{
			items: []money.Money{
				{
					Amount:   60_00,
					Currency: "RUB",
				},
				{
					Amount:   60_00,
					Currency: "RUB",
				},
				{
					Amount:   600_00,
					Currency: "RUB",
				},
				{
					Amount:   100_00,
					Currency: "USD",
				},
			},
			want: money.Money{
				Amount:   112_00,
				Currency: "USD",
			},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(fmt.Sprintf("sum %s", c.items), func(t *testing.T) {
			var s money.Sum
			for _, i := range c.items {
				s.Add(i)
			}
			m, err := s.Calculate("USD", converter)
			require.NoError(t, err)
			require.Equal(t, c.want, m)
		})
	}
}

func ExampleSum_Calculate() {
	converter := money.RatesTableConverter{
		"USD": 1,
		"RUB": 60,
	}
	var sum money.Sum
	sum.Add(money.FromFloat(100, "USD"))
	sum.Add(money.FromFloat(600, "RUB"))
	sumValue, _ := sum.Calculate("USD", converter)
	fmt.Println(sumValue)

	// Output: 110.00 USD
}
