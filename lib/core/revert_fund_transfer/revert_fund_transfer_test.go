package revertfundtransfer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRevertFundTransferGeneratedXML(t *testing.T) {
	test := []struct {
		name   string
		param  Params
		expect []string
	}{
		{
			name: "Validate Revert Fund Transfer",
			param: Params{
				Username:      "SUPERAPP",
				Password:      "123456",
				TransactionID: "FT21343GPBZ6",
			},
			expect: []string{
				`<transactionId>FT21343GPBZ6</transactionId>`,
			},
		},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			xmlRequest := NewRevertFundTransfer(tc.param)
			for _, expectedStr := range tc.expect {
				assert.Contains(t, xmlRequest, expectedStr)
			}
		})
	}
}
