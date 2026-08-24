package core

import (
	"context"
	"fmt"
	"io"
	"maps"
	"strings"
	"time"

	"github.com/hugokessem/coreio/core/internal"
	accountcreation "github.com/hugokessem/coreio/lib/core/account/account_creation"
	accountlist "github.com/hugokessem/coreio/lib/core/account/account_list"
	accountlookup "github.com/hugokessem/coreio/lib/core/account/account_lookup"
	billpayment "github.com/hugokessem/coreio/lib/core/bill_payment"
	billpaymentverify "github.com/hugokessem/coreio/lib/core/bill_payment/bill_payment_verify"
	cardreplace "github.com/hugokessem/coreio/lib/core/card/card_replace"
	cardrequest "github.com/hugokessem/coreio/lib/core/card/card_request"
	"github.com/hugokessem/coreio/lib/core/fayda"
	frauddetection "github.com/hugokessem/coreio/lib/core/fraud_detection"
	namelookup "github.com/hugokessem/coreio/lib/core/name_lookup"
	splitpayment "github.com/hugokessem/coreio/lib/core/split_payment"
	"github.com/redis/go-redis/v9"
	"gitlab.com/bersufekadgetachew/cbe-super-app-shared/shared/survey"

	customercreation "github.com/hugokessem/coreio/lib/core/customer/customer_creation"
	customerdetail "github.com/hugokessem/coreio/lib/core/customer/customer_detail"
	customerfetch "github.com/hugokessem/coreio/lib/core/customer/customer_fetch"
	customerlimitamendbycif "github.com/hugokessem/coreio/lib/core/customer/customer_limit_amend_by_cif"
	customerlimitfetchbycif "github.com/hugokessem/coreio/lib/core/customer/customer_limit_fetch_by_cif"
	customerlimitfetchbycifreturnservice "github.com/hugokessem/coreio/lib/core/customer/customer_limit_fetch_by_cif_return_service"
	customerlimitfetchbyservice "github.com/hugokessem/coreio/lib/core/customer/customer_limit_fetch_by_service"
	customerlookup "github.com/hugokessem/coreio/lib/core/customer/customer_lookup"
	verifyaml "github.com/hugokessem/coreio/lib/core/customer/verify_aml"
	eligibility "github.com/hugokessem/coreio/lib/core/eligibility"
	exchangerate "github.com/hugokessem/coreio/lib/core/exchange_rate"
	fundtransfer "github.com/hugokessem/coreio/lib/core/fund_transfer/fund_transfer"
	fundtransfercheck "github.com/hugokessem/coreio/lib/core/fund_transfer/fund_transfer_check"
	fundtransferverify "github.com/hugokessem/coreio/lib/core/fund_transfer/fund_transfer_verify"
	statuscheck "github.com/hugokessem/coreio/lib/core/fund_transfer/status_check"
	lockedamountcreate "github.com/hugokessem/coreio/lib/core/locked_amount/locked_amount_create"
	lockedamountft "github.com/hugokessem/coreio/lib/core/locked_amount/locked_amount_ft"
	lockedamountlist "github.com/hugokessem/coreio/lib/core/locked_amount/locked_amount_list"
	lockedamountrelease "github.com/hugokessem/coreio/lib/core/locked_amount/locked_amount_release"
	ministatementbydaterange "github.com/hugokessem/coreio/lib/core/mini_statement/mini_statement_by_date_range"
	ministatementbylimit "github.com/hugokessem/coreio/lib/core/mini_statement/mini_statement_by_limit"
	phonelookup "github.com/hugokessem/coreio/lib/core/phone_lookup"
	revertfundtransfer "github.com/hugokessem/coreio/lib/core/revert_fund_transfer"
	servicedetail "github.com/hugokessem/coreio/lib/core/service/service_detail"
	servicelimit "github.com/hugokessem/coreio/lib/core/service/service_limit"
	standingordercancel "github.com/hugokessem/coreio/lib/core/standing_order/standing_order_cancel"
	standingordercreate "github.com/hugokessem/coreio/lib/core/standing_order/standing_order_create"
	standingorderlist "github.com/hugokessem/coreio/lib/core/standing_order/standing_order_list"
	standingorderlisthistory "github.com/hugokessem/coreio/lib/core/standing_order/standing_order_list_history"
	standingorderupdate "github.com/hugokessem/coreio/lib/core/standing_order/standing_order_update"
	statuschange "github.com/hugokessem/coreio/lib/core/super_app/status_change"
	valueobject "github.com/hugokessem/coreio/lib/core/super_app/status_change/value_object"
	superappsubscribe "github.com/hugokessem/coreio/lib/core/super_app/subscribe"
	superappunsubscribe "github.com/hugokessem/coreio/lib/core/super_app/unsbscribe"
	"github.com/hugokessem/coreio/utils"
)

type BillPaymentVerifyParam = billpaymentverify.BillPaymentVerifyParam
type BillPaymentVerifyResult = billpaymentverify.BillPaymentVerifyResult

type ExchangeRatesResult = exchangerate.ExchangeRateResult
type AccountLookupParam = accountlookup.AccountLookupParam
type AccountLookupResult = accountlookup.AccountLookupResult
type AccountListParam = accountlist.AccountListParams
type AccountListResult = accountlist.AccountListResult
type ServiceLimitParam = servicelimit.ServiceLimitParam
type ServiceLimitResult = servicelimit.ServiceLimitResult

type SplitPaymentParam = splitpayment.SplitPaymentParam
type SplitPaymentResult = splitpayment.SplitPaymentResult

type NameLookupParam = namelookup.NameLookupParam
type NameLookupResult = namelookup.NameLookupResult

type CreateCustomerParam = customercreation.CreateCustomerParams
type CreateCustomerResult = customercreation.CustomerCreationResult

type AccountCreationParam = accountcreation.AccountCreationParams
type AccountCreationResult = accountcreation.AccountCreationResult

type FundTransferParam = fundtransfer.FundTransferParam
type FundTransferResult = fundtransfer.FundTransferResult
type FundTransferCheckParam = fundtransfercheck.FundTransferCheckParams
type FundTransferCheckResult = fundtransfercheck.FundTransferCheckResult
type FundTransferVerifyParam = fundtransferverify.FundTransferVerifyParams
type FundTransferVerifyResult = fundtransferverify.FundTransferVerifyResult
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

type VerifyAMLParam = verifyaml.VerifyAMLParam
type VerifyAMLResult = verifyaml.VerifyAMLResult

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

type StatusCheckParam = statuscheck.StatusCheckParam
type StatusCheckResult = statuscheck.StatusCheckResult

type CustomerDetailParam = customerdetail.CustomerDetailParam
type CustomerDetailResult = customerdetail.CustomerDetailResult

type ServiceDetailParam = servicedetail.ServiceDetailParams
type ServiceDetailResult = servicedetail.ServiceDetailResult

