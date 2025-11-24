package lockedamountrelease

import (
	"crypto/tls"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationReleaseLockedAmount(t *testing.T) {
	params := Params{
		Username:      "SUPERAPP",
		Password:      "123456",
		TransactionID: "ACLK213436WRSG",
	}

	xmlRequest := NewReleaseLockedAmount(params)
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

	result, err := ParseCancleLockedAmountSOAP(string(responseData))

	assert.NoError(t, err)
	assert.NotNil(t, result, "Expected result to be non-nil")

	// Check that the lookup succeeded
	assert.True(t, result.Success)
	assert.NotNil(t, result.Detail)

	if result.Detail != nil {
		assert.Equal(t, "1000000006924", result.Detail.AccountNumber)
		assert.Equal(t, "ACLK213436WRSG", result.Detail.TransactionID)
		assert.Equal(t, "4510.00", result.Detail.LockedAmount)
	} else {
		t.Error("Expected Detail to be non-nil")
	}
}
