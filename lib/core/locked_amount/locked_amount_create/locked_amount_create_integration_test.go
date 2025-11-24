package lockedamountcreate

import (
	"crypto/tls"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationCreateLockedAmount(t *testing.T) {
	params := Params{
		Username:      "SUPERAPP",
		Password:      "123456",
		AccountNumber: "1000000006924",
		Description:   "3 Click Payment",
		From:          "20251109",
		To:            "20251111",
		LockedAmount:  "250",
	}

	xmlRequest := NewCreateLockedAmount(params)
	endpoint := "https://devapisuperapp.cbe.com.et/superapp/parser/proxy/CBESUPERAPP/services?target=http%3A%2F%2F10.1.15.195%3A8080&wsdl=null"

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

	result, err := ParseCreateLockedAmountSOAP(string(responseData))

	assert.NoError(t, err)
	assert.NotNil(t, result, "Expected result to be non-nil")

	// Check that the lookup succeededt
	t.Log(result)

	assert.True(t, result.Success)
	assert.NotNil(t, result.Detail)

	if result.Detail != nil {
		assert.Equal(t, "20251111", result.Detail.To)
		assert.Equal(t, "20251109", result.Detail.From)
		assert.Equal(t, "250.00", result.Detail.LockedAmount)
		assert.Equal(t, "3 Click Payment", result.Detail.Description)
		assert.Equal(t, "1000000006924", result.Detail.AccountNumber)
	} else {
		t.Error("Expected Detail to be non-nil")
	}
}
