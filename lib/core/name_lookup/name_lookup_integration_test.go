package namelookup

import (
	"crypto/tls"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationAccountLookup(t *testing.T) {
	params := Params{
		Username:      "SUPERAPP",
		Password:      "123456",
		AccountNumber: "1000517052152",
	}

	xmlRequest := NewNameLookup(params)
	t.Logf("xmlRequest %v", xmlRequest)
	endpoint := "https://devapisuperapp.cbe.com.et/superapp/parser/proxy/CBESUPERAPP/services?target=http://10.1.15.195%3A8080&wsdl=null"

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(xmlRequest))
	assert.NoError(t, err)

	req.Header.Add("Content-Type", "text/xml; charset=utf-8")

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: false,
			},
		},
	}

	resp, err := client.Do(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, responseData, "Expected response body to be non-empty")

	result, err := ParseNameLookupSOAP(string(responseData))
	assert.NoError(t, err)
	assert.NotNil(t, result, "Expected result to be non-nil")

	// Check that the lookup succeeded
	// if result == nil {
	// 	t.Fatal("Expected result to be non-nil")
	// }

	assert.True(t, result.Success)
	assert.NotNil(t, result.Detail)
	t.Logf("Details: %+v", result.Detail)
	// t.Logf("Currency: %+v", result.Detail.Currency)
	// t.Logf("Customer Number: %+v", result.Detail.CustomerNumber)
	// t.Log("resultDetail", result)

	// if result.Detail != nil {
	// 	assert.Equal(t, "1000517052152", result.Detail.AccountNumber)
	// 	assert.Equal(t, "ETB", result.Detail.Currency)
	// } else {
	// 	t.Error("Expected Detail to be non-nil")
	// }
}
