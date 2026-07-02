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
	coreSOAPURL          = "https://devapisuperapp.cbe.com.et/superapp/parser/proxy/CBESUPERAPP/services?target=http://10.1.15.195%3A8080&wsdl=null"
	sandboxOAuthTokenURL = "https://devapisuperapp.cbe.com.et/superapp/parser/proxy/cbe-dev/sandbox/oauth-mb-cbebirr/oauth2/token?target=https%3A%2F%2Fapi-gw-uat-gateway-apic-nonprod.apps.cp4itest.cbe.local"
	sandboxCustomerURL   = "https://devapisuperapp.cbe.com.et/superapp/parser/proxy/cbe-dev/sandbox/cust_creation?target=https%3A%2F%2Fapi-gw-uat-gateway-apic-nonprod.apps.cp4itest.cbe.local"
	sandboxAccountURL    = "https://devapisuperapp.cbe.com.et/superapp/parser/proxy/cbe-dev/sandbox/acc-creation?target=https%3A%2F%2Fapi-gw-uat-gateway-apic-nonprod.apps.cp4itest.cbe.local"
	sandboxOAuthBody     = "grant_type=client_credentials&client_id=f1ceebd8d6d5b802dc7fd8332ab33603&client_secret=05ac01f2f134bfff2549669bc11bd6cc&scope=mb-cbebirr-scope"
	accountCategory      = "6501"
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

func stepLog(step int, total int, title string) {
	log.Printf("[Step %d/%d] %s", step, total, title)
}

func stepDone(detail string) {
	log.Printf("          -> %s", detail)
}

func stepFail(detail string) {
	log.Printf("          !! %s", detail)
}

func getAccessToken(ctx context.Context) (string, error) {
	stepLog(1, 8, "Fetching sandbox OAuth access token")
	stepDone(fmt.Sprintf("POST %s", sandboxOAuthTokenURL))

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

	stepDone(fmt.Sprintf("token received (type=%s, length=%d)", tokenResp.TokenType, len(tokenResp.AccessToken)))
	return tokenResp.AccessToken, nil
}

func (c *CoreAPI) checkIfUserExistsWithLogs(
	ctx context.Context,
	param core.UserExistsParam,
) (*core.UserExistsResult, error) {
	stepLog(3, 8, "CheckIfUserExists")
	stepDone("flow: PhoneLookup -> Fayda")
	stepDone(fmt.Sprintf("phone=%s nid=%s", param.PhoneNumber, param.NationalId))

	result, err := c.coreInterface.CheckIfUserExists(ctx, param)
	if err != nil {
		stepFail(fmt.Sprintf("unexpected error: %v", err))
		return nil, err
	}

	stepDone(fmt.Sprintf("success=%v messages=%v", result.Success, result.Messages))

	if result.PhoneLookupDetail != nil {
		stepLog(3, 8, "User found via PhoneLookup")
		if result.PhoneLookupDetail.Detail != nil {
			stepDone(fmt.Sprintf("customer_id=%s phone=%s",
				result.PhoneLookupDetail.Detail.CustomerID,
				result.PhoneLookupDetail.Detail.PhoneNumber))
		}
		return result, nil
	}

	if result.FaydaDetail != nil {
		stepLog(4, 8, "User found or blocked via Fayda")
		if result.FaydaDetail.Detail != nil {
			stepDone(fmt.Sprintf("flag=%s customer_id=%s name=%s",
				result.FaydaDetail.Detail.CustomerFlag,
				result.FaydaDetail.Detail.CustomerID,
				result.FaydaDetail.Detail.CustomerName))
		}
		return result, nil
	}

	if len(result.Messages) > 0 && !result.Success {
		stepFail(fmt.Sprintf("user existence check failed: %v", result.Messages))
		return result, nil
	}

	if result.Success {
		stepDone("user not found in phone lookup or fayda — proceeding to account creation")
	}

	return result, nil
}

func (c *CoreAPI) accountCreateWithLogs(
	ctx context.Context,
	param core.CreateCustomerParam,
	accountCreateURL, category string,
) (*core.CusteomerAccountCreationResponse, error) {
	stepLog(5, 8, "AccountCreate")
	stepDone("flow: CreateCustomer -> AML check -> AccountCreation")
	stepDone(fmt.Sprintf("POST customer_url=%s", param.Url))
	stepDone(fmt.Sprintf("POST account_url=%s category=%s", accountCreateURL, category))

	result, err := c.coreInterface.AccountCreate(ctx, param, accountCreateURL, category)
	if err != nil {
		stepFail(fmt.Sprintf("unexpected error: %v", err))
		return nil, err
	}

	if result.CustomerCreationDetail != nil {
		stepDone(fmt.Sprintf("customer creation success=%v messages=%v",
			result.CustomerCreationDetail.Success, result.CustomerCreationDetail.Messages))
		if result.CustomerCreationDetail.Detail != nil {
			stepDone(fmt.Sprintf("customer_number=%s mnemonic=%s",
				result.CustomerCreationDetail.Detail.CustomerNumber,
				result.CustomerCreationDetail.Detail.Menmonic))
		}
	}

	if result.AccountCreationDetail == nil && result.CustomerCreationDetail != nil && len(result.Messages) > 0 {
		stepFail(fmt.Sprintf("stopped before account creation: %v", result.Messages))
		return result, nil
	}

	if result.AccountCreationDetail != nil {
		stepLog(7, 8, "AccountCreation completed")
		stepDone(fmt.Sprintf("success=%v messages=%v",
			result.AccountCreationDetail.Success, result.AccountCreationDetail.Messages))
		if result.AccountCreationDetail.Detail != nil {
			stepDone(fmt.Sprintf("account_number=%s title=%s",
				result.AccountCreationDetail.Detail.AccountNumber,
				result.AccountCreationDetail.Detail.AccountTitle))
		}
	}

	if result.AccountCreationDetail != nil && result.AccountCreationDetail.Success {
		stepLog(8, 8, "Account create flow completed successfully")
	}

	return result, nil
}

