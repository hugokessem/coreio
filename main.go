package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/hugokessem/coreio/core"
)

const (
	sandboxOAuthTokenURL = "https://devapisuperapp.cbe.com.et/superapp/parser/proxy/cbe-dev/sandbox/oauth-mb-cbebirr/oauth2/token?target=https%3A%2F%2Fapi-gw-uat-gateway-apic-nonprod.apps.cp4itest.cbe.local"
	sandboxCustomerURL   = "https://devapisuperapp.cbe.com.et/superapp/parser/proxy/cbe-dev/sandbox/cust_creation?target=https%3A%2F%2Fapi-gw-uat-gateway-apic-nonprod.apps.cp4itest.cbe.local"
	sandboxOAuthBody     = "grant_type=client_credentials&client_id=f1ceebd8d6d5b802dc7fd8332ab33603&client_secret=05ac01f2f134bfff2549669bc11bd6cc&scope=mb-cbebirr-scope"
)

type CoreAPI struct {
	coreInterface core.CBECoreAPIInterface
}

type oauthTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

func InitCoreAPICalls(username, password, url string) CoreAPI {
	return CoreAPI{
		coreInterface: core.NewCBECoreAPI(&core.CBECoreCredential{
			Username: username,
			Password: password,
			Url:      url,
		}),
	}
}

func (c *CoreAPI) CreateCustomer(ctx context.Context, param core.CreateCustomerParam) (*core.CreateCustomerResult, error) {
	result, err := c.coreInterface.CreateCustomer(ctx, param)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func getAccessToken(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, sandboxOAuthTokenURL, strings.NewReader(sandboxOAuthBody))
	if err != nil {
		return "", fmt.Errorf("create oauth request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("oauth request failed: %w", err)
	}
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read oauth response: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("oauth request returned %s: %s", resp.Status, string(responseData))
	}

	var tokenResp oauthTokenResponse
	if err := json.Unmarshal(responseData, &tokenResp); err != nil {
		return "", fmt.Errorf("parse oauth response: %w", err)
	}
	if tokenResp.AccessToken == "" {
		return "", fmt.Errorf("oauth response missing access_token")
	}

	return tokenResp.AccessToken, nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	accessToken, err := getAccessToken(ctx)
	if err != nil {
		log.Fatalf("failed to get access token: %v", err)
	}

	calls := InitCoreAPICalls(
		"SUPERAPP",
		"123456",
		sandboxCustomerURL,
	)

	tinNumber := fmt.Sprintf("%016d", rand.Intn(10000000000000000))
	nationalID := fmt.Sprintf("%016d", rand.Intn(10000000000000000))
	email := fmt.Sprintf("sampletet%d@gmail.com", rand.Intn(10000000000000000))
	phoneNumber := fmt.Sprintf("+25191%06d", rand.Intn(1000000))
	legalID := fmt.Sprintf("%016d", rand.Intn(10000000000000000))

	customer := core.CreateCustomerParam{
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
		UniqueID:           legalID,
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
		NationalId:         nationalID,
		Url:                sandboxCustomerURL,
		Header: map[string]string{
			"Authorization": "Bearer " + accessToken,
		},
	}

	result, err := calls.CreateCustomer(ctx, customer)
	if err != nil {
		log.Fatalf("create customer failed: %v", err)
	}

	if !result.Success {
		log.Fatalf("create customer rejected: %v", result.Messages)
	}

	fmt.Printf("customer created successfully\n")
	if result.Detail != nil {
		fmt.Printf("mnemonic: %s\n", result.Detail.Menmonic)
		fmt.Printf("full name: %s\n", result.Detail.FullName)
		fmt.Printf("phone: %s\n", result.Detail.PhoneNumber)
		fmt.Printf("national id: %s\n", result.Detail.NationalId)
		fmt.Printf("cocode: %s\n", result.Detail.Cocode)
		fmt.Printf("customer number: %s\n", result.Detail.Address)
	}
}
