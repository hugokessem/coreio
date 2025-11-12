package standingordercancel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCancleStandingPaymentGenerateXML(t *testing.T) {
	test := []struct {
		name   string
		param  Params
		expect []string
	}{
		{
			name: "Validate Cancle Standing Payment",
			param: Params{
				Username:      "SUPERAPP",
				Password:      "123456",
				AccountNumber: "1000648501521",
				OrderId:       "3",
			},
			expect: []string{
				"<transactionId>1000648501521.3</transactionId>",
			},
		},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			xmlRequest := NewCancleStandingOrder(tc.param)
			for _, expectedStr := range tc.expect {
				assert.Contains(t, xmlRequest, expectedStr)
			}
		})
	}
}
