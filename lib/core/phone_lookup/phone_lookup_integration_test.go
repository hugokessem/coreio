package phonelookup

import (
	"crypto/tls"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationPhoneLookup(t *testing.T) {
	params := Params{
		Username:    "SUPERAPP",
		Password:    "123456",
		PhoneNumber: "Y911706628",
	}

	xmlRequest := NewPhoneLookup(params)
	endpoint := "https://devapisuperapp.cbe.com.et/superapp/parser/proxy/CBESUPERAPP/services?target=http://10.1.15.195%3A8080&wsdl=null"

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

	result, err := ParsePhoneLookupSOAP(string(responseData))
	assert.NoError(t, err)
	assert.NotNil(t, result, "Expected result to be non-nil")

	// Handle the case where no details are found (valid API response)
	if !result.Success && result.Message == "no details found" {
		t.Log("No customer details found for the given phone number - this is a valid API response")
		return
	}

	// Check that the lookup succeeded
	assert.True(t, result.Success)
	assert.NotNil(t, result.Detail)
	t.Logf("result: %v", result.Detail)

	if result.Detail != nil {
		assert.NotEmpty(t, result.Detail.CustomerID)
		assert.NotEmpty(t, result.Detail.PhoneNumber)
		t.Logf("Phone lookup result: CustomerID=%s, PhoneNumber=%s, Email=%s",
			result.Detail.CustomerID, result.Detail.PhoneNumber, result.Detail.Email)
		t.Log("Integration test passed")
	} else {
		t.Error("Expected Detail to be non-nil")
	}
}
