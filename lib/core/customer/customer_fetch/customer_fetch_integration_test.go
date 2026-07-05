package customerfetch

import (
	"crypto/tls"
	"io"
	"net/http"
	"strings"
	"testing"

	valueobject "github.com/hugokessem/coreio/lib/core/customer/customer_fetch/value_object"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationCustomerFetchByCustomerNumber(t *testing.T) {
	params := Params{
		Username:       "SUPERAPP",
		Password:       "123456",
		FetchBy:        valueobject.FetchByAccountNumber.String(),
		CustomerNumber: "1000301764441",
	}

	xmlRequest := NewCustomerFetch(params)
	t.Log("XML Request:", xmlRequest)

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

	result, err := ParseCustomerFetchSOAP(string(responseData))
	assert.NoError(t, err)
	assert.NotNil(t, result, "Expected result to be non-nil")

	assert.True(t, result.Success)
	t.Log("Details:", result.Details)
	lastIndex := len(result.Details) - 1
	t.Log("Last index:", lastIndex)
	t.Log("PhoneNumber: ", result.Details[lastIndex].Phone)
}

func TestIntegrationCustomerFetchByAccountNumber(t *testing.T) {
	params := Params{
		Username:       "SUPERAPP",
		Password:       "123456",
		FetchBy:        valueobject.FetchByAccountNumber.String(),
		CustomerNumber: "1000041045384696",
	}

	xmlRequest := NewCustomerFetch(params)
	t.Log("XML Request:", xmlRequest)

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

	result, err := ParseCustomerFetchSOAP(string(responseData))
	assert.NoError(t, err)
	assert.NotNil(t, result, "Expected result to be non-nil")

	assert.True(t, result.Success)
	t.Log("Details:", result.Details)
}
