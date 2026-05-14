package statuschange

import (
	"encoding/xml"
	"fmt"

	valueobject "github.com/hugokessem/coreio/lib/core/super_app/status_change/value_object"
)

type Param struct {
	Username       string
	Password       string
	SuperappUserID string
	Status         valueobject.Status
}

type StatusParam struct {
	SuperappUserID string
	Status         valueobject.Status
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
        </cbes:SuperAppUserChangeStatus>
    </soapenv:Body>
</soapenv:Envelope>
	`, param.Password, param.Username, param.SuperappUserID, param.Status.String()), nil
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	SuperAppUserChangeStatusResponse SuperAppUserChangeStatusResponse `xml:"SuperAppUserChangeStatusResponse"`
}

type SuperAppUserChangeStatusResponse struct {
	Status *struct {
		TransactionID    string   `xml:"transactionId"`
		Messages         []string `xml:"messages"`
		SuccessIndicator string   `xml:"successIndicator"`
		Application      string   `xml:"application"`
	} `xml:"Status"`
}
