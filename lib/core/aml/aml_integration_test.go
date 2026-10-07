package aml

import (
	"crypto/tls"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const amlLookupEndpoint = "https://devapisuperapp.cbe.com.et/superapp/parser/proxy/CBESUPERAPP/services?target=http%3A%2F%2F10.1.15.195%3A8080&wsdl=null"

func TestIntegrationAMLLookup(t *testing.T) {
	params := Params{
		Username:      "SUPERAPP",
		Password:      "123456",
		AccountNumber: "1000259876854",
	}

	xmlRequest := NewAMLLookup(params)
	t.Logf("xmlRequest: %s", xmlRequest)
	require.NotEmpty(t, xmlRequest)

	req, err := http.NewRequest(http.MethodPost, amlLookupEndpoint, strings.NewReader(xmlRequest))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("SOAPAction", `"http://temenos.com/CBESUPERAPP/AccountLookup"`)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
		},
		Timeout: 60 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Logf("Network error (endpoint may be unreachable): %v", err)
		t.Skip("Skipping test due to network error - endpoint may be unreachable")
		return
	}
	require.NotNil(t, resp)
	defer resp.Body.Close()

	t.Logf("Response status: %s", resp.Status)

	responseData, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NotEmpty(t, responseData)
	t.Logf("Raw response: %s", string(responseData))

	result, err := ParseAMLLookupSOAP(string(responseData))
	require.NoError(t, err)
	require.NotNil(t, result)

	t.Logf("Result Success: %v", result.Success)
	if len(result.Messages) > 0 {
		t.Logf("Messages: %v", result.Messages)
	}

	if result.Success {
		require.NotNil(t, result.Detail)
		assert.Equal(t, "1000259876854", result.Detail.AccountNumber)
		assert.Equal(t, "AZEMERAW TEFERA TIBEBU", result.Detail.AccountName)
		assert.Equal(t, "ETB", result.Detail.Currency)
		assert.Equal(t, "1036095547", result.Detail.CustomerNumber)
		t.Logf("AML lookup detail: %+v", result.Detail)
		return
	}

	require.NotEmpty(t, result.Messages, "failed AML lookup should return messages")
}
