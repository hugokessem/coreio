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

func TestIntegrationCustomerFundTransfer(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ftNumber := strconv.FormatInt(r.Int63(), 10)

	params := Params{
		FTNumber:               ftNumber,
		Password:               "8eZVmhR2RmGWW/1P8DjLDpHiiiLUle0u",
		Timestamp:              time.Now().Format("20060102150405"),
		SecurityCredential:     "BWJ3KefDOdp+GHqRnA9Yfo2RbsZM60sw",
		ThirdPartyIdentifier:   "USSDPushCaller",
		PrimaryParty:           "000099",
		ReceiverParty:          "251913170005",
		Amount:                 "10.00",
		Currency:               "ETB",
		Narative:               "Integration test",
		DebitAccountNumber:     "1000184084108",
		DebitAccountHolderName: "Elnatan Michael Michael",
	}

	xmlRequest := NewCustomerFundTransfer(params)
	endpoint := "https://devapisuperapp.cbe.com.et/superapp/parser/proxy/cbe-dev/sandbox/mb_cbebirr_sync?target=https://api-gw-uat-gateway-apic-nonprod.apps.cp4itest.cbe.local"

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(xmlRequest))
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/xml")
	req.Header.Set("iib_authorization", "Basic VW5pZmllZDpQYXNzd29yZA==")
	req.Header.Set("Authorization", "Bearer AAIgZjFjZWViZDhkNmQ1YjgwMmRjN2ZkODMzMmFiMzM2MDO3fYtB73GxhqX4-4KfZLCDbwE3FSIeTnjmWmtg2VbFzelBO5qfkl1yWX4-0MpU05fkUbqsakF6JkAPJ1_Pj_bAjz3p2QL4SKqKkB9y4T7ooYQ1GQicAT0Ps4S584ZJ9YXfcnZ1RoM0-l0woO_wjVJ3")

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
	t.Log("responseData", string(responseData))
	assert.NoError(t, err)
	assert.NotEmpty(t, responseData, "Expected response body to be non-empty")

	result, err := ParserCustomreFundTransfer(string(responseData))
	t.Log("result", result)
	assert.NoError(t, err)
	assert.NotNil(t, result, "Expected result to be non-nil")

	if result == nil {
		t.Fatal("Expected result to be non-nil")
	}

	assert.True(t, result.Status)
	if result.Detail != nil {
		assert.NotEmpty(t, result.Detail.FTNumber, "Expected FTNumber to be non-empty")
		assert.NotEmpty(t, result.Detail.ConverstationIdentifier, "Expected ConversationIdentifier to be non-empty")
		assert.Greater(t, len(result.Detail.ReferenceDetail), 0, "Expected ReferenceDetail to be non-empty")
	} else {
		t.Error("Expected Detail to be non-nil")
	}
}
