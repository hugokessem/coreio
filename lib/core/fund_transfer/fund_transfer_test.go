package fundtransfer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFundTransferGeneratedXML(t *testing.T) {
	test := []struct {
		name   string
		param  Params
		expect []string
	}{
		{
			name: "Validate Fund Transfer",
			param: Params{
				Username:            "SUPERAPP",
				Password:            "123456",
				DebitAccountNumber:  "10009876543",
				CreditAccountNumber: "10001234567",
				DebitReference:      "Payment for invoice",
				CreditReference:     "Received payment",
				DebitAmount:         "1500.75",
				TransationID:        "TXN123456",
				PaymentDetail:       "Urgent transfer",
			},
			expect: []string{
				"<fun:DEBITACCTNO>10009876543</fun:DEBITACCTNO>",
				"<fun:DEBITAMOUNT>1500.75</fun:DEBITAMOUNT>",
				"<fun:DEBITTHEIRREF>Payment for invoice</fun:DEBITTHEIRREF>",
				"<fun:CREDITTHEIRREF>Received payment</fun:CREDITTHEIRREF>",
				"<fun:CREDITACCTNO>10001234567</fun:CREDITACCTNO>",
			},
		},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			xmlRequest := NewFundTransfer(tc.param)
			for _, expectedStr := range tc.expect {
				assert.Contains(t, xmlRequest, expectedStr)
			}
		})
	}
}
