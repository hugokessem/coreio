package cardreplace

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCardReplace(t *testing.T) {
	tests := []struct {
		name   string
		param  Params
		expect []string
	}{
		{
			name: "Validate card replace XML generation",
			param: Params{
				Username:          "SUPERAPP",
				Password:          "123456",
				AccountNumber:     "1000000006924",
				BranchCode:        "ET0010222",
				PhoneNumber:       "+251913323918",
				CardType:          "VIEL",
				ReplacementReason: "TEST",
				ProductType:       "PLASSTIC",
			},
			expect: []string{
				`<password>123456</password>`,
				`<userName>SUPERAPP</userName>`,
				`<atm:ACCOUNT>1000000006924</atm:ACCOUNT>`,
				`<atm:BRANCHCODE>ET0010222</atm:BRANCHCODE>`,
				`<atm:PHONENO>+251913323918</atm:PHONENO>`,
				`<atm:CARDTYPE>VIEL</atm:CARDTYPE>`,
				`<atm:ReplacementReason>TEST</atm:ReplacementReason>`,
				`<atm:ProductType>PLASSTIC</atm:ProductType>`,
				`<soapenv:Envelope`,
				`<soapenv:Body>`,
				`<cbes:ATMCardReplacementRequest>`,
			},
		},
		{
			name: "Validate card replace with different values",
			param: Params{
				Username:          "TESTUSER",
				Password:          "PASSWORD123",
				AccountNumber:     "2000000000001",
				BranchCode:        "ET0010001",
				PhoneNumber:       "+251911111111",
				CardType:          "WLCD",
				ReplacementReason: "LOST",
				ProductType:       "VIRTUAL",
			},
			expect: []string{
				`<password>PASSWORD123</password>`,
				`<userName>TESTUSER</userName>`,
				`<atm:ACCOUNT>2000000000001</atm:ACCOUNT>`,
				`<atm:BRANCHCODE>ET0010001</atm:BRANCHCODE>`,
				`<atm:PHONENO>+251911111111</atm:PHONENO>`,
				`<atm:CARDTYPE>WLCD</atm:CARDTYPE>`,
				`<atm:ReplacementReason>LOST</atm:ReplacementReason>`,
				`<atm:ProductType>VIRTUAL</atm:ProductType>`,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			xmlRequest := NewCardReplace(tc.param)

			// Validate XML structure
			assert.Contains(t, xmlRequest, "<soapenv:Envelope")
			assert.Contains(t, xmlRequest, "<soapenv:Body>")
			assert.Contains(t, xmlRequest, "<cbes:ATMCardReplacementRequest>")

			// Validate all expected strings are present
			for _, expectedStr := range tc.expect {
				assert.Contains(t, xmlRequest, expectedStr)
			}

			// Validate XML is not empty
			assert.NotEmpty(t, xmlRequest)
		})
	}
}

