package aml

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewAMLLookup(t *testing.T) {
	xmlRequest := NewAMLLookup(Params{
		Username:      "SUPERAPP",
		Password:      "123456",
		AccountNumber: "1000259876854",
	})

	assert.Contains(t, xmlRequest, "<cbes:AccountNameLookup>")
	assert.Contains(t, xmlRequest, "<password>123456</password>")
	assert.Contains(t, xmlRequest, "<userName>SUPERAPP</userName>")
	assert.Contains(t, xmlRequest, "<columnName>ID</columnName>")
	assert.Contains(t, xmlRequest, "<criteriaValue>1000259876854</criteriaValue>")
	assert.Contains(t, xmlRequest, "<operand>EQ</operand>")
}

func TestParseAMLLookupSOAP_Success(t *testing.T) {
	xmlData := `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
  <S:Body>
    <ns61:AccountNameLookupResponse xmlns:ns56="http://temenos.com/NAMELOOKUPSUPERAPP" xmlns:ns61="http://temenos.com/CBESUPERAPP">
      <Status>
        <successIndicator>Success</successIndicator>
      </Status>
      <NAMELOOKUPSUPERAPPType>
        <ns56:gNAMELOOKUPSUPERAPPDetailType>
          <ns56:mNAMELOOKUPSUPERAPPDetailType>
            <ns56:AccountNumber>1000259876854</ns56:AccountNumber>
            <ns56:AccountName>AZEMERAW TEFERA TIBEBU</ns56:AccountName>
            <ns56:Currency>ETB</ns56:Currency>
            <ns56:RestrictionType></ns56:RestrictionType>
            <ns56:InactiveFlag></ns56:InactiveFlag>
            <ns56:CIF>1036095547</ns56:CIF>
            <ns56:Industry>1499</ns56:Industry>
            <ns56:CustomerSegment>MASS</ns56:CustomerSegment>
            <ns56:SubSegment>MASS</ns56:SubSegment>
            <ns56:CustomerGroup>RETAIL</ns56:CustomerGroup>
            <ns56:Category>6501</ns56:Category>
            <ns56:Sector>1000</ns56:Sector>
            <ns56:TINNo></ns56:TINNo>
          </ns56:mNAMELOOKUPSUPERAPPDetailType>
        </ns56:gNAMELOOKUPSUPERAPPDetailType>
      </NAMELOOKUPSUPERAPPType>
    </ns61:AccountNameLookupResponse>
  </S:Body>
</S:Envelope>`

	result, err := ParseAMLLookupSOAP(xmlData)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.True(t, result.Success)
	require.NotNil(t, result.Detail)
	assert.Equal(t, "1000259876854", result.Detail.AccountNumber)
	assert.Equal(t, "AZEMERAW TEFERA TIBEBU", result.Detail.AccountName)
	assert.Equal(t, "ETB", result.Detail.Currency)
	assert.Equal(t, "1036095547", result.Detail.CustomerNumber)
	assert.Equal(t, "1499", result.Detail.Industry)
	assert.Equal(t, "MASS", result.Detail.CustomerSegment)
	assert.Equal(t, "MASS", result.Detail.SubSegment)
	assert.Equal(t, "RETAIL", result.Detail.CustomerGroup)
	assert.Equal(t, "6501", result.Detail.Category)
	assert.Equal(t, "1000", result.Detail.Sector)
}

func TestParseAMLLookupSOAP_Failure(t *testing.T) {
	xmlData := `<?xml version="1.0" encoding="UTF-8"?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
  <S:Body>
    <ns61:AccountNameLookupResponse xmlns:ns61="http://temenos.com/CBESUPERAPP">
      <Status>
        <successIndicator>T24Error</successIndicator>
        <messages>No records were found that matched the selection criteria</messages>
      </Status>
    </ns61:AccountNameLookupResponse>
  </S:Body>
</S:Envelope>`

	result, err := ParseAMLLookupSOAP(xmlData)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.False(t, result.Success)
	assert.Nil(t, result.Detail)
	assert.Equal(t, []string{"No records were found that matched the selection criteria"}, result.Messages)
}

func TestParseAMLLookupSOAP_InvalidResponseType(t *testing.T) {
	xmlData := `<?xml version="1.0" encoding="UTF-8"?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
  <S:Body>
    <SomeOtherResponse/>
  </S:Body>
</S:Envelope>`

	result, err := ParseAMLLookupSOAP(xmlData)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.False(t, result.Success)
	assert.Equal(t, []string{"Invalid response type"}, result.Messages)
}
