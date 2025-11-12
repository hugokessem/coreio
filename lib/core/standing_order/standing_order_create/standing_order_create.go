package standingordercreate

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Params struct {
	Username            string
	Password            string
	DebitAccountNumber  string
	CreditAccountNumber string
	CurrentDate         string
	Amount              string
	Currency            string
	Frequency           string
	PaymentDetail       string
}
type CreateStandingOrderParams struct {
	DebitAccountNumber  string
	CreditAccountNumber string
	CurrentDate         string
	Amount              string
	Currency            string
	Frequency           string
	PaymentDetail       string
}

func NewCreateStandingOrder(param Params) string {
	return fmt.Sprintf(
		`<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cbes="http://temenos.com/CBESUPERAPPV2" xmlns:stan="http://temenos.com/STANDINGORDERMANAGEORDERSUPERAPP">
    <soapenv:Header/>
    <soapenv:Body>
        <cbes:CreateUpdateStandingOrder>
            <WebRequestCommon>
                <company/>
                <password>%s</password>
                <userName>%s</userName>
            </WebRequestCommon>
            <OfsFunction></OfsFunction>
            <STANDINGORDERMANAGEORDERSUPERAPPType id="%s.">
                <stan:CURRENCY>%s</stan:CURRENCY>
                <stan:CURRENTAMOUNTBAL>%s</stan:CURRENTAMOUNTBAL>
                <stan:CURRENTFREQUENCY>%s</stan:CURRENTFREQUENCY>
                <stan:CURRENTENDDATE>%s</stan:CURRENTENDDATE>
                <stan:gPAYMENTDETAILS g="1">
                    <stan:PAYMENTDETAILS>%s</stan:PAYMENTDETAILS>
                </stan:gPAYMENTDETAILS>
                <stan:CPTYACCTNO>%s</stan:CPTYACCTNO>
            </STANDINGORDERMANAGEORDERSUPERAPPType>
        </cbes:CreateUpdateStandingOrder>
    </soapenv:Body>
</soapenv:Envelope>`, param.Password, param.Username, param.DebitAccountNumber, param.Currency, param.Amount, param.Frequency, param.CurrentDate, param.PaymentDetail, param.CreditAccountNumber)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	StandingOrderResponse *StandingOrderResponse `xml:"CreateUpdateStandingOrderResponse"`
}

type StandingOrderDetail struct {
	Type                string `xml:"TYPE"`
	PaymentMethod       string `xml:"PAYMETHOD"`
	Currency            string `xml:"CURRENCY"`
	Amount              string `xml:"CURRENTAMOUNTBAL"`
	Frequency           string `xml:"CURRENTFREQUENCY"`
	CurrentDate         string `xml:"CURRENTENDDATE"`
	CreditAccountNumber string `xml:"CPTYACCTNO"`
	NextPayment         string `xml:"CURRFREQDATE"`
}

type StandingOrderResponse struct {
	Status *struct {
		SuccessIndicator string `xml:"successIndicator"`
		MessageId        string `xml:"messageId"`
		Application      string `xml:"application"`
		TransactionId    string `xml:"transactionId"`
	} `xml:"Status"`
	StandingOrderType *StandingOrderDetail `xml:"STANDINGORDERType"`
}

type StandingOrderResult struct {
	Success bool
	Detail  *StandingOrderDetail
	Message []string
}

func ParseCreateStandingOrderSOAP(xmlData string) (*StandingOrderResult, error) {
	var env Envelope
	if err := xml.Unmarshal([]byte(xmlData), &env); err != nil {
		return nil, err
	}

	if env.Body.StandingOrderResponse != nil {
		resp := env.Body.StandingOrderResponse
		if resp.Status == nil {
			return &StandingOrderResult{
				Success: false,
				Message: []string{"Missing Status"},
			}, nil
		}

		if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
			return &StandingOrderResult{
				Success: false,
				Message: []string{"API returned failure"},
			}, nil
		}

		if resp.StandingOrderType == nil {
			return &StandingOrderResult{
				Success: false,
				Message: []string{},
			}, nil
		}

		return &StandingOrderResult{
			Success: true,
			Detail: &StandingOrderDetail{
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

	return &StandingOrderResult{
		Success: false,
		Message: []string{"Invalid response type"},
	}, nil
}