func TestParseCardReplaceResponse(t *testing.T) {
	tests := []struct {
		name             string
		xmlData          string
		expectedSuccess  bool
		expectedError    bool
		expectedDetail   bool
		expectedMessages []string
	}{
		{
			name: "Parse successful response",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns26:ATMCardReplacementRequestResponse xmlns:ns26="http://temenos.com/CBESUPERAPP">
            <Status>
                <transactionId>102134334628</transactionId>
                <messageId></messageId>
                <successIndicator>Success</successIndicator>
                <application>ATM.CARD.REG.DET</application>
            </Status>
            <ATMCARDREGDETType id="102134334628">
                <ACCOUNT>1000000006924</ACCOUNT>
                <ACCOUNTTITLE1>ABIY HAILEYESUS MENGISTU</ACCOUNTTITLE1>
                <ADDRESS>AABO</ADDRESS>
                <BRANCHNAME>Addis Ababa Branch</BRANCHNAME>
                <ACOPENDATE>20080703</ACOPENDATE>
                <RESIDENCE>ET</RESIDENCE>
                <INDUSTRY>1006</INDUSTRY>
                <BRANCHCODE>ET0010222</BRANCHCODE>
                <PHONENO>+251913323918</PHONENO>
                <CARDTYPE>VIEL</CARDTYPE>
                <SEX>M</SEX>
                <CIVILSTATUS>MARRIED</CIVILSTATUS>
                <HOLDERNO>102134334628</HOLDERNO>
                <CURRNO>1</CURRNO>
                <gDATETIME>
                    <DATETIME>2512021442</DATETIME>
                </gDATETIME>
                <AUTHORISER>3457_SUPERAPP.1_OFS_GCS</AUTHORISER>
                <COCODE>ET0010001</COCODE>
                <DEPTCODE>1</DEPTCODE>
                <DATE>20211209</DATE>
                <REQUESTTYPE>REPLACEMENT</REQUESTTYPE>
                <REPLREASON>TEST</REPLREASON>
                <PRODUCTTYPE>PLASSTIC</PRODUCTTYPE>
                <VIRTUALFLAG>YES</VIRTUALFLAG>
            </ATMCARDREGDETType>
        </ns26:ATMCardReplacementRequestResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess: true,
			expectedError:   false,
			expectedDetail:  true,
		},
		{
			name: "Parse response with failure status",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns26:ATMCardReplacementRequestResponse xmlns:ns26="http://temenos.com/CBESUPERAPP">
            <Status>
                <transactionId>102134334628</transactionId>
                <messageId></messageId>
                <successIndicator>Failure</successIndicator>
                <application>ATM.CARD.REG.DET</application>
            </Status>
        </ns26:ATMCardReplacementRequestResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess:  false,
			expectedError:    false,
			expectedDetail:   false,
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
			expectedDetail:  false,
		},
		{
			name: "Parse response without Status",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns26:ATMCardReplacementRequestResponse xmlns:ns26="http://temenos.com/CBESUPERAPP">
        </ns26:ATMCardReplacementRequestResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess:  false,
			expectedError:    false,
			expectedDetail:   false,
			expectedMessages: []string{"Missing Status"},
		},
		{
			name: "Parse response without ATMCardReplacementRequestDetail",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns26:ATMCardReplacementRequestResponse xmlns:ns26="http://temenos.com/CBESUPERAPP">
            <Status>
                <transactionId>102134334628</transactionId>
                <messageId></messageId>
                <successIndicator>Success</successIndicator>
                <application>ATM.CARD.REG.DET</application>
            </Status>
        </ns26:ATMCardReplacementRequestResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess:  false,
			expectedError:    false,
			expectedDetail:   false,
			expectedMessages: []string{"Missing ATMCardReplacementRequestDetail"},
		},
		{
			name: "Parse response with invalid response structure",
			xmlData: `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <OtherResponse xmlns:ns26="http://temenos.com/CBESUPERAPP">
        </OtherResponse>
    </S:Body>
</S:Envelope>`,
			expectedSuccess:  false,
			expectedError:    false,
			expectedDetail:   false,
			expectedMessages: []string{"Invalid response"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ParseCardReplaceResponse(tc.xmlData)

			if tc.expectedError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tc.expectedSuccess, result.Success)

				if tc.expectedDetail {
					assert.NotNil(t, result.Detail)
					if result.Detail != nil {
						assert.NotEmpty(t, result.Detail.AccountNumber)
					}
				} else {
					if result != nil {
						assert.Nil(t, result.Detail)
					}
				}

				if len(tc.expectedMessages) > 0 {
					assert.Equal(t, tc.expectedMessages, result.Messages)
				}
			}
		})
	}
}

