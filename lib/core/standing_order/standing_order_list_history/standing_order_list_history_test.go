package standingorderlisthistory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStandingOrderListHistory(t *testing.T) {
	tests := []struct {
		name   string
		param  Params
		expect []string
	}{
		{
			name: "Validate standing order list history XML generation",
			param: Params{
				Username:      "SUPERAPP",
				Password:      "123456",
				AccountNumber: "1000373456776",
			},
			expect: []string{
				`<password>123456</password>`,
				`<userName>SUPERAPP</userName>`,
				`<columnName>ACCOUNT</columnName>`,
				`<criteriaValue>1000373456776</criteriaValue>`,
				`<operand>CT</operand>`,
				`<soapenv:Envelope`,
				`<soapenv:Body>`,
				`<cbes:StandingOrderHistorylistbyAc>`,
			},
		},
		{
			name: "Validate standing order list history with different values",
			param: Params{
				Username:      "TESTUSER",
				Password:      "PASSWORD123",
				AccountNumber: "1000000006924",
			},
			expect: []string{
				`<password>PASSWORD123</password>`,
				`<userName>TESTUSER</userName>`,
				`<criteriaValue>1000000006924</criteriaValue>`,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			xmlRequest := NewListStandingOrderHistory(tc.param)

			// Validate XML structure
			assert.Contains(t, xmlRequest, "<soapenv:Envelope")
			assert.Contains(t, xmlRequest, "<soapenv:Body>")
			assert.Contains(t, xmlRequest, "<cbes:StandingOrderHistorylistbyAc>")

			// Validate all expected strings are present
			for _, expectedStr := range tc.expect {
				assert.Contains(t, xmlRequest, expectedStr)
			}

			// Validate XML is not empty
			assert.NotEmpty(t, xmlRequest)
		})
	}
}

