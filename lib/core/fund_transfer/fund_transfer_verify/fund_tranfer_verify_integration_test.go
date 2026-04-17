package fundtrasferverify

import (
	"context"
	"crypto/tls"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntegrationFundTransferVerify(t *testing.T) {
	// Test parameters matching the curl request
	clientReference, _ := uuid.NewV7()
	params := Params{
		Username:           "SUPERAPP",
		Password:           "123456",
		DebitAccountNumber: "1000446113608",
		DebitCurrency:      "USD",
		DebitAmount:        "10000",
		// DebitAccountNumber:  "1000446113608",
		// DebitCurrency:       "USD",
		// CreditAmount:        "50",
		CreditAccountNumber: "1000357597823",
		CreditCurrency:      "ETB",
		DebitReference:      "DEBIT NARRATIVE",
		CreditReference:     "CREDIT NARRATIVE",
		PaymentDetails:      "TEST PAYMENT",
		ClientReference:     clientReference.String(),
		ServiceCode:         "CBE",
		CustomerSegment:     "MASS",
		ChannelType:         "USSD",
	}

	xmlRequest := NewFundTransferVerify(params)
	t.Logf("Generated XML: %s", xmlRequest)
	require.NotEmpty(t, xmlRequest, "Generated XML should not be empty")

	// Use the production endpoint from curl request
	endpoint := "https://devapisuperapp.cbe.com.et/superapp/parser/proxy/CBESUPERAPP/services?target=http%3A%2F%2F10.1.15.195%3A8080&wsdl=null"

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, strings.NewReader(xmlRequest))
	require.NoError(t, err, "Failed to create HTTP request")

	// Set headers matching the curl request
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("SOAPAction", `"http://temenos.com/CBESUPERAPP/AccountTransfer_Validate"`)

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
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	require.NoError(t, err, "Failed to read response body")
	require.NotEmpty(t, responseData, "Expected response body to be non-empty")

	result, err := ParseFundTransferVerifySOAP(string(responseData))
	require.NoError(t, err, "Failed to parse SOAP response")
	require.NotNil(t, result, "Expected result to be non-nil")

	// Check that "the lookup succeeded
	t.Logf("result: %+v", result)
	t.Log(result.Detail)

	assert.True(t, result.Success)
	assert.NotNil(t, result.Detail)

	if result.Detail != nil {
		assert.Equal(t, "USSD", result.Detail.TransactionChannel)
		assert.Equal(t, "50.00", result.Detail.DebitAmount)
		assert.Equal(t, "ETB", result.Detail.DebitCurrency)
	} else {
		t.Error("Expected Detail to be non-nil")
	}
}
