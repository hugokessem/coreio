package qrpayment

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewQrPayment(t *testing.T) {
	tests := []struct {
		name   string
		param  Params
		expect []string
	}{
		{
			name: "Validate QR Payment XML generation with all fields",
			param: Params{
				DebitBankBIC:              "CBETETAA",
				CreditBankBIC:             "ETSETAA",
				BizMessageIdentifier:      "CBETETAA553572981",
				CreditDate:                "2023-06-24T00:00:00.000Z",
				MessageIdentifier:         "CBETETAA847772981",
				CreditDateTime:            "2023-06-24T00:00:00.000+03:00",
				EndToEndIdentifier:        "CBETETAA897809371",
				TransactionIdentifier:     "CBETETAAFT18978092",
				InterBankSettlementAmount: "1000.00",
				AccptanceDtatTime:         "2023-06-24T12:00:00.000Z",
				InstructedAmount:          "1000.00",
				DebitAccountNumber:        "1000000006924",
				CreditAccountNumber:       "1000382499388",
				CreditAccountHolderName:   "John Doe",
				Narative:                  "QR Payment Test",
				DebiterInformation: DebiterInformation{
					Name:           "Debiter Name",
					StreetName:     "Main Street",
					BuildingNumber: "123",
					PostalCode:     "1000",
					TownName:       "Addis Ababa",
					Country:        "ET",
				},
				CreditInformation: CrediterInformation{
					Name:           "Crediter Name",
					StreetName:     "Second Street",
					BuildingNumber: "456",
					PostalCode:     "2000",
					City:           "Addis Ababa",
					TownName:       "Addis Ababa",
					Country:        "ET",
					AddressLine:    "Additional Address Info",
				},
			},
			expect: []string{
				`<soapenv:Envelope`,
				`<soapenv:Body>`,
				`<mb:Payment>`,
				`<urn:Id>CBETETAA</urn:Id>`,
				`<urn:Id>ETSETAA</urn:Id>`,
				`<urn:BizMsgIdr>CBETETAA553572981</urn:BizMsgIdr>`,
				`<urn:CreDt>2023-06-24T00:00:00.000Z</urn:CreDt>`,
				`<urn1:MsgId>CBETETAA847772981</urn1:MsgId>`,
				`<urn1:CreDtTm>2023-06-24T00:00:00.000+03:00</urn1:CreDtTm>`,
				`<urn1:EndToEndId>CBETETAA897809371</urn1:EndToEndId>`,
				`<urn1:TxId>CBETETAAFT18978092</urn1:TxId>`,
				`<urn1:IntrBkSttlmAmt Ccy="ETB">1000.00</urn1:IntrBkSttlmAmt>`,
				`<urn1:AccptncDtTm>2023-06-24T12:00:00.000Z</urn1:AccptncDtTm>`,
				`<urn1:InstdAmt Ccy="ETB">1000.00</urn1:InstdAmt>`,
				`<urn1:Nm>Debiter Name</urn1:Nm>`,
				`<urn1:StrtNm>Main Street</urn1:StrtNm>`,
				`<urn1:BldgNb>123</urn1:BldgNb>`,
				`<urn1:PstCd>1000</urn1:PstCd>`,
				`<urn1:TwnNm>Addis Ababa</urn1:TwnNm>`,
				`<urn1:Ctry>ET</urn1:Ctry>`,
				`<urn1:Id>1000000006924</urn1:Id>`,
				`<urn1:Id>1000382499388</urn1:Id>`,
				`<urn1:Nm>Crediter Name</urn1:Nm>`,
				`<urn1:StrtNm>Second Street</urn1:StrtNm>`,
				`<urn1:BldgNb>456</urn1:BldgNb>`,
				`<urn1:PstCd>2000</urn1:PstCd>`,
				`<urn1:TwnNm>Addis Ababa</urn1:TwnNm>`,
				`<urn1:Ctry>ET</urn1:Ctry>`,
				`<urn1:AdrLine>Additional Address Info</urn1:AdrLine>`,
				`<urn1:Ustrd>QR Payment Test</urn1:Ustrd>`,
			},
		},
		{
			name: "Validate QR Payment XML with different values",
			param: Params{
				DebitBankBIC:              "AWINETAA",
				CreditBankBIC:             "CBETETAA",
				BizMessageIdentifier:      "AWINETAA999999999",
				CreditDate:                "2024-01-15T00:00:00.000Z",
				MessageIdentifier:         "AWINETAA888888888",
				CreditDateTime:            "2024-01-15T10:30:00.000+03:00",
				EndToEndIdentifier:        "AWINETAA777777777",
				TransactionIdentifier:     "AWINETAAFT66666666",
				InterBankSettlementAmount: "500.50",
				AccptanceDtatTime:         "2024-01-15T11:00:00.000Z",
				InstructedAmount:          "500.50",
				DebitAccountNumber:        "9876543210",
				CreditAccountNumber:       "1234567890",
				CreditAccountHolderName:   "Jane Smith",
				Narative:                  "Payment for services",
				DebiterInformation: DebiterInformation{
					Name:           "Test Debiter",
					StreetName:     "Test Street",
					BuildingNumber: "789",
					PostalCode:     "3000",
					TownName:       "Dire Dawa",
					Country:        "ET",
				},
				CreditInformation: CrediterInformation{
					Name:           "Test Crediter",
					StreetName:     "Another Street",
					BuildingNumber: "321",
					PostalCode:     "4000",
					City:           "Mekelle",
					TownName:       "Mekelle",
					Country:        "ET",
					AddressLine:    "Extra info",
				},
			},
			expect: []string{
				`<urn:Id>AWINETAA</urn:Id>`,
				`<urn:Id>CBETETAA</urn:Id>`,
				`<urn:BizMsgIdr>AWINETAA999999999</urn:BizMsgIdr>`,
				`<urn:CreDt>2024-01-15T00:00:00.000Z</urn:CreDt>`,
				`<urn1:MsgId>AWINETAA888888888</urn1:MsgId>`,
				`<urn1:CreDtTm>2024-01-15T10:30:00.000+03:00</urn1:CreDtTm>`,
				`<urn1:EndToEndId>AWINETAA777777777</urn1:EndToEndId>`,
				`<urn1:TxId>AWINETAAFT66666666</urn1:TxId>`,
				`<urn1:IntrBkSttlmAmt Ccy="ETB">500.50</urn1:IntrBkSttlmAmt>`,
				`<urn1:AccptncDtTm>2024-01-15T11:00:00.000Z</urn1:AccptncDtTm>`,
				`<urn1:InstdAmt Ccy="ETB">500.50</urn1:InstdAmt>`,
				`<urn1:Nm>Test Debiter</urn1:Nm>`,
				`<urn1:StrtNm>Test Street</urn1:StrtNm>`,
				`<urn1:BldgNb>789</urn1:BldgNb>`,
				`<urn1:PstCd>3000</urn1:PstCd>`,
				`<urn1:TwnNm>Dire Dawa</urn1:TwnNm>`,
				`<urn1:Id>9876543210</urn1:Id>`,
				`<urn1:Id>1234567890</urn1:Id>`,
				`<urn1:Nm>Test Crediter</urn1:Nm>`,
				`<urn1:StrtNm>Another Street</urn1:StrtNm>`,
				`<urn1:BldgNb>321</urn1:BldgNb>`,
				`<urn1:PstCd>4000</urn1:PstCd>`,
				`<urn1:TwnNm>Mekelle</urn1:TwnNm>`,
				`<urn1:AdrLine>Extra info</urn1:AdrLine>`,
				`<urn1:Ustrd>Payment for services</urn1:Ustrd>`,
			},
		},
		{
			name: "Validate QR Payment XML with empty strings",
			param: Params{
				DebitBankBIC:              "CBETETAA",
				CreditBankBIC:             "ETSETAA",
				BizMessageIdentifier:      "CBETETAA123456789",
				CreditDate:                "2023-06-24T00:00:00.000Z",
				MessageIdentifier:         "CBETETAA987654321",
				CreditDateTime:            "2023-06-24T00:00:00.000+03:00",
				EndToEndIdentifier:        "CBETETAA111111111",
				TransactionIdentifier:     "CBETETAAFT22222222",
				InterBankSettlementAmount: "0.00",
				AccptanceDtatTime:         "2023-06-24T12:00:00.000Z",
				InstructedAmount:          "0.00",
				DebitAccountNumber:        "",
				CreditAccountNumber:       "",
				CreditAccountHolderName:   "",
				Narative:                  "",
				DebiterInformation: DebiterInformation{
					Name:           "",
					StreetName:     "",
					BuildingNumber: "",
					PostalCode:     "",
					TownName:       "",
					Country:        "",
				},
				CreditInformation: CrediterInformation{
					Name:           "",
					StreetName:     "",
					BuildingNumber: "",
					PostalCode:     "",
					City:           "",
					TownName:       "",
					Country:        "",
					AddressLine:    "",
				},
			},
			expect: []string{
				`<soapenv:Envelope`,
				`<soapenv:Body>`,
				`<mb:Payment>`,
				`<urn:BizMsgIdr>CBETETAA123456789</urn:BizMsgIdr>`,
				`<urn1:IntrBkSttlmAmt Ccy="ETB">0.00</urn1:IntrBkSttlmAmt>`,
				`<urn1:InstdAmt Ccy="ETB">0.00</urn1:InstdAmt>`,
			},
		},
		{
			name: "Validate QR Payment XML structure elements",
			param: Params{
				DebitBankBIC:              "CBETETAA",
				CreditBankBIC:             "ETSETAA",
				BizMessageIdentifier:      "CBETETAA553572981",
				CreditDate:                "2023-06-24T00:00:00.000Z",
				MessageIdentifier:         "CBETETAA847772981",
				CreditDateTime:            "2023-06-24T00:00:00.000+03:00",
				EndToEndIdentifier:        "CBETETAA897809371",
				TransactionIdentifier:     "CBETETAAFT18978092",
				InterBankSettlementAmount: "1000.00",
				AccptanceDtatTime:         "2023-06-24T12:00:00.000Z",
				InstructedAmount:          "1000.00",
				DebitAccountNumber:        "1000000006924",
				CreditAccountNumber:       "1000382499388",
				CreditAccountHolderName:   "John Doe",
				Narative:                  "QR Payment Test",
				DebiterInformation: DebiterInformation{
					Name:           "Debiter Name",
					StreetName:     "Main Street",
					BuildingNumber: "123",
					PostalCode:     "1000",
					TownName:       "Addis Ababa",
					Country:        "ET",
				},
				CreditInformation: CrediterInformation{
					Name:           "Crediter Name",
					StreetName:     "Second Street",
					BuildingNumber: "456",
					PostalCode:     "2000",
					City:           "Addis Ababa",
					TownName:       "Addis Ababa",
					Country:        "ET",
					AddressLine:    "Additional Address Info",
				},
			},
			expect: []string{
				`<urn1:FIToFICstmrCdtTrf>`,
				`<urn1:GrpHdr>`,
				`<urn1:CdtTrfTxInf>`,
				`<urn1:PmtId>`,
				`<urn1:UltmtDbtr>`,
				`<urn1:Dbtr>`,
				`<urn1:DbtrAcct>`,
				`<urn1:DbtrAgt>`,
				`<urn1:CdtrAgt>`,
				`<urn1:Cdtr>`,
				`<urn1:CdtrAcct>`,
				`<urn1:UltmtCdtr>`,
				`<urn1:Purp>`,
				`<urn1:Tax>`,
				`<urn1:RmtInf>`,
				`<urn1:SttlmInf>`,
				`<urn1:PmtTpInf>`,
				`<urn1:LclInstrm>`,
				`<urn1:CtgyPurp>`,
				`<urn1:SttlmMtd>CLRG</urn1:SttlmMtd>`,
				`<urn1:Prtry>FP</urn1:Prtry>`,
				`<urn1:Prtry>CRTRM</urn1:Prtry>`,
				`<urn1:Prtry>C2BSQR</urn1:Prtry>`,
				`<urn1:ChrgBr>SLEV</urn1:ChrgBr>`,
				`<urn1:CtryOfRes>ET</urn1:CtryOfRes>`,
				`<urn1:Prtry>GOVTP</urn1:Prtry>`,
				`<urn1:TaxId>123456789</urn1:TaxId>`,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			xmlRequest := NewQrPayment(tc.param)

			// Validate XML is not empty
			assert.NotEmpty(t, xmlRequest, "XML request should not be empty")

			// Validate XML structure
			assert.Contains(t, xmlRequest, "<soapenv:Envelope", "Should contain SOAP envelope")
			assert.Contains(t, xmlRequest, "<soapenv:Body>", "Should contain SOAP body")
			assert.Contains(t, xmlRequest, "<mb:Payment>", "Should contain Payment element")

			// Validate all expected strings are present
			for _, expectedStr := range tc.expect {
				assert.Contains(t, xmlRequest, expectedStr, "XML should contain: %s", expectedStr)
			}

			// Validate that all required namespaces are present
			assert.Contains(t, xmlRequest, `xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"`)
			assert.Contains(t, xmlRequest, `xmlns:mb="http://MB_IPS"`)
			assert.Contains(t, xmlRequest, `xmlns:urn="urn:iso:std:iso:20022:tech:xsd:head.001.001.03"`)
			assert.Contains(t, xmlRequest, `xmlns:urn1="urn:iso:std:iso:20022:tech:xsd:pacs.008.001.10"`)

			// Validate that the message definition identifier is correct
			assert.Contains(t, xmlRequest, `<urn:MsgDefIdr>pacs.008.001.10</urn:MsgDefIdr>`)

			// Validate that number of transactions is set to 1
			assert.Contains(t, xmlRequest, `<urn1:NbOfTxs>1</urn1:NbOfTxs>`)

			// Validate currency is ETB
			assert.Contains(t, xmlRequest, `Ccy="ETB"`)

			// Validate hardcoded values
			assert.Contains(t, xmlRequest, `<urn1:Nm>MSCWT</urn1:Nm>`, "Should contain hardcoded MSCWT")
			assert.Contains(t, xmlRequest, `<urn1:Id>MOBN</urn1:Id>`, "Should contain hardcoded MOBN")
			assert.Contains(t, xmlRequest, `<urn1:Prtry>LPNB</urn1:Prtry>`, "Should contain hardcoded LPNB")
			assert.Contains(t, xmlRequest, `<urn1:Prtry>ACCT</urn1:Prtry>`, "Should contain hardcoded ACCT")
			assert.Contains(t, xmlRequest, `<urn1:Issr>C</urn1:Issr>`, "Should contain hardcoded Issr C")
			assert.Contains(t, xmlRequest, `<urn1:Issr>ATM</urn1:Issr>`, "Should contain hardcoded Issr ATM")
			assert.Contains(t, xmlRequest, `<urn1:ChanlTp>QRCP</urn1:ChanlTp>`, "Should contain hardcoded QRCP")
			assert.Contains(t, xmlRequest, `<urn1:Id>QRCPS</urn1:Id>`, "Should contain hardcoded QRCPS")
		})
	}
}

