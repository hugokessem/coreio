package eligibility

import (
	"testing"

	valueobject "github.com/hugokessem/coreio/lib/core/eligibility/value_object"
	"github.com/stretchr/testify/assert"
)

func TestNewEligibility(t *testing.T) {
	tests := []struct {
		name   string
		param  Param
		expect []string
	}{
		{
			name: "Validate eligibility XML generation by account number",
			param: Param{
				Username:      "SUPERAPP",
				Password:      "123456",
				FetchBy:       valueobject.FetchByAccountNumber.String(),
				CriticalValue: "1000041045384696",
			},
			expect: []string{
				`<password>123456</password>`,
				`<userName>SUPERAPP</userName>`,
				`<columnName>CUS.ID</columnName>`,
				`<columnName>ACCT.ID</columnName>`,
				`<criteriaValue></criteriaValue>`,
				`<criteriaValue>1000041045384696</criteriaValue>`,
				`<operand>EQ</operand>`,
				`<soapenv:Envelope`,
				`<soapenv:Body>`,
				`<cbes:CustomerEligibilityforOnboarding>`,
				`<ACCOUNTINFOSUPERAPPRESTRICTType>`,
			},
		},
		{
			name: "Validate eligibility XML generation by customer number",
			param: Param{
				Username:      "SUPERAPP",
				Password:      "123456",
				FetchBy:       valueobject.FetchByCustomerNumber.String(),
				CriticalValue: "1045384696",
			},
			expect: []string{
				`<password>123456</password>`,
				`<userName>SUPERAPP</userName>`,
				`<criteriaValue>1045384696</criteriaValue>`,
				`<criteriaValue></criteriaValue>`,
				`<cbes:CustomerEligibilityforOnboarding>`,
			},
		},
		{
			name: "Validate eligibility XML with different credentials",
			param: Param{
				Username:      "TESTUSER",
				Password:      "PASSWORD123",
				FetchBy:       valueobject.FetchByAccountNumber.String(),
				CriticalValue: "1000123456789",
			},
			expect: []string{
				`<password>PASSWORD123</password>`,
				`<userName>TESTUSER</userName>`,
				`<criteriaValue>1000123456789</criteriaValue>`,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			xmlRequest := NewEligibility(tc.param)

			assert.Contains(t, xmlRequest, "<soapenv:Envelope")
			assert.Contains(t, xmlRequest, "<soapenv:Body>")
			assert.Contains(t, xmlRequest, "<cbes:CustomerEligibilityforOnboarding>")

			for _, expectedStr := range tc.expect {
				assert.Contains(t, xmlRequest, expectedStr)
			}

			assert.NotEmpty(t, xmlRequest)
		})
	}
}

