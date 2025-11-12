package standingordercreate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateStandingOrderGeneratedXML(t *testing.T) {
	test := []struct {
		name   string
		param  Params
		expect []string
	}{
		{
			name: "Validate Create Standing Order",
			param: Params{
				Username:            "SUPERAPP",
				Password:            "123456",
				DebitAccountNumber:  "1000648501521",
				CreditAccountNumber: "1000035279327",
				CurrentDate:         "20260101",
				Amount:              "7000",
				Currency:            "ETB",
				Frequency:           "20220512 e0Y e0M e1W e0D e0F",
				PaymentDetail:       "TEST",
			},
			expect: []string{
				"<stan:CURRENCY>ETB</stan:CURRENCY>",
				"<stan:CURRENTAMOUNTBAL>7000</stan:CURRENTAMOUNTBAL>",
				"<stan:CURRENTFREQUENCY>20220512 e0Y e0M e1W e0D e0F</stan:CURRENTFREQUENCY>",
				"<stan:CURRENTENDDATE>20260101</stan:CURRENTENDDATE>",
				"<stan:PAYMENTDETAILS>TEST</stan:PAYMENTDETAILS>",
				"<stan:CPTYACCTNO>1000035279327</stan:CPTYACCTNO>",
			},
		},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			xmlRequest := NewCreateStandingOrder(tc.param)
			for _, expectedStr := range tc.expect {
				assert.Contains(t, xmlRequest, expectedStr)
			}
		})
	}
}
