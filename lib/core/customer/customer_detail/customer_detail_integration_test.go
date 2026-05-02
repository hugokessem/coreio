package customerdetail

import (
	"context"
	"crypto/tls"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestIntegrationFundTransfer(t *testing.T) {
	// Test parameters matching the curl request

	params := Params{
		Username:       "SUPERAPP",
		Password:       "123456",
		CustomerNumber: "1027260729",
	}

	xmlRequest := NewCustomerDetail(params)
	t.Logf("Generated XML Request: %s", xmlRequest)
	require.NotEmpty(t, xmlRequest, "Generated XML should not be empty")

	// Use the production endpoint from curl request
	endpoint := "https://devapisuperapp.cbe.com.et/superapp/parser/proxy/CBESUPERAPP/services?target=http%3A%2F%2F10.1.15.195%3A8080&wsdl=null"

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, strings.NewReader(xmlRequest))
	require.NoError(t, err, "Failed to create HTTP request")

	// Set headers matching the curl request
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")

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
	require.NotNil(t, resp, "Response should not be nil")
	defer resp.Body.Close()

	// Accept both 200 OK and other success status codes
	if resp.StatusCode != http.StatusOK {
		t.Logf("Received status code: %d", resp.StatusCode)
		// Continue with parsing even if status is not 200
	}

	responseData, err := io.ReadAll(resp.Body)
	require.NoError(t, err, "Failed to read response body")
	require.NotEmpty(t, responseData, "Response body should not be empty")

	result, err := ParseCustomerDetailSOAP(string(responseData))
	require.NoError(t, err, "Failed to parse SOAP response")
	require.NotNil(t, result, "Parsed result should not be nil")

	t.Logf("Result: %v", result)
	t.Logf("Result Success: %v", result.Success)
	t.Logf("Customer Name: %s", result.CustomerInfos)
	t.Logf("CustomerGroup: %v", result.CustomerInfos.CustomerGroup)
	t.Logf("CustomerSegment: %v", result.CustomerInfos.CustomerSegment)
	t.Logf("CustomerSubSegment: %v", result.CustomerInfos.CustomerSubSegment)
	if len(result.Message) > 0 {
		t.Logf("Messages: %v", result.Message)
	}

}
