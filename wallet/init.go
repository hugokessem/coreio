// coreio target: wallet/init.go
// Replace wallet/init.go after copying the 4 lib/wallet/cbe_birr parser files.
package wallet

import (
	"context"
	"io"
	"strings"
	"time"

	agent_accountlookup "github.com/hugokessem/coreio/lib/wallet/cbe_birr/agent/account_lookup"
	agent_fundtransfer "github.com/hugokessem/coreio/lib/wallet/cbe_birr/agent/fund_transfer"
	cutomer_accountlookup "github.com/hugokessem/coreio/lib/wallet/cbe_birr/customer/account_lookup"
	cutomer_fundtransfer "github.com/hugokessem/coreio/lib/wallet/cbe_birr/customer/fund_transfer"
	"github.com/hugokessem/coreio/utils"
	"github.com/hugokessem/coreio/wallet/internal"
)

type WalletCredentials struct {
	Url                         string
	Password                    string
	Authorization               string // "Bearer <token>" from OAuth
	IIBAuthorization            string
	SecurityCredential          string // KYC credential
	ThirdPartyIdentifier        string
	KYCThirdPartyIdentifier     string
	PaymentThirdPartyIdentifier string
	PaymentSecurityCredential   string
	ShortCode                   string
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
	AgentAccountLookup(ctx context.Context, param AgentAccountLookupParams) (*AgentAccountLookupResult, error)
	AgentFundTransfer(ctx context.Context, param AgentFundTransferParams) (*AgentFundTransferResult, error)
	CustomerAccountLookup(ctx context.Context, param CustomerAccountLookupParam) (*CustomerAccountLookupResult, error)
	CustomerFundTransfer(ctx context.Context, param CustomerFundTransferParams) (*CustomerFundTransferResult, error)
}

func NewWalletAPI(param WalletCredentials) WalletInterface {
	config := internal.SetConfig(
		param.Url,
		param.Password,
		param.Authorization,
		param.IIBAuthorization,
		param.SecurityCredential,
		param.ThirdPartyIdentifier,
		param.KYCThirdPartyIdentifier,
		param.PaymentThirdPartyIdentifier,
		param.PaymentSecurityCredential,
		param.ShortCode,
	)
	return &WalletAPI{config: config}
}

const (
	timeout    = 120 * time.Second
	maxRetries = 1
)

func (w *WalletAPI) postHeaders() map[string]string {
	return map[string]string{
		"Content-Type":      "application/xml",
		"iib_authorization": w.config.IIBAuthorization,
		"Authorization":     w.config.Authorization,
	}
}

