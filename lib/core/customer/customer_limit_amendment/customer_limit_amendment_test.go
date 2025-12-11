package customerlimitamendment

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCustomerLimitAmendment(t *testing.T) {
	tests := []struct {
		name   string
		param  Params
		expect []string
	}{
		{
			name: "Validate customer limit amendment XML generation",
			param: Params{
				Username:         "SUPERAPP",
				Password:         "123456",
				CustomerID:       "1026582446",
				AppUserMaxLimit:  "800000",
				USSDUserMaxLimit: "150000",
			},
			expect: []string{
				`<password>123456</password>`,
				`<userName>SUPERAPP</userName>`,
				`<soapenv:Envelope`,
				`<soapenv:Body>`,
				`<cbes:CustomerLimitAmendment>`,
				`<WebRequestCommon>`,
				`<CUSTOMERLIMITCUSTOMLIMITType id="1026582446">`,
				`<cus:USERCHANNELTYPE>APP</cus:USERCHANNELTYPE>`,
				`<cus:USERMAXLIMIT>800000</cus:USERMAXLIMIT>`,
				`<cus:USERCHANNELTYPE>USSD</cus:USERCHANNELTYPE>`,
				`<cus:USERMAXLIMIT>150000</cus:USERMAXLIMIT>`,
			},
		},
		{
			name: "Validate customer limit amendment with different values",
			param: Params{
				Username:         "TESTUSER",
				Password:         "PASSWORD123",
				CustomerID:       "1234567890",
				AppUserMaxLimit:  "500000",
				USSDUserMaxLimit: "100000",
			},
			expect: []string{
				`<password>PASSWORD123</password>`,
				`<userName>TESTUSER</userName>`,
				`id="1234567890"`,
				`<cus:USERMAXLIMIT>500000</cus:USERMAXLIMIT>`,
				`<cus:USERMAXLIMIT>100000</cus:USERMAXLIMIT>`,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			xmlRequest := NewCustomerLimitAmendment(tc.param)

			// Validate XML structure
			assert.Contains(t, xmlRequest, "<soapenv:Envelope")
			assert.Contains(t, xmlRequest, "<soapenv:Body>")
			assert.Contains(t, xmlRequest, "<cbes:CustomerLimitAmendment>")

			// Validate all expected strings are present
			for _, expectedStr := range tc.expect {
				assert.Contains(t, xmlRequest, expectedStr)
			}

			// Validate XML is not empty
			assert.NotEmpty(t, xmlRequest)
		})
	}
}

func TestParseCustomerLimitAmendmentSOAP(t *testing.T) {
	tests := []struct {
		name            string
		xmlData         string
		expectedSuccess bool
		expectedError   bool
		expectedDetail  bool
		expectedMessage string
	}{
		{
			name: "Parse successful response",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns32:CustomerLimitAmendmentResponse xmlns:ns29="http://temenos.com/CUSTOMERLIMITCUSTOMLIMIT" xmlns:ns30="http://temenos.com/CUSTOMERLIMIT" xmlns:ns32="http://temenos.com/CBESUPERAPP">
            <Status>
                <transactionId>1026582446</transactionId>
                <messageId></messageId>
                <successIndicator>Success</successIndicator>
                <application>CUSTOMER.LIMIT</application>
            </Status>
            <CUSTOMERLIMITType id="1026582446">
                <ns29:gUSERCHANNELTYPE>
                    <ns29:mUSERCHANNELTYPE>
                        <ns29:USERCHANNELTYPE>APP</ns29:USERCHANNELTYPE>
                        <ns29:USERMAXLIMIT>800000</ns29:USERMAXLIMIT>
                    </ns29:mUSERCHANNELTYPE>
                    <ns29:mUSERCHANNELTYPE>
                        <ns29:USERCHANNELTYPE>USSD</ns29:USERCHANNELTYPE>
                        <ns29:USERMAXLIMIT>150000</ns29:USERMAXLIMIT>
                    </ns29:mUSERCHANNELTYPE>
                </ns29:gUSERCHANNELTYPE>
                <ns29:CURRNO>1</ns29:CURRNO>
                <ns29:gINPUTTER>
                    <ns29:INPUTTER>10731_SUPERAPP.1__OFS_GCS</ns29:INPUTTER>
                </ns29:gINPUTTER>
                <ns29:gDATETIME>
                    <ns29:DATETIME>2512120054</ns29:DATETIME>
                </ns29:gDATETIME>
                <ns29:AUTHORISER>10731_SUPERAPP.1_OFS_GCS</ns29:AUTHORISER>
                <ns29:COCODE>ET0010001</ns29:COCODE>
                <ns29:DEPTCODE>1</ns29:DEPTCODE>
            </CUSTOMERLIMITType>
        </ns32:CustomerLimitAmendmentResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess: true,
			expectedError:   false,
			expectedDetail:  true,
		},
		{
			name: "Parse response with failure status",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns32:CustomerLimitAmendmentResponse xmlns:ns32="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>Failure</successIndicator>
            </Status>
        </ns32:CustomerLimitAmendmentResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess: false,
			expectedError:   false,
			expectedDetail:  false,
			expectedMessage: "API returned failure",
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
        <ns32:CustomerLimitAmendmentResponse xmlns:ns32="http://temenos.com/CBESUPERAPP">
        </ns32:CustomerLimitAmendmentResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess: false,
			expectedError:   false,
			expectedDetail:  false,
			expectedMessage: "Missing Status",
		},
		{
			name: "Parse response with no customer limit type",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns32:CustomerLimitAmendmentResponse xmlns:ns32="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>Success</successIndicator>
            </Status>
        </ns32:CustomerLimitAmendmentResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess: false,
			expectedError:   false,
			expectedDetail:  false,
		},
		{
			name: "Parse response with case-insensitive success indicator",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns32:CustomerLimitAmendmentResponse xmlns:ns29="http://temenos.com/CUSTOMERLIMITCUSTOMLIMIT" xmlns:ns32="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>SUCCESS</successIndicator>
            </Status>
            <CUSTOMERLIMITType>
                <ns29:gUSERCHANNELTYPE>
                    <ns29:mUSERCHANNELTYPE>
                        <ns29:USERCHANNELTYPE>APP</ns29:USERCHANNELTYPE>
                        <ns29:USERMAXLIMIT>800000</ns29:USERMAXLIMIT>
                    </ns29:mUSERCHANNELTYPE>
                </ns29:gUSERCHANNELTYPE>
            </CUSTOMERLIMITType>
        </ns32:CustomerLimitAmendmentResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess: true,
			expectedError:   false,
			expectedDetail:  true,
		},
		{
			name: "Parse response with invalid response structure",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <OtherResponse xmlns:ns32="http://temenos.com/CBESUPERAPP">
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
			result, err := ParseCustomerLimitAmendmentSOAP(tc.xmlData)

			if tc.expectedError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tc.expectedSuccess, result.Success)

				if tc.expectedDetail {
					assert.NotNil(t, result.Detail)
				} else {
					if result != nil {
						if tc.expectedMessage != "" {
							assert.NotEmpty(t, result.Message)
							if len(result.Message) > 0 {
								assert.Contains(t, result.Message[0], tc.expectedMessage)
							}
						}
					}
				}
			}
		})
	}
}