func TestNewQrPayment_ParameterSubstitution(t *testing.T) {
	param := Params{
		DebitBankBIC:              "TESTBIC1",
		CreditBankBIC:             "TESTBIC2",
		BizMessageIdentifier:      "TESTBIZ123",
		CreditDate:                "2023-01-01T00:00:00.000Z",
		MessageIdentifier:         "TESTMSG456",
		CreditDateTime:            "2023-01-01T10:00:00.000+03:00",
		EndToEndIdentifier:        "TESTEND789",
		TransactionIdentifier:     "TESTTX012",
		InterBankSettlementAmount: "1234.56",
		AccptanceDtatTime:         "2023-01-01T11:00:00.000Z",
		InstructedAmount:          "1234.56",
		DebitAccountNumber:        "DEBIT123",
		CreditAccountNumber:       "CREDIT456",
		CreditAccountHolderName:   "Test Holder",
		Narative:                  "Test Narrative",
		DebiterInformation: DebiterInformation{
			Name:           "Test Debiter",
			StreetName:     "Test Street",
			BuildingNumber: "999",
			PostalCode:     "9999",
			TownName:       "Test Town",
			Country:        "US",
		},
		CreditInformation: CrediterInformation{
			Name:           "Test Crediter",
			StreetName:     "Credit Street",
			BuildingNumber: "888",
			PostalCode:     "8888",
			City:           "Credit City",
			TownName:       "Credit Town",
			Country:        "CA",
			AddressLine:    "Credit Address Line",
		},
	}

	xmlRequest := NewQrPayment(param)

	// Verify all parameters are correctly substituted
	assert.Contains(t, xmlRequest, "TESTBIC1", "Should contain DebitBankBIC")
	assert.Contains(t, xmlRequest, "TESTBIC2", "Should contain CreditBankBIC")
	assert.Contains(t, xmlRequest, "TESTBIZ123", "Should contain BizMessageIdentifier")
	assert.Contains(t, xmlRequest, "2023-01-01T00:00:00.000Z", "Should contain CreditDate")
	assert.Contains(t, xmlRequest, "TESTMSG456", "Should contain MessageIdentifier")
	assert.Contains(t, xmlRequest, "2023-01-01T10:00:00.000+03:00", "Should contain CreditDateTime")
	assert.Contains(t, xmlRequest, "TESTEND789", "Should contain EndToEndIdentifier")
	assert.Contains(t, xmlRequest, "TESTTX012", "Should contain TransactionIdentifier")
	assert.Contains(t, xmlRequest, "1234.56", "Should contain InterBankSettlementAmount")
	assert.Contains(t, xmlRequest, "2023-01-01T11:00:00.000Z", "Should contain AccptanceDtatTime")
	assert.Contains(t, xmlRequest, "DEBIT123", "Should contain DebitAccountNumber")
	assert.Contains(t, xmlRequest, "CREDIT456", "Should contain CreditAccountNumber")
	assert.Contains(t, xmlRequest, "Test Debiter", "Should contain DebiterInformation.Name")
	assert.Contains(t, xmlRequest, "Test Street", "Should contain DebiterInformation.StreetName")
	assert.Contains(t, xmlRequest, "999", "Should contain DebiterInformation.BuildingNumber")
	assert.Contains(t, xmlRequest, "9999", "Should contain DebiterInformation.PostalCode")
	assert.Contains(t, xmlRequest, "Test Town", "Should contain DebiterInformation.TownName")
	assert.Contains(t, xmlRequest, "US", "Should contain DebiterInformation.Country")
	assert.Contains(t, xmlRequest, "Test Crediter", "Should contain CreditInformation.Name")
	assert.Contains(t, xmlRequest, "Credit Street", "Should contain CreditInformation.StreetName")
	assert.Contains(t, xmlRequest, "888", "Should contain CreditInformation.BuildingNumber")
	assert.Contains(t, xmlRequest, "8888", "Should contain CreditInformation.PostalCode")
	assert.Contains(t, xmlRequest, "Credit Town", "Should contain CreditInformation.TownName")
	assert.Contains(t, xmlRequest, "CA", "Should contain CreditInformation.Country")
	assert.Contains(t, xmlRequest, "Credit Address Line", "Should contain CreditInformation.AddressLine")
	assert.Contains(t, xmlRequest, "Test Narrative", "Should contain Narative")
}

