package actionautorization

import (
	"context"
	"crypto/tls"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const actionAutorizationEndpoint = "https://devapisuperapp.cbe.com.et/superapp/parser/proxy/IIBONBOARDING/services?target=http%3A%2F%2F172.31.6.115%3A9095&wsdl=null"

func TestIntegrationActionAutorization(t *testing.T) {
	params := Param{
		Username:       "SUPERAPP.AUTH",
		Password:       "123456",
		CustomerNumber: "1277588454",
	}

	xmlRequest := NewActionAutorization(params)
	t.Logf("Generated XML Request: %s", xmlRequest)
	require.NotEmpty(t, xmlRequest)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, actionAutorizationEndpoint, strings.NewReader(xmlRequest))
	require.NoError(t, err, "Failed to create HTTP request")

	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("SOAPAction", `"http://temenos.com/IIBONBOARDING/DeleteCustomerIDSuperApp"`)
	if token := os.Getenv("CBE_SUPERAPP_BEARER_TOKEN"); token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

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
	require.NoError(t, err, "Failed to read response body")
	require.NotEmpty(t, responseData, "Response body should not be empty")
	t.Logf("Raw response: %s", string(responseData))

	result, err := ParseActionAutorizationSOAP(string(responseData))
	require.NoError(t, err, "Failed to parse SOAP response")
	require.NotNil(t, result, "Parsed result should not be nil")

	t.Logf("Result Success: %v", result.Success)
	if len(result.Messages) > 0 {
		t.Logf("Messages: %v", result.Messages)
	}

	if result.Success {
		require.NotNil(t, result.Detail)
		assert.NotEmpty(t, result.Detail.CustomerNumber)
		assert.Equal(t, "IDEL", result.Detail.RecordStatus)
		t.Logf("Authorized customer: %s recordStatus=%s shortName=%s", result.Detail.CustomerNumber, result.Detail.RecordStatus, result.Detail.ShortName)
		return
	}

	require.NotEmpty(t, result.Messages, "failed authorization should return messages")
	t.Logf("Action authorization was not successful. Messages: %v", result.Messages)
}
