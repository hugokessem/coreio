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
		FTNumber: "FT21343FGHD6",
	}

	xmlRequest := NewFundTransferCheck(params)
	t.Logf("xmlRequest: %s", xmlRequest)
	endpoint := "https://devapisuperapp.cbe.com.et/superapp/parser/proxy/CBESUPERAPP/services?target=http://10.1.15.195%3A8080&wsdl=null"

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(xmlRequest))

	assert.NoError(t, err)

	req.Header.Set("Content-Type", "text/xml; charset=utf-8")

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
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
	detail := result.Detail
	t.Logf("Debit Amount: %s", detail.DebitAmount)
	t.Logf("Debit Amount With Currency: %s", detail.DebitAmountWithCurrency)
	t.Logf("Credit Amount With Currency: %s", detail.CreditAmountWithCurrency)
	t.Logf("Total Commision With Comission: %s", detail.TotalCommisionWithComission)
	t.Logf("Desaster Recovery Fund: %s", detail.DisasterReservedFund)
	t.Logf("Processing Date: %s", detail.ProcessingDate)
	t.Logf("Service Code: %s", detail.ServiceCode)

	// Check that the lookup succeeded
	if result.Status {
		assert.NotNil(t, result.Detail)
		if result.Detail != nil {
			t.Log("-----------------------------")
			t.Logf("ChargeCommissionDisplay: %s", result.Detail.ChargeCommissionDisplay)
			t.Log("-----------------------------")
			t.Logf("Transaction Type: %s", result.Detail.TransactionType)
			t.Logf("Debit Account: %s", result.Detail.DebitAccountNumber)
			t.Logf("Credit Account: %s", result.Detail.CreditAccountNumber)
			t.Logf("Debit Amount: %s", result.Detail.DebitAmount)
			t.Logf("Credit Amount: %s", result.Detail.CreditAmount)
			t.Logf("Processing Date: %s", result.Detail.ProcessingDate)
			t.Logf("Amount Debited: %s", result.Detail.DebitAmountWithCurrency)
			t.Logf("Amount Credited: %s", result.Detail.CreditAmountWithCurrency)
			t.Logf("Total Charge Amount: %s", result.Detail.TotalChargeAmount)
		}
	} else {
		t.Logf("Fund transfer check failed with message: %s", result.Messages)
	}
}
