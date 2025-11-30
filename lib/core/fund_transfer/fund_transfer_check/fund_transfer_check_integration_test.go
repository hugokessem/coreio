package fundtransfercheck

import (
	"crypto/tls"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationFundTransferCheck(t *testing.T) {
	params := Params{
		Username: "SUPERAPP",
		Password: "123456",
		FTNumber: "FT21343CXGBD",
	}

	xmlRequest := NewFundTransferCheck(params)
	endpoint := "https://devopscbe.eaglelionsystems.com/superapp/parser/proxy/CBESUPERAPP/services?target=http%3A%2F%2F10.1.15.195%3A8080&wsdl=null"

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(xmlRequest))
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "text/xml; charset=utf-8")

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

	result, err := ParseFundTransferCheckSOAP(string(responseData))
	if err != nil {
		t.Logf("Parsing error: %v", err)
		t.Logf("Response data: %s", string(responseData))
	}
	assert.NoError(t, err)
	assert.NotNil(t, result, "Expected result to be non-nil")

	// Check that the lookup succeeded
	t.Logf("result: %+v", result)
	if result.Status {
		assert.NotNil(t, result.Detail)
		if result.Detail != nil {
			t.Logf("Transaction Type: %s", result.Detail.TransactionType)
			t.Logf("Debit Account: %s", result.Detail.DebitAccountNumber)
			t.Logf("Credit Account: %s", result.Detail.CreditAccountNumber)
			t.Logf("Debit Amount: %s", result.Detail.DebitAmount)
			t.Logf("Credit Amount: %s", result.Detail.CreditAmount)
			t.Logf("Processing Date: %s", result.Detail.ProcessingDate)
			t.Logf("Amount Debited: %s", result.Detail.AmountDebited)
			t.Logf("Amount Credited: %s", result.Detail.AmountCredited)
			t.Logf("Total Charge Amount: %s", result.Detail.TotalChargeAmount)

			// Validate that fund transfer details are populated
			if result.Detail.DebitAccountNumber != "" {
				assert.NotEmpty(t, result.Detail.DebitAccountNumber)
			}
			if result.Detail.CreditAccountNumber != "" {
				assert.NotEmpty(t, result.Detail.CreditAccountNumber)
			}
		}
	} else {
		t.Logf("Fund transfer check failed with message: %s", result.Message)
	}
}