func TestParseQrPaymentSOAP(t *testing.T) {
	tests := []struct {
		name           string
		xmlData        string
		expectedResult *QrPaymentResult
		expectError    bool
	}{
		{
			name: "Successful QR Payment response with ACSC status",
			xmlData: `<?xml version="1.0" encoding="utf-8"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
    <soapenv:Body>
        <NS1:PaymentResponse xmlns:NS1="http://MB_IPS">
            <output1>
                <NS2:AppHdr xmlns:NS2="urn:iso:std:iso:20022:tech:xsd:head.001.001.03">
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
                                    <NS2:Id>CBETETAA</NS2:Id>
                                </NS2:Othr>
                            </NS2:FinInstnId>
                        </NS2:FIId>
                    </NS2:To>
                    <NS2:BizMsgIdr>ETSETAA535018266346350319329</NS2:BizMsgIdr>
                    <NS2:MsgDefIdr>pacs.002.001.12</NS2:MsgDefIdr>
                    <NS2:CreDt>2025-12-16T18:35:55.055Z</NS2:CreDt>
                    <NS2:Rltd>
                        <NS2:Fr>
                            <NS2:FIId>
                                <NS2:FinInstnId>
                                    <NS2:Othr>
                                        <NS2:Id>CBETETAA</NS2:Id>
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
                        <NS2:BizMsgIdr>CBETETAA1947809198</NS2:BizMsgIdr>
                        <NS2:MsgDefIdr>pacs.008.001.10</NS2:MsgDefIdr>
                        <NS2:CreDt>2024-06-05T00:00:00.000Z</NS2:CreDt>
                    </NS2:Rltd>
                </NS2:AppHdr>
                <NS3:Document xmlns:NS3="urn:iso:std:iso:20022:tech:xsd:pacs.002.001.12">
                    <NS3:FIToFIPmtStsRpt>
                        <NS3:GrpHdr>
                            <NS3:MsgId>ETSETAA535018266346350319329</NS3:MsgId>
                            <NS3:CreDtTm>2025-12-16T18:35:55.960+03:00</NS3:CreDtTm>
                            <NS3:InstgAgt>
                                <NS3:FinInstnId>
                                    <NS3:Othr>
                                        <NS3:Id>CBETETAA</NS3:Id>
                                    </NS3:Othr>
                                </NS3:FinInstnId>
                            </NS3:InstgAgt>
                            <NS3:InstdAgt>
                                <NS3:FinInstnId>
                                    <NS3:Othr>
                                        <NS3:Id>ETSETAA</NS3:Id>
                                    </NS3:Othr>
                                </NS3:FinInstnId>
                            </NS3:InstdAgt>
                        </NS3:GrpHdr>
                        <NS3:OrgnlGrpInfAndSts>
                            <NS3:OrgnlMsgId>CBETETAA1947809198</NS3:OrgnlMsgId>
                            <NS3:OrgnlMsgNmId>pacs.008.001.10</NS3:OrgnlMsgNmId>
                            <NS3:OrgnlCreDtTm>2024-06-05T00:00:00.000+03:00</NS3:OrgnlCreDtTm>
                        </NS3:OrgnlGrpInfAndSts>
                        <NS3:TxInfAndSts>
                            <NS3:OrgnlEndToEndId>CBETETAA098795427</NS3:OrgnlEndToEndId>
                            <NS3:OrgnlTxId>CBETETAAFT098795438</NS3:OrgnlTxId>
                            <NS3:TxSts>ACSC</NS3:TxSts>
                            <NS3:AccptncDtTm>2024-06-05T17:42:39.071+06:00</NS3:AccptncDtTm>
                            <NS3:OrgnlTxRef>
                                <NS3:IntrBkSttlmAmt Ccy="ETB">10</NS3:IntrBkSttlmAmt>
                                <NS3:Amt>
                                    <NS3:InstdAmt Ccy="ETB">10</NS3:InstdAmt>
                                </NS3:Amt>
                                <NS3:PmtTpInf>
                                    <NS3:CtgyPurp>
                                        <NS3:Prtry>C2BSQR</NS3:Prtry>
                                    </NS3:CtgyPurp>
                                </NS3:PmtTpInf>
                                <NS3:RmtInf>
                                    <NS3:Ustrd>Transferring my funds</NS3:Ustrd>
                                    <NS3:Strd>
                                        <NS3:RfrdDocInf>
                                            <NS3:Tp>
                                                <NS3:CdOrPrtry>
                                                    <NS3:Prtry>0123456789</NS3:Prtry>
                                                </NS3:CdOrPrtry>
                                            </NS3:Tp>
                                        </NS3:RfrdDocInf>
                                    </NS3:Strd>
                                </NS3:RmtInf>
                                <NS3:Dbtr>
                                    <NS3:Pty>
                                        <NS3:Nm>MSCWT</NS3:Nm>
                                        <NS3:PstlAdr>
                                            <NS3:StrtNm>Tito St.</NS3:StrtNm>
                                        </NS3:PstlAdr>
                                        <NS3:Id>
                                            <NS3:PrvtId>
                                                <NS3:Othr>
                                                    <NS3:Id>MOBN</NS3:Id>
                                                    <NS3:SchmeNm>
                                                        <NS3:Prtry>LPNB</NS3:Prtry>
                                                    </NS3:SchmeNm>
                                                </NS3:Othr>
                                            </NS3:PrvtId>
                                        </NS3:Id>
                                        <NS3:CtctDtls></NS3:CtctDtls>
                                    </NS3:Pty>
                                </NS3:Dbtr>
                                <NS3:DbtrAcct>
                                    <NS3:Id>
                                        <NS3:Othr>
                                            <NS3:Id>1234567890</NS3:Id>
                                            <NS3:SchmeNm>
                                                <NS3:Prtry>ACCT</NS3:Prtry>
                                            </NS3:SchmeNm>
                                            <NS3:Issr>C</NS3:Issr>
                                        </NS3:Othr>
                                    </NS3:Id>
                                </NS3:DbtrAcct>
                                <NS3:Cdtr>
                                    <NS3:Pty>
                                        <NS3:Nm>Merchant Name</NS3:Nm>
                                        <NS3:PstlAdr>
                                            <NS3:StrtNm>Tito St.</NS3:StrtNm>
                                            <NS3:BldgNb>17</NS3:BldgNb>
                                            <NS3:PstCd>18444</NS3:PstCd>
                                            <NS3:TwnNm>Addis Ababa</NS3:TwnNm>
                                            <NS3:Ctry>AA</NS3:Ctry>
                                            <NS3:AdrLine>Kazanchis</NS3:AdrLine>
                                        </NS3:PstlAdr>
                                        <NS3:CtryOfRes>ET</NS3:CtryOfRes>
                                        <NS3:CtctDtls>
                                            <NS3:Nm>Merchant Name</NS3:Nm>
                                            <NS3:Dept>MerchantLabel</NS3:Dept>
                                            <NS3:Othr>
                                                <NS3:ChanlTp>QRCP</NS3:ChanlTp>
                                                <NS3:Id>QRCPS</NS3:Id>
                                            </NS3:Othr>
                                        </NS3:CtctDtls>
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
                                <NS3:UltmtCdtr>
                                    <NS3:Pty>
                                        <NS3:Id>
                                            <NS3:PrvtId>
                                                <NS3:Othr>
                                                    <NS3:Id>MOBN</NS3:Id>
                                                    <NS3:SchmeNm>
                                                        <NS3:Prtry>MOBN</NS3:Prtry>
                                                    </NS3:SchmeNm>
                                                </NS3:Othr>
                                            </NS3:PrvtId>
                                        </NS3:Id>
                                    </NS3:Pty>
                                </NS3:UltmtCdtr>
                                <NS3:Purp>
                                    <NS3:Prtry>GOVTP</NS3:Prtry>
                                </NS3:Purp>
                            </NS3:OrgnlTxRef>
                        </NS3:TxInfAndSts>
                    </NS3:FIToFIPmtStsRpt>
                </NS3:Document>
            </output1>
        </NS1:PaymentResponse>
    </soapenv:Body>
</soapenv:Envelope>`,
			expectedResult: &QrPaymentResult{
				Success: true,
				Detail: &QrPaymentDetail{
					FromBankBIC:                   "FP",
					ToBankBIC:                     "CBETETAA",
					BizMessageIdentifier:          "ETSETAA535018266346350319329",
					CreditDate:                    "2025-12-16T18:35:55.055Z",
					RelatedFromBankBIC:            "CBETETAA",
					RelatedToBankBIC:              "ETSETAA",
					RelatedBizMessageIdentifier:   "CBETETAA1947809198",
					RelatedMessageDefinitionId:    "pacs.008.001.10",
					RelatedCreditDate:             "2024-06-05T00:00:00.000Z",
					MessageId:                     "ETSETAA535018266346350319329",
					CreationDateTime:              "2025-12-16T18:35:55.960+03:00",
					InstructingAgent:              "CBETETAA",
					InstructedAgent:               "ETSETAA",
					OriginalMessageIdentifier:     "CBETETAA1947809198",
					OriginalMessageNameId:         "pacs.008.001.10",
					OriginalCreditDateTime:        "2024-06-05T00:00:00.000+03:00",
					OriginalEndToEndIdentifier:    "CBETETAA098795427",
					OriginalTransactionIdentifier: "CBETETAAFT098795438",
					TransactionStatus:             "ACSC",
					AcceptanceDateTime:            "2024-06-05T17:42:39.071+06:00",
					InterBankSettlementAmount:     "10",
					InterBankSettlementCurrency:   "ETB",
					InstructedAmount:              "10",
					InstructedAmountCurrency:      "ETB",
					PaymentTypeCategoryPurpose:    "C2BSQR",
					RemittanceInformation:         "Transferring my funds",
					RemittanceReference:           "0123456789",
					DebtorName:                    "MSCWT",
					DebtorStreetName:              "Tito St.",
					DebtorAccountNumber:           "1234567890",
					CreditorName:                  "Merchant Name",
					CreditorStreetName:            "Tito St.",
					CreditorBuildingNumber:        "17",
					CreditorPostalCode:            "18444",
					CreditorTownName:              "Addis Ababa",
					CreditorCountry:               "AA",
					CreditorAddressLine:           "Kazanchis",
					CreditorCountryOfResidence:    "ET",
					CreditorContactName:           "Merchant Name",
					CreditorContactDepartment:     "MerchantLabel",
					CreditorAccountNumber:         "1234567890",
					Purpose:                       "GOVTP",
				},
			},
			expectError: false,
		},
		{
			name: "Rejected QR Payment response with RJCT status",
			xmlData: `<?xml version="1.0" encoding="utf-8"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
    <soapenv:Body>
        <NS1:PaymentResponse xmlns:NS1="http://MB_IPS">
            <output1>
                <NS2:AppHdr xmlns:NS2="urn:iso:std:iso:20022:tech:xsd:head.001.001.03">
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
                                    <NS2:Id>CBETETAA</NS2:Id>
                                </NS2:Othr>
                            </NS2:FinInstnId>
                        </NS2:FIId>
                    </NS2:To>
                    <NS2:BizMsgIdr>ETSETAA535018266346350319329</NS2:BizMsgIdr>
                    <NS2:MsgDefIdr>pacs.002.001.12</NS2:MsgDefIdr>
                    <NS2:CreDt>2025-12-16T18:35:55.055Z</NS2:CreDt>
                </NS2:AppHdr>
                <NS3:Document xmlns:NS3="urn:iso:std:iso:20022:tech:xsd:pacs.002.001.12">
                    <NS3:FIToFIPmtStsRpt>
                        <NS3:GrpHdr>
                            <NS3:MsgId>ETSETAA535018266346350319329</NS3:MsgId>
                            <NS3:CreDtTm>2025-12-16T18:35:55.960+03:00</NS3:CreDtTm>
                            <NS3:InstgAgt>
                                <NS3:FinInstnId>
                                    <NS3:Othr>
                                        <NS3:Id>CBETETAA</NS3:Id>
                                    </NS3:Othr>
                                </NS3:FinInstnId>
                            </NS3:InstgAgt>
                            <NS3:InstdAgt>
                                <NS3:FinInstnId>
                                    <NS3:Othr>
                                        <NS3:Id>ETSETAA</NS3:Id>
                                    </NS3:Othr>
                                </NS3:FinInstnId>
                            </NS3:InstdAgt>
                        </NS3:GrpHdr>
                        <NS3:OrgnlGrpInfAndSts>
                            <NS3:OrgnlMsgId>CBETETAA1947809198</NS3:OrgnlMsgId>
                            <NS3:OrgnlMsgNmId>pacs.008.001.10</NS3:OrgnlMsgNmId>
                            <NS3:OrgnlCreDtTm>2024-06-05T00:00:00.000+03:00</NS3:OrgnlCreDtTm>
                        </NS3:OrgnlGrpInfAndSts>
                        <NS3:TxInfAndSts>
                            <NS3:OrgnlEndToEndId>CBETETAA098795427</NS3:OrgnlEndToEndId>
                            <NS3:OrgnlTxId>CBETETAAFT098795438</NS3:OrgnlTxId>
                            <NS3:TxSts>RJCT</NS3:TxSts>
                            <NS3:AccptncDtTm>2024-06-05T17:42:39.071+06:00</NS3:AccptncDtTm>
                            <NS3:StsRsnInf>
                                <NS3:Orgtr>
                                    <NS3:Id>
                                        <NS3:OrgId>
                                            <NS3:Othr>
                                                <NS3:Id>FP</NS3:Id>
                                            </NS3:Othr>
                                        </NS3:OrgId>
                                    </NS3:Id>
                                </NS3:Orgtr>
                                <NS3:Rsn>
                                    <NS3:Prtry>DTID</NS3:Prtry>
                                </NS3:Rsn>
                                <NS3:AddtlInf>TxID duplicate in IPS</NS3:AddtlInf>
                            </NS3:StsRsnInf>
                        </NS3:TxInfAndSts>
                    </NS3:FIToFIPmtStsRpt>
                </NS3:Document>
            </output1>
        </NS1:PaymentResponse>
    </soapenv:Body>
</soapenv:Envelope>`,
			expectedResult: &QrPaymentResult{
				Success:  false,
				Messages: []string{"TxID duplicate in IPS"},
				Detail: &QrPaymentDetail{
					TransactionStatus:          "RJCT",
					StatusReason:               "DTID",
					StatusReasonAdditionalInfo: "TxID duplicate in IPS",
				},
			},
			expectError: false,
		},
		{
			name:        "Invalid XML",
			xmlData:     `<invalid xml>`,
			expectError: true,
		},
		{
			name: "Missing PaymentResponse",
			xmlData: `<?xml version="1.0" encoding="utf-8"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
    <soapenv:Body>
        <OtherResponse>
        </OtherResponse>
    </soapenv:Body>
</soapenv:Envelope>`,
			expectedResult: &QrPaymentResult{
				Success:  false,
				Messages: []string{"Invalid Response: Missing PaymentResponse"},
			},
			expectError: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ParseQrPaymentSOAP(tc.xmlData)

			if tc.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)

				if tc.expectedResult != nil {
					assert.Equal(t, tc.expectedResult.Success, result.Success)
					assert.Equal(t, tc.expectedResult.Messages, result.Messages)

					if tc.expectedResult.Detail != nil {
						assert.NotNil(t, result.Detail)
						expected := tc.expectedResult.Detail
						actual := result.Detail

						if expected.FromBankBIC != "" {
							assert.Equal(t, expected.FromBankBIC, actual.FromBankBIC)
						}
						if expected.ToBankBIC != "" {
							assert.Equal(t, expected.ToBankBIC, actual.ToBankBIC)
						}
						if expected.BizMessageIdentifier != "" {
							assert.Equal(t, expected.BizMessageIdentifier, actual.BizMessageIdentifier)
						}
						if expected.TransactionStatus != "" {
							assert.Equal(t, expected.TransactionStatus, actual.TransactionStatus)
						}
						if expected.StatusReason != "" {
							assert.Equal(t, expected.StatusReason, actual.StatusReason)
						}
						if expected.StatusReasonAdditionalInfo != "" {
							assert.Equal(t, expected.StatusReasonAdditionalInfo, actual.StatusReasonAdditionalInfo)
						}
						if expected.OriginalEndToEndIdentifier != "" {
							assert.Equal(t, expected.OriginalEndToEndIdentifier, actual.OriginalEndToEndIdentifier)
						}
						if expected.OriginalTransactionIdentifier != "" {
							assert.Equal(t, expected.OriginalTransactionIdentifier, actual.OriginalTransactionIdentifier)
						}
						if expected.CreditorName != "" {
							assert.Equal(t, expected.CreditorName, actual.CreditorName)
						}
						if expected.CreditorAccountNumber != "" {
							assert.Equal(t, expected.CreditorAccountNumber, actual.CreditorAccountNumber)
						}
						if expected.RemittanceInformation != "" {
							assert.Equal(t, expected.RemittanceInformation, actual.RemittanceInformation)
						}
					} else {
						assert.Nil(t, result.Detail)
					}
				}
			}
		})
	}
}

