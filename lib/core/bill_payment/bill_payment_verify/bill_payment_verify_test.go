package billpaymentverify

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewBillPaymentVerify(t *testing.T) {
	tests := []struct {
		name   string
		param  Param
		expect []string
	}{
		{
			name: "generates validate request XML",
			param: Param{
				Username:            "SUPERAPP",
				Password:            "123456",
				DebitAccountNumber:  "1000000006924",
				DebitCurrency:       "ETB",
				DebitAmount:         "30.00",
				DebitReference:      "7098700000000075",
				CreditReference:     "159358542456",
				CreditAccountNumber: "1000083962554",
				CreditCurrency:      "ETB",
				CreditAmount:        "",
				PaymentDetails:      "TELEBIRR",
				ClientReference:     "client-ref-1",
				ServiceDescription:  "test Service Description",
			},
			expect: []string{
				"<soapenv:Envelope",
				"<cbes:FTBillPayment_Validate>",
				"<password>123456</password>",
				"<userName>SUPERAPP</userName>",
				"<fun:DEBITACCTNO>1000000006924</fun:DEBITACCTNO>",
				"<fun:DEBITCURRENCY>ETB</fun:DEBITCURRENCY>",
				"<fun:DEBITAMOUNT>30.00</fun:DEBITAMOUNT>",
				"<fun:DEBITTHEIRREF>7098700000000075</fun:DEBITTHEIRREF>",
				"<fun:CREDITTHEIRREF>159358542456</fun:CREDITTHEIRREF>",
				"<fun:CREDITACCTNO>1000083962554</fun:CREDITACCTNO>",
				"<fun:CREDITCURRENCY>ETB</fun:CREDITCURRENCY>",
				"<fun:PAYMENTDETAILS>TELEBIRR</fun:PAYMENTDETAILS>",
				"<fun:ClientReference>client-ref-1</fun:ClientReference>",
				"<fun:ServiceDescription>test Service Description</fun:ServiceDescription>",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			xmlRequest := NewBillPaymentVerify(tc.param)
			assert.NotEmpty(t, xmlRequest)
			for _, expected := range tc.expect {
				assert.Contains(t, xmlRequest, expected)
			}
		})
	}
}

