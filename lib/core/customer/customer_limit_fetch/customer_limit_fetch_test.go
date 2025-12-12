package customerlimitfetch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCustomerLimitFetch(t *testing.T) {
	tests := []struct {
		name   string
		param  Param
		expect []string
	}{
		{
			name: "Validate customer limit fetch XML generation",
			param: Param{
				Username:       "SUPERAPP",
				Password:       "123456",
				CustomerNumber: "1026582446",
			},
			expect: []string{
				`<password>123456</password>`,
				`<userName>SUPERAPP</userName>`,
				`<transactionId>1026582446</transactionId>`,
				`<soapenv:Envelope`,
				`<soapenv:Body>`,
				`<cbes:CustomerLimitView>`,
				`<WebRequestCommon>`,
				`<CUSTOMERLIMITCUSTOMLIMITType>`,
			},
		},
		{
			name: "Validate customer limit fetch with different values",
			param: Param{
				Username:       "TESTUSER",
				Password:       "PASSWORD123",
				CustomerNumber: "1234567890",
			},
			expect: []string{
				`<password>PASSWORD123</password>`,
				`<userName>TESTUSER</userName>`,
				`<transactionId>1234567890</transactionId>`,
			},
		},
		{
			name: "Validate customer limit fetch with empty customer number",
			param: Param{
				Username:       "SUPERAPP",
				Password:       "123456",
				CustomerNumber: "",
			},
			expect: []string{
				`<password>123456</password>`,
				`<userName>SUPERAPP</userName>`,
				`<transactionId></transactionId>`,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			xmlRequest := NewCustomerLimitFetch(tc.param)

			// Validate XML structure
			assert.Contains(t, xmlRequest, "<soapenv:Envelope")
			assert.Contains(t, xmlRequest, "<soapenv:Body>")
			assert.Contains(t, xmlRequest, "<cbes:CustomerLimitView>")

			// Validate all expected strings are present
			for _, expectedStr := range tc.expect {
				assert.Contains(t, xmlRequest, expectedStr)
			}

			// Validate XML is not empty
			assert.NotEmpty(t, xmlRequest)
		})
	}
}

func TestParseCustomerLimitFetchSOAP(t *testing.T) {
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
        <ns30:CustomerLimitViewResponse xmlns:ns29="http://temenos.com/CUSTOMERLIMITCUSTOMLIMIT" xmlns:ns30="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>Success</successIndicator>
                <transactionId>1026582446</transactionId>
                <messageId>MSG123</messageId>
                <application>CUSTOMER.LIMIT</application>
            </Status>
            <CUSTOMERLIMITType>
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
        </ns30:CustomerLimitViewResponse>
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
        <ns30:CustomerLimitViewResponse xmlns:ns30="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>Failure</successIndicator>
            </Status>
        </ns30:CustomerLimitViewResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess: false,
			expectedError:   false,
			expectedDetail:  false,
			expectedMessage: "missing customer limit detail",
		},
		{
			name: "Parse response with failure status but has detail",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns30:CustomerLimitViewResponse xmlns:ns29="http://temenos.com/CUSTOMERLIMITCUSTOMLIMIT" xmlns:ns30="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>Failure</successIndicator>
            </Status>
            <CUSTOMERLIMITType>
                <ns29:gUSERCHANNELTYPE>
                    <ns29:mUSERCHANNELTYPE>
                        <ns29:USERCHANNELTYPE>APP</ns29:USERCHANNELTYPE>
                        <ns29:USERMAXLIMIT>800000</ns29:USERMAXLIMIT>
                    </ns29:mUSERCHANNELTYPE>
                </ns29:gUSERCHANNELTYPE>
            </CUSTOMERLIMITType>
        </ns30:CustomerLimitViewResponse>
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
        <ns30:CustomerLimitViewResponse xmlns:ns30="http://temenos.com/CBESUPERAPP">
        </ns30:CustomerLimitViewResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess: false,
			expectedError:   false,
			expectedDetail:  false,
			expectedMessage: "missing status",
		},
		{
			name: "Parse response with no customer limit detail",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns30:CustomerLimitViewResponse xmlns:ns30="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>Success</successIndicator>
            </Status>
        </ns30:CustomerLimitViewResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess: false,
			expectedError:   false,
			expectedDetail:  false,
			expectedMessage: "missing customer limit detail",
		},
		{
			name: "Parse response with case-insensitive success indicator",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns30:CustomerLimitViewResponse xmlns:ns29="http://temenos.com/CUSTOMERLIMITCUSTOMLIMIT" xmlns:ns30="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>success</successIndicator>
            </Status>
            <CUSTOMERLIMITType>
                <ns29:gUSERCHANNELTYPE>
                    <ns29:mUSERCHANNELTYPE>
                        <ns29:USERCHANNELTYPE>APP</ns29:USERCHANNELTYPE>
                        <ns29:USERMAXLIMIT>800000</ns29:USERMAXLIMIT>
                    </ns29:mUSERCHANNELTYPE>
                </ns29:gUSERCHANNELTYPE>
            </CUSTOMERLIMITType>
        </ns30:CustomerLimitViewResponse>
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
        <OtherResponse xmlns:ns30="http://temenos.com/CBESUPERAPP">
        </OtherResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess: false,
			expectedError:   false,
			expectedDetail:  false,
			expectedMessage: "invalid response type",
		},
		{
			name: "Parse response with uppercase SUCCESS indicator",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns30:CustomerLimitViewResponse xmlns:ns29="http://temenos.com/CUSTOMERLIMITCUSTOMLIMIT" xmlns:ns30="http://temenos.com/CBESUPERAPP">
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
        </ns30:CustomerLimitViewResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess: true,
			expectedError:   false,
			expectedDetail:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ParseCustomerLimitFetchSOAP(tc.xmlData)

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
							assert.Contains(t, result.Message, tc.expectedMessage)
						}
					}
				}
			}
		})
	}
}

