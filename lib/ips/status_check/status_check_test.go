package statuscheck

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsePaymentStatusSOAP_Success(t *testing.T) {
	xmlData := `<?xml version="1.0" encoding="utf-8"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
    <soapenv:Body>
        <NS1:PaymentStatusResponse xmlns:NS1="http://MB_IPS">
            <output1>
                <NS2:AppHdr xmlns:NS2="urn:iso:std:iso:20022:tech:xsd:head.001.001.03">
                    <NS2:Fr>
                        <NS2:FIId>
                            <NS2:FinInstnId>
                                <NS2:Othr>
                                    <NS2:Id>ETSETAA</NS2:Id>
                                </NS2:Othr>
                            </NS2:FinInstnId>
                        </NS2:FIId>
                    </NS2:Fr>
                    <NS2:To>
                        <NS2:FIId>
                            <NS2:FinInstnId>
                                <NS2:Othr>
                                    <NS2:Id>FP</NS2:Id>
                                </NS2:Othr>
                            </NS2:FinInstnId>
                        </NS2:FIId>
                    </NS2:To>
                    <NS2:BizMsgIdr>ETSETAA44fd9da03a1e4009a</NS2:BizMsgIdr>
                    <NS2:MsgDefIdr>pacs.002.001.12</NS2:MsgDefIdr>
                    <NS2:CreDt>2025-11-10T15:50:37.037Z</NS2:CreDt>
                    <NS2:Rltd>
                        <NS2:Fr>
                            <NS2:FIId>
                                <NS2:FinInstnId>
                                    <NS2:Othr>
                                        <NS2:Id>FP</NS2:Id>
                                    </NS2:Othr>
                                </NS2:FinInstnId>
                            </NS2:FIId>
                        </NS2:Fr>
                        <NS2:To>
                            <NS2:FIId>
                                <NS2:FinInstnId>
                                    <NS2:Othr>
                                        <NS2:Id>ETSETAA</NS2:Id>
                                    </NS2:Othr>
                                </NS2:FinInstnId>
                            </NS2:FIId>
                        </NS2:To>
                        <NS2:BizMsgIdr>CBETETAA531415266346350005549</NS2:BizMsgIdr>
                        <NS2:MsgDefIdr>pacs.008.001.10</NS2:MsgDefIdr>
                        <NS2:CreDt>2025-11-10T15:50:37.037Z</NS2:CreDt>
                    </NS2:Rltd>
                </NS2:AppHdr>
                <NS3:Document xmlns:NS3="urn:iso:std:iso:20022:tech:xsd:pacs.002.001.12">
                    <NS3:FIToFIPmtStsRpt>
                        <NS3:GrpHdr>
                            <NS3:MsgId>ETSETAA44fd9da03a1e4009a</NS3:MsgId>
                            <NS3:CreDtTm>2025-11-10T13:50:39.043374235+01:00</NS3:CreDtTm>
                            <NS3:InstgAgt>
                                <NS3:FinInstnId>
                                    <NS3:Othr>
                                        <NS3:Id>ETSETAA</NS3:Id>
                                    </NS3:Othr>
                                </NS3:FinInstnId>
                            </NS3:InstgAgt>
                            <NS3:InstdAgt>
                                <NS3:FinInstnId>
                                    <NS3:Othr>
                                        <NS3:Id>CBETETAA</NS3:Id>
                                    </NS3:Othr>
                                </NS3:FinInstnId>
                            </NS3:InstdAgt>
                        </NS3:GrpHdr>
                        <NS3:TxInfAndSts>
                            <NS3:OrgnlGrpInf>
                                <NS3:OrgnlMsgId>CBETETAA897809371</NS3:OrgnlMsgId>
                                <NS3:OrgnlMsgNmId>pacs.008.001.10</NS3:OrgnlMsgNmId>
                                <NS3:OrgnlCreDtTm>2023-07-25T00:00:00.000+03:00</NS3:OrgnlCreDtTm>
                            </NS3:OrgnlGrpInf>
                            <NS3:OrgnlEndToEndId>CBETETAA897809371</NS3:OrgnlEndToEndId>
                            <NS3:OrgnlTxId>CBETETAAFT18978092</NS3:OrgnlTxId>
                            <NS3:TxSts>ACSC</NS3:TxSts>
                            <NS3:AccptncDtTm>2025-11-10T12:50:39.043417711Z</NS3:AccptncDtTm>
                            <NS3:OrgnlTxRef>
                                <NS3:IntrBkSttlmAmt Ccy="ETB">10</NS3:IntrBkSttlmAmt>
                                <NS3:Amt>
                                    <NS3:InstdAmt Ccy="ETB">10</NS3:InstdAmt>
                                </NS3:Amt>
                                <NS3:RmtInf>
                                    <NS3:Ustrd>Transferring my funds</NS3:Ustrd>
                                </NS3:RmtInf>
                                <NS3:Dbtr>
                                    <NS3:Pty>
                                        <NS3:Nm>MSCWT</NS3:Nm>
                                        <NS3:PstlAdr>
                                            <NS3:AdrLine>MOSCOW</NS3:AdrLine>
                                        </NS3:PstlAdr>
                                    </NS3:Pty>
                                </NS3:Dbtr>
                                <NS3:DbtrAcct>
                                    <NS3:Id>
                                        <NS3:Othr>
                                            <NS3:Id>1234567890</NS3:Id>
                                            <NS3:SchmeNm>
                                                <NS3:Prtry>ACCT</NS3:Prtry>
                                            </NS3:SchmeNm>
                                        </NS3:Othr>
                                    </NS3:Id>
                                </NS3:DbtrAcct>
                                <NS3:Cdtr>
                                    <NS3:Pty>
                                        <NS3:Nm>test</NS3:Nm>
                                    </NS3:Pty>
                                </NS3:Cdtr>
                                <NS3:CdtrAcct>
                                    <NS3:Id>
                                        <NS3:Othr>
                                            <NS3:Id>1234567890</NS3:Id>
                                            <NS3:SchmeNm>
                                                <NS3:Prtry>ACCT</NS3:Prtry>
                                            </NS3:SchmeNm>
                                        </NS3:Othr>
                                    </NS3:Id>
                                </NS3:CdtrAcct>
                            </NS3:OrgnlTxRef>
                        </NS3:TxInfAndSts>
                    </NS3:FIToFIPmtStsRpt>
                </NS3:Document>
            </output1>
        </NS1:PaymentStatusResponse>
    </soapenv:Body>
</soapenv:Envelope>`

	result, err := ParsePaymentStatusSOAP(xmlData)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Success)
	assert.NotNil(t, result.Detail)

	detail := result.Detail
	assert.Equal(t, "ETSETAA", detail.FromBankBIC)
	assert.Equal(t, "FP", detail.ToBankBIC)
	assert.Equal(t, "ETSETAA44fd9da03a1e4009a", detail.BizMessageIdentifier)
	assert.Equal(t, "pacs.002.001.12", detail.MessageDefinitionIdentifier)
	assert.Equal(t, "2025-11-10T15:50:37.037Z", detail.CreationDate)
	assert.Equal(t, "FP", detail.RelatedFromBankBIC)
	assert.Equal(t, "ETSETAA", detail.RelatedToBankBIC)
	assert.Equal(t, "CBETETAA531415266346350005549", detail.RelatedBizMessageIdentifier)
	assert.Equal(t, "pacs.008.001.10", detail.RelatedMessageDefinitionId)
	assert.Equal(t, "ETSETAA44fd9da03a1e4009a", detail.MessageId)
	assert.Equal(t, "2025-11-10T13:50:39.043374235+01:00", detail.CreationDateTime)
	assert.Equal(t, "ETSETAA", detail.InstructingAgent)
	assert.Equal(t, "CBETETAA", detail.InstructedAgent)
	assert.Equal(t, "CBETETAA897809371", detail.OriginalMessageId)
	assert.Equal(t, "pacs.008.001.10", detail.OriginalMessageNameId)
	assert.Equal(t, "2023-07-25T00:00:00.000+03:00", detail.OriginalCreationDateTime)
	assert.Equal(t, "CBETETAA897809371", detail.OriginalEndToEndId)
	assert.Equal(t, "CBETETAAFT18978092", detail.OriginalTransactionId)
	assert.Equal(t, "ACSC", detail.TransactionStatus)
	assert.Equal(t, "2025-11-10T12:50:39.043417711Z", detail.AcceptanceDateTime)
	assert.Equal(t, "10", detail.InterBankSettlementAmount)
	assert.Equal(t, "ETB", detail.InterBankSettlementCurrency)
	assert.Equal(t, "10", detail.InstructedAmount)
	assert.Equal(t, "ETB", detail.InstructedAmountCurrency)
	assert.Equal(t, "Transferring my funds", detail.RemittanceInformation)
	assert.Equal(t, "MSCWT", detail.DebtorName)
	assert.Equal(t, "MOSCOW", detail.DebtorAddress)
	assert.Equal(t, "1234567890", detail.DebtorAccountNumber)
	assert.Equal(t, "test", detail.CreditorName)
	assert.Equal(t, "1234567890", detail.CreditorAccountNumber)
}

