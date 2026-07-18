package customercreation

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	sandboxOAuthTokenURL = "https://devapisuperapp.cbe.com.et/superapp/parser/proxy/cbe-dev/sandbox/oauth-mb-cbebirr/oauth2/token?target=https%3A%2F%2Fapi-gw-uat-gateway-apic-nonprod.apps.cp4itest.cbe.local"
	sandboxCustomerURL   = "https://devapisuperapp.cbe.com.et/superapp/parser/proxy/cbe-dev/sandbox/cust_creation?target=https%3A%2F%2Fapi-gw-uat-gateway-apic-nonprod.apps.cp4itest.cbe.local"
	sandboxOAuthBody     = "grant_type=client_credentials&client_id=f1ceebd8d6d5b802dc7fd8332ab33603&client_secret=05ac01f2f134bfff2549669bc11bd6cc&scope=mb-cbebirr-scope"
)

type oauthTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

func fetchSandboxAccessToken(t *testing.T, client *http.Client) string {
	t.Helper()

	t.Logf("fetching sandbox access token from: %s\n", sandboxOAuthTokenURL)

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

	t.Logf("oauth response status: %s\n", resp.Status)
	t.Logf("oauth response body: %s\n", string(responseData))

	var tokenResp oauthTokenResponse
	require.NoError(t, json.Unmarshal(responseData, &tokenResp))
	require.NotEmpty(t, tokenResp.AccessToken)

	t.Logf("access token received, type: %s, length: %d\n", tokenResp.TokenType, len(tokenResp.AccessToken))

	return tokenResp.AccessToken
}

func TestIntegrationCreateCustomer(t *testing.T) {
	tinNumber := fmt.Sprintf("%016d", rand.Int63n(1e16))
	nationalId := fmt.Sprintf("%018d%018d", rand.Int63n(1e18), rand.Int63n(1e18))
	email := fmt.Sprintf("sampletet%d@gmail.com", rand.Int63n(1e16))
	phoneNumber := fmt.Sprintf("+25191%06d", rand.Intn(1000000))
	legalId := fmt.Sprintf("WS%d", rand.Int63n(1e14))

	params := Params{
		Username:           "SUPERAPP",
		Password:           "123456",
		Company:            "ET0011859",
		FirstName:          "KETEM",
		MiddleName:         "HAILU",
		LastName:           "TAYE",
		Menmonic:           "K" + phoneNumber[len(phoneNumber)-9:],
		PhoneNumber:        phoneNumber,
		Address:            "ADDIS ABABA",
		PostalCode:         "4144",
		ISOCountryCode:     "ET",
		ISONationalityCode: "ET",
		ISOResidentCode:    "ET",
		UniqueID:           legalId,
		LegalDocumenetName: "NATIONAL.ID",
		IssuesBy:           "FAYDA",
		IssuedDate:         "20210504",
		ExpiryDate:         "20500101",
		Title:              "MR",
		Gender:             "MALE",
		DateOfBirth:        "19900310",
		MaritalStatus:      "SINGLE",
		NoOfDependents:     "0",
		Email:              email,
		EmploymentStatus:   "EMPLOYED",
		Occupation:         "HIRED",
		CustomerCurrency:   "ETB",
		Salary:             "75000",
		Street:             "AM",
		TownCountry:        "ADDIS ABABA",
		TinNumber:          tinNumber,
		CustomerOccupation: "Banker",
		EducationStatus:    "First Degree",
		MotherName:         "TIGIST ADAM FEKADU",
		FATCACompliant:     "NO",
		USPerson:           "NO",
		KebeleHNO:          "1544/02",
		CustomerSubSegment: "MASS",
		CustomerSegment:    "MASS",
		GrandFatherName:    "TAYE",
		CustomerGroup:      "RETAIL",
		NationalId:         nationalId,
	}

	client := &http.Client{
		Timeout: 60 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
		},
	}

	accessToken := fetchSandboxAccessToken(t, client)

	xmlRequest := NewCustomerCreation(params)
	t.Logf("request params: %+v\n", params)
	t.Logf("xmlRequest: %s\n", xmlRequest)

	req, err := http.NewRequest(http.MethodPost, sandboxCustomerURL, strings.NewReader(xmlRequest))
	require.NoError(t, err)

	req.Header.Set("Content-Type", "application/xml")
	req.Header.Set("Authorization", "Bearer "+accessToken)
	t.Logf("customer creation endpoint: %s\n", sandboxCustomerURL)
	t.Logf("request headers: Content-Type=%s, Authorization=Bearer <token>\n", req.Header.Get("Content-Type"))

	resp, err := client.Do(req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	defer resp.Body.Close()

	t.Logf("customer creation response status: %s\n", resp.Status)

	responseData, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NotEmpty(t, responseData, "Expected response body to be non-empty")

	t.Logf("core response: %s\n", string(responseData))
	result, err := ParseCustomerCreationSOAP(string(responseData))
	require.NoError(t, err)
	require.NotNil(t, result, "Expected result to be non-nil")

	t.Logf("parsed result success: %v", result.Success)
	t.Logf("parsed result messages: %v", result.Messages)
	t.Logf("parsed result detail: %+v", result.Detail)

	if !result.Success {
		t.Logf("customer creation failed: %v", result.Messages)
		require.FailNow(t, "customer creation API returned failure", "messages=%v", result.Messages)
	}

	require.NotNil(t, result.Detail)
	t.Logf("Customer Number: %s", result.Detail.CustomerNumber)
	t.Logf("customer mnemonic: %s", result.Detail.Menmonic)
	t.Logf("customer full name: %s", result.Detail.FullName)
	t.Logf("customer phone: %s", result.Detail.PhoneNumber)
	t.Logf("customer national id: %s", result.Detail.NationalId)
	t.Logf("customer ownership: %s", result.Detail.Ownership)
	t.Logf("customer cocode: %s", result.Detail.Cocode)

	assert.NotEmpty(t, result.Detail.CustomerNumber)
	assert.NotEmpty(t, result.Detail.Menmonic)
	assert.NotEmpty(t, result.Detail.FullName)
	assert.Equal(t, "ADDIS ABABA", result.Detail.Address)
	assert.Equal(t, "ET", result.Detail.Country)
	assert.Equal(t, "ET", result.Detail.Nationality)
	assert.Equal(t, "MALE", result.Detail.Gender)
	assert.Equal(t, "EMPLOYED", result.Detail.EmploymentStatus)
	assert.Equal(t, "ETB", result.Detail.Currency)
	assert.Equal(t, "RETAIL", result.Detail.CustomerGroup)
	assert.Equal(t, "3000", result.Detail.Ownership)
}
