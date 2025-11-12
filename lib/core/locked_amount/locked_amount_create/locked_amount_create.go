package lockedamountcreate

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Params struct {
	Username      string
	Password      string
	AccountNumber string
	Description   string
	From          string
	To            string
	LockedAmount  string
}

type CreateLockedAmountParam struct {
	AccountNumber string
	Description   string
	From          string
	To            string
	LockedAmount  string
}

func NewCreateLockedAmount(param Params) string {
	return fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cbes="http://temenos.com/CBESUPERAPP" xmlns:acl="http://temenos.com/ACLOCKEDEVENTSCREATELOCKSUPERAPP">
    <soapenv:Header/>
    <soapenv:Body>
        <cbes:CreateAccountLock>
            <WebRequestCommon>
                <company/>
                <password>%s</password>
                <userName>%s</userName>
            </WebRequestCommon>
            <OfsFunction></OfsFunction>
            <ACLOCKEDEVENTSCREATELOCKSUPERAPPType id="">
                <acl:ACCOUNTNUMBER>%s</acl:ACCOUNTNUMBER>
                <acl:DESCRIPTION>%s</acl:DESCRIPTION>
                <acl:FROMDATE>%s</acl:FROMDATE>
                <acl:TODATE>%s</acl:TODATE>
                <acl:LOCKEDAMOUNT>%s</acl:LOCKEDAMOUNT>
            </ACLOCKEDEVENTSCREATELOCKSUPERAPPType>
        </cbes:CreateAccountLock>
    </soapenv:Body>
</soapenv:Envelope>`, param.Password, param.Username, param.AccountNumber, param.Description, param.From, param.To, param.LockedAmount)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	CreateLockedAmountResponse *CreateLockedAmountResponse `xml:"CreateAccountLockResponse"`
}

type CreateLockedAmountResponse struct {
	Status *struct {
		SuccessIndicator string `xml:"successIndicator"`
		TransactionID    string `xml:"transactionId"`
		Application      string `xml:"application"`
		MessageId        string `xml:"messageId"`
	} `xml:"Status"`
	CreateLockedAmount *CreateLockedAmountDetail `xml:"ACLOCKEDEVENTSType"`
}

type CreateLockedAmountDetail struct {
	AccountNumber string   `xml:"ACCOUNTNUMBER"`
	XMLName       xml.Name `xml:"ACLOCKEDEVENTSType"`
	TransactionID string   `xml:"id,attr"`
	Description   string   `xml:"DESCRIPTION"`
	From          string   `xml:"FROMDATE"`
	To            string   `xml:"TODATE"`
	LockedAmount  string   `xml:"LOCKEDAMOUNT"`
}

type CreateLockedAmountResult struct {
	Success  bool
	Detail   *CreateLockedAmountDetail
	Messages []string
}

func ParseCreateLockedAmountSOAP(xmlData string) (*CreateLockedAmountResult, error) {
	var env Envelope
	if err := xml.Unmarshal([]byte(xmlData), &env); err != nil {
		return nil, err
	}

	if env.Body.CreateLockedAmountResponse != nil {
		resp := env.Body.CreateLockedAmountResponse
		if resp.Status == nil {
			return &CreateLockedAmountResult{
				Success:  false,
				Messages: []string{"Missing Status"},
			}, nil
		}

		if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
			return &CreateLockedAmountResult{
				Success:  false,
				Messages: []string{"API returned failure"},
			}, nil
		}

		if resp.CreateLockedAmount == nil {
			return &CreateLockedAmountResult{
				Success:  false,
				Messages: []string{},
			}, nil
		}

		return &CreateLockedAmountResult{
			Success: true,
			Detail: &CreateLockedAmountDetail{
				To:            resp.CreateLockedAmount.To,
				From:          resp.CreateLockedAmount.From,
				Description:   resp.CreateLockedAmount.Description,
				LockedAmount:  resp.CreateLockedAmount.LockedAmount,
				AccountNumber: resp.CreateLockedAmount.AccountNumber,
				TransactionID: resp.CreateLockedAmount.TransactionID,
			},
		}, nil
	}

	return &CreateLockedAmountResult{
		Success:  false,
		Messages: []string{"Invalid response type"},
	}, nil
}
