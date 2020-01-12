package money_test

import (
	"fmt"
	"testing"

	"github.com/gruzovator/money"
	"github.com/stretchr/testify/require"
)

func init() {
	testScales := map[string]int{
		"USD": 2,
		"KRW": 0,
		"TND": 3,
	}
	money.ReplaceCurrenciesSystem(testScales, 2)
}

func TestMake(t *testing.T) {
	cases := []struct {
		code   string
		amount int64
		scale  int

		want money.Money
	}{
		{
			code:   "USD",
			amount: 1,
			scale:  0,
			want: money.Money{
				Amount:   100,
				Currency: "USD",
			},
		},
		{
			code:   "USD",
			amount: 1,
			scale:  2,
			want: money.Money{
				Amount:   1,
				Currency: "USD",
			},
		},
		{
			code:   "USD",
			amount: 1,
			scale:  -6, //millions
			want: money.Money{
				Amount:   1_000_000_00,
				Currency: "USD",
			},
		},
		{
			code:   "KRW",
			amount: -1,
			scale:  0,
			want: money.Money{
				Amount:   -1,
				Currency: "KRW",
			},
		},
		{
			code:   "KRW",
			amount: -50,
			scale:  2,
			want: money.Money{
				Amount:   -1, // -50 * 10^(-2) and math.Round
				Currency: "KRW",
			},
		},
	}

	for _, c := range cases {
		caseName := fmt.Sprintf("%s_%d_%d", c.code, c.amount, c.scale)
		c := c
		t.Run(caseName, func(t *testing.T) {
			m := money.Make(c.amount, c.scale, c.code)
			require.Equal(t, c.want, m)
		})
	}
}

func TestFromFloat(t *testing.T) {
	cases := []struct {
		code string
		val  float64

		want money.Money
	}{
		{
			code: "USD",
			val:  100.0,
			want: money.Money{
				Amount:   100_00,
				Currency: "USD",
			},
		},
		{
			code: "USD",
			val:  -1_000_000_001.425,
			want: money.Money{
				Amount:   -1_000_000_001_43,
				Currency: "USD",
			},
		},
		{
			code: "KRW",
			val:  42.1234,
			want: money.Money{
				Amount:   42,
				Currency: "KRW",
			},
		},
		{
			code: "KRW",
			val:  0.5,
			want: money.Money{
				Amount:   1,
				Currency: "KRW",
			},
		},
		{
			code: "TND", // scale is 3
			val:  0.1234,
			want: money.Money{
				Amount:   123,
				Currency: "TND",
			},
		},
	}

	for _, c := range cases {
		c := c
		caseName := fmt.Sprintf("%s_%.4f", c.code, c.val)
		t.Run(caseName, func(t *testing.T) {
			m := money.FromFloat(c.val, c.code)
			require.Equal(t, c.want, m)
		})
	}
}

func TestMoney_String(t *testing.T) {
	cases := []struct {
		m    money.Money
		want string
	}{
		{
			m:    money.Money{},
			want: "0.00",
		},
		{
			m:    money.FromFloat(42, "USD"),
			want: "42.00 USD",
		},
		{
			m:    money.FromFloat(1.01, "USD"),
			want: "1.01 USD",
		},
		{
			m:    money.FromFloat(-0.12, "USD"),
			want: "-0.12 USD",
		},
		{
			m:    money.FromFloat(1234, "KRW"),
			want: "1234 KRW",
		},
		{
			m:    money.FromFloat(1.23, "TND"),
			want: "1.230 TND",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(fmt.Sprintf("%d_%s", c.m.Amount, c.m.Currency), func(t *testing.T) {
			require.Equal(t, c.want, c.m.String())
		})
	}
}

func TestFromString(t *testing.T) {
	cases := []struct {
		s    string
		want money.Money
	}{
		{
			s: "12.1234",
			want: money.Money{
				Amount:   12_12,
				Currency: "",
			},
		},
		{
			s: "100 USD",
			want: money.Money{
				Amount:   100_00,
				Currency: "USD",
			},
		},
		{
			s: "-100.123 USD",
			want: money.Money{
				Amount:   -100_12,
				Currency: "USD",
			},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.s, func(t *testing.T) {
			m, err := money.FromString(c.s)
			require.NoError(t, err)
			require.Equal(t, c.want, m)
		})
	}
}

func TestFromString_Error(t *testing.T) {
	cases := []string{
		"",
		"USD 12",
		"1.2a USD",
	}
	for _, c := range cases {
		c := c
		t.Run(c, func(t *testing.T) {
			_, err := money.FromString(c)
			require.Error(t, err)
		})
	}
}

func Example() {
	_1usd := money.Make(1, 0, "USD")
	fmt.Println(_1usd)

	_1usd = money.Make(100, 2, "USD")
	fmt.Println(_1usd)

	_10usd := money.FromFloat(10.001, "USD")
	fmt.Println(_10usd)

	_100usd, _ := money.FromString("100.00 USD")
	fmt.Println(_100usd)

	//Output:
	// 1.00 USD
	// 1.00 USD
	// 10.00 USD
	// 100.00 USD
}