type SuperAppSubscribeParam = superappsubscribe.Params
type SuperAppSubscribeResult = superappsubscribe.SubscribeResult
type SuperAppUnsubscribeParam = superappunsubscribe.Params
type SuperAppUnsubscribeResult = superappunsubscribe.UnsubscribeResult
type SuperAppStatusChangeParam = statuschange.StatusParam
type SuperAppStatusChangeResult = statuschange.StatusChangeResult

type CustomerFetchParam = customerfetch.CustomerFetchParam
type CustomerFetchResult = customerfetch.CustomerFetchResult

type CustomerLimitFetchByCIFReturnServiceParam = customerlimitfetchbycifreturnservice.CustomerLimitFetchByCIFReturnServiceParam
type CustomerLimitFetchByCIFReturnServiceResult = customerlimitfetchbycifreturnservice.CustomerLimitFetchByCIFReturnServiceResult

type BillPaymentParam = billpayment.BillPaymentParams
type BillPaymentResult = billpayment.BillPaymentResult

type FaydaParam = fayda.FaydaParam
type FaydaResult = fayda.FaydaResult

type EligibilityParam = eligibility.EligibilityParam
type EligibilityResult = eligibility.CustomerEligibilityResult

type CBECoreAPIInterface interface {
	CustomerLimitFetchByCustomerNumber(ctx context.Context, param CustomerLimitFetchByCIFParam) (*CustomerLimitFetchByCIFResult, error)
	CustomerLimitAmendByCustomerNumber(ctx context.Context, param CustomerLimitAmendByCIFParam) (*CustomerLimitAmendByCIFResult, error)
	CustomerLimitFetchByService(ctx context.Context, param CustomerLimitFetchByServiceParam) (*CustomerLimitFetchByServiceResult, error)

	ServiceLimit(ctx context.Context, param ServiceLimitParam) (*ServiceLimitResult, error)
	FundTransferVerify(ctx context.Context, param FundTransferVerifyParam) (*FundTransferVerifyResult, error)
	FundTransfer(ctx context.Context, param FundTransferParam) (*FundTransferResult, error)
	FundTransferCheck(ctx context.Context, param FundTransferCheckParam) (*FundTransferCheckResult, error)
	RevertFundTransfer(ctx context.Context, param RevertFundTransferParam) (*RevertFundTransferResult, error)
	AccountLookup(ctx context.Context, param AccountLookupParam) (*AccountLookupResult, error)
	AccountCreation(ctx context.Context, param AccountCreationParam) (*AccountCreationResult, error)

	LockedAmountFT(ctx context.Context, param LockedAmountFTParam) (*LockedAmountFTResult, error)
	ListLockedAmount(ctx context.Context, param ListLockedAmountParam) (*ListLockedAmountResult, error)
	CreateLockedAmount(ctx context.Context, param CreateLockedAmountParam) (*CreateLockedAmountResult, error)
	ReleaseLockedAmount(ctx context.Context, param ReleaseLockedAmountParam) (*ReleaseLockedAmountResult, error)

	ListStandingOrder(ctx context.Context, param ListStandingOrderParam) (*ListStandingOrderResult, error)
	UpdateStandingOrder(ctx context.Context, param UpdateStandingOrderParam) (*UpdateStandingOrderResult, error)
	CreateStandingOrder(ctx context.Context, param CreateStandingOrderParam) (*CreateStandingOrderResult, error)
	CancleStandingOrder(ctx context.Context, param CancleStandingOrderParam) (*CancelStandingOrderResult, error)
	ListStandingOrderHistory(ctx context.Context, param ListStandingOrderHistoryParam) (*ListStandingOrderHistoryResult, error)

	MiniStatementByLimit(ctx context.Context, param MiniStatementByLimitParams) (*MiniStatementByLimitResult, error)
	MiniStatementByDateRange(ctx context.Context, param MiniStatementByDateRangeParam) (*MiniStatementByDateRangeResult, error)

	PhoneLookup(ctx context.Context, param PhoneLookupParam) (*PhoneLookupResult, error)

	CustomerLookup(ctx context.Context, param CustomerLookupParam) (*CustomerLookupResult, error)
	VerifyAML(ctx context.Context, param VerifyAMLParam) (*VerifyAMLResult, error)
	AccountList(ctx context.Context, param AccountListParam) (*AccountListResult, error)
	CardReplace(ctx context.Context, param CardReplaceParam) (*CardReplaceResult, error)
	CardRequest(ctx context.Context, param CardRequestParam) (*CardRequestResult, error)

	StatusCheck(ctx context.Context, param StatusCheckParam) (*StatusCheckResult, error)
	ExchangeRates(ctx context.Context) (*ExchangeRatesResult, error)
	CustomerDetail(ctx context.Context, param CustomerDetailParam) (*CustomerDetailResult, error)

	SplitPayment(ctx context.Context, param SplitPaymentParam) (*SplitPaymentResult, error)
	ServiceDetail(ctx context.Context, param ServiceDetailParam) (*ServiceDetailResult, error)

	NameLookup(ctx context.Context, param NameLookupParam) (*NameLookupResult, error)

	SuperAppSubscribe(ctx context.Context, param SuperAppSubscribeParam) (*SuperAppSubscribeResult, error)
	SuperAppUnsubscribe(ctx context.Context, param SuperAppUnsubscribeParam) (*SuperAppUnsubscribeResult, error)
	SuperAppStatusChange(ctx context.Context, param SuperAppStatusChangeParam) (*SuperAppStatusChangeResult, error)
	CustomerLimitFetchByCIFReturnService(ctx context.Context, param CustomerLimitFetchByCIFReturnServiceParam) (*CustomerLimitFetchByCIFReturnServiceResult, error)
	CreateCustomer(ctx context.Context, param CreateCustomerParam) (*CreateCustomerResult, error)
	CustomerFetch(ctx context.Context, param CustomerFetchParam) (*CustomerFetchResult, error)
	Eligibility(ctx context.Context, param EligibilityParam) (*EligibilityResult, error)
	BillPayment(ctx context.Context, param BillPaymentParam) (*BillPaymentResult, error)
	AccountCreate(ctx context.Context, param CreateCustomerParam, accountCreateURL, category string) (*CusteomerAccountCreationResponse, error)
	Fayda(ctx context.Context, param FaydaParam) (*FaydaResult, error)
	CheckIfUserExists(ctx context.Context, param UserExistsParam) (*UserExistsResult, error)
	BillPaymentVerify(ctx context.Context, param BillPaymentVerifyParam) (*BillPaymentVerifyResult, error)
}

// commented
type CBECoreCredential struct {
	Username           string
	Password           string
	Url                string
	FraudAPICredential FraudAPICredential
	RedisClient        *redis.Client
	RedisTTL           time.Duration
}

type FraudAPICredential struct {
	Authorization string
	ForwardHost   string
	Url           string
}

type CBECoreAPI struct {
	config *internal.Config
}