func TestParsePaymentStatusSOAP_NonACSCStatus(t *testing.T) {
	xmlData := `<?xml version="1.0" encoding="utf-8"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
    <soapenv:Body>
        <NS1:PaymentStatusResponse xmlns:NS1="http://MB_IPS">
            <output1>
                <NS2:AppHdr xmlns:NS2="urn:iso:std:iso:20022:tech:xsd:head.001.001.03">
                    <NS2:Fr>
                        <NS2:FIId>
                            <NS2:FinInstnId>
                                <NS2:Othr>
                                    <NS2:Id>ETSETAA</NS2:Id>
                                </NS2:Othr>
                            </NS2:FinInstnId>
                        </NS2:FIId>
                    </NS2:Fr>
                    <NS2:To>
                        <NS2:FIId>
                            <NS2:FinInstnId>
                                <NS2:Othr>
                                    <NS2:Id>FP</NS2:Id>
                                </NS2:Othr>
                            </NS2:FinInstnId>
                        </NS2:FIId>
                    </NS2:To>
                    <NS2:BizMsgIdr>ETSETAA44fd9da03a1e4009a</NS2:BizMsgIdr>
                    <NS2:MsgDefIdr>pacs.002.001.12</NS2:MsgDefIdr>
                    <NS2:CreDt>2025-11-10T15:50:37.037Z</NS2:CreDt>
                </NS2:AppHdr>
                <NS3:Document xmlns:NS3="urn:iso:std:iso:20022:tech:xsd:pacs.002.001.12">
                    <NS3:FIToFIPmtStsRpt>
                        <NS3:GrpHdr>
                            <NS3:MsgId>ETSETAA44fd9da03a1e4009a</NS3:MsgId>
                            <NS3:CreDtTm>2025-11-10T13:50:39.043374235+01:00</NS3:CreDtTm>
                            <NS3:InstgAgt>
                                <NS3:FinInstnId>
                                    <NS3:Othr>
                                        <NS3:Id>ETSETAA</NS3:Id>
                                    </NS3:Othr>
                                </NS3:FinInstnId>
                            </NS3:InstgAgt>
                            <NS3:InstdAgt>
                                <NS3:FinInstnId>
                                    <NS3:Othr>
                                        <NS3:Id>CBETETAA</NS3:Id>
                                    </NS3:Othr>
                                </NS3:FinInstnId>
                            </NS3:InstdAgt>
                        </NS3:GrpHdr>
                        <NS3:TxInfAndSts>
                            <NS3:OrgnlGrpInf>
                                <NS3:OrgnlMsgId>CBETETAA897809371</NS3:OrgnlMsgId>
                                <NS3:OrgnlMsgNmId>pacs.008.001.10</NS3:OrgnlMsgNmId>
                                <NS3:OrgnlCreDtTm>2023-07-25T00:00:00.000+03:00</NS3:OrgnlCreDtTm>
                            </NS3:OrgnlGrpInf>
                            <NS3:OrgnlEndToEndId>CBETETAA897809371</NS3:OrgnlEndToEndId>
                            <NS3:OrgnlTxId>CBETETAAFT18978092</NS3:OrgnlTxId>
                            <NS3:TxSts>RJCT</NS3:TxSts>
                            <NS3:AccptncDtTm>2025-11-10T12:50:39.043417711Z</NS3:AccptncDtTm>
                        </NS3:TxInfAndSts>
                    </NS3:FIToFIPmtStsRpt>
                </NS3:Document>
            </output1>
        </NS1:PaymentStatusResponse>
    </soapenv:Body>
</soapenv:Envelope>`

	result, err := ParsePaymentStatusSOAP(xmlData)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.False(t, result.Success)
	assert.Contains(t, result.Messages[0], "Transaction status is not ACSC")
	// Should still have detail even if status is not ACSC
	assert.NotNil(t, result.Detail)
	if result.Detail != nil {
		assert.Equal(t, "RJCT", result.Detail.TransactionStatus)
	}
}

