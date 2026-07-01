package fayda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFayda(t *testing.T) {
	tests := []struct {
		name   string
		param  Params
		expect []string
	}{
		{
			name: "Validate Fayda XML generation",
			param: Params{
				Username: "SUPERAPP",
				Password: "123456",
				NID:      "1234567890123456",
			},
			expect: []string{
				`<password>123456</password>`,
				`<userName>SUPERAPP</userName>`,
				`<columnName>NID</columnName>`,
				`<criteriaValue>1234567890123456</criteriaValue>`,
				`<operand>EQ</operand>`,
				`<soapenv:Envelope`,
				`<soapenv:Body>`,
				`<cbes:CustomerUniqueVerification>`,
				`<CUSTOMERVERIFYENQSUPERAPPType>`,
			},
		},
		{
			name: "Validate Fayda XML with different values",
			param: Params{
				Username: "TESTUSER",
				Password: "PASSWORD123",
				NID:      "9876543210987654",
			},
			expect: []string{
				`<password>PASSWORD123</password>`,
				`<userName>TESTUSER</userName>`,
				`<criteriaValue>9876543210987654</criteriaValue>`,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			xmlRequest := NewFayda(tc.param)

			assert.Contains(t, xmlRequest, "<soapenv:Envelope")
			assert.Contains(t, xmlRequest, "<soapenv:Body>")
			assert.Contains(t, xmlRequest, "<cbes:CustomerUniqueVerification>")

			for _, expectedStr := range tc.expect {
				assert.Contains(t, xmlRequest, expectedStr)
			}

			assert.NotEmpty(t, xmlRequest)
		})
	}
}