func TestParseBillPaymentVerifySOAP_Success(t *testing.T) {
	xmlData := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns58:FTBillPayment_ValidateResponse xmlns:ns20="http://temenos.com/FUNDSTRANSFER" xmlns:ns58="http://temenos.com/CBESUPERAPP">
            <Status>
                <transactionId>FT21343MVX5S</transactionId>
                <messageId></messageId>
                <successIndicator>Success</successIndicator>
                <application>FUNDS.TRANSFER</application>
            </Status>
            <FUNDSTRANSFERType id="FT21343MVX5S">
                <ns20:TRANSACTIONTYPE>ACSA</ns20:TRANSACTIONTYPE>
                <ns20:DEBITACCTNO>1000000006924</ns20:DEBITACCTNO>
                <ns20:CURRENCYMKTDR>1</ns20:CURRENCYMKTDR>
                <ns20:DEBITCURRENCY>ETB</ns20:DEBITCURRENCY>
                <ns20:DEBITAMOUNT>30.00</ns20:DEBITAMOUNT>
                <ns20:DEBITVALUEDATE>20211209</ns20:DEBITVALUEDATE>
                <ns20:DEBITTHEIRREF>7098700000000075</ns20:DEBITTHEIRREF>
                <ns20:CREDITTHEIRREF>159358542456</ns20:CREDITTHEIRREF>
                <ns20:CREDITACCTNO>1000083962554</ns20:CREDITACCTNO>
                <ns20:CURRENCYMKTCR>1</ns20:CURRENCYMKTCR>
                <ns20:CREDITCURRENCY>ETB</ns20:CREDITCURRENCY>
                <ns20:CREDITVALUEDATE>20211209</ns20:CREDITVALUEDATE>
                <ns20:PROCESSINGDATE>20211209</ns20:PROCESSINGDATE>
                <ns20:gPAYMENTDETAILS>
                    <ns20:PAYMENTDETAILS>TELEBIRR</ns20:PAYMENTDETAILS>
                </ns20:gPAYMENTDETAILS>
                <ns20:CHARGECOMDISPLAY>NO</ns20:CHARGECOMDISPLAY>
                <ns20:COMMISSIONCODE>Debit Plus Charges</ns20:COMMISSIONCODE>
                <ns20:gCOMMISSIONTYPE>
                    <ns20:mCOMMISSIONTYPE>
                        <ns20:COMMISSIONTYPE>CBEWALLETSP</ns20:COMMISSIONTYPE>
                        <ns20:COMMISSIONAMT>ETB9.20</ns20:COMMISSIONAMT>
                    </ns20:mCOMMISSIONTYPE>
                    <ns20:mCOMMISSIONTYPE>
                        <ns20:COMMISSIONTYPE>DRFWALLETSP</ns20:COMMISSIONTYPE>
                        <ns20:COMMISSIONAMT>ETB0.10</ns20:COMMISSIONAMT>
                    </ns20:mCOMMISSIONTYPE>
                    <ns20:mCOMMISSIONTYPE>
                        <ns20:COMMISSIONTYPE>ELWALLETSP</ns20:COMMISSIONTYPE>
                        <ns20:COMMISSIONAMT>ETB0.80</ns20:COMMISSIONAMT>
                    </ns20:mCOMMISSIONTYPE>
                </ns20:gCOMMISSIONTYPE>
                <ns20:CHARGECODE>WAIVE</ns20:CHARGECODE>
                <ns20:PROFITCENTRECUST>1000080127</ns20:PROFITCENTRECUST>
                <ns20:RETURNTODEPT>NO</ns20:RETURNTODEPT>
                <ns20:FEDFUNDS>NO</ns20:FEDFUNDS>
                <ns20:POSITIONTYPE>TR</ns20:POSITIONTYPE>
                <ns20:gTAXTYPE>
                    <ns20:mTAXTYPE>
                        <ns20:TAXTYPE>16</ns20:TAXTYPE>
                        <ns20:TAXAMT>ETB1.50</ns20:TAXAMT>
                    </ns20:mTAXTYPE>
                </ns20:gTAXTYPE>
                <ns20:AMOUNTDEBITED>ETB41.60</ns20:AMOUNTDEBITED>
                <ns20:AMOUNTCREDITED>ETB30.00</ns20:AMOUNTCREDITED>
                <ns20:TOTALCHARGEAMT>ETB11.60</ns20:TOTALCHARGEAMT>
                <ns20:TOTALTAXAMOUNT>ETB1.50</ns20:TOTALTAXAMOUNT>
                <ns20:CREDITCOMPCODE>ET0010535</ns20:CREDITCOMPCODE>
                <ns20:DEBITCOMPCODE>ET0010222</ns20:DEBITCOMPCODE>
                <ns20:LOCAMTDEBITED>41.60</ns20:LOCAMTDEBITED>
                <ns20:LOCAMTCREDITED>30.00</ns20:LOCAMTCREDITED>
                <ns20:LOCTOTTAXAMT>1.50</ns20:LOCTOTTAXAMT>
                <ns20:LOCALCHARGEAMT>11.60</ns20:LOCALCHARGEAMT>
                <ns20:LOCPOSCHGSAMT>11.60</ns20:LOCPOSCHGSAMT>
                <ns20:CUSTGROUPLEVEL>1</ns20:CUSTGROUPLEVEL>
                <ns20:DEBITCUSTOMER>1000080127</ns20:DEBITCUSTOMER>
                <ns20:CREDITCUSTOMER>1010782079</ns20:CREDITCUSTOMER>
                <ns20:DRADVICEREQDYN>Y</ns20:DRADVICEREQDYN>
                <ns20:CRADVICEREQDYN>Y</ns20:CRADVICEREQDYN>
                <ns20:CHARGEDCUSTOMER>1000080127</ns20:CHARGEDCUSTOMER>
                <ns20:TOTRECCOMM>0</ns20:TOTRECCOMM>
                <ns20:TOTRECCOMMLCL>0</ns20:TOTRECCOMMLCL>
                <ns20:TOTRECCHG>0</ns20:TOTRECCHG>
                <ns20:TOTRECCHGLCL>0</ns20:TOTRECCHGLCL>
                <ns20:RATEFIXING>NO</ns20:RATEFIXING>
                <ns20:TOTRECCHGCRCCY>0</ns20:TOTRECCHGCRCCY>
                <ns20:TOTSNDCHGCRCCY>11.60</ns20:TOTSNDCHGCRCCY>
                <ns20:ROUNDTYPE>NATURAL</ns20:ROUNDTYPE>
                <ns20:gOVERRIDE>
                    <ns20:OVERRIDE>STAFF.TXN.AUTH}{</ns20:OVERRIDE>
                </ns20:gOVERRIDE>
                <ns20:COCODE>ET0010001</ns20:COCODE>
                <ns20:SEPACATEG>6502</ns20:SEPACATEG>
                <ns20:PHONENUM>+251913798907</ns20:PHONENUM>
                <ns20:MTOREF>846518494316420061</ns20:MTOREF>
                <ns20:SENDERNAME>ABIY HAILEYESUS MENGISTU</ns20:SENDERNAME>
                <ns20:RECEIVERNAME>BEZA BERHANU ZEGEYE</ns20:RECEIVERNAME>
                <ns20:CEKCS>1713313381.9</ns20:CEKCS>
                <ns20:GPONU>65698.52</ns20:GPONU>
                <ns20:PHONE>+251913323918</ns20:PHONE>
                <ns20:SASEGMENT>,,</ns20:SASEGMENT>
                <ns20:SERVICEDESC>test Service Description</ns20:SERVICEDESC>
            </FUNDSTRANSFERType>
        </ns58:FTBillPayment_ValidateResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseBillPaymentVerifySOAP(xmlData)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.True(t, result.Success)
	require.NotNil(t, result.Detail)

	detail := result.Detail
	assert.Equal(t, "FT21343MVX5S", detail.FTNumber)
	assert.Equal(t, "ACSA", detail.TransactionType)
	assert.Equal(t, "1000000006924", detail.DebitAccountNumber)
	assert.Equal(t, "ETB", detail.DebitCurrency)
	assert.Equal(t, "30.00", detail.DebitAmount)
	assert.Equal(t, "7098700000000075", detail.DebitReference)
	assert.Equal(t, "159358542456", detail.CreditReference)
	assert.Equal(t, "1000083962554", detail.CreditAccountNumber)
	assert.Equal(t, "ETB", detail.CreditCurrency)
	assert.Equal(t, "TELEBIRR", detail.PaymentDetails.PaymentDetail)
	assert.Equal(t, "ETB41.60", detail.DebitAmountWithCurrency)
	assert.Equal(t, "ETB30.00", detail.CreditAmountWithCurrency)
	assert.Equal(t, "ETB11.60", detail.TotalChargeAmount)
	assert.Equal(t, "1.5000", detail.TotalTaxAmount)
	assert.Equal(t, "ETB1.5000", detail.LocalTotalTaxAmount)
	assert.Equal(t, "ETB0.1000", detail.DisasterReservedFund)
	assert.Equal(t, "ETB10.0000", detail.TotalCommisionWithComission)
	assert.Equal(t, "ETB30.0000", detail.OriginalPaidAmount)
	assert.Equal(t, "ABIY HAILEYESUS MENGISTU", detail.DebitAccountHolderName)
	assert.Equal(t, "BEZA BERHANU ZEGEYE", detail.ReceiverName)
	assert.Equal(t, "test Service Description", detail.ServiceDescription)
	assert.Equal(t, "6502", detail.AccountCategory)
	assert.Equal(t, "846518494316420061", detail.TransactionID)
	assert.Len(t, detail.GlobalCommissionType.MultipleCommissionType, 3)
}

