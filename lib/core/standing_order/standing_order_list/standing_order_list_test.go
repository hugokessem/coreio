package standingorderlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListStandingOrderGeneratedXML(t *testing.T) {
	test := []struct {
		name   string
		param  Params
		expect []string
	}{
		{
			name: "Validate List Standing Order",
			param: Params{
				Username:      "SUPERAPP",
				Password:      "123456",
				AccountNumber: "10009876543",
			},
			expect: []string{
				"<criteriaValue>10009876543</criteriaValue>",
			},
		},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			xmlRequest := NewListStandingOrder(tc.param)
			for _, expectedStr := range tc.expect {
				assert.Contains(t, xmlRequest, expectedStr)
			}
		})
	}
}

func TestParseListStandingOrderSOAP(t *testing.T) {
	xmlResponse := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns18:ListStandingOrdersResponse xmlns:ns2="http://temenos.com/ACCTSTOLISTHISSUPERAPP" xmlns:ns3="http://temenos.com/FUNDSTRANSFERFTREVERSESUPERAPP" xmlns:ns4="http://temenos.com/FUNDSTRANSFER" xmlns:ns5="http://temenos.com/FUNDSTRANSFERFTTXNSUPERAPP" xmlns:ns6="http://temenos.com/ACCTLOCKEDAMOUNTSSUPERAPP" xmlns:ns7="http://temenos.com/FUNDSTRANSFERBILLPAYSUPERAPP" xmlns:ns8="http://temenos.com/STANDINGORDERTXNLISTSUPERAPP" xmlns:ns9="http://temenos.com/ACCTSTOLISTSUPERAPP" xmlns:ns10="http://temenos.com/ACCTSTMTRGSUPERAPP" xmlns:ns11="http://temenos.com/ACLOCKEDEVENTSCREATELOCKSUPERAPP" xmlns:ns12="http://temenos.com/ACLOCKEDEVENTS" xmlns:ns13="http://temenos.com/ACLOCKEDEVENTSRELEASELOCKSUPERAPP" xmlns:ns14="http://temenos.com/ACCOUNTENQUIRYSUPERAPP" xmlns:ns15="http://temenos.com/STANDINGORDERMANAGEORDERSUPERAPP" xmlns:ns16="http://temenos.com/STANDINGORDER" xmlns:ns17="http://temenos.com/CBEMINISTMTENQ" xmlns:ns18="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>Success</successIndicator>
            </Status>
            <ACCTSTOLISTSUPERAPPType>
                <ns9:gACCTSTOLISTSUPERAPPDetailType>
                    <ns9:mACCTSTOLISTSUPERAPPDetailType>
                        <ns9:StandingOrderID>1000000006924.13</ns9:StandingOrderID>
                        <ns9:KTYPE>Fixed</ns9:KTYPE>
                        <ns9:PAYMENTDETAILS>Fund transfer</ns9:PAYMENTDETAILS>
                        <ns9:CURRENCY>ETB</ns9:CURRENCY>
                        <ns9:CURRENTAMOUNTBAL>124.00</ns9:CURRENTAMOUNTBAL>
                        <ns9:CURRENTFREQUENCY>Weekly</ns9:CURRENTFREQUENCY>
                        <ns9:CPTYACCTNO>1000357597823</ns9:CPTYACCTNO>
                        <ns9:TOACCTNAME>SHIMALIS ZERFU A/C FOR SOSINA SHIMA</ns9:TOACCTNAME>
                        <ns9:STOSTARTDATE>20220512</ns9:STOSTARTDATE>
                        <ns9:CURRENTENDDATE>20260101</ns9:CURRENTENDDATE>
                        <ns9:CURRFREQDATE>20220512</ns9:CURRFREQDATE>
                        <ns9:DEBITACCTDESC>ABIY HAILEYESUS MENGISTU</ns9:DEBITACCTDESC>
                        <ns9:DebitAccount>100000000692</ns9:DebitAccount>
                    </ns9:mACCTSTOLISTSUPERAPPDetailType>
                </ns9:gACCTSTOLISTSUPERAPPDetailType>
            </ACCTSTOLISTSUPERAPPType>
        </ns18:ListStandingOrdersResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseListStandingOrderSOAP(xmlResponse)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Success)
	assert.Len(t, result.Details, 1)
	
	if len(result.Details) > 0 {
		detail := result.Details[0]
		assert.Equal(t, "1000000006924.13", detail.StandingOrderId)
		assert.Equal(t, "Fixed", detail.OrderType)
		assert.Equal(t, "Fund transfer", detail.PaymentDetail)
		assert.Equal(t, "ETB", detail.Currency)
		assert.Equal(t, "124.00", detail.Amount)
		assert.Equal(t, "Weekly", detail.Frequency)
		assert.Equal(t, "1000357597823", detail.CreditAccountNumber)
		assert.Equal(t, "SHIMALIS ZERFU A/C FOR SOSINA SHIMA", detail.CreditAccountHolderName)
		assert.Equal(t, "20260101", detail.CurrentDate)
		assert.Equal(t, "ABIY HAILEYESUS MENGISTU", detail.DebitAccountHolderName)
		assert.Equal(t, "100000000692", detail.DebitAccountNumber)
	}
}
