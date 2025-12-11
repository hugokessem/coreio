package customerlimit

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type Params struct {
	Username      string
	Password      string
	TransactionID string
}

type CustomerLimitParams struct {
	TransactionID string
}

func NewCustomerLimit(param Params) string {
	return fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"
xmlns:cbes="http://temenos.com/CBESUPERAPP">
    <soapenv:Header/>
    <soapenv:Body>
        <cbes:CustomerLimitView>
            <WebRequestCommon>
                <company/>
                <password>%s</password>
                <userName>%s</userName>
            </WebRequestCommon>
            <CUSTOMERLIMITVIEWType>
                <transactionId>%s</transactionId>
            </CUSTOMERLIMITVIEWType>
        </cbes:CustomerLimitView>
    </soapenv:Body>
</soapenv:Envelope>`, param.Password, param.Username, param.TransactionID)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	CustomerLimitResponse *CustomerLimitResponse `xml:"CustomerLimitViewResponse"`
}

type CustomerLimitResponse struct {
	Status *struct {
		SuccessIndicator string `xml:"successIndicator"`
		TransactionID    string `xml:"transactionId"`
		MessageID        string `xml:"messageId"`
		Application      string `xml:"application"`
	} `xml:"Status"`
	CustomerLimitType *CustomerLimitDetail `xml:"CUSTOMERLIMITType"`
}

type CustomerLimitDetail struct {
	ChannelTypeGroup *struct {
		Details []ChannelType `xml:"mCHANNELTYPE"`
	} `xml:"gCHANNELTYPE"`
	ConvRoleTypeGroup *struct {
		Details []ConvRoleType `xml:"mCONVROLETYPE"`
	} `xml:"gCONVROLETYPE"`
	IfbRoleTypeGroup *struct {
		Details []IfbRoleType `xml:"mIFBROLETYPE"`
	} `xml:"gIFBROLETYPE"`
	CustomerMaxLimit string `xml:"CUSTMAXLIMIT"`
	CustomerMinLimit string `xml:"CUSTMINLIMIT"`
	CustomerCount    string `xml:"CUSTCOUNT"`
	AccountMaxLimit  string `xml:"ACCTMAXLIMIT"`
	AccountMinLimit  string `xml:"ACCTMINLIMIT"`
	AccountCount     string `xml:"ACCTCOUNT"`
	CurrencyNo       string `xml:"CURRNO"`
	GlobalDateTime   struct {
		DateTime string `xml:"DATETIME"`
	} `xml:"gDATETIME"`
	GlobalInputter struct {
		Inputter string `xml:"INPUTTER"`
	} `xml:"gINPUTTER"`
	Authoriser string `xml:"AUTHORISER"`
	CoCode     string `xml:"COCODE"`
	DeptCode   string `xml:"DEPTCODE"`
}

type ChannelType struct {
	ChannelType     string `xml:"CHANNELTYPE"`
	ChannelMaxLimit string `xml:"CHANNELMAXLIMIT"`
	ChannelMinLimit string `xml:"CHANNELMINLIMIT"`
}

type ConvRoleType struct {
	ConvRoleType     string `xml:"CONVROLETYPE"`
	ConvRoleMaxLimit string `xml:"CONVROLEMAXLIMIT"`
	ConvRoleMinLimit string `xml:"CONVROLEMINLIMIT"`
}

type IfbRoleType struct {
	IfbRoleType     string `xml:"IFBROLETYPE"`
	IfbRoleMaxLimit string `xml:"IFBROLEMAXLIMIT"`
	IfbRoleMinLimit string `xml:"IFBROLEMINLIMIT"`
}
type CustomerLimitResult struct {
	Success bool
	Detail  *CustomerLimitDetail
	Message []string
}

func ParseCustomerLimitSOAP(xmlData string) (*CustomerLimitResult, error) {
	var env Envelope
	err := xml.Unmarshal([]byte(xmlData), &env)
	if err != nil {
		return nil, err
	}
	if env.Body.CustomerLimitResponse != nil {
		resp := env.Body.CustomerLimitResponse
		if resp.Status == nil {
			return &CustomerLimitResult{
				Success: false,
				Message: []string{"Missing Status"},
			}, nil
		}
		if resp.Status.SuccessIndicator != "Success" {
			return &CustomerLimitResult{
				Success: false,
				Message: []string{"API returned failure"},
			}, nil
		}
		if resp.CustomerLimitType == nil {
			return &CustomerLimitResult{
				Success: false,
				Message: []string{"No details found"},
			}, nil
		}

		return &CustomerLimitResult{
			Success: true,
			Detail: &CustomerLimitDetail{
				ChannelTypeGroup:  resp.CustomerLimitType.ChannelTypeGroup,
				ConvRoleTypeGroup: resp.CustomerLimitType.ConvRoleTypeGroup,
				IfbRoleTypeGroup:  resp.CustomerLimitType.IfbRoleTypeGroup,
				CustomerMaxLimit:  resp.CustomerLimitType.CustomerMaxLimit,
				CustomerMinLimit:  resp.CustomerLimitType.CustomerMinLimit,
				CustomerCount:     resp.CustomerLimitType.CustomerCount,
				AccountMaxLimit:   resp.CustomerLimitType.AccountMaxLimit,
				AccountMinLimit:   resp.CustomerLimitType.AccountMinLimit,
				AccountCount:      resp.CustomerLimitType.AccountCount,
				CurrencyNo:        resp.CustomerLimitType.CurrencyNo,
				GlobalDateTime:    resp.CustomerLimitType.GlobalDateTime,
				GlobalInputter:    resp.CustomerLimitType.GlobalInputter,
				Authoriser:        resp.CustomerLimitType.Authoriser,
				CoCode:            resp.CustomerLimitType.CoCode,
				DeptCode:          resp.CustomerLimitType.DeptCode,
			},
		}, nil
	}
	return nil, errors.New("invalid response type")
}
