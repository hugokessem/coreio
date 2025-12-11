package customerlimit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCustomerLimit(t *testing.T) {
	tests := []struct {
		name   string
		param  Params
		expect []string
	}{
		{
			name: "Validate customer limit XML generation",
			param: Params{
				Username:      "SUPERAPP",
				Password:      "123456",
				TransactionID: "Global",
			},
			expect: []string{
				`<password>123456</password>`,
				`<userName>SUPERAPP</userName>`,
				`<transactionId>Global</transactionId>`,
				`<soapenv:Envelope`,
				`<soapenv:Body>`,
				`<cbes:CustomerLimitView>`,
				`<WebRequestCommon>`,
				`<CUSTOMERLIMITSETUPType>`,
			},
		},
		{
			name: "Validate customer limit with different values",
			param: Params{
				Username:      "TESTUSER",
				Password:      "PASSWORD123",
				TransactionID: "TXN123456",
			},
			expect: []string{
				`<password>PASSWORD123</password>`,
				`<userName>TESTUSER</userName>`,
				`<transactionId>TXN123456</transactionId>`,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			xmlRequest := NewCustomerLimit(tc.param)

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

func TestParseCustomerLimitSOAP(t *testing.T) {
	tests := []struct {
		name                 string
		xmlData              string
		expectedSuccess      bool
		expectedError        bool
		expectedDetail       bool
		expectedMessage      string
		expectedCustomerMax  string
		expectedCustomerMin  string
		expectedAccountMax   string
		expectedAccountMin   string
	}{
		{
			name: "Parse successful response",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns30:CustomerLimitViewResponse xmlns:ns29="http://temenos.com/CUSTOMERLIMIT" xmlns:ns30="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>Success</successIndicator>
                <transactionId>Global</transactionId>
                <messageId></messageId>
                <application>CUSTOMER.LIMIT</application>
            </Status>
            <CUSTOMERLIMITType>
                <ns29:gCHANNELTYPE>
                    <ns29:mCHANNELTYPE>
                        <ns29:CHANNELTYPE>APP</ns29:CHANNELTYPE>
                        <ns29:CHANNELMAXLIMIT>1000000</ns29:CHANNELMAXLIMIT>
                        <ns29:CHANNELMINLIMIT>0.1</ns29:CHANNELMINLIMIT>
                    </ns29:mCHANNELTYPE>
                </ns29:gCHANNELTYPE>
                <ns29:CUSTMAXLIMIT>1000000</ns29:CUSTMAXLIMIT>
                <ns29:CUSTMINLIMIT>0.1</ns29:CUSTMINLIMIT>
                <ns29:CUSTCOUNT>25</ns29:CUSTCOUNT>
                <ns29:ACCTMAXLIMIT>1000000</ns29:ACCTMAXLIMIT>
                <ns29:ACCTMINLIMIT>0.1</ns29:ACCTMINLIMIT>
                <ns29:ACCTCOUNT>25</ns29:ACCTCOUNT>
                <ns29:CURRNO>1</ns29:CURRNO>
                <ns29:gDATETIME>
                    <ns29:DATETIME>2512051827</ns29:DATETIME>
                </ns29:gDATETIME>
                <ns29:gINPUTTER>
                    <ns29:INPUTTER>5020_SUPERMAKER__OFS_CBELIVETC</ns29:INPUTTER>
                </ns29:gINPUTTER>
                <ns29:AUTHORISER>5020_SUPERMAKER_OFS_CBELIVETC</ns29:AUTHORISER>
                <ns29:COCODE>ET0010001</ns29:COCODE>
                <ns29:DEPTCODE>1</ns29:DEPTCODE>
            </CUSTOMERLIMITType>
        </ns30:CustomerLimitViewResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess:     true,
			expectedError:       false,
			expectedDetail:      true,
			expectedCustomerMax: "1000000",
			expectedCustomerMin: "0.1",
			expectedAccountMax:  "1000000",
			expectedAccountMin:  "0.1",
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
			expectedMessage: "Missing Status",
		},
		{
			name: "Parse response with no details",
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
			expectedMessage: "No details found",
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
			expectedError:   true,
			expectedDetail:  false,
		},
		{
			name: "Parse response with channel types",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns30:CustomerLimitViewResponse xmlns:ns29="http://temenos.com/CUSTOMERLIMIT" xmlns:ns30="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>Success</successIndicator>
            </Status>
            <CUSTOMERLIMITType>
                <ns29:gCHANNELTYPE>
                    <ns29:mCHANNELTYPE>
                        <ns29:CHANNELTYPE>APP</ns29:CHANNELTYPE>
                        <ns29:CHANNELMAXLIMIT>1000000</ns29:CHANNELMAXLIMIT>
                        <ns29:CHANNELMINLIMIT>0.1</ns29:CHANNELMINLIMIT>
                    </ns29:mCHANNELTYPE>
                    <ns29:mCHANNELTYPE>
                        <ns29:CHANNELTYPE>USSD</ns29:CHANNELTYPE>
                        <ns29:CHANNELMAXLIMIT>200000</ns29:CHANNELMAXLIMIT>
                        <ns29:CHANNELMINLIMIT>0.1</ns29:CHANNELMINLIMIT>
                    </ns29:mCHANNELTYPE>
                </ns29:gCHANNELTYPE>
                <ns29:CUSTMAXLIMIT>1000000</ns29:CUSTMAXLIMIT>
                <ns29:CUSTMINLIMIT>0.1</ns29:CUSTMINLIMIT>
                <ns29:CUSTCOUNT>25</ns29:CUSTCOUNT>
                <ns29:ACCTMAXLIMIT>1000000</ns29:ACCTMAXLIMIT>
                <ns29:ACCTMINLIMIT>0.1</ns29:ACCTMINLIMIT>
                <ns29:ACCTCOUNT>25</ns29:ACCTCOUNT>
            </CUSTOMERLIMITType>
        </ns30:CustomerLimitViewResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess:     true,
			expectedError:       false,
			expectedDetail:      true,
			expectedCustomerMax: "1000000",
			expectedCustomerMin: "0.1",
			expectedAccountMax:  "1000000",
			expectedAccountMin:  "0.1",
		},
		{
			name: "Parse response with conv role types",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns30:CustomerLimitViewResponse xmlns:ns29="http://temenos.com/CUSTOMERLIMIT" xmlns:ns30="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>Success</successIndicator>
            </Status>
            <CUSTOMERLIMITType>
                <ns29:gCONVROLETYPE>
                    <ns29:mCONVROLETYPE>
                        <ns29:CONVROLETYPE>Mass</ns29:CONVROLETYPE>
                        <ns29:CONVROLEMAXLIMIT>600000</ns29:CONVROLEMAXLIMIT>
                        <ns29:CONVROLEMINLIMIT>0.1</ns29:CONVROLEMINLIMIT>
                    </ns29:mCONVROLETYPE>
                    <ns29:mCONVROLETYPE>
                        <ns29:CONVROLETYPE>Middle</ns29:CONVROLETYPE>
                        <ns29:CONVROLEMAXLIMIT>800000</ns29:CONVROLEMAXLIMIT>
                        <ns29:CONVROLEMINLIMIT>0.1</ns29:CONVROLEMINLIMIT>
                    </ns29:mCONVROLETYPE>
                </ns29:gCONVROLETYPE>
                <ns29:CUSTMAXLIMIT>1000000</ns29:CUSTMAXLIMIT>
                <ns29:CUSTMINLIMIT>0.1</ns29:CUSTMINLIMIT>
                <ns29:ACCTMAXLIMIT>1000000</ns29:ACCTMAXLIMIT>
                <ns29:ACCTMINLIMIT>0.1</ns29:ACCTMINLIMIT>
            </CUSTOMERLIMITType>
        </ns30:CustomerLimitViewResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess:     true,
			expectedError:       false,
			expectedDetail:      true,
			expectedCustomerMax: "1000000",
			expectedCustomerMin: "0.1",
			expectedAccountMax:  "1000000",
			expectedAccountMin:  "0.1",
		},
		{
			name: "Parse response with ifb role types",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns30:CustomerLimitViewResponse xmlns:ns29="http://temenos.com/CUSTOMERLIMIT" xmlns:ns30="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>Success</successIndicator>
            </Status>
            <CUSTOMERLIMITType>
                <ns29:gIFBROLETYPE>
                    <ns29:mIFBROLETYPE>
                        <ns29:IFBROLETYPE>Mass</ns29:IFBROLETYPE>
                        <ns29:IFBROLEMAXLIMIT>600000</ns29:IFBROLEMAXLIMIT>
                        <ns29:IFBROLEMINLIMIT>0.1</ns29:IFBROLEMINLIMIT>
                    </ns29:mIFBROLETYPE>
                </ns29:gIFBROLETYPE>
                <ns29:CUSTMAXLIMIT>1000000</ns29:CUSTMAXLIMIT>
                <ns29:CUSTMINLIMIT>0.1</ns29:CUSTMINLIMIT>
                <ns29:ACCTMAXLIMIT>1000000</ns29:ACCTMAXLIMIT>
                <ns29:ACCTMINLIMIT>0.1</ns29:ACCTMINLIMIT>
            </CUSTOMERLIMITType>
        </ns30:CustomerLimitViewResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess:     true,
			expectedError:       false,
			expectedDetail:      true,
			expectedCustomerMax: "1000000",
			expectedCustomerMin: "0.1",
			expectedAccountMax:  "1000000",
			expectedAccountMin:  "0.1",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ParseCustomerLimitSOAP(tc.xmlData)

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
						assert.Equal(t, tc.expectedCustomerMax, result.Detail.CustomerMaxLimit)
						assert.Equal(t, tc.expectedCustomerMin, result.Detail.CustomerMinLimit)
						assert.Equal(t, tc.expectedAccountMax, result.Detail.AccountMaxLimit)
						assert.Equal(t, tc.expectedAccountMin, result.Detail.AccountMinLimit)
					}
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

func TestParseCustomerLimitSOAP_DetailFields(t *testing.T) {
	xmlData := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns30:CustomerLimitViewResponse xmlns:ns29="http://temenos.com/CUSTOMERLIMIT" xmlns:ns30="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>Success</successIndicator>
                <transactionId>Global</transactionId>
                <messageId>MSG123</messageId>
                <application>CUSTOMER.LIMIT</application>
            </Status>
            <CUSTOMERLIMITType>
                <ns29:gCHANNELTYPE>
                    <ns29:mCHANNELTYPE>
                        <ns29:CHANNELTYPE>APP</ns29:CHANNELTYPE>
                        <ns29:CHANNELMAXLIMIT>1000000</ns29:CHANNELMAXLIMIT>
                        <ns29:CHANNELMINLIMIT>0.1</ns29:CHANNELMINLIMIT>
                    </ns29:mCHANNELTYPE>
                </ns29:gCHANNELTYPE>
                <ns29:gCONVROLETYPE>
                    <ns29:mCONVROLETYPE>
                        <ns29:CONVROLETYPE>Mass</ns29:CONVROLETYPE>
                        <ns29:CONVROLEMAXLIMIT>600000</ns29:CONVROLEMAXLIMIT>
                        <ns29:CONVROLEMINLIMIT>0.1</ns29:CONVROLEMINLIMIT>
                    </ns29:mCONVROLETYPE>
                </ns29:gCONVROLETYPE>
                <ns29:gIFBROLETYPE>
                    <ns29:mIFBROLETYPE>
                        <ns29:IFBROLETYPE>Affluent</ns29:IFBROLETYPE>
                        <ns29:IFBROLEMAXLIMIT>1000000</ns29:IFBROLEMAXLIMIT>
                        <ns29:IFBROLEMINLIMIT>0.1</ns29:IFBROLEMINLIMIT>
                    </ns29:mIFBROLETYPE>
                </ns29:gIFBROLETYPE>
                <ns29:CUSTMAXLIMIT>1000000</ns29:CUSTMAXLIMIT>
                <ns29:CUSTMINLIMIT>0.1</ns29:CUSTMINLIMIT>
                <ns29:CUSTCOUNT>25</ns29:CUSTCOUNT>
                <ns29:ACCTMAXLIMIT>1000000</ns29:ACCTMAXLIMIT>
                <ns29:ACCTMINLIMIT>0.1</ns29:ACCTMINLIMIT>
                <ns29:ACCTCOUNT>25</ns29:ACCTCOUNT>
                <ns29:CURRNO>1</ns29:CURRNO>
                <ns29:gDATETIME>
                    <ns29:DATETIME>2512051827</ns29:DATETIME>
                </ns29:gDATETIME>
                <ns29:gINPUTTER>
                    <ns29:INPUTTER>5020_SUPERMAKER__OFS_CBELIVETC</ns29:INPUTTER>
                </ns29:gINPUTTER>
                <ns29:AUTHORISER>5020_SUPERMAKER_OFS_CBELIVETC</ns29:AUTHORISER>
                <ns29:COCODE>ET0010001</ns29:COCODE>
                <ns29:DEPTCODE>1</ns29:DEPTCODE>
            </CUSTOMERLIMITType>
        </ns30:CustomerLimitViewResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseCustomerLimitSOAP(xmlData)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Success)
	assert.NotNil(t, result.Detail)

	if result.Detail != nil {
		assert.Equal(t, "1000000", result.Detail.CustomerMaxLimit)
		assert.Equal(t, "0.1", result.Detail.CustomerMinLimit)
		assert.Equal(t, "25", result.Detail.CustomerCount)
		assert.Equal(t, "1000000", result.Detail.AccountMaxLimit)
		assert.Equal(t, "0.1", result.Detail.AccountMinLimit)
		assert.Equal(t, "25", result.Detail.AccountCount)
		assert.Equal(t, "1", result.Detail.CurrencyNo)
		assert.Equal(t, "2512051827", result.Detail.GlobalDateTime.DateTime)
		assert.Equal(t, "5020_SUPERMAKER__OFS_CBELIVETC", result.Detail.GlobalInputter.Inputter)
		assert.Equal(t, "5020_SUPERMAKER_OFS_CBELIVETC", result.Detail.Authoriser)
		assert.Equal(t, "ET0010001", result.Detail.CoCode)
		assert.Equal(t, "1", result.Detail.DeptCode)

		// Check ChannelTypeGroup
		assert.NotNil(t, result.Detail.ChannelTypeGroup)
		if result.Detail.ChannelTypeGroup != nil {
			assert.Len(t, result.Detail.ChannelTypeGroup.Details, 1)
			assert.Equal(t, "APP", result.Detail.ChannelTypeGroup.Details[0].ChannelType)
			assert.Equal(t, "1000000", result.Detail.ChannelTypeGroup.Details[0].ChannelMaxLimit)
			assert.Equal(t, "0.1", result.Detail.ChannelTypeGroup.Details[0].ChannelMinLimit)
		}

		// Check ConvRoleTypeGroup
		assert.NotNil(t, result.Detail.ConvRoleTypeGroup)
		if result.Detail.ConvRoleTypeGroup != nil {
			assert.Len(t, result.Detail.ConvRoleTypeGroup.Details, 1)
			assert.Equal(t, "Mass", result.Detail.ConvRoleTypeGroup.Details[0].ConvRoleType)
			assert.Equal(t, "600000", result.Detail.ConvRoleTypeGroup.Details[0].ConvRoleMaxLimit)
			assert.Equal(t, "0.1", result.Detail.ConvRoleTypeGroup.Details[0].ConvRoleMinLimit)
		}

		// Check IfbRoleTypeGroup
		assert.NotNil(t, result.Detail.IfbRoleTypeGroup)
		if result.Detail.IfbRoleTypeGroup != nil {
			assert.Len(t, result.Detail.IfbRoleTypeGroup.Details, 1)
			assert.Equal(t, "Affluent", result.Detail.IfbRoleTypeGroup.Details[0].IfbRoleType)
			assert.Equal(t, "1000000", result.Detail.IfbRoleTypeGroup.Details[0].IfbRoleMaxLimit)
			assert.Equal(t, "0.1", result.Detail.IfbRoleTypeGroup.Details[0].IfbRoleMinLimit)
		}
	}
}

func TestParseCustomerLimitSOAP_ExactSuccessIndicator(t *testing.T) {
	xmlData := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns30:CustomerLimitViewResponse xmlns:ns30="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>SUCCESS</successIndicator>
            </Status>
            <CUSTOMERLIMITType>
                <CUSTMAXLIMIT>1000000</CUSTMAXLIMIT>
                <CUSTMINLIMIT>0.1</CUSTMINLIMIT>
            </CUSTOMERLIMITType>
        </ns30:CustomerLimitViewResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseCustomerLimitSOAP(xmlData)
	// Note: The code checks for exact "Success" (case-sensitive), so "SUCCESS" should return failure
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.False(t, result.Success)
	assert.NotEmpty(t, result.Message)
	if len(result.Message) > 0 {
		assert.Contains(t, result.Message[0], "API returned failure")
	}
}

