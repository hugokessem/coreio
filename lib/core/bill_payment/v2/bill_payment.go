package billpaymentv2

import (
	"fmt"
)

type Params struct {
	Username            string
	Password            string
	DebitAccountNumber  string
	DebitCurrency       string
	DebitAmount         string
	DebitReference      string
	CrediterReference   string
	CreditAccountNumber string
	CreditCurrency      string
	ClientReference     string
	ServiceCode         string
	CustomerRole        string
	ChannelType         string
	BudgetType          string
	UserID              string
}

func NewBillPayment(param Params) string {
	return fmt.Sprintf(`
		<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cbes="http://temenos.com/CBESUPERAPP" xmlns:fun="http://temenos.com/FUNDSTRANSFERBILLPAYSUPERAPP">
		<soapenv:Header/>
		<soapenv:Body>
			<cbes:FTBillPayment>
				<WebRequestCommon>
					<company/>
					<password>%s</password>
					<userName>%s</userName>
				</WebRequestCommon>
				<OfsFunction/>
				<FUNDSTRANSFERBILLPAYSUPERAPPType id="">
					<fun:DEBITACCTNO>%s</fun:DEBITACCTNO>
					<fun:DEBITCURRENCY>%s</fun:DEBITCURRENCY>
					<fun:DEBITAMOUNT>%s</fun:DEBITAMOUNT>
					<fun:DEBITTHEIRREF>%s</fun:DEBITTHEIRREF>
					<fun:CREDITTHEIRREF>%s</fun:CREDITTHEIRREF>
					<fun:CREDITACCTNO>%s</fun:CREDITACCTNO>
					<fun:CREDITCURRENCY>%s</fun:CREDITCURRENCY>
					<fun:CREDITAMOUNT/>
					<fun:ClientReference>%s</fun:ClientReference>
					<fun:ServiceCode>%s</fun:ServiceCode>
					<fun:CustomerRole>%s</fun:CustomerRole>
					<fun:ChannelType>%s</fun:ChannelType>
					<fun:BudgetType>%s</fun:BudgetType>
					<fun:UserID>%s</fun:UserID>
				</FUNDSTRANSFERBILLPAYSUPERAPPType>
			</cbes:FTBillPayment>
		</soapenv:Body>
	</soapenv:Envelope>
	`, param.Password, param.Username, param.DebitAccountNumber, param.DebitCurrency, param.DebitAmount, param.DebitReference, param.CrediterReference, param.CreditAccountNumber, param.CreditCurrency, param.ClientReference, param.ServiceCode, param.CustomerRole, param.ChannelType, param.BudgetType, param.UserID)
}

type BillPaymentDetail struct {
}

type BillPaymentResult struct {
	Status  bool
	Detail  *BillPaymentDetail
	Message []string
}
