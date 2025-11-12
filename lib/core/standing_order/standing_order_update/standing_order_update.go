package standingorderupdate

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Params struct {
	Username            string
	Password            string
	DebitAccountNumber  string
	OrderId             string
	Currency            string
	Amount              string
	Frequency           string
	CurrentDate         string
	PaymentDetail       string
	CreditAccountNumber string
}

type UpdateStandingOrderParam struct {
	DebitAccountNumber  string
	OrderId             string
	Currency            string
	Amount              string
	Frequency           string
	CurrentDate         string
	PaymentDetail       string
	CreditAccountNumber string
}

func NewUpdateStandingOrder(param Params) string {
	return fmt.Sprintf(`<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cbes="http://temenos.com/CBESUPERAPPV2" xmlns:stan="http://temenos.com/STANDINGORDERMANAGEORDERSUPERAPP">
    <soapenv:Header/>
    <soapenv:Body>
        <cbes:CreateUpdateStandingOrder>
            <WebRequestCommon>
                <company/>
                <password>%s</password>
                <userName>%s</userName>
            </WebRequestCommon>
            <OfsFunction/>
            <STANDINGORDERMANAGEORDERSUPERAPPType id="%s.%s">
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
</soapenv:Envelope>`, param.Password, param.Username, param.DebitAccountNumber, param.OrderId, param.Currency, param.Amount, param.Frequency, param.CurrentDate, param.PaymentDetail, param.CreditAccountNumber)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	UpdateStandingOrderResponse *UpdateStandingOrderResponse `xml:"CreateUpdateStandingOrderResponse"`
}

type UpdateStandingOrderDetail struct {
	Type                string `xml:"TYPE"`
	PaymentMethod       string `xml:"PAYMETHOD"`
	Currency            string `xml:"CURRENCY"`
	Amount              string `xml:"CURRENTAMOUNTBAL"`
	Frequency           string `xml:"CURRENTFREQUENCY"`
	CurrentDate         string `xml:"CURRENTENDDATE"`
	CreditAccountNumber string `xml:"CPTYACCTNO"`
	NextPayment         string `xml:"CURRFREQDATE"`
}

type UpdateStandingOrderResponse struct {
	Status *struct {
		SuccessIndicator string `xml:"successIndicator"`
		MessageId        string `xml:"messageId"`
		Application      string `xml:"application"`
		TransactionId    string `xml:"transactionId"`
	} `xml:"Status"`
	StandingOrderType *UpdateStandingOrderDetail `xml:"STANDINGORDERType"`
}

type UpdateStandingOrderResult struct {
	Success bool
	Detail  *UpdateStandingOrderDetail
	Message []string
}

func ParseUpdateStandingOrderSOAP(xmlData string) (*UpdateStandingOrderResult, error) {
	var env Envelope
	if err := xml.Unmarshal([]byte(xmlData), &env); err != nil {
		return nil, err
	}

	if env.Body.UpdateStandingOrderResponse != nil {
		resp := env.Body.UpdateStandingOrderResponse
		if resp.Status == nil {
			return &UpdateStandingOrderResult{
				Success: false,
				Message: []string{"Missing Status"},
			}, nil
		}

		if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
			return &UpdateStandingOrderResult{
				Success: false,
				Message: []string{"API returned failure"},
			}, nil
		}

		if resp.StandingOrderType == nil {
			return &UpdateStandingOrderResult{
				Success: false,
				Message: []string{},
			}, nil
		}

		return &UpdateStandingOrderResult{
			Success: true,
			Detail: &UpdateStandingOrderDetail{
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

	return &UpdateStandingOrderResult{
		Success: false,
		Message: []string{"Invalid response type"},
	}, nil
}
