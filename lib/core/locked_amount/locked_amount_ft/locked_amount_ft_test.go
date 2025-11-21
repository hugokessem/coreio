package lockedamountft

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLockedAmountFtGeneratedXML(t *testing.T) {
	test := []struct {
		name   string
		param  Params
		expect []string
	}{
		{
			name: "Validate Locked Amount Fund Transfer",
			param: Params{
				Username:            "SUPERAPP",
				Password:            "123456",
				CreditCurrent:       "ETB",
				CreditAccountNumber: "1000000006924",
				CrediterReference:   "Credit reference",
				DebitAmount:         "1000.00",
				DebitAccountNumber:  "1000382499388",
				DebitCurrency:       "ETB",
				DebiterReference:    "Debit reference",
				ClientReference:     "CLIENT123456",
				LockID:              "LOCK123456",
			},
			expect: []string{
				`<password>123456</password>`,
				`<userName>SUPERAPP</userName>`,
				`<fun:DEBITACCTNO>1000382499388</fun:DEBITACCTNO>`,
				`<fun:DEBITCURRENCY>ETB</fun:DEBITCURRENCY>`,
				`<fun:DEBITAMOUNT>1000.00</fun:DEBITAMOUNT>`,
				`<fun:DEBITTHEIRREF>Debit reference</fun:DEBITTHEIRREF>`,
				`<fun:CREDITTHEIRREF>Credit reference</fun:CREDITTHEIRREF>`,
				`<fun:CREDITACCTNO>1000000006924</fun:CREDITACCTNO>`,
				`<fun:CREDITCURRENCY>ETB</fun:CREDITCURRENCY>`,
				`<fun:ClientReference>CLIENT123456</fun:ClientReference>`,
				`<fun:LockID>LOCK123456</fun:LockID>`,
			},
		},
		{
			name: "Validate Locked Amount Fund Transfer with different values",
			param: Params{
				Username:            "TESTUSER",
				Password:            "PASSWORD123",
				CreditCurrent:       "USD",
				CreditAccountNumber: "2000000000001",
				CrediterReference:   "Test credit ref",
				DebitAmount:         "500.50",
				DebitAccountNumber:  "2000000000002",
				DebitCurrency:       "USD",
				DebiterReference:    "Test debit ref",
				ClientReference:     "TEST123",
				LockID:              "TESTLOCK456",
			},
			expect: []string{
				`<password>PASSWORD123</password>`,
				`<userName>TESTUSER</userName>`,
				`<fun:DEBITACCTNO>2000000000002</fun:DEBITACCTNO>`,
				`<fun:DEBITCURRENCY>USD</fun:DEBITCURRENCY>`,
				`<fun:DEBITAMOUNT>500.50</fun:DEBITAMOUNT>`,
				`<fun:DEBITTHEIRREF>Test debit ref</fun:DEBITTHEIRREF>`,
				`<fun:CREDITTHEIRREF>Test credit ref</fun:CREDITTHEIRREF>`,
				`<fun:CREDITACCTNO>2000000000001</fun:CREDITACCTNO>`,
				`<fun:CREDITCURRENCY>USD</fun:CREDITCURRENCY>`,
				`<fun:ClientReference>TEST123</fun:ClientReference>`,
				`<fun:LockID>TESTLOCK456</fun:LockID>`,
			},
		},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			xmlRequest := NewLockedAmountFt(tc.param)

			// Validate XML structure
			assert.Contains(t, xmlRequest, "<soapenv:Envelope")
			assert.Contains(t, xmlRequest, "<soapenv:Body>")
			assert.Contains(t, xmlRequest, "<cbes:AccountTransfer>")
			assert.Contains(t, xmlRequest, "<FUNDSTRANSFERFTTXNSUPERAPPType")

			// Validate all expected strings are present
			for _, expectedStr := range tc.expect {
				assert.Contains(t, xmlRequest, expectedStr)
			}

			// Validate XML is not empty
			assert.NotEmpty(t, xmlRequest)
		})
	}
}
