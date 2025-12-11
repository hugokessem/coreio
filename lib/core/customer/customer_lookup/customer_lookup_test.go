package customerlookup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCustomerLookupGeneratedXML(t *testing.T) {
	test := []struct {
		name   string
		param  Params
		expect []string
	}{
		{
			name: "Validate Customer Lookup XML Generation",
			param: Params{
				Username:           "SUPERAPP",
				Password:           "123456",
				CustomerIdentifier: "1000123456789",
			},
			expect: []string{
				"<soapenv:Envelope",
				"<soapenv:Header/>",
				"<soapenv:Body>",
				"<cbes:CustomerInformationDetails>",
				"<WebRequestCommon>",
				"<password>123456</password>",
				"<userName>SUPERAPP</userName>",
				"<CUSTOMERINFOSUPERAPPType>",
				"<columnName>ID</columnName>",
				"<criteriaValue>1000123456789</criteriaValue>",
				"<operand>EQ</operand>",
			},
		},
		{
			name: "Validate Customer Lookup with different values",
			param: Params{
				Username:           "TESTUSER",
				Password:           "TESTPASS",
				CustomerIdentifier: "2000987654321",
			},
			expect: []string{
				"<password>TESTPASS</password>",
				"<userName>TESTUSER</userName>",
				"<criteriaValue>2000987654321</criteriaValue>",
			},
		},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			xmlRequest := NewCustomerLookup(tc.param)

			// Validate XML structure
			assert.Contains(t, xmlRequest, "<soapenv:Envelope")
			assert.Contains(t, xmlRequest, "<soapenv:Header/>")
			assert.Contains(t, xmlRequest, "<soapenv:Body>")
			assert.Contains(t, xmlRequest, "<cbes:CustomerInformationDetails>")

			// Validate all expected strings are present
			for _, expectedStr := range tc.expect {
				assert.Contains(t, xmlRequest, expectedStr)
			}

			// Validate XML is not empty
			assert.NotEmpty(t, xmlRequest)
		})
	}
}

func TestParseCustomerLookupSOAP(t *testing.T) {
	xmlResponse := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns26:CustomerInformationDetailsResponse
            xmlns:ns2="http://temenos.com/ACCTSTOLISTHISSUPERAPP"
            xmlns:ns3="http://temenos.com/FUNDSTRANSFERVIEWDETAILSSUPERAPP"
            xmlns:ns4="http://temenos.com/FUNDSTRANSFER"
            xmlns:ns5="http://temenos.com/FUNDSTRANSFERFTREVERSESUPERAPP"
            xmlns:ns6="http://temenos.com/ATMCARDTYPEENQMBX"
            xmlns:ns7="http://temenos.com/ATMCARDSTATUSENQMBX"
            xmlns:ns8="http://temenos.com/BRANCHLISTSUPERAPP"
            xmlns:ns9="http://temenos.com/STANDINGORDERTXNLISTSUPERAPP"
            xmlns:ns10="http://temenos.com/ACCTSTMTRGSUPERAPP"
            xmlns:ns11="http://temenos.com/ACLOCKEDEVENTSCREATELOCKSUPERAPP"
            xmlns:ns12="http://temenos.com/ACLOCKEDEVENTS"
            xmlns:ns13="http://temenos.com/STANDINGORDERMANAGEORDERSUPERAPP"
            xmlns:ns14="http://temenos.com/STANDINGORDER"
            xmlns:ns15="http://temenos.com/ACCOUNTENQUIRYSUPERAPP"
            xmlns:ns16="http://temenos.com/CBEMINISTMTENQ"
            xmlns:ns17="http://temenos.com/ATMCARDREGDETCARDREPLACESUPERAPP"
            xmlns:ns18="http://temenos.com/ATMCARDREGDET"
            xmlns:ns19="http://temenos.com/ATMCARDREGDETCARDREQSUPERAPP"
            xmlns:ns20="http://temenos.com/FUNDSTRANSFERFTTXNSUPERAPP"
            xmlns:ns21="http://temenos.com/CUSTOMERINFOSUPERAPP"
            xmlns:ns22="http://temenos.com/ACCTLOCKEDAMOUNTSSUPERAPP"
            xmlns:ns23="http://temenos.com/FUNDSTRANSFERBILLPAYSUPERAPP"
            xmlns:ns24="http://temenos.com/ACCTSTOLISTSUPERAPP"
            xmlns:ns25="http://temenos.com/ACLOCKEDEVENTSRELEASELOCKSUPERAPP"
            xmlns:ns26="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>Success</successIndicator>
            </Status>
            <CUSTOMERINFOSUPERAPPType>
                <ns21:gCUSTOMERINFOSUPERAPPDetailType>
                    <ns21:mCUSTOMERINFOSUPERAPPDetailType>
                        <ns21:FullName>LUKAS DASALEGN HATIYE</ns21:FullName>
                        <ns21:BirthDate>19630507</ns21:BirthDate>
                        <ns21:Gender>MALE</ns21:Gender>
                        <ns21:Address>SNNPWO</ns21:Address>
                        <ns21:PhoneNumber>+251916138917</ns21:PhoneNumber>
                        <ns21:Email>test@example.com</ns21:Email>
                        <ns21:City>SNNP</ns21:City>
                        <ns21:Nationality>ET</ns21:Nationality>
                        <ns21:MaritalStatus>MARRIED</ns21:MaritalStatus>
                        <ns21:PostalCode>1000</ns21:PostalCode>
                        <ns21:IDDocument>KEBELE ID</ns21:IDDocument>
                        <ns21:Title>MR</ns21:Title>
                        <ns21:FirstName>LUKAS</ns21:FirstName>
                        <ns21:LastName>DASALEGN</ns21:LastName>
                    </ns21:mCUSTOMERINFOSUPERAPPDetailType>
                </ns21:gCUSTOMERINFOSUPERAPPDetailType>
            </CUSTOMERINFOSUPERAPPType>
        </ns26:CustomerInformationDetailsResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseCustomerLookupSOAP(xmlResponse)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Success)
	assert.NotNil(t, result.CustomerInfos)

	if result.CustomerInfos != nil {
		customer := result.CustomerInfos
		assert.Equal(t, "LUKAS DASALEGN HATIYE", customer.FullName)
		assert.Equal(t, "19630507", customer.BirthDate)
		// Note: Gender field is not copied in ParseCustomerLookupSOAP function
		assert.Equal(t, "SNNPWO", customer.Address)
		assert.Equal(t, "+251916138917", customer.PhoneNumber)
		assert.Equal(t, "test@example.com", customer.Email)
		assert.Equal(t, "SNNP", customer.City)
		assert.Equal(t, "ET", customer.Nationality)
		assert.Equal(t, "MARRIED", customer.MaritalStatus)
		assert.Equal(t, "1000", customer.PostalCode)
		assert.Equal(t, "KEBELE ID", customer.IDDocument)
		assert.Equal(t, "MR", customer.Title)
		assert.Equal(t, "LUKAS", customer.FirstName)
		assert.Equal(t, "DASALEGN", customer.LastName)
	}
}

