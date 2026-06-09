package statuschange

import (
	"encoding/xml"
	"fmt"
	"strings"

	valueobject "github.com/hugokessem/coreio/lib/core/super_app/status_change/value_object"
)

type Param struct {
	Username string
	Password string
	UserCode string
	Status   valueobject.Status
}

type StatusParam struct {
	UserCode string
	Status   string
}

func NewStatusChange(param Param) (string, error) {
	if err := param.Status.IsValid(); err != nil {
		return "", err
	}

	return fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cbes="http://temenos.com/CBESUPERAPP" xmlns:sup="http://temenos.com/SUPERAPPUSERSTATUS">
    <soapenv:Header/>
    <soapenv:Body>
        <cbes:SuperAppUserChangeStatus>
            <WebRequestCommon>
                <company/>
                <password>%s</password>
                <userName>%s</userName>
            </WebRequestCommon>
            <OfsFunction></OfsFunction>
            <SUPERAPPUSERSTATUSType id="%s">
                <sup:STATUS>%s</sup:STATUS>
            </SUPERAPPUSERSTATUSType>
        </cbes:SuperAppUserChangeStatus
    </soapenv:Body>
</soapenv:Envelope>
	`, param.Password, param.Username, param.UserCode, param.Status.String()), nil
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	SuperAppUserChangeStatusResponse *SuperAppUserChangeStatusResponse `xml:"SuperAppUserChangeStatusResponse"`
}

type SuperAppUserChangeStatusResponse struct {
	Status *struct {
		TransactionID    string   `xml:"transactionId"`
		Messages         []string `xml:"messages"`
		SuccessIndicator string   `xml:"successIndicator"`
		Application      string   `xml:"application"`
	} `xml:"Status"`
	SuperAppUserChangeStatusDetail *SuperAppUserChangeStatusDetail `xml:"SUPERAPPUSERType"`
}

type SuperAppUserChangeStatusDetail struct {
	CustomerName    string `xml:"CUSTOMERNAME"`
	CustomerId      string `xml:"CUSTOMERID"`
	Status          string `xml:"STATUS"`
	PhoneNumber     string `xml:"PHONENO"`
	Email           string `xml:"EMAIL"`
	Channel         string `xml:"CHANNEL"`
	Description     string `xml:"DESCRIPTION"`
	LastUpdatedDate string `xml:"LASTUPDATEDDATE"`
	BranchCode      string `xml:"BRANCHCODE"`
	SalesId         string `xml:"SALESID"`
	CurrNo          string `xml:"CURRNO"`
	Datetime        struct {
		DateTime string `xml:"DATETIME"`
	} `xml:"gDATETIME"`
	Authoriser string `xml:"AUTHORISER"`
	Code       string `xml:"COCODE"`
	DeptCode   string `xml:"DEPTCODE"`
}

type StatusChangeResult struct {
	Success  bool
	Detail   *SuperAppUserChangeStatusDetail
	Messages []string
}

func ParseSuperAppUserChangeStatusResponseSOAP(xmlData string) (*StatusChangeResult, error) {
	var env Envelope
	err := xml.Unmarshal([]byte(xmlData), &env)
	if err != nil {
		return nil, err
	}
	if env.Body.SuperAppUserChangeStatusResponse != nil {
		resp := env.Body.SuperAppUserChangeStatusResponse
		if resp.Status == nil {
			return &StatusChangeResult{
				Success:  false,
				Messages: []string{"Missing Status"},
			}, nil
		}

		if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
			return &StatusChangeResult{
				Success:  false,
				Messages: resp.Status.Messages,
			}, nil
		}

		if resp.SuperAppUserChangeStatusDetail == nil {
			return &StatusChangeResult{
				Success:  false,
				Messages: []string{"Missing SuperAppUserChangeStatusDetail"},
			}, nil
		}

		return &StatusChangeResult{
			Success:  true,
			Detail:   resp.SuperAppUserChangeStatusDetail,
			Messages: resp.Status.Messages,
		}, nil
	}
	return &StatusChangeResult{
		Success:  false,
		Messages: []string{"Invalid response type"},
	}, nil
}
