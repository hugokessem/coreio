package standingorderlisthistory

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

type ListStandingOrderHistoryParams struct {
	AccountNumber string
}

func NewListStandingOrderHistory(param Params) string {
	return fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cbes="http://temenos.com/CBESUPERAPP">
   <soapenv:Header/>
   <soapenv:Body>
      <cbes:StandingOrderHistorylistbyAc>
         <WebRequestCommon>
            <company></company>
            <password>%s</password>
            <userName>%s</userName>
         </WebRequestCommon>
         <ACCTSTOLISTHISSUPERAPPType>
            <enquiryInputCollection>
               <columnName>ACCOUNT</columnName>
               <criteriaValue>%s</criteriaValue>
               <operand>CT</operand>
            </enquiryInputCollection>
         </ACCTSTOLISTHISSUPERAPPType>
      </cbes:StandingOrderHistorylistbyAc>
   </soapenv:Body>
</soapenv:Envelope>`, param.Password, param.Username, param.AccountNumber)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	StandingOrderHistorylistbyAcResponse *StandingOrderHistorylistbyAcResponse `xml:"StandingOrderHistorylistbyAcResponse"`
}

type StandingOrderHistorylistbyAcResponse struct {
	Status *struct {
		SuccessIndicator string `xml:"successIndicator"`
	} `xml:"Status"`
	ACCTSTOLISTHISSUPERAPPType *struct {
		Group *struct {
			Details []StandingOrderHistorylistbyAcDetail `xml:"mACCTSTOLISTHISSUPERAPPDetailType"`
		} `xml:"gACCTSTOLISTHISSUPERAPPDetailType"`
	} `xml:"ACCTSTOLISTHISSUPERAPPType"`
}

type StandingOrderHistorylistbyAcDetail struct {
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
	CurrentFrequencyDate    string `xml:"CURRFREQDATE"`
	EndDate                 string `xml:"CURRENTENDDATE"`
	StartDate               string `xml:"STOSTARTDATE"`
}

type StandingOrderListHistoryResult struct {
	Success bool
	Details []StandingOrderHistorylistbyAcDetail
	Message []string
}

func ParseStandingOrderListHistorySOAP(xmlData string) (*StandingOrderListHistoryResult, error) {
	var env Envelope
	if err := xml.Unmarshal([]byte(xmlData), &env); err != nil {
		return nil, err
	}

	if env.Body.StandingOrderHistorylistbyAcResponse != nil {
		resp := env.Body.StandingOrderHistorylistbyAcResponse
		if resp.Status == nil {
			return &StandingOrderListHistoryResult{
				Success: false,
				Message: []string{"Missing Status"},
			}, nil
		}
		if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
			return &StandingOrderListHistoryResult{
				Success: false,
				Message: []string{"API returned failure"},
			}, nil
		}
		if resp.ACCTSTOLISTHISSUPERAPPType == nil || resp.ACCTSTOLISTHISSUPERAPPType.Group == nil || len(resp.ACCTSTOLISTHISSUPERAPPType.Group.Details) == 0 {
			return &StandingOrderListHistoryResult{
				Success: true,
				Message: []string{"No Standing Order History Found!"},
			}, nil
		}
		details := resp.ACCTSTOLISTHISSUPERAPPType.Group.Details
		detailsList := make([]StandingOrderHistorylistbyAcDetail, len(details))
		for i, detail := range details {
			detailsList[i] = StandingOrderHistorylistbyAcDetail{
				StandingOrderId:         detail.StandingOrderId,
				OrderType:               detail.OrderType,
				PaymentDetail:           detail.PaymentDetail,
				Currency:                detail.Currency,
				Amount:                  detail.Amount,
				Frequency:               detail.Frequency,
				DebitAccountNumber:      detail.DebitAccountNumber,
				DebitAccountHolderName:  detail.DebitAccountHolderName,
				CreditAccountNumber:     detail.CreditAccountNumber,
				CreditAccountHolderName: detail.CreditAccountHolderName,
				CurrentFrequencyDate:    detail.CurrentFrequencyDate,
				EndDate:                 detail.EndDate,
				StartDate:               detail.StartDate,
			}
		}
		return &StandingOrderListHistoryResult{
			Success: true,
			Details: detailsList,
		}, nil
	}

	return &StandingOrderListHistoryResult{
		Success: false,
		Message: []string{"Invalid response type"},
	}, nil
}
