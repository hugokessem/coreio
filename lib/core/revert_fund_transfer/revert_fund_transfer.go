package revertfundtransfer

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Params struct {
	Username      string
	Password      string
	TransactionID string
}
type RevertFundTransferParams struct {
	TransactionID string
}

func NewRevertFundTransfer(param Params) string {
	return fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cbes="http://temenos.com/CBESUPERAPPV2">
    <soapenv:Header/>
    <soapenv:Body>
        <cbes:AccountTransferReversal>
            <WebRequestCommon>
                <company/>
                <password>%s</password>
                <userName>%s</userName>
            </WebRequestCommon>
            <OfsFunction/>
            <FUNDSTRANSFERFTREVERSESUPERAPPType>
                <transactionId>%s</transactionId>
            </FUNDSTRANSFERFTREVERSESUPERAPPType>
        </cbes:AccountTransferReversal>
    </soapenv:Body>
</soapenv:Envelope>`, param.Password, param.Username, param.TransactionID)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	RevertFundTransferResponse *RevertFundTransferResponse `xml:"AccountTransferReversalResponse"`
}

type RevertFundTransferDetail struct {
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

type RevertFundTransferResponse struct {
	Status *struct {
		SuccessIndicator string `xml:"successIndicator"`
		TransactionID    string `xml:"transactionId"`
		Application      string `xml:"application"`
		MessageId        string `xml:"messageId"`
	} `xml:"Status"`
	FundTransferType *RevertFundTransferDetail `xml:"FUNDSTRANSFERType"`
}

type RevertFundTransferResult struct {
	Success  bool
	Detail   *RevertFundTransferDetail
	Messages []string
}

func ParseRevertFundTransferSOAP(xmlData string) (*RevertFundTransferResult, error) {
	var env Envelope
	if err := xml.Unmarshal([]byte(xmlData), &env); err != nil {
		return nil, err
	}

	if env.Body.RevertFundTransferResponse != nil {
		resp := env.Body.RevertFundTransferResponse
		if resp.Status == nil {
			return &RevertFundTransferResult{
				Success:  false,
				Messages: []string{"Missing Status"},
			}, nil
		}

		if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
			return &RevertFundTransferResult{
				Success:  false,
				Messages: []string{"API returned failure"},
			}, nil
		}

		if resp.FundTransferType == nil {
			return &RevertFundTransferResult{
				Success:  true,
				Messages: []string{},
			}, nil
		}

		return &RevertFundTransferResult{
			Success: true,
			Detail: &RevertFundTransferDetail{
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

	return &RevertFundTransferResult{
		Success:  false,
		Messages: []string{"Invalid response type"},
	}, nil
}
