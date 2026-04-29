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
		Username:            "SUPERAPP",
		Password:            "123456",
		DebitAccountNumber:  "1000263525144",
		DebitCurrency:       "ETB",
		CreditAccountNumber: "1000298095649",
		CreditCurrency:      "ETB",
		DebitAmount:         "10",
		// DebitAccountNumber: "1000446113608",
		// DebitCurrency:      "USD",
		// DebitAmount:        "1000",

		// CreditAccountNumber: "1000446116286",
		// CreditCurrency:      "GBP",
		// CreditAmount:        "50",

		// DebitAccountNumber:  "1000446113608",
		// DebitCurrency:       "USD",

		// CreditAccountNumber: "1000446115875",
		// CreditCurrency:      "USD",
		// CreditAmount:        "50",

		// CreditAccountNumber: "1000357597823",
		// CreditCurrency:      "ETB",
		DebitReference:  "DEBIT NARRATIVE",
		CreditReference: "CREDIT NARRATIVE",
		PaymentDetails:  "TEST PAYMENT",
		ClientReference: clientReference.String(),
		ServiceCode:     "CBE",
		CustomerSegment: "MASS",
		ChannelType:     "APP",
		UserID:          "SA1036559081",
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
	t.Logf("DisasterReservedFund: %+v", result.Detail.DisasterReservedFund)
	t.Logf("OriginalPaidAmount: %+v", result.Detail.OriginalPaidAmount)
	t.Logf("TotalCommisionWithComission: %+v", result.Detail.TotalCommisionWithComission)

	assert.True(t, result.Success)
	assert.NotNil(t, result.Detail)

	detail := result.Detail

	t.Logf("FT Number: %s", detail.FTNumber)
	t.Logf("Transaction ID: %s", detail.TransactionID)
	t.Logf("Debit Amount: %s", detail.DebitAmount)
	t.Logf("Debit Amount With Currency: %s", detail.DebitAmountWithCurrency)
	t.Logf("Credit Amount With Currency: %s", detail.CreditAmountWithCurrency)
	t.Logf("Total Commision With Comission: %s", detail.TotalCommisionWithComission)
	t.Logf("Total Tax Amount: %s", detail.TotalTaxAmount)
	t.Logf("Desaster Recovery Fund: %s", detail.DisasterReservedFund)
	t.Logf("Processing Date: %s", detail.ProcessingDate)
	t.Logf("Debit Account Holder: %s", detail.DebitAccountHolderName)
	t.Logf("Receiver Name: %s", detail.ReceiverName)
	t.Logf("Service Code: %s", detail.ServiceCode)
	if result.Detail != nil {
		assert.Equal(t, "USSD", result.Detail.TransactionChannel)
		assert.Equal(t, "50.00", result.Detail.DebitAmount)
		assert.Equal(t, "ETB", result.Detail.DebitCurrency)
	} else {
		t.Error("Expected Detail to be non-nil")
	}
}
