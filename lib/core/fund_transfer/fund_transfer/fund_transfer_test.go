package fundtransfer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFundTransfer(t *testing.T) {
	tests := []struct {
		name   string
		param  Params
		expect []string
	}{
		{
			name: "Validate Fund Transfer XML Generation - matches curl request",
			param: Params{
				Username:            "SUPERAPP",
				Password:            "123456",
				DebitAccountNumber:  "1000000006924",
				DebitCurrency:       "ETB",
				CreditAccountNumber: "1000357597823",
				CreditCurrency:      "ETB",
				DebitReference:      "DEBIT NARRATIVE",
				CreditReference:     "CREDIT NARRATIVE",
				DebitAmount:         "50",
				TransactionID:       "12385824578895",
				PaymentDetail:       "TEST PAYMENT",
				ServiceCode:         "GLOBAL",
			},
			expect: []string{
				"<fun:DEBITACCTNO>1000000006924</fun:DEBITACCTNO>",
				"<fun:DEBITCURRENCY>ETB</fun:DEBITCURRENCY>",
				"<fun:DEBITAMOUNT>50</fun:DEBITAMOUNT>",
				"<fun:DEBITTHEIRREF>DEBIT NARRATIVE</fun:DEBITTHEIRREF>",
				"<fun:CREDITTHEIRREF>CREDIT NARRATIVE</fun:CREDITTHEIRREF>",
				"<fun:CREDITACCTNO>1000357597823</fun:CREDITACCTNO>",
				"<fun:CREDITCURRENCY>ETB</fun:CREDITCURRENCY>",
				"<fun:PAYMENTDETAILS>TEST PAYMENT</fun:PAYMENTDETAILS>",
				"<fun:ClientReference>12385824578895</fun:ClientReference>",
				"<fun:ServiceCode>GLOBAL</fun:ServiceCode>",
				"<password>123456</password>",
				"<userName>SUPERAPP</userName>",
				"<fun:gPAYMENTDETAILS g=\"1\">",
				"<fun:gCOMMISSIONTYPE g=\"1\">",
			},
		},
		{
			name: "Validate Fund Transfer with different values",
			param: Params{
				Username:            "SUPERAPP",
				Password:            "123456",
				DebitAccountNumber:  "10009876543",
				DebitCurrency:       "ETB",
				CreditAccountNumber: "10001234567",
				CreditCurrency:      "ETB",
				DebitReference:      "Payment for invoice",
				CreditReference:     "Received payment",
				DebitAmount:         "1500.75",
				TransactionID:       "TXN123456",
				PaymentDetail:       "Urgent transfer",
				ServiceCode:         "GLOBAL",
			},
			expect: []string{
				"<fun:DEBITACCTNO>10009876543</fun:DEBITACCTNO>",
				"<fun:DEBITAMOUNT>1500.75</fun:DEBITAMOUNT>",
				"<fun:ClientReference>TXN123456</fun:ClientReference>",
				"<fun:ServiceCode>GLOBAL</fun:ServiceCode>",
			},
		},
		{
			name: "Validate Fund Transfer with different currency",
			param: Params{
				Username:            "TESTUSER",
				Password:            "PASSWORD123",
				DebitAccountNumber:  "2000000001",
				DebitCurrency:       "USD",
				CreditAccountNumber: "2000000002",
				CreditCurrency:      "USD",
				DebitReference:      "International transfer",
				CreditReference:     "Received USD payment",
				DebitAmount:         "500.00",
				TransactionID:       "TXN789012",
				PaymentDetail:       "International payment",
				ServiceCode:         "GLOBAL",
			},
			expect: []string{
				"<fun:DEBITCURRENCY>USD</fun:DEBITCURRENCY>",
				"<fun:CREDITCURRENCY>USD</fun:CREDITCURRENCY>",
				"<fun:DEBITAMOUNT>500.00</fun:DEBITAMOUNT>",
				"<fun:ClientReference>TXN789012</fun:ClientReference>",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			xmlRequest := NewFundTransfer(tc.param)
			assert.NotEmpty(t, xmlRequest, "Generated XML should not be empty")
			// Validate XML structure
			assert.Contains(t, xmlRequest, "<soapenv:Envelope", "Should contain SOAP envelope")
			assert.Contains(t, xmlRequest, "<cbes:AccountTransfer>", "Should contain AccountTransfer")
			assert.Contains(t, xmlRequest, "<FUNDSTRANSFERFTTXNSUPERAPPType", "Should contain FUNDSTRANSFERFTTXNSUPERAPPType")
		})
	}
}

