package customerlimitfetchbycifreturnservice

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Params struct {
	Username       string
	Password       string
	CustomerNumber string
}

type CustomerLimitFetchByCIFReturnServiceParam struct {
	CustomerNumber string
}

func NewCustomerLimitFetchByCIFReturnService(param Params) string {
	return fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cbes="http://temenos.com/CBESUPERAPP">
    <soapenv:Header/>
    <soapenv:Body>
        <cbes:ViewCustomedLimit>
            <WebRequestCommon>
                <company/>
                <password>%s</password>
                <userName>%s</userName>
            </WebRequestCommon>
            <CUSTOMLIMITAMENDENQType>
                <enquiryInputCollection>
                    <columnName>ID</columnName>
                    <criteriaValue>%s</criteriaValue>
                    <operand>EQ</operand>
                </enquiryInputCollection>
            </CUSTOMLIMITAMENDENQType>
        </cbes:ViewCustomedLimit>
    </soapenv:Body>
</soapenv:Envelope>`, param.Password, param.Username, param.CustomerNumber)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	ViewCustomedLimitResponse *ViewCustomedLimitResponse `xml:"ViewCustomedLimitResponse"`
}

type ViewCustomedLimitResponse struct {
	Status *struct {
		SuccessIndicator string   `xml:"successIndicator"`
		Messages         []string `xml:"messages"`
		Application      string   `xml:"application"`
		TransactionId    string   `xml:"transactionId"`
	} `xml:"Status"`
	CustomerLimitType *struct {
		GCustomerLimitDetailType struct {
			CustomerLimitDetailType []CustomerLimitDetailType `xml:"mCUSTOMLIMITAMENDENQDetailType"`
		} `xml:"gCUSTOMLIMITAMENDENQDetailType"`
	} `xml:"CUSTOMLIMITAMENDENQType"`
}

type CustomerLimitDetailType struct {
	CIF         string `xml:"CIF"`
	Channel     string `xml:"Channel"`
	ServiceCode string `xml:"ServiceCode"`
	ServiceName string `xml:"ServiceName"`
	Limit       string `xml:"Limit"`
	Count       string `xml:"Count"`
}

type CustomerLimitFetchByCIFReturnServiceResult struct {
	Success  bool
	Detail   []CustomerLimitDetailType
	Messages []string
}

func ParseCustomerLimitFetchByCIFReturnServiceSOAP(xmlData string) (*CustomerLimitFetchByCIFReturnServiceResult, error) {
	var env Envelope
	if err := xml.Unmarshal([]byte(xmlData), &env); err != nil {
		return nil, err
	}
	if env.Body.ViewCustomedLimitResponse != nil {
		resp := env.Body.ViewCustomedLimitResponse
		if resp.Status == nil {
			return &CustomerLimitFetchByCIFReturnServiceResult{
				Success:  false,
				Messages: []string{"missing status"},
			}, nil
		}
		if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
			return &CustomerLimitFetchByCIFReturnServiceResult{
				Success:  false,
				Messages: resp.Status.Messages,
			}, nil
		}

		detail := resp.CustomerLimitType.GCustomerLimitDetailType.CustomerLimitDetailType
		if len(detail) == 0 {
			return &CustomerLimitFetchByCIFReturnServiceResult{
				Success:  false,
				Messages: []string{"missing detail"},
			}, nil
		}

		return &CustomerLimitFetchByCIFReturnServiceResult{
			Success:  true,
			Detail:   detail,
			Messages: resp.Status.Messages,
		}, nil
	}

	return &CustomerLimitFetchByCIFReturnServiceResult{
		Success:  false,
		Messages: []string{"invalid response"},
	}, nil
}
