package accountlookup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountLookupGeneratedXML(t *testing.T) {
	test := []struct {
		name   string
		param  Params
		expect []string
	}{
		{
			name: "Validate Account Lookups",
			param: Params{
				CreditAccountNumber:  "1234567890",
				DebitBankBIC:         "CBETETAA",
				CreditBankBIC:        "ETSETAA",
				BizMessageIdentifier: "CBETETAA553572981",
				MessageIdentifier:    "CBETETAA847772981",
				CreditDateTime:       "2023-06-24T00:00:00.000+03:00",
				CreditDate:           "2023-06-24T00:00:00.000Z",
			},
			expect: []string{
				`<urn:BizMsgIdr>CBETETAA553572981</urn:BizMsgIdr>`,
				`<urn:CreDt>2023-06-24T00:00:00.000Z</urn:CreDt>`,
				`<urn1:MsgId>CBETETAA847772981</urn1:MsgId>`,
				`<urn1:Id>1234567890</urn1:Id>`,
			},
		},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			xmlRequest := NewAccountLookup(tc.param)
			for _, expectedStr := range tc.expect {
				assert.Contains(t, xmlRequest, expectedStr)
			}
		})
	}
}
