package fundtransfer

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

func TestIntegrationAgentFundTransfer(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ftNumber := r.Intn(100078546981)

	params := Params{
		FTNumber:               strconv.Itoa(ftNumber),
		Timestamp:              time.Now().Format("20060102150405"),
		PrimaryParty:           "000099",
		ReceiverParty:          "251000",
		Amount:                 "124.00",
		Currency:               "ETB",
		Narative:               "Integration test",
		ThirdPartyIdentifier:   "USSDPushCaller",
		Password:               "8eZVmhR2RmGWW/1P8DjLDpHiiiLUle0u",
		SecurityCredential:     "BWJ3KefDOdp+GHqRnA9Yfo2RbsZM60sw",
		DebitAccountNumber:     "1000184084108",
		DebitAccountHolderName: "Elnatan Michael Michael",
	}

	xmlRequest := NewAgentFundTransfer(params)
	endpoint := "https://devapisuperapp.cbe.com.et/superapp/parser/proxy/cbe-dev/sandbox/mb_cbebirr_sync?target=https://api-gw-uat-gateway-apic-nonprod.apps.cp4itest.cbe.local"

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(xmlRequest))
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/xml")
	req.Header.Set("iib_authorization", "Basic VW5pZmllZDpQYXNzd29yZA==")
	req.Header.Set("Authorization", "Bearer AAIgZjFjZWViZDhkNmQ1YjgwMmRjN2ZkODMzMmFiMzM2MDMvZwY6NerYvcf5bD5UHo6YCGJ5bjjyk1kp77SK2DmFnPTMuorkYxE090DDB0XEY9PF8MkJTyYuyGrd4g2ywbxX7YoKJpisSnF5_YYl0nNGspW0evSOfcxhrd73m2Hr5_8YchqCuszX3AiNbfzBv5lT")

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

	result, err := ParseAgentFundTransfer(string(responseData))
	assert.NoError(t, err)
	assert.NotNil(t, result, "Expected result to be non-nil")

	if result == nil {
		t.Fatal("Expected result to be non-nil")
	}

	if result.Detail != nil {
		assert.True(t, result.Status)
		assert.NotEmpty(t, result.Detail.FTNumber)
		assert.NotEmpty(t, result.Detail.ConversationIdentifier)
		assert.Greater(t, len(result.Detail.ReferenceDetail), 0)
	} else {
		t.Error("Expected Detail to be non-nil")
	}
}
