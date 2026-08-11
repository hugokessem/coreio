package customersupport

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type CustomerSupportAPI struct {
	AuthURL     string
	BaseURL     string
	CustomerKey string
	AccessToken string
	Client      http.Client
}

type CustomerSupportCredential struct {
	AuthURL     string
	BaseURL     string
	CustomerKey string
	AccessToken string
}

func NewCustomerSupportAPI(param CustomerSupportCredential) CustomerSupportAPIInterface {
	return &CustomerSupportAPI{
		AuthURL:     strings.TrimRight(param.AuthURL, "/"),
		BaseURL:     strings.TrimRight(param.BaseURL, "/"),
		CustomerKey: param.CustomerKey,
		AccessToken: param.AccessToken,
		Client: http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: false,
					MinVersion:         tls.VersionTLS12,
				},
				DisableKeepAlives: true,
				IdleConnTimeout:   10 * time.Second,
			},
		},
	}
}

func (c *CustomerSupportAPI) SetAccessToken(token string) {
	c.AccessToken = token
}

func (c *CustomerSupportAPI) WorkerLogin(ctx context.Context, param WorkerLoginParam) (*WorkerLoginResult, error) {
	url := c.AuthURL + "/workerlogin_Prod"
	var result WorkerLoginResult
	if err := c.doJSON(ctx, http.MethodPost, url, param, c.CustomerKey, false, http.StatusOK, &result); err != nil {
		return nil, err
	}
	if result.AccessToken != "" {
		c.AccessToken = result.AccessToken
	}
	return &result, nil
}

func (c *CustomerSupportAPI) IssueProducts(ctx context.Context) (*IssueProductsResult, error) {
	url := c.BaseURL + "/issueProducts"
	var result IssueProductsResult
	if err := c.doJSON(ctx, http.MethodGet, url, nil, "", true, http.StatusOK, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *CustomerSupportAPI) IssuesList(ctx context.Context, param IssuesListParam) (*IssuesListResult, error) {
	url := c.BaseURL + "/cbeAdminIssuesList"
	var result IssuesListResult
	if err := c.doJSON(ctx, http.MethodPost, url, param, "", true, http.StatusOK, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *CustomerSupportAPI) ViewMessage(ctx context.Context, param ViewMessageParam) (*ViewMessageResult, error) {
	url := c.BaseURL + "/cbeCustomerviewMessage"
	var result ViewMessageResult
	if err := c.doJSON(ctx, http.MethodPost, url, param, "", true, http.StatusOK, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *CustomerSupportAPI) SendMessage(ctx context.Context, param SendMessageParam) (*SendMessageResult, error) {
	url := c.BaseURL + "/cbeAdminSendMessage"
	var result SendMessageResult
	if err := c.doJSON(ctx, http.MethodPost, url, param, "", true, http.StatusCreated, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *CustomerSupportAPI) ExportActivityLog(ctx context.Context) (*ExportActivityLogResult, error) {
	url := c.BaseURL + "/activity/exportlog"
	var result ExportActivityLogResult
	if err := c.doJSON(ctx, http.MethodGet, url, nil, "", true, http.StatusOK, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *CustomerSupportAPI) CreateIssueForCustomer(ctx context.Context, param CreateIssueParam) (*CreateIssueResult, error) {
	url := c.BaseURL + "/issuesCreateForCustomer"
	var result CreateIssueResult
	if err := c.doJSON(ctx, http.MethodPost, url, param, "", true, http.StatusCreated, &result); err != nil {
		return nil, err
	}
	if result.Status == "success" {
		result.Success = true
	}
	return &result, nil
}

func (c *CustomerSupportAPI) doJSON(
	ctx context.Context,
	method, url string,
	payload any,
	rawAuthorization string,
	useBearer bool,
	expectedStatus int,
	out any,
) error {
	var bodyReader io.Reader
	if payload != nil {
		body, err := json.Marshal(payload)
		if err != nil {
			return fmt.Errorf("failed to marshal request: %w", err)
		}
		bodyReader = bytes.NewReader(body)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	switch {
	case rawAuthorization != "":
		req.Header.Set("Authorization", rawAuthorization)
	case useBearer:
		if c.AccessToken == "" {
			return fmt.Errorf("access token is required")
		}
		req.Header.Set("Authorization", "Bearer "+c.AccessToken)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to call customer support api: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != expectedStatus {
		return fmt.Errorf("customer support api failed (%d): %s", resp.StatusCode, body)
	}

	if out == nil {
		return nil
	}

	if err := json.Unmarshal(body, out); err != nil {
		return fmt.Errorf("failed to decode response body: %w", err)
	}
	return nil
}