func (w *WalletAPI) doPost(ctx context.Context, xmlRequest string) ([]byte, error) {
	cfg := utils.Config{MaxRetries: maxRetries, Timeout: timeout}
	resp, err := utils.DoPost(ctx, w.config.Url, xmlRequest, cfg, w.postHeaders())
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

// AgentAccountLookup implements WalletInterface.
func (w *WalletAPI) AgentAccountLookup(ctx context.Context, param AgentAccountLookupParams) (*AgentAccountLookupResult, error) {
	_ = param.OriginalConverstationIdentifier // always generate unique ID (platform/cbebirr)
	originatorID := agent_accountlookup.GenerateOriginatorConversationID()

	internalParams := agent_accountlookup.Params{
		OriginalConverstationIdentifier: originatorID,
		Timestamp:                       param.Timestamp,
		PhoneNumber:                     param.PhoneNumber, // agent code
		ThirdPartyIdentifier:            w.config.ThirdPartyIdentifier,
		Password:                        w.config.Password,
		InitiatorIdentifier:             w.config.KYCThirdPartyIdentifier,
		SecurityCredential:              w.config.SecurityCredential,
	}

	responseData, err := w.doPost(ctx, agent_accountlookup.NewAgentAccountLookup(internalParams))
	if err != nil {
		return nil, err
	}
	return agent_accountlookup.ParseAgentLookupSOAP(string(responseData))
}

// AgentFundTransfer implements WalletInterface.
func (w *WalletAPI) AgentFundTransfer(ctx context.Context, param AgentFundTransferParams) (*AgentFundTransferResult, error) {
	currency := strings.TrimSpace(param.Currency)
	if currency == "" {
		currency = "ETB"
	}

	internalParams := agent_fundtransfer.Params{
		FTNumber:               param.FTNumber,
		Timestamp:              param.Timestamp,
		PrimaryParty:           param.PrimaryParty,
		ReceiverParty:          param.ReceiverParty,
		Amount:                 param.Amount,
		Currency:               currency,
		DebitAccountNumber:     param.DebitAccountNumber,
		DebitAccountHolderName: param.DebitAccountHolderName,
		ThirdPartyIdentifier:   w.config.ThirdPartyIdentifier,
		Password:               w.config.Password,
		InitiatorIdentifier:    w.config.PaymentThirdPartyIdentifier,
		SecurityCredential:     w.config.PaymentSecurityCredential,
	}

	responseData, err := w.doPost(ctx, agent_fundtransfer.NewAgentFundTransfer(internalParams))
	if err != nil {
		return nil, err
	}
	return agent_fundtransfer.ParseAgentFundTransfer(string(responseData))
}

// CustomerAccountLookup implements WalletInterface.
func (w *WalletAPI) CustomerAccountLookup(ctx context.Context, param CustomerAccountLookupParam) (*CustomerAccountLookupResult, error) {
	_ = param.OriginalConverstationIdentifier // always generate unique ID (platform/cbebirr)
	originatorID := cutomer_accountlookup.GenerateOriginatorConversationID()

	internalParams := cutomer_accountlookup.Params{
		OriginalConverstationIdentifier: originatorID,
		Timestamp:                       param.Timestamp,
		PhoneNumber:                     param.PhoneNumber,
		ThirdPartyIdentifier:            w.config.ThirdPartyIdentifier,
		Password:                        w.config.Password,
		InitiatorIdentifier:             w.config.KYCThirdPartyIdentifier,
		SecurityCredential:              w.config.SecurityCredential,
	}

	responseData, err := w.doPost(ctx, cutomer_accountlookup.NewCustomerAccountLookup(internalParams))
	if err != nil {
		return nil, err
	}
	return cutomer_accountlookup.ParseCustomerLookupSOAP(string(responseData), param.PhoneNumber)
}

// CustomerFundTransfer implements WalletInterface.
func (w *WalletAPI) CustomerFundTransfer(ctx context.Context, param CustomerFundTransferParams) (*CustomerFundTransferResult, error) {
	currency := strings.TrimSpace(param.Currency)
	if currency == "" {
		currency = "ETB"
	}
	shortCode := strings.TrimSpace(param.PrimaryParty)
	if shortCode == "" {
		shortCode = w.config.ShortCode
	}

	internalParams := cutomer_fundtransfer.Params{
		FTNumber:               param.FTNumber,
		Timestamp:              param.Timestamp,
		ShortCode:              shortCode,
		ReceiverParty:          param.ReceiverParty,
		Amount:                 param.Amount,
		Currency:               currency,
		DebitAccountNumber:     param.DebitAccountNumber,
		DebitAccountHolderName: param.DebitAccountHolderName,
		ThirdPartyIdentifier:   w.config.ThirdPartyIdentifier,
		Password:               w.config.Password,
		InitiatorIdentifier:    w.config.PaymentThirdPartyIdentifier,
		SecurityCredential:     w.config.PaymentSecurityCredential,
	}

	responseData, err := w.doPost(ctx, cutomer_fundtransfer.NewCustomerFundTransfer(internalParams))
	if err != nil {
		return nil, err
	}
	return cutomer_fundtransfer.ParserCustomreFundTransfer(string(responseData))
}
