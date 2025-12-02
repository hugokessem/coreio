package cardrequest

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Params struct {
	Username      string
	Password      string
	AccountNumner string
	BranchCode    string
	PhoneNumber   string
	CardType      string
}

type CardRequestParam struct {
	AccountNumber string
	BranchCode    string
	PhoneNumber   string
	CardType      string
}

func NewCardRequest(param Params) string {
	return fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"
xmlns:cbes="http://temenos.com/CBESUPERAPP"
xmlns:atm="http://temenos.com/ATMCARDREGDETCARDREQSUPERAPP">
<soapenv:Header/>
<soapenv:Body>
<cbes:ATMCardNewRequest>
<WebRequestCommon>
<company/>
<password>%s</password>
<userName>%s</userName>
</WebRequestCommon>
<OfsFunction/>
<ATMCARDREGDETCARDREQSUPERAPPType id="">
<atm:ACCOUNT>%s</atm:ACCOUNT>
<atm:BRANCHCODE>%s</atm:BRANCHCODE>
<atm:PHONENO>%s</atm:PHONENO>
<atm:CARDTYPE>%s</atm:CARDTYPE>
</ATMCARDREGDETCARDREQSUPERAPPType>
</cbes:ATMCardNewRequest>
</soapenv:Body>
</soapenv:Envelope>
	`, param.Password, param.Username, param.AccountNumner, param.BranchCode, param.PhoneNumber, param.CardType)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	ATMCardNewRequestResponse *ATMCardNewRequestResponse `xml:"ATMCardNewRequestResponse"`
}

type ATMCardNewRequestResponse struct {
	Status *struct {
		SuccessIndicator string `xml:"successIndicator"`
		MessageId        string `xml:"messageId"`
		Application      string `xml:"application"`
		TransactionId    string `xml:"transactionId"`
	} `xml:"Status"`
	ATMCardRequestDetail *ATMCardRequestDetail `xml:"ATMCARDREGDETType"`
}

type ATMCardRequestDetail struct {
	AccountNumber      string `xml:"ACCOUNT"`
	AccouuntHolderName string `xml:"ACCOUNTTITLE1"`
	Address            string `xml:"ADDRESS"`
	BranchCode         string `xml:"BRANCHCODE"`
	OpenDate           string `xml:"ACOPENDATE"`
	Residence          string `xml:"RESIDENCE"`
	Industry           string `xml:"INDUSTRY"`
	PhoneNumber        string `xml:"PHONENO"`
	CardType           string `xml:"CARDTYPE"`
	Sex                string `xml:"SEX"`
	CivilStatus        string `xml:"CIVILSTATUS"`
	HolderNo           string `xml:"HOLDERNO"`
	CardNumber         string `xml:"CURRNO"`
	GDatetime          struct {
		DateTime string `xml:"DATETIME"`
	} `xml:"gDATETIME"`
	Authoriser  string `xml:"AUTHORISER"`
	CoCode      string `xml:"COCODE"`
	DeptCode    string `xml:"DEPTCODE"`
	Date        string `xml:"DATE"`
	RequestType string `xml:"REQUESTTYPE"`
	ProductType string `xml:"PRODUCTTYPE"`
	VirtualFlag string `xml:"VIRTUALFLAG"`
}

type CardRequestResult struct {
	Success  bool
	Detail   *ATMCardRequestDetail
	Messages []string
}

func ParseATMCardRequestSOAP(xmlData string) (*CardRequestResult, error) {
	var env Envelope
	err := xml.Unmarshal([]byte(xmlData), &env)
	if err != nil {
		return nil, err
	}

	if env.Body.ATMCardNewRequestResponse == nil {
		resp := env.Body.ATMCardNewRequestResponse
		if resp.Status == nil {
			return &CardRequestResult{
				Success:  false,
				Messages: []string{"Missing Status"},
			}, nil
		}
		if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
			return &CardRequestResult{
				Success:  false,
				Messages: []string{"API returned failure"},
			}, nil
		}
		if resp.ATMCardRequestDetail == nil {
			return &CardRequestResult{
				Success:  false,
				Messages: []string{"Missing ATMCardRequestDetail"},
			}, nil
		}

		return &CardRequestResult{
			Success: true,
			Detail:  resp.ATMCardRequestDetail,
		}, nil
	}

	return &CardRequestResult{
		Success: true,
		Detail:  env.Body.ATMCardNewRequestResponse.ATMCardRequestDetail,
	}, nil
}
