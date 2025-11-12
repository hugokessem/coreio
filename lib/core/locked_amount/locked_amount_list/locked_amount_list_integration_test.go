package lockedamountlist

import (
	"crypto/tls"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationLockedAmount(t *testing.T) {
	params := Params{
		Username:      "SUPERAPP",
		Password:      "123456",
		AccountNumber: "1000000006924",
	}

	xmlRequest := NewListLockedAmount(params)
	endpoint := "https://devapisuperapp.cbe.com.et/superapp/parser/proxy/CBESUPERAPPV2/services?target=http%3A%2F%2F10.1.15.195%3A8080&wsdl=null"

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

	result, err := ParseListLockedAmountSOAP(string(responseData))

	assert.NoError(t, err)
	assert.NotNil(t, result, "Expected result to be non-nil")

	// Check that the lookup succeeded
	assert.True(t, result.Success)
	assert.Greater(t, len(result.Details), 0)

	if len(result.Details) > 0 {
		assert.Equal(t, "1000000006924", result.Details[0].AccountNumber)
		assert.Equal(t, "20120501", result.Details[0].LockedDate)
		assert.Equal(t, "100", result.Details[0].LockedAmount)
	} else {
		t.Error("Expected Detail to be non-nil")
	}
}