const (
	Key        = "Content-Type"
	Value      = "text/xml; charset=utf-8"
	timeout    = 120 * time.Second
	maxRetries = 1
)

func (c *CBECoreAPI) BillPaymentVerify(ctx context.Context, param BillPaymentVerifyParam) (*BillPaymentVerifyResult, error) {
	params := billpaymentverify.Param{
		Username:            c.config.Username,
		Password:            c.config.Password,
		DebitAccountNumber:  param.DebitAccountNumber,
		DebitCurrency:       param.DebitCurrency,
		DebitAmount:         param.DebitAmount,
		DebitReference:      param.DebitReference,
		CreditReference:     param.CreditReference,
		CreditAccountNumber: param.CreditAccountNumber,
		CreditCurrency:      param.CreditCurrency,
		ClientReference:     param.ClientReference,
		ServiceCode:         param.ServiceCode,
		SuperappUserCode:    param.SuperappUserCode,
	}
	fmt.Println("params", params)
	xmlRequest := billpaymentverify.NewBillPaymentVerify(params)
	headers := map[string]string{
		Key: Value,
	}
	fmt.Println("xmlRequest", xmlRequest)
	resp, err := utils.DoPost(ctx, c.config.Url, xmlRequest, utils.Config{
		Timeout:    timeout,
		MaxRetries: maxRetries,
	}, headers)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println("responseData", string(responseData))
	result, err := billpaymentverify.ParseBillPaymentVerifySOAP(string(responseData))
	if err != nil {
		return nil, err
	}
	fmt.Println("result", result)
	return result, nil
}

