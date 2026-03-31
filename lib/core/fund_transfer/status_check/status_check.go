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
	SuperAppTxnStatusResponse *StatusCheckResponse `xml:"SuperAppTxnStatusResponse"`
}

type StatusCheckResponse struct {
	Status *struct {
		SuccessIndicator string   `xml:"successIndicator"`
		Messages         []string `xml:"messages"`
	} `xml:"Status"`
	FundTransferStatusCheck *struct {
		OuterSuperappDetailType struct {
			InnerSuperappDetailType StatusCheckDetail `xml:"mTXNSTATUSSUPERAPPDetailType"`
		} `xml:"gTXNSTATUSSUPERAPPDetailType"`
	} `xml:"TXNSTATUSSUPERAPPType"`
}

type StatusCheckDetail struct {
	ServiceCode   string `xml:"SERVICECODE"`
	DebitAccount  string `xml:"DEBITACCOUNT"`
	CreditAccount string `xml:"CREDITACCOUNT"`
	DebitCurrency string `xml:"TXNCURRENCY"`
	DebitAmount   string `xml:"TXNAMOUNT"`
	Channel       string `xml:"CHANNEL"`
	FTReference   string `xml:"CBEREFERENCE"`
}

type StatusCheckResult struct {
	Success  bool
	Detail   *StatusCheckDetail
	Messages []string
}

func ParseStatusCheckSOAP(response string) (*StatusCheckResult, error) {
	var env Envelope
	if err := xml.Unmarshal([]byte(response), &env); err != nil {
		return nil, err
	}

	if env.Body.SuperAppTxnStatusResponse != nil {
		resp := env.Body.SuperAppTxnStatusResponse
		if resp.Status == nil {
			return &StatusCheckResult{
				Success:  false,
				Messages: []string{"Missing Status"},
			}, nil
		}
		if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
			return &StatusCheckResult{
				Success:  false,
				Messages: resp.Status.Messages,
			}, nil
		}

		if resp.FundTransferStatusCheck != nil {
			detail := resp.FundTransferStatusCheck.OuterSuperappDetailType.InnerSuperappDetailType
			return &StatusCheckResult{
				Success: true,
				Detail: &StatusCheckDetail{
					ServiceCode:   detail.ServiceCode,
					DebitAccount:  detail.DebitAccount,
					CreditAccount: detail.CreditAccount,
					DebitCurrency: detail.DebitCurrency,
					DebitAmount:   detail.DebitAmount,
					Channel:       detail.Channel,
					FTReference:   detail.FTReference,
				},
			}, nil
		}
	}

	return &StatusCheckResult{}, nil
}
