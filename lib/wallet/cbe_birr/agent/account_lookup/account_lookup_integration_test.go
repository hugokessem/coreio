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
		Password:                        "8eZVmhR2RmGWW/1P8DjLDpHiiiLUle0u",
		OriginalConverstationIdentifier: strconv.Itoa(n),
		ThirdPartyIdentifier:            "USSDPushCaller",
		Timestamp:                       time.Now().Format("20060102150405"),
		SecurityCredential:              "BWJ3KefDOdp+GHqRnA9Yfo2RbsZM60sw",
		PhoneNumber:                     "251000",
	}

	xmlRequest := NewAgentAccountLookup(params)
	endpoint := "https://devapisuperapp.cbe.com.et/superapp/parser/proxy/cbe-dev/sandbox/mb_cbebirr_sync?target=https://api-gw-uat-gateway-apic-nonprod.apps.cp4itest.cbe.local"

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(xmlRequest))
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/xml")
	req.Header.Set("iib_authorization", "Basic VW5pZmllZDpQYXNzd29yZA==")
	req.Header.Set("Authorization", "Bearer AIgZjFjZWViZDhkNmQ1YjgwMmRjN2ZkODMzMmFiMzM2MDMvZwY6NerYvcf5bD5UHo6YCGJ5bjjyk1kp77SK2DmFnPTMuorkYxE090DDB0XEY9PF8MkJTyYuyGrd4g2ywbxX7YoKJpisSnF5_YYl0nNGspW0evSOfcxhrd73m2Hr5_8YchqCuszX3AiNbfzBv5lT")

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
		assert.Equal(t, strconv.Itoa(n), result.Detail.OriginalConverstationIdentifier)
		assert.Equal(t, "1.0", result.Detail.Version)
	} else {
		t.Error("Expected Detail to be non-nil")
	}
}
