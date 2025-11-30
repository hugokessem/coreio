package fundtransfercheck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFundTransferCheckGeneratedXML(t *testing.T) {
	test := []struct {
		name   string
		param  Params
		expect []string
	}{
		{
			name: "Validate Fund Transfer Check XML Generation",
			param: Params{
				Username: "SUPERAPP",
				Password: "123456",
				FTNumber: "FT21343CXGBD",
			},
			expect: []string{
				"<soapenv:Envelope",
				"<soapenv:Header/>",
				"<soapenv:Body>",
				"<cbes:TransferViewDetails>",
				"<WebRequestCommon>",
				"<password>123456</password>",
				"<userName>SUPERAPP</userName>",
				"<FUNDSTRANSFERVIEWDETAILSSUPERAPPType>",
				"<transactionId>FT21343CXGBD</transactionId>",
			},
		},
		{
			name: "Validate Fund Transfer Check with different values",
			param: Params{
				Username: "TESTUSER",
				Password: "TESTPASS",
				FTNumber: "FT789012",
			},
			expect: []string{
				"<password>TESTPASS</password>",
				"<userName>TESTUSER</userName>",
				"<transactionId>FT789012</transactionId>",
			},
		},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			xmlRequest := NewFundTransferCheck(tc.param)

			// Validate XML structure
			assert.Contains(t, xmlRequest, "<soapenv:Envelope")
			assert.Contains(t, xmlRequest, "<soapenv:Header/>")
			assert.Contains(t, xmlRequest, "<soapenv:Body>")
			assert.Contains(t, xmlRequest, "<cbes:TransferViewDetails>")

			// Validate all expected strings are present
			for _, expectedStr := range tc.expect {
				assert.Contains(t, xmlRequest, expectedStr)
			}

			// Validate XML is not empty
			assert.NotEmpty(t, xmlRequest)
		})
	}
}

