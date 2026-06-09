package unsbscribe

import (
	"crypto/tls"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationRevertFundTransfer(t *testing.T) {
	params := Params{
		Username: "SUPERAPP",
		Password: "123456",
		UserCode: "SA1771239173",
	}

	xmlRequest := NewUnsubscribe(params)
	endpoint := "https://devapisuperapp.cbe.com.et/superapp/parser/proxy/CBESUPERAPP/services?target=http://10.1.15.195%3A8080&wsdl=null"

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(xmlRequest))
	assert.NoError(t, err)

	req.Header.Add("Content-Type", "text/xml; charset=utf-8")

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

	result, err := ParseUnsubscribeResponseSOAP(string(responseData))
	assert.NoError(t, err)
	assert.NotNil(t, result, "Expected result to be non-nil")

	// Check that the lookup succeeded
	assert.True(t, result.Success)
	assert.NotNil(t, result.Detail)

	if result.Detail != nil {
		assert.Equal(t, "SA1771239173", result.Detail.CustomerId)
		assert.Equal(t, "Mr Yohannes Teshome", result.Detail.CustomerName)
		assert.Equal(t, "+251911706628", result.Detail.PhoneNumber)
		assert.Equal(t, "yohannes@yml.com", result.Detail.Email)
		assert.Equal(t, "ALL", result.Detail.Channel)
		assert.Equal(t, "New User", result.Detail.Description)
	} else {
		t.Error("Expected Detail to be non-nil")
	}
}
