// Package fundtransfer provides SOAP XML generators and parsers
// for core banking fund transfer requests.
package fundtransfer

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Params struct {
	Username            string
	Password            string
	DebitAccountNumber  string
	DebitCurrency       string
	CreditAccountNumber string
	CreditCurrency      string
	DebitReference      string
	CreditReference     string
	DebitAmount         string
	TransactionID       string
	PaymentDetail       string
}

type FundTransferParam struct {
	DebitAccountNumber  string
	DebitCurrency       string
	CreditAccountNumber string
	CreditCurrency      string
	DebitReference      string
	CreditReference     string
	DebitAmount         string
	TransactionID       string
	PaymentDetail       string
}

func NewFundTransfer(param Params) string {
	return fmt.Sprintf(
		`<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cbes="http://temenos.com/CBESUPERAPP" xmlns:fun="http://temenos.com/FUNDSTRANSFERFTTXNSUPERAPP">
		<soapenv:Header/>
		<soapenv:Body>
			<cbes:AccountTransfer>
				<WebRequestCommon>
					<company></company>
					<password>%s</password>
					<userName>%s</userName>
				</WebRequestCommon>
				<OfsFunction></OfsFunction>
				<FUNDSTRANSFERFTTXNSUPERAPPType id="">
					<fun:DEBITACCTNO>%s</fun:DEBITACCTNO>
					<fun:DEBITCURRENCY>%s</fun:DEBITCURRENCY>
					<fun:DEBITAMOUNT>%s</fun:DEBITAMOUNT>
					<fun:DEBITTHEIRREF>%s</fun:DEBITTHEIRREF>
					<fun:CREDITTHEIRREF>%s</fun:CREDITTHEIRREF>
					<fun:CREDITACCTNO>%s</fun:CREDITACCTNO>
					<fun:CREDITCURRENCY>%s</fun:CREDITCURRENCY>
					<fun:CREDITAMOUNT></fun:CREDITAMOUNT>
					<fun:gPAYMENTDETAILS g="1">
					<fun:PAYMENTDETAILS>%s</fun:PAYMENTDETAILS>
					</fun:gPAYMENTDETAILS>
					<fun:COMMISSIONCODE></fun:COMMISSIONCODE>
					<fun:CHARGECODE></fun:CHARGECODE>
					<fun:ClientReference>%s</fun:ClientReference>
				</FUNDSTRANSFERFTTXNSUPERAPPType>
			</cbes:AccountTransfer>
		</soapenv:Body>
		</soapenv:Envelope>
`, param.Password, param.Username, param.DebitAccountNumber, param.DebitCurrency, param.DebitAmount, param.DebitReference, param.CreditReference, param.CreditAccountNumber, param.CreditCurrency, param.PaymentDetail, param.TransactionID)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	FundTransferResponse *FundTransferResponse `xml:"AccountTransferResponse"`
}

type FundTransferDetail struct {
	TransactionType                   string   `xml:"TRANSACTIONTYPE"`
	XMLName                           xml.Name `xml:"FUNDSTRANSFERType"`
	FTNumber                          string   `xml:"id,attr"`
	TransactionID                     string   `xml:"MTOREF"`
	DebitAccountNumber                string   `xml:"DEBITACCTNO"`
	DebitAccountHolderName            string   `xml:"SENDERNAME"`
	DebitAccountCurrentWorkingBalance string   `xml:"CEKCS"`
	DebitReference                    string   `xml:"DEBITTHEIRREF"`
	DebitCurrency                     string   `xml:"DEBITCURRENCY"`
	DebitAmount                       string   `xml:"DEBITAMOUNT"`
	DebitAmountWithCurrency           string   `xml:"AMOUNTDEBITED"`

	CreditAccountHolderName            string `xml:"RECEIVERNAME"`
	CreditAccountCurrentWorkingBalance string `xml:"GPONU"`
	CreditAccountNumber                string `xml:"CREDITACCTNO"`
	CreditAmountWithCurrency           string `xml:"AMOUNTCREDITED"`
	CreditReference                    string `xml:"CREDITTHEIRREF"`
	CreditCurrency                     string `xml:"CREDITCURRENCY"`
	CreditValidationDare               string `xml:"CREDITVALUEDATE"`
	ProcessingDate                     string `xml:"PROCESSINGDATE"`
	ChargeCommisionDisplay             string `xml:"CHARGECOMDISPLAY"`
}

type FundTransferResponse struct {
	Status           FundTransferStatus  `xml:"Status"`
	FundTransferType *FundTransferDetail `xml:"FUNDSTRANSFERType"`
}

type FundTransferStatus *struct {
	SuccessIndicator string   `xml:"successIndicator"`
	TransactionID    string   `xml:"transactionId"`
	Application      string   `xml:"application"`
	Messages         []string `xml:"messages"`
}

type FundTransferResult struct {
	Success  bool
	Detail   *FundTransferDetail
	Messages []string
}

func ParseFundTransferSOAP(xmlData string) (*FundTransferResult, error) {
	var env Envelope
	if err := xml.Unmarshal([]byte(xmlData), &env); err != nil {
		return nil, err
	}

	if env.Body.FundTransferResponse != nil {
		resp := env.Body.FundTransferResponse
		if resp.Status == nil {
			return &FundTransferResult{
				Success:  false,
				Messages: []string{"Missing Status"},
			}, nil
		}

		if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
			return &FundTransferResult{
				Success:  false,
				Messages: resp.Status.Messages,
			}, nil
		}

		if resp.FundTransferType == nil {
			return &FundTransferResult{
				Success:  true,
				Messages: []string{},
			}, nil
		}

		return &FundTransferResult{
			Success: true,
			Detail: &FundTransferDetail{
				TransactionType:                    resp.FundTransferType.TransactionType,
				FTNumber:                           resp.FundTransferType.FTNumber,
				TransactionID:                      resp.FundTransferType.TransactionID,
				DebitAccountNumber:                 resp.FundTransferType.DebitAccountNumber,
				DebitAccountHolderName:             resp.FundTransferType.DebitAccountHolderName,
				DebitAmount:                        resp.FundTransferType.DebitAmount,
				DebitCurrency:                      resp.FundTransferType.DebitCurrency,
				DebitAmountWithCurrency:            resp.FundTransferType.DebitAmountWithCurrency,
				DebitAccountCurrentWorkingBalance:  resp.FundTransferType.DebitAccountCurrentWorkingBalance,
				DebitReference:                     resp.FundTransferType.DebitReference,
				CreditAccountNumber:                resp.FundTransferType.CreditAccountNumber,
				CreditAccountHolderName:            resp.FundTransferType.CreditAccountHolderName,
				CreditAccountCurrentWorkingBalance: resp.FundTransferType.CreditAccountCurrentWorkingBalance,
				CreditAmountWithCurrency:           resp.FundTransferType.CreditAmountWithCurrency,
				CreditReference:                    resp.FundTransferType.CreditReference,
				CreditCurrency:                     resp.FundTransferType.CreditCurrency,
				CreditValidationDare:               resp.FundTransferType.CreditValidationDare,
				ProcessingDate:                     resp.FundTransferType.ProcessingDate,
				ChargeCommisionDisplay:             resp.FundTransferType.ChargeCommisionDisplay,
			},
		}, nil
	}

	return &FundTransferResult{
		Success:  false,
		Messages: []string{"Invalid response type"},
	}, nil
}
