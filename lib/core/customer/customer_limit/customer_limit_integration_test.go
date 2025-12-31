package customerlimit

import (
	"crypto/tls"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationCustomerLimit(t *testing.T) {
	params := Params{
		Username:      "SUPERAPP",
		Password:      "123456",
		TransactionID: "GLOBAL",
	}

	xmlRequest := NewCustomerLimit(params)
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

	result, err := ParseCustomerLimitSOAP(string(responseData))
	assert.NoError(t, err)
	assert.NotNil(t, result, "Expected result to be non-nil")

	// Check that the lookup succeeded
	assert.True(t, result.Success)
	assert.NotNil(t, result.Detail)
	t.Logf("result: %v", result.Detail)
	t.Logf("result: %v", result.Detail)

	if result.Detail != nil {
		assert.NotEmpty(t, result.Detail.CustomerMaxLimit)
		assert.NotEmpty(t, result.Detail.CustomerMinLimit)
		assert.NotEmpty(t, result.Detail.AccountMaxLimit)
		assert.NotEmpty(t, result.Detail.AccountMinLimit)
		t.Logf("Customer Limit result: CustomerMaxLimit=%s, CustomerMinLimit=%s, AccountMaxLimit=%s, AccountMinLimit=%s",
			result.Detail.CustomerMaxLimit, result.Detail.CustomerMinLimit,
			result.Detail.AccountMaxLimit, result.Detail.AccountMinLimit)

		// Log channel types if available
		if result.Detail.ChannelTypeGroup != nil && len(result.Detail.ChannelTypeGroup.Details) > 0 {
			t.Logf("Channel Types: %d found", len(result.Detail.ChannelTypeGroup.Details))
			for i, ct := range result.Detail.ChannelTypeGroup.Details {
				t.Logf("  [%d] Type=%s, Max=%s, Min=%s", i+1, ct.ChannelType, ct.ChannelMaxLimit, ct.ChannelMinLimit)
			}
		}

		// Log conv role types if available
		if result.Detail.ConvRoleTypeGroup != nil && len(result.Detail.ConvRoleTypeGroup.Details) > 0 {
			t.Logf("Conv Role Types: %d found", len(result.Detail.ConvRoleTypeGroup.Details))
			for i, crt := range result.Detail.ConvRoleTypeGroup.Details {
				t.Logf("  [%d] Type=%s, Max=%s, Min=%s", i+1, crt.ConvRoleType, crt.ConvRoleMaxLimit, crt.ConvRoleMinLimit)
			}
		}

		// Log ifb role types if available
		if result.Detail.IfbRoleTypeGroup != nil && len(result.Detail.IfbRoleTypeGroup.Details) > 0 {
			t.Logf("IFB Role Types: %d found", len(result.Detail.IfbRoleTypeGroup.Details))
			for i, irt := range result.Detail.IfbRoleTypeGroup.Details {
				t.Logf("  [%d] Type=%s, Max=%s, Min=%s", i+1, irt.IfbRoleType, irt.IfbRoleMaxLimit, irt.IfbRoleMinLimit)
			}
		}

		t.Log("Integration test passed")
	} else {
		t.Error("Expected Detail to be non-nil")
	}
}
