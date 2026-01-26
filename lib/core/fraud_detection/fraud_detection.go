package frauddetection

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type FraudAPI struct {
	Authorization string
	ForwardHost   string
	Url           string
	Client        http.Client
}

type FraudAPIPayload struct {
	TranasctionID              string `json:"transaction_id"`
	AccountID                  string `json:"account_id"`
	CustomerName               string `json:"customer_name"`
	CustomerPhoneMobileSMS     string `json:"customer_phone_mobile_sms"`
	BeneficiaryAccountID       string `json:"beneficiary_account_id"`
	BeneficiaryName            string `json:"beneficiary_name"`
	AccountCategory            string `json:"account_category"`
	AccountCurrency            string `json:"account_currency"`
	TransactionConvertedAmount string `json:"transaction_converted_amount"`
	TransactionType            string `json:"transaction_type"`
	SourceUser                 string `json:"source_user"`
	ChangeInPhoneEmail         string `json:"change_in_phone_email"`
	TransactionTimestamp       string `json:"transaction_timestamp"`
	ChangeInPIN                string `json:"change_in_pin"`
	ChangeInPassword           string `json:"change_in_password"`
	ChangeInDevice             string `json:"change_in_device"`
}

type FraudAPIResponse struct {
	TransactionID string `json:"transaction_id"`
	Violation     string `json:"violation"`
}

func NewFraudAPI(authorization string, forwardHost string, url string) FraudAPIInterface {
	return FraudAPI{
		Authorization: authorization,
		ForwardHost:   forwardHost,
		Url:           url,
		Client: http.Client{
			Timeout: 10 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: false,
					MinVersion:         tls.VersionTLS13,
				},
				DisableKeepAlives: true,
				IdleConnTimeout:   10 * time.Second,
			},
		},
	}
}

func NewFraudAPIPayload(param FraudAPIPayload) string {
	return fmt.Sprintf(`{
	"transaction":{
		"transaction_id": "%s",
		"account_id": "%s",
		"customer_name": "%s",
		"customer_phone_mobile_sms": "%s",
		"beneficiary_account_id": "%s",
		"beneficiary_name": "%s",
		"account_category": "%s",
		"account_currency": "%s",
		"transaction_converted_amount": "%s",
		"transaction_type": "%s",
		"source_user": "%s",
		"change_in_phone_email": "%s",
		"transaction_timestamp": "%s",
		"change_in_pin": "%s",
		"change_in_password": "%s",
		"change_in_device": "%s"
		}
	}`, param.TranasctionID, param.AccountID, param.CustomerName, param.CustomerPhoneMobileSMS,
		param.BeneficiaryAccountID, param.BeneficiaryName, param.AccountCategory, param.AccountCurrency,
		param.TransactionConvertedAmount, param.TransactionType, param.SourceUser, param.ChangeInPhoneEmail,
		param.TransactionTimestamp, param.ChangeInPIN, param.ChangeInPassword, param.ChangeInDevice)
}

func (f FraudAPI) Call(param FraudAPIPayload) (*FraudAPIResponse, error) {
	payload := NewFraudAPIPayload(param)

	method := "POST"
	url := f.Url

	req, err := http.NewRequest(method, url, strings.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", f.Authorization)
	req.Header.Add("x-forward-host", f.ForwardHost)

	resp, err := f.Client.Do(req)

	if err != nil {
		return nil, errors.New("failed to call fraud api")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("failed to read response body")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fraud api failed: %s", body)
	}

	var result FraudAPIResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("failed to decode response body")
	}

	fmt.Println("Response Payload:", string(body))
	return &result, nil
}
