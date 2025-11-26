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

func TestParseAgentLookupSOAP(t *testing.T) {
	xmlResponse := `<?xml version="1.0" encoding="utf-8"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
    <soapenv:Body>
        <api:Result xmlns:api="http://cps.huawei.com/synccpsinterface/api_requestmgr" xmlns:res="http://cps.huawei.com/synccpsinterface/result">
            <res:Header>
                <res:Version>1.0</res:Version>
                <res:OriginatorConversationID>S_334220042321001</res:OriginatorConversationID>
                <res:ConversationID>AG_20251127_7010236952534809c7c2</res:ConversationID>
            </res:Header>
            <res:Body>
                <res:ResultType>0</res:ResultType>
                <res:ResultCode>0</res:ResultCode>
                <res:ResultDesc>Process service request successfully.</res:ResultDesc>
                <res:QueryOrganizationInfoResult>
                    <res:BOCompletedTime>20251127012642</res:BOCompletedTime>
                    <res:OrganizationBasicData>
                        <res:ShortCode>251000</res:ShortCode>
                        <res:OrganizationName>gedaplc</res:OrganizationName>
                        <res:IdentityStatus>03</res:IdentityStatus>
                        <res:CreationDate>20160810</res:CreationDate>
                        <res:TrustLevel>11</res:TrustLevel>
                        <res:TrustLevelName>Top Organization Trust Level</res:TrustLevelName>
                        <res:RuleProfileID>11024</res:RuleProfileID>
                        <res:RuleProfileName>Default Organization Rule Profile</res:RuleProfileName>
                        <res:ChargeProfileID>92</res:ChargeProfileID>
                        <res:ChargeProfileName>Bank Branch Charge Profile</res:ChargeProfileName>
                        <res:AggregatorAcctModel>Accounting Model for Aggregator</res:AggregatorAcctModel>
                        <res:HierarchyLevel>1</res:HierarchyLevel>
                        <res:HierarchyModel>Aggregator</res:HierarchyModel>
                    </res:OrganizationBasicData>
                </res:QueryOrganizationInfoResult>
            </res:Body>
        </api:Result>
    </soapenv:Body>
</soapenv:Envelope>`

	result, err := ParseAgentLookupSOAP(xmlResponse)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Success)
	assert.NotNil(t, result.Detail)
	
	if result.Detail != nil {
		assert.Equal(t, "1.0", result.Detail.Version)
		assert.Equal(t, "S_334220042321001", result.Detail.OriginalConverstationIdentifier)
		assert.Equal(t, "AG_20251127_7010236952534809c7c2", result.Detail.ConversationIdentifier)
		
		if result.Detail.OrganizationBasicData.OrganizationBasicData != nil {
			orgData := result.Detail.OrganizationBasicData.OrganizationBasicData
			assert.Equal(t, "251000", orgData.ShortCode)
			assert.Equal(t, "gedaplc", orgData.OrganizationName)
			assert.Equal(t, "03", orgData.IdentityStatus)
			assert.Equal(t, "20160810", orgData.CreationDate)
			assert.Equal(t, "11", orgData.TrustLevel)
			assert.Equal(t, "Top Organization Trust Level", orgData.TrustLevelName)
			assert.Equal(t, "11024", orgData.RuleProfileID)
			assert.Equal(t, "Default Organization Rule Profile", orgData.RuleProfileName)
			assert.Equal(t, "92", orgData.ChargeProfileID)
			assert.Equal(t, "Bank Branch Charge Profile", orgData.ChargeProfileName)
			assert.Equal(t, "Accounting Model for Aggregator", orgData.AggregatorAcctModel)
			assert.Equal(t, "1", orgData.HierarchyLevel)
			assert.Equal(t, "Aggregator", orgData.HierarchyModel)
		}
	}
}