func TestParseCustomerLimitFetchSOAP_DetailFields(t *testing.T) {
	xmlData := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns30:CustomerLimitViewResponse xmlns:ns29="http://temenos.com/CUSTOMERLIMITCUSTOMLIMIT" xmlns:ns30="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>Success</successIndicator>
                <transactionId>1026582446</transactionId>
                <messageId>MSG123</messageId>
                <application>CUSTOMER.LIMIT</application>
            </Status>
            <CUSTOMERLIMITType>
                <ns29:gUSERCHANNELTYPE>
                    <ns29:mUSERCHANNELTYPE>
                        <ns29:USERCHANNELTYPE>APP</ns29:USERCHANNELTYPE>
                        <ns29:USERMAXLIMIT>800000</ns29:USERMAXLIMIT>
                    </ns29:mUSERCHANNELTYPE>
                    <ns29:mUSERCHANNELTYPE>
                        <ns29:USERCHANNELTYPE>USSD</ns29:USERCHANNELTYPE>
                        <ns29:USERMAXLIMIT>150000</ns29:USERMAXLIMIT>
                    </ns29:mUSERCHANNELTYPE>
                    <ns29:mUSERCHANNELTYPE>
                        <ns29:USERCHANNELTYPE>ATM</ns29:USERCHANNELTYPE>
                        <ns29:USERMAXLIMIT>500000</ns29:USERMAXLIMIT>
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
        </ns30:CustomerLimitViewResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseCustomerLimitFetchSOAP(xmlData)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Success)
	assert.NotNil(t, result.Detail)

	if result.Detail != nil {
		assert.Equal(t, "1", result.Detail.CurrNo)
		assert.Equal(t, "10731_SUPERAPP.1__OFS_GCS", result.Detail.GlobalInputter.Inputter)
		assert.Equal(t, "2512120054", result.Detail.GlobalDatetime.Datetime)
		assert.Equal(t, "10731_SUPERAPP.1_OFS_GCS", result.Detail.Authoriser)
		assert.Equal(t, "ET0010001", result.Detail.CoCode)
		assert.Equal(t, "1", result.Detail.DeptCode)

		// Check UserChannelType
		assert.NotNil(t, result.Detail.UserChannelType)
		if result.Detail.UserChannelType != nil {
			assert.Len(t, result.Detail.UserChannelType.Details, 3)

			// Check first channel type
			assert.Equal(t, "APP", result.Detail.UserChannelType.Details[0].UserChannelType)
			assert.Equal(t, "800000", result.Detail.UserChannelType.Details[0].UserMaxLimit)

			// Check second channel type
			assert.Equal(t, "USSD", result.Detail.UserChannelType.Details[1].UserChannelType)
			assert.Equal(t, "150000", result.Detail.UserChannelType.Details[1].UserMaxLimit)

			// Check third channel type
			assert.Equal(t, "ATM", result.Detail.UserChannelType.Details[2].UserChannelType)
			assert.Equal(t, "500000", result.Detail.UserChannelType.Details[2].UserMaxLimit)
		}
	}
}

func TestParseCustomerLimitFetchSOAP_EmptyUserChannelType(t *testing.T) {
	xmlData := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns30:CustomerLimitViewResponse xmlns:ns29="http://temenos.com/CUSTOMERLIMITCUSTOMLIMIT" xmlns:ns30="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>Success</successIndicator>
            </Status>
            <CUSTOMERLIMITType>
                <ns29:gUSERCHANNELTYPE>
                </ns29:gUSERCHANNELTYPE>
                <ns29:CURRNO>1</ns29:CURRNO>
            </CUSTOMERLIMITType>
        </ns30:CustomerLimitViewResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseCustomerLimitFetchSOAP(xmlData)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Success)
	assert.NotNil(t, result.Detail)

	if result.Detail != nil {
		// UserChannelType may be nil or have empty details when gUSERCHANNELTYPE is empty
		// Both cases are acceptable
		if result.Detail.UserChannelType != nil {
			// If it exists, details should be empty or nil
			if result.Detail.UserChannelType.Details != nil {
				assert.Len(t, result.Detail.UserChannelType.Details, 0)
			}
		}
		assert.Equal(t, "1", result.Detail.CurrNo)
	}
}

func TestParseCustomerLimitFetchSOAP_NilPointers(t *testing.T) {
	xmlData := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns30:CustomerLimitViewResponse xmlns:ns30="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>Success</successIndicator>
            </Status>
            <CUSTOMERLIMITType>
                <CURRNO>1</CURRNO>
            </CUSTOMERLIMITType>
        </ns30:CustomerLimitViewResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseCustomerLimitFetchSOAP(xmlData)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Success)
	assert.NotNil(t, result.Detail)

	if result.Detail != nil {
		assert.Equal(t, "1", result.Detail.CurrNo)
		// UserChannelType, GlobalInputter, and GlobalDatetime may be nil
		// This is acceptable as they are optional fields
	}
}