func TestParseBillPaymentVerifySOAP_MissingStatus(t *testing.T) {
	xmlData := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <FTBillPayment_ValidateResponse>
            <FUNDSTRANSFERType id="FT21343MVX5S">
                <DEBITACCTNO>1000000006924</DEBITACCTNO>
            </FUNDSTRANSFERType>
        </FTBillPayment_ValidateResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseBillPaymentVerifySOAP(xmlData)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.False(t, result.Success)
	assert.Equal(t, []string{"Missing Status!"}, result.Messages)
}

func TestParseBillPaymentVerifySOAP_FailureStatus(t *testing.T) {
	xmlData := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <FTBillPayment_ValidateResponse>
            <Status>
                <transactionId>FT21343MVX5S</transactionId>
                <successIndicator>Failure</successIndicator>
                <messages>Insufficient funds</messages>
            </Status>
        </FTBillPayment_ValidateResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseBillPaymentVerifySOAP(xmlData)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.False(t, result.Success)
	assert.Equal(t, []string{"Insufficient funds"}, result.Messages)
	assert.Nil(t, result.Detail)
}

func TestParseBillPaymentVerifySOAP_InvalidResponseType(t *testing.T) {
	xmlData := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <OtherResponse>
            <Status>
                <successIndicator>Success</successIndicator>
            </Status>
        </OtherResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseBillPaymentVerifySOAP(xmlData)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.False(t, result.Success)
	assert.Equal(t, []string{"Invalid response type"}, result.Messages)
}

func TestParseBillPaymentVerifySOAP_InvalidXML(t *testing.T) {
	xmlData := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <FTBillPayment_ValidateResponse>
            <Status>
                <successIndicator>Success</successIndicator>
            </Status>
            <FUNDSTRANSFERType id="FT21343MVX5S">
                <unclosedTag>
            </FUNDSTRANSFERType>
        </FTBillPayment_ValidateResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseBillPaymentVerifySOAP(xmlData)
	assert.Error(t, err)
	assert.Nil(t, result)
}
