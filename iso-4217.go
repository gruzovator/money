package money

import (
	"encoding/xml"
	"fmt"
	"io"
	"strconv"
)

// CurrencyScalesFromISO4217 parses currencies codes and scales from ISO417 XML data.
//
// See:
//	- https://www.iso.org/iso-4217-currency-codes.html
//	- https://www.currency-iso.org/dam/downloads/lists/list_one.xml
func CurrencyScalesFromISO4217(xmlContentReader io.Reader) (map[string]int, error) {
	data := struct {
		CcyTbl struct {
			CcyNtry []struct {
				CcyNm      string
				Ccy        string
				CcyMnrUnts string // e.g. "N.A."
			}
		}
	}{}
	if err := xml.NewDecoder(xmlContentReader).Decode(&data); err != nil {
		return nil, fmt.Errorf("XML data decode: %s", err)
	}

	scales := make(map[string]int, len(data.CcyTbl.CcyNtry))
	for _, e := range data.CcyTbl.CcyNtry {
		if e.CcyMnrUnts == "N.A." || e.CcyMnrUnts == "" {
			continue
		}
		scale, err := strconv.ParseInt(e.CcyMnrUnts, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("bad scale value %q for %s", e.CcyMnrUnts, e.CcyNm)
		}
		// check that items with same currency code have same scales
		if s, ok := scales[e.Ccy]; ok {
			if s != int(scale) {
				return nil, fmt.Errorf("bad data: different scales for currency code: %s", e.Ccy)
			}
		}
		scales[e.Ccy] = int(scale)
	}
	return scales, nil
}
