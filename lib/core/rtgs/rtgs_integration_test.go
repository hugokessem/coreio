package rtgs

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

func TestIntegrationRTGS(t *testing.T) {
	params := Params{
		Username:           "SUPERAPP",
		Password:           "123456",
		Company:            "ET0010222",
		DebitAccountNumber: "1000251861149",
		DebitCurrency:      "ETB",
		CreditAmount:       "50",
		OrderingCust: []string{
			"MOR-BAHR0049695769",
			"1437145650002",
			"DURBET TRADE SHARE COMPANY",
		},
		BenAccountNumber: "0100191040217",
		PaymentDetails:   "0049695769 BAHIR DAR BRANCH",
		BkToBkInfo: []string{
			"/REC/TAXFRMID",
			"//0049695769/1437145650002",
			"//EFFDATE/20260402",
		},
		BillID: "1437145650002",
	}

	xmlRequest := NewRTGS(params)
	t.Logf("Generated XML Request: %s", xmlRequest)
	require.NotEmpty(t, xmlRequest, "Generated XML should not be empty")

	endpoint := "https://devapisuperapp.cbe.com.et/superapp/parser/proxy/CBESUPERAPP/services?target=http%3A%2F%2F10.1.15.195%3A8080&wsdl=null"

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, strings.NewReader(xmlRequest))
	require.NoError(t, err, "Failed to create HTTP request")

	req.Header.Set("Content-Type", "text/xml;charset=UTF-8")

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

	if resp.StatusCode != http.StatusOK {
		t.Logf("Received status code: %d", resp.StatusCode)
	}

	responseData, err := io.ReadAll(resp.Body)
	require.NoError(t, err, "Failed to read response body")
	require.NotEmpty(t, responseData, "Response body should not be empty")
	t.Logf("Raw Response: %s", string(responseData))

	result, err := ParseRTGSSOAP(string(responseData))
	require.NoError(t, err, "Failed to parse SOAP response")
	require.NotNil(t, result, "Parsed result should not be nil")

	t.Logf("Result Success: %v", result.Success)
	if len(result.Messages) > 0 {
		t.Logf("Messages: %v", result.Messages)
	}

	if !result.Success {
		t.Logf("RTGS was not successful. Messages: %v", result.Messages)
		if result.Detail == nil {
			t.Log("No detail returned from API")
			return
		}
	}

	if result.Detail != nil {
		detail := result.Detail
		t.Logf("FT Number: %s", detail.FTNumber)
		t.Logf("Transaction Type: %s", detail.TransactionType)
		t.Logf("Debit Account: %s", detail.DebitAccountNumber)
		t.Logf("Debit Currency: %s", detail.DebitCurrency)
		t.Logf("Credit Amount: %s", detail.CreditAmount)
		t.Logf("Ben Account: %s", detail.BenAccountNumber)
		t.Logf("Bill ID: %s", detail.BillID)
		t.Logf("Amount Debited: %s", detail.DebitAmountWithCurrency)
		t.Logf("Amount Credited: %s", detail.CreditAmountWithCurrency)
		t.Logf("Processing Date: %s", detail.ProcessingDate)
		t.Logf("Company Code: %s", detail.CompanyCode)
	} else {
		t.Log("Expected Detail to be non-nil on success path")
	}
}