func (c *CBECoreAPI) BillPayment(ctx context.Context, param BillPaymentParam) (*BillPaymentResult, error) {
	params := billpayment.Params{
		Username:            c.config.Username,
		Password:            c.config.Password,
		DebitAccountNumber:  param.DebitAccountNumber,
		DebitCurrency:       param.DebitCurrency,
		DebitAmount:         param.DebitAmount,
		DebitReference:      param.DebitReference,
		CrediterReference:   param.CrediterReference,
		CreditAccountNumber: param.CreditAccountNumber,
		CreditCurrency:      param.CreditCurrency,
		ServiceCode:         param.ServiceCode,
		SuperappUserCode:    param.SuperappUserCode,
		ClientReference:     param.ClientReference,
		CreditAmount:        param.CreditAmount,
	}

	xmlRequest := billpayment.NewBillPayment(params)
	headers := map[string]string{
		Key: Value,
	}
	resp, err := utils.DoPost(ctx, c.config.Url, xmlRequest, utils.Config{
		Timeout:    timeout,
		MaxRetries: maxRetries,
	}, headers)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result, err := billpayment.ParseBillPaymentSOAP(string(responseData))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *CBECoreAPI) CustomerFetch(ctx context.Context, param CustomerFetchParam) (*CustomerFetchResult, error) {

	params := customerfetch.Params{
		Username:       c.config.Username,
		Password:       c.config.Password,
		CustomerNumber: param.CustomerNumber,
		FetchBy:        param.FetchBy,
	}

	xmlRequest := customerfetch.NewCustomerFetch(params)
	headers := map[string]string{
		Key: Value,
	}

	resp, err := utils.DoPost(ctx, c.config.Url, xmlRequest, utils.Config{
		Timeout:    timeout,
		MaxRetries: maxRetries,
	}, headers)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	result, err := customerfetch.ParseCustomerFetchSOAP(string(responseData))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *CBECoreAPI) Eligibility(ctx context.Context, param EligibilityParam) (*EligibilityResult, error) {
	params := eligibility.Param{
		Username:      c.config.Username,
		Password:      c.config.Password,
		FetchBy:       param.FetchBy,
		CriticalValue: param.CriticalValue,
	}

	xmlRequest := eligibility.NewEligibility(params)
	headers := map[string]string{
		Key: Value,
	}

	resp, err := utils.DoPost(ctx, c.config.Url, xmlRequest, utils.Config{
		Timeout:    timeout,
		MaxRetries: maxRetries,
	}, headers)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result, err := eligibility.ParseCustomerEligibilitySOAP(string(responseData))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *CBECoreAPI) AccountCreation(ctx context.Context, param AccountCreationParam) (*AccountCreationResult, error) {
	params := accountcreation.Params{
		Username:       c.config.Username,
		Password:       c.config.Password,
		AccountOfficer: param.AccountOfficer,
		CustomerNumber: param.CustomerNumber,
		Category:       param.Category,
		Currency:       param.Currency,
		Url:            param.Url,
		Header:         param.Header,
	}
	xmlRequest := accountcreation.NewAccountCreation(params)
	headers := map[string]string{
		"Content-Type": "application/xml",
	}

	if param.Header != nil {
		maps.Copy(headers, param.Header)
	}

	url := c.config.Url
	if param.Url != "" {
		url = param.Url
	}
	resp, err := utils.DoPost(ctx, url, xmlRequest, utils.Config{
		Timeout:    timeout,
		MaxRetries: maxRetries,
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

func GetMenmonicNumber(phoneNumber string, index int, names ...string) string {
	lastNineDigit := phoneNumber[len(phoneNumber)-9:]
	return fmt.Sprintf("%s%s", string(names[index][0]), lastNineDigit)
}

func (c *CBECoreAPI) CreateCustomer(ctx context.Context, param CreateCustomerParam) (*CreateCustomerResult, error) {
	params := customercreation.Params{
		Username:                c.config.Username,
		Password:                c.config.Password,
		Company:                 param.Company,
		FirstName:               param.FirstName,
		MiddleName:              param.MiddleName,
		LastName:                param.LastName,
		Menmonic:                param.Menmonic,
		PhoneNumber:             param.PhoneNumber,
		Address:                 param.Address,
		PostalCode:              param.PostalCode,
		ISOCountryCode:          param.ISOCountryCode,
		AccountOffice:           param.AccountOffice,
		Industry:                param.Industry,
		ISONationalityCode:      param.ISONationalityCode,
		ISOResidentCode:         param.ISOResidentCode,
		UniqueID:                param.UniqueID,
		IssuesBy:                param.IssuesBy,
		IssuedDate:              param.IssuedDate,
		ExpiryDate:              param.ExpiryDate,
		Title:                   param.Title,
		Gender:                  param.Gender,
		Street:                  param.Street,
		DateOfBirth:             param.DateOfBirth,
		MaritalStatus:           param.MaritalStatus,
		NoOfDependents:          param.NoOfDependents,
		Email:                   param.Email,
		EmploymentStatus:        param.EmploymentStatus,
		Occupation:              param.Occupation,
		EmployerName:            param.EmployerName,
		EmployerAddress:         param.EmployerAddress,
		EmployerBusiness:        param.EmployerBusiness,
		CustomerCurrency:        param.CustomerCurrency,
		Salary:                  param.Salary,
		AnnualBonus:             param.AnnualBonus,
		NetMonthlyIncome:        param.NetMonthlyIncome,
		NetMonthlyExpence:       param.NetMonthlyExpence,
		TownCountry:             param.TownCountry,
		LegalDocumenetName:      param.LegalDocumenetName,
		TinNumber:               param.TinNumber,
		CustomerOccupation:      param.CustomerOccupation,
		EducationStatus:         param.EducationStatus,
		MotherName:              param.MotherName,
		FATCACompliant:          param.FATCACompliant,
		USPerson:                param.USPerson,
		KebeleHNO:               param.KebeleHNO,
		CustomerSubSegment:      param.CustomerSubSegment,
		CustomerSegment:         param.CustomerSegment,
		GrandFatherName:         param.GrandFatherName,
		CustomerGroup:           param.CustomerGroup,
		NationalId:              param.NationalId,
		CommunicationPreference: param.CommunicationPreference,
		Url:                     param.Url,
		Header:                  param.Header,
	}

	xmlRequest := customercreation.NewCustomerCreation(params)
	headers := map[string]string{
		"Content-Type": "application/xml",
	}

	if param.Header != nil {
		maps.Copy(headers, param.Header)
	}
	url := param.Url
	resp, err := utils.DoPost(ctx, url, xmlRequest, utils.Config{
		Timeout:    timeout,
		MaxRetries: maxRetries,
	}, headers)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	result, err := customercreation.ParseCustomerCreationSOAP(string(responseData))
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *CBECoreAPI) Fayda(ctx context.Context, param FaydaParam) (*FaydaResult, error) {
	params := fayda.Params{
		Username: c.config.Username,
		Password: c.config.Password,
		NID:      param.NID,
	}
	xmlRequest := fayda.NewFayda(params)
	headers := map[string]string{
		Key: Value,
	}
	resp, err := utils.DoPost(ctx, c.config.Url, xmlRequest, utils.Config{
		Timeout:    timeout,
		MaxRetries: maxRetries,
	}, headers)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	responseData, err := utils.ReadResponseBody(resp, "fayda")
	if err != nil {
		return nil, err
	}
	result, err := fayda.ParseFaydaSOAP(string(responseData))
	if err != nil {
		return nil, err
	}

	return result, nil
}

type CusteomerAccountCreationResponse struct {
	CustomerCreationDetail *CreateCustomerResult
	AccountCreationDetail  *AccountCreationResult
	PhoneLookupDetail      *PhoneLookupResult
	FaydaDetail            *FaydaResult
	Messages               []string
}

type UserExistsResult struct {
	PhoneLookupDetail *PhoneLookupResult
	FaydaDetail       *FaydaResult
	Success           bool
	Messages          []string
}

type UserExistsParam struct {
	PhoneNumber string
	NationalId  string
}

func (c *CBECoreAPI) CheckIfUserExists(ctx context.Context, param UserExistsParam) (*UserExistsResult, error) {
	nemonic, err := c.PhoneLookup(ctx, PhoneLookupParam{
		PhoneNumber: param.PhoneNumber,
	})
	if err != nil {
		return &UserExistsResult{
			Messages: []string{fmt.Sprintf("phone lookup failed: %v", err)},
		}, nil
	}

	if nemonic.Success {
		return &UserExistsResult{
			PhoneLookupDetail: nemonic,
			Success:           false,
			Messages:          nemonic.Message,
		}, nil
	}

	faydaResult, err := c.Fayda(ctx, FaydaParam{
		NID: param.NationalId,
	})
	if err != nil {
		return &UserExistsResult{
			Messages: []string{fmt.Sprintf("fayda verification failed: %v", err)},
			Success:  false,
		}, nil
	}

	if !faydaResult.Success {
		return &UserExistsResult{
			FaydaDetail: faydaResult,
			Messages:    faydaResult.Message,
			Success:     false,
		}, nil
	}

	if faydaResult.Detail.CustomerFlag == "FOUND" {
		return &UserExistsResult{
			FaydaDetail: faydaResult,
			Messages:    faydaResult.Message,
			Success:     false,
		}, nil
	}

	return &UserExistsResult{
		PhoneLookupDetail: nil,
		FaydaDetail:       nil,
		Messages:          []string{"user not found in both phone lookup and fayda, you are good to go!"},
		Success:           true,
	}, nil
}

func (c *CBECoreAPI) AccountCreate(ctx context.Context, param CreateCustomerParam, accountCreateURL, category string) (*CusteomerAccountCreationResponse, error) {

	customer, err := c.CreateCustomer(ctx, param)
	if err != nil {
		messages := []string{err.Error()}
		if customer != nil && len(customer.Messages) > 0 {
			messages = customer.Messages
		}
		return &CusteomerAccountCreationResponse{
			CustomerCreationDetail: customer,
			Messages:               messages,
		}, nil
	}

	if customer == nil || !customer.Success || customer.Detail == nil {
		messages := []string{"customer creation failed"}
		if customer != nil && len(customer.Messages) > 0 {
			messages = customer.Messages
		}
		return &CusteomerAccountCreationResponse{
			CustomerCreationDetail: customer,
			Messages:               messages,
		}, nil
	}

	var isAMLCheck bool
	var amlMessages []string
	for _, value := range customer.Detail.Override {
		if strings.Contains(value, "AML") {
			isAMLCheck = true
			amlMessages = append(amlMessages, value)
		}
	}

	if isAMLCheck {
		return &CusteomerAccountCreationResponse{
			CustomerCreationDetail: customer,
			Messages:               amlMessages,
		}, nil
	}

	account, err := c.AccountCreation(ctx, AccountCreationParam{
		CustomerNumber: customer.Detail.CustomerNumber,
		Category:       category,
		Currency:       param.CustomerCurrency,
		AccountOfficer: param.AccountOffice,
		Header:         param.Header,
		Url:            accountCreateURL,
	})
	if err != nil {
		messages := []string{err.Error()}
		if account != nil && len(account.Messages) > 0 {
			messages = account.Messages
		}
		return &CusteomerAccountCreationResponse{
			CustomerCreationDetail: customer,
			Messages:               messages,
		}, nil
	}

	if !account.Success {
		return &CusteomerAccountCreationResponse{
			CustomerCreationDetail: customer,
			Messages:               account.Messages,
		}, nil
	}

	return &CusteomerAccountCreationResponse{
		AccountCreationDetail:  account,
		CustomerCreationDetail: customer,
	}, nil
}

func (c *CBECoreAPI) CustomerLimitFetchByCIFReturnService(ctx context.Context, param CustomerLimitFetchByCIFReturnServiceParam) (*CustomerLimitFetchByCIFReturnServiceResult, error) {
	params := customerlimitfetchbycifreturnservice.Params{
		Username:       c.config.Username,
		Password:       c.config.Password,
		CustomerNumber: param.CustomerNumber,
	}
	xmlRequest := customerlimitfetchbycifreturnservice.NewCustomerLimitFetchByCIFReturnService(params)
	headers := map[string]string{
		Key: Value,
	}
	resp, err := utils.DoPost(ctx, c.config.Url, xmlRequest, utils.Config{
		Timeout:    timeout,
		MaxRetries: maxRetries,
	}, headers)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	result, err := customerlimitfetchbycifreturnservice.ParseCustomerLimitFetchByCIFReturnServiceSOAP(string(responseData))
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *CBECoreAPI) SuperAppSubscribe(ctx context.Context, param SuperAppSubscribeParam) (*SuperAppSubscribeResult, error) {
	params := superappsubscribe.Params{
		Username: c.config.Username,
		Password: c.config.Password,
		UserCode: param.UserCode,
	}

	xmlRequest := superappsubscribe.NewSubscribe(params)
	headers := map[string]string{
		Key: Value,
	}
	resp, err := utils.DoPost(ctx, c.config.Url, xmlRequest, utils.Config{
		Timeout:    timeout,
		MaxRetries: maxRetries,
	}, headers)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	result, err := superappsubscribe.ParseSubscribeResponseSOAP(string(responseData))
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *CBECoreAPI) SuperAppUnsubscribe(ctx context.Context, param SuperAppUnsubscribeParam) (*SuperAppUnsubscribeResult, error) {
	params := superappunsubscribe.Params{
		Username: c.config.Username,
		Password: c.config.Password,
		UserCode: param.UserCode,
	}
	xmlRequest := superappunsubscribe.NewUnsubscribe(params)
	headers := map[string]string{
		Key: Value,
	}
	resp, err := utils.DoPost(ctx, c.config.Url, xmlRequest, utils.Config{
		Timeout:    timeout,
		MaxRetries: maxRetries,
	}, headers)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	result, err := superappunsubscribe.ParseUnsubscribeResponseSOAP(string(responseData))
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *CBECoreAPI) SuperAppStatusChange(ctx context.Context, param SuperAppStatusChangeParam) (*SuperAppStatusChangeResult, error) {
	params := statuschange.Param{
		Username: c.config.Username,
		Password: c.config.Password,
		UserCode: param.UserCode,
		Status:   valueobject.Status(param.Status),
	}

	xmlRequest, err := statuschange.NewStatusChange(params)
	if err != nil {
		return nil, err
	}
	headers := map[string]string{
		Key: Value,
	}
	resp, err := utils.DoPost(ctx, c.config.Url, xmlRequest, utils.Config{
		Timeout:    timeout,
		MaxRetries: maxRetries,
	}, headers)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	result, err := statuschange.ParseSuperAppUserChangeStatusResponseSOAP(string(responseData))
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (c *CBECoreAPI) NameLookup(ctx context.Context, param NameLookupParam) (*NameLookupResult, error) {
	params := namelookup.Params{
		Username:      c.config.Username,
		Password:      c.config.Password,
		AccountNumber: param.AccountNumber,
	}

	xmlRequest := namelookup.NewNameLookup(params)
	headers := map[string]string{
		Key: Value,
	}

	resp, err := utils.DoPost(ctx, c.config.Url, xmlRequest, utils.Config{
		Timeout:    timeout,
		MaxRetries: maxRetries,
	}, headers)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result, err := namelookup.ParseNameLookupSOAP(string(responseData))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *CBECoreAPI) ServiceDetail(ctx context.Context, param ServiceDetailParam) (*ServiceDetailResult, error) {
	params := servicedetail.Params{
		Username:    c.config.Username,
		Password:    c.config.Password,
		ServiceCode: param.ServiceCode,
	}

	xmlRequest := servicedetail.NewServiceDetail(params)
	headers := map[string]string{
		Key: Value,
	}

	resp, err := utils.DoPost(ctx, c.config.Url, xmlRequest, utils.Config{
		Timeout:    timeout,
		MaxRetries: maxRetries,
	}, headers)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result, err := servicedetail.ParseServiceDetailSOAP(string(responseData))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *CBECoreAPI) CustomerDetail(ctx context.Context, param CustomerDetailParam) (*CustomerDetailResult, error) {
	params := customerdetail.Params{
		Username:       c.config.Username,
		Password:       c.config.Password,
		CustomerNumber: param.CustomerNumber,
	}

	xmlRequest := customerdetail.NewCustomerDetail(params)
	headers := map[string]string{
		Key: Value,
	}

	resp, err := utils.DoPost(ctx, c.config.Url, xmlRequest, utils.Config{
		Timeout:    timeout,
		MaxRetries: maxRetries,
	}, headers)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result, err := customerdetail.ParseCustomerDetailSOAP(string(responseData))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *CBECoreAPI) StatusCheck(ctx context.Context, param StatusCheckParam) (*StatusCheckResult, error) {
	params := statuscheck.Params{
		Username:      c.config.Username,
		Password:      c.config.Password,
		TransactionID: param.TransactionID,
	}

	xmlRequest := statuscheck.NewStatusCheck(params)
	headers := map[string]string{
		Key: Value,
	}

	resp, err := utils.DoPost(ctx, c.config.Url, xmlRequest, utils.Config{
		Timeout:    timeout,
		MaxRetries: maxRetries,
	}, headers)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result, err := statuscheck.ParseStatusCheckSOAP(string(responseData))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *CBECoreAPI) FundTransferVerify(ctx context.Context, param FundTransferVerifyParam) (*FundTransferVerifyResult, error) {
	params := fundtransferverify.Params{
		Username:            c.config.Username,
		Password:            c.config.Password,
		DebitAccountNumber:  param.DebitAccountNumber,
		DebitCurrency:       param.DebitCurrency,
		DebitAmount:         param.DebitAmount,
		DebitReference:      param.DebitReference,
		CreditReference:     param.CreditReference,
		CreditAccountNumber: param.CreditAccountNumber,
		PaymentDetails:      param.PaymentDetails,
		ClientReference:     param.ClientReference,
		ServiceCode:         param.ServiceCode,
		CustomerSegment:     param.CustomerSegment,
		ChannelType:         param.ChannelType,
		CreditCurrency:      param.CreditCurrency,
		CreditAmount:        param.CreditAmount,
		SuperappUserCode:    param.SuperappUserCode,
	}

	xmlRequest := fundtransferverify.NewFundTransferVerify(params)
	headers := map[string]string{
		Key: Value,
	}

	resp, err := utils.DoPost(ctx, c.config.Url, xmlRequest, utils.Config{
		Timeout:    timeout,
		MaxRetries: maxRetries,
	}, headers)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result, err := fundtransferverify.ParseFundTransferVerifySOAP(string(responseData))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *CBECoreAPI) CustomerLimitFetchByCustomerNumber(ctx context.Context, param CustomerLimitFetchByCIFParam) (*CustomerLimitFetchByCIFResult, error) {
	params := customerlimitfetchbycif.Params{
		Username:       c.config.Username,
		Password:       c.config.Password,
		CustomerNumber: param.CustomerNumber,
	}

	xmlRequest := customerlimitfetchbycif.NewCustomerLimitFetchByCIF(params)
	headers := map[string]string{
		Key: Value,
	}

	resp, err := utils.DoPost(ctx, c.config.Url, xmlRequest, utils.Config{
		Timeout:    timeout,
		MaxRetries: maxRetries,
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
func (c *CBECoreAPI) CustomerLimitFetchByService(ctx context.Context, param CustomerLimitFetchByServiceParam) (*CustomerLimitFetchByServiceResult, error) {
	params := customerlimitfetchbyservice.Params{
		Username:    c.config.Username,
		Password:    c.config.Password,
		ServiceCode: param.ServiceCode,
	}
	xmlRequest := customerlimitfetchbyservice.NewCustomerLimitFetchByService(params)
	headers := map[string]string{
		Key: Value,
	}

	resp, err := utils.DoPost(ctx, c.config.Url, xmlRequest, utils.Config{
		Timeout:    timeout,
		MaxRetries: maxRetries,
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
func (c *CBECoreAPI) CustomerLimitAmendByCustomerNumber(ctx context.Context, param CustomerLimitAmendByCIFParam) (*CustomerLimitAmendByCIFResult, error) {
	params := customerlimitamendbycif.Params{
		Username:       c.config.Username,
		Password:       c.config.Password,
		CustomerNumber: param.CustomerNumber,
		ChannelLimit:   param.ChannelLimit,
	}
	xmlRequest := customerlimitamendbycif.NewCustomerLimitAmendByCIF(params)
	fmt.Println("xmlRequest: ", xmlRequest)
	headers := map[string]string{
		Key: Value,
	}

	resp, err := utils.DoPost(ctx, c.config.Url, xmlRequest, utils.Config{
		Timeout:    timeout,
		MaxRetries: maxRetries,
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
	fmt.Println("result:", result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *CBECoreAPI) SplitPayment(ctx context.Context, param SplitPaymentParam) (*SplitPaymentResult, error) {
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
		Key: Value,
	}
	resp, err := utils.DoPost(ctx, c.config.Url, xmlRequest, utils.Config{
		Timeout:    timeout,
		MaxRetries: maxRetries,
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

func (c *CBECoreAPI) ServiceLimit(ctx context.Context, param ServiceLimitParam) (*ServiceLimitResult, error) {
	params := servicelimit.Params{
		Username: c.config.Username,
		Password: c.config.Password,
	}
	xmlRequest := servicelimit.NewServiceLimit(params)
	headers := map[string]string{
		Key: Value,
	}
	resp, err := utils.DoPost(ctx, c.config.Url, xmlRequest, utils.Config{
		Timeout:    timeout,
		MaxRetries: maxRetries,
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

func (c *CBECoreAPI) ExchangeRates(ctx context.Context) (*ExchangeRatesResult, error) {
	params := exchangerate.Params{
		Username: c.config.Username,
		Password: c.config.Password,
	}
	xmlRequest := exchangerate.NewExchangeRate(params)
	headers := map[string]string{
		Key: Value,
	}
	resp, err := utils.DoPost(ctx, c.config.Url, xmlRequest, utils.Config{
		Timeout:    timeout,
		MaxRetries: maxRetries,
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

func (c *CBECoreAPI) PhoneLookup(ctx context.Context, param PhoneLookupParam) (*PhoneLookupResult, error) {
	params := phonelookup.Params{
		Username:    c.config.Username,
		Password:    c.config.Password,
		PhoneNumber: param.PhoneNumber,
	}
	xmlRequest := phonelookup.NewPhoneLookup(params)
	headers := map[string]string{
		Key: Value,
	}
	resp, err := utils.DoPost(ctx, c.config.Url, xmlRequest, utils.Config{
		Timeout:    timeout,
		MaxRetries: maxRetries,
	}, headers)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseData, err := utils.ReadResponseBody(resp, "phone lookup")
	if err != nil {
		return nil, err
	}
	result, err := phonelookup.ParsePhoneLookupSOAP(string(responseData))
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *CBECoreAPI) RevertFundTransfer(ctx context.Context, param RevertFundTransferParam) (*RevertFundTransferResult, error) {
	params := revertfundtransfer.Params{
		Username:      c.config.Username,
		Password:      c.config.Password,
		TransactionID: param.TransactionID,
	}

	xmlRequest := revertfundtransfer.NewRevertFundTransfer(params)
	headers := map[string]string{
		Key: Value,
	}
	resp, err := utils.DoPost(ctx, c.config.Url, xmlRequest, utils.Config{
		Timeout:    timeout,
		MaxRetries: maxRetries,
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

func (c *CBECoreAPI) AccountLookup(ctx context.Context, param AccountLookupParam) (*AccountLookupResult, error) {
	params := accountlookup.Params{
		Username:       c.config.Username,
		Password:       c.config.Password,
		AccountNumber:  param.AccountNumber,
		CustomerNumber: param.CustomerNumber,
	}

	xmlRequest := accountlookup.NewAccountLookup(params)
	headers := map[string]string{
		Key: Value,
	}
	resp, err := utils.DoPost(ctx, c.config.Url, xmlRequest, utils.Config{
		Timeout:    timeout,
		MaxRetries: maxRetries,
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

func (c *CBECoreAPI) AccountList(ctx context.Context, param AccountListParam) (*AccountListResult, error) {
	params := accountlist.Params{
		Username:      c.config.Username,
		Password:      c.config.Password,
		ColumnName:    param.ColumnName,
		CriteriaValue: param.CriteriaValue,
	}

	xmlRequest := accountlist.NewAccountList(params)
	headers := map[string]string{
		Key: Value,
	}
	resp, err := utils.DoPost(ctx, c.config.Url, xmlRequest, utils.Config{
		Timeout:    timeout,
		MaxRetries: maxRetries,
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

func (c *CBECoreAPI) FundTransfer(ctx context.Context, param FundTransferParam) (*FundTransferResult, error) {
	var messages []string
	if strings.TrimSpace(param.ServiceCode) == "" {
		messages = append(messages, "Service Code Not Found!")
	}
	if strings.TrimSpace(param.CustomerSegment) == "" {
		messages = append(messages, "Customer Segnment Not Found!")
	}
	if strings.TrimSpace(param.ChannelType) == "" {
		messages = append(messages, "Channele Type Not Found!")
	}
	if len(messages) > 0 {
		return &FundTransferResult{
			Success:  false,
			Messages: messages,
		}, nil
	}

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
		CustomerSegment:     param.CustomerSegment,
		ChannelType:         param.ChannelType,
		SuperappUserCode:    param.SuperappUserCode,
		Meta:                param.Meta,
		IsFraudCheckEnabled: param.IsFraudCheckEnabled,
		Key:                 param.Key,
		BranchCode:          param.BranchCode,
	}

	go func(meta frauddetection.FraudAPIPayload) {
		if param.IsFraudCheckEnabled {
			fraud := frauddetection.NewFraudAPI(
				c.config.FraudAPIConfig.Authorization,
				c.config.FraudAPIConfig.ForwardHost,
				c.config.FraudAPIConfig.Url,
			)

			meta.TransactionType = "ACSA"

			layout := "2006-01-02 15:04:05"
			meta.TransactionTimestamp = time.Now().Add(3 * time.Hour).UTC().Format(layout)
			fraud.Call(meta)
			// if err != nil {
			// 	return nil, fmt.Errorf("fraud detection call failed: %w", err)
			// }

			// if response.Result != "approved" {
			// 	return &fundtransfer.FundTransferResult{
			// 		Success:  false,
			// 		Messages: []string{"transaction blocked by fraud detection"},
			// 	}, nil
			// }
		}
	}(param.Meta)

	xmlRequest := fundtransfer.NewFundTransfer(params)
	headers := map[string]string{
		Key: Value,
	}

	resp, err := utils.DoPost(ctx, c.config.Url, xmlRequest, utils.Config{
		Timeout:    timeout,
		MaxRetries: maxRetries,
	}, headers)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		if c.config.RedisClient != nil {
			key := c.redisKey("failed_ft_count", param.Key)
			c.config.RedisClient.Incr(ctx, key)
			if ttl, err := c.config.RedisClient.TTL(ctx, key).Result(); err == nil && ttl < 0 {
				c.config.RedisClient.Expire(ctx, key, c.config.RedisTTL)
			}
		}
		return nil, err
	}

	result, err := fundtransfer.ParseFundTransferSOAP(string(responseData))
	if err != nil {
		return nil, err
	}

	// increment on every success transaction
	if c.config.RedisClient != nil {
		key := c.redisKey("success_ft_count", param.Key)
		c.config.RedisClient.Incr(ctx, key)
		if ttl, err := c.config.RedisClient.TTL(ctx, key).Result(); err == nil && ttl < 0 {
			c.config.RedisClient.Expire(ctx, key, c.config.RedisTTL)
		}
	}

	var surveyResults survey.SurveyResult
	if param.IsSurveyEnabled || len(param.SurveyRules) > 0 {
		surveyResults = c.initSurvey(
			ctx,
			c.redisKey("success_ft_count", param.Key),
			param.BranchCode,
			param.CustomerSegment,
			param.DebitAmount,
			param.DebitCurrency,
			c.mapSurveyResult(param.SurveyRules),
		)

		result.SurveyResults = &surveyResults
		if result.SurveyResults.Result || result.SurveyResults.Url != nil {
			temp := GenerateSurveyHashURL(*result.SurveyResults.Url, param.Meta.CustomerName)
			result.SurveyResults.Url = &temp
		}
	}

	return result, nil
}

func (c *CBECoreAPI) FundTransferCheck(ctx context.Context, param FundTransferCheckParam) (*FundTransferCheckResult, error) {
	params := fundtransfercheck.Params{
		Username: c.config.Username,
		Password: c.config.Password,
		FTNumber: param.FTNumber,
	}

	xmlRequest := fundtransfercheck.NewFundTransferCheck(params)
	headers := map[string]string{
		Key: Value,
	}
	resp, err := utils.DoPost(ctx, c.config.Url, xmlRequest, utils.Config{
		Timeout:    timeout,
		MaxRetries: maxRetries,
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

func (c *CBECoreAPI) ListLockedAmount(ctx context.Context, param ListLockedAmountParam) (*ListLockedAmountResult, error) {
	params := lockedamountlist.Params{
		Username:      c.config.Username,
		Password:      c.config.Password,
		AccountNumber: param.AccountNumber,
	}

	xmlRequest := lockedamountlist.NewListLockedAmount(params)
	headers := map[string]string{
		Key: Value,
	}
	resp, err := utils.DoPost(ctx, c.config.Url, xmlRequest, utils.Config{
		Timeout:    timeout,
		MaxRetries: maxRetries,
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

func (c *CBECoreAPI) CreateLockedAmount(ctx context.Context, param CreateLockedAmountParam) (*CreateLockedAmountResult, error) {
	params := lockedamountcreate.Params{
		Username:      c.config.Username,
		Password:      c.config.Password,
		AccountNumber: param.AccountNumber,
		Description:   param.Description,
		From:          param.From,
		To:            param.To,
		LockedAmount:  param.LockedAmount,
		UserID:        param.UserID,
		ServiceCode:   param.ServiceCode,
	}

	fmt.Println("** [1]create locked amount params **", params)

	xmlRequest := lockedamountcreate.NewCreateLockedAmount(params)
	headers := map[string]string{
		Key: Value,
	}
	fmt.Println("** [2]create locked amount request **", xmlRequest)
	resp, err := utils.DoPost(ctx, c.config.Url, xmlRequest, utils.Config{
		Timeout:    timeout,
		MaxRetries: maxRetries,
	}, headers)

	if err != nil {
		fmt.Println("** [3]create locked amount error **", err)
		return nil, err
	}
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("** [4]create locked amount error **", err)
		return nil, err
	}

	result, err := lockedamountcreate.ParseCreateLockedAmountSOAP(string(responseData))
	fmt.Println("** [5]create locked amount response **", result)
	if err != nil {
		fmt.Println("** [6]create locked amount error **", err)
		return nil, err
	}

	return result, nil
}

func (c *CBECoreAPI) LockedAmountFT(ctx context.Context, param LockedAmountFTParam) (*LockedAmountFTResult, error) {
	params := lockedamountft.Params{
		Username:            c.config.Username,
		Password:            c.config.Password,
		CreditCurrency:      param.CreditCurrency,
		CreditAccountNumber: param.CreditAccountNumber,
		CreditAmount:        param.CreditAmount,
		DebitAmount:         param.DebitAmount,
		DebitAccountNumber:  param.DebitAccountNumber,
		DebitCurrency:       param.DebitCurrency,
		CreditReference:     param.CreditReference,
		DebitReference:      param.DebitReference,
		ClientReference:     param.ClientReference,
		ServiceCode:         param.ServiceCode,
		LockID:              param.LockID,
		PaymentDetails:      param.PaymentDetails,
		CustomerRole:        param.CustomerRole,
		ChannelType:         param.ChannelType,
		SuperappUserCode:    param.SuperappUserCode,
	}

	xmlRequest := lockedamountft.NewLockedAmountFt(params)
	headers := map[string]string{
		Key: Value,
	}
	resp, err := utils.DoPost(ctx, c.config.Url, xmlRequest, utils.Config{
		Timeout:    timeout,
		MaxRetries: maxRetries,
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

func (c *CBECoreAPI) ReleaseLockedAmount(ctx context.Context, param ReleaseLockedAmountParam) (*ReleaseLockedAmountResult, error) {
	params := lockedamountrelease.Params{
		Username:      c.config.Username,
		Password:      c.config.Password,
		TransactionID: param.TransactionID,
	}

	xmlRequest := lockedamountrelease.NewReleaseLockedAmount(params)
	headers := map[string]string{
		Key: Value,
	}
	resp, err := utils.DoPost(ctx, c.config.Url, xmlRequest, utils.Config{
		Timeout:    timeout,
		MaxRetries: maxRetries,
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

func (c *CBECoreAPI) CancleStandingOrder(ctx context.Context, param CancleStandingOrderParam) (*CancelStandingOrderResult, error) {
	params := standingordercancel.Params{
		Username:      c.config.Username,
		Password:      c.config.Password,
		AccountNumber: param.AccountNumber,
		OrderId:       param.OrderId,
	}

	xmlRequest := standingordercancel.NewCancleStandingOrder(params)
	headers := map[string]string{
		Key: Value,
	}
	resp, err := utils.DoPost(ctx, c.config.Url, xmlRequest, utils.Config{
		Timeout:    timeout,
		MaxRetries: maxRetries,
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

func (c *CBECoreAPI) UpdateStandingOrder(ctx context.Context, param UpdateStandingOrderParam) (*UpdateStandingOrderResult, error) {
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
		Key: Value,
	}
	resp, err := utils.DoPost(ctx, c.config.Url, xmlRequest, utils.Config{
		Timeout:    timeout,
		MaxRetries: maxRetries,
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

func (c *CBECoreAPI) ListStandingOrderHistory(ctx context.Context, param ListStandingOrderHistoryParam) (*ListStandingOrderHistoryResult, error) {
	params := standingorderlisthistory.Params{
		Username:      c.config.Username,
		Password:      c.config.Password,
		AccountNumber: param.AccountNumber,
	}
	xmlRequest := standingorderlisthistory.NewListStandingOrderHistory(params)
	headers := map[string]string{
		Key: Value,
	}
	resp, err := utils.DoPost(ctx, c.config.Url, xmlRequest, utils.Config{
		Timeout:    timeout,
		MaxRetries: maxRetries,
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
func (c *CBECoreAPI) CreateStandingOrder(ctx context.Context, param CreateStandingOrderParam) (*CreateStandingOrderResult, error) {
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
		Key: Value,
	}
	resp, err := utils.DoPost(ctx, c.config.Url, xmlRequest, utils.Config{
		Timeout:    timeout,
		MaxRetries: maxRetries,
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

func (c *CBECoreAPI) ListStandingOrder(ctx context.Context, param ListStandingOrderParam) (*ListStandingOrderResult, error) {
	params := standingorderlist.Params{
		Username:      c.config.Username,
		Password:      c.config.Password,
		AccountNumber: param.AccountNumber,
	}

	xmlRequest := standingorderlist.NewListStandingOrder(params)
	headers := map[string]string{
		Key: Value,
	}
	resp, err := utils.DoPost(ctx, c.config.Url, xmlRequest, utils.Config{
		Timeout:    timeout,
		MaxRetries: maxRetries,
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

func (c *CBECoreAPI) MiniStatementByLimit(ctx context.Context, param MiniStatementByLimitParams) (*MiniStatementByLimitResult, error) {
	params := ministatementbylimit.Params{
		Username:            c.config.Username,
		Password:            c.config.Password,
		AccountNumber:       param.AccountNumber,
		NumberOfTransaction: param.NumberOfTransaction,
	}

	xmlRequest := ministatementbylimit.NewMiniStatementByLimit(params)
	headers := map[string]string{
		Key: Value,
	}
	resp, err := utils.DoPost(ctx, c.config.Url, xmlRequest, utils.Config{
		Timeout:    timeout,
		MaxRetries: maxRetries,
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

func (c *CBECoreAPI) MiniStatementByDateRange(ctx context.Context, param MiniStatementByDateRangeParam) (*MiniStatementByDateRangeResult, error) {
	params := ministatementbydaterange.Params{
		Username:      c.config.Username,
		Password:      c.config.Password,
		AccountNumber: param.AccountNumber,
		From:          param.From,
		To:            param.To,
	}

	xmlRequest := ministatementbydaterange.NewMiniStatementByDateRange(params)
	headers := map[string]string{
		Key: Value,
	}

	resp, err := utils.DoPost(ctx, c.config.Url, xmlRequest, utils.Config{
		Timeout:    timeout,
		MaxRetries: maxRetries,
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
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *CBECoreAPI) CustomerLookup(ctx context.Context, param CustomerLookupParam) (*CustomerLookupResult, error) {
	params := customerlookup.Params{
		Username:           c.config.Username,
		Password:           c.config.Password,
		CustomerIdentifier: param.CustomerIdentifier,
	}
	xmlRequest := customerlookup.NewCustomerLookup(params)
	headers := map[string]string{
		Key: Value,
	}

	resp, err := utils.DoPost(ctx, c.config.Url, xmlRequest, utils.Config{
		Timeout:    timeout,
		MaxRetries: maxRetries,
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
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *CBECoreAPI) VerifyAML(ctx context.Context, param VerifyAMLParam) (*VerifyAMLResult, error) {
	params := verifyaml.Param{
		Password:       c.config.Password,
		UserName:       c.config.Username,
		CustomerNumber: param.CustomerNumber,
	}
	xmlRequest := verifyaml.NewVerifyAML(params)
	headers := map[string]string{
		Key: Value,
	}

	resp, err := utils.DoPost(ctx, c.config.Url, xmlRequest, utils.Config{
		Timeout:    timeout,
		MaxRetries: maxRetries,
	}, headers)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result, err := verifyaml.ParseVerifyAMLSOAP(string(responseData))
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *CBECoreAPI) CardReplace(ctx context.Context, param CardReplaceParam) (*CardReplaceResult, error) {
	params := cardreplace.Params{
		Username:      c.config.Username,
		Password:      c.config.Password,
		AccountNumber: param.AccountNumber,
	}
	xmlRequest := cardreplace.NewCardReplace(params)
	headers := map[string]string{
		Key: Value,
	}
	resp, err := utils.DoPost(ctx, c.config.Url, xmlRequest, utils.Config{
		Timeout:    timeout,
		MaxRetries: maxRetries,
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

func (c *CBECoreAPI) CardRequest(ctx context.Context, param CardRequestParam) (*CardRequestResult, error) {
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
		Key: Value,
	}
	resp, err := utils.DoPost(ctx, c.config.Url, xmlRequest, utils.Config{
		Timeout:    timeout,
		MaxRetries: maxRetries,
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

func NewCBECoreAPI(param *CBECoreCredential) CBECoreAPIInterface {
	config := internal.SetConfig(
		param.Username,
		param.Password,
		param.Url,
		param.FraudAPICredential.Authorization,
		param.FraudAPICredential.Url,
		param.FraudAPICredential.ForwardHost,
		param.RedisClient,
	)
	return &CBECoreAPI{
		config: config,
	}
}
