package lockedamountrelease

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
			name: "Validate Relese Locked Amount",
			param: Params{
				Username:      "SUPERAPP",
				Password:      "123456",
				TransactionID: "ACLK213436WRSG",
			},
			expect: []string{
				`<transactionId>ACLK213436WRSG</transactionId>`,
			},
		},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			xmlRequest := NewReleaseLockedAmount(tc.param)
			for _, expectedStr := range tc.expect {
				assert.Contains(t, xmlRequest, expectedStr)
			}
		})
	}
}
