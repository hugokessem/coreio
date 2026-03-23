package statuscheck

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

type StatusCheckParam struct {
	TransactionID string
}

func NewStatusCheck(param Params) string {
	return fmt.Sprintf(`<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cbes="http://temenos.com/CBESUPERAPP">
   <soapenv:Header/>
   <soapenv:Body>
      <cbes:SuperAppTxnStatus>
         <WebRequestCommon>
            <company/>
            <password>%s</password>
            <userName>%s</userName>
         </WebRequestCommon>
         <TXNSTATUSSUPERAPPType>
            <enquiryInputCollection>
               <columnName>ID</columnName>
               <criteriaValue>%s</criteriaValue>
               <operand>EQ</operand>
            </enquiryInputCollection>
         </TXNSTATUSSUPERAPPType>
      </cbes:SuperAppTxnStatus>
   </soapenv:Body>
</soapenv:Envelope>
`, param.Password, param.Username, param.TransactionID)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	SuperAppTxnStatusResponse *SuperAppTxnStatusResponse `xml:"SuperAppTxnStatusResponse"`
}

type SuperAppTxnStatusResponse struct {
	Status *Status `xml:"Status"`
}

type Status struct {
	SuccessIndicator string   `xml:"successIndicator"`
	Messages         []string `xml:"messages"`
}

type StatusCheckResult struct {
	Success  bool
	Messages []string
}

func ParseStatusCheckSOAP(xmlData string) (*StatusCheckResult, error) {
	var env Envelope
	if err := xml.Unmarshal([]byte(xmlData), &env); err != nil {
		return nil, err
	}

	if env.Body.SuperAppTxnStatusResponse == nil {
		return &StatusCheckResult{
			Success:  false,
			Messages: []string{"Invalid response type"},
		}, nil
	}

	resp := env.Body.SuperAppTxnStatusResponse
	if resp.Status == nil {
		return &StatusCheckResult{
			Success:  false,
			Messages: []string{"Missing Status"},
		}, nil
	}

	success := strings.EqualFold(resp.Status.SuccessIndicator, "success")
	messages := resp.Status.Messages
	if len(messages) == 0 {
		if success {
			messages = []string{}
		} else {
			messages = []string{"API returned failure"}
		}
	}

	return &StatusCheckResult{
		Success:  success,
		Messages: messages,
	}, nil
}
