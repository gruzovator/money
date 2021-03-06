# Golang Money

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Documentation](https://godoc.org/github.com/gruzovator/money?status.svg)](https://godoc.org/github.com/gruzovator/money)
[![Go Report Card](https://goreportcard.com/badge/github.com/gruzovator/money)](https://goreportcard.com/report/github.com/gruzovator/money)


Money is value type built from currency code and money value in minor units:

```go
// Money is amount of money as integer number of minor units and currency code.
// E.g.: one dollar is Money{100, "USD"} cause "USD" currency has scale 2.
type Money struct {
	Amount   int64
	Currency string
}
```

Currencies scales (or "exponents") are package-level constants from ISO4217 data file.
For unknown currencies scale '2' is used. 

## Usage examples

### Money creation & representation

```go

// create money from amount, non-standard scale and currency code
_1usd := money.Make(1, 0, "USD") 
fmt.Println(_1usd)

_1usd = money.Make(100, 2, "USD")
fmt.Println(_1usd)

// create money from float value rounding to currency minor units.
_10usd := money.FromFloat(10.001, "USD")
fmt.Println(_10usd)

_100usd, _ := money.FromString("100.00 USD")
fmt.Println(_100usd)

//Output:
// 1.00 USD
// 1.00 USD
// 10.00 USD
// 100.00 USD
```

### Money Sum

Sum object keeps money amounts per currency. Total sum is calculated by request using 
currency converter.

```go
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
```     

## TODO

* Distribute