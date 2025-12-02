package cardreplace

import (
	"crypto/tls"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationCardReplace(t *testing.T) {
	params := Params{
		Username:          "SUPERAPP",
		Password:          "123456",
		AccountNumber:     "1000000006924",
		BranchCode:        "ET0010222",
		PhoneNumber:       "+251913323918",
		CardType:          "VIEL",
		ReplacementReason: "TEST",
		ProductType:       "PLASSTIC",
	}

	xmlRequest := NewCardReplace(params)
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

	result, err := ParseCardReplaceResponse(string(responseData))
	assert.NoError(t, err)
	assert.NotNil(t, result, "Expected result to be non-nil")

	// Check that the request succeeded
	assert.True(t, result.Success)
	assert.NotNil(t, result.Detail)
	t.Logf("result: %v", result.Detail)

	if result.Detail != nil {
		assert.Equal(t, "1000000006924", result.Detail.AccountNumber)
		assert.NotEmpty(t, result.Detail.AccountHolderName)
		assert.NotEmpty(t, result.Detail.BranchCode)
		assert.NotEmpty(t, result.Detail.PhoneNumber)
		assert.Equal(t, "VIEL", result.Detail.CardType)
		assert.Equal(t, "REPLACEMENT", result.Detail.RequestType)
		assert.Equal(t, "TEST", result.Detail.ReplacementReason)
		t.Log("Integration test passed")
	} else {
		t.Error("Expected Detail to be non-nil")
	}
}
