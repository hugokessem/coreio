package phonelookup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPhoneLookup(t *testing.T) {
	tests := []struct {
		name   string
		param  Params
		expect []string
	}{
		{
			name: "Validate phone lookup XML generation",
			param: Params{
				Username:    "SUPERAPP",
				Password:    "123456",
				PhoneNumber: "+251911706628",
			},
			expect: []string{
				`<password>123456</password>`,
				`<userName>SUPERAPP</userName>`,
				`<columnName>MNEMONIC</columnName>`,
				`<criteriaValue>+251911706628</criteriaValue>`,
				`<operand>EQ</operand>`,
				`<soapenv:Envelope`,
				`<soapenv:Body>`,
				`<cbes:GetCustomerPhoneNo>`,
			},
		},
		{
			name: "Validate phone lookup with different values",
			param: Params{
				Username:    "TESTUSER",
				Password:    "PASSWORD123",
				PhoneNumber: "251913323918",
			},
			expect: []string{
				`<password>PASSWORD123</password>`,
				`<userName>TESTUSER</userName>`,
				`<criteriaValue>251913323918</criteriaValue>`,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			xmlRequest := NewPhoneLookup(tc.param)

			// Validate XML structure
			assert.Contains(t, xmlRequest, "<soapenv:Envelope")
			assert.Contains(t, xmlRequest, "<soapenv:Body>")
			assert.Contains(t, xmlRequest, "<cbes:GetCustomerPhoneNo>")

			// Validate all expected strings are present
			for _, expectedStr := range tc.expect {
				assert.Contains(t, xmlRequest, expectedStr)
			}

			// Validate XML is not empty
			assert.NotEmpty(t, xmlRequest)
		})
	}
}

func TestParsePhoneLookupSOAP(t *testing.T) {
	tests := []struct {
		name                 string
		xmlData              string
		expectedSuccess      bool
		expectedError        bool
		expectedDetail       bool
		expectedMessage      string
		expectedCustomerID   string
		expectedPhoneNumber  string
		expectedEmail        string
	}{
		{
			name: "Parse successful response",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns30:GetCustomerPhoneNoResponse xmlns:ns30="http://temenos.com/CBESUPERAPP" xmlns:ns10="http://temenos.com/GETPHONECUSTOMER">
            <Status>
                <successIndicator>Success</successIndicator>
            </Status>
            <GETPHONECUSTOMERType>
                <ns10:gGETPHONECUSTOMERDetailType>
                    <ns10:mGETPHONECUSTOMERDetailType>
                        <ns10:CustomerID>1771239173</ns10:CustomerID>
                        <ns10:PhoneNumber>+251911706628</ns10:PhoneNumber>
                        <ns10:Email>test@example.com</ns10:Email>
                    </ns10:mGETPHONECUSTOMERDetailType>
                </ns10:gGETPHONECUSTOMERDetailType>
            </GETPHONECUSTOMERType>
        </ns30:GetCustomerPhoneNoResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess:     true,
			expectedError:       false,
			expectedDetail:      true,
			expectedCustomerID:  "1771239173",
			expectedPhoneNumber:  "+251911706628",
			expectedEmail:        "test@example.com",
		},
		{
			name: "Parse successful response with empty email",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns30:GetCustomerPhoneNoResponse xmlns:ns30="http://temenos.com/CBESUPERAPP" xmlns:ns10="http://temenos.com/GETPHONECUSTOMER">
            <Status>
                <successIndicator>Success</successIndicator>
            </Status>
            <GETPHONECUSTOMERType>
                <ns10:gGETPHONECUSTOMERDetailType>
                    <ns10:mGETPHONECUSTOMERDetailType>
                        <ns10:CustomerID>1771239173</ns10:CustomerID>
                        <ns10:PhoneNumber>+251911706628</ns10:PhoneNumber>
                        <ns10:Email></ns10:Email>
                    </ns10:mGETPHONECUSTOMERDetailType>
                </ns10:gGETPHONECUSTOMERDetailType>
            </GETPHONECUSTOMERType>
        </ns30:GetCustomerPhoneNoResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess:     true,
			expectedError:       false,
			expectedDetail:      true,
			expectedCustomerID:  "1771239173",
			expectedPhoneNumber: "+251911706628",
			expectedEmail:        "",
		},
		{
			name: "Parse response with failure status",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns30:GetCustomerPhoneNoResponse xmlns:ns30="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>Failure</successIndicator>
            </Status>
        </ns30:GetCustomerPhoneNoResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess:  false,
			expectedError:    false,
			expectedDetail:   false,
			expectedMessage:  "API returned failure",
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
        <ns30:GetCustomerPhoneNoResponse xmlns:ns30="http://temenos.com/CBESUPERAPP">
        </ns30:GetCustomerPhoneNoResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess:  false,
			expectedError:     false,
			expectedDetail:   false,
			expectedMessage:   "missing status",
		},
		{
			name: "Parse response with no details",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns30:GetCustomerPhoneNoResponse xmlns:ns30="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>Success</successIndicator>
            </Status>
            <GETPHONECUSTOMERType>
            </GETPHONECUSTOMERType>
        </ns30:GetCustomerPhoneNoResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess:  true,
			expectedError:     false,
			expectedDetail:    false,
			expectedMessage:   "no details found",
		},
		{
			name: "Parse response with empty group",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns30:GetCustomerPhoneNoResponse xmlns:ns30="http://temenos.com/CBESUPERAPP" xmlns:ns10="http://temenos.com/GETPHONECUSTOMER">
            <Status>
                <successIndicator>Success</successIndicator>
            </Status>
            <GETPHONECUSTOMERType>
                <ns10:gGETPHONECUSTOMERDetailType>
                </ns10:gGETPHONECUSTOMERDetailType>
            </GETPHONECUSTOMERType>
        </ns30:GetCustomerPhoneNoResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess:  true,
			expectedError:     false,
			expectedDetail:    false,
			expectedMessage:   "no details found",
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
			expectedSuccess:  false,
			expectedError:    false,
			expectedDetail:   false,
			expectedMessage:  "no details found",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ParsePhoneLookupSOAP(tc.xmlData)

			if tc.expectedError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tc.expectedSuccess, result.Success)

				if tc.expectedDetail {
					assert.NotNil(t, result.Detail)
					if result.Detail != nil {
						assert.Equal(t, tc.expectedCustomerID, result.Detail.CustomerID)
						assert.Equal(t, tc.expectedPhoneNumber, result.Detail.PhoneNumber)
						assert.Equal(t, tc.expectedEmail, result.Detail.Email)
					}
				} else {
					if result != nil {
						assert.Nil(t, result.Detail)
						if tc.expectedMessage != "" {
							assert.Equal(t, tc.expectedMessage, result.Message)
						}
					}
				}
			}
		})
	}
}