func TestParseFundTransferSOAP_Success(t *testing.T) {
	xmlResponse := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns32:AccountTransferResponse xmlns:ns14="http://temenos.com/FUNDSTRANSFER" xmlns:ns32="http://temenos.com/CBESUPERAPP">
            <Status>
                <transactionId>FT21343Y0P1L</transactionId>
                <successIndicator>Success</successIndicator>
                <application>FUNDS.TRANSFER</application>
            </Status>
            <FUNDSTRANSFERType id="FT21343Y0P1L">
                <ns14:TRANSACTIONTYPE>ACSA</ns14:TRANSACTIONTYPE>
                <ns14:DEBITACCTNO>1000000006924</ns14:DEBITACCTNO>
                <ns14:CURRENCYMKTDR>1</ns14:CURRENCYMKTDR>
                <ns14:DEBITCURRENCY>ETB</ns14:DEBITCURRENCY>
                <ns14:DEBITAMOUNT>50.00</ns14:DEBITAMOUNT>
                <ns14:DEBITVALUEDATE>20211209</ns14:DEBITVALUEDATE>
                <ns14:DEBITTHEIRREF>DEBIT NARRATIVE</ns14:DEBITTHEIRREF>
                <ns14:CREDITTHEIRREF>CREDIT NARRATIVE</ns14:CREDITTHEIRREF>
                <ns14:CREDITACCTNO>1000357597823</ns14:CREDITACCTNO>
                <ns14:CURRENCYMKTCR>1</ns14:CURRENCYMKTCR>
                <ns14:CREDITCURRENCY>ETB</ns14:CREDITCURRENCY>
                <ns14:CREDITVALUEDATE>20211209</ns14:CREDITVALUEDATE>
                <ns14:PROCESSINGDATE>20211209</ns14:PROCESSINGDATE>
                <ns14:gPAYMENTDETAILS>
                    <ns14:PAYMENTDETAILS>TEST PAYMENT</ns14:PAYMENTDETAILS>
                </ns14:gPAYMENTDETAILS>
                <ns14:CHARGECOMDISPLAY>NO</ns14:CHARGECOMDISPLAY>
                <ns14:COMMISSIONCODE>DEBIT PLUS CHARGES</ns14:COMMISSIONCODE>
                <ns14:gCOMMISSIONTYPE>
                    <ns14:mCOMMISSIONTYPE>
                        <ns14:COMMISSIONTYPE>COMMLMT</ns14:COMMISSIONTYPE>
                        <ns14:COMMISSIONAMT>ETB3.00</ns14:COMMISSIONAMT>
                    </ns14:mCOMMISSIONTYPE>
                </ns14:gCOMMISSIONTYPE>
                <ns14:CHARGECODE>WAIVE</ns14:CHARGECODE>
                <ns14:PROFITCENTRECUST>1000080127</ns14:PROFITCENTRECUST>
                <ns14:RETURNTODEPT>NO</ns14:RETURNTODEPT>
                <ns14:FEDFUNDS>NO</ns14:FEDFUNDS>
                <ns14:POSITIONTYPE>TR</ns14:POSITIONTYPE>
                <ns14:gTAXTYPE>
                    <ns14:mTAXTYPE>
                        <ns14:TAXTYPE>17</ns14:TAXTYPE>
                        <ns14:TAXAMT>ETB0.45</ns14:TAXAMT>
                    </ns14:mTAXTYPE>
                </ns14:gTAXTYPE>
                <ns14:AMOUNTDEBITED>ETB53.45</ns14:AMOUNTDEBITED>
                <ns14:AMOUNTCREDITED>ETB50.00</ns14:AMOUNTCREDITED>
                <ns14:TOTALCHARGEAMT>ETB3.45</ns14:TOTALCHARGEAMT>
                <ns14:TOTALTAXAMOUNT>ETB0.45</ns14:TOTALTAXAMOUNT>
                <ns14:gDELIVERYOUTREF>
                    <ns14:DELIVERYOUTREF>D20251212376084602801-900.1.1       DEBIT ADVICE</ns14:DELIVERYOUTREF>
                    <ns14:DELIVERYOUTREF>D20251212376084602802-910.2.1       CREDIT ADVICE</ns14:DELIVERYOUTREF>
                </ns14:gDELIVERYOUTREF>
                <ns14:CREDITCOMPCODE>ET0010434</ns14:CREDITCOMPCODE>
                <ns14:DEBITCOMPCODE>ET0010222</ns14:DEBITCOMPCODE>
                <ns14:LOCAMTDEBITED>53.45</ns14:LOCAMTDEBITED>
                <ns14:LOCAMTCREDITED>50.00</ns14:LOCAMTCREDITED>
                <ns14:LOCTOTTAXAMT>0.45</ns14:LOCTOTTAXAMT>
                <ns14:LOCALCHARGEAMT>3.45</ns14:LOCALCHARGEAMT>
                <ns14:LOCPOSCHGSAMT>3.45</ns14:LOCPOSCHGSAMT>
                <ns14:CUSTGROUPLEVEL>1</ns14:CUSTGROUPLEVEL>
                <ns14:DEBITCUSTOMER>1000080127</ns14:DEBITCUSTOMER>
                <ns14:CREDITCUSTOMER>1047608946</ns14:CREDITCUSTOMER>
                <ns14:DRADVICEREQDYN>Y</ns14:DRADVICEREQDYN>
                <ns14:CRADVICEREQDYN>Y</ns14:CRADVICEREQDYN>
                <ns14:CHARGEDCUSTOMER>1000080127</ns14:CHARGEDCUSTOMER>
                <ns14:TOTRECCOMM>0</ns14:TOTRECCOMM>
                <ns14:TOTRECCOMMLCL>0</ns14:TOTRECCOMMLCL>
                <ns14:TOTRECCHG>0</ns14:TOTRECCHG>
                <ns14:TOTRECCHGLCL>0</ns14:TOTRECCHGLCL>
                <ns14:RATEFIXING>NO</ns14:RATEFIXING>
                <ns14:TOTRECCHGCRCCY>0</ns14:TOTRECCHGCRCCY>
                <ns14:TOTSNDCHGCRCCY>3.45</ns14:TOTSNDCHGCRCCY>
                <ns14:AUTHDATE>20211209</ns14:AUTHDATE>
                <ns14:ROUNDTYPE>NATURAL</ns14:ROUNDTYPE>
                <ns14:gSTMTNOS>
                    <ns14:STMTNOS>211663760846027.00</ns14:STMTNOS>
                    <ns14:STMTNOS>1-3</ns14:STMTNOS>
                    <ns14:STMTNOS>1</ns14:STMTNOS>
                    <ns14:STMTNOS>ET0010222</ns14:STMTNOS>
                </ns14:gSTMTNOS>
                <ns14:gOVERRIDE>
                    <ns14:OVERRIDE>STAFF.TXN.AUTH}{</ns14:OVERRIDE>
                </ns14:gOVERRIDE>
                <ns14:CURRNO>1</ns14:CURRNO>
                <ns14:gINPUTTER>
                    <ns14:INPUTTER>37608_SUPERAPP.1__OFS_GCS</ns14:INPUTTER>
                </ns14:gINPUTTER>
                <ns14:gDATETIME>
                    <ns14:DATETIME>2512121247</ns14:DATETIME>
                </ns14:gDATETIME>
                <ns14:AUTHORISER>37608_SUPERAPP.1_OFS_GCS</ns14:AUTHORISER>
                <ns14:COCODE>ET0010001</ns14:COCODE>
                <ns14:DEPTCODE>1</ns14:DEPTCODE>
                <ns14:LINPUTVERSION>FUNDS.TRANSFER,FT.TXN.SUPERAPP</ns14:LINPUTVERSION>
                <ns14:LAUTHVERSION>FUNDS.TRANSFER,FT.TXN.SUPERAPP</ns14:LAUTHVERSION>
                <ns14:MTOREF>12385824578895</ns14:MTOREF>
                <ns14:SENDERNAME>ABIY HAILEYESUS MENGISTU</ns14:SENDERNAME>
                <ns14:RECEIVERNAME>SHIMALIS ZERFU A/C FOR SOSINA SHIMA</ns14:RECEIVERNAME>
                <ns14:SERVICECODE>GLOBAL</ns14:SERVICECODE>
                <ns14:CEKCS>5587591375.85</ns14:CEKCS>
                <ns14:GPONU>3233.8</ns14:GPONU>
            </FUNDSTRANSFERType>
        </ns32:AccountTransferResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseFundTransferSOAP(xmlResponse)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Success)
	assert.NotNil(t, result.Detail)

	detail := result.Detail
	// Validate all fields match the curl request response
	assert.Equal(t, "FT21343Y0P1L", detail.FTNumber)
	assert.Equal(t, "ACSA", detail.TransactionType)
	assert.Equal(t, "1000000006924", detail.DebitAccountNumber)
	assert.Equal(t, "1000357597823", detail.CreditAccountNumber)
	assert.Equal(t, "ETB", detail.DebitCurrency)
	assert.Equal(t, "ETB", detail.CreditCurrency)
	assert.Equal(t, "50.00", detail.DebitAmount)
	assert.Equal(t, "DEBIT NARRATIVE", detail.DebitReference)
	assert.Equal(t, "CREDIT NARRATIVE", detail.CreditReference)
	assert.Equal(t, "TEST PAYMENT", detail.PaymentDetails.PaymentDetail)
	assert.Equal(t, "GLOBAL", detail.ServiceCode)
	assert.Equal(t, "12385824578895", detail.TransactionID)
	assert.Equal(t, "ABIY HAILEYESUS MENGISTU", detail.DebitAccountHolderName)
	assert.Equal(t, "SHIMALIS ZERFU A/C FOR SOSINA SHIMA", detail.ReceiverName)
}