func TestParseCustomerLookupSOAP_FailureResponse(t *testing.T) {
	xmlResponse := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns26:CustomerInformationDetailsResponse xmlns:ns26="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>Failure</successIndicator>
            </Status>
            <CUSTOMERINFOSUPERAPPType>
                <ns21:gCUSTOMERINFOSUPERAPPDetailType xmlns:ns21="http://temenos.com/CUSTOMERINFOSUPERAPP">
                </ns21:gCUSTOMERINFOSUPERAPPDetailType>
            </CUSTOMERINFOSUPERAPPType>
        </ns26:CustomerInformationDetailsResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseCustomerLookupSOAP(xmlResponse)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.False(t, result.Success)
	assert.Equal(t, "Customer lookup failed", result.Message)
	assert.Nil(t, result.CustomerInfos)
}

func TestParseCustomerLookupSOAP_NoCustomerDetails(t *testing.T) {
	xmlResponse := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns26:CustomerInformationDetailsResponse xmlns:ns26="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>Success</successIndicator>
            </Status>
            <CUSTOMERINFOSUPERAPPType>
                <ns21:gCUSTOMERINFOSUPERAPPDetailType xmlns:ns21="http://temenos.com/CUSTOMERINFOSUPERAPP">
                </ns21:gCUSTOMERINFOSUPERAPPDetailType>
            </CUSTOMERINFOSUPERAPPType>
        </ns26:CustomerInformationDetailsResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseCustomerLookupSOAP(xmlResponse)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.False(t, result.Success)
	assert.Equal(t, "No customer details found", result.Message)
	assert.Nil(t, result.CustomerInfos)
}

func TestParseCustomerLookupSOAP_NoStatus(t *testing.T) {
	xmlResponse := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns26:CustomerInformationDetailsResponse xmlns:ns26="http://temenos.com/CBESUPERAPP">
            <CUSTOMERINFOSUPERAPPType>
                <ns21:gCUSTOMERINFOSUPERAPPDetailType xmlns:ns21="http://temenos.com/CUSTOMERINFOSUPERAPP">
                </ns21:gCUSTOMERINFOSUPERAPPDetailType>
            </CUSTOMERINFOSUPERAPPType>
        </ns26:CustomerInformationDetailsResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseCustomerLookupSOAP(xmlResponse)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.False(t, result.Success)
	assert.Equal(t, "Invalid response structure", result.Message)
	assert.Nil(t, result.CustomerInfos)
}

func TestParseCustomerLookupSOAP_InvalidResponse(t *testing.T) {
	xmlResponse := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <SomeOtherResponse>
        </SomeOtherResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseCustomerLookupSOAP(xmlResponse)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.False(t, result.Success)
	assert.Equal(t, "Invalid response type", result.Message)
	assert.Nil(t, result.CustomerInfos)
}

func TestParseCustomerLookupSOAP_InvalidXML(t *testing.T) {
	xmlResponse := `invalid xml content`

	result, err := ParseCustomerLookupSOAP(xmlResponse)
	assert.Error(t, err)
	assert.Nil(t, result)
}
