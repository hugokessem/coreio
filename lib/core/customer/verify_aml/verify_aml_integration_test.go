package verifyaml

import (
	"context"
	"crypto/tls"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntegrationVerifyAML(t *testing.T) {
	params := Param{
		Password:       "123456",
		UserName:       "SUPERAPP",
		CustomerNumber: "1072666796",
	}

	xmlRequest := NewVerifyAML(params)
	t.Logf("Generated XML Request: %s", xmlRequest)
	require.NotEmpty(t, xmlRequest)

	endpoint := "https://devapisuperapp.cbe.com.et/superapp/parser/proxy/CBESUPERAPP/services?target=http%3A%2F%2F10.1.15.195%3A8080&wsdl=null"

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, strings.NewReader(xmlRequest))
	require.NoError(t, err, "Failed to create HTTP request")

	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("SOAPAction", `"http://temenos.com/CBESUPERAPP/AMLStatusCheckSuperApp"`)

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

	if resp.StatusCode != http.StatusOK {
		t.Logf("Received status code: %d", resp.StatusCode)
	}

	responseData, err := io.ReadAll(resp.Body)
	require.NoError(t, err, "Failed to read response body")
	require.NotEmpty(t, responseData, "Response body should not be empty")
	t.Logf("Raw response: %s", string(responseData))

	result, err := ParseVerifyAMLSOAP(string(responseData))
	require.NoError(t, err, "Failed to parse SOAP response")
	require.NotNil(t, result, "Parsed result should not be nil")

	t.Logf("Result Success: %v", result.Success)
	t.Logf("FCM Status: %s", result.FCMStatus)
	if len(result.Messages) > 0 {
		t.Logf("Messages: %v", result.Messages)
	}

	if result.Success {
		assert.NotEmpty(t, result.FCMStatus)
	} else {
		t.Logf("AML verify was not successful. Messages: %v", result.Messages)
	}
}
