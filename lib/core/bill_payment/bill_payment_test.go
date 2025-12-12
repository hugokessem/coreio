package billpayment

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBillPayment(t *testing.T) {
	tests := []struct {
		name   string
		param  Params
		expect []string
	}{
		{
			name: "Validate Bill Payment XML Generation",
			param: Params{
				Username:            "SUPERAPP",
				Password:            "123456",
				DebitAccountNumber:  "1000000006924",
				DebitCurrency:       "ETB",
				DebitAmount:         "150.00",
				DebitReference:      "Dr Narrative",
				CrediterReference:   "AAF124578",
				CreditAccountNumber: "1000357597823",
				CreditCurrency:      "ETB",
				PaymentDetail:       "EEU BILL PAYMENT",
				ClientReference:     "",
			},
			expect: []string{
				`<password>123456</password>`,
				`<userName>SUPERAPP</userName>`,
				`<fun:DEBITACCTNO>1000000006924</fun:DEBITACCTNO>`,
				`<fun:DEBITCURRENCY>ETB</fun:DEBITCURRENCY>`,
				`<fun:DEBITAMOUNT>150.00</fun:DEBITAMOUNT>`,
				`<fun:DEBITTHEIRREF>Dr Narrative</fun:DEBITTHEIRREF>`,
				`<fun:CREDITTHEIRREF>AAF124578</fun:CREDITTHEIRREF>`,
				`<fun:CREDITACCTNO>1000357597823</fun:CREDITACCTNO>`,
				`<fun:CREDITCURRENCY>ETB</fun:CREDITCURRENCY>`,
				`<fun:PAYMENTDETAILS>EEU BILL PAYMENT</fun:PAYMENTDETAILS>`,
				`<fun:ClientReference></fun:ClientReference>`,
			},
		},
		{
			name: "Validate Bill Payment with different values",
			param: Params{
				Username:            "TESTUSER",
				Password:            "PASSWORD123",
				DebitAccountNumber:  "2000000000001",
				DebitCurrency:       "USD",
				DebitAmount:         "500.00",
				DebitReference:      "Test Debit Ref",
				CrediterReference:   "TEST123456",
				CreditAccountNumber: "3000000000002",
				CreditCurrency:      "USD",
				PaymentDetail:       "Test Payment",
				ClientReference:     "",
			},
			expect: []string{
				`<password>PASSWORD123</password>`,
				`<userName>TESTUSER</userName>`,
				`<fun:DEBITACCTNO>2000000000001</fun:DEBITACCTNO>`,
				`<fun:DEBITCURRENCY>USD</fun:DEBITCURRENCY>`,
				`<fun:DEBITAMOUNT>500.00</fun:DEBITAMOUNT>`,
				`<fun:DEBITTHEIRREF>Test Debit Ref</fun:DEBITTHEIRREF>`,
				`<fun:CREDITTHEIRREF>TEST123456</fun:CREDITTHEIRREF>`,
				`<fun:CREDITACCTNO>3000000000002</fun:CREDITACCTNO>`,
				`<fun:CREDITCURRENCY>USD</fun:CREDITCURRENCY>`,
				`<fun:PAYMENTDETAILS>Test Payment</fun:PAYMENTDETAILS>`,
				`<fun:ClientReference></fun:ClientReference>`,
			},
		},
		{
			name: "Validate Bill Payment with empty fields",
			param: Params{
				Username:            "USER",
				Password:            "PASS",
				DebitAccountNumber:  "",
				DebitCurrency:       "",
				DebitAmount:         "",
				DebitReference:      "",
				CrediterReference:   "",
				CreditAccountNumber: "",
				CreditCurrency:      "",
				PaymentDetail:       "",
			},
			expect: []string{
				`<password>PASS</password>`,
				`<userName>USER</userName>`,
				`<fun:DEBITACCTNO></fun:DEBITACCTNO>`,
				`<fun:DEBITCURRENCY></fun:DEBITCURRENCY>`,
				`<fun:DEBITAMOUNT></fun:DEBITAMOUNT>`,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			xmlRequest := NewBillPayment(tc.param)

			// Validate XML structure
			assert.Contains(t, xmlRequest, "<soapenv:Envelope")
			assert.Contains(t, xmlRequest, "<soapenv:Body>")
			assert.Contains(t, xmlRequest, "<cbes:FTBillPayment>")
			assert.Contains(t, xmlRequest, "<FUNDSTRANSFERBILLPAYSUPERAPPType")

			// Validate all expected strings are present
			for _, expectedStr := range tc.expect {
				assert.Contains(t, xmlRequest, expectedStr)
			}

			// Validate XML is not empty
			assert.NotEmpty(t, xmlRequest)
		})
	}
}