func TestParsePhoneLookupSOAP_DetailFields(t *testing.T) {
	xmlData := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns30:GetCustomerPhoneNoResponse xmlns:ns30="http://temenos.com/CBESUPERAPP" xmlns:ns10="http://temenos.com/GETPHONECUSTOMER">
            <Status>
                <successIndicator>Success</successIndicator>
            </Status>
            <GETPHONECUSTOMERType>
                <ns10:gGETPHONECUSTOMERDetailType>
                    <ns10:mGETPHONECUSTOMERDetailType>
                        <ns10:CustomerID>1771239173</ns10:CustomerID>
                        <ns10:PhoneNumber>+251911706628</ns10:PhoneNumber>
                        <ns10:Email>yohhanes@example.com</ns10:Email>
                    </ns10:mGETPHONECUSTOMERDetailType>
                </ns10:gGETPHONECUSTOMERDetailType>
            </GETPHONECUSTOMERType>
        </ns30:GetCustomerPhoneNoResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParsePhoneLookupSOAP(xmlData)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Success)
	assert.NotNil(t, result.Detail)

	if result.Detail != nil {
		assert.Equal(t, "1771239173", result.Detail.CustomerID)
		assert.Equal(t, "+251911706628", result.Detail.PhoneNumber)
		assert.Equal(t, "yohhanes@example.com", result.Detail.Email)
	}
}

func TestParsePhoneLookupSOAP_ExactSuccessIndicator(t *testing.T) {
	xmlData := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns30:GetCustomerPhoneNoResponse xmlns:ns30="http://temenos.com/CBESUPERAPP" xmlns:ns10="http://temenos.com/GETPHONECUSTOMER">
            <Status>
                <successIndicator>SUCCESS</successIndicator>
            </Status>
            <GETPHONECUSTOMERType>
                <ns10:gGETPHONECUSTOMERDetailType>
                    <ns10:mGETPHONECUSTOMERDetailType>
                        <ns10:CustomerID>1771239173</ns10:CustomerID>
                        <ns10:PhoneNumber>+251911706628</ns10:PhoneNumber>
                        <ns10:Email></ns10:Email>
                    </ns10:mGETPHONECUSTOMERDetailType>
                </ns10:gGETPHONECUSTOMERDetailType>
            </GETPHONECUSTOMERType>
        </ns30:GetCustomerPhoneNoResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParsePhoneLookupSOAP(xmlData)
	// Note: The code checks for exact "Success" (case-sensitive), so "SUCCESS" should return failure
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.False(t, result.Success)
	assert.Equal(t, "API returned failure", result.Message)
}

