package ips

import (
	"io"
	"net/http"
	"time"

	"github.com/hugokessem/coreio/ips/internal"
	accountlookup "github.com/hugokessem/coreio/lib/ips/account_lookup"
	fundtransfer "github.com/hugokessem/coreio/lib/ips/fund_transfer"
	statuscheck "github.com/hugokessem/coreio/lib/ips/status_check"
	"github.com/hugokessem/coreio/utils"
)

type HttpPostWithRetryFunc func(url string, xmlBody string, config utils.Config) (*http.Response, error)
type AccountLookupParam = accountlookup.Params
type AccountLookupResult = accountlookup.AccountLookupResult
type FundTransferParam = fundtransfer.Params
type FundTransferResult = fundtransfer.FundTransferResult

type StatusCheckParam = statuscheck.Params
type StatusCheckResult = statuscheck.PaymentStatusResult

type IPSCoreAPIInterface interface {
	AccountLookup(param AccountLookupParam) (*AccountLookupResult, error)
	FundTransfer(param FundTransferParam) (*FundTransferResult, error)
	StatusCheck(param StatusCheckParam) (*StatusCheckResult, error)
}

type IPSCredentials struct {
	Username        string
	Password        string
	GrantType       string
	JwtAssertion    string
	MBAuthorization string
	Authorization   string
	Url             string
}

type IPSCoreAPI struct {
	config *internal.Config
}

func (c *IPSCoreAPI) StatusCheck(param StatusCheckParam) (*StatusCheckResult, error) {
	config := utils.Config{
		MaxRetries: 6,
		Timeout:    30 * time.Second,
	}

	xmlRequest := statuscheck.NewStatusCheck(param)
	headers := map[string]string{
		"Content-Type":     "application/xml",
		"username":         c.config.Username,
		"password":         c.config.Password,
		"grant_type":       c.config.GrantType,
		"Jwt_Assertion":    c.config.JwtAssertion,
		"MB_authorization": c.config.MBAuthorization,
		"Authorization":    c.config.Authorization,
	}
	resp, err := utils.DoPostWithRetry(c.config.Url, xmlRequest, config, headers)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseDate, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result, err := statuscheck.ParsePaymentStatusSOAP(string(responseDate))
	if err != nil {
		return nil, err
	}

	return result, nil
}

// AccountLookupParam implements CBEIspAPIInterface.
func (c *IPSCoreAPI) AccountLookup(param AccountLookupParam) (*AccountLookupResult, error) {
	config := utils.Config{
		MaxRetries: 6,
		Timeout:    30 * time.Second,
	}

	xmlRequest := accountlookup.NewAccountLookup(param)
	headers := map[string]string{
		"Content-Type":     "application/xml",
		"username":         c.config.Username,
		"password":         c.config.Password,
		"grant_type":       c.config.GrantType,
		"Jwt_Assertion":    c.config.JwtAssertion,
		"MB_authorization": c.config.MBAuthorization,
		"Authorization":    c.config.Authorization,
	}

	resp, err := utils.DoPostWithRetry(c.config.Url, xmlRequest, config, headers)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseDate, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result, err := accountlookup.ParseAccountLookupSOAP(string(responseDate))
	if err != nil {
		return nil, err
	}

	return result, nil
}

// FundTransfer implements CBEIspAPIInterface.
func (c *IPSCoreAPI) FundTransfer(param FundTransferParam) (*FundTransferResult, error) {
	config := utils.Config{
		MaxRetries: 6,
		Timeout:    30 * time.Second,
	}

	xmlRequest := fundtransfer.NewFundTransfer(param)
	headers := map[string]string{
		"Content-Type":     "application/xml",
		"username":         c.config.Username,
		"password":         c.config.Password,
		"grant_type":       c.config.GrantType,
		"Jwt_Assertion":    c.config.JwtAssertion,
		"MB_authorization": c.config.MBAuthorization,
		"Authorization":    c.config.Authorization,
	}

	resp, err := utils.DoPostWithRetry(c.config.Url, xmlRequest, config, headers)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseDate, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result, err := fundtransfer.ParseFundTransferSOAP(string(responseDate))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func NewCBEIpsAPI(param IPSCredentials) IPSCoreAPIInterface {
	config := internal.SetConfig(
		param.Username,
		param.Password,
		param.GrantType,
		param.JwtAssertion,
		param.MBAuthorization,
		param.Authorization,
		param.Url,
	)
	return &IPSCoreAPI{
		config: config,
	}
}
