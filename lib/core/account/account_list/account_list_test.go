package accountlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAccountList(t *testing.T) {
	tests := []struct {
		name   string
		param  Params
		expect []string
	}{
		{
			name: "Validate account list XML generation",
			param: Params{
				Username:      "SUPERAPP",
				Password:      "123456",
				ColumnName:    "CUS.ID",
				CriteriaValue: "1000000015",
			},
			expect: []string{
				`<password>123456</password>`,
				`<userName>SUPERAPP</userName>`,
				`<columnName>CUS.ID</columnName>`,
				`<criteriaValue>1000000015</criteriaValue>`,
				`<soapenv:Envelope`,
				`<soapenv:Body>`,
				`<cbes:AccountListByCIF>`,
				`<operand>EQ</operand>`,
			},
		},
		{
			name: "Validate account list with different values",
			param: Params{
				Username:      "TESTUSER",
				Password:      "PASSWORD123",
				ColumnName:    "ACCOUNT",
				CriteriaValue: "1000000006924",
			},
			expect: []string{
				`<password>PASSWORD123</password>`,
				`<userName>TESTUSER</userName>`,
				`<columnName>ACCOUNT</columnName>`,
				`<criteriaValue>1000000006924</criteriaValue>`,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			xmlRequest := NewAccountList(tc.param)

			// Validate XML structure
			assert.Contains(t, xmlRequest, "<soapenv:Envelope")
			assert.Contains(t, xmlRequest, "<soapenv:Body>")
			assert.Contains(t, xmlRequest, "<cbes:AccountListByCIF>")

			// Validate all expected strings are present
			for _, expectedStr := range tc.expect {
				assert.Contains(t, xmlRequest, expectedStr)
			}

			// Validate XML is not empty
			assert.NotEmpty(t, xmlRequest)

			// Validate that CUS.ID enquiryInputCollection is always present
			assert.Contains(t, xmlRequest, `<columnName>CUS.ID</columnName>`)
		})
	}
}

