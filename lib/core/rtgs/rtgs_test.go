package rtgs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRTGS(t *testing.T) {
	xmlRequest := NewRTGS(Params{
		Username:           "SUPERAPP",
		Password:           "123456",
		Company:            "ET0010222",
		DebitAccountNumber: "1000251861149",
		DebitCurrency:      "ETB",
		CreditAmount:       "50",
		OrderingCust: []string{
			"MOR-BAHR0049695769",
			"1437145650002",
			"DURBET TRADE SHARE COMPANY",
		},
		BenAccountNumber: "0100191040217",
		PaymentDetails:   "0049695769 BAHIR DAR BRANCH",
		BkToBkInfo: []string{
			"/REC/TAXFRMID",
			"//0049695769/1437145650002",
			"//EFFDATE/20260402",
		},
		BillID: "1437145650002",
	})

	assert.Contains(t, xmlRequest, `<cbes:MinistryOfRevenueMORRTGSFT>`)
	assert.Contains(t, xmlRequest, `<FUNDSTRANSFERMORSUPERAPPType id="">`)
	assert.Contains(t, xmlRequest, `<company>ET0010222</company>`)
	assert.Contains(t, xmlRequest, `<password>123456</password>`)
	assert.Contains(t, xmlRequest, `<userName>SUPERAPP</userName>`)
	assert.Contains(t, xmlRequest, `<fun:DEBITACCTNO>1000251861149</fun:DEBITACCTNO>`)
	assert.Contains(t, xmlRequest, `<fun:DEBITCURRENCY>ETB</fun:DEBITCURRENCY>`)
	assert.Contains(t, xmlRequest, `<fun:CREDITAMOUNT>50</fun:CREDITAMOUNT>`)
	assert.Contains(t, xmlRequest, `<fun:ORDERINGCUST>MOR-BAHR0049695769</fun:ORDERINGCUST>`)
	assert.Contains(t, xmlRequest, `<fun:ORDERINGCUST>1437145650002</fun:ORDERINGCUST>`)
	assert.Contains(t, xmlRequest, `<fun:ORDERINGCUST>DURBET TRADE SHARE COMPANY</fun:ORDERINGCUST>`)
	assert.Contains(t, xmlRequest, `<fun:BENACCTNO>0100191040217</fun:BENACCTNO>`)
	assert.Contains(t, xmlRequest, `<fun:PAYMENTDETAILS>0049695769 BAHIR DAR BRANCH</fun:PAYMENTDETAILS>`)
	assert.Contains(t, xmlRequest, `<fun:BKTOBKINFO>/REC/TAXFRMID</fun:BKTOBKINFO>`)
	assert.Contains(t, xmlRequest, `<fun:BKTOBKINFO>//0049695769/1437145650002</fun:BKTOBKINFO>`)
	assert.Contains(t, xmlRequest, `<fun:BKTOBKINFO>//EFFDATE/20260402</fun:BKTOBKINFO>`)
	assert.Contains(t, xmlRequest, `<fun:BILLID>1437145650002</fun:BILLID>`)
}

func TestParseRTGSSOAP_Success(t *testing.T) {
	xmlResponse := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns32:MinistryOfRevenueMORRTGSFTResponse xmlns:ns14="http://temenos.com/FUNDSTRANSFER" xmlns:ns32="http://temenos.com/CBESUPERAPP">
            <Status>
                <transactionId>FT21343Y0P1L</transactionId>
                <successIndicator>Success</successIndicator>
                <application>FUNDS.TRANSFER</application>
            </Status>
            <FUNDSTRANSFERType id="FT21343Y0P1L">
                <ns14:TRANSACTIONTYPE>ACRT</ns14:TRANSACTIONTYPE>
                <ns14:DEBITACCTNO>1000251861149</ns14:DEBITACCTNO>
                <ns14:DEBITCURRENCY>ETB</ns14:DEBITCURRENCY>
                <ns14:CREDITAMOUNT>50.00</ns14:CREDITAMOUNT>
                <ns14:BENACCTNO>0100191040217</ns14:BENACCTNO>
                <ns14:BILLID>1437145650002</ns14:BILLID>
                <ns14:AMOUNTDEBITED>ETB50.00</ns14:AMOUNTDEBITED>
                <ns14:AMOUNTCREDITED>ETB50.00</ns14:AMOUNTCREDITED>
                <ns14:PROCESSINGDATE>20260402</ns14:PROCESSINGDATE>
                <ns14:COCODE>ET0010222</ns14:COCODE>
            </FUNDSTRANSFERType>
        </ns32:MinistryOfRevenueMORRTGSFTResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseRTGSSOAP(xmlResponse)
	assert.NoError(t, err)
	assert.True(t, result.Success)
	assert.NotNil(t, result.Detail)
	assert.Equal(t, "FT21343Y0P1L", result.Detail.FTNumber)
	assert.Equal(t, "1000251861149", result.Detail.DebitAccountNumber)
	assert.Equal(t, "0100191040217", result.Detail.BenAccountNumber)
	assert.Equal(t, "1437145650002", result.Detail.BillID)
}

func TestParseRTGSSOAP_Failure(t *testing.T) {
	xmlResponse := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns32:MinistryOfRevenueMORRTGSFTResponse xmlns:ns32="http://temenos.com/CBESUPERAPP">
            <Status>
                <transactionId></transactionId>
                <successIndicator>T24Error</successIndicator>
                <messages>Insufficient funds</messages>
            </Status>
        </ns32:MinistryOfRevenueMORRTGSFTResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseRTGSSOAP(xmlResponse)
	assert.NoError(t, err)
	assert.False(t, result.Success)
	assert.Contains(t, result.Messages, "Insufficient funds")
}
