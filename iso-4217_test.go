package money_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gruzovator/money"
	"github.com/stretchr/testify/require"
)

func TestCurrencyScalesFromISO4217(t *testing.T) {
	xmlData := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<ISO_4217 Pblshd="2018-08-29">
	<CcyTbl>
		<CcyNtry>
			<CtryNm>AFGHANISTAN</CtryNm>
			<CcyNm>Afghani</CcyNm>
			<Ccy>AFN</Ccy>
			<CcyNbr>971</CcyNbr>
			<CcyMnrUnts>2</CcyMnrUnts>
		</CcyNtry>
		<CcyNtry>
			<CtryNm>ÅLAND ISLANDS</CtryNm>
			<CcyNm>Euro</CcyNm>
			<Ccy>EUR</Ccy>
			<CcyNbr>978</CcyNbr>
			<CcyMnrUnts>2</CcyMnrUnts>
		</CcyNtry>
	</CcyTbl>
</ISO_4217>
`
	scales, err := money.CurrencyScalesFromISO4217(strings.NewReader(xmlData))

	assert := require.New(t)
	assert.NoError(err)
	assert.Equal(map[string]int{
		"AFN": 2,
		"EUR": 2,
	}, scales)
}

func TestCurrencyScalesFromISO4217_Errors(t *testing.T) {
	cases := []struct {
		name          string
		content       string
		wantErrPrefix string
	}{
		{
			name:          "bad xml",
			content:       `<?xml version="1.0" encoding="UTF-8" standalone="yes"`,
			wantErrPrefix: "XML data decode",
		},
		{
			name: "bad currency scale",
			content: `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<ISO_4217 Pblshd="2018-08-29">
	<CcyTbl>
		<CcyNtry>
			<CtryNm>AFGHANISTAN</CtryNm>
			<CcyNm>Afghani</CcyNm>
			<Ccy>AFN</Ccy>
			<CcyNbr>971</CcyNbr>
			<CcyMnrUnts>?</CcyMnrUnts>
		</CcyNtry>
	</CcyTbl>
</ISO_4217>
`,
			wantErrPrefix: "bad scale value",
		},
		{
			name: "a currency code with different scales",
			content: `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<ISO_4217 Pblshd="2018-08-29">
	<CcyTbl>
		<CcyNtry>
			<CtryNm>AFGHANISTAN</CtryNm>
			<CcyNm>Afghani</CcyNm>
			<Ccy>AFN</Ccy>
			<CcyNbr>971</CcyNbr>
			<CcyMnrUnts>2</CcyMnrUnts>
		</CcyNtry>
		<CcyNtry>
			<CtryNm>AFGHANISTAN</CtryNm>
			<CcyNm>Afghani</CcyNm>
			<Ccy>AFN</Ccy>
			<CcyNbr>971</CcyNbr>
			<CcyMnrUnts>3</CcyMnrUnts>
		</CcyNtry>
	</CcyTbl>
</ISO_4217>
`,
			wantErrPrefix: "bad data:",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			_, err := money.CurrencyScalesFromISO4217(strings.NewReader(c.content))
			require.Error(t, err)
			require.Regexp(t, fmt.Sprintf("^%s.*", c.wantErrPrefix), err.Error())
		})
	}
}