func TestParseStandingOrderListHistorySOAP(t *testing.T) {
	tests := []struct {
		name             string
		xmlData          string
		expectedSuccess  bool
		expectedError    bool
		expectedDetails  bool
		expectedCount    int
		expectedMessages []string
	}{
		{
			name: "Parse successful response with multiple standing orders",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns30:StandingOrderHistorylistbyAcResponse xmlns:ns30="http://temenos.com/CBESUPERAPP" xmlns:ns2="http://temenos.com/ACCTSTOLISTHISSUPERAPP">
            <Status>
                <successIndicator>Success</successIndicator>
            </Status>
            <ACCTSTOLISTHISSUPERAPPType>
                <ns2:gACCTSTOLISTHISSUPERAPPDetailType>
                    <ns2:mACCTSTOLISTHISSUPERAPPDetailType>
                        <ns2:StandingOrderID>1000373456776.2</ns2:StandingOrderID>
                        <ns2:KTYPE>Fixed</ns2:KTYPE>
                        <ns2:PAYMENTDETAILS>BILISUMA BAYISA</ns2:PAYMENTDETAILS>
                        <ns2:CURRENCY>ETB</ns2:CURRENCY>
                        <ns2:CURRENTAMOUNTBAL>550.00</ns2:CURRENTAMOUNTBAL>
                        <ns2:CURRENTFREQUENCY>Daily</ns2:CURRENTFREQUENCY>
                        <ns2:CPTYACCTNO>1000374758822</ns2:CPTYACCTNO>
                        <ns2:TOACCTNAME>BETHEL ACADEMY</ns2:TOACCTNAME>
                        <ns2:STOSTARTDATE>20210301</ns2:STOSTARTDATE>
                        <ns2:CURRENTENDDATE>20210310</ns2:CURRENTENDDATE>
                        <ns2:CURRFREQDATE>20210310</ns2:CURRFREQDATE>
                        <ns2:DEBITACCTDESC>BILISUMA BAYISA MINOR URGE AYANA JE</ns2:DEBITACCTDESC>
                        <ns2:DebitAccount>1000373456776</ns2:DebitAccount>
                    </ns2:mACCTSTOLISTHISSUPERAPPDetailType>
                    <ns2:mACCTSTOLISTHISSUPERAPPDetailType>
                        <ns2:StandingOrderID>1000373456776.3</ns2:StandingOrderID>
                        <ns2:KTYPE>Fixed</ns2:KTYPE>
                        <ns2:PAYMENTDETAILS>BILISUMA BAYISA KG 1D</ns2:PAYMENTDETAILS>
                        <ns2:CURRENCY>ETB</ns2:CURRENCY>
                        <ns2:CURRENTAMOUNTBAL>550.00</ns2:CURRENTAMOUNTBAL>
                        <ns2:CURRENTFREQUENCY>Daily</ns2:CURRENTFREQUENCY>
                        <ns2:CPTYACCTNO>1000374758822</ns2:CPTYACCTNO>
                        <ns2:TOACCTNAME>BETHEL ACADEMY</ns2:TOACCTNAME>
                        <ns2:STOSTARTDATE>20210301</ns2:STOSTARTDATE>
                        <ns2:CURRENTENDDATE>20210410</ns2:CURRENTENDDATE>
                        <ns2:CURRFREQDATE>20210410</ns2:CURRFREQDATE>
                        <ns2:DEBITACCTDESC>BILISUMA BAYISA MINOR URGE AYANA JE</ns2:DEBITACCTDESC>
                        <ns2:DebitAccount>1000373456776</ns2:DebitAccount>
                    </ns2:mACCTSTOLISTHISSUPERAPPDetailType>
                </ns2:gACCTSTOLISTHISSUPERAPPDetailType>
            </ACCTSTOLISTHISSUPERAPPType>
        </ns30:StandingOrderHistorylistbyAcResponse>
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
        <ns30:StandingOrderHistorylistbyAcResponse xmlns:ns30="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>Failure</successIndicator>
            </Status>
        </ns30:StandingOrderHistorylistbyAcResponse>
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
        <ns30:StandingOrderHistorylistbyAcResponse xmlns:ns30="http://temenos.com/CBESUPERAPP">
        </ns30:StandingOrderHistorylistbyAcResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess:  false,
			expectedError:    false,
			expectedDetails:  false,
			expectedCount:    0,
			expectedMessages: []string{"Missing Status"},
		},
		{
			name: "Parse response with no standing order history",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns30:StandingOrderHistorylistbyAcResponse xmlns:ns30="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>Success</successIndicator>
            </Status>
            <ACCTSTOLISTHISSUPERAPPType>
            </ACCTSTOLISTHISSUPERAPPType>
        </ns30:StandingOrderHistorylistbyAcResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess:  true,
			expectedError:    false,
			expectedDetails:  false,
			expectedCount:    0,
			expectedMessages: []string{"No Standing Order History Found!"},
		},
		{
			name: "Parse response with empty group",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns30:StandingOrderHistorylistbyAcResponse xmlns:ns30="http://temenos.com/CBESUPERAPP" xmlns:ns2="http://temenos.com/ACCTSTOLISTHISSUPERAPP">
            <Status>
                <successIndicator>Success</successIndicator>
            </Status>
            <ACCTSTOLISTHISSUPERAPPType>
                <ns2:gACCTSTOLISTHISSUPERAPPDetailType>
                </ns2:gACCTSTOLISTHISSUPERAPPDetailType>
            </ACCTSTOLISTHISSUPERAPPType>
        </ns30:StandingOrderHistorylistbyAcResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess:  true,
			expectedError:    false,
			expectedDetails:  false,
			expectedCount:    0,
			expectedMessages: []string{"No Standing Order History Found!"},
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
			expectedMessages: []string{"Invalid response type"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ParseStandingOrderListHistorySOAP(tc.xmlData)

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

func TestParseStandingOrderListHistorySOAP_DetailFields(t *testing.T) {
	xmlData := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns30:StandingOrderHistorylistbyAcResponse xmlns:ns30="http://temenos.com/CBESUPERAPP" xmlns:ns2="http://temenos.com/ACCTSTOLISTHISSUPERAPP">
            <Status>
                <successIndicator>Success</successIndicator>
            </Status>
            <ACCTSTOLISTHISSUPERAPPType>
                <ns2:gACCTSTOLISTHISSUPERAPPDetailType>
                    <ns2:mACCTSTOLISTHISSUPERAPPDetailType>
                        <ns2:StandingOrderID>1000373456776.2</ns2:StandingOrderID>
                        <ns2:KTYPE>Fixed</ns2:KTYPE>
                        <ns2:PAYMENTDETAILS>BILISUMA BAYISA</ns2:PAYMENTDETAILS>
                        <ns2:CURRENCY>ETB</ns2:CURRENCY>
                        <ns2:CURRENTAMOUNTBAL>550.00</ns2:CURRENTAMOUNTBAL>
                        <ns2:CURRENTFREQUENCY>Daily</ns2:CURRENTFREQUENCY>
                        <ns2:DebitAccount>1000373456776</ns2:DebitAccount>
                        <ns2:DEBITACCTDESC>BILISUMA BAYISA MINOR URGE AYANA JE</ns2:DEBITACCTDESC>
                        <ns2:CPTYACCTNO>1000374758822</ns2:CPTYACCTNO>
                        <ns2:TOACCTNAME>BETHEL ACADEMY</ns2:TOACCTNAME>
                        <ns2:CURRFREQDATE>20210310</ns2:CURRFREQDATE>
                        <ns2:CURRENTENDDATE>20210310</ns2:CURRENTENDDATE>
                        <ns2:STOSTARTDATE>20210301</ns2:STOSTARTDATE>
                    </ns2:mACCTSTOLISTHISSUPERAPPDetailType>
                </ns2:gACCTSTOLISTHISSUPERAPPDetailType>
            </ACCTSTOLISTHISSUPERAPPType>
        </ns30:StandingOrderHistorylistbyAcResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseStandingOrderListHistorySOAP(xmlData)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Success)
	assert.NotNil(t, result.Details)
	assert.Equal(t, 1, len(result.Details))

	if len(result.Details) > 0 {
		detail := result.Details[0]
		assert.Equal(t, "1000373456776.2", detail.StandingOrderId)
		assert.Equal(t, "Fixed", detail.OrderType)
		assert.Equal(t, "BILISUMA BAYISA", detail.PaymentDetail)
		assert.Equal(t, "ETB", detail.Currency)
		assert.Equal(t, "550.00", detail.Amount)
		assert.Equal(t, "Daily", detail.Frequency)
		assert.Equal(t, "1000373456776", detail.DebitAccountNumber)
		assert.Equal(t, "BILISUMA BAYISA MINOR URGE AYANA JE", detail.DebitAccountHolderName)
		assert.Equal(t, "1000374758822", detail.CreditAccountNumber)
		assert.Equal(t, "BETHEL ACADEMY", detail.CreditAccountHolderName)
		assert.Equal(t, "20210310", detail.CurrentFrequencyDate)
		assert.Equal(t, "20210310", detail.EndDate)
		assert.Equal(t, "20210301", detail.StartDate)
	}
}

func TestParseStandingOrderListHistorySOAP_MultipleOrders(t *testing.T) {
	xmlData := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns30:StandingOrderHistorylistbyAcResponse xmlns:ns30="http://temenos.com/CBESUPERAPP" xmlns:ns2="http://temenos.com/ACCTSTOLISTHISSUPERAPP">
            <Status>
                <successIndicator>Success</successIndicator>
            </Status>
            <ACCTSTOLISTHISSUPERAPPType>
                <ns2:gACCTSTOLISTHISSUPERAPPDetailType>
                    <ns2:mACCTSTOLISTHISSUPERAPPDetailType>
                        <ns2:StandingOrderID>1000373456776.1</ns2:StandingOrderID>
                        <ns2:KTYPE>Fixed</ns2:KTYPE>
                        <ns2:PAYMENTDETAILS>EDUCATION FEE</ns2:PAYMENTDETAILS>
                        <ns2:CURRENCY>ETB</ns2:CURRENCY>
                        <ns2:CURRENTAMOUNTBAL>550.00</ns2:CURRENTAMOUNTBAL>
                        <ns2:CURRENTFREQUENCY>Monthly</ns2:CURRENTFREQUENCY>
                        <ns2:DebitAccount>1000373456776</ns2:DebitAccount>
                        <ns2:CPTYACCTNO>1000374758822</ns2:CPTYACCTNO>
                    </ns2:mACCTSTOLISTHISSUPERAPPDetailType>
                    <ns2:mACCTSTOLISTHISSUPERAPPDetailType>
                        <ns2:StandingOrderID>1000373456776.2</ns2:StandingOrderID>
                        <ns2:KTYPE>Fixed</ns2:KTYPE>
                        <ns2:PAYMENTDETAILS>BILISUMA BAYISA</ns2:PAYMENTDETAILS>
                        <ns2:CURRENCY>ETB</ns2:CURRENCY>
                        <ns2:CURRENTAMOUNTBAL>550.00</ns2:CURRENTAMOUNTBAL>
                        <ns2:CURRENTFREQUENCY>Daily</ns2:CURRENTFREQUENCY>
                        <ns2:DebitAccount>1000373456776</ns2:DebitAccount>
                        <ns2:CPTYACCTNO>1000374758822</ns2:CPTYACCTNO>
                    </ns2:mACCTSTOLISTHISSUPERAPPDetailType>
                    <ns2:mACCTSTOLISTHISSUPERAPPDetailType>
                        <ns2:StandingOrderID>1000373456776.3</ns2:StandingOrderID>
                        <ns2:KTYPE>Fixed</ns2:KTYPE>
                        <ns2:PAYMENTDETAILS>BILISUMA BAYISA KG 1D</ns2:PAYMENTDETAILS>
                        <ns2:CURRENCY>ETB</ns2:CURRENCY>
                        <ns2:CURRENTAMOUNTBAL>550.00</ns2:CURRENTAMOUNTBAL>
                        <ns2:CURRENTFREQUENCY>Daily</ns2:CURRENTFREQUENCY>
                        <ns2:DebitAccount>1000373456776</ns2:DebitAccount>
                        <ns2:CPTYACCTNO>1000374758822</ns2:CPTYACCTNO>
                    </ns2:mACCTSTOLISTHISSUPERAPPDetailType>
                </ns2:gACCTSTOLISTHISSUPERAPPDetailType>
            </ACCTSTOLISTHISSUPERAPPType>
        </ns30:StandingOrderHistorylistbyAcResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseStandingOrderListHistorySOAP(xmlData)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Success)
	assert.NotNil(t, result.Details)
	assert.Equal(t, 3, len(result.Details))

	// Validate all standing orders are parsed
	standingOrderIds := []string{"1000373456776.1", "1000373456776.2", "1000373456776.3"}
	paymentDetails := []string{"EDUCATION FEE", "BILISUMA BAYISA", "BILISUMA BAYISA KG 1D"}
	frequencies := []string{"Monthly", "Daily", "Daily"}

	for i, expectedId := range standingOrderIds {
		assert.Equal(t, expectedId, result.Details[i].StandingOrderId)
		assert.Equal(t, paymentDetails[i], result.Details[i].PaymentDetail)
		assert.Equal(t, frequencies[i], result.Details[i].Frequency)
		assert.Equal(t, "ETB", result.Details[i].Currency)
		assert.Equal(t, "550.00", result.Details[i].Amount)
		assert.Equal(t, "1000373456776", result.Details[i].DebitAccountNumber)
		assert.Equal(t, "1000374758822", result.Details[i].CreditAccountNumber)
	}
}

func TestParseStandingOrderListHistorySOAP_CaseInsensitiveSuccess(t *testing.T) {
	xmlData := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns30:StandingOrderHistorylistbyAcResponse xmlns:ns30="http://temenos.com/CBESUPERAPP" xmlns:ns2="http://temenos.com/ACCTSTOLISTHISSUPERAPP">
            <Status>
                <successIndicator>SUCCESS</successIndicator>
            </Status>
            <ACCTSTOLISTHISSUPERAPPType>
                <ns2:gACCTSTOLISTHISSUPERAPPDetailType>
                    <ns2:mACCTSTOLISTHISSUPERAPPDetailType>
                        <ns2:StandingOrderID>1000373456776.1</ns2:StandingOrderID>
                        <ns2:KTYPE>Fixed</ns2:KTYPE>
                        <ns2:PAYMENTDETAILS>EDUCATION FEE</ns2:PAYMENTDETAILS>
                        <ns2:CURRENCY>ETB</ns2:CURRENCY>
                        <ns2:CURRENTAMOUNTBAL>550.00</ns2:CURRENTAMOUNTBAL>
                        <ns2:CURRENTFREQUENCY>Monthly</ns2:CURRENTFREQUENCY>
                    </ns2:mACCTSTOLISTHISSUPERAPPDetailType>
                </ns2:gACCTSTOLISTHISSUPERAPPDetailType>
            </ACCTSTOLISTHISSUPERAPPType>
        </ns30:StandingOrderHistorylistbyAcResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseStandingOrderListHistorySOAP(xmlData)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Success)
	assert.NotNil(t, result.Details)
	assert.Equal(t, 1, len(result.Details))
}
