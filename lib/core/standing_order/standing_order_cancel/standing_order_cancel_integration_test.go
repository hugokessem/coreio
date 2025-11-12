package standingordercancel

import (
	"crypto/tls"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationCreateStandingOrder(t *testing.T) {
	params := Params{
		Username:      "SUPERAPP",
		Password:      "123456",
		AccountNumber: "1000000006924",
		OrderId:       "10",
	}

	xmlRequest := NewCancleStandingOrder(params)
	endpoint := "https://devapisuperapp.cbe.com.et/superapp/parser/proxy/CBESUPERAPPV2/services?target=http://10.1.15.195:8080&wsdl=null"

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

	result, err := ParseCancelStandingOrderSOAP(string(responseData))
	assert.NoError(t, err)
	assert.NotNil(t, result, "Expected result to be non-nil")

	// Check that the lookup succeeded
	t.Logf("result: %v", result)
	assert.True(t, result.Success)
	assert.NotNil(t, result.Detail)

	if result.Detail != nil {
		assert.Equal(t, "1000357597823", result.Detail.CreditAccountNumber)
		assert.Equal(t, "124.00", result.Detail.Amount)
		assert.Equal(t, "20220512 e0Y e0M e1W e0D e0F", result.Detail.Frequency)
	} else {
		t.Error("Expected Detail to be non-nil")
	}
}