func TestParseAccountListSOAP(t *testing.T) {
	tests := []struct {
		name            string
		xmlData          string
		expectedSuccess  bool
		expectedError    bool
		expectedDetails  bool
		expectedCount    int
		expectedMessages []string
	}{
		{
			name: "Parse successful response with multiple accounts",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns30:AccountListByCIFResponse xmlns:ns30="http://temenos.com/CBESUPERAPP" xmlns:ns24="http://temenos.com/ACCOUNTINFOSUPERAPP">
            <Status>
                <successIndicator>Success</successIndicator>
            </Status>
            <ACCOUNTINFOSUPERAPPType>
                <ns24:gACCOUNTINFOSUPERAPPDetailType>
                    <ns24:mACCOUNTINFOSUPERAPPDetailType>
                        <ns24:AccountNumber>1000446094649</ns24:AccountNumber>
                        <ns24:CustomerName>UNITED BANK-ECX</ns24:CustomerName>
                        <ns24:Restriction>NO</ns24:Restriction>
                        <ns24:Currency>ETB</ns24:Currency>
                        <ns24:CustomerID>1000000015</ns24:CustomerID>
                        <ns24:Category>1000</ns24:Category>
                        <ns24:AccountType>CURRENT ACCOUNT</ns24:AccountType>
                        <ns24:BranchCode>ET0010222</ns24:BranchCode>
                        <ns24:BranchName>Addis Ababa Branch</ns24:BranchName>
                        <ns24:DistrictName>ARADA</ns24:DistrictName>
                        <ns24:PhoneNo>+251911540610</ns24:PhoneNo>
                        <ns24:Industry>3002</ns24:Industry>
                        <ns24:Sector>3001</ns24:Sector>
                        <ns24:Ownership>3000</ns24:Ownership>
                        <ns24:CustomerSegment>CORPORATE</ns24:CustomerSegment>
                        <ns24:Target>1</ns24:Target>
                    </ns24:mACCOUNTINFOSUPERAPPDetailType>
                    <ns24:mACCOUNTINFOSUPERAPPDetailType>
                        <ns24:AccountNumber>1000446094673</ns24:AccountNumber>
                        <ns24:CustomerName>UNITED BANK-ECX</ns24:CustomerName>
                        <ns24:Restriction>NO</ns24:Restriction>
                        <ns24:Currency>ETB</ns24:Currency>
                        <ns24:CustomerID>1000000015</ns24:CustomerID>
                        <ns24:Category>1000</ns24:Category>
                        <ns24:AccountType>SAVINGS ACCOUNT</ns24:AccountType>
                        <ns24:BranchCode>ET0010001</ns24:BranchCode>
                        <ns24:BranchName>CBE-HEAD OFFICE</ns24:BranchName>
                        <ns24:DistrictName></ns24:DistrictName>
                        <ns24:PhoneNo>+251911540610</ns24:PhoneNo>
                        <ns24:Industry>3002</ns24:Industry>
                        <ns24:Sector>3001</ns24:Sector>
                        <ns24:Ownership>3000</ns24:Ownership>
                        <ns24:CustomerSegment>CORPORATE</ns24:CustomerSegment>
                        <ns24:Target>1</ns24:Target>
                    </ns24:mACCOUNTINFOSUPERAPPDetailType>
                </ns24:gACCOUNTINFOSUPERAPPDetailType>
            </ACCOUNTINFOSUPERAPPType>
        </ns30:AccountListByCIFResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess: true,
			expectedError:   false,
			expectedDetails: true,
			expectedCount:   2,
		},
		{
			name: "Parse response with failure status",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns30:AccountListByCIFResponse xmlns:ns30="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>Failure</successIndicator>
            </Status>
        </ns30:AccountListByCIFResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess:  false,
			expectedError:    false,
			expectedDetails:  false,
			expectedCount:    0,
			expectedMessages: []string{"API returned failure"},
		},
		{
			name: "Parse invalid XML",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<InvalidXML>
    <Broken>
</InvalidXML>`,
			expectedSuccess: false,
			expectedError:   true,
			expectedDetails: false,
			expectedCount:   0,
		},
		{
			name: "Parse response without Status",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns30:AccountListByCIFResponse xmlns:ns30="http://temenos.com/CBESUPERAPP">
        </ns30:AccountListByCIFResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess:  false,
			expectedError:    false,
			expectedDetails:  false,
			expectedCount:    0,
			expectedMessages: []string{"Missing Status"},
		},
		{
			name: "Parse response with no account details",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns30:AccountListByCIFResponse xmlns:ns30="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>Success</successIndicator>
            </Status>
            <ACCOUNTINFOSUPERAPPType>
            </ACCOUNTINFOSUPERAPPType>
        </ns30:AccountListByCIFResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess:  true,
			expectedError:    false,
			expectedDetails:  false,
			expectedCount:    0,
			expectedMessages: []string{"No account details found"},
		},
		{
			name: "Parse response with empty group",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns30:AccountListByCIFResponse xmlns:ns30="http://temenos.com/CBESUPERAPP" xmlns:ns24="http://temenos.com/ACCOUNTINFOSUPERAPP">
            <Status>
                <successIndicator>Success</successIndicator>
            </Status>
            <ACCOUNTINFOSUPERAPPType>
                <ns24:gACCOUNTINFOSUPERAPPDetailType>
                </ns24:gACCOUNTINFOSUPERAPPDetailType>
            </ACCOUNTINFOSUPERAPPType>
        </ns30:AccountListByCIFResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess:  true,
			expectedError:    false,
			expectedDetails:  false,
			expectedCount:    0,
			expectedMessages: []string{"No account details found"},
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
			expectedDetails:  false,
			expectedCount:    0,
			expectedMessages: []string{"Invalid response format"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ParseAccountListSOAP(tc.xmlData)

			if tc.expectedError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tc.expectedSuccess, result.Success)
				assert.Equal(t, tc.expectedCount, len(result.Details))

				if tc.expectedDetails {
					assert.NotNil(t, result.Details)
					assert.Greater(t, len(result.Details), 0)
				} else {
					if result != nil {
						assert.Equal(t, 0, len(result.Details))
					}
				}

				if len(tc.expectedMessages) > 0 {
					assert.Equal(t, tc.expectedMessages, result.Message)
				}
			}
		})
	}
}

func TestParseAccountListSOAP_DetailFields(t *testing.T) {
	xmlData := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns30:AccountListByCIFResponse xmlns:ns30="http://temenos.com/CBESUPERAPP" xmlns:ns24="http://temenos.com/ACCOUNTINFOSUPERAPP">
            <Status>
                <successIndicator>Success</successIndicator>
            </Status>
            <ACCOUNTINFOSUPERAPPType>
                <ns24:gACCOUNTINFOSUPERAPPDetailType>
                    <ns24:mACCOUNTINFOSUPERAPPDetailType>
                        <ns24:AccountNumber>1000000938738</ns24:AccountNumber>
                        <ns24:CustomerName>UNITED BANK-ECX</ns24:CustomerName>
                        <ns24:Restriction>NO</ns24:Restriction>
                        <ns24:Currency>ETB</ns24:Currency>
                        <ns24:CustomerID>1000000015</ns24:CustomerID>
                        <ns24:Category>1001</ns24:Category>
                        <ns24:AccountType>CURRENT ACCOUNT</ns24:AccountType>
                        <ns24:BranchCode>ET0010222</ns24:BranchCode>
                        <ns24:BranchName>Addis Ababa Branch</ns24:BranchName>
                        <ns24:DistrictName>ARADA</ns24:DistrictName>
                        <ns24:PhoneNo>+251911540610</ns24:PhoneNo>
                        <ns24:Industry>3002</ns24:Industry>
                        <ns24:Sector>3001</ns24:Sector>
                        <ns24:Ownership>3000</ns24:Ownership>
                        <ns24:CustomerSegment>CORPORATE</ns24:CustomerSegment>
                        <ns24:Target>1</ns24:Target>
                    </ns24:mACCOUNTINFOSUPERAPPDetailType>
                </ns24:gACCOUNTINFOSUPERAPPDetailType>
            </ACCOUNTINFOSUPERAPPType>
        </ns30:AccountListByCIFResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseAccountListSOAP(xmlData)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Success)
	assert.NotNil(t, result.Details)
	assert.Equal(t, 1, len(result.Details))

	if len(result.Details) > 0 {
		detail := result.Details[0]
		assert.Equal(t, "1000000938738", detail.AccountNumber)
		assert.Equal(t, "UNITED BANK-ECX", detail.CustomerName)
		assert.Equal(t, "NO", detail.Restriction)
		assert.Equal(t, "ETB", detail.Currency)
		assert.Equal(t, "1000000015", detail.CustomerID)
		assert.Equal(t, "1001", detail.Category)
		assert.Equal(t, "CURRENT ACCOUNT", detail.AccountType)
		assert.Equal(t, "ET0010222", detail.BranchCode)
		assert.Equal(t, "Addis Ababa Branch", detail.BranchName)
		assert.Equal(t, "ARADA", detail.DistrictName)
		assert.Equal(t, "+251911540610", detail.PhoneNo)
		assert.Equal(t, "3002", detail.Industry)
		assert.Equal(t, "3001", detail.Sector)
		assert.Equal(t, "3000", detail.Ownership)
		assert.Equal(t, "CORPORATE", detail.CustomerSegment)
		assert.Equal(t, "1", detail.Target)
	}
}

func TestParseAccountListSOAP_MultipleAccounts(t *testing.T) {
	xmlData := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns30:AccountListByCIFResponse xmlns:ns30="http://temenos.com/CBESUPERAPP" xmlns:ns24="http://temenos.com/ACCOUNTINFOSUPERAPP">
            <Status>
                <successIndicator>Success</successIndicator>
            </Status>
            <ACCOUNTINFOSUPERAPPType>
                <ns24:gACCOUNTINFOSUPERAPPDetailType>
                    <ns24:mACCOUNTINFOSUPERAPPDetailType>
                        <ns24:AccountNumber>1000446094649</ns24:AccountNumber>
                        <ns24:CustomerName>UNITED BANK-ECX</ns24:CustomerName>
                        <ns24:Currency>ETB</ns24:Currency>
                        <ns24:CustomerID>1000000015</ns24:CustomerID>
                    </ns24:mACCOUNTINFOSUPERAPPDetailType>
                    <ns24:mACCOUNTINFOSUPERAPPDetailType>
                        <ns24:AccountNumber>1000446094673</ns24:AccountNumber>
                        <ns24:CustomerName>UNITED BANK-ECX</ns24:CustomerName>
                        <ns24:Currency>ETB</ns24:Currency>
                        <ns24:CustomerID>1000000015</ns24:CustomerID>
                    </ns24:mACCOUNTINFOSUPERAPPDetailType>
                    <ns24:mACCOUNTINFOSUPERAPPDetailType>
                        <ns24:AccountNumber>1000446094878</ns24:AccountNumber>
                        <ns24:CustomerName>UNITED BANK-ECX</ns24:CustomerName>
                        <ns24:Currency>ETB</ns24:Currency>
                        <ns24:CustomerID>1000000015</ns24:CustomerID>
                    </ns24:mACCOUNTINFOSUPERAPPDetailType>
                </ns24:gACCOUNTINFOSUPERAPPDetailType>
            </ACCOUNTINFOSUPERAPPType>
        </ns30:AccountListByCIFResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseAccountListSOAP(xmlData)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Success)
	assert.NotNil(t, result.Details)
	assert.Equal(t, 3, len(result.Details))

	// Validate all accounts are parsed
	accountNumbers := []string{"1000446094649", "1000446094673", "1000446094878"}
	for i, expectedAccount := range accountNumbers {
		assert.Equal(t, expectedAccount, result.Details[i].AccountNumber)
		assert.Equal(t, "UNITED BANK-ECX", result.Details[i].CustomerName)
		assert.Equal(t, "ETB", result.Details[i].Currency)
		assert.Equal(t, "1000000015", result.Details[i].CustomerID)
	}
}

func TestParseAccountListSOAP_CaseInsensitiveSuccess(t *testing.T) {
	xmlData := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns30:AccountListByCIFResponse xmlns:ns30="http://temenos.com/CBESUPERAPP" xmlns:ns24="http://temenos.com/ACCOUNTINFOSUPERAPP">
            <Status>
                <successIndicator>SUCCESS</successIndicator>
            </Status>
            <ACCOUNTINFOSUPERAPPType>
                <ns24:gACCOUNTINFOSUPERAPPDetailType>
                    <ns24:mACCOUNTINFOSUPERAPPDetailType>
                        <ns24:AccountNumber>1000446094649</ns24:AccountNumber>
                        <ns24:CustomerName>UNITED BANK-ECX</ns24:CustomerName>
                        <ns24:Currency>ETB</ns24:Currency>
                    </ns24:mACCOUNTINFOSUPERAPPDetailType>
                </ns24:gACCOUNTINFOSUPERAPPDetailType>
            </ACCOUNTINFOSUPERAPPType>
        </ns30:AccountListByCIFResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseAccountListSOAP(xmlData)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Success)
	assert.NotNil(t, result.Details)
	assert.Equal(t, 1, len(result.Details))
}

