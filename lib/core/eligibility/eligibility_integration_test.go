package eligibility

import (
	"crypto/tls"
	"io"
	"net/http"
	"strings"
	"testing"

	valueobject "github.com/hugokessem/coreio/lib/core/eligibility/value_object"
	"github.com/stretchr/testify/assert"
)

const eligibilityEndpoint = "https://devapisuperapp.cbe.com.et/superapp/parser/proxy/CBESUPERAPP/services?target=http://10.1.15.195%3A8080&wsdl=null"

func postEligibilityRequest(t *testing.T, xmlRequest string) string {
	t.Helper()

	req, err := http.NewRequest("POST", eligibilityEndpoint, strings.NewReader(xmlRequest))
	assert.NoError(t, err)
	req.Header.Add("Content-Type", "text/xml; charset=utf-8")

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Logf("Network error (endpoint may be unreachable): %v", err)
		t.Skip("Skipping test due to network error - endpoint may be unreachable")
	}
	assert.NotNil(t, resp)
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, responseData, "Expected response body to be non-empty")

	return string(responseData)
}

func TestIntegrationEligibilityByAccountNumber(t *testing.T) {
	param := Param{
		Username:      "SUPERAPP",
		Password:      "123456",
		FetchBy:       valueobject.FetchByAccountNumber.String(),
		CriticalValue: "1000446123703",
	}

	xmlRequest := NewEligibility(param)
	t.Log("XML Request:", xmlRequest)

	responseData := postEligibilityRequest(t, xmlRequest)
	t.Log("Response:", responseData)

	result, err := ParseCustomerEligibilitySOAP(responseData)
	assert.NoError(t, err)
	assert.NotNil(t, result, "Expected result to be non-nil")

	t.Logf("Result: Success=%v, Messages=%v", result.Success, result.Messages)

	if !result.Success {
		return
	}

	assert.NotEmpty(t, result.Details)
	t.Log("Details:", result.Details)
}

func TestIntegrationEligibilityByCustomerNumber(t *testing.T) {
	param := Param{
		Username:      "SUPERAPP",
		Password:      "123456",
		FetchBy:       valueobject.FetchByCustomerNumber.String(),
		CriticalValue: "1045384696",
	}

	xmlRequest := NewEligibility(param)
	t.Log("XML Request:", xmlRequest)

	responseData := postEligibilityRequest(t, xmlRequest)
	t.Log("Response:", responseData)

	result, err := ParseCustomerEligibilitySOAP(responseData)
	assert.NoError(t, err)
	assert.NotNil(t, result, "Expected result to be non-nil")

	t.Logf("Result: Success=%v, Messages=%v", result.Success, result.Messages)

	if !result.Success {
		return
	}

	assert.NotEmpty(t, result.Details)
	t.Log("Details:", result.Details)
}
