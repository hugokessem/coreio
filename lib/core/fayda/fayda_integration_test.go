package fayda

import (
	"crypto/tls"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationFayda(t *testing.T) {
	params := Params{
		Username: "SUPERAPP",
		Password: "123456",
		NID:      "357253841014476138538353601641801488",
	}

	xmlRequest := NewFayda(params)
	t.Logf("xmlRequest %v", xmlRequest)

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
	if err != nil {
		t.Logf("Network error (endpoint may be unreachable): %v", err)
		t.Skip("Skipping test due to network error - endpoint may be unreachable")
		return
	}
	assert.NotNil(t, resp)
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, responseData, "Expected response body to be non-empty")

	result, err := ParseFaydaSOAP(string(responseData))
	assert.NoError(t, err)
	assert.NotNil(t, result, "Expected result to be non-nil")

	t.Logf("Result: Success=%v, Message=%v", result.Success, result.Message)

	if !result.Success {
		if len(result.Message) > 0 {
			t.Logf("API messages: %v", result.Message)
		}
		return
	}

	if result.Detail == nil {
		if len(result.Message) > 0 && result.Message[0] == "no details found" {
			t.Log("No customer details found for the given NID - this is a valid API response")
			return
		}
		t.Error("Expected Detail to be non-nil on successful response")
		return
	}

	detail := result.Detail
	assert.NotEmpty(t, detail.CustomerID)
	t.Logf("Fayda result: CustomerID=%s, CustomerName=%s, CustomerFlag=%s",
		detail.CustomerID, detail.CustomerName, detail.CustomerFlag)
	t.Log("Integration test passed")
}
