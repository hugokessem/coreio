package wallet

import (
	"io"
	"time"

	agent_accountlookup "github.com/hugokessem/coreio/lib/wallet/cbe_birr/agent/account_lookup"
	agent_fundtransfer "github.com/hugokessem/coreio/lib/wallet/cbe_birr/agent/fund_transfer"
	cutomer_accountlookup "github.com/hugokessem/coreio/lib/wallet/cbe_birr/customer/account_lookup"
	cutomer_fundtransfer "github.com/hugokessem/coreio/lib/wallet/cbe_birr/customer/fund_transfer"
	"github.com/hugokessem/coreio/utils"
	internal "github.com/hugokessem/coreio/wallet/internal"
)

type WalletCredentials struct {
	Url                  string
	Password             string
	Authorization        string
	IIBAuthorization     string
	SecurityCredential   string
	ThirdPartyIdentifier string
}

type WalletAPI struct {
	config *internal.Config
}

type AgentFundTransferParams = agent_fundtransfer.AgentFundTransferParams
type AgentFundTransferResult = agent_fundtransfer.AgentFundTransferResult
type AgentAccountLookupParams = agent_accountlookup.AgentAccountLookupParams
type AgentAccountLookupResult = agent_accountlookup.AgentAccountLookupResult

type CustomerFundTransferParams = cutomer_fundtransfer.CustomerFundTransferParams
type CustomerFundTransferResult = cutomer_fundtransfer.CustomerFundTransferResult
type CustomerAccountLookupParam = cutomer_accountlookup.CustomerAccountLookupParams
type CustomerAccountLookupResult = cutomer_accountlookup.CustomerAccountLookupResult

type WalletInterface interface {
	AgentAccountLookup(param AgentAccountLookupParams) (*AgentAccountLookupResult, error)
	AgentFundTransfer(param AgentFundTransferParams) (*AgentFundTransferResult, error)
	CustomerAccountLookup(param CustomerAccountLookupParam) (*CustomerAccountLookupResult, error)
	CustomerFundTransfer(param CustomerFundTransferParams) (*CustomerFundTransferResult, error)
}

func NewWalletAPI(param WalletCredentials) WalletInterface {
	config := internal.SetConfig(
		param.Url,
		param.Password,
		param.Authorization,
		param.IIBAuthorization,
		param.SecurityCredential,
		param.ThirdPartyIdentifier,
	)

	return &WalletAPI{
		config: config,
	}
}

// AgentAccountLookup implements WalletInterface.
func (w *WalletAPI) AgentAccountLookup(param AgentAccountLookupParams) (*AgentAccountLookupResult, error) {
	config := utils.Config{
		MaxRetries: 6,
		Timeout:    30 * time.Second,
	}

	// Convert public params to internal params with config values
	internalParams := agent_accountlookup.Params{
		OriginalConverstationIdentifier: param.OriginalConverstationIdentifier,
		Timestamp:                       param.Timestamp,
		PhoneNumber:                     param.PhoneNumber,
		ThirdPartyIdentifier:            w.config.ThirdPartyIdentifier,
		Password:                        w.config.Password,
		SecurityCredential:              w.config.SecurityCredential,
	}

	xmlRequest := agent_accountlookup.NewAgentAccountLookup(internalParams)
	headers := map[string]string{
		"Content-Type":      "application/xml",
		"iib_authorization": w.config.IIBAuthorization,
		"Authorization":     w.config.Authorization,
	}

	resp, err := utils.DoPostWithRetry(w.config.Url, xmlRequest, config, headers)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result, err := agent_accountlookup.ParseAgentLookupSOAP(string(responseData))
	if err != nil {
		return nil, err
	}

	return result, nil
}

