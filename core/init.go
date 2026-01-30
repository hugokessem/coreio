package core

import (
	"fmt"
	"io"
	"time"

	"github.com/hugokessem/coreio/core/internal"
	accountcreation "github.com/hugokessem/coreio/lib/core/account/acccount_creation"
	accountlist "github.com/hugokessem/coreio/lib/core/account/account_list"
	accountlookup "github.com/hugokessem/coreio/lib/core/account/account_lookup"
	cardreplace "github.com/hugokessem/coreio/lib/core/card/card_replace"
	cardrequest "github.com/hugokessem/coreio/lib/core/card/card_request"
	frauddetection "github.com/hugokessem/coreio/lib/core/fraud_detection"
	splitpayment "github.com/hugokessem/coreio/lib/core/split_payment"

	customerlimitamendbycif "github.com/hugokessem/coreio/lib/core/customer/customer_limit_amend_by_cif"
	customerlimitfetchbycif "github.com/hugokessem/coreio/lib/core/customer/customer_limit_fetch_by_cif"
	customerlimitfetchbyservice "github.com/hugokessem/coreio/lib/core/customer/customer_limit_fetch_by_service"
	customerlookup "github.com/hugokessem/coreio/lib/core/customer/customer_lookup"
	exchangerate "github.com/hugokessem/coreio/lib/core/exchange_rate"
	fundtransfer "github.com/hugokessem/coreio/lib/core/fund_transfer/fund_transfer"
	fundtransfercheck "github.com/hugokessem/coreio/lib/core/fund_transfer/fund_transfer_check"
	lockedamountcreate "github.com/hugokessem/coreio/lib/core/locked_amount/locked_amount_create"
	lockedamountft "github.com/hugokessem/coreio/lib/core/locked_amount/locked_amount_ft"
	lockedamountlist "github.com/hugokessem/coreio/lib/core/locked_amount/locked_amount_list"
	lockedamountrelease "github.com/hugokessem/coreio/lib/core/locked_amount/locked_amount_release"
	ministatementbydaterange "github.com/hugokessem/coreio/lib/core/mini_statement/mini_statement_by_date_range"
	ministatementbylimit "github.com/hugokessem/coreio/lib/core/mini_statement/mini_statement_by_limit"
	phonelookup "github.com/hugokessem/coreio/lib/core/phone_lookup"
	revertfundtransfer "github.com/hugokessem/coreio/lib/core/revert_fund_transfer"
	servicelimit "github.com/hugokessem/coreio/lib/core/service/service_limit"
	standingordercancel "github.com/hugokessem/coreio/lib/core/standing_order/standing_order_cancel"
	standingordercreate "github.com/hugokessem/coreio/lib/core/standing_order/standing_order_create"
	standingorderlist "github.com/hugokessem/coreio/lib/core/standing_order/standing_order_list"
	standingorderlisthistory "github.com/hugokessem/coreio/lib/core/standing_order/standing_order_list_history"
	standingorderupdate "github.com/hugokessem/coreio/lib/core/standing_order/standing_order_update"
	"github.com/hugokessem/coreio/utils"
)

type ExchangeRatesResult = exchangerate.ExchangeRateResult
type AccountLookupParam = accountlookup.AccountLookupParam
type AccountLookupResult = accountlookup.AccountLookupResult
type AccountListParam = accountlist.AccountListParams
type AccountListResult = accountlist.AccountListResult
type ServiceLimitParam = servicelimit.ServiceLimitParam
type ServiceLimitResult = servicelimit.ServiceLimitResult

type SplitPaymentParam = splitpayment.SplitPaymentParam
type SplitPaymentResult = splitpayment.SplitPaymentResult

type AccountCreationParam = accountcreation.AccountCreationParams
type AccountCreationResult = accountcreation.AccountCreationResult

type FundTransferParam = fundtransfer.FundTransferParam
type FundTransferResult = fundtransfer.FundTransferResult
type FundTransferCheckParam = fundtransfercheck.FundTransferCheckParams
type FundTransferCheckResult = fundtransfercheck.FundTransferCheckResult
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
type ListStandingOrderHistoryParam = standingorderlisthistory.ListStandingOrderHistoryParams
type ListStandingOrderHistoryResult = standingorderlisthistory.StandingOrderListHistoryResult
type UpdateStandingOrderParam = standingorderupdate.UpdateStandingOrderParam
type UpdateStandingOrderResult = standingorderupdate.UpdateStandingOrderResult
type CancleStandingOrderParam = standingordercancel.CancelStandingOrderParams
type CancelStandingOrderResult = standingordercancel.CancelStandingOrderResult

