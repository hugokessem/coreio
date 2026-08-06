package verifyaml

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewVerifyAML(t *testing.T) {
	tests := []struct {
		name   string
		param  Param
		expect []string
	}{
		{
			name: "generates AML status check request XML",
			param: Param{
				Password:       "123456",
				UserName:       "SUPERAPP",
				CustomerNumber: "1000123456789",
			},
			expect: []string{
				"<soapenv:Envelope",
				"<soapenv:Header/>",
				"<soapenv:Body>",
				"<cbes:AMLStatusCheckSuperApp>",
				"<WebRequestCommon>",
				"<password>123456</password>",
				"<userName>SUPERAPP</userName>",
				"<AMLSTATUSSUPERAPPType>",
				"<columnName>ID</columnName>",
				"<criteriaValue>1000123456789</criteriaValue>",
				"<operand>CT</operand>",
			},
		},
		{
			name: "generates XML with different credentials",
			param: Param{
				Password:       "TESTPASS",
				UserName:       "TESTUSER",
				CustomerNumber: "2000987654321",
			},
			expect: []string{
				"<password>TESTPASS</password>",
				"<userName>TESTUSER</userName>",
				"<criteriaValue>2000987654321</criteriaValue>",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			xmlRequest := NewVerifyAML(tc.param)
			assert.NotEmpty(t, xmlRequest)
			for _, expected := range tc.expect {
				assert.Contains(t, xmlRequest, expected)
			}
		})
	}
}

func TestParseVerifyAMLSOAP_Success(t *testing.T) {
	xmlData := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns60:AMLStatusCheckSuperAppResponse xmlns:ns14="http://temenos.com/AMLSTATUSSUPERAPP" xmlns:ns60="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>Success</successIndicator>
            </Status>
            <AMLSTATUSSUPERAPPType>
                <ns14:gAMLSTATUSSUPERAPPDetailType>
                    <ns14:mAMLSTATUSSUPERAPPDetailType>
                        <ns14:FCMSTATUS>APPROVED</ns14:FCMSTATUS>
                    </ns14:mAMLSTATUSSUPERAPPDetailType>
                </ns14:gAMLSTATUSSUPERAPPDetailType>
            </AMLSTATUSSUPERAPPType>
        </ns60:AMLStatusCheckSuperAppResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseVerifyAMLSOAP(xmlData)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.True(t, result.Success)
	assert.Equal(t, "APPROVED", result.FCMStatus)
	assert.Nil(t, result.Messages)
}

func TestParseVerifyAMLSOAP_FailureResponse(t *testing.T) {
	xmlData := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns60:AMLStatusCheckSuperAppResponse xmlns:ns60="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>Failure</successIndicator>
                <messages>AML check rejected</messages>
            </Status>
            <AMLSTATUSSUPERAPPType>
            </AMLSTATUSSUPERAPPType>
        </ns60:AMLStatusCheckSuperAppResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseVerifyAMLSOAP(xmlData)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.False(t, result.Success)
	assert.Empty(t, result.FCMStatus)
	assert.Equal(t, []string{"AML check rejected"}, result.Messages)
}

func TestParseVerifyAMLSOAP_FailureWithoutMessages(t *testing.T) {
	xmlData := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns60:AMLStatusCheckSuperAppResponse xmlns:ns60="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>Failure</successIndicator>
            </Status>
        </ns60:AMLStatusCheckSuperAppResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseVerifyAMLSOAP(xmlData)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.False(t, result.Success)
	assert.Equal(t, []string{"AML status check failed"}, result.Messages)
}

func TestParseVerifyAMLSOAP_NoStatus(t *testing.T) {
	xmlData := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns60:AMLStatusCheckSuperAppResponse xmlns:ns60="http://temenos.com/CBESUPERAPP">
            <AMLSTATUSSUPERAPPType>
            </AMLSTATUSSUPERAPPType>
        </ns60:AMLStatusCheckSuperAppResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseVerifyAMLSOAP(xmlData)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.False(t, result.Success)
	assert.Equal(t, []string{"Missing Status"}, result.Messages)
}

func TestParseVerifyAMLSOAP_NoAMLDetails(t *testing.T) {
	xmlData := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns60:AMLStatusCheckSuperAppResponse xmlns:ns14="http://temenos.com/AMLSTATUSSUPERAPP" xmlns:ns60="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>Success</successIndicator>
            </Status>
            <AMLSTATUSSUPERAPPType>
                <ns14:gAMLSTATUSSUPERAPPDetailType>
                </ns14:gAMLSTATUSSUPERAPPDetailType>
            </AMLSTATUSSUPERAPPType>
        </ns60:AMLStatusCheckSuperAppResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseVerifyAMLSOAP(xmlData)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.False(t, result.Success)
	assert.Equal(t, []string{"No AML status details found"}, result.Messages)
}

func TestParseVerifyAMLSOAP_SuccessWithoutDetailsUsesStatusMessages(t *testing.T) {
	xmlData := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns60:AMLStatusCheckSuperAppResponse xmlns:ns60="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>Success</successIndicator>
                <messages>No records were found that matched the selection criteria</messages>
            </Status>
        </ns60:AMLStatusCheckSuperAppResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseVerifyAMLSOAP(xmlData)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.False(t, result.Success)
	assert.Equal(t, []string{"No records were found that matched the selection criteria"}, result.Messages)
}

func TestParseVerifyAMLSOAP_InvalidResponseType(t *testing.T) {
	xmlData := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <SomeOtherResponse>
        </SomeOtherResponse>
    </S:Body>
</S:Envelope>`

	result, err := ParseVerifyAMLSOAP(xmlData)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.False(t, result.Success)
	assert.Equal(t, []string{"Invalid response type"}, result.Messages)
}

func TestParseVerifyAMLSOAP_InvalidXML(t *testing.T) {
	result, err := ParseVerifyAMLSOAP(`invalid xml content`)
	assert.Error(t, err)
	assert.Nil(t, result)
}