func TestParseFaydaSOAP(t *testing.T) {
	tests := []struct {
		name                string
		xmlData             string
		expectedSuccess     bool
		expectedError       bool
		expectedDetail      bool
		expectedMessage     string
		expectedCustomerID  string
		expectedName        string
		expectedCustomerFlag string
	}{
		{
			name: "Parse successful response",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns58:CustomerUniqueVerificationResponse xmlns:ns58="http://temenos.com/CBESUPERAPP" xmlns:ns3="http://temenos.com/CUSTOMERVERIFYENQSUPERAPP">
            <Status>
                <successIndicator>Success</successIndicator>
            </Status>
            <CUSTOMERVERIFYENQSUPERAPPType>
                <ns3:gCUSTOMERVERIFYENQSUPERAPPDetailType>
                    <ns3:mCUSTOMERVERIFYENQSUPERAPPDetailType>
                        <ns3:CustomerFlag>FOUND</ns3:CustomerFlag>
                        <ns3:CustomerID>1328828703</ns3:CustomerID>
                        <ns3:SHORTNAME>HIWOT  BEKELE</ns3:SHORTNAME>
                    </ns3:mCUSTOMERVERIFYENQSUPERAPPDetailType>
                </ns3:gCUSTOMERVERIFYENQSUPERAPPDetailType>
            </CUSTOMERVERIFYENQSUPERAPPType>
        </ns58:CustomerUniqueVerificationResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess:      true,
			expectedError:        false,
			expectedDetail:       true,
			expectedCustomerID:   "1328828703",
			expectedName:         "HIWOT  BEKELE",
			expectedCustomerFlag: "FOUND",
		},
		{
			name: "Parse response with failure status",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns58:CustomerUniqueVerificationResponse xmlns:ns58="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>Failure</successIndicator>
                <messages>Verification failed</messages>
            </Status>
        </ns58:CustomerUniqueVerificationResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess: false,
			expectedError:   false,
			expectedDetail:  false,
			expectedMessage: "Verification failed",
		},
		{
			name: "Parse invalid XML",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<InvalidXML>
    <Broken>
</InvalidXML>`,
			expectedSuccess: false,
			expectedError:   true,
			expectedDetail:  false,
		},
		{
			name: "Parse response without Status",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns58:CustomerUniqueVerificationResponse xmlns:ns58="http://temenos.com/CBESUPERAPP">
        </ns58:CustomerUniqueVerificationResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess: false,
			expectedError:   false,
			expectedDetail:  false,
			expectedMessage: "missing status",
		},
		{
			name: "Parse response with no details",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns58:CustomerUniqueVerificationResponse xmlns:ns58="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>Success</successIndicator>
            </Status>
            <CUSTOMERVERIFYENQSUPERAPPType>
            </CUSTOMERVERIFYENQSUPERAPPType>
        </ns58:CustomerUniqueVerificationResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess: false,
			expectedError:   false,
			expectedDetail:  false,
			expectedMessage: "no details found",
		},
		{
			name: "Parse response with empty group",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns58:CustomerUniqueVerificationResponse xmlns:ns58="http://temenos.com/CBESUPERAPP" xmlns:ns3="http://temenos.com/CUSTOMERVERIFYENQSUPERAPP">
            <Status>
                <successIndicator>Success</successIndicator>
            </Status>
            <CUSTOMERVERIFYENQSUPERAPPType>
                <ns3:gCUSTOMERVERIFYENQSUPERAPPDetailType>
                </ns3:gCUSTOMERVERIFYENQSUPERAPPDetailType>
            </CUSTOMERVERIFYENQSUPERAPPType>
        </ns58:CustomerUniqueVerificationResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess: false,
			expectedError:   false,
			expectedDetail:  false,
			expectedMessage: "no details found",
		},
		{
			name: "Parse response with invalid response structure",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <OtherResponse xmlns:ns58="http://temenos.com/CBESUPERAPP">
        </OtherResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess: false,
			expectedError:   false,
			expectedDetail:  false,
			expectedMessage: "Invalid response type",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ParseFaydaSOAP(tc.xmlData)

			if tc.expectedError {
				assert.Error(t, err)
				assert.Nil(t, result)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, result)
			assert.Equal(t, tc.expectedSuccess, result.Success)

			if tc.expectedDetail {
				assert.NotNil(t, result.Detail)
				if result.Detail != nil {
					assert.Equal(t, tc.expectedCustomerID, result.Detail.CustomerID)
					assert.Equal(t, tc.expectedName, result.Detail.CustomerName)
					assert.Equal(t, tc.expectedCustomerFlag, result.Detail.CustomerFlag)
				}
				return
			}

			assert.Nil(t, result.Detail)
			if tc.expectedMessage != "" {
				assert.Contains(t, result.Message, tc.expectedMessage)
			}
		})
	}
}

func TestParseFaydaSOAP_CaseInsensitiveSuccessIndicator(t *testing.T) {
	xmlData := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns58:CustomerUniqueVerificationResponse xmlns:ns58="http://temenos.com/CBESUPERAPP" xmlns:ns3="http://temenos.com/CUSTOMERVERIFYENQSUPERAPP">
            <Status>
                <successIndicator>SUCCESS</successIndicator>
            </Status>
            <CUSTOMERVERIFYENQSUPERAPPType>
                <ns3:gCUSTOMERVERIFYENQSUPERAPPDetailType>
                    <ns3:mCUSTOMERVERIFYENQSUPERAPPDetailType>
                        <ns3:CustomerFlag>FOUND</ns3:CustomerFlag>
                        <ns3:CustomerID>1328828703</ns3:CustomerID>
                        <ns3:SHORTNAME>HIWOT  BEKELE</ns3:SHORTNAME>
                    </ns3:mCUSTOMERVERIFYENQSUPERAPPDetailType>
                </ns3:gCUSTOMERVERIFYENQSUPERAPPDetailType>
            </CUSTOMERVERIFYENQSUPERAPPType>
        </ns58:CustomerUniqueVerificationResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseFaydaSOAP(xmlData)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Success)
	assert.NotNil(t, result.Detail)
	assert.Equal(t, "1328828703", result.Detail.CustomerID)
}
