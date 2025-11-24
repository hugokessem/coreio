package core

import (
	"fmt"
	"io"
	"time"

	"github.com/hugokessem/coreio/core/internal"
	accountlookup "github.com/hugokessem/coreio/lib/core/account_lookup"
	fundtransfer "github.com/hugokessem/coreio/lib/core/fund_transfer"
	lockedamountcreate "github.com/hugokessem/coreio/lib/core/locked_amount/locked_amount_create"
	lockedamountft "github.com/hugokessem/coreio/lib/core/locked_amount/locked_amount_ft"
	lockedamountlist "github.com/hugokessem/coreio/lib/core/locked_amount/locked_amount_list"
	lockedamountrelease "github.com/hugokessem/coreio/lib/core/locked_amount/locked_amount_release"
	ministatementbydaterange "github.com/hugokessem/coreio/lib/core/mini_statement/mini_statement_by_date_range"
	ministatementbylimit "github.com/hugokessem/coreio/lib/core/mini_statement/mini_statement_by_limit"
	revertfundtransfer "github.com/hugokessem/coreio/lib/core/revert_fund_transfer"
	standingordercancel "github.com/hugokessem/coreio/lib/core/standing_order/standing_order_cancel"
	standingordercreate "github.com/hugokessem/coreio/lib/core/standing_order/standing_order_create"
	standingorderlist "github.com/hugokessem/coreio/lib/core/standing_order/standing_order_list"
	standingorderupdate "github.com/hugokessem/coreio/lib/core/standing_order/standing_order_update"
	"github.com/hugokessem/coreio/utils"
)

type AccountLookupParam = accountlookup.AccountLookupParam
type AccountLookupResult = accountlookup.AccountLookupResult
type FundTransferParam = fundtransfer.FundTransferParam
type FundTransferResult = fundtransfer.FundTransferResult
type RevertFundTransferParam = revertfundtransfer.RevertFundTransferParams
type RevertFundTransferResult = revertfundtransfer.RevertFundTransferResult

type LockedAmountFTParam = lockedamountft.LockedAmountFTParams
type LockedAmountFTResult = lockedamountft.LockedAmountFTResult
type ListLockedAmountParam = lockedamountlist.ListLockedAmountParam
type ListLockedAmountResult = lockedamountlist.ListLockedAmountResult
type CreateLockedAmountParam = lockedamountcreate.CreateLockedAmountParam
type CreateLockedAmountResult = lockedamountcreate.CreateLockedAmountResult
type ReleaseLockedAmountParam = lockedamountrelease.ReleaseLockedAmountParam
type ReleaseLockedAmountResult = lockedamountrelease.ReleaseAccountLockedResult

type CreateStandingOrderParam = standingordercreate.CreateStandingOrderParams
type CreateStandingOrderResult = standingordercreate.StandingOrderResult
type ListStandingOrderParam = standingorderlist.ListStandingOrderParams
type ListStandingOrderResult = standingorderlist.ListStandingOrderResult
type UpdateStandingOrderParam = standingorderupdate.UpdateStandingOrderParam
type UpdateStandingOrderResult = standingorderupdate.UpdateStandingOrderResult
type CancleStandingOrderParam = standingordercancel.CancelStandingOrderParams
type CancelStandingOrderResult = standingordercancel.CancelStandingOrderResult

type MiniStatementByLimitParams = ministatementbylimit.MiniStatementByLimitParams
type MiniStatementByLimitResult = ministatementbylimit.MiniStatementByLimitResult
type MiniStatementByDateRangeParam = ministatementbydaterange.MiniStatementByDateRangeParam
type MiniStatementByDateRangeResult = ministatementbydaterange.MiniStatementByDateRangeResult