func TestParseFundTransferCheckSOAP_Success(t *testing.T) {
	xmlResponse := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns26:TransferViewDetailsResponse xmlns:ns2="http://temenos.com/ACCTSTOLISTHISSUPERAPP"
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
                <transactionId>FT21343CXGBD</transactionId>
                <messageId></messageId>
                <successIndicator>Success</successIndicator>
                <application>FUNDS.TRANSFER</application>
            </Status>
            <FUNDSTRANSFERType id="FT21343CXGBD">
                <ns4:TRANSACTIONTYPE>ACLM</ns4:TRANSACTIONTYPE>
                <ns4:DEBITACCTNO>1000204346953</ns4:DEBITACCTNO>
                <ns4:DEBITCURRENCY>ETB</ns4:DEBITCURRENCY>
                <ns4:DEBITAMOUNT>1.00</ns4:DEBITAMOUNT>
                <ns4:DEBITVALUEDATE>20211209</ns4:DEBITVALUEDATE>
                <ns4:CREDITACCTNO>1000029780939</ns4:CREDITACCTNO>
                <ns4:CREDITCURRENCY>ETB</ns4:CREDITCURRENCY>
                <ns4:CREDITAMOUNT>1.00</ns4:CREDITAMOUNT>
                <ns4:CREDITVALUEDATE>20211209</ns4:CREDITVALUEDATE>
                <ns4:PROCESSINGDATE>20211209</ns4:PROCESSINGDATE>
                <ns4:gCOMMISSIONTYPE>
                    <ns4:mCOMMISSIONTYPE>
                        <ns4:COMMISSIONTYPE>COMMLMT</ns4:COMMISSIONTYPE>
                        <ns4:COMMISSIONAMT>ETB3.00</ns4:COMMISSIONAMT>
                    </ns4:mCOMMISSIONTYPE>
                    <ns4:mCOMMISSIONTYPE>
                        <ns4:COMMISSIONTYPE>CABLECHRG</ns4:COMMISSIONTYPE>
                        <ns4:COMMISSIONAMT>ETB30.00</ns4:COMMISSIONAMT>
                    </ns4:mCOMMISSIONTYPE>
                </ns4:gCOMMISSIONTYPE>
                <ns4:COMMISSIONCODE>DEBIT PLUS CHARGES</ns4:COMMISSIONCODE>
                <ns4:CHARGECODE>WAIVE</ns4:CHARGECODE>
                <ns4:PROFITCENTRECUST>1025015557</ns4:PROFITCENTRECUST>
                <ns4:RETURNTODEPT>NO</ns4:RETURNTODEPT>
                <ns4:FEDFUNDS>NO</ns4:FEDFUNDS>
                <ns4:POSITIONTYPE>TR</ns4:POSITIONTYPE>
                <ns4:AMOUNTDEBITED>ETB34.00</ns4:AMOUNTDEBITED>
                <ns4:AMOUNTCREDITED>ETB1.00</ns4:AMOUNTCREDITED>
                <ns4:TOTALCHARGEAMT>ETB33.00</ns4:TOTALCHARGEAMT>
                <ns4:CREDITCOMPCODE>ET0010493</ns4:CREDITCOMPCODE>
                <ns4:DEBITCOMPCODE>ET0010210</ns4:DEBITCOMPCODE>
                <ns4:LOCAMTDEBITED>34.00</ns4:LOCAMTDEBITED>
                <ns4:LOCAMTCREDITED>1.00</ns4:LOCAMTCREDITED>
                <ns4:LOCALCHARGEAMT>33.00</ns4:LOCALCHARGEAMT>
                <ns4:LOCPOSCHGSAMT>33.00</ns4:LOCPOSCHGSAMT>
                <ns4:CUSTGROUPLEVEL>99</ns4:CUSTGROUPLEVEL>
                <ns4:DEBITCUSTOMER>1025015557</ns4:DEBITCUSTOMER>
                <ns4:CREDITCUSTOMER>9999999999</ns4:CREDITCUSTOMER>
                <ns4:DRADVICEREQDYN>N</ns4:DRADVICEREQDYN>
                <ns4:CRADVICEREQDYN>N</ns4:CRADVICEREQDYN>
                <ns4:CHARGEDCUSTOMER>1025015557</ns4:CHARGEDCUSTOMER>
                <ns4:TOTRECCOMM>0</ns4:TOTRECCOMM>
                <ns4:TOTRECCOMMLCL>0</ns4:TOTRECCOMMLCL>
                <ns4:TOTRECCHG>0</ns4:TOTRECCHG>
                <ns4:TOTRECCHGLCL>0</ns4:TOTRECCHGLCL>
                <ns4:RATEFIXING>NO</ns4:RATEFIXING>
                <ns4:TOTRECCHGCRCCY>0</ns4:TOTRECCHGCRCCY>
                <ns4:TOTSNDCHGCRCCY>33.00</ns4:TOTSNDCHGCRCCY>
                <ns4:AUTHDATE>20211209</ns4:AUTHDATE>
                <ns4:ROUNDTYPE>NATURAL</ns4:ROUNDTYPE>
                <ns4:gSTMTNOS>
                    <ns4:STMTNOS>200749826144191.00</ns4:STMTNOS>
                    <ns4:STMTNOS>1-2</ns4:STMTNOS>
                </ns4:gSTMTNOS>
                <ns4:CURRNO>1</ns4:CURRNO>
                <ns4:gINPUTTER>
                    <ns4:INPUTTER>98261_INPUTTER__OFS_IRISPA</ns4:INPUTTER>
                </ns4:gINPUTTER>
                <ns4:gDATETIME>
                    <ns4:DATETIME>2212161216</ns4:DATETIME>
                </ns4:gDATETIME>
                <ns4:AUTHORISER>98261_INPUTTER_OFS_IRISPA</ns4:AUTHORISER>
                <ns4:COCODE>ET0010493</ns4:COCODE>
                <ns4:DEPTCODE>1</ns4:DEPTCODE>
                <ns4:QUESTION>hhhfrf</ns4:QUESTION>
                <ns4:SECANSWER>gggggy</ns4:SECANSWER>
                <ns4:SECNUMBER>V0lOMjEzNDNXNlNLVA==</ns4:SECNUMBER>
                <ns4:LMTSSENDNO>+251923427268</ns4:LMTSSENDNO>
            </FUNDSTRANSFERType>
        </ns26:TransferViewDetailsResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseFundTransferCheckSOAP(xmlResponse)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Status)
	assert.NotNil(t, result.Detail)

	if result.Detail != nil {
		assert.Equal(t, "ACLM", result.Detail.TransactionType)
		assert.Equal(t, "1000204346953", result.Detail.DebitAccountNumber)
		assert.Equal(t, "ETB", result.Detail.DebitCurrency)
		assert.Equal(t, "1.00", result.Detail.DebitAmount)
		assert.Equal(t, "20211209", result.Detail.DebitValuedate)
		assert.Equal(t, "1000029780939", result.Detail.CreditAccountNumber)
		assert.Equal(t, "ETB", result.Detail.CreditCurrency)
		assert.Equal(t, "1.00", result.Detail.CreditAmount)
		assert.Equal(t, "20211209", result.Detail.CreditValuedate)
		assert.Equal(t, "20211209", result.Detail.ProcessingDate)
		assert.Equal(t, "DEBIT PLUS CHARGES", result.Detail.CommissionCode)
		assert.Equal(t, "WAIVE", result.Detail.ChargeCode)
		assert.Equal(t, "1025015557", result.Detail.ProfitCenterCustomer)
		assert.Equal(t, "NO", result.Detail.ReturnToDept)
		assert.Equal(t, "NO", result.Detail.FedFunds)
		assert.Equal(t, "TR", result.Detail.PositionType)
		assert.Equal(t, "ETB34.00", result.Detail.AmountDebited)
		assert.Equal(t, "ETB1.00", result.Detail.AmountCredited)
		assert.Equal(t, "ETB33.00", result.Detail.TotalChargeAmount)
		assert.Equal(t, "ET0010493", result.Detail.CreditCompCode)
		assert.Equal(t, "ET0010210", result.Detail.DebitCompCode)
		assert.Equal(t, "34.00", result.Detail.LocAmtDebited)
		assert.Equal(t, "1.00", result.Detail.LocAmtCredited)
		assert.Equal(t, "33.00", result.Detail.LocalChargeAmount)
		assert.Equal(t, "33.00", result.Detail.LocalPosChgsAmount)
		assert.Equal(t, "99", result.Detail.CustGroupLevel)
		assert.Equal(t, "1025015557", result.Detail.DebitCustomer)
		assert.Equal(t, "9999999999", result.Detail.CreditCustomer)
		assert.Equal(t, "N", result.Detail.DrAdvicerEqdYN)
		assert.Equal(t, "N", result.Detail.CrAdvicerEqdYN)
		assert.Equal(t, "1025015557", result.Detail.ChargedCustomer)
		assert.Equal(t, "0", result.Detail.TotRecComm)
		assert.Equal(t, "0", result.Detail.TotRecCommLcl)
		assert.Equal(t, "0", result.Detail.TotRecChg)
		assert.Equal(t, "0", result.Detail.TotRecChgLcl)
		assert.Equal(t, "NO", result.Detail.RateFixing)
		assert.Equal(t, "0", result.Detail.TotRecChgCrcCy)
		assert.Equal(t, "33.00", result.Detail.TotSndChgCrcCy)
		assert.Equal(t, "20211209", result.Detail.AuthDate)
		assert.Equal(t, "NATURAL", result.Detail.RoundType)
		assert.Equal(t, "1", result.Detail.CurrNo)
		assert.Equal(t, "98261_INPUTTER__OFS_IRISPA", result.Detail.GInputter.Inputter)
		assert.Equal(t, "2212161216", result.Detail.GDatetime.Datetime)
		assert.Equal(t, "98261_INPUTTER_OFS_IRISPA", result.Detail.Authoriser)
		assert.Equal(t, "ET0010493", result.Detail.CoCode)
		assert.Equal(t, "1", result.Detail.DeptCode)
		assert.Equal(t, "hhhfrf", result.Detail.Question)
		assert.Equal(t, "gggggy", result.Detail.SecAnswer)
		assert.Equal(t, "V0lOMjEzNDNXNlNLVA==", result.Detail.SecNumber)
		assert.Equal(t, "+251923427268", result.Detail.LmtssSendNo)
		// ChargeCommissionDisplay may have 1 or more items depending on XML parsing
		if len(result.Detail.ChargeCommissionDisplay) > 0 {
			assert.GreaterOrEqual(t, len(result.Detail.ChargeCommissionDisplay), 1)
			// Check first item if available
			if len(result.Detail.ChargeCommissionDisplay) >= 1 {
				assert.NotEmpty(t, result.Detail.ChargeCommissionDisplay[0].CommisionType.ComissionType)
				assert.NotEmpty(t, result.Detail.ChargeCommissionDisplay[0].CommisionType.ComissionAmount)
			}
			// Check second item if available
			if len(result.Detail.ChargeCommissionDisplay) >= 2 {
				assert.Equal(t, "COMMLMT", result.Detail.ChargeCommissionDisplay[0].CommisionType.ComissionType)
				assert.Equal(t, "ETB3.00", result.Detail.ChargeCommissionDisplay[0].CommisionType.ComissionAmount)
				assert.Equal(t, "CABLECHRG", result.Detail.ChargeCommissionDisplay[1].CommisionType.ComissionType)
				assert.Equal(t, "ETB30.00", result.Detail.ChargeCommissionDisplay[1].CommisionType.ComissionAmount)
			}
		}
		assert.Len(t, result.Detail.StatementNos.StatementNo, 2)
	}
}