func logUserExistsResult(result *core.UserExistsResult) {
	log.Println("========== User Exists Check ==========")
	log.Printf("success (clear to onboard): %v", result.Success)
	if len(result.Messages) > 0 {
		log.Printf("messages: %v", result.Messages)
	}
	if result.PhoneLookupDetail != nil && result.PhoneLookupDetail.Detail != nil {
		detail := result.PhoneLookupDetail.Detail
		log.Printf("phone lookup: customer_id=%s phone=%s email=%s",
			detail.CustomerID, detail.PhoneNumber, detail.Email)
	}
	if result.FaydaDetail != nil && result.FaydaDetail.Detail != nil {
		detail := result.FaydaDetail.Detail
		log.Printf("fayda: customer_id=%s name=%s flag=%s",
			detail.CustomerID, detail.CustomerName, detail.CustomerFlag)
	}
	log.Println("=======================================")
}

func logAccountCreateResult(result *core.CusteomerAccountCreationResponse) {
	log.Println("========== Account Create Result ==========")

	if len(result.Messages) > 0 {
		log.Printf("messages: %v", result.Messages)
	}

	if result.CustomerCreationDetail != nil {
		log.Printf("customer creation success: %v", result.CustomerCreationDetail.Success)
		if result.CustomerCreationDetail.Detail != nil {
			detail := result.CustomerCreationDetail.Detail
			log.Printf("  mnemonic: %s", detail.Menmonic)
			log.Printf("  full name: %s", detail.FullName)
			log.Printf("  phone: %s", detail.PhoneNumber)
			log.Printf("  national id: %s", detail.NationalId)
			log.Printf("  cocode: %s", detail.Cocode)
			log.Printf("  customer number: %s", detail.CustomerNumber)
		}
	} else {
		log.Println("customer creation detail: <nil>")
	}

	if result.AccountCreationDetail != nil {
		log.Printf("account creation success: %v", result.AccountCreationDetail.Success)
		if result.AccountCreationDetail.Detail != nil {
			detail := result.AccountCreationDetail.Detail
			log.Printf("  account number: %s", detail.AccountNumber)
			log.Printf("  customer: %s", detail.Customer)
			log.Printf("  account title: %s", detail.AccountTitle)
			log.Printf("  currency: %s", detail.Currency)
			log.Printf("  category: %s", detail.Category)
			log.Printf("  account officer: %s", detail.AccountOfficer)
			log.Printf("  cocode: %s", detail.CoCode)
		}
	} else {
		log.Println("account creation detail: <nil>")
	}

	log.Println("===========================================")
}

func main() {
	log.SetFlags(log.Ltime)
	log.Println("Starting account create flow")

	ctx, cancel := context.WithTimeout(context.Background(), 180*time.Second)
	defer cancel()

	accessToken, err := getAccessToken(ctx)
	if err != nil {
		log.Fatalf("failed to get access token: %v", err)
	}

	stepLog(2, 8, "Initializing core API client")
	stepDone(fmt.Sprintf("username=SUPERAPP core_url=%s", coreSOAPURL))
	stepDone(fmt.Sprintf("customer_url=%s", sandboxCustomerURL))
	stepDone(fmt.Sprintf("account_url=%s category=%s", sandboxAccountURL, accountCategory))

	calls := InitCoreAPICalls(
		"SUPERAPP",
		"123456",
		coreSOAPURL,
	)

	tinNumber := fmt.Sprintf("%016d", rand.Intn(10000000000000000))
	nationalID := "357253841014476138538353601641801488"
	email := fmt.Sprintf("sampletet%d@gmail.com", rand.Intn(10000000000000000))
	phoneNumber := "Y911706608"
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

	log.Printf("test payload: phone=%s email=%s national_id=%s tin=%s legal_id=%s",
		phoneNumber, email, nationalID, tinNumber, legalID)

	existsResult, err := calls.checkIfUserExistsWithLogs(ctx, core.UserExistsParam{
		PhoneNumber: phoneNumber,
		NationalId:  nationalID,
	})
	if err != nil {
		log.Fatalf("check if user exists failed: %v", err)
	}

	logUserExistsResult(existsResult)

	if !existsResult.Success {
		log.Println("stopping: user already exists or validation failed")
		return
	}

	accountResult, err := calls.accountCreateWithLogs(ctx, customer, sandboxAccountURL, accountCategory)
	if err != nil {
		log.Fatalf("account create failed: %v", err)
	}

	logAccountCreateResult(accountResult)

	if accountResult.AccountCreationDetail != nil && accountResult.AccountCreationDetail.Success {
		return
	}
	if len(accountResult.Messages) > 0 || accountResult.CustomerCreationDetail != nil {
		log.Println("flow completed with partial result or failure")
		return
	}
	log.Fatalf("account create rejected with no result details")
}
