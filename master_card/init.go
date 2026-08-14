package mastercard

import (
	customersupport "github.com/hugokessem/coreio/lib/master_card/customer_support"
	"github.com/hugokessem/coreio/master_card/internal"
)

type WorkerLoginParam = customersupport.WorkerLoginParam
type WorkerLoginResult = customersupport.WorkerLoginResult

type IssueProductsResult = customersupport.IssueProductsResult

type IssuesListParam = customersupport.IssuesListParam
type IssuesListResult = customersupport.IssuesListResult

type ViewMessageParam = customersupport.ViewMessageParam
type ViewMessageResult = customersupport.ViewMessageResult

type SendMessageParam = customersupport.SendMessageParam
type SendMessageResult = customersupport.SendMessageResult

type ExportActivityLogResult = customersupport.ExportActivityLogResult
type ActivityLog = customersupport.ActivityLog

type CreateIssueParam = customersupport.CreateIssueParam
type CreateIssueResult = customersupport.CreateIssueResult

type MasterCardAPIInterface = customersupport.CustomerSupportAPIInterface
type MasterCardCredential = customersupport.CustomerSupportCredential

func NewMasterCardAPI(param *MasterCardCredential) MasterCardAPIInterface {
	config := internal.SetConfig(param.AuthURL, param.BaseURL, param.CustomerKey)
	return customersupport.NewCustomerSupportAPI(customersupport.CustomerSupportCredential{
		AuthURL:     config.AuthURL,
		BaseURL:     config.BaseURL,
		CustomerKey: config.CustomerKey,
		AccessToken: param.AccessToken,
	})
}