func TestParseCardReplaceResponse_DetailFields(t *testing.T) {
	xmlData := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns26:ATMCardReplacementRequestResponse xmlns:ns26="http://temenos.com/CBESUPERAPP">
            <Status>
                <transactionId>102134334628</transactionId>
                <messageId></messageId>
                <successIndicator>Success</successIndicator>
                <application>ATM.CARD.REG.DET</application>
            </Status>
            <ATMCARDREGDETType id="102134334628">
                <ACCOUNT>1000000006924</ACCOUNT>
                <ACCOUNTTITLE1>ABIY HAILEYESUS MENGISTU</ACCOUNTTITLE1>
                <ADDRESS>AABO</ADDRESS>
                <BRANCHNAME>Addis Ababa Branch</BRANCHNAME>
                <ACOPENDATE>20080703</ACOPENDATE>
                <RESIDENCE>ET</RESIDENCE>
                <INDUSTRY>1006</INDUSTRY>
                <BRANCHCODE>ET0010222</BRANCHCODE>
                <PHONENO>+251913323918</PHONENO>
                <CARDTYPE>VIEL</CARDTYPE>
                <SEX>M</SEX>
                <CIVILSTATUS>MARRIED</CIVILSTATUS>
                <HOLDERNO>102134334628</HOLDERNO>
                <CURRNO>1</CURRNO>
                <INPUTTER>3457_SUPERAPP.1__OFS_GCS</INPUTTER>
                <gDATETIME>
                    <DATETIME>2512021442</DATETIME>
                </gDATETIME>
                <AUTHORISER>3457_SUPERAPP.1_OFS_GCS</AUTHORISER>
                <COCODE>ET0010001</COCODE>
                <DEPTCODE>1</DEPTCODE>
                <DATE>20211209</DATE>
                <REQUESTTYPE>REPLACEMENT</REQUESTTYPE>
                <REPLREASON>TEST</REPLREASON>
                <PRODUCTTYPE>PLASSTIC</PRODUCTTYPE>
                <VIRTUALFLAG>YES</VIRTUALFLAG>
            </ATMCARDREGDETType>
        </ns26:ATMCardReplacementRequestResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseCardReplaceResponse(xmlData)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Success)
	assert.NotNil(t, result.Detail)

	if result.Detail != nil {
		assert.Equal(t, "1000000006924", result.Detail.AccountNumber)
		assert.Equal(t, "ABIY HAILEYESUS MENGISTU", result.Detail.AccountHolderName)
		assert.Equal(t, "AABO", result.Detail.Address)
		assert.Equal(t, "Addis Ababa Branch", result.Detail.BranchName)
		assert.Equal(t, "ET0010222", result.Detail.BranchCode)
		assert.Equal(t, "20080703", result.Detail.OpenDate)
		assert.Equal(t, "ET", result.Detail.Residence)
		assert.Equal(t, "1006", result.Detail.Industry)
		assert.Equal(t, "+251913323918", result.Detail.PhoneNumber)
		assert.Equal(t, "VIEL", result.Detail.CardType)
		assert.Equal(t, "M", result.Detail.Sex)
		assert.Equal(t, "MARRIED", result.Detail.CivilStatus)
		assert.Equal(t, "102134334628", result.Detail.HolderNumber)
		assert.Equal(t, "1", result.Detail.CurruntNumber)
		assert.Equal(t, "2512021442", result.Detail.GDatetime.DateTime)
		assert.Equal(t, "3457_SUPERAPP.1__OFS_GCS", result.Detail.Inputter)
		assert.Equal(t, "3457_SUPERAPP.1_OFS_GCS", result.Detail.Authoriser)
		assert.Equal(t, "ET0010001", result.Detail.CoCode)
		assert.Equal(t, "1", result.Detail.DeptCode)
		assert.Equal(t, "20211209", result.Detail.Date)
		assert.Equal(t, "REPLACEMENT", result.Detail.RequestType)
		assert.Equal(t, "TEST", result.Detail.ReplacementReason)
		assert.Equal(t, "PLASSTIC", result.Detail.ProductType)
		assert.Equal(t, "YES", result.Detail.VirtualFlag)
	}
}

func TestParseCardReplaceResponse_CaseInsensitiveSuccess(t *testing.T) {
	xmlData := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns26:ATMCardReplacementRequestResponse xmlns:ns26="http://temenos.com/CBESUPERAPP">
            <Status>
                <transactionId>102134334628</transactionId>
                <messageId></messageId>
                <successIndicator>SUCCESS</successIndicator>
                <application>ATM.CARD.REG.DET</application>
            </Status>
            <ATMCARDREGDETType id="102134334628">
                <ACCOUNT>1000000006924</ACCOUNT>
                <ACCOUNTTITLE1>ABIY HAILEYESUS MENGISTU</ACCOUNTTITLE1>
                <BRANCHCODE>ET0010222</BRANCHCODE>
                <PHONENO>+251913323918</PHONENO>
                <CARDTYPE>VIEL</CARDTYPE>
            </ATMCARDREGDETType>
        </ns26:ATMCardReplacementRequestResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseCardReplaceResponse(xmlData)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Success)
	assert.NotNil(t, result.Detail)
}
