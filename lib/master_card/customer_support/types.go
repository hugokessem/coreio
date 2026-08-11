package customersupport

import "encoding/json"

type WorkerLoginParam struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	YabxTxnID string `json:"yabx_txn_id"`
}

type WorkerLoginResult struct {
	Success     bool   `json:"success"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	ForceChange bool   `json:"forceChange"`
	Role        string `json:"role"`
	UserStage   string `json:"userStage"`
	Message     string `json:"message"`
	AccessToken string `json:"accessToken"`
}

type IssueProductsResult struct {
	Success bool            `json:"success"`
	Status  string          `json:"status"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

type IssuesListParam struct {
	UserT24ID   string `json:"usert24id"`
	UserPhone   string `json:"userPhone"`
	IssueStatus string `json:"issueStatus"`
}

type IssuesListResult struct {
	Success bool            `json:"success"`
	Status  string          `json:"status"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

type ViewMessageParam struct {
	UserID string `json:"userId"`
	Ticket string `json:"ticket"`
}

type ViewMessageResult struct {
	Success bool            `json:"success"`
	Status  string          `json:"status"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

type SendMessageParam struct {
	UserT24ID   string `json:"usert24id"`
	Username    string `json:"username"`
	Ticket      string `json:"ticket"`
	Chat        string `json:"chat"`
	MessageType string `json:"messageType"`
	SentFrom    string `json:"sentFrom"`
	Flag        string `json:"flag"`
}

type SendMessageResult struct {
	Success bool   `json:"success"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ActivityLog struct {
	LogID      int    `json:"log_id"`
	UserID     int    `json:"user_id"`
	IssueID    int64  `json:"issue_id"`
	CustomerID string `json:"customerID"`
	UserEmail  string `json:"userEmail"`
	ActionType string `json:"action_type"`
	Details    string `json:"details"`
	UserFrom   string `json:"userFrom"`
	CreatedAt  string `json:"created_at"`
	Ticket     string `json:"ticket"`
}

type ExportActivityLogResult struct {
	Success bool          `json:"success"`
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Data    []ActivityLog `json:"data"`
}

type CreateIssueParam struct {
	CustomerID      string `json:"customerid"`
	Username        string `json:"username"`
	UserPhone       string `json:"userphone"`
	AccountNo       string `json:"accountNo"`
	Branch          string `json:"branch"`
	Place           string `json:"place"`
	ATMInfo         string `json:"atminfo"`
	Description     string `json:"description"`
	Money           string `json:"money"`
	IssueTypeMenu   string `json:"issuetypeMenu"`
	IssueTypeTitle  string `json:"issuetypeTitle"`
	IssueTypeDetail string `json:"issuetypeDetail"`
	Date            string `json:"date"`
	FTReference     string `json:"ftReference"`
	BeneAcc         string `json:"beneAcc"`
	BenePhone       string `json:"benePhone"`
	BillRef         string `json:"billRef"`
	AgentID         string `json:"agentId"`
	CardNo          string `json:"cardNo"`
	POSTerminalID   string `json:"posTerminalid"`
	Passport        string `json:"passport"`
	Email           string `json:"email"`
	AccessStartDate string `json:"Access_startDate"`
	AccessEndDate   string `json:"Access_endDate"`
	CardSixDigits   string `json:"cardSixDigits"`
	CardFourDigits  string `json:"cardFourDigits"`
	Flag            string `json:"flag"`
}

type CreateIssueResult struct {
	Status  string `json:"status"`
	Ticket  string `json:"ticket"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}
