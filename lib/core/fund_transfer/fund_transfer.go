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
                <fun:gCHARGETYPE g="1">
                    <fun:CHARGETYPE/>
                </fun:gCHARGETYPE>
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
	TransactionType          string   `xml:"TRANSACTIONTYPE"`
	XMLName                  xml.Name `xml:"FUNDSTRANSFERType"`
	FTNumber                 string   `xml:"id,attr"`
	TransactionID            string   `xml:"MTOREF"`
	DebitAccountNumber       string   `xml:"DEBITACCTNO"`
	DebitReference           string   `xml:"DEBITTHEIRREF"`
	DebitCurrency            string   `xml:"DEBITCURRENCY"`
	DebitAmount              string   `xml:"DEBITAMOUNT"`
	DebitAmountWithCurrency  string   `xml:"AMOUNTDEBITED"`
	CreditAmountWithCurrency string   `xml:"AMOUNTCREDITED"`
	CreditAccountNumber      string   `xml:"CREDITACCTNO"`
	CreditReference          string   `xml:"CREDITTHEIRREF"`
	CreditCurrency           string   `xml:"CREDITCURRENCY"`
	CreditValidationDare     string   `xml:"CREDITVALUEDATE"`
	ProcessingDate           string   `xml:"PROCESSINGDATE"`
	ChargeCommisionDisplay   string   `xml:"CHARGECOMDISPLAY"`
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
				TransactionType:          resp.FundTransferType.TransactionType,
				FTNumber:                 resp.FundTransferType.FTNumber,
				TransactionID:            resp.FundTransferType.TransactionID,
				DebitAccountNumber:       resp.FundTransferType.DebitAccountNumber,
				DebitReference:           resp.FundTransferType.DebitReference,
				DebitCurrency:            resp.FundTransferType.DebitCurrency,
				DebitAmount:              resp.FundTransferType.DebitAmount,
				CreditAmountWithCurrency: resp.FundTransferType.CreditAmountWithCurrency,
				DebitAmountWithCurrency:  resp.FundTransferType.DebitAmountWithCurrency,
				CreditAccountNumber:      resp.FundTransferType.CreditAccountNumber,
				CreditReference:          resp.FundTransferType.CreditReference,
				CreditCurrency:           resp.FundTransferType.CreditCurrency,
				CreditValidationDare:     resp.FundTransferType.CreditValidationDare,
				ProcessingDate:           resp.FundTransferType.ProcessingDate,
				ChargeCommisionDisplay:   resp.FundTransferType.ChargeCommisionDisplay,
			},
		}, nil
	}

	return &FundTransferResult{
		Success:  false,
		Messages: []string{"Invalid response type"},
	}, nil
}
