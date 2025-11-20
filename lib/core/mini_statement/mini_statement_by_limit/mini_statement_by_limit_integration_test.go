package ministatementbylimit

import (
	"crypto/tls"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationMiniStatement(t *testing.T) {
	params := Params{
		Username:            "SUPERAPP",
		Password:            "123456",
		AccountNumber:       "1000030677308",
		NumberOfTransaction: "3",
	}

	xmlRequest := NewMiniStatement(params)
	endpoint := "https://devopscbe.eaglelionsystems.com/superapp/parser/proxy/CBESUPERAPP/services?target=http%3A%2F%2F10.1.15.195%3A8080&wsdl=null"

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(xmlRequest))
	assert.NoError(t, err)
	req.Header.Add("Content-Type", "text/xml; charset=utf-8")

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	resp, err := client.Do(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, responseData, "Expected response body to be non-empty")

	result, err := ParseMiniStatementSOAP(string(responseData))
	assert.NoError(t, err)
	assert.NotNil(t, result, "Expected result to be non-nil")

	// Check that the lookup succeeded
	assert.True(t, result.Success)
	assert.NotNil(t, result.Details)

	if len(result.Details) > 0 {
		assert.Equal(t, "ABIY HAILEYESUS MENGISTU", result.Details[0].OtherPartyAccount)
		assert.Equal(t, "FT2134349MRK", result.Details[0].TransactionReference)
		assert.Equal(t, "ETB", result.Details[0].Currency)
		assert.Equal(t, "-2.00", result.Details[0].Amount)
	} else {
		t.Error("Expected Detail to be non-nil")
	}
}
