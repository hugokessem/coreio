package standingorderupdate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateStandingOrderGeneratedXML(t *testing.T) {
	test := []struct {
		name   string
		param  Params
		expect []string
	}{
		{
			name: "Validate Update Standing Order",
			param: Params{
				Username:            "SUPERAPP",
				Password:            "123456",
				DebitAccountNumber:  "1000648501521",
				OrderId:             "3",
				Currency:            "ETB",
				Amount:              "2500",
				Frequency:           "20220512 e0Y e0M e0W e1D e0F",
				CurrentDate:         "20260101",
				PaymentDetail:       "TEST",
				CreditAccountNumber: "1000035279327",
			},
			expect: []string{
				`<STANDINGORDERMANAGEORDERSUPERAPPType id="1000648501521.3">`,
				`<stan:CURRENCY>ETB</stan:CURRENCY>`,
				`<stan:CURRENTAMOUNTBAL>2500</stan:CURRENTAMOUNTBAL>`,
				`<stan:CURRENTFREQUENCY>20220512 e0Y e0M e0W e1D e0F</stan:CURRENTFREQUENCY>`,
				`<stan:CURRENTENDDATE>20260101</stan:CURRENTENDDATE>`,
				`<stan:PAYMENTDETAILS>TEST</stan:PAYMENTDETAILS>`,
				`<stan:CPTYACCTNO>1000035279327</stan:CPTYACCTNO>`,
			},
		},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			xmlRequest := NewUpdateStandingOrder(tc.param)
			for _, expectStr := range tc.expect {
				assert.Contains(t, xmlRequest, expectStr)
			}
		})
	}

}
