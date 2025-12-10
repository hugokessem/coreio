package accountlist

import (
	"crypto/tls"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationAccountList(t *testing.T) {
	params := Params{
		Username:      "SUPERAPP",
		Password:      "123456",
		ColumnName:    "CUS.ID",
		CriteriaValue: "1000000015",
	}

	xmlRequest := NewAccountList(params)
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

	result, err := ParseAccountListSOAP(string(responseData))
	assert.NoError(t, err)
	assert.NotNil(t, result, "Expected result to be non-nil")

	// Check that the request succeeded
	assert.True(t, result.Success)
	assert.NotNil(t, result.Details)
	t.Logf("result: %v", result.Details)

	if len(result.Details) > 0 {
		// Validate at least one account detail
		detail := result.Details[0]
		assert.NotEmpty(t, detail.AccountNumber)
		assert.NotEmpty(t, detail.CustomerName)
		assert.NotEmpty(t, detail.Currency)
		assert.NotEmpty(t, detail.CustomerID)
		t.Logf("First account: AccountNumber=%s, CustomerName=%s, Currency=%s",
			detail.AccountNumber, detail.CustomerName, detail.Currency)
		t.Log("Integration test passed")
	} else {
		t.Log("No account details found, but request was successful")
	}
}
