package cardreplace

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Params struct {
	Username          string
	Password          string
	AccountNumber     string
	BranchCode        string
	PhoneNumber       string
	CardType          string
	ReplacementReason string
	ProductType       string
}

type CardReplaceParam struct {
	AccountNumber     string
	BranchCode        string
	PhoneNumber       string
	CardType          string
	ReplacementReason string
	ProductType       string
}

func NewCardReplace(param Params) string {
	return fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cbes="http://temenos.com/CBESUPERAPP" xmlns:atm="http://temenos.com/ATMCARDREGDETCARDREPLACESUPERAPP">
		<soapenv:Header/>
		<soapenv:Body>
		<cbes:ATMCardReplacementRequest>
		<WebRequestCommon>
		<company/>
		<password>%s</password>
		<userName>%s</userName>
		</WebRequestCommon>
		<OfsFunction/>
		<ATMCARDREGDETCARDREPLACESUPERAPPType id="">
		<atm:ACCOUNT>%s</atm:ACCOUNT>
		<atm:BRANCHCODE>%s</atm:BRANCHCODE>
		<atm:PHONENO>%s</atm:PHONENO>
		<atm:CARDTYPE>%s</atm:CARDTYPE>
		<atm:ReplacementReason>%s</atm:ReplacementReason>
		<atm:ProductType>%s</atm:ProductType>
		</ATMCARDREGDETCARDREPLACESUPERAPPType>
		</cbes:ATMCardReplacementRequest>
		</soapenv:Body>
</soapenv:Envelope>`, param.Password, param.Username, param.AccountNumber, param.BranchCode, param.PhoneNumber, param.CardType, param.ReplacementReason, param.ProductType)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	CardReplaceResponse *CardReplaceResponse `xml:"ATMCardReplacementRequestResponse"`
}

type CardReplaceResponse struct {
	Status *struct {
		SuccessIndicator string `xml:"successIndicator"`
		TransactionId    string `xml:"transactionId"`
		Application      string `xml:"application"`
		MessageId        string `xml:"messageId"`
	} `xml:"Status"`
	ATMCardReplacementRequestDetail *ATMCardReplacementRequestDetail `xml:"ATMCARDREGDETType"`
}

type ATMCardReplacementRequestDetail struct {
	AccountNumber     string `xml:"ACCOUNT"`
	AccountHolderName string `xml:"ACCOUNTTITLE1"`
	Address           string `xml:"ADDRESS"`
	BranchName        string `xml:"BRANCHNAME"`
	OpenDate          string `xml:"ACOPENDATE"`
	Residence         string `xml:"RESIDENCE"`
	Industry          string `xml:"INDUSTRY"`
	BranchCode        string `xml:"BRANCHCODE"`
	PhoneNumber       string `xml:"PHONENO"`
	Sex               string `xml:"SEX"`
	CivilStatus       string `xml:"CIVILSTATUS"`
	HolderNumber      string `xml:"HOLDERNO"`
	CurruntNumber     string `xml:"CURRNO"`
	CardType          string `xml:"CARDTYPE"`
	Inputter          string `xml:"INPUTTER"`
	Datetime          string `xml:"DATETIME"`
	GDatetime         struct {
		DateTime string `xml:"DATETIME"`
	} `xml:"gDATETIME"`
	Authoriser        string `xml:"AUTHORISER"`
	CoCode            string `xml:"COCODE"`
	DeptCode          string `xml:"DEPTCODE"`
	Date              string `xml:"DATE"`
	RequestType       string `xml:"REQUESTTYPE"`
	ReplacementReason string `xml:"REPLREASON"`
	ProductType       string `xml:"PRODUCTTYPE"`
	VirtualFlag       string `xml:"VIRTUALFLAG"`
}

type CardReplaceResult struct {
	Success  bool
	Detail   *ATMCardReplacementRequestDetail
	Messages []string
}

func ParseCardReplaceResponse(xmlData string) (*CardReplaceResult, error) {
	var env Envelope
	err := xml.Unmarshal([]byte(xmlData), &env)
	if err != nil {
		return nil, err
	}

	if env.Body.CardReplaceResponse != nil {
		resp := env.Body.CardReplaceResponse
		if resp.Status == nil {
			return &CardReplaceResult{
				Success:  false,
				Messages: []string{"Missing Status"},
			}, nil
		}
		if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
			return &CardReplaceResult{
				Success:  false,
				Messages: []string{"API returned failure"},
			}, nil
		}
		if resp.ATMCardReplacementRequestDetail == nil {
			return &CardReplaceResult{
				Success:  false,
				Messages: []string{"Missing ATMCardReplacementRequestDetail"},
			}, nil
		}
		return &CardReplaceResult{
			Success: true,
			Detail:  resp.ATMCardReplacementRequestDetail,
		}, nil
	}
	return &CardReplaceResult{
		Success:  false,
		Messages: []string{"Invalid response"},
	}, nil
}