func TestParseFundTransferSOAP_Failure(t *testing.T) {
	xmlResponse := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns32:AccountTransferResponse xmlns:ns32="http://temenos.com/CBESUPERAPP">
            <Status>
                <transactionId>FT21343Y0P1L</transactionId>
                <successIndicator>Failure</successIndicator>
                <application>FUNDS.TRANSFER</application>
                <messages>Error message</messages>
            </Status>
        </ns32:AccountTransferResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseFundTransferSOAP(xmlResponse)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.False(t, result.Success)
	assert.NotEmpty(t, result.Messages)
}

func TestParseFundTransferSOAP_MissingStatus(t *testing.T) {
	xmlResponse := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns32:AccountTransferResponse xmlns:ns32="http://temenos.com/CBESUPERAPP">
        </ns32:AccountTransferResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseFundTransferSOAP(xmlResponse)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.False(t, result.Success)
	assert.Contains(t, result.Messages, "Missing Status")
}

func TestParseFundTransferSOAP_NoDetail(t *testing.T) {
	xmlResponse := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns32:AccountTransferResponse xmlns:ns32="http://temenos.com/CBESUPERAPP">
            <Status>
                <transactionId>FT21343Y0P1L</transactionId>
                <successIndicator>Success</successIndicator>
                <application>FUNDS.TRANSFER</application>
            </Status>
        </ns32:AccountTransferResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseFundTransferSOAP(xmlResponse)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Success)
	assert.Nil(t, result.Detail)
}

func TestParseFundTransferSOAP_InvalidXML(t *testing.T) {
	invalidXML := `<invalid>xml</structure>`

	result, err := ParseFundTransferSOAP(invalidXML)
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestParseFundTransferSOAP_InvalidResponseType(t *testing.T) {
	xmlResponse := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <OtherResponse>
        </OtherResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseFundTransferSOAP(xmlResponse)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.False(t, result.Success)
	assert.Contains(t, result.Messages, "Invalid response type")
}
