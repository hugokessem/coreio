package unsbscribe

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Params struct {
	Username string
	Password string
	UserCode string
}

type UnsubscribeParam struct {
	UserCode string
}

func NewUnsubscribe(param Params) string {
	return fmt.Sprintf(`
		<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"
	xmlns:cbes="http://temenos.com/CBESUPERAPP">
		<soapenv:Header/>
		<soapenv:Body>
			<cbes:DeleteSuperAppUser>
				<WebRequestCommon>
					<company/>
					<password>%s</password>
					<userName>%s</userName>
				</WebRequestCommon>
				<OfsFunction/>
				<SUPERAPPUSERUNSUBSCRIBEType>
					<transactionId>%s</transactionId>
				</SUPERAPPUSERUNSUBSCRIBEType>
			</cbes:DeleteSuperAppUser>
		</soapenv:Body>
</soapenv:Envelope>`, param.Password, param.Username, param.UserCode)
}

// response
/*
<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns56:DeleteSuperAppUserResponse xmlns:ns2="http://temenos.com/CIFINFOSUPERAPP" xmlns:ns3="http://temenos.com/CUSTOMERVERIFYENQSUPERAPP" xmlns:ns4="http://temenos.com/CUSTOMERLIMITCUSTOMERSERVICEAMEND" xmlns:ns5="http://temenos.com/CUSTOMERLIMIT" xmlns:ns6="http://temenos.com/BRANCHLISTSUPERAPP" xmlns:ns7="http://temenos.com/GETPHONECUSTOMER" xmlns:ns8="http://temenos.com/CUSTOMERLIMITVIEW" xmlns:ns9="http://temenos.com/STANDINGORDERMANAGEORDERSUPERAPP" xmlns:ns10="http://temenos.com/STANDINGORDER" xmlns:ns11="http://temenos.com/PREMASTERCARDREGDETMCARDREGSUPERAPP" xmlns:ns12="http://temenos.com/PREMASTERCARDREGDET" xmlns:ns13="http://temenos.com/ATMCARDREGDETCARDREPLACESUPERAPP" xmlns:ns14="http://temenos.com/ATMCARDREGDET" xmlns:ns15="http://temenos.com/ACCTLOCKEDAMOUNTSSUPERAPP" xmlns:ns16="http://temenos.com/CUSTOMERLOOKUPSUPERAPP" xmlns:ns17="http://temenos.com/ACCTSTOLISTSUPERAPP" xmlns:ns18="http://temenos.com/FUNDSTRANSFERYABXDISBURSESUPERAPP" xmlns:ns19="http://temenos.com/FUNDSTRANSFER" xmlns:ns20="http://temenos.com/ACLOCKEDEVENTSRELEASELOCKSUPERAPP" xmlns:ns21="http://temenos.com/ACLOCKEDEVENTS" xmlns:ns22="http://temenos.com/ACCOUNTINFOSUPERAPPRESTRICT" xmlns:ns23="http://temenos.com/ACCTSTOLISTHISSUPERAPP" xmlns:ns24="http://temenos.com/SERVICELIMITCREATE" xmlns:ns25="http://temenos.com/SERVICELIMIT" xmlns:ns26="http://temenos.com/FUNDSTRANSFERVIEWDETAILSSUPERAPP" xmlns:ns27="http://temenos.com/FUNDSTRANSFERFTREVERSESUPERAPP" xmlns:ns28="http://temenos.com/ATMCARDREGDETCARDREQSUPERAPP" xmlns:ns29="http://temenos.com/GLOBALLIMITVIEWSUPERAPP" xmlns:ns30="http://temenos.com/EXCHANGERATESUPERAPP" xmlns:ns31="http://temenos.com/STANDINGORDERTXNLISTSUPERAPP" xmlns:ns32="http://temenos.com/TXNSTATUSSUPERAPP" xmlns:ns33="http://temenos.com/CARDTYPELISTSUPERAPP" xmlns:ns34="http://temenos.com/ACCTSTMTRGSUPERAPP" xmlns:ns35="http://temenos.com/ACLOCKEDEVENTSCREATELOCKSUPERAPP" xmlns:ns36="http://temenos.com/CUSTOMLIMITAMENDENQ" xmlns:ns37="http://temenos.com/CARDSTATUSSUPERAPP" xmlns:ns38="http://temenos.com/ACCOUNTENQUIRYSUPERAPP" xmlns:ns39="http://temenos.com/CBEMINISTMTENQ" xmlns:ns40="http://temenos.com/SUPERAPPUSERUNSUBSCRIBE" xmlns:ns41="http://temenos.com/SUPERAPPUSER" xmlns:ns42="http://temenos.com/FTBULKCREDITACSPLITPAYMENTACSUPERAPP" xmlns:ns43="http://temenos.com/FTBULKCREDITAC" xmlns:ns44="http://temenos.com/FUNDSTRANSFERFTTXNSUPERAPP" xmlns:ns45="http://temenos.com/CUSTOMERINFOSUPERAPP" xmlns:ns46="http://temenos.com/SUPERAPPUSERSTATUS" xmlns:ns47="http://temenos.com/NAMELOOKUPSUPERAPP" xmlns:ns48="http://temenos.com/ACCOUNTINFOSUPERAPP" xmlns:ns49="http://temenos.com/FUNDSTRANSFERBILLPAYSUPERAPP" xmlns:ns50="http://temenos.com/SUPERAPPUSERCREATE" xmlns:ns51="http://temenos.com/SERVICELISTVIEWSUPERAPP" xmlns:ns52="http://temenos.com/MMMONEYMARKETSUPERAPP" xmlns:ns53="http://temenos.com/MMMONEYMARKET" xmlns:ns54="http://temenos.com/FUNDSTRANSFERYABXREPAYMENTSUPERAPP" xmlns:ns55="http://temenos.com/CUSTOMERLIMITCUSTOMLIMIT" xmlns:ns56="http://temenos.com/CBESUPERAPP">
            <Status>
                <transactionId>SA1002027256</transactionId>
                <messageId></messageId>
                <successIndicator>Success</successIndicator>
                <application>SUPERAPP.USER</application>
            </Status>
            <SUPERAPPUSERType id="SA1002027256">
                <ns41:CUSTOMERNAME>Mr Yohannes Teshome</ns41:CUSTOMERNAME>
                <ns41:CUSTOMERID>1002027256</ns41:CUSTOMERID>
                <ns41:STATUS>ACTIVE</ns41:STATUS>
                <ns41:PHONENO>+251911706628</ns41:PHONENO>
                <ns41:EMAIL>yohannes@yml.com</ns41:EMAIL>
                <ns41:CHANNEL>ALL</ns41:CHANNEL>
                <ns41:DESCRIPTION>Register Again</ns41:DESCRIPTION>
                <ns41:LASTUPDATEDDATE>20211209</ns41:LASTUPDATEDDATE>
                <ns41:BRANCHCODE>6043</ns41:BRANCHCODE>
                <ns41:SALESID>CBE059562</ns41:SALESID>
                <ns41:RECORDSTATUS>REVE</ns41:RECORDSTATUS>
                <ns41:CURRNO>2</ns41:CURRNO>
                <ns41:gINPUTTER>
                    <ns41:INPUTTER>84365_SUPERAPP.1__OFS_GCS</ns41:INPUTTER>
                </ns41:gINPUTTER>
                <ns41:gDATETIME>
                    <ns41:DATETIME>2606081115</ns41:DATETIME>
                </ns41:gDATETIME>
                <ns41:AUTHORISER>84365_SUPERAPP.1_OFS_GCS</ns41:AUTHORISER>
                <ns41:COCODE>ET0010001</ns41:COCODE>
                <ns41:DEPTCODE>1</ns41:DEPTCODE>
            </SUPERAPPUSERType>
        </ns56:DeleteSuperAppUserResponse>
    </S:Body>
</S:Envelope>
*/

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	DeleteSuperAppUserResponse *DeleteSuperAppUserResponse `xml:"DeleteSuperAppUserResponse"`
}

