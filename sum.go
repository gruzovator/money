package money

// Sum aggregates money values by currency.
type Sum struct {
	AmountsByCurrency map[string]int64
}

func (s *Sum) Add(m Money) {
	if s.AmountsByCurrency == nil {
		s.AmountsByCurrency = make(map[string]int64)
	}
	s.AmountsByCurrency[m.Currency] += m.Amount
}

// Calcuelate calculates sum value in target currency.
func (s *Sum) Calculate(targetCurrency string, converter Converter) (Money, error) {
	var sumAmount int64
	for c, a := range s.AmountsByCurrency {
		converted, err := converter.Convert(Money{
			Amount:   a,
			Currency: c,
		}, targetCurrency)
		if err != nil {
			return Money{}, err
		}
		sumAmount += converted.Amount
	}
	return Money{
		Amount:   sumAmount,
		Currency: targetCurrency,
	}, nil
}
