package accountlookup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAgentAccountLookupGeneratedXML(t *testing.T) {
	test := []struct {
		name   string
		param  Params
		expect []string
	}{
		{
			name: "Validate Agent Account Lookup",
			param: Params{
				OriginalConverstationIdentifier: "CONV123456",
				ThirdPartyIdentifier:            "USSDPushCaller",
				Password:                        "8eZVmhR2RmGWW/991P8DjLDpHiiiLUle0u",
				Timestamp:                       "20130402152345",
				SecurityCredential:              "BWJ3KefDOdp+GHqRnA9Yfo2RbsZM60sw",
				PhoneNumber:                     "251000",
			},
			expect: []string{
				`<req:Version>1.0</req:Version>`,
				`<req:CommandID>QueryOrganizationInfo</req:CommandID>`,
				`<req:OriginatorConversationID>CONV123456</req:OriginatorConversationID>`,
				`<req:CallerType>2</req:CallerType>`,
				`<req:ThirdPartyID>USSDPushCaller</req:ThirdPartyID>`,
				`<req:Password>8eZVmhR2RmGWW/991P8DjLDpHiiiLUle0u</req:Password>`,
				`<req:Timestamp>20130402152345</req:Timestamp>`,
				`<req:IdentifierType>14</req:IdentifierType>`,
				`<req:Identifier>Anamail</req:Identifier>`,
				`<req:SecurityCredential>BWJ3KefDOdp+GHqRnA9Yfo2RbsZM60sw</req:SecurityCredential>`,
				`<req:IdentifierType>4</req:IdentifierType>`,
				`<req:Identifier>251000</req:Identifier>`,
				`<req:QueryOrganizationInfoRequest/>`,
			},
		},
		{
			name: "Validate Agent Account Lookup with different values",
			param: Params{
				OriginalConverstationIdentifier: "CONV789012",
				ThirdPartyIdentifier:            "TESTCALLER",
				Password:                        "TESTPASSWORD123",
				Timestamp:                       "20250101120000",
				SecurityCredential:              "TESTSECURITY123",
				PhoneNumber:                     "251911",
			},
			expect: []string{
				`<req:OriginatorConversationID>CONV789012</req:OriginatorConversationID>`,
				`<req:ThirdPartyID>TESTCALLER</req:ThirdPartyID>`,
				`<req:Password>TESTPASSWORD123</req:Password>`,
				`<req:Timestamp>20250101120000</req:Timestamp>`,
				`<req:SecurityCredential>TESTSECURITY123</req:SecurityCredential>`,
				`<req:Identifier>251911</req:Identifier>`,
			},
		},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			xmlRequest := NewAgentAccountLookup(tc.param)

			// Validate XML structure
			assert.Contains(t, xmlRequest, "<soapenv:Envelope")
			assert.Contains(t, xmlRequest, "<soapenv:Header/>")
			assert.Contains(t, xmlRequest, "<soapenv:Body>")
			assert.Contains(t, xmlRequest, "<api:Request>")
			assert.Contains(t, xmlRequest, "<req:Header>")
			assert.Contains(t, xmlRequest, "<req:Body>")
			assert.Contains(t, xmlRequest, "<req:Identity>")

			// Validate all expected strings are present
			for _, expectedStr := range tc.expect {
				assert.Contains(t, xmlRequest, expectedStr)
			}

			// Validate XML is not empty
			assert.NotEmpty(t, xmlRequest)
		})
	}
}
