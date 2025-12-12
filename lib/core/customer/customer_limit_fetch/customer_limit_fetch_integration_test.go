package customerlimitfetch

import (
	"crypto/tls"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationCustomerLimitFetch(t *testing.T) {
	params := Param{
		Username:       "SUPERAPP",
		Password:       "123456",
		CustomerNumber: "1026582446",
	}

	xmlRequest := NewCustomerLimitFetch(params)
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

	result, err := ParseCustomerLimitFetchSOAP(string(responseData))
	assert.NoError(t, err)
	assert.NotNil(t, result, "Expected result to be non-nil")

	// Check that the lookup succeeded
	assert.True(t, result.Success)
	assert.NotNil(t, result.Detail)
	t.Logf("result: %v", result.Detail)

	if result.Detail != nil {
		// Log basic fields
		if result.Detail.CurrNo != "" {
			t.Logf("Currency Number: %s", result.Detail.CurrNo)
		}
		if result.Detail.Authoriser != "" {
			t.Logf("Authoriser: %s", result.Detail.Authoriser)
		}
		if result.Detail.CoCode != "" {
			t.Logf("Company Code: %s", result.Detail.CoCode)
		}
		if result.Detail.DeptCode != "" {
			t.Logf("Department Code: %s", result.Detail.DeptCode)
		}

		// Log user channel types if available
		if result.Detail.UserChannelType != nil && len(result.Detail.UserChannelType.Details) > 0 {
			t.Logf("User Channel Types: %d found", len(result.Detail.UserChannelType.Details))
			for i, uct := range result.Detail.UserChannelType.Details {
				t.Logf("  [%d] ChannelType=%s, MaxLimit=%s", i+1, uct.UserChannelType, uct.UserMaxLimit)
			}
		} else {
			t.Log("No user channel types found in response")
		}

		// Log global inputter if available
		if result.Detail.GlobalInputter != nil && result.Detail.GlobalInputter.Inputter != "" {
			t.Logf("Inputter: %s", result.Detail.GlobalInputter.Inputter)
		}

		// Log global datetime if available
		if result.Detail.GlobalDatetime != nil && result.Detail.GlobalDatetime.Datetime != "" {
			t.Logf("DateTime: %s", result.Detail.GlobalDatetime.Datetime)
		}

		t.Log("Integration test passed")
	} else {
		t.Error("Expected Detail to be non-nil")
	}
}

func TestIntegrationCustomerLimitFetch_WithDifferentCustomerNumber(t *testing.T) {
	params := Param{
		Username:       "SUPERAPP",
		Password:       "123456",
		CustomerNumber: "GLOBAL",
	}

	xmlRequest := NewCustomerLimitFetch(params)
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

	result, err := ParseCustomerLimitFetchSOAP(string(responseData))
	assert.NoError(t, err)
	assert.NotNil(t, result, "Expected result to be non-nil")

	// The result may or may not be successful depending on the customer number
	// We just verify that parsing works correctly
	if result.Success {
		assert.NotNil(t, result.Detail)
		t.Logf("Successfully fetched customer limit for customer: %s", params.CustomerNumber)
	} else {
		t.Logf("Customer limit fetch returned failure for customer: %s, Message: %s", params.CustomerNumber, result.Message)
	}
}

