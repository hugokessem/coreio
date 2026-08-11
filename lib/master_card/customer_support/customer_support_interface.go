package customersupport

import "context"

type CustomerSupportAPIInterface interface {
	WorkerLogin(ctx context.Context, param WorkerLoginParam) (*WorkerLoginResult, error)
	IssueProducts(ctx context.Context) (*IssueProductsResult, error)
	IssuesList(ctx context.Context, param IssuesListParam) (*IssuesListResult, error)
	ViewMessage(ctx context.Context, param ViewMessageParam) (*ViewMessageResult, error)
	SendMessage(ctx context.Context, param SendMessageParam) (*SendMessageResult, error)
	ExportActivityLog(ctx context.Context) (*ExportActivityLogResult, error)
	CreateIssueForCustomer(ctx context.Context, param CreateIssueParam) (*CreateIssueResult, error)
	SetAccessToken(token string)
}
