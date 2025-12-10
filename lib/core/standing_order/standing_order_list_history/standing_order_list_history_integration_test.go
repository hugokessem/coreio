package standingorderlisthistory

import (
	"crypto/tls"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationStandingOrderListHistory(t *testing.T) {
	params := Params{
		Username:      "SUPERAPP",
		Password:      "123456",
		AccountNumber: "1000373456776",
	}

	xmlRequest := NewListStandingOrderHistory(params)
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

	result, err := ParseStandingOrderListHistorySOAP(string(responseData))
	assert.NoError(t, err)
	assert.NotNil(t, result, "Expected result to be non-nil")

	// Check that the request succeeded
	assert.True(t, result.Success)
	t.Logf("result: %v", result.Details)

	if len(result.Details) > 0 {
		// Validate at least one standing order history detail
		detail := result.Details[0]
		assert.NotEmpty(t, detail.StandingOrderId)
		assert.NotEmpty(t, detail.Currency)
		assert.NotEmpty(t, detail.DebitAccountNumber)
		t.Logf("First standing order history: StandingOrderId=%s, Currency=%s, Amount=%s, Frequency=%s",
			detail.StandingOrderId, detail.Currency, detail.Amount, detail.Frequency)
		t.Log("Integration test passed")
	} else {
		t.Log("No standing order history found, but request was successful")
	}
}