// AgentFundTransfer implements WalletInterface.
func (w *WalletAPI) AgentFundTransfer(param AgentFundTransferParams) (*AgentFundTransferResult, error) {
	config := utils.Config{
		MaxRetries: 6,
		Timeout:    30 * time.Second,
	}

	// Convert public params to internal params with config values
	internalParams := agent_fundtransfer.Params{
		FTNumber:               param.FTNumber,
		Timestamp:              param.Timestamp,
		PrimaryParty:           param.PrimaryParty,
		ReceiverParty:          param.ReceiverParty,
		Amount:                 param.Amount,
		Currency:               param.Currency,
		Narative:               param.Narative,
		DebitAccountNumber:     param.DebitAccountNumber,
		DebitAccountHolderName: param.DebitAccountHolderName,
		ThirdPartyIdentifier:   w.config.ThirdPartyIdentifier,
		Password:               w.config.Password,
		SecurityCredential:     w.config.SecurityCredential,
	}

	xmlRequest := agent_fundtransfer.NewAgentFundTransfer(internalParams)
	headers := map[string]string{
		"Content-Type":      "application/xml",
		"iib_authorization": w.config.IIBAuthorization,
		"Authorization":     w.config.Authorization,
	}

	resp, err := utils.DoPostWithRetry(w.config.Url, xmlRequest, config, headers)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result, err := agent_fundtransfer.ParseAgentFundTransfer(string(responseData))
	if err != nil {
		return nil, err
	}

	return result, nil
}

// CustomerAccountLookup implements WalletInterface.
func (w *WalletAPI) CustomerAccountLookup(param CustomerAccountLookupParam) (*CustomerAccountLookupResult, error) {
	config := utils.Config{
		MaxRetries: 6,
		Timeout:    30 * time.Second,
	}

	// Convert public params to internal params with config values
	internalParams := cutomer_accountlookup.Params{
		OriginalConverstationIdentifier: param.OriginalConverstationIdentifier,
		Timestamp:                       param.Timestamp,
		PhoneNumber:                     param.PhoneNumber,
		ThirdPartyIdentifier:            w.config.ThirdPartyIdentifier,
		Password:                        w.config.Password,
		SecurityCredential:              w.config.SecurityCredential,
	}

	xmlRequest := cutomer_accountlookup.NewCustomerAccountLookup(internalParams)
	headers := map[string]string{
		"Content-Type":      "application/xml",
		"iib_authorization": w.config.IIBAuthorization,
		"Authorization":     w.config.Authorization,
	}

	resp, err := utils.DoPostWithRetry(w.config.Url, xmlRequest, config, headers)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result, err := cutomer_accountlookup.ParseCustomerLookupSOAP(string(responseData))
	if err != nil {
		return nil, err
	}

	return result, nil
}

// CustomerFundTransfer implements WalletInterface.
func (w *WalletAPI) CustomerFundTransfer(param CustomerFundTransferParams) (*CustomerFundTransferResult, error) {
	config := utils.Config{
		MaxRetries: 6,
		Timeout:    30 * time.Second,
	}

	// Convert public params to internal params with config values
	internalParams := cutomer_fundtransfer.Params{
		FTNumber:               param.FTNumber,
		Timestamp:              param.Timestamp,
		PrimaryParty:           param.PrimaryParty,
		ReceiverParty:          param.ReceiverParty,
		Amount:                 param.Amount,
		Currency:               param.Currency,
		Narative:               param.Narative,
		DebitAccountNumber:     param.DebitAccountNumber,
		DebitAccountHolderName: param.DebitAccountHolderName,
		ThirdPartyIdentifier:   w.config.ThirdPartyIdentifier,
		Password:               w.config.Password,
		SecurityCredential:     w.config.SecurityCredential,
	}

	xmlRequest := cutomer_fundtransfer.NewCustomerFundTransfer(internalParams)
	headers := map[string]string{
		"Content-Type":      "application/xml",
		"iib_authorization": w.config.IIBAuthorization,
		"Authorization":     w.config.Authorization,
	}

	resp, err := utils.DoPostWithRetry(w.config.Url, xmlRequest, config, headers)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result, err := cutomer_fundtransfer.ParserCustomreFundTransfer(string(responseData))
	if err != nil {
		return nil, err
	}

	return result, nil
}
