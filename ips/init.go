package ips

import (
	"io"
	"net/http"
	"time"

	accountlookup "github.com/hugokessem/coreio/lib/ips/account_lookup"
	fundtransfer "github.com/hugokessem/coreio/lib/ips/fund_transfer"
	"github.com/hugokessem/coreio/utils"
)

type HttpPostWithRetryFunc func(url string, xmlBody string, config utils.Config) (*http.Response, error)
type AccountLookupParam = accountlookup.Params
type AccountLookupResult = accountlookup.AccountLookupResult
type FundTransferParam = fundtransfer.Params
type FundTransferResult = fundtransfer.FundTransferResult

type CBEIspAPIInterface interface {
	AccountLookup(param AccountLookupParam) (*AccountLookupResult, error)
	FundTransfer(param FundTransferParam) (*FundTransferResult, error)
}

type CBEIpsAPIHTTPHandler struct {
	Url string
}

// AccountLookupParam implements CBEIspAPIInterface.
func (c *CBEIpsAPIHTTPHandler) AccountLookup(param AccountLookupParam) (*AccountLookupResult, error) {
	config := utils.Config{
		MaxRetries: 6,
		Timeout:    30 * time.Second,
	}

	xmlRequest := accountlookup.NewAccountLookup(param)
	resp, err := utils.DoPostWithRetry(c.Url, xmlRequest, config)
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
func (c *CBEIpsAPIHTTPHandler) FundTransfer(param FundTransferParam) (*FundTransferResult, error) {
	config := utils.Config{
		MaxRetries: 6,
		Timeout:    30 * time.Second,
	}

	xmlRequest := fundtransfer.NewFundTransfer(param)
	resp, err := utils.DoPostWithRetry(c.Url, xmlRequest, config)
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

func NewCBEIpsAPI(url string) CBEIspAPIInterface {
	return &CBEIpsAPIHTTPHandler{
		Url: url,
	}
}
