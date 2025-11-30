package standingorderlist

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Params struct {
	Username      string
	Password      string
	AccountNumber string
}

type ListStandingOrderParams struct {
	AccountNumber string
}

func NewListStandingOrder(param Params) string {
	return fmt.Sprintf(
		`<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cbes="http://temenos.com/CBESUPERAPP">
    <soapenv:Header/>
    <soapenv:Body>
        <cbes:ListStandingOrders>
            <WebRequestCommon>
                <company></company>
                <password>%s</password>
                <userName>%s</userName>
            </WebRequestCommon>
            <ACCTSTOLISTSUPERAPPType>
                <enquiryInputCollection>
                    <columnName>ID</columnName>
                    <criteriaValue>%s</criteriaValue>
                    <operand>CT</operand>
                </enquiryInputCollection>
            </ACCTSTOLISTSUPERAPPType>
        </cbes:ListStandingOrders>
    </soapenv:Body>
</soapenv:Envelope>`, param.Password, param.Username, param.AccountNumber)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	ListStandingOrderResponse *ListStandingOrderResponse `xml:"ListStandingOrdersResponse"`
}

type ListStandingOrderResponse struct {
	Status *struct {
		SuccessIndicator string `xml:"successIndicator"`
	} `xml:"Status"`
	ListStandingOrderType *struct {
		Group *struct {
			Details []ListStandingOrderDetail `xml:"mACCTSTOLISTSUPERAPPDetailType"`
		} `xml:"gACCTSTOLISTSUPERAPPDetailType"`
	} `xml:"ACCTSTOLISTSUPERAPPType"`
}

type ListStandingOrderDetail struct {
	StandingOrderId         string `xml:"StandingOrderID"`
	OrderType               string `xml:"KTYPE"`
	PaymentDetail           string `xml:"PAYMENTDETAILS"`
	Currency                string `xml:"CURRENCY"`
	Amount                  string `xml:"CURRENTAMOUNTBAL"`
	Frequency               string `xml:"CURRENTFREQUENCY"`
	DebitAccountNumber      string `xml:"DebitAccount"`
	DebitAccountHolderName  string `xml:"DEBITACCTDESC"`
	CreditAccountNumber     string `xml:"CPTYACCTNO"`
	CreditAccountHolderName string `xml:"TOACCTNAME"`
	CurrentDate             string `xml:"CURRENTENDDATE"`
}

type ListStandingOrderResult struct {
	Success bool
	Details []ListStandingOrderDetail
	Message []string
}

func ParseListStandingOrderSOAP(xmlData string) (*ListStandingOrderResult, error) {
	var env Envelope
	if err := xml.Unmarshal([]byte(xmlData), &env); err != nil {
		return nil, err
	}

	if env.Body.ListStandingOrderResponse != nil {
		resp := env.Body.ListStandingOrderResponse
		if resp.Status == nil {
			return &ListStandingOrderResult{
				Success: false,
				Message: []string{"Missing Status"},
			}, nil
		}

		if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
			return &ListStandingOrderResult{
				Success: false,
				Message: []string{"API returned failure"},
			}, nil
		}

		if resp.ListStandingOrderType == nil ||
			resp.ListStandingOrderType.Group == nil ||
			len(resp.ListStandingOrderType.Group.Details) == 0 {
			return &ListStandingOrderResult{
				Success: true,
				Message: []string{"No Standing Order Found!"},
			}, nil
		}

		return &ListStandingOrderResult{
			Success: true,
			Details: resp.ListStandingOrderType.Group.Details,
		}, nil
	}

	return &ListStandingOrderResult{
		Success: false,
		Message: []string{"Invalid response format"},
	}, nil
}