type MiniStatementByLimitParams = ministatementbylimit.MiniStatementByLimitParams
type MiniStatementByLimitResult = ministatementbylimit.MiniStatementByLimitResult
type MiniStatementByDateRangeParam = ministatementbydaterange.MiniStatementByDateRangeParam
type MiniStatementByDateRangeResult = ministatementbydaterange.MiniStatementByDateRangeResult

type CustomerLookupParam = customerlookup.CustomerLookupParam
type CustomerLookupResult = customerlookup.CustomerLookupResult

type PhoneLookupParam = phonelookup.PhoneLookupParam
type PhoneLookupResult = phonelookup.PhoneLookupResult

type CardReplaceParam = cardreplace.CardReplaceParam
type CardReplaceResult = cardreplace.CardReplaceResult
type CardRequestParam = cardrequest.CardRequestParam
type CardRequestResult = cardrequest.CardRequestResult

type CustomerLimitFetchByCIFResult = customerlimitfetchbycif.CustomerLimitViewResult
type CustomerLimitFetchByCIFParam = customerlimitfetchbycif.CustomerLimitFetchByCIFParam
type CustomerLimitFetchByServiceResult = customerlimitfetchbyservice.CustomerLimitFetchResult
type CustomerLimitFetchByServiceParam = customerlimitfetchbyservice.CustomerLimitFetchByServiceParam
type CustomerLimitAmendByCIFResult = customerlimitamendbycif.CustomerLimitAmendResult
type CustomerLimitAmendByCIFParam = customerlimitamendbycif.CustomerLimitAmendByCIFParam
type CBECoreAPIInterface interface {
	CustomerLimitFetchByCustomerNumber(param CustomerLimitFetchByCIFParam) (*CustomerLimitFetchByCIFResult, error)
	CustomerLimitAmendByCustomerNumber(param CustomerLimitAmendByCIFParam) (*CustomerLimitAmendByCIFResult, error)
	CustomerLimitFetchByService(param CustomerLimitFetchByServiceParam) (*CustomerLimitFetchByServiceResult, error)

	ServiceLimit(param ServiceLimitParam) (*ServiceLimitResult, error)
	FundTransfer(param FundTransferParam) (*FundTransferResult, error)
	FundTransferCheck(param FundTransferCheckParam) (*FundTransferCheckResult, error)
	RevertFundTransfer(param RevertFundTransferParam) (*RevertFundTransferResult, error)
	AccountLookup(param AccountLookupParam) (*AccountLookupResult, error)
	AccountCreation(param AccountCreationParam) (*AccountCreationResult, error)

	LockedAmountFT(param LockedAmountFTParam) (*LockedAmountFTResult, error)
	ListLockedAmount(param ListLockedAmountParam) (*ListLockedAmountResult, error)
	CreateLockedAmount(param CreateLockedAmountParam) (*CreateLockedAmountResult, error)
	ReleaseLockedAmount(param ReleaseLockedAmountParam) (*ReleaseLockedAmountResult, error)

	ListStandingOrder(param ListStandingOrderParam) (*ListStandingOrderResult, error)
	UpdateStandingOrder(param UpdateStandingOrderParam) (*UpdateStandingOrderResult, error)
	CreateStandingOrder(param CreateStandingOrderParam) (*CreateStandingOrderResult, error)
	CancleStandingOrder(param CancleStandingOrderParam) (*CancelStandingOrderResult, error)
	ListStandingOrderHistory(param ListStandingOrderHistoryParam) (*ListStandingOrderHistoryResult, error)

	MiniStatementByLimit(param MiniStatementByLimitParams) (*MiniStatementByLimitResult, error)
	MiniStatementByDateRange(param MiniStatementByDateRangeParam) (*MiniStatementByDateRangeResult, error)

	PhoneLookup(param PhoneLookupParam) (*PhoneLookupResult, error)

	CustomerLookup(param CustomerLookupParam) (*CustomerLookupResult, error)
	AccountList(param AccountListParam) (*AccountListResult, error)
	CardReplace(param CardReplaceParam) (*CardReplaceResult, error)
	CardRequest(param CardRequestParam) (*CardRequestResult, error)

	ExchangeRates() (*ExchangeRatesResult, error)

	SplitPayment(param SplitPaymentParam) (*SplitPaymentResult, error)
}

type CBECoreCredential struct {
	Username           string
	Password           string
	Url                string
	FraudAPICredential FraudAPICredential
}

