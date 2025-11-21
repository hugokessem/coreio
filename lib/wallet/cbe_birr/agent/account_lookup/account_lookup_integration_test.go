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

func TestIntegrationAgentAccountLookup(t *testing.T) {

	// Use a local rand.Rand seeded from time to avoid the deprecated global rand.Seed.
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(100078546981)
	params := Params{
		Password:                        "8eZVmhR2RmGWW/991P8DjLDpHiiiLUle0u",
		OriginalConverstationIdentifier: strconv.Itoa(n),
		ThirdPartyIdentifier:            "USSDPushCaller",
		Timestamp:                       "20130402152345",
		SecurityCredential:              "BWJ3KefDOdp+GHqRnA9Yfo2RbsZM60sw",
		PhoneNumber:                     "251000",
	}

	xmlRequest := NewAgentAccountLookup(params)
	endpoint := "https://devapisuperapp.cbe.com.et/superapp/parser/proxy/cbe-dev/sandbox/mb_cbebirr_sync?target=https://api-gw-uat-gateway-apic-nonprod.apps.cp4itest.cbe.local"

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(xmlRequest))
	assert.NoError(t, err)

	req.Header.Add("Content-Type", "application/xml")
	req.Header.Add("iib_authorization", "Basic VW5pZmllZDpQYXNzd29yZA==")
	req.Header.Add("Authorization", "Bearer AAIgZjFjZWViZDhkNmQ1YjgwMmRjN2ZkODMzMmFiMzM2MDMU0v4nVOKXPo-Deygtqcvx5L5NhqfsNi-8Xu9idc6hCCD_hgaJ1X3mwCboftG3UThc-7aa7Xfb2E9fr6QKCoayTCJfidHktrJh33pjlC678iTjP3VofIqlnWqNSHB5Zgv4lfP-ckHekIk1yFNbCXyo")

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

	result, err := ParseAgentLookupSOAP(string(responseData))
	assert.NoError(t, err)
	assert.NotNil(t, result, "Expected result to be non-nil")

	// Check that the lookup succeeded
	t.Logf("result: %v", result)
	assert.True(t, result.Success)
	assert.NotNil(t, result.Detail)

	if result.Detail != nil {
		assert.Equal(t, strconv.Itoa(n), result.Detail.ConversationIdentifier)
		assert.Equal(t, "1.0", result.Detail.Version)
	} else {
		t.Error("Expected Detail to be non-nil")
	}
}
