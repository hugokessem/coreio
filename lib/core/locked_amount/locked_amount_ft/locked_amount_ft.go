package lockedamountft

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Params struct {
	Username            string
	Password            string
	CreditCurrent       string
	CreditAccountNumber string
	CrediterReference   string
	DebitAmount         string
	DebitAccountNumber  string
	DebitCurrency       string
	DebiterReference    string
	ClientReference     string
	ServiceCode         string
	LockID              string
	PaymentDetails      string
}

type LockedAmountFTParams struct {
	CreditCurrent       string
	CreditAccountNumber string
	CrediterReference   string
	DebitAmount         string
	DebitAccountNumber  string
	DebitCurrency       string
	DebiterReference    string
	ClientReference     string
	ServiceCode         string
	LockID              string
	PaymentDetails      string
}

func NewLockedAmountFt(param Params) string {
	return fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"
xmlns:cbes="http://temenos.com/CBESUPERAPP"
xmlns:fun="http://temenos.com/FUNDSTRANSFERFTTXNSUPERAPP">
    <soapenv:Header/>
    <soapenv:Body>
        <cbes:AccountTransfer>
            <WebRequestCommon>
                <company/>
                <password>%s</password>
                <userName>%s</userName>
            </WebRequestCommon>
            <OfsFunction/>
            <FUNDSTRANSFERFTTXNSUPERAPPType id="">
                <fun:DEBITACCTNO>%s</fun:DEBITACCTNO>
                <fun:DEBITCURRENCY>%s</fun:DEBITCURRENCY>
                <fun:DEBITAMOUNT>%s</fun:DEBITAMOUNT>
                <fun:DEBITTHEIRREF>%s</fun:DEBITTHEIRREF>
                <fun:CREDITTHEIRREF>%s</fun:CREDITTHEIRREF>
                <fun:CREDITACCTNO>%s</fun:CREDITACCTNO>
                <fun:CREDITCURRENCY>%s</fun:CREDITCURRENCY>
                <fun:CREDITAMOUNT/>
                <fun:gPAYMENTDETAILS g="1">
                    <fun:PAYMENTDETAILS>%s</fun:PAYMENTDETAILS>
                </fun:gPAYMENTDETAILS>
                <fun:COMMISSIONCODE/>
                <fun:CHARGECODE/>
                <fun:ClientReference>%s</fun:ClientReference>
                <fun:LockID>%s</fun:LockID>
				<fun:ServiceCode>%s</fun:ServiceCode>
            </FUNDSTRANSFERFTTXNSUPERAPPType>
        </cbes:AccountTransfer>
    </soapenv:Body>
</soapenv:Envelope>
	`, param.Password, param.Username, param.DebitAccountNumber, param.DebitCurrency, param.DebitAmount, param.DebiterReference, param.CrediterReference, param.CreditAccountNumber, param.CreditCurrent, param.PaymentDetails, param.ClientReference, param.LockID, param.ServiceCode)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	AccountTransferResponse *AccountTransferResponse `xml:"AccountTransferResponse"`
}

type AccountTransferResponse struct {
	Status *struct {
		SuccessIndicator string   `xml:"successIndicator"`
		TransactionId    string   `xml:"transactionId"`
		Application      string   `xml:"application"`
		Messages         []string `xml:"messages"`
	} `xml:"Status"`
	FundTransferResponse *FundTransferResponse `xml:"FUNDSTRANSFERType"`
}

type FundTransferResponse struct {
	TransactionId               string
	TransactionType             string `xml:"TRANSACTIONTYPE"`
	DebitAccountNumber          string `xml:"DEBITACCTNO"`
	DebitAccountCustomerNumber  string `xml:"DEBITCUSTOMER"`
	DebitCurrency               string `xml:"DEBITCURRENCY"`
	DebitAmount                 string `xml:"DEBITAMOUNT"`
	DebitedDate                 string `xml:"DEBITVALUEDATE"`
	DebiterReference            string `xml:"DEBITTHEIRREF"`
	CreditAccountNumber         string `xml:"CREDITACCTNO"`
	CreidtAccountCustomerNumber string `xml:"CREDITCUSTOMER"`
	CreditCurrenct              string `xml:"CREDITCURRENCY"`
	CreditedDate                string `xml:"CREDITVALUEDATE"`
	CrediterReference           string `xml:"CREDITTHEIRREF"`
	ComissionCode               string `xml:"COMMISSIONCODE"`
	ChargeCode                  string `xml:"CHARGECODE"`
	CreditAmountWithCurrency    string `xml:"AMOUNTCREDITED"`
	DebitAmountWithCurrency     string `xml:"AMOUNTDEBITED"`
	DebitedLockedAmount         string `xml:"LOCAMTDEBITED"`
	CreditLockedAmount          string `xml:"LOCAMTCREDITED"`
	LockId                      string `xml:"ACLOCKID"`
	LocalAmountDebited          string `xml:"LOCAMTDEBITED"`
	LocalAmountCredited         string `xml:"LOCAMTCREDITED"`
	LocalTotalTaxAmount         string `xml:"LOCTOTTAXAMT"`
	LocalChargeAmount           string `xml:"LOCALCHARGEAMT"`
	LocalPositionChargesAmount  string `xml:"LOCPOSCHGSAMT"`
}

type LockedAmountFTResult struct {
	Success  bool
	Detail   FundTransferResponse
	Messages []string
}

func ParseLockedAmountFTSOAP(xmlData string) (*LockedAmountFTResult, error) {
	var env Envelope
	if err := xml.Unmarshal([]byte(xmlData), &env); err != nil {
		return nil, err
	}

	if env.Body.AccountTransferResponse != nil {
		resp := env.Body.AccountTransferResponse
		if resp.Status == nil {
			return &LockedAmountFTResult{
				Success:  false,
				Messages: []string{"Missing Status"},
			}, nil
		}

		if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
			return &LockedAmountFTResult{
				Success:  false,
				Messages: resp.Status.Messages,
			}, nil
		}

		if resp.FundTransferResponse == nil {
			return &LockedAmountFTResult{
				Success:  false,
				Messages: []string{"Missing FundTransferResponse"},
			}, nil
		}

		return &LockedAmountFTResult{
			Success: true,
			Detail: FundTransferResponse{
				TransactionId:               resp.Status.TransactionId,
				TransactionType:             resp.FundTransferResponse.TransactionType,
				DebitAccountNumber:          resp.FundTransferResponse.DebitAccountNumber,
				DebitAccountCustomerNumber:  resp.FundTransferResponse.DebitAccountCustomerNumber,
				DebitCurrency:               resp.FundTransferResponse.DebitCurrency,
				DebitAmount:                 resp.FundTransferResponse.DebitAmount,
				DebitedDate:                 resp.FundTransferResponse.DebitedDate,
				DebiterReference:            resp.FundTransferResponse.DebiterReference,
				CreditAccountNumber:         resp.FundTransferResponse.CreditAccountNumber,
				CreidtAccountCustomerNumber: resp.FundTransferResponse.CreidtAccountCustomerNumber,
				CreditCurrenct:              resp.FundTransferResponse.CreditCurrenct,
				CreditedDate:                resp.FundTransferResponse.CreditedDate,
				CrediterReference:           resp.FundTransferResponse.CrediterReference,
				ComissionCode:               resp.FundTransferResponse.ComissionCode,
				ChargeCode:                  resp.FundTransferResponse.ChargeCode,
				CreditAmountWithCurrency:    resp.FundTransferResponse.CreditAmountWithCurrency,
				DebitAmountWithCurrency:     resp.FundTransferResponse.DebitAmountWithCurrency,
				DebitedLockedAmount:         resp.FundTransferResponse.DebitedLockedAmount,
				CreditLockedAmount:          resp.FundTransferResponse.CreditLockedAmount,
				LockId:                      resp.FundTransferResponse.LockId,
				LocalAmountDebited:          resp.FundTransferResponse.LocalAmountDebited,
				LocalAmountCredited:         resp.FundTransferResponse.LocalAmountCredited,
				LocalTotalTaxAmount:         resp.FundTransferResponse.LocalTotalTaxAmount,
				LocalChargeAmount:           resp.FundTransferResponse.LocalChargeAmount,
				LocalPositionChargesAmount:  resp.FundTransferResponse.LocalPositionChargesAmount,
			},
		}, nil
	}

	return &LockedAmountFTResult{
		Success:  false,
		Messages: []string{"Invalid response type"},
	}, nil
}
