package ministatementbydaterange

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMiniStatementByDateRangeGeneratedXML(t *testing.T) {
	test := []struct {
		name   string
		param  Params
		expect []string
	}{
		{
			name: "Validate Mini Statement by Date Range",
			param: Params{
				Username:      "SUPERAPP",
				Password:      "123456",
				AccountNumber: "1000030677308",
				From:          "20250101",
				To:            "20250131",
			},
			expect: []string{
				`<password>123456</password>`,
				`<userName>SUPERAPP</userName>`,
				`<columnName>ACCOUNT</columnName>`,
				`<criteriaValue>1000030677308</criteriaValue>`,
				`<columnName>BOOKING.DATE</columnName>`,
				`<criteriaValue>20250101 20250131</criteriaValue>`,
				`<operand>EQ</operand>`,
			},
		},
		{
			name: "Validate Mini Statement by Date Range with different values",
			param: Params{
				Username:      "TESTUSER",
				Password:      "PASSWORD123",
				AccountNumber: "2000000000001",
				From:          "20241201",
				To:            "20241231",
			},
			expect: []string{
				`<password>PASSWORD123</password>`,
				`<userName>TESTUSER</userName>`,
				`<criteriaValue>2000000000001</criteriaValue>`,
				`<criteriaValue>20241201 20241231</criteriaValue>`,
			},
		},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			xmlRequest := NewMiniStatementByDateRange(tc.param)

			// Validate XML structure
			assert.Contains(t, xmlRequest, "<soapenv:Envelope")
			assert.Contains(t, xmlRequest, "<soapenv:Body>")
			assert.Contains(t, xmlRequest, "<cbes:AccountStatementByRange>")
			assert.Contains(t, xmlRequest, "<ACCTSTMTRGSUPERAPPType>")

			// Validate all expected strings are present
			for _, expectedStr := range tc.expect {
				assert.Contains(t, xmlRequest, expectedStr)
			}

			// Validate XML is not empty
			assert.NotEmpty(t, xmlRequest)
		})
	}
}
