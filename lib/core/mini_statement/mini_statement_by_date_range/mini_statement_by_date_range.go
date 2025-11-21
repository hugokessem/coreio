package ministatementbydaterange

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Params struct {
	Username      string
	Password      string
	AccountNumber string
	From          string
	To            string
}
type MiniStatementByDateRangeParam struct {
	AccountNumber string
	From          string
	To            string
}

func NewMiniStatementByDateRange(param Params) string {
	return fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"
xmlns:cbes="http://temenos.com/CBESUPERAPP">
    <soapenv:Header/>
    <soapenv:Body>
        <cbes:AccountStatementByRange>
            <WebRequestCommon>
                <company/>
                <password>%s</password>
                <userName>%s</userName>
            </WebRequestCommon>
            <ACCTSTMTRGSUPERAPPType>
                <enquiryInputCollection>
                    <columnName>ACCOUNT</columnName>
                    <criteriaValue>%s</criteriaValue>
                    <operand>EQ</operand>
                </enquiryInputCollection>
                <enquiryInputCollection>
                    <columnName>BOOKING.DATE</columnName>
                    <criteriaValue>%s %s</criteriaValue>
                    <operand>EQ</operand>
                </enquiryInputCollection>
            </ACCTSTMTRGSUPERAPPType>
        </cbes:AccountStatementByRange>
    </soapenv:Body>
</soapenv:Envelope>`, param.Password, param.Username, param.AccountNumber, param.From, param.To)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	AccountStatementByRangeResponse *AccountStatementByRangeResponse `xml:"AccountStatementByRangeResponse"`
}

type AccountStatementByRangeResponse struct {
	Status *struct {
		SuccessIndicator string `xml:"successIndicator"`
	} `xml:"Status"`
	MiniStatementResponse *MiniStatementResponse `xml:"ACCTSTMTRGSUPERAPPType"`
}

type MiniStatementResponse struct {
	AccountNumber   string `xml:"DACC"`
	CustomerNumber  string `xml:"CUS"`
	Currency        string `xml:"DCCY"`
	StartingBalance string `xml:"AMTBFWD"`
	EndBalance      string `xml:"TOTAL"`
	Group           *struct {
		Details []MiniStatementDetail `xml:"mACCTSTMTRGSUPERAPPDetailType"`
	} `xml:"gACCTSTMTRGSUPERAPPDetailType"`
}

type MiniStatementDetail struct {
	Date                 string `xml:"VALDESC"`
	Order                string `xml:"PDESC"`
	TransactionReference string `xml:"REFNO"`
	Post                 string `xml:"POST"`
	Amount               string `xml:"PAMT"`
}

type MiniStatementByDateRangeResult struct {
	Success bool
	Detail  *MiniStatementResponse
	Message []string
}

func ParseMiniStatementByDateRangeSOAP(xmlData string) (*MiniStatementByDateRangeResult, error) {
	var env Envelope
	err := xml.Unmarshal([]byte(xmlData), &env)
	if err != nil {
		return nil, err
	}

	if env.Body.AccountStatementByRangeResponse != nil {
		resp := env.Body.AccountStatementByRangeResponse

		if resp.Status == nil {
			return &MiniStatementByDateRangeResult{
				Success: false,
				Message: []string{"Missing Status"},
			}, nil
		}

		if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
			return &MiniStatementByDateRangeResult{
				Success: false,
				Message: []string{"API returned failure"},
			}, nil
		}

		if resp.MiniStatementResponse == nil {
			return &MiniStatementByDateRangeResult{
				Success: false,
				Message: []string{},
			}, nil
		}

		return &MiniStatementByDateRangeResult{
			Success: true,
			Detail:  resp.MiniStatementResponse,
		}, nil
	}

	return &MiniStatementByDateRangeResult{
		Success: false,
		Message: []string{"Invalid response type"},
	}, nil
}
