package customerlimitamendment

import (
	"crypto/tls"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationCustomerLimitAmendment(t *testing.T) {
	params := Params{
		Username:         "SUPERAPP",
		Password:         "123456",
		CustomerID:       "1026582446",
		AppUserMaxLimit:  "800005",
		USSDUserMaxLimit: "150005",
	}

	xmlRequest := NewCustomerLimitAmendment(params)
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

	result, err := ParseCustomerLimitAmendmentSOAP(string(responseData))
	if err != nil {
		t.Logf("Parse error: %v", err)
		t.Logf("Response data (first 500 chars): %s", string(responseData)[:min(500, len(responseData))])
	}
	assert.NoError(t, err)
	assert.NotNil(t, result, "Expected result to be non-nil")

	if result == nil {
		t.Fatal("Result is nil, cannot continue")
	}

	// Check that the amendment succeeded
	if !result.Success {
		t.Logf("Amendment failed. Messages: %v", result.Message)
		t.Logf("Response data (first 1000 chars): %s", string(responseData)[:min(1000, len(responseData))])
	}
	assert.True(t, result.Success)
	assert.NotNil(t, result.Detail)
	t.Logf("result: %v", result.Detail)

	if result.Detail != nil {
		// Log user channel types if available
		if len(result.Detail.UserChannelType) > 0 {
			t.Logf("User Channel Types: %d found", len(result.Detail.UserChannelType))
			for i, uct := range result.Detail.UserChannelType {
				t.Logf("  [%d] ChannelType=%s, MaxLimit=%s", i+1, uct.UserChannelType, uct.UserMaxLimit)
			}
		}

		// Validate that detail fields are populated
		assert.Greater(t, len(result.Detail.UserChannelType), 0, "Expected at least one user channel type")

		// Log other detail fields
		t.Logf("Customer Limit Amendment result: UserMaxLimit=%s, UserCustomerId=%s, Account=%s, Currency=%s",
			result.Detail.UserMaxLimit, result.Detail.UserCustomerId,
			result.Detail.Account, result.Detail.Currency)
		t.Logf("Inputter=%s, Datetime=%s, Authoriser=%s, CoCode=%s, DeptCode=%s",
			result.Detail.Inputter, result.Detail.Datetime, result.Detail.Authoriser,
			result.Detail.Cocode, result.Detail.Deptcode)

		t.Log("Integration test passed")
	} else {
		t.Error("Expected Detail to be non-nil")
	}
}
