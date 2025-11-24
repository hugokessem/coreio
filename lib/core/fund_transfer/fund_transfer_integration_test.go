package fundtransfer

import (
	"crypto/tls"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationFundTransfer(t *testing.T) {
	params := Params{
		Username:            "SUPERAPP",
		Password:            "123456",
		DebitAccountNumber:  "1000000006924",
		DebitCurrency:       "ETB",
		CreditAccountNumber: "1000357597823",
		CreditCurrency:      "ETB",
		DebitAmount:         "124.00",
		TransationID:        "TXN123456789",
		DebitReference:      "Payment",
		CreditReference:     "Received payment",
		PaymentDetail:       "Fund transfer",
	}

	xmlRequest := NewFundTransfer(params)
	endpoint := "https://devopscbe.eaglelionsystems.com/superapp/parser/proxy/CBESUPERAPP/services?target=http%3A%2F%2F10.1.15.195%3A8080&wsdl=null"

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(xmlRequest))
	assert.NoError(t, err)

	req.Header.Add("Content-Type", "text/xml; charset=utf-8")

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

	result, err := ParseFundTransferSOAP(string(responseData))
	t.Log(result.Messages)
	assert.NoError(t, err)
	assert.NotNil(t, result, "Expected result to be non-nil")

	// Check that the lookup succeeded
	assert.True(t, result.Success)
	assert.NotNil(t, result.Detail)
	t.Log("result", result.Detail)

	if result.Detail != nil {
		assert.Equal(t, "1000000006924", result.Detail.DebitAccountNumber)
		assert.Equal(t, "1000357597823", result.Detail.CreditAccountNumber)
		assert.Equal(t, "ETB124.00", result.Detail.CreditAmountWithCurrency)
	} else {
		t.Error("Expected Detail to be non-nil")
	}
}