func TestParsePaymentStatusSOAP_InvalidResponse(t *testing.T) {
	xmlData := `<?xml version="1.0" encoding="utf-8"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
    <soapenv:Body>
        <OtherResponse>
        </OtherResponse>
    </soapenv:Body>
</soapenv:Envelope>`

	result, err := ParsePaymentStatusSOAP(xmlData)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.False(t, result.Success)
	assert.Contains(t, result.Messages[0], "Invalid response type")
}

func TestParsePaymentStatusSOAP_ErrorResponse(t *testing.T) {
	// Test errorResponse without Envelope wrapper
	xmlData := `<errorResponse>
		<code>500</code>
		<message>Internal Server Error</message>
		<description>Transaction not found</description>
	</errorResponse>`

	result, err := ParsePaymentStatusSOAP(xmlData)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.False(t, result.Success)
	assert.Contains(t, result.Messages[0], "API Error")
	assert.Contains(t, result.Messages[0], "500")
}

func TestParsePaymentStatusSOAP_ErrorResponseInEnvelope(t *testing.T) {
	// Test errorResponse within Envelope
	xmlData := `<?xml version="1.0" encoding="utf-8"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
    <soapenv:Body>
        <errorResponse>
            <code>404</code>
            <message>Not Found</message>
            <description>Transaction ID not found</description>
        </errorResponse>
    </soapenv:Body>
</soapenv:Envelope>`

	result, err := ParsePaymentStatusSOAP(xmlData)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.False(t, result.Success)
	assert.Contains(t, result.Messages[0], "API Error")
	// The error message should contain the code, message, or description
	assert.True(t, 
		strings.Contains(result.Messages[0], "404") ||
		strings.Contains(result.Messages[0], "Not Found") ||
		strings.Contains(result.Messages[0], "Transaction ID not found"),
		"Error message should contain error details: %s", result.Messages[0])
}

func TestParsePaymentStatusSOAP_InvalidXML(t *testing.T) {
	invalidXML := `<invalid>xml</structure>`

	result, err := ParsePaymentStatusSOAP(invalidXML)
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestNewStatusCheck(t *testing.T) {
	params := Params{
		DebitBankBIC:                  "ETSETAA",
		BizMessageIdentifier:          "ETSETAA44fd9da03a1e4009a",
		MessageIdentifier:             "ETSETAA44fd9da03a1e4009a",
		CreditDateTime:                "2025-11-10T13:50:39.043374235+01:00",
		CreditDate:                    "2025-11-10T15:50:37.037Z",
		OriginalTransactionIdentifier: "CBETETAAFT18978092",
	}

	xmlRequest := NewStatusCheck(params)
	assert.NotEmpty(t, xmlRequest)
	assert.Contains(t, xmlRequest, "<mb:PaymentStatus>")
	assert.Contains(t, xmlRequest, "<urn1:FIToFIPmtStsReq>")
	assert.Contains(t, xmlRequest, "ETSETAA")
	assert.Contains(t, xmlRequest, "CBETETAAFT18978092")
	assert.Contains(t, xmlRequest, "ETSETAA44fd9da03a1e4009a")
}