type CBECoreAPIInterface interface {
	FundTransfer(param FundTransferParam) (*FundTransferResult, error)
	RevertFundTransfer(param RevertFundTransferParam) (*RevertFundTransferResult, error)
	AccountLookup(param AccountLookupParam) (*AccountLookupResult, error)

	LockedAmountFT(param LockedAmountFTParam) (*LockedAmountFTResult, error)
	ListLockedAmount(param ListLockedAmountParam) (*ListLockedAmountResult, error)
	CreateLockedAmount(param CreateLockedAmountParam) (*CreateLockedAmountResult, error)
	ReleaseLockedAmount(param ReleaseLockedAmountParam) (*ReleaseLockedAmountResult, error)

	ListStandingOrder(param ListStandingOrderParam) (*ListStandingOrderResult, error)
	UpdateStandingOrder(param UpdateStandingOrderParam) (*UpdateStandingOrderResult, error)
	CreateStandingOrder(param CreateStandingOrderParam) (*CreateStandingOrderResult, error)
	CancleStandingOrder(param CancleStandingOrderParam) (*CancelStandingOrderResult, error)
	MiniStatementByLimit(param MiniStatementByLimitParams) (*MiniStatementByLimitResult, error)
	MiniStatementByDateRange(param MiniStatementByDateRangeParam) (*MiniStatementByDateRangeResult, error)
}

type CBECoreCredential struct {
	Username string
	Password string
	Url      string
}

type CBECoreAPI struct {
	config *internal.Config
}

