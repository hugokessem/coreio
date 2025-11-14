package standingordercancel

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Params struct {
	Username      string
	Password      string
	AccountNumber string
	OrderId       string
}

type CancelStandingOrderParams struct {
	AccountNumber string
	OrderId       string
}

func NewCancleStandingOrder(param Params) string {
	return fmt.Sprintf(
		`<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cbes="http://temenos.com/CBESUPERAPP">
    <soapenv:Header/>
    <soapenv:Body>
        <cbes:CancelStandingOrder>
            <WebRequestCommon>
                <company/>
                <password>%s</password>
                <userName>%s</userName>
            </WebRequestCommon>
            <OfsFunction/>
            <STANDINGORDERMANAGEORDERSUPERAPPType>
                <transactionId>%s.%s</transactionId>
            </STANDINGORDERMANAGEORDERSUPERAPPType>
        </cbes:CancelStandingOrder>
    </soapenv:Body>
</soapenv:Envelope>`, param.Password, param.Username, param.AccountNumber, param.OrderId)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	CancelStandingOrderResponse *CancelStandingOrderResponse `xml:"CancelStandingOrderResponse"`
}

type CancelStandingOrderDetail struct {
	Type                string `xml:"TYPE"`
	PaymentMethod       string `xml:"PAYMETHOD"`
	Currency            string `xml:"CURRENCY"`
	Amount              string `xml:"CURRENTAMOUNTBAL"`
	Frequency           string `xml:"CURRENTFREQUENCY"`
	CurrentDate         string `xml:"CURRENTENDDATE"`
	CreditAccountNumber string `xml:"CPTYACCTNO"`
	NextPayment         string `xml:"CURRFREQDATE"`
}

type CancelStandingOrderResponse struct {
	Status *struct {
		SuccessIndicator string `xml:"successIndicator"`
		MessageId        string `xml:"messageId"`
		Application      string `xml:"application"`
		TransactionId    string `xml:"transactionId"`
	} `xml:"Status"`
	StandingOrderType *CancelStandingOrderDetail `xml:"STANDINGORDERType"`
}

type CancelStandingOrderResult struct {
	Success bool
	Detail  *CancelStandingOrderDetail
	Message []string
}

func ParseCancelStandingOrderSOAP(xmlData string) (*CancelStandingOrderResult, error) {
	var env Envelope
	if err := xml.Unmarshal([]byte(xmlData), &env); err != nil {
		return nil, err
	}

	if env.Body.CancelStandingOrderResponse != nil {
		resp := env.Body.CancelStandingOrderResponse
		if resp.Status == nil {
			return &CancelStandingOrderResult{
				Success: false,
				Message: []string{"Missing Status"},
			}, nil
		}

		if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
			return &CancelStandingOrderResult{
				Success: false,
				Message: []string{"API returned failure"},
			}, nil
		}

		if resp.StandingOrderType == nil {
			return &CancelStandingOrderResult{
				Success: false,
				Message: []string{},
			}, nil
		}

		return &CancelStandingOrderResult{
			Success: true,
			Detail: &CancelStandingOrderDetail{
				Type:                resp.StandingOrderType.Type,
				Amount:              resp.StandingOrderType.Amount,
				Currency:            resp.StandingOrderType.Currency,
				Frequency:           resp.StandingOrderType.Frequency,
				CurrentDate:         resp.StandingOrderType.CurrentDate,
				NextPayment:         resp.StandingOrderType.NextPayment,
				PaymentMethod:       resp.StandingOrderType.PaymentMethod,
				CreditAccountNumber: resp.StandingOrderType.CreditAccountNumber,
			},
		}, nil
	}

	return &CancelStandingOrderResult{
		Success: false,
		Message: []string{"Invalid response type"},
	}, nil
}
