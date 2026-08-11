package mastercard

import (
	"context"

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

type MasterCardAPIInterface interface {
	WorkerLogin(ctx context.Context, param WorkerLoginParam) (*WorkerLoginResult, error)
	IssueProducts(ctx context.Context) (*IssueProductsResult, error)
	IssuesList(ctx context.Context, param IssuesListParam) (*IssuesListResult, error)
	ViewMessage(ctx context.Context, param ViewMessageParam) (*ViewMessageResult, error)
	SendMessage(ctx context.Context, param SendMessageParam) (*SendMessageResult, error)
	ExportActivityLog(ctx context.Context) (*ExportActivityLogResult, error)
	CreateIssueForCustomer(ctx context.Context, param CreateIssueParam) (*CreateIssueResult, error)
	SetAccessToken(token string)
}

type MasterCardCredential struct {
	AuthURL     string
	BaseURL     string
	CustomerKey string
	AccessToken string
}

type MasterCardAPI struct {
	config  *internal.Config
	support customersupport.CustomerSupportAPIInterface
}

func NewMasterCardAPI(param *MasterCardCredential) MasterCardAPIInterface {
	config := internal.SetConfig(
		param.AuthURL,
		param.BaseURL,
		param.CustomerKey,
	)

	return &MasterCardAPI{
		config: config,
		support: customersupport.NewCustomerSupportAPI(customersupport.CustomerSupportCredential{
			AuthURL:     config.AuthURL,
			BaseURL:     config.BaseURL,
			CustomerKey: config.CustomerKey,
		}),
	}
}

func (m *MasterCardAPI) SetAccessToken(token string) {
	m.support.SetAccessToken(token)
}

func (m *MasterCardAPI) WorkerLogin(ctx context.Context, param WorkerLoginParam) (*WorkerLoginResult, error) {
	result, err := m.support.WorkerLogin(ctx, param)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (m *MasterCardAPI) IssueProducts(ctx context.Context) (*IssueProductsResult, error) {
	return m.support.IssueProducts(ctx)
}

func (m *MasterCardAPI) IssuesList(ctx context.Context, param IssuesListParam) (*IssuesListResult, error) {
	return m.support.IssuesList(ctx, param)
}

func (m *MasterCardAPI) ViewMessage(ctx context.Context, param ViewMessageParam) (*ViewMessageResult, error) {
	return m.support.ViewMessage(ctx, param)
}

func (m *MasterCardAPI) SendMessage(ctx context.Context, param SendMessageParam) (*SendMessageResult, error) {
	return m.support.SendMessage(ctx, param)
}

func (m *MasterCardAPI) ExportActivityLog(ctx context.Context) (*ExportActivityLogResult, error) {
	return m.support.ExportActivityLog(ctx)
}

func (m *MasterCardAPI) CreateIssueForCustomer(ctx context.Context, param CreateIssueParam) (*CreateIssueResult, error) {
	return m.support.CreateIssueForCustomer(ctx, param)
}
