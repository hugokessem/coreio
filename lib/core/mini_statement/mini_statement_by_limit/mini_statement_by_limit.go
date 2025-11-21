package ministatementbylimit

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Params struct {
	Username            string
	Password            string
	AccountNumber       string
	NumberOfTransaction string
}

type MiniStatementByLimitParams struct {
	AccountNumber       string
	NumberOfTransaction string
}

func NewMiniStatementByLimit(param Params) string {
	return fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cbes="http://temenos.com/CBESUPERAPP">
    <soapenv:Header/>
    <soapenv:Body>
        <cbes:AccountMiniStatement>
            <WebRequestCommon>
                <company/>
                <password>%s</password>
                <userName>%s</userName>
            </WebRequestCommon>
            <CBEMINISTMTENQType>
                <enquiryInputCollection>
                    <columnName>ACCOUNT</columnName>
                    <criteriaValue>%s</criteriaValue>
                    <operand>EQ</operand>
                </enquiryInputCollection>
                <enquiryInputCollection>
                    <columnName>NO.OF.TXNS</columnName>
                    <criteriaValue>%s</criteriaValue>
                    <operand>EQ</operand>
                </enquiryInputCollection>
            </CBEMINISTMTENQType>
        </cbes:AccountMiniStatement>
    </soapenv:Body>
</soapenv:Envelope>`, param.Password, param.Username, param.AccountNumber, param.NumberOfTransaction)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	MiniStatementResponse *MiniStatementResponse `xml:"AccountMiniStatementResponse"`
}

type MiniStatementResponse struct {
	Status *struct {
		SuccessIndicator string `xml:"successIndicator"`
	} `xml:"Status"`
	MiniStatementType *struct {
		Group *struct {
			Details []MiniStatementDetail `xml:"mCBEMINISTMTENQDetailType"`
		} `xml:"gCBEMINISTMTENQDetailType"`
	} `xml:"CBEMINISTMTENQType"`
}

type MiniStatementDetail struct {
	ValueDate            string `xml:"ValueDate"`
	Description          string `xml:"Description"`
	TransactionReference string `xml:"TxnReference"`
	Amount               string `xml:"Amount"`
	Currency             string `xml:"Currency"`
	OtherPartyAccount    string `xml:"OtherPartyAcc"`
	PaymentDetails       string `xml:"PaymentDetails"`
	DateTime             string `xml:"DateTime"`
}

type MiniStatementByLimitResult struct {
	Success bool
	Details []MiniStatementDetail
	Message []string
}

func ParseMiniStatementByLimitSOAP(xmlData string) (*MiniStatementByLimitResult, error) {
	var env Envelope
	err := xml.Unmarshal([]byte(xmlData), &env)
	if err != nil {
		return nil, err
	}

	if env.Body.MiniStatementResponse != nil {
		resp := env.Body.MiniStatementResponse
		if resp.Status == nil {
			return &MiniStatementByLimitResult{
				Success: false,
				Message: []string{"Missing Status"},
			}, nil
		}

		if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
			return &MiniStatementByLimitResult{
				Success: false,
				Message: []string{"API returned failure"},
			}, nil
		}

		if resp.MiniStatementType == nil {
			return &MiniStatementByLimitResult{
				Success: false,
				Message: []string{},
			}, nil
		}

		if resp.MiniStatementType.Group == nil {
			return &MiniStatementByLimitResult{
				Success: true,
				Details: []MiniStatementDetail{},
			}, nil
		}

		return &MiniStatementByLimitResult{
			Success: true,
			Details: resp.MiniStatementType.Group.Details,
		}, nil
	}

	return &MiniStatementByLimitResult{
		Success: false,
		Message: []string{"Invalid response type"},
	}, nil
}
