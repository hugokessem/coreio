package ministatementbydaterange

import (
	"crypto/tls"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationMiniStatementByDate(t *testing.T) {
	params := Params{
		Username:      "SUPERAPP",
		Password:      "123456",
		AccountNumber: "1000184349713",
		From:          "20200101",
		To:            "20200105",
	}

	xmlRequest := NewMiniStatementByDateRange(params)
	endpoint := "https://devapisuperapp.cbe.com.et/superapp/parser/proxy/CBESUPERAPP/services?target=http://10.1.15.195%3A8080&wsdl=null"

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(xmlRequest))
	assert.NoError(t, err)
	req.Header.Add("Content-Type", "text/xml; charset=utf-8")

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
		},
	}

	resp, err := client.Do(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, responseData, "Expected response body to be non-empty")

	result, err := ParseMiniStatementByDateRangeSOAP(string(responseData))
	assert.NoError(t, err)
	assert.NotNil(t, result, "Expected result to be non-nil")

	// Check that the lookup succeeded
	assert.True(t, result.Success)
	assert.NotNil(t, result.Detail)
	t.Log(result.Detail.Group)

	assert.Equal(t, "1000184349713", result.Detail.AccountNumber)
	assert.Equal(t, "1026902114", result.Detail.CustomerNumber)
	assert.Equal(t, "ETB", result.Detail.Currency)
	if len(result.Detail.Group.Details) > 0 {
		assert.Equal(t, `TT20001YXSWL\KH1`, result.Detail.Group.Details[0].TransactionReference)
		assert.Equal(t, "-14,000.00", result.Detail.Group.Details[0].Amount)
	} else {
		t.Error("Expected Detail to be non-nil")
	}
}