type FraudAPICredential struct {
	Authorization string
	ForwardHost   string
	Url           string
}

type CBECoreAPI struct {
	config *internal.Config
}

func (c *CBECoreAPI) CustomerLimitFetchByCustomerNumber(param CustomerLimitFetchByCIFParam) (*CustomerLimitFetchByCIFResult, error) {
	params := customerlimitfetchbycif.Params{
		Username:       c.config.Username,
		Password:       c.config.Password,
		CustomerNumber: param.CustomerNumber,
	}

	xmlRequest := customerlimitfetchbycif.NewCustomerLimitFetchByCIF(params)
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

	result, err := customerlimitfetchbycif.ParseCustomerLimitFetchByCIFSOAP(string(responseData))
	if err != nil {
		return nil, err
	}

	return result, nil
}
func (c *CBECoreAPI) CustomerLimitFetchByService(param CustomerLimitFetchByServiceParam) (*CustomerLimitFetchByServiceResult, error) {
	params := customerlimitfetchbyservice.Params{
		Username:    c.config.Username,
		Password:    c.config.Password,
		ServiceCode: param.ServiceCode,
	}
	xmlRequest := customerlimitfetchbyservice.NewCustomerLimitFetchByService(params)
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
	result, err := customerlimitfetchbyservice.ParseCustomerLimitFetchByServiceSOAP(string(responseData))
	if err != nil {
		return nil, err
	}

	return result, nil
}
func (c *CBECoreAPI) CustomerLimitAmendByCustomerNumber(param CustomerLimitAmendByCIFParam) (*CustomerLimitAmendByCIFResult, error) {
	params := customerlimitamendbycif.Params{
		Username:       c.config.Username,
		Password:       c.config.Password,
		CustomerNumber: param.CustomerNumber,
		ChannelLimit:   param.ChannelLimit,
	}
	xmlRequest := customerlimitamendbycif.NewCustomerLimitAmendByCIF(params)
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

	result, err := customerlimitamendbycif.ParseCustomerLimitAmendByCIFSOAP(string(responseData))
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *CBECoreAPI) SplitPayment(param SplitPaymentParam) (*SplitPaymentResult, error) {
	params := splitpayment.Param{
		Username:                 c.config.Username,
		Password:                 c.config.Password,
		DebitCurrency:            param.DebitCurrency,
		DebitAccount:             param.DebitAccount,
		DebitReference:           param.DebitReference,
		CreditCurrency:           param.CreditCurrency,
		CreditAccountInformation: param.CreditAccountInformation,
	}

	xmlRequest := splitpayment.NewSplitPayment(params)
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

	result, err := splitpayment.ParseSplitPaymentSOAP(string(responseData))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *CBECoreAPI) AccountCreation(param AccountCreationParam) (*AccountCreationResult, error) {
	params := accountcreation.Params{
		Username: c.config.Username,
		Password: c.config.Password,
		Customer: param.Customer,
		Category: param.Category,
		Currency: param.Currency,
	}
	xmlRequest := accountcreation.NewAccountCreation(params)
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
	result, err := accountcreation.ParseAccountCreationSOAP(string(responseData))
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *CBECoreAPI) ServiceLimit(param ServiceLimitParam) (*ServiceLimitResult, error) {
	params := servicelimit.Params{
		Username: c.config.Username,
		Password: c.config.Password,
	}
	xmlRequest := servicelimit.NewServiceLimit(params)
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
	result, err := servicelimit.ParseServiceLimitSOAP(string(responseData))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *CBECoreAPI) ExchangeRates() (*ExchangeRatesResult, error) {
	params := exchangerate.Params{
		Username: c.config.Username,
		Password: c.config.Password,
	}
	xmlRequest := exchangerate.NewExchangeRate(params)
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
	result, err := exchangerate.ParseExchangeRateSOAP(string(responseData))

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *CBECoreAPI) PhoneLookup(param PhoneLookupParam) (*PhoneLookupResult, error) {
	params := phonelookup.Params{
		Username:    c.config.Username,
		Password:    c.config.Password,
		PhoneNumber: param.PhoneNumber,
	}
	xmlRequest := phonelookup.NewPhoneLookup(params)
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
	result, err := phonelookup.ParsePhoneLookupSOAP(string(responseData))
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *CBECoreAPI) RevertFundTransfer(param RevertFundTransferParam) (*RevertFundTransferResult, error) {
	params := revertfundtransfer.Params{
		Username:      c.config.Username,
		Password:      c.config.Password,
		TransactionID: param.TransactionID,
	}

	xmlRequest := revertfundtransfer.NewRevertFundTransfer(params)
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

	result, err := accountlookup.ParseAccountLookupSOAP(string(responseData))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *CBECoreAPI) AccountList(param AccountListParam) (*AccountListResult, error) {
	params := accountlist.Params{
		Username:      c.config.Username,
		Password:      c.config.Password,
		ColumnName:    param.ColumnName,
		CriteriaValue: param.CriteriaValue,
	}

	xmlRequest := accountlist.NewAccountList(params)
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

	result, err := accountlist.ParseAccountListSOAP(string(responseData))
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
		CreditAmount:        param.CreditAmount,
		ServiceCode:         param.ServiceCode,
		DebitReference:      param.DebitReference,
		CreditReference:     param.CreditReference,
		DebitAmount:         param.DebitAmount,
		TransactionID:       param.TransactionID,
		PaymentDetail:       param.PaymentDetail,
		Meta:                param.Meta,
	}

	fraud := frauddetection.NewFraudAPI(
		c.config.FraudAPIConfig.Authorization,
		c.config.FraudAPIConfig.ForwardHost,
		c.config.FraudAPIConfig.Url,
	)
	response, err := fraud.Call(param.Meta)
	if err != nil {
		return nil, fmt.Errorf("fraud detection call failed: %w", err)
	}

	if response.Result != "approved" {
		// return nil, errors.New("transaction blocked by fraud detection")
		return &fundtransfer.FundTransferResult{
			Success:  false,
			Messages: []string{"transaction blocked by fraud detection"},
		}, nil
	}

	xmlRequest := fundtransfer.NewFundTransfer(params)
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

	result, err := fundtransfer.ParseFundTransferSOAP(string(responseData))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *CBECoreAPI) FundTransferCheck(param FundTransferCheckParam) (*FundTransferCheckResult, error) {
	params := fundtransfercheck.Params{
		Username: c.config.Username,
		Password: c.config.Password,
		FTNumber: param.FTNumber,
	}

	xmlRequest := fundtransfercheck.NewFundTransferCheck(params)
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

	result, err := fundtransfercheck.ParseFundTransferCheckSOAP(string(responseData))
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
		ServiceCode:         param.ServiceCode,
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

func (c *CBECoreAPI) ListStandingOrderHistory(param ListStandingOrderHistoryParam) (*ListStandingOrderHistoryResult, error) {
	params := standingorderlisthistory.Params{
		Username:      c.config.Username,
		Password:      c.config.Password,
		AccountNumber: param.AccountNumber,
	}
	xmlRequest := standingorderlisthistory.NewListStandingOrderHistory(params)
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
	result, err := standingorderlisthistory.ParseStandingOrderListHistorySOAP(string(responseData))
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

func (c *CBECoreAPI) CustomerLookup(param CustomerLookupParam) (*CustomerLookupResult, error) {
	params := customerlookup.Params{
		Username:           c.config.Username,
		Password:           c.config.Password,
		CustomerIdentifier: param.CustomerIdentifier,
	}
	xmlRequest := customerlookup.NewCustomerLookup(params)
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

	result, err := customerlookup.ParseCustomerLookupSOAP(string(responseData))
	fmt.Println("result", result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *CBECoreAPI) CardReplace(param CardReplaceParam) (*CardReplaceResult, error) {
	params := cardreplace.Params{
		Username:      c.config.Username,
		Password:      c.config.Password,
		AccountNumber: param.AccountNumber,
	}
	xmlRequest := cardreplace.NewCardReplace(params)
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

	result, err := cardreplace.ParseCardReplaceResponse(string(responseData))
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *CBECoreAPI) CardRequest(param CardRequestParam) (*CardRequestResult, error) {
	params := cardrequest.Params{
		Username:      c.config.Username,
		Password:      c.config.Password,
		AccountNumner: param.AccountNumber,
		BranchCode:    param.BranchCode,
		PhoneNumber:   param.PhoneNumber,
		CardType:      param.CardType,
	}

	xmlRequest := cardrequest.NewCardRequest(params)
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

	result, err := cardrequest.ParseATMCardRequestSOAP(string(responseData))
	if err != nil {
		return nil, err
	}
	return result, nil
}

func NewCBECoreAPI(param CBECoreCredential) CBECoreAPIInterface {
	config := internal.SetConfig(
		param.Username,
		param.Password,
		param.Url,
		param.FraudAPICredential.Authorization,
		param.FraudAPICredential.Url,
		param.FraudAPICredential.ForwardHost,
	)
	return &CBECoreAPI{
		config: config,
	}
}