func TestParseCustomerLimitAmendmentSOAP_DetailFields(t *testing.T) {
	xmlData := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns32:CustomerLimitAmendmentResponse xmlns:ns29="http://temenos.com/CUSTOMERLIMITCUSTOMLIMIT" xmlns:ns30="http://temenos.com/CUSTOMERLIMIT" xmlns:ns32="http://temenos.com/CBESUPERAPP">
            <Status>
                <transactionId>1026582446</transactionId>
                <messageId>MSG123</messageId>
                <successIndicator>Success</successIndicator>
                <application>CUSTOMER.LIMIT</application>
            </Status>
            <CUSTOMERLIMITType id="1026582446">
                <ns29:gUSERCHANNELTYPE>
                    <ns29:mUSERCHANNELTYPE>
                        <ns29:USERCHANNELTYPE>APP</ns29:USERCHANNELTYPE>
                        <ns29:USERMAXLIMIT>800000</ns29:USERMAXLIMIT>
                    </ns29:mUSERCHANNELTYPE>
                    <ns29:mUSERCHANNELTYPE>
                        <ns29:USERCHANNELTYPE>USSD</ns29:USERCHANNELTYPE>
                        <ns29:USERMAXLIMIT>150000</ns29:USERMAXLIMIT>
                    </ns29:mUSERCHANNELTYPE>
                </ns29:gUSERCHANNELTYPE>
                <ns29:CURRNO>1</ns29:CURRNO>
                <ns29:gINPUTTER>
                    <ns29:INPUTTER>10731_SUPERAPP.1__OFS_GCS</ns29:INPUTTER>
                </ns29:gINPUTTER>
                <ns29:gDATETIME>
                    <ns29:DATETIME>2512120054</ns29:DATETIME>
                </ns29:gDATETIME>
                <ns29:AUTHORISER>10731_SUPERAPP.1_OFS_GCS</ns29:AUTHORISER>
                <ns29:COCODE>ET0010001</ns29:COCODE>
                <ns29:DEPTCODE>1</ns29:DEPTCODE>
            </CUSTOMERLIMITType>
        </ns32:CustomerLimitAmendmentResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseCustomerLimitAmendmentSOAP(xmlData)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Success)
	assert.NotNil(t, result.Detail)

	if result.Detail != nil {
		// Check UserChannelType array
		assert.NotNil(t, result.Detail.UserChannelType)
		if len(result.Detail.UserChannelType) > 0 {
			// Check first channel type
			assert.Equal(t, "APP", result.Detail.UserChannelType[0].UserChannelType)
			assert.Equal(t, "800000", result.Detail.UserChannelType[0].UserMaxLimit)

			// Check second channel type if present
			if len(result.Detail.UserChannelType) > 1 {
				assert.Equal(t, "USSD", result.Detail.UserChannelType[1].UserChannelType)
				assert.Equal(t, "150000", result.Detail.UserChannelType[1].UserMaxLimit)
			}
		}
	}
}

func TestParseCustomerLimitAmendmentSOAP_EmptyUserChannelType(t *testing.T) {
	xmlData := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns32:CustomerLimitAmendmentResponse xmlns:ns29="http://temenos.com/CUSTOMERLIMITCUSTOMLIMIT" xmlns:ns32="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>Success</successIndicator>
            </Status>
            <CUSTOMERLIMITType>
                <ns29:gUSERCHANNELTYPE>
                </ns29:gUSERCHANNELTYPE>
            </CUSTOMERLIMITType>
        </ns32:CustomerLimitAmendmentResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseCustomerLimitAmendmentSOAP(xmlData)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Success)
	assert.NotNil(t, result.Detail)

	if result.Detail != nil {
		// UserChannelType should be empty or nil
		// It's okay if it's empty when gUSERCHANNELTYPE is empty
		if result.Detail.UserChannelType != nil {
			assert.Equal(t, 0, len(result.Detail.UserChannelType))
		}
	}
}

