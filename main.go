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

func (c *CoreAPI) AccountCreateWithLogs(
	ctx context.Context,
	param core.CreateCustomerParam,
	accountCreateURL, category string,
) (*core.CusteomerAccountCreationResponse, error) {
	stepLog(3, 8, "Account create orchestration")
	stepDone("flow: PhoneLookup -> Fayda -> CreateCustomer -> AccountCreation")
	stepDone(fmt.Sprintf("phone=%s nid=%s", param.PhoneNumber, param.NationalId))

	result, err := c.coreInterface.AccountCreate(ctx, param, accountCreateURL, category)
	if err != nil {
		stepFail(fmt.Sprintf("unexpected error: %v", err))
		return nil, err
	}

	if len(result.Messages) > 0 && result.PhoneLookupDetail == nil && result.FaydaDetail == nil && result.CustomerCreationDetail == nil {
		stepFail(fmt.Sprintf("flow failed during initial validation: %v", result.Messages))
		return result, nil
	}

	if result.PhoneLookupDetail != nil {
		stepLog(3, 8, "Stopped at PhoneLookup (customer already exists)")
		stepDone(fmt.Sprintf("success=%v messages=%v", result.PhoneLookupDetail.Success, result.PhoneLookupDetail.Message))
		if result.PhoneLookupDetail.Detail != nil {
			stepDone(fmt.Sprintf("customer_id=%s phone=%s",
				result.PhoneLookupDetail.Detail.CustomerID, result.PhoneLookupDetail.Detail.PhoneNumber))
		}
		return result, nil
	}

	if result.FaydaDetail != nil {
		stepLog(4, 8, "Stopped at Fayda (customer found via NID)")
		stepDone(fmt.Sprintf("success=%v messages=%v", result.FaydaDetail.Success, result.FaydaDetail.Message))
		if result.FaydaDetail.Detail != nil {
			stepDone(fmt.Sprintf("flag=%s customer_id=%s name=%s",
				result.FaydaDetail.Detail.CustomerFlag,
				result.FaydaDetail.Detail.CustomerID,
				result.FaydaDetail.Detail.CustomerName))
		}
		return result, nil
	}

	if result.CustomerCreationDetail != nil {
		stepLog(5, 8, "CreateCustomer completed")
		stepDone(fmt.Sprintf("success=%v messages=%v", result.CustomerCreationDetail.Success, result.CustomerCreationDetail.Messages))
		if result.CustomerCreationDetail.Detail != nil {
			stepDone(fmt.Sprintf("customer_number=%s mnemonic=%s",
				result.CustomerCreationDetail.Detail.CustomerNumber,
				result.CustomerCreationDetail.Detail.Menmonic))
		}
	}

	if result.AccountCreationDetail == nil && result.CustomerCreationDetail != nil {
		if len(result.Messages) > 0 {
			stepFail(fmt.Sprintf("stopped before account creation: %v", result.Messages))
		}
		return result, nil
	}

	if result.AccountCreationDetail != nil {
		stepLog(7, 8, "AccountCreation completed")
		stepDone(fmt.Sprintf("success=%v messages=%v", result.AccountCreationDetail.Success, result.AccountCreationDetail.Messages))
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

func logFinalResult(result *core.CusteomerAccountCreationResponse) {
	log.Println("========== Final Result ==========")

	if len(result.Messages) > 0 {
		log.Printf("messages: %v", result.Messages)
	}

	if result.PhoneLookupDetail != nil {
		log.Println("phone lookup detail:")
		if result.PhoneLookupDetail.Detail != nil {
			detail := result.PhoneLookupDetail.Detail
			log.Printf("  customer id: %s", detail.CustomerID)
			log.Printf("  phone: %s", detail.PhoneNumber)
			log.Printf("  email: %s", detail.Email)
			log.Printf("  full name: %s", detail.FullName)
		}
	}

	if result.FaydaDetail != nil {
		log.Println("fayda detail:")
		if result.FaydaDetail.Detail != nil {
			detail := result.FaydaDetail.Detail
			log.Printf("  customer id: %s", detail.CustomerID)
			log.Printf("  name: %s", detail.CustomerName)
			log.Printf("  flag: %s", detail.CustomerFlag)
		}
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

	log.Println("==================================")
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
	nationalID := fmt.Sprintf("%016d", rand.Intn(10000000000000000))
	email := fmt.Sprintf("sampletet%d@gmail.com", rand.Intn(10000000000000000))
	phoneNumber := fmt.Sprintf("+25191%06d", rand.Intn(1000000))
	legalID := fmt.Sprintf("%016d", rand.Intn(10000000000000000))

	customer := core.CreateCustomerParam{
		FirstName:          "KETEM",
		MiddleName:         "HAILU",
		LastName:           "TAYE",
		PhoneNumber:        "Y911706608",
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
		NationalId:         "357253841014476138538353601641801888",
		Url:                sandboxCustomerURL,
		Header: map[string]string{
			"Authorization": "Bearer " + accessToken,
		},
	}

	log.Printf("generated test payload: phone=%s email=%s national_id=%s tin=%s legal_id=%s",
		phoneNumber, email, nationalID, tinNumber, legalID)

	result, err := calls.AccountCreateWithLogs(ctx, customer, sandboxAccountURL, accountCategory)
	if err != nil {
		log.Fatalf("account create failed: %v", err)
	}

	logFinalResult(result)

	if result.AccountCreationDetail != nil && result.AccountCreationDetail.Success {
		return
	}
	if result.PhoneLookupDetail != nil || result.FaydaDetail != nil {
		log.Println("flow completed with early exit (customer already exists)")
		return
	}
	if len(result.Messages) > 0 {
		log.Println("flow completed with partial result or failure")
		return
	}
	log.Fatalf("account create rejected with no result details")
}