func TestParseBillPaymentSOAP(t *testing.T) {
	tests := []struct {
		name           string
		xmlData        string
		expectedStatus bool
		expectedError  bool
		expectedMsg    string
		validateDetail func(t *testing.T, detail *BillPaymentDetail)
	}{
		{
			name: "Parse successful response",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <FTBillPaymentResponse>
            <Status>
                <transactionId>FT2134353FMC</transactionId>
                <messageId></messageId>
                <successIndicator>success</successIndicator>
                <application>FUNDS.TRANSFER</application>
            </Status>
            <FUNDSTRANSFERType id="FT2134353FMC">
                <TRANSACTIONTYPE>ACSA</TRANSACTIONTYPE>
                <DEBITACCTNO>1000000006924</DEBITACCTNO>
                <DEBITCURRENCY>ETB</DEBITCURRENCY>
                <DEBITAMOUNT>150.00</DEBITAMOUNT>
                <DEBITTHEIRREF>Dr Narrative</DEBITTHEIRREF>
                <CREDITTHEIRREF>AAF124578</CREDITTHEIRREF>
                <CREDITACCTNO>1000357597823</CREDITACCTNO>
                <CREDITCURRENCY>ETB</CREDITCURRENCY>
                <gPAYMENTDETAILS>
                    <PAYMENTDETAILS>EEU BILL PAYMENT</PAYMENTDETAILS>
                </gPAYMENTDETAILS>
                <CHARGECOMDISPLAY>NO</CHARGECOMDISPLAY>
                <COMMISSIONCODE>Debit Plus Charges</COMMISSIONCODE>
                <COMMISSIONTYPE>WAIVE</COMMISSIONTYPE>
                <PROFITCENTRECUST>1000080127</PROFITCENTRECUST>
                <RETURNTODEPT>NO</RETURNTODEPT>
                <FEDFUNDS>NO</FEDFUNDS>
                <POSITIONTYPE>TR</POSITIONTYPE>
                <AMOUNTDEBITED>ETB150.00</AMOUNTDEBITED>
                <AMOUNTCREDITED>ETB150.00</AMOUNTCREDITED>
                <LOCAMTDEBITED>150.00</LOCAMTDEBITED>
                <LOCAMTCREDITED>150.00</LOCAMTCREDITED>
                <CUSTGROUPLEVEL>1</CUSTGROUPLEVEL>
                <DEBITCUSTOMER>1000080127</DEBITCUSTOMER>
                <CREDITCUSTOMER>1047608946</CREDITCUSTOMER>
                <DRADVICEREQDYN>Y</DRADVICEREQDYN>
                <CRADVICEREQDYN>Y</CRADVICEREQDYN>
                <CHARGEDCUSTOMER>1000080127</CHARGEDCUSTOMER>
                <TOTRECCOMM>0</TOTRECCOMM>
                <TOTRECCOMMLCL>0</TOTRECCOMMLCL>
                <TOTRECCOMMCUR>0</TOTRECCOMMCUR>
                <TOTRECCOMMLCLCUR>0</TOTRECCOMMLCLCUR>
                <CHARGECODE>WAIVE</CHARGECODE>
                <ClientReference>AAF124578</ClientReference>
                <DELIVERYINREF>D20251212827631062301</DELIVERYINREF>
                <DELIVERYOUTREF>D20251212827631062302</DELIVERYOUTREF>
                <CHARGETYPE>WAIVE</CHARGETYPE>
                <CHARGEAMOUNT>0</CHARGEAMOUNT>
                <CHARGECURRENCY>ETB</CHARGECURRENCY>
                <CHARGEDATE>20211209</CHARGEDATE>
                <CHARGETIME>0257</CHARGETIME>
                <CHARGESTATUS>SUCCESS</CHARGESTATUS>
                <CHARGEDESCRIPTION>Charge waived</CHARGEDESCRIPTION>
            </FUNDSTRANSFERType>
        </FTBillPaymentResponse>
    </S:Body>
</S:Envelope>`,
			expectedStatus: true,
			expectedError:  false,
			validateDetail: func(t *testing.T, detail *BillPaymentDetail) {
				assert.NotNil(t, detail)
				assert.Equal(t, "FT2134353FMC", detail.TransactionId)
				assert.Equal(t, "ACSA", detail.TransactionType)
				assert.Equal(t, "1000000006924", detail.DebitAccountNumber)
				assert.Equal(t, "ETB", detail.DebitCurrency)
				assert.Equal(t, "150.00", detail.DebitAmount)
				assert.Equal(t, "Dr Narrative", detail.DebitReference)
				assert.Equal(t, "AAF124578", detail.CrediterReference)
				assert.Equal(t, "1000357597823", detail.CreditAccountNumber)
				assert.Equal(t, "ETB", detail.CreditCurrency)
				assert.Equal(t, "EEU BILL PAYMENT", detail.GlobalPaymentDetail.PaymentDetail)
				assert.Equal(t, "AAF124578", detail.ClientReference)
			},
		},
		{
			name: "Parse response with missing status",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <FTBillPaymentResponse>
            <FUNDSTRANSFERType id="FT2134353FMC">
                <DEBITACCTNO>1000000006924</DEBITACCTNO>
            </FUNDSTRANSFERType>
        </FTBillPaymentResponse>
    </S:Body>
</S:Envelope>`,
			expectedStatus: false,
			expectedError:  false,
			expectedMsg:    "Missing Status",
		},
		{
			name: "Parse response with failure status",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <FTBillPaymentResponse>
            <Status>
                <transactionId>FT2134353FMC</transactionId>
                <successIndicator>failure</successIndicator>
            </Status>
        </FTBillPaymentResponse>
    </S:Body>
</S:Envelope>`,
			expectedStatus: false,
			expectedError:  false,
			expectedMsg:    "API returned failure",
		},
		{
			name: "Parse response with missing bill payment detail",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <FTBillPaymentResponse>
            <Status>
                <transactionId>FT2134353FMC</transactionId>
                <successIndicator>success</successIndicator>
            </Status>
        </FTBillPaymentResponse>
    </S:Body>
</S:Envelope>`,
			expectedStatus: false,
			expectedError:  false,
			expectedMsg:    "Missing Bill Payment Detail",
		},
		{
			name: "Parse invalid response type",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <OtherResponse>
            <Status>
                <successIndicator>success</successIndicator>
            </Status>
        </OtherResponse>
    </S:Body>
</S:Envelope>`,
			expectedStatus: false,
			expectedError:  false,
			expectedMsg:    "Invalid response type",
		},
		{
			name: "Parse invalid XML",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <FTBillPaymentResponse>
            <Status>
                <transactionId>FT2134353FMC</transactionId>
                <successIndicator>success</successIndicator>
            </Status>
            <FUNDSTRANSFERType id="FT2134353FMC">
                <DEBITACCTNO>1000000006924</DEBITACCTNO>
                <unclosedTag>
            </FUNDSTRANSFERType>
        </FTBillPaymentResponse>
    </S:Body>
</S:Envelope>`,
			expectedStatus: false,
			expectedError:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ParseBillPaymentSOAP(tc.xmlData)

			if tc.expectedError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tc.expectedStatus, result.Status)
				if tc.expectedMsg != "" {
					assert.Equal(t, tc.expectedMsg, result.Message)
				}
				if tc.validateDetail != nil {
					tc.validateDetail(t, result.Detail)
				}
			}
		})
	}
}

