package lockedamountrelease

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

type ReleaseLockedAmountParam struct {
	TransactionID string
}

func NewReleaseLockedAmount(param Params) string {
	return fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cbes="http://temenos.com/CBESUPERAPP">
    <soapenv:Header/>
    <soapenv:Body>
        <cbes:ReleaseAccountLock>
            <WebRequestCommon>
                <company/>
                <password>%s</password>
                <userName>%s</userName>
            </WebRequestCommon>
            <OfsFunction></OfsFunction>
            <ACLOCKEDEVENTSRELEASELOCKSUPERAPPType>
                <transactionId>%s</transactionId>
            </ACLOCKEDEVENTSRELEASELOCKSUPERAPPType>
        </cbes:ReleaseAccountLock>
    </soapenv:Body>
</soapenv:Envelope>
	`, param.Password, param.Username, param.TransactionID)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	ReleaseLockedAmountResponse *ReleaseLockedAmountResponse `xml:"ReleaseAccountLockResponse"`
}

type ReleaseLockedAmountResponse struct {
	Status *struct {
		SuccessIndicator string `xml:"successIndicator"`
		TransactionID    string `xml:"transactionId"`
		Application      string `xml:"application"`
		MessageId        string `xml:"messageId"`
	} `xml:"Status"`
	ReleaseLockedAmountDetail *ReleaseLockedAmountDetail `xml:"ACLOCKEDEVENTSType"`
}

type ReleaseLockedAmountDetail struct {
	To            string   `xml:"TODATE"`
	From          string   `xml:"FROMDATE"`
	XMLName       xml.Name `xml:"ACLOCKEDEVENTSType"`
	TransactionID string   `xml:"id,attr"`
	Description   string   `xml:"DESCRIPTION"`
	LockedAmount  string   `xml:"LOCKEDAMOUNT"`
	AccountNumber string   `xml:"ACCOUNTNUMBER"`
	Status        string   `xml:"RECORDSTATUS"`
}

type ReleaseAccountLockedResult struct {
	Success  bool
	Detail   *ReleaseLockedAmountDetail
	Messages []string
}

func ParseCancleLockedAmountSOAP(xmlData string) (*ReleaseAccountLockedResult, error) {
	var env Envelope
	if err := xml.Unmarshal([]byte(xmlData), &env); err != nil {
		return nil, err
	}

	if env.Body.ReleaseLockedAmountResponse != nil {
		resp := env.Body.ReleaseLockedAmountResponse
		if resp.Status == nil {
			return &ReleaseAccountLockedResult{
				Success:  false,
				Messages: []string{"Missing Status"},
			}, nil
		}

		if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
			return &ReleaseAccountLockedResult{
				Success:  false,
				Messages: []string{"API returned failure"},
			}, nil
		}

		if resp.ReleaseLockedAmountDetail == nil {
			return &ReleaseAccountLockedResult{
				Success:  false,
				Messages: []string{},
			}, nil
		}

		return &ReleaseAccountLockedResult{
			Success: true,
			Detail: &ReleaseLockedAmountDetail{
				To:            resp.ReleaseLockedAmountDetail.To,
				From:          resp.ReleaseLockedAmountDetail.From,
				TransactionID: resp.ReleaseLockedAmountDetail.TransactionID,
				LockedAmount:  resp.ReleaseLockedAmountDetail.LockedAmount,
				AccountNumber: resp.ReleaseLockedAmountDetail.AccountNumber,
				Description:   resp.ReleaseLockedAmountDetail.Description,
				Status:        resp.ReleaseLockedAmountDetail.Status,
			},
		}, nil
	}

	return &ReleaseAccountLockedResult{
		Success:  false,
		Messages: []string{"Invalid response type"},
	}, nil
}
