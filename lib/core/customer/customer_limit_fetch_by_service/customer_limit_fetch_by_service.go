package customerlimitfetchbyservice

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Params struct {
	Username    string
	Password    string
	ServiceCode string
}

type LimitFetchByService struct {
	ServiceCode string
}

func NewCustomerLimitFetchByService(param Params) string {
	return fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cbes="http://temenos.com/CBESUPERAPP">
   <soapenv:Header/>
   <soapenv:Body>
      <cbes:GenericLimitView>
         <WebRequestCommon>
            <company/>
            <password>%s</password>
            <userName>%s</userName>
         </WebRequestCommon>
         <CUSTOMERLIMITVIEWType>
            <transactionId>%s</transactionId>
         </CUSTOMERLIMITVIEWType>
      </cbes:GenericLimitView>
   </soapenv:Body>
</soapenv:Envelope>`, param.Password, param.Username, param.ServiceCode)
}

type Envelope struct {
	Body Body `xml:"Body"`
}

type Body struct {
	GenericLimitViewResponse GenericLimitViewResponse `xml:"GenericLimitViewResponse"`
}

type GenericLimitViewResponse struct {
	Status *struct {
		TransactionId    string `xml:"transactionId"`
		MessageId        string `xml:"messageId"`
		SuccessIndicator string `xml:"successIndicator"`
		Application      string `xml:"application"`
	} `xml:"Status"`
	Detail *CustomerLimitDetail `xml:"CUSTOMERLIMITType"`
}

type CustomerLimitDetail struct {
	GChannelType struct {
		MChannelType struct {
			CHANNELTYPE   string `xml:"CHANNELTYPE"`
			SGServiceType struct {
				GServiceType struct {
					GServiceType    string `xml:"GSERVICETYPE"`
					ChannelMaxLimit string `xml:"CHANNELMAXLIMIT"`
				} `xml:"GSERVICETYPE"`
			} `xml:"sgGSERVICETYPE"`
		} `xml:"mCHANNELTYPE"`
	} `xml:"gCHANNELTYPE"`
	ChargeCode        string `xml:"CHARGECODE"`
	TransactionCount  string `xml:"TXNCOUNT"`
	TransactionAmount string `xml:"TXNAMOUNT"`
	CurrentNo         string `xml:"CURRNO"`
	CoCode            string `xml:"COCODE"`
	DeptCode          string `xml:"DEPTCODE"`
	Authoriser        string `xml:"AUTHORISER"`
	GDateTime         struct {
		DateTime string `xml:"DATETIME"`
	} `xml:"gDATETIME"`
}

type CustomerLimitResult struct {
	Success  bool
	Detail   *CustomerLimitDetail
	Messages []string
}

func ParseCustomerLimitFetchByServiceSOAP(xmlData string) (*CustomerLimitResult, error) {
	var env Envelope
	if err := xml.Unmarshal([]byte(xmlData), &env); err != nil {
		return nil, err
	}

	if env.Body.GenericLimitViewResponse.Detail != nil {
		resp := env.Body.GenericLimitViewResponse

		if resp.Status == nil {
			return &CustomerLimitResult{
				Success:  false,
				Messages: []string{"Missing Status"},
			}, nil
		}

		if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
			return &CustomerLimitResult{
				Success:  false,
				Messages: []string{resp.Status.MessageId},
			}, nil
		}

		if resp.Detail == nil {
			return &CustomerLimitResult{
				Success:  true,
				Messages: []string{},
			}, nil
		}

		return &CustomerLimitResult{
			Success:  true,
			Detail:   resp.Detail,
			Messages: []string{resp.Status.MessageId},
		}, nil
	}

	return &CustomerLimitResult{
		Success:  false,
		Messages: []string{"Invalid response structure"},
	}, nil
}
