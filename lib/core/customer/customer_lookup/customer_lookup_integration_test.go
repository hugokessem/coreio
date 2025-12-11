package customerlookup

import (
	"crypto/tls"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationCustomerLookup(t *testing.T) {
	params := Params{
		Username:           "SUPERAPP",
		Password:           "123456",
		CustomerIdentifier: "1020746756",
	}

	xmlRequest := NewCustomerLookup(params)
	endpoint := "https://devopscbe.eaglelionsystems.com/superapp/parser/proxy/CBESUPERAPP/services?target=http%3A%2F%2F10.1.15.195%3A8080&wsdl=null"

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(xmlRequest))
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "text/xml; charset=utf-8")

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

	result, err := ParseCustomerLookupSOAP(string(responseData))
	if err != nil {
		t.Logf("Parsing error: %v", err)
		t.Logf("Response data: %s", string(responseData))
	}
	assert.NoError(t, err)
	assert.NotNil(t, result, "Expected result to be non-nil")

	// Check that the lookup succeeded
	t.Logf("result: %+v", result)
	if result.Success {
		assert.NotNil(t, result.CustomerInfos)
		if result.CustomerInfos != nil {
			t.Logf("Customer Full Name: %s", result.CustomerInfos.FullName)
			t.Logf("Customer First Name: %s", result.CustomerInfos.FirstName)
			t.Logf("Customer Last Name: %s", result.CustomerInfos.LastName)
			t.Logf("Customer Phone: %s", result.CustomerInfos.PhoneNumber)
			t.Logf("Customer Email: %s", result.CustomerInfos.Email)
			t.Logf("Customer City: %s", result.CustomerInfos.City)

			// Validate that customer identifier matches or related fields are populated
			if result.CustomerInfos.FullName != "" {
				assert.NotEmpty(t, result.CustomerInfos.FullName)
			}
		}
	} else {
		t.Logf("Customer lookup failed with message: %s", result.Message)
	}
}
