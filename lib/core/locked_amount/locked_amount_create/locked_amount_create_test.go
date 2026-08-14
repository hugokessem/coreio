package lockedamountcreate

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateLockedAmountGeneratedXML(t *testing.T) {
	test := []struct {
		name   string
		param  Params
		expect []string
	}{
		{
			name: "Validate Create Locked Amount",
			param: Params{
				Username:      "SUPERAPP",
				Password:      "123456",
				AccountNumber: "1000000006924",
				Description:   "3 Click",
				From:          "20251108",
				To:            "20251111",
				LockedAmount:  "1200",
			},
			expect: []string{
				`<acl:ACCOUNTNUMBER>1000000006924</acl:ACCOUNTNUMBER>`,
				`<acl:DESCRIPTION>3 Click</acl:DESCRIPTION>`,
				`<acl:FROMDATE>20251108</acl:FROMDATE>`,
				`<acl:TODATE>20251111</acl:TODATE>`,
				`<acl:LOCKEDAMOUNT>1200</acl:LOCKEDAMOUNT>`,
			},
		},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			xmlRequest := NewCreateLockedAmount(tc.param)
			for _, expectedStr := range tc.expect {
				assert.Contains(t, xmlRequest, expectedStr)
			}
		})
	}
}

func TestParseCreateLockedAmountSOAP_InvalidResponseType(t *testing.T) {
	xmlData := `<?xml version="1.0" encoding="UTF-8"?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
  <S:Body>
    <SomeOtherResponse/>
  </S:Body>
</S:Envelope>`

	result, err := ParseCreateLockedAmountSOAP(xmlData)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.False(t, result.Success)
	assert.Equal(t, []string{"Invalid response type"}, result.Messages)
}

func TestParseCreateLockedAmountSOAP_MissingStatus(t *testing.T) {
	xmlData := `<?xml version="1.0" encoding="UTF-8"?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
  <S:Body>
    <ns:CreateAccountLockResponse xmlns:ns="http://temenos.com/CBESUPERAPP">
      <ACLOCKEDEVENTSType id="ACL123"/>
    </ns:CreateAccountLockResponse>
  </S:Body>
</S:Envelope>`

	result, err := ParseCreateLockedAmountSOAP(xmlData)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.False(t, result.Success)
	assert.Equal(t, []string{"Missing Status"}, result.Messages)
}

func TestParseCreateLockedAmountSOAP_T24Error(t *testing.T) {
	xmlData := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns60:CreateAccountLockResponse xmlns:ns60="http://temenos.com/CBESUPERAPP">
            <Status>
                <transactionId>ACLK21343XZ0YC</transactionId>
                <messageId></messageId>
                <successIndicator>T24Error</successIndicator>
                <application>AC.LOCKED.EVENTS</application>
                <messages>LOCKED.AMOUNT:1:1=INVALID MINUS</messages>
                <messages>LOCKED.AMOUNT:1:1=INPUT MISSING</messages>
            </Status>
        </ns60:CreateAccountLockResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseCreateLockedAmountSOAP(xmlData)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.False(t, result.Success)
	assert.Nil(t, result.Detail)
	assert.Equal(t, []string{
		"LOCKED.AMOUNT:1:1=INVALID MINUS",
		"LOCKED.AMOUNT:1:1=INPUT MISSING",
	}, result.Messages)
}

func TestParseCreateLockedAmountSOAP_Success(t *testing.T) {
	xmlData := `<?xml version="1.0" encoding="UTF-8"?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
  <S:Body>
    <ns:CreateAccountLockResponse xmlns:ns="http://temenos.com/CBESUPERAPP" xmlns:acl="http://temenos.com/ACLOCKEDEVENTS">
      <Status>
        <successIndicator>Success</successIndicator>
        <transactionId>ACL123</transactionId>
      </Status>
      <ACLOCKEDEVENTSType id="ACL123">
        <acl:ACCOUNTNUMBER>1000517052152</acl:ACCOUNTNUMBER>
        <acl:DESCRIPTION>54</acl:DESCRIPTION>
        <acl:FROMDATE>20260814</acl:FROMDATE>
        <acl:TODATE>20260913</acl:TODATE>
        <acl:LOCKEDAMOUNT>363.41</acl:LOCKEDAMOUNT>
      </ACLOCKEDEVENTSType>
    </ns:CreateAccountLockResponse>
  </S:Body>
</S:Envelope>`

	result, err := ParseCreateLockedAmountSOAP(xmlData)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.True(t, result.Success)
	require.NotNil(t, result.Detail)
	assert.Equal(t, "ACL123", result.Detail.TransactionID)
	assert.Equal(t, "1000517052152", result.Detail.AccountNumber)
	assert.Equal(t, "363.41", result.Detail.LockedAmount)
}