func (c *CBECoreAPI) RevertFundTransfer(param RevertFundTransferParam) (*RevertFundTransferResult, error) {
	params := revertfundtransfer.Params{
		Username:      c.config.Username,
		Password:      c.config.Password,
		TransactionID: param.TransactionID,
	}

	xmlRequest := revertfundtransfer.NewRevertFundTransfer(params)
	headers := map[string]string{
		"Content-Type": "application/xml",
	}
	resp, err := utils.DoPostWithRetry(c.config.Url, xmlRequest, utils.Config{
		Timeout:    30 * time.Second,
		MaxRetries: 6,
	}, headers)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result, err := revertfundtransfer.ParseRevertFundTransferSOAP(string(responseData))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *CBECoreAPI) AccountLookup(param AccountLookupParam) (*AccountLookupResult, error) {
	params := accountlookup.Params{
		Username:      c.config.Username,
		Password:      c.config.Password,
		AccountNumber: param.AccountNumber,
	}

	xmlRequest := accountlookup.NewAccountLookup(params)
	headers := map[string]string{
		"Content-Type": "application/xml",
	}
	resp, err := utils.DoPostWithRetry(c.config.Url, xmlRequest, utils.Config{
		Timeout:    30 * time.Second,
		MaxRetries: 6,
	}, headers)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result, err := accountlookup.ParseAccountLookupSOAP(string(responseData))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *CBECoreAPI) FundTransfer(param FundTransferParam) (*FundTransferResult, error) {
	params := fundtransfer.Params{
		Username:            c.config.Username,
		Password:            c.config.Password,
		DebitAccountNumber:  param.DebitAccountNumber,
		DebitCurrency:       param.DebitCurrency,
		CreditAccountNumber: param.CreditAccountNumber,
		CreditCurrency:      param.CreditCurrency,
		DebitReference:      param.DebitReference,
		CreditReference:     param.CreditReference,
		DebitAmount:         param.DebitAmount,
		TransationID:        param.TransationID,
		PaymentDetail:       param.PaymentDetail,
	}

	xmlRequest := fundtransfer.NewFundTransfer(params)
	headers := map[string]string{
		"Content-Type": "application/xml",
	}
	resp, err := utils.DoPostWithRetry(c.config.Url, xmlRequest, utils.Config{
		Timeout:    30 * time.Second,
		MaxRetries: 6,
	}, headers)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result, err := fundtransfer.ParseFundTransferSOAP(string(responseData))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *CBECoreAPI) ListLockedAmount(param ListLockedAmountParam) (*ListLockedAmountResult, error) {
	params := lockedamountlist.Params{
		Username:      c.config.Username,
		Password:      c.config.Password,
		AccountNumber: param.AccountNumber,
	}

	xmlRequest := lockedamountlist.NewListLockedAmount(params)
	headers := map[string]string{
		"Content-Type": "application/xml",
	}
	resp, err := utils.DoPostWithRetry(c.config.Url, xmlRequest, utils.Config{
		Timeout:    30 * time.Second,
		MaxRetries: 6,
	}, headers)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result, err := lockedamountlist.ParseListLockedAmountSOAP(string(responseData))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *CBECoreAPI) CreateLockedAmount(param CreateLockedAmountParam) (*CreateLockedAmountResult, error) {
	params := lockedamountcreate.Params{
		Username:      c.config.Username,
		Password:      c.config.Password,
		AccountNumber: param.AccountNumber,
		Description:   param.Description,
		From:          param.From,
		To:            param.To,
		LockedAmount:  param.LockedAmount,
	}

	xmlRequest := lockedamountcreate.NewCreateLockedAmount(params)
	headers := map[string]string{
		"Content-Type": "text/xml; charset=utf-8",
	}
	resp, err := utils.DoPostWithRetry(c.config.Url, xmlRequest, utils.Config{
		Timeout:    30 * time.Second,
		MaxRetries: 6,
	}, headers)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result, err := lockedamountcreate.ParseCreateLockedAmountSOAP(string(responseData))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *CBECoreAPI) LockedAmountFT(param LockedAmountFTParam) (*LockedAmountFTResult, error) {
	params := lockedamountft.Params{
		Username:            c.config.Username,
		Password:            c.config.Password,
		CreditCurrent:       param.CreditCurrent,
		CreditAccountNumber: param.CreditAccountNumber,
		CrediterReference:   param.CrediterReference,
		DebitAmount:         param.DebitAmount,
		DebitAccountNumber:  param.DebitAccountNumber,
		DebitCurrency:       param.DebitCurrency,
		DebiterReference:    param.DebiterReference,
		ClientReference:     param.ClientReference,
		LockID:              param.LockID,
	}

	xmlRequest := lockedamountft.NewLockedAmountFt(params)
	headers := map[string]string{
		"Content-Type": "text/xml; charset=utf-8",
	}
	resp, err := utils.DoPostWithRetry(c.config.Url, xmlRequest, utils.Config{
		Timeout:    30 * time.Second,
		MaxRetries: 6,
	}, headers)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result, err := lockedamountft.ParseLockedAmountFTSOAP(string(responseData))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *CBECoreAPI) ReleaseLockedAmount(param ReleaseLockedAmountParam) (*ReleaseLockedAmountResult, error) {
	params := lockedamountrelease.Params{
		Username:      c.config.Username,
		Password:      c.config.Password,
		TransactionID: param.TransactionID,
	}

	xmlRequest := lockedamountrelease.NewReleaseLockedAmount(params)
	headers := map[string]string{
		"Content-Type": "text/xml; charset=utf-8",
	}
	resp, err := utils.DoPostWithRetry(c.config.Url, xmlRequest, utils.Config{
		Timeout:    30 * time.Second,
		MaxRetries: 6,
	}, headers)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result, err := lockedamountrelease.ParseCancleLockedAmountSOAP(string(responseData))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *CBECoreAPI) CancleStandingOrder(param CancleStandingOrderParam) (*CancelStandingOrderResult, error) {
	params := standingordercancel.Params{
		Username:      c.config.Username,
		Password:      c.config.Password,
		AccountNumber: param.AccountNumber,
		OrderId:       param.OrderId,
	}

	xmlRequest := standingordercancel.NewCancleStandingOrder(params)
	headers := map[string]string{
		"Content-Type": "text/xml; charset=utf-8",
	}
	resp, err := utils.DoPostWithRetry(c.config.Url, xmlRequest, utils.Config{
		Timeout:    30 * time.Second,
		MaxRetries: 6,
	}, headers)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result, err := standingordercancel.ParseCancelStandingOrderSOAP(string(responseData))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *CBECoreAPI) UpdateStandingOrder(param UpdateStandingOrderParam) (*UpdateStandingOrderResult, error) {
	params := standingorderupdate.Params{
		Username:            c.config.Username,
		Password:            c.config.Password,
		Amount:              param.Amount,
		OrderId:             param.OrderId,
		Currency:            param.Currency,
		Frequency:           param.Frequency,
		CurrentDate:         param.CurrentDate,
		PaymentDetail:       param.PaymentDetail,
		DebitAccountNumber:  param.DebitAccountNumber,
		CreditAccountNumber: param.CreditAccountNumber,
	}

	xmlRequest := standingorderupdate.NewUpdateStandingOrder(params)
	headers := map[string]string{
		"Content-Type": "text/xml; charset=utf-8",
	}
	resp, err := utils.DoPostWithRetry(c.config.Url, xmlRequest, utils.Config{
		Timeout:    30 * time.Second,
		MaxRetries: 6,
	}, headers)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result, err := standingorderupdate.ParseUpdateStandingOrderSOAP(string(responseData))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *CBECoreAPI) CreateStandingOrder(param CreateStandingOrderParam) (*CreateStandingOrderResult, error) {
	params := standingordercreate.Params{
		Username:            c.config.Username,
		Password:            c.config.Password,
		DebitAccountNumber:  param.DebitAccountNumber,
		CreditAccountNumber: param.CreditAccountNumber,
		CurrentDate:         param.CurrentDate,
		Frequency:           param.Frequency,
		Currency:            param.Currency,
		PaymentDetail:       param.PaymentDetail,
		Amount:              param.Amount,
	}

	xmlRequest := standingordercreate.NewCreateStandingOrder(params)
	headers := map[string]string{
		"Content-Type": "text/xml; charset=utf-8",
	}
	resp, err := utils.DoPostWithRetry(c.config.Url, xmlRequest, utils.Config{
		Timeout:    30 * time.Second,
		MaxRetries: 6,
	}, headers)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result, err := standingordercreate.ParseCreateStandingOrderSOAP(string(responseData))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *CBECoreAPI) ListStandingOrder(param ListStandingOrderParam) (*ListStandingOrderResult, error) {
	params := standingorderlist.Params{
		Username:      c.config.Username,
		Password:      c.config.Password,
		AccountNumber: param.AccountNumber,
	}

	xmlRequest := standingorderlist.NewListStandingOrder(params)
	headers := map[string]string{
		"Content-Type": "text/xml; charset=utf-8",
	}
	resp, err := utils.DoPostWithRetry(c.config.Url, xmlRequest, utils.Config{
		Timeout:    30 * time.Second,
		MaxRetries: 6,
	}, headers)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result, err := standingorderlist.ParseListStandingOrderSOAP(string(responseData))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *CBECoreAPI) MiniStatementByLimit(param MiniStatementByLimitParams) (*MiniStatementByLimitResult, error) {
	params := ministatementbylimit.Params{
		Username:            c.config.Username,
		Password:            c.config.Password,
		AccountNumber:       param.AccountNumber,
		NumberOfTransaction: param.NumberOfTransaction,
	}

	xmlRequest := ministatementbylimit.NewMiniStatementByLimit(params)
	headers := map[string]string{
		"Content-Type": "text/xml; charset=utf-8",
	}
	resp, err := utils.DoPostWithRetry(c.config.Url, xmlRequest, utils.Config{
		Timeout:    30 * time.Second,
		MaxRetries: 6,
	}, headers)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result, err := ministatementbylimit.ParseMiniStatementByLimitSOAP(string(responseData))
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *CBECoreAPI) MiniStatementByDateRange(param MiniStatementByDateRangeParam) (*MiniStatementByDateRangeResult, error) {
	params := ministatementbydaterange.Params{
		Username:      c.config.Username,
		Password:      c.config.Password,
		AccountNumber: param.AccountNumber,
		From:          param.From,
		To:            param.To,
	}

	xmlRequest := ministatementbydaterange.NewMiniStatementByDateRange(params)
	fmt.Println("xmlRequest", xmlRequest)
	headers := map[string]string{
		"Content-Type": "text/xml; charset=utf-8",
	}

	resp, err := utils.DoPostWithRetry(c.config.Url, xmlRequest, utils.Config{
		Timeout:    30 * time.Second,
		MaxRetries: 6,
	}, headers)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result, err := ministatementbydaterange.ParseMiniStatementByDateRangeSOAP(string(responseData))
	fmt.Println("result", result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func NewCBECoreAPI(param CBECoreCredential) CBECoreAPIInterface {
	config := internal.SetConfig(param.Username, param.Password, param.Url)
	return &CBECoreAPI{
		config: config,
	}
}