func TestParseQrPaymentSOAP_WithSampleFile(t *testing.T) {
	// Test with the actual sample.xml file if it exists
	xmlData, err := os.ReadFile("../../../sample.xml")
	if err != nil {
		t.Skip("sample.xml file not found, skipping test")
		return
	}

	result, err := ParseQrPaymentSOAP(string(xmlData))
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// The sample XML has RJCT status, so it should not be successful
	assert.False(t, result.Success)
	assert.NotNil(t, result.Detail)
	assert.Equal(t, "RJCT", result.Detail.TransactionStatus)
	assert.Equal(t, "DTID", result.Detail.StatusReason)
	assert.Contains(t, result.Detail.StatusReasonAdditionalInfo, "TxID duplicate")

	// For rejected transactions, some fields might be empty if not in OrgnlTxRef
	// But we should still have the basic transaction info
	assert.Equal(t, "CBETETAA098795427", result.Detail.OriginalEndToEndIdentifier)
	assert.Equal(t, "CBETETAAFT098795438", result.Detail.OriginalTransactionIdentifier)

	// FromBankBIC and ToBankBIC might be empty for rejected transactions
	// if they're only in OrgnlTxRef which might not be fully populated
	if result.Detail.FromBankBIC != "" {
		assert.Equal(t, "FP", result.Detail.FromBankBIC)
	}
	if result.Detail.ToBankBIC != "" {
		assert.Equal(t, "CBETETAA", result.Detail.ToBankBIC)
	}
}
