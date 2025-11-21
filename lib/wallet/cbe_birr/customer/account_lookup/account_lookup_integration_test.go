package accountlookup

import (
	"crypto/tls"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationCustomerAccountLookup(t *testing.T) {

	// Use a local rand.Rand seeded from time to avoid the deprecated global rand.Seed.
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(100078546981)
	params := Params{
		Password:                        "8eZVmhR2RmGWW/1P8DjLDpHiiiLUle0u",
		OriginalConverstationIdentifier: strconv.Itoa(n),
		ThirdPartyIdentifier:            "USSDPushCaller",
		Timestamp:                       "20130402152345",
		SecurityCredential:              "BWJ3KefDOdp+GHqRnA9Yfo2RbsZM60sw",
		PhoneNumber:                     "251913170005",
	}

	xmlRequest := NewCustomerAccountLookup(params)
	endpoint := "https://devapisuperapp.cbe.com.et/superapp/parser/proxy/cbe-dev/sandbox/mb_cbebirr_sync?target=https://api-gw-uat-gateway-apic-nonprod.apps.cp4itest.cbe.local"

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(xmlRequest))
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/xml")
	req.Header.Set("iib_authorization", "Basic VW5pZmllZDpQYXNzd29yZA==")
	req.Header.Set("Authorization", "Bearer AAIgZjFjZWViZDhkNmQ1YjgwMmRjN2ZkODMzMmFiMzM2MDMLanMADavHe19ILn7YTaiCDJug3yHSpJB-vMw7oY7XO3k1U9vpjVVqG1JjwQ7eWfkkn-68xAzQ0s9AT5mBblvrzetgk9H1NZyOWnGV15Pb3YD9vUhRIfLaRqyETnl5-rn_K-fKuek2fBcFxAV06xv-")

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
	if result == nil {
		t.Fatal("Expected result to be non-nil")
	}

	// Check that the lookup succeeded
	assert.True(t, result.Success)
	assert.NotNil(t, result.Detail)

	if result.Detail != nil {
		t.Logf("result: %v", result.Detail.CustomerKYCData)
		assert.Equal(t, strconv.Itoa(n), result.Detail.OriginalConverstationIdentifier)
		assert.Equal(t, "1.0", result.Detail.Version)
	} else {
		t.Error("Expected Detail to be non-nil")
	}
}
