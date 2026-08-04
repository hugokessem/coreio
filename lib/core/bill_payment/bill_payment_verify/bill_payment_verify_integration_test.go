package billpaymentverify

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

func TestIntegrationBillPaymentVerify(t *testing.T) {
	clientReference, _ := uuid.NewV7()
	creditReference, _ := uuid.NewV7()
	debitReference, _ := uuid.NewV7()

	params := Param{
		Username:            "SUPERAPP",
		Password:            "123456",
		DebitAccountNumber:  "1000263525144",
		DebitCurrency:       "ETB",
		DebitAmount:         "30",
		DebitReference:      debitReference.String()[3:13],
		CreditReference:     creditReference.String()[1:11],
		CreditAccountNumber: "1000357597823",
		CreditCurrency:      "ETB",
		CreditAmount:        "",
		PaymentDetails:      "TELEBIRR",
		ClientReference:     clientReference.String(),
		ServiceDescription:  "test Service Description",
	}

	xmlRequest := NewBillPaymentVerify(params)
	t.Logf("Generated XML Request: %s", xmlRequest)
	require.NotEmpty(t, xmlRequest)

	endpoint := "https://devapisuperapp.cbe.com.et/superapp/parser/proxy/CBESUPERAPP/services?target=http%3A%2F%2F10.1.15.195%3A8080&wsdl=null"

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, strings.NewReader(xmlRequest))
	require.NoError(t, err, "Failed to create HTTP request")

	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("SOAPAction", `"http://temenos.com/CBESUPERAPP/FTBillPayment_Validate"`)

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

	result, err := ParseBillPaymentVerifySOAP(string(responseData))
	require.NoError(t, err, "Failed to parse SOAP response")
	require.NotNil(t, result, "Parsed result should not be nil")

	t.Logf("Result Success: %v", result.Success)
	if len(result.Messages) > 0 {
		t.Logf("Messages: %v", result.Messages)
	}

	if !result.Success {
		t.Logf("Bill payment verify was not successful. Messages: %v", result.Messages)
		if result.Detail == nil {
			t.Log("No detail returned from API")
			return
		}
	}

	if result.Detail != nil {
		detail := result.Detail
		t.Logf("FT Number: %s", detail.FTNumber)
		t.Logf("Transaction ID: %s", detail.TransactionID)
		t.Logf("Debit Amount: %s", detail.DebitAmount)
		t.Logf("Debit Amount With Currency: %s", detail.DebitAmountWithCurrency)
		t.Logf("Credit Amount With Currency: %s", detail.CreditAmountWithCurrency)
		t.Logf("Total Commision: %s", detail.TotalCommisionWithComission)
		t.Logf("Total Tax Amount: %s", detail.TotalTaxAmount)
		t.Logf("Disaster Recovery Fund: %s", detail.DisasterReservedFund)
		t.Logf("Original Paid Amount: %s", detail.OriginalPaidAmount)
		t.Logf("Processing Date: %s", detail.ProcessingDate)
		t.Logf("Debit Account Holder: %s", detail.DebitAccountHolderName)
		t.Logf("Receiver Name: %s", detail.ReceiverName)
		t.Logf("Service Description: %s", detail.ServiceDescription)
		t.Logf("Account Category: %s", detail.AccountCategory)
		assert.NotEmpty(t, detail.FTNumber)
	} else {
		t.Error("Expected Detail to be non-nil")
	}
}