type DeleteSuperAppUserResponse struct {
	Status *struct {
		SuccessIndicator string   `xml:"successIndicator"`
		Messages         []string `xml:"messages"`
		Application      string   `xml:"application"`
		TransactionId    string   `xml:"transactionId"`
	} `xml:"Status"`
	SuperAppUserType *SuperAppUserType `xml:"SUPERAPPUSERType"`
}

type SuperAppUserType struct {
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
	} `xml:"DATETIME"`
	Authoriser string `xml:"AUTHORISER"`
	Code       string `xml:"COCODE"`
	DeptCode   string `xml:"DEPTCODE"`
}

type UnsubscribeResult struct {
	Success  bool
	Detail   *SuperAppUserType
	Messages []string
}

func ParseUnsubscribeResponseSOAP(xmlData string) (*UnsubscribeResult, error) {
	var env Envelope
	err := xml.Unmarshal([]byte(xmlData), &env)
	if err != nil {
		return nil, err
	}
	if env.Body.DeleteSuperAppUserResponse != nil {
		resp := env.Body.DeleteSuperAppUserResponse
		if resp.Status == nil {
			return &UnsubscribeResult{
				Success:  false,
				Messages: []string{"Missing Status"},
			}, nil
		}
		if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
			return &UnsubscribeResult{
				Success:  false,
				Messages: resp.Status.Messages,
			}, nil
		}
		if resp.SuperAppUserType == nil {
			return &UnsubscribeResult{
				Success:  false,
				Messages: []string{"Missing SuperAppUserType"},
			}, nil
		}
		return &UnsubscribeResult{
			Success:  true,
			Detail:   resp.SuperAppUserType,
			Messages: resp.Status.Messages,
		}, nil
	}
	return &UnsubscribeResult{
		Success:  false,
		Messages: []string{"Invalid response type"},
	}, nil
}
