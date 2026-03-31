package statuscheck

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

func TestIntegrationStatusCheck(t *testing.T) {
	params := Params{
		Username:      "SUPERAPP",
		Password:      "123456",
		TransactionID: "1238582457889",
	}

	xmlRequest := NewStatusCheck(params)
	t.Logf("Generated XML Request: %s", xmlRequest)
	require.NotEmpty(t, xmlRequest, "Generated XML should not be empty")

	endpoint := "https://devapisuperapp.cbe.com.et/superapp/parser/proxy/CBESUPERAPP/services?target=http%3A%2F%2F10.1.15.195%3A8080&wsdl=null"

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, strings.NewReader(xmlRequest))
	require.NoError(t, err, "Failed to create HTTP request")

	// Set headers matching the curl request
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("SOAPAction", `"http://temenos.com/CBESUPERAPP/AccountTransfer"`)

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

	result, err := ParseStatusCheckSOAP(string(responseData))
	require.NoError(t, err, "Failed to parse SOAP response")
	require.NotNil(t, result, "Parsed result should not be nil")

	t.Logf("Result: %v", result)

	t.Logf("Result Success: %v", result.Success)
	if len(result.Messages) > 0 {
		t.Logf("Messages: %v", result.Messages)
	}

	if !result.Success {
		t.Logf("Status Check was not successful. Messages: %v", result.Messages)
		if result.Detail == nil {
			t.Log("No detail returned from API")
			return
		}
	}

	if result.Detail != nil {
		detail := result.Detail
		assert.NotEmpty(t, detail.ServiceCode, "ServiceCode should not be empty")
		assert.NotEmpty(t, detail.DebitAccount, "DebitAccount should not be empty")
		assert.NotEmpty(t, detail.CreditAccount, "CreditAccount should not be empty")
		assert.NotEmpty(t, detail.DebitCurrency, "DebitCurrency should not be empty")
		assert.NotEmpty(t, detail.DebitAmount, "DebitAmount should not be empty")
		assert.NotEmpty(t, detail.Channel, "Channel should not be empty")
		assert.NotEmpty(t, detail.FTReference, "FTReference should not be empty")

		t.Logf("Service Code: %s", detail.ServiceCode)
		t.Logf("DebitAccount: %s", detail.DebitAccount)
		t.Logf("CreditAccount: %s", detail.CreditAccount)
		t.Logf("Debit Amount: %s", detail.DebitAmount)
		t.Logf("DebitCurrency: %s", detail.DebitCurrency)
		t.Logf("Channel: %s", detail.Channel)
		t.Logf("FTReference: %s", detail.FTReference)
	} else {
		t.Error("Expected Detail to be non-nil")
	}
}
