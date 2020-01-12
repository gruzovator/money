# Golang Money

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Documentation](https://godoc.org/github.com/gruzovator/money?status.svg)](https://godoc.org/github.com/gruzovator/money)
[![Go Report Card](https://goreportcard.com/badge/github.com/gruzovator/money)](https://goreportcard.com/report/github.com/gruzovator/money)


Money is value type built from currency code and money value in minor units:

```go
type Money struct {
	Amount   int64
	Currency string
}
```

Currencies scales (or "exponents") are package-level constants from ISO4217 data file.
For unknown currencies scale '2' is used. 

## Usage example

```go
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
```     

## TODO

* Currency Converter interface
* Sum object to sum money value (per currency) with Value(c Converter) methods to get sum value
* Distribute