func TestParseCustomerEligibilitySOAP(t *testing.T) {
	tests := []struct {
		name            string
		xmlData         string
		expectedSuccess bool
		expectedError   bool
		expectedDetails int
		expectedMessage string
	}{
		{
			name: "Parse successful response",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns30:CustomerEligibilityforOnboardingResponse xmlns:ns30="http://temenos.com/CBESUPERAPP" xmlns:ns10="http://temenos.com/ACCOUNTINFOSUPERAPPRESTRICT">
            <Status>
                <successIndicator>success</successIndicator>
            </Status>
            <ACCOUNTINFOSUPERAPPRESTRICTType>
                <ns10:gACCOUNTINFOSUPERAPPRESTRICTDetailType>
                    <ns10:mACCOUNTINFOSUPERAPPRESTRICTDetailType>
                        <ns10:AccountNumber>1000041045384696</ns10:AccountNumber>
                        <ns10:CustomerName>JOHN DOE</ns10:CustomerName>
                        <ns10:Restriction>NONE</ns10:Restriction>
                        <ns10:Currency>ETB</ns10:Currency>
                        <ns10:CustomerID>1045384696</ns10:CustomerID>
                        <ns10:PhoneNo>+251911706628</ns10:PhoneNo>
                        <ns10:Email>john@example.com</ns10:Email>
                    </ns10:mACCOUNTINFOSUPERAPPRESTRICTDetailType>
                </ns10:gACCOUNTINFOSUPERAPPRESTRICTDetailType>
            </ACCOUNTINFOSUPERAPPRESTRICTType>
        </ns30:CustomerEligibilityforOnboardingResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess: true,
			expectedError:   false,
			expectedDetails: 1,
		},
		{
			name: "Parse response with failure status",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns30:CustomerEligibilityforOnboardingResponse xmlns:ns30="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>failure</successIndicator>
                <messages>Customer not eligible</messages>
            </Status>
        </ns30:CustomerEligibilityforOnboardingResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess: false,
			expectedError:   false,
			expectedDetails: 0,
			expectedMessage: "Customer not eligible",
		},
		{
			name: "Parse invalid XML",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<InvalidXML>
    <Broken>
</InvalidXML>`,
			expectedSuccess: false,
			expectedError:   true,
			expectedDetails: 0,
		},
		{
			name: "Parse response without Status",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns30:CustomerEligibilityforOnboardingResponse xmlns:ns30="http://temenos.com/CBESUPERAPP">
        </ns30:CustomerEligibilityforOnboardingResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess: false,
			expectedError:   false,
			expectedDetails: 0,
			expectedMessage: "Missing status",
		},
		{
			name: "Parse response with no details",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns30:CustomerEligibilityforOnboardingResponse xmlns:ns30="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>success</successIndicator>
            </Status>
            <ACCOUNTINFOSUPERAPPRESTRICTType>
            </ACCOUNTINFOSUPERAPPRESTRICTType>
        </ns30:CustomerEligibilityforOnboardingResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess: false,
			expectedError:   false,
			expectedDetails: 0,
			expectedMessage: "Missing details",
		},
		{
			name: "Parse response with empty group",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns30:CustomerEligibilityforOnboardingResponse xmlns:ns30="http://temenos.com/CBESUPERAPP" xmlns:ns10="http://temenos.com/ACCOUNTINFOSUPERAPPRESTRICT">
            <Status>
                <successIndicator>success</successIndicator>
            </Status>
            <ACCOUNTINFOSUPERAPPRESTRICTType>
                <ns10:gACCOUNTINFOSUPERAPPRESTRICTDetailType>
                </ns10:gACCOUNTINFOSUPERAPPRESTRICTDetailType>
            </ACCOUNTINFOSUPERAPPRESTRICTType>
        </ns30:CustomerEligibilityforOnboardingResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess: false,
			expectedError:   false,
			expectedDetails: 0,
			expectedMessage: "Missing details",
		},
		{
			name: "Parse response with invalid response structure",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <OtherResponse xmlns:ns30="http://temenos.com/CBESUPERAPP">
        </OtherResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess: false,
			expectedError:   false,
			expectedDetails: 0,
			expectedMessage: "Failed to parse eligibility response",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ParseCustomerEligibilitySOAP(tc.xmlData)

			if tc.expectedError {
				assert.Error(t, err)
				assert.Nil(t, result)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, result)
			assert.Equal(t, tc.expectedSuccess, result.Success)
			assert.Len(t, result.Details, tc.expectedDetails)

			if tc.expectedMessage != "" {
				assert.Contains(t, result.Messages, tc.expectedMessage)
			}
		})
	}
}

func TestParseCustomerEligibilitySOAP_DetailFields(t *testing.T) {
	xmlData := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns30:CustomerEligibilityforOnboardingResponse xmlns:ns30="http://temenos.com/CBESUPERAPP" xmlns:ns10="http://temenos.com/ACCOUNTINFOSUPERAPPRESTRICT">
            <Status>
                <successIndicator>success</successIndicator>
            </Status>
            <ACCOUNTINFOSUPERAPPRESTRICTType>
                <ns10:gACCOUNTINFOSUPERAPPRESTRICTDetailType>
                    <ns10:mACCOUNTINFOSUPERAPPRESTRICTDetailType>
                        <ns10:AccountNumber>1000041045384696</ns10:AccountNumber>
                        <ns10:CustomerName>JOHN DOE</ns10:CustomerName>
                        <ns10:Restriction>NONE</ns10:Restriction>
                        <ns10:Currency>ETB</ns10:Currency>
                        <ns10:CustomerID>1045384696</ns10:CustomerID>
                        <ns10:Category>1001</ns10:Category>
                        <ns10:AccountType>Savings</ns10:AccountType>
                        <ns10:BranchCode>001</ns10:BranchCode>
                        <ns10:BranchName>Main Branch</ns10:BranchName>
                        <ns10:PhoneNo>+251911706628</ns10:PhoneNo>
                        <ns10:Email>john@example.com</ns10:Email>
                        <ns10:RestrictionType>0</ns10:RestrictionType>
                    </ns10:mACCOUNTINFOSUPERAPPRESTRICTDetailType>
                </ns10:gACCOUNTINFOSUPERAPPRESTRICTDetailType>
            </ACCOUNTINFOSUPERAPPRESTRICTType>
        </ns30:CustomerEligibilityforOnboardingResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseCustomerEligibilitySOAP(xmlData)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Success)
	assert.Len(t, result.Details, 1)

	detail := result.Details[0]
	assert.Equal(t, "1000041045384696", detail.AccountNumber)
	assert.Equal(t, "JOHN DOE", detail.CustomerName)
	assert.Equal(t, "NONE", detail.Restriction)
	assert.Equal(t, "ETB", detail.Currency)
	assert.Equal(t, "1045384696", detail.CustomerID)
	assert.Equal(t, "+251911706628", detail.PhoneNo)
	assert.Equal(t, "john@example.com", detail.Email)
}

func TestParseCustomerEligibilitySOAP_SuccessIndicatorCase(t *testing.T) {
	xmlData := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns30:CustomerEligibilityforOnboardingResponse xmlns:ns30="http://temenos.com/CBESUPERAPP" xmlns:ns10="http://temenos.com/ACCOUNTINFOSUPERAPPRESTRICT">
            <Status>
                <successIndicator>Success</successIndicator>
            </Status>
            <ACCOUNTINFOSUPERAPPRESTRICTType>
                <ns10:gACCOUNTINFOSUPERAPPRESTRICTDetailType>
                    <ns10:mACCOUNTINFOSUPERAPPRESTRICTDetailType>
                        <ns10:AccountNumber>1000041045384696</ns10:AccountNumber>
                        <ns10:CustomerName>JOHN DOE</ns10:CustomerName>
                    </ns10:mACCOUNTINFOSUPERAPPRESTRICTDetailType>
                </ns10:gACCOUNTINFOSUPERAPPRESTRICTDetailType>
            </ACCOUNTINFOSUPERAPPRESTRICTType>
        </ns30:CustomerEligibilityforOnboardingResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseCustomerEligibilitySOAP(xmlData)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.False(t, result.Success)
}
