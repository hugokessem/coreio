package fundtransfer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAgentFundTransferGeneratedXML(t *testing.T) {
	test := []struct {
		name   string
		param  Params
		expect []string
	}{
		{
			name: "Validate Agent Fund Transfer",
			param: Params{
				FTNumber:               "FT123456",
				Timestamp:              "20130402152345",
				PrimaryParty:           "251000",
				ReceiverParty:          "251911",
				Amount:                 "1000.00",
				Currency:               "ETB",
				Narative:               "Payment",
				ThirdPartyIdentifier:   "USSDPushCaller",
				Password:               "8eZVmhR2RmGWW/991P8DjLDpHiiiLUle0u",
				SecurityCredential:     "BWJ3KefDOdp+GHqRnA9Yfo2RbsZM60sw",
				DebitAccountNumber:     "1000000006924",
				DebitAccountHolderName: "John Doe",
			},
			expect: []string{
				`<req:Version>1.0</req:Version>`,
				`<req:OriginatorConversationID>FT123456</req:OriginatorConversationID>`,
				`<req:ThirdPartyID>USSDPushCaller</req:ThirdPartyID>`,
				`<req:Password>8eZVmhR2RmGWW/991P8DjLDpHiiiLUle0u</req:Password>`,
				`<req:Timestamp>20130402152345</req:Timestamp>`,
				`<req:SecurityCredential>BWJ3KefDOdp+GHqRnA9Yfo2RbsZM60sw</req:SecurityCredential>`,
				`<req:Amount>1000.00</req:Amount>`,
				`<req:Currency>ETB</req:Currency>`,
				`<req:ReasonType>Payment</req:ReasonType>`,
				`<com:Key>Debited shortcode</com:Key>`,
				`<com:Value>251000</com:Value>`,
				`<com:Key>Debited Customer Name</com:Key>`,
				`<com:Value>John Doe</com:Value>`,
				`<com:Key>Debited Acct</com:Key>`,
				`<com:Value>1000000006924</com:Value>`,
				`<com:Key>MB txnID</com:Key>`,
				`<com:Value>FT123456</com:Value>`,
			},
		},
		{
			name: "Validate Agent Fund Transfer with different values",
			param: Params{
				FTNumber:               "FT789012",
				Timestamp:              "20250101120000",
				PrimaryParty:           "251922",
				ReceiverParty:          "251933",
				Amount:                 "500.50",
				Currency:               "ETB",
				Narative:               "Transfer",
				ThirdPartyIdentifier:   "TESTCALLER",
				Password:               "TESTPASSWORD123",
				SecurityCredential:     "TESTSECURITY123",
				DebitAccountNumber:     "2000000000001",
				DebitAccountHolderName: "Jane Smith",
			},
			expect: []string{
				`<req:OriginatorConversationID>FT789012</req:OriginatorConversationID>`,
				`<req:ThirdPartyID>TESTCALLER</req:ThirdPartyID>`,
				`<req:Password>TESTPASSWORD123</req:Password>`,
				`<req:Timestamp>20250101120000</req:Timestamp>`,
				`<req:SecurityCredential>TESTSECURITY123</req:SecurityCredential>`,
				`<req:Identifier>251922</req:Identifier>`,
				`<req:Identifier>251933</req:Identifier>`,
				`<req:Amount>500.50</req:Amount>`,
				`<req:Currency>ETB</req:Currency>`,
				`<req:ReasonType>Transfer</req:ReasonType>`,
				`<com:Value>Jane Smith</com:Value>`,
				`<com:Value>2000000000001</com:Value>`,
				`<com:Value>FT789012</com:Value>`,
			},
		},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			xmlRequest := NewAgentFundTransfer(tc.param)

			// Validate XML structure
			assert.Contains(t, xmlRequest, "<soapenv:Envelope")
			assert.Contains(t, xmlRequest, "<soapenv:Header/>")
			assert.Contains(t, xmlRequest, "<soapenv:Body>")
			assert.Contains(t, xmlRequest, "<api:Request>")
			assert.Contains(t, xmlRequest, "<req:Header>")
			assert.Contains(t, xmlRequest, "<req:Body>")
			assert.Contains(t, xmlRequest, "<req:Identity>")
			assert.Contains(t, xmlRequest, "<req:TransactionRequest>")
			assert.Contains(t, xmlRequest, "<req:ReferenceData>")

			// Validate all expected strings are present
			for _, expectedStr := range tc.expect {
				assert.Contains(t, xmlRequest, expectedStr)
			}

			// Validate XML is not empty
			assert.NotEmpty(t, xmlRequest)
		})
	}
}
