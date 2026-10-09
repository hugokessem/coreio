package rtgs

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Params struct {
	Username           string
	Password           string
	Company            string
	DebitAccountNumber string
	DebitCurrency      string
	CreditAmount       string
	OrderingCust       []string
	BenAccountNumber   string
	PaymentDetails     string
	BkToBkInfo         []string
	BillID             string
}

type RTGSParam struct {
	Company            string
	DebitAccountNumber string
	DebitCurrency      string
	CreditAmount       string
	OrderingCust       []string
	BenAccountNumber   string
	PaymentDetails     string
	BkToBkInfo         []string
	BillID             string
}

func buildRepeatingElements(tag string, values []string) string {
	if len(values) == 0 {
		return ""
	}
	var b strings.Builder
	for _, v := range values {
		b.WriteString(fmt.Sprintf("\n\t\t\t\t\t<fun:%s>%s</fun:%s>", tag, v, tag))
	}
	return b.String()
}

func NewRTGS(params Params) string {
	orderingCust := buildRepeatingElements("ORDERINGCUST", params.OrderingCust)
	bkToBkInfo := buildRepeatingElements("BKTOBKINFO", params.BkToBkInfo)

	return fmt.Sprintf(`<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"
xmlns:cbes="http://temenos.com/CBESUPERAPP" xmlns:fun="http://temenos.com/FUNDSTRANSFERMORSUPERAPP">
<soapenv:Header/>
<soapenv:Body>
<cbes:MinistryOfRevenueMORRTGSFT>
<WebRequestCommon>
<company>%s</company>
<password>%s</password>
<userName>%s</userName>
</WebRequestCommon>
<OfsFunction/>
<FUNDSTRANSFERMORSUPERAPPType id="">
<fun:DEBITACCTNO>%s</fun:DEBITACCTNO>
<fun:DEBITCURRENCY>%s</fun:DEBITCURRENCY>
<fun:CREDITAMOUNT>%s</fun:CREDITAMOUNT>
<fun:gORDERINGCUST g="1">%s
</fun:gORDERINGCUST>
<fun:BENACCTNO>%s</fun:BENACCTNO>
<fun:gPAYMENTDETAILS g="1">
<fun:PAYMENTDETAILS>%s</fun:PAYMENTDETAILS>
</fun:gPAYMENTDETAILS>
<fun:gBKTOBKINFO g="1">%s
</fun:gBKTOBKINFO>
<fun:BILLID>%s</fun:BILLID>
</FUNDSTRANSFERMORSUPERAPPType>
</cbes:MinistryOfRevenueMORRTGSFT>
</soapenv:Body>
</soapenv:Envelope>`,
		params.Company,
		params.Password,
		params.Username,
		params.DebitAccountNumber,
		params.DebitCurrency,
		params.CreditAmount,
		orderingCust,
		params.BenAccountNumber,
		params.PaymentDetails,
		bkToBkInfo,
		params.BillID,
	)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	MinistryOfRevenueMORRTGSFTResponse *MinistryOfRevenueMORRTGSFTResponse `xml:"MinistryOfRevenueMORRTGSFTResponse"`
}

type MinistryOfRevenueMORRTGSFTResponse struct {
	Status *struct {
		SuccessIndicator string   `xml:"successIndicator"`
		TransactionID    string   `xml:"transactionId"`
		Application      string   `xml:"application"`
		MessageId        string   `xml:"messageId"`
		Messages         []string `xml:"messages"`
	} `xml:"Status"`
	FundTransferType *RTGSDetail `xml:"FUNDSTRANSFERType"`
}

type RTGSDetail struct {
	XMLName                  xml.Name `xml:"FUNDSTRANSFERType"`
	FTNumber                 string   `xml:"id,attr"`
	TransactionType          string   `xml:"TRANSACTIONTYPE"`
	DebitAccountNumber       string   `xml:"DEBITACCTNO"`
	DebitCurrency            string   `xml:"DEBITCURRENCY"`
	DebitAmount              string   `xml:"DEBITAMOUNT"`
	CreditAmount             string   `xml:"CREDITAMOUNT"`
	BenAccountNumber         string   `xml:"BENACCTNO"`
	BillID                   string   `xml:"BILLID"`
	DebitAmountWithCurrency  string   `xml:"AMOUNTDEBITED"`
	CreditAmountWithCurrency string   `xml:"AMOUNTCREDITED"`
	ProcessingDate           string   `xml:"PROCESSINGDATE"`
	CompanyCode              string   `xml:"COCODE"`
	GlobalOrderingCust       struct {
		OrderingCust []string `xml:"ORDERINGCUST"`
	} `xml:"gORDERINGCUST"`
	GlobalPaymentDetails struct {
		PaymentDetails string `xml:"PAYMENTDETAILS"`
	} `xml:"gPAYMENTDETAILS"`
	GlobalBkToBkInfo struct {
		BkToBkInfo []string `xml:"BKTOBKINFO"`
	} `xml:"gBKTOBKINFO"`
	GlobalOverride struct {
		Override []string `xml:"OVERRIDE"`
	} `xml:"gOVERRIDE"`
}

type RTGSResult struct {
	Success  bool
	Detail   *RTGSDetail
	Messages []string
}

func ParseRTGSSOAP(xmlData string) (*RTGSResult, error) {
	var env Envelope
	if err := xml.Unmarshal([]byte(xmlData), &env); err != nil {
		return nil, err
	}

	if env.Body.MinistryOfRevenueMORRTGSFTResponse != nil {
		resp := env.Body.MinistryOfRevenueMORRTGSFTResponse
		if resp.Status == nil {
			return &RTGSResult{
				Success:  false,
				Messages: []string{"Missing Status"},
			}, nil
		}

		if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
			messages := resp.Status.Messages
			if len(messages) == 0 {
				messages = []string{"API returned failure"}
			}
			return &RTGSResult{
				Success:  false,
				Messages: messages,
			}, nil
		}

		if resp.FundTransferType == nil {
			return &RTGSResult{
				Success:  true,
				Messages: resp.Status.Messages,
			}, nil
		}

		return &RTGSResult{
			Success:  true,
			Detail:   resp.FundTransferType,
			Messages: resp.Status.Messages,
		}, nil
	}

	return &RTGSResult{
		Success:  false,
		Messages: []string{"Invalid response type"},
	}, nil
}
