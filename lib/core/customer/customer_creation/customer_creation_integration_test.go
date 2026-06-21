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
	tinNumber := fmt.Sprintf("%016d", rand.Intn(10000000000000000))
	// nationalId := fmt.Sprintf("%016d", rand.Intn(10000000000000000))
	email := fmt.Sprintf("sampletet%d@gmail.com", rand.Intn(10000000000000000))
	phoneNumber := fmt.Sprintf("+25191%06d", rand.Intn(1000000))
	// legalId := fmt.Sprintf("%016d", rand.Intn(10000000000000000))

	params := Params{
		Username:           "SUPERAPP",
		Password:           "123456",
		FirstName:          "KETEM",
		MiddleName:         "HAILU",
		LastName:           "TAYE",
		PhoneNumber:        phoneNumber,
		Address:            "ADDIS ABABA",
		PostalCode:         "4144",
		ISOCountryCode:     "ET",
		AccountOffice:      "7124",
		Industry:           "1499",
		ISONationalityCode: "ET",
		ISOResidentCode:    "ET",
		UniqueID:           "09as8df0a9s8f",
		IssuesBy:           "FAYDA",
		IssuedDate:         "20210504",
		ExpiryDate:         "20500101",
		Gender:             "MALE",
		DateOfBirth:        "19900310",
		MaritalStatus:      "SINGLE",
		Email:              email,
		EmploymentStatus:   "EMPLOYED",
		Occupation:         "HIRED",
		EmployerName:       "MIDRO",
		EmployerAddress:    "ADDIS ABABA",
		EmployerBusiness:   "SHARE COMPANY",
		CustomerCurrency:   "ETB",
		Salary:             "75000",
		AnnualBonus:        "50000",
		NetMonthlyIncome:   "55000",
		NetMonthlyExpence:  "42000",
		TinNumber:          tinNumber,
		MotherName:         "TIGIST ADAM FEKADU",
		CustomerGroup:      "RETAIL",
		NationalId:         "0787a0980987",
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
	t.Logf("Customr Number: %+v", result.Detail.CustomerNumber)

	if !result.Success {
		t.Logf("customer creation failed")
	}

	assert.True(t, result.Success)
	assert.NotNil(t, result.Detail)

	if result.Detail != nil {
		t.Logf("customer mnemonic: %s", result.Detail.Menmonic)
		t.Logf("customer full name: %s", result.Detail.FullName)
		t.Logf("customer phone: %s", result.Detail.PhoneNumber)
		t.Logf("customer national id: %s", result.Detail.NationalId)
		t.Logf("customer ownership: %s", result.Detail.Ownership)
		t.Logf("customer cocode: %s", result.Detail.Cocode)

		assert.Equal(t, "K912895609", result.Detail.Menmonic)
		assert.Equal(t, "KETEM HAILU TAYE", result.Detail.FullName)
		assert.Equal(t, "ADDIS ABABA", result.Detail.Address)
		assert.Equal(t, "4144", result.Detail.PostalCode)
		assert.Equal(t, "ET", result.Detail.Country)
		assert.Equal(t, "7124", result.Detail.AccountOfficer)
		assert.Equal(t, "1499", result.Detail.Industry)
		assert.Equal(t, "ET", result.Detail.Nationality)
		assert.Equal(t, "20210504", result.Detail.IssuedDate)
		assert.Equal(t, "20500101", result.Detail.ExpiryDate)
		assert.Equal(t, "ET0010001", result.Detail.CompanyBook)
		assert.Equal(t, "MALE", result.Detail.Gender)
		assert.Equal(t, "19900310", result.Detail.DateOfBirth)
		assert.Equal(t, "SINGLE", result.Detail.MaritalStatus)
		assert.Equal(t, "+251912895689", result.Detail.PhoneNumber)
		assert.Equal(t, "sampletet@gmail.com", result.Detail.Email)
		assert.Equal(t, "EMPLOYED", result.Detail.EmploymentStatus)
		assert.Equal(t, "75000.00", result.Detail.Salary)
		assert.Equal(t, "MIDRO", result.Detail.Customer)
		assert.Equal(t, "50000.00", result.Detail.AnnualBonus)
		assert.Equal(t, "ETB", result.Detail.Currency)
		assert.Equal(t, "7775854855128587", result.Detail.TinNumber)
		assert.Equal(t, "TIGIST ADAM FEKADU", result.Detail.MotherName)
		assert.Equal(t, "RETAIL", result.Detail.CustomerGroup)
		assert.Equal(t, "4455567887777455", result.Detail.NationalId)
		assert.Equal(t, "3000", result.Detail.Ownership)
		assert.Equal(t, "ET0010001", result.Detail.Cocode)
	} else {
		t.Error("Expected Detail to be non-nil")
	}
}