func TestParseBillPaymentSOAP_WithSampleXML(t *testing.T) {
	// Test with the actual sample XML structure from sample.xml
	xmlData := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns32:FTBillPaymentResponse xmlns:ns32="http://temenos.com/CBESUPERAPP">
            <Status>
                <transactionId>FT2134353FMC</transactionId>
                <messageId></messageId>
                <successIndicator>Success</successIndicator>
                <application>FUNDS.TRANSFER</application>
            </Status>
            <FUNDSTRANSFERType id="FT2134353FMC">
                <ns14:TRANSACTIONTYPE xmlns:ns14="http://temenos.com/FUNDSTRANSFER">ACSA</ns14:TRANSACTIONTYPE>
                <ns14:DEBITACCTNO xmlns:ns14="http://temenos.com/FUNDSTRANSFER">1000000006924</ns14:DEBITACCTNO>
                <ns14:DEBITCURRENCY xmlns:ns14="http://temenos.com/FUNDSTRANSFER">ETB</ns14:DEBITCURRENCY>
                <ns14:DEBITAMOUNT xmlns:ns14="http://temenos.com/FUNDSTRANSFER">150.00</ns14:DEBITAMOUNT>
                <ns14:DEBITTHEIRREF xmlns:ns14="http://temenos.com/FUNDSTRANSFER">Dr Narrative</ns14:DEBITTHEIRREF>
                <ns14:CREDITTHEIRREF xmlns:ns14="http://temenos.com/FUNDSTRANSFER">AAF124578</ns14:CREDITTHEIRREF>
                <ns14:CREDITACCTNO xmlns:ns14="http://temenos.com/FUNDSTRANSFER">1000357597823</ns14:CREDITACCTNO>
                <ns14:CREDITCURRENCY xmlns:ns14="http://temenos.com/FUNDSTRANSFER">ETB</ns14:CREDITCURRENCY>
                <ns14:gPAYMENTDETAILS xmlns:ns14="http://temenos.com/FUNDSTRANSFER">
                    <ns14:PAYMENTDETAILS>EEU BILL PAYMENT</ns14:PAYMENTDETAILS>
                </ns14:gPAYMENTDETAILS>
            </FUNDSTRANSFERType>
        </ns32:FTBillPaymentResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseBillPaymentSOAP(xmlData)
	
	// Now that we use case-insensitive comparison, "Success" should be recognized as success
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Status, "Should recognize 'Success' as success")
	if result.Detail != nil {
		assert.Equal(t, "FT2134353FMC", result.Detail.TransactionId)
		assert.Equal(t, "ACSA", result.Detail.TransactionType)
		assert.Equal(t, "1000000006924", result.Detail.DebitAccountNumber)
		assert.Equal(t, "ETB", result.Detail.DebitCurrency)
		assert.Equal(t, "150.00", result.Detail.DebitAmount)
		assert.Equal(t, "Dr Narrative", result.Detail.DebitReference)
		assert.Equal(t, "AAF124578", result.Detail.CrediterReference)
		assert.Equal(t, "1000357597823", result.Detail.CreditAccountNumber)
		assert.Equal(t, "ETB", result.Detail.CreditCurrency)
		assert.Equal(t, "EEU BILL PAYMENT", result.Detail.GlobalPaymentDetail.PaymentDetail)
	}
}