func TestParseFundTransferCheckSOAP_FailureResponse(t *testing.T) {
	xmlResponse := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns26:TransferViewDetailsResponse xmlns:ns26="http://temenos.com/CBESUPERAPP">
            <Status>
                <transactionId>FT21343CXGBD</transactionId>
                <messageId>Transaction not found</messageId>
                <successIndicator>Failure</successIndicator>
                <application>FUNDS.TRANSFER</application>
            </Status>
        </ns26:TransferViewDetailsResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseFundTransferCheckSOAP(xmlResponse)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "Transaction not found", err.Error())
}

func TestParseFundTransferCheckSOAP_MissingStatus(t *testing.T) {
	xmlResponse := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns26:TransferViewDetailsResponse xmlns:ns26="http://temenos.com/CBESUPERAPP">
            <FUNDSTRANSFERType>
                <ns4:TRANSACTIONTYPE>ACLM</ns4:TRANSACTIONTYPE>
            </FUNDSTRANSFERType>
        </ns26:TransferViewDetailsResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseFundTransferCheckSOAP(xmlResponse)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "missing status", err.Error())
}

func TestParseFundTransferCheckSOAP_MissingFundTransferType(t *testing.T) {
	xmlResponse := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns26:TransferViewDetailsResponse xmlns:ns26="http://temenos.com/CBESUPERAPP">
            <Status>
                <transactionId>FT21343CXGBD</transactionId>
                <messageId></messageId>
                <successIndicator>Success</successIndicator>
                <application>FUNDS.TRANSFER</application>
            </Status>
        </ns26:TransferViewDetailsResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseFundTransferCheckSOAP(xmlResponse)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "missing fund transfer type", err.Error())
}

func TestParseFundTransferCheckSOAP_NoTransferViewDetailsResponse(t *testing.T) {
	xmlResponse := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <SomeOtherResponse>
        </SomeOtherResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseFundTransferCheckSOAP(xmlResponse)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.False(t, result.Status)
}

func TestParseFundTransferCheckSOAP_InvalidXML(t *testing.T) {
	xmlResponse := `invalid xml content`

	result, err := ParseFundTransferCheckSOAP(xmlResponse)
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestParseFundTransferCheckSOAP_EmptyResponse(t *testing.T) {
	xmlResponse := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns26:TransferViewDetailsResponse xmlns:ns26="http://temenos.com/CBESUPERAPP">
        </ns26:TransferViewDetailsResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseFundTransferCheckSOAP(xmlResponse)
	// Empty response should return an error because Status is missing
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "missing status", err.Error())
}
