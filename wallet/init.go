package wallet

import (
	agent_accountlookup "github.com/hugokessem/coreio/lib/wallet/cbe_birr/agent/account_lookup"
	agent_fundtransfer "github.com/hugokessem/coreio/lib/wallet/cbe_birr/agent/fund_transfer"
	accountlookup "github.com/hugokessem/coreio/lib/wallet/cbe_birr/customer/account_lookup"
	fundtransfer "github.com/hugokessem/coreio/lib/wallet/cbe_birr/customer/fund_transfer"
	"github.com/hugokessem/coreio/wallet/internal"
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

type CustomerFundTransferParams = fundtransfer.CustomerFundTransferParams
type CustomerFundTransferResult = fundtransfer.CustomerFundTransferResult
type CustomerAccountLookupParam = accountlookup.CustomerAccountLookupParams
type CustomerAccountLookupResult = accountlookup.CustomerAccountLookupResult

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
