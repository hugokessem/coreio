package acccountcreation

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	sandboxOAuthTokenURL = "https://devapisuperapp.cbe.com.et/superapp/parser/proxy/cbe-dev/sandbox/oauth-mb-cbebirr/oauth2/token?target=https%3A%2F%2Fapi-gw-uat-gateway-apic-nonprod.apps.cp4itest.cbe.local"
	sandboxAccountURL    = "https://devapisuperapp.cbe.com.et/superapp/parser/proxy/cbe-dev/sandbox/acc-creation?target=https%3A%2F%2Fapi-gw-uat-gateway-apic-nonprod.apps.cp4itest.cbe.local"
	sandboxOAuthBody     = "grant_type=client_credentials&client_id=f1ceebd8d6d5b802dc7fd8332ab33603&client_secret=05ac01f2f134bfff2549669bc11bd6cc&scope=mb-cbebirr-scope"
)

type oauthTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

func fetchSandboxAccessToken(t *testing.T, client *http.Client) string {
	t.Helper()

	t.Logf("fetching sandbox access token from: %s", sandboxOAuthTokenURL)

	req, err := http.NewRequest(http.MethodPost, sandboxOAuthTokenURL, strings.NewReader(sandboxOAuthBody))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NotEmpty(t, responseData)

	t.Logf("oauth response status: %s", resp.Status)
	t.Logf("oauth response body: %s", string(responseData))

	var tokenResp oauthTokenResponse
	require.NoError(t, json.Unmarshal(responseData, &tokenResp))
	require.NotEmpty(t, tokenResp.AccessToken)

	t.Logf("access token received, type: %s, length: %d", tokenResp.TokenType, len(tokenResp.AccessToken))

	return tokenResp.AccessToken
}

func TestIntegrationCreateAccount(t *testing.T) {
	params := Params{
		Username:       "SUPERAPP",
		Password:       "123456",
		CustomerNumber: "1027958756",
		Category:       "6501",
		Currency:       "ETB",
		AccountOfficer: "7016",
		Url:            sandboxAccountURL,
	}

	client := &http.Client{
		Timeout: 60 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
		},
	}

	accessToken := fetchSandboxAccessToken(t, client)

	xmlRequest := NewAccountCreation(params)
	t.Logf("request params: %+v", params)
	t.Logf("xmlRequest: %s", xmlRequest)

	req, err := http.NewRequest(http.MethodPost, sandboxAccountURL, strings.NewReader(xmlRequest))
	require.NoError(t, err)

	req.Header.Set("Content-Type", "application/xml")
	req.Header.Set("Authorization", "Bearer "+accessToken)
	t.Logf("account creation endpoint: %s", sandboxAccountURL)
	t.Logf("request headers: Content-Type=%s, Authorization=Bearer <token>", req.Header.Get("Content-Type"))

	resp, err := client.Do(req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	defer resp.Body.Close()

	t.Logf("account creation response status: %s", resp.Status)

	responseData, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NotEmpty(t, responseData, "Expected response body to be non-empty")

	t.Logf("core response: %s", string(responseData))
	result, err := ParseAccountCreationSOAP(string(responseData))
	require.NoError(t, err)
	require.NotNil(t, result, "Expected result to be non-nil")

	t.Logf("parsed result success: %v", result.Success)
	t.Logf("parsed result messages: %v", result.Messages)
	t.Logf("parsed result detail: %+v", result.Detail)

	if !result.Success {
		t.Logf("account creation failed")
	}

	assert.True(t, result.Success)
	assert.NotNil(t, result.Detail)

	if result.Detail != nil {
		t.Logf("account number: %s", result.Detail.AccountNumber)
		t.Logf("customer number: %s", result.Detail.Customer)
		t.Logf("account title: %s", result.Detail.AccountTitle)
		t.Logf("currency: %s", result.Detail.Currency)
		t.Logf("account officer: %s", result.Detail.AccountOfficer)
		t.Logf("cocode: %s", result.Detail.CoCode)

		assert.NotEmpty(t, result.Detail.AccountNumber)
		assert.NotEmpty(t, result.Detail.Customer)
		assert.Equal(t, "6501", result.Detail.Category)
		assert.NotEmpty(t, result.Detail.AccountTitle)
		assert.Equal(t, "ETB", result.Detail.Currency)
		assert.Equal(t, "7016", result.Detail.AccountOfficer)
		assert.NotEmpty(t, result.Detail.OpeningDate)
		assert.NotEmpty(t, result.Detail.CoCode)
		assert.NotEmpty(t, result.Detail.ProductType)
		assert.NotEmpty(t, result.Detail.AltAcctTypes)
	} else {
		t.Error("Expected Detail to be non-nil")
	}
}
