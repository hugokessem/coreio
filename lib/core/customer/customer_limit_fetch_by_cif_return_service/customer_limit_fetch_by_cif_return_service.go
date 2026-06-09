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

// response
/*
<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns56:ViewCustomedLimitResponse xmlns:ns2="http://temenos.com/CIFINFOSUPERAPP" xmlns:ns3="http://temenos.com/CUSTOMERVERIFYENQSUPERAPP" xmlns:ns4="http://temenos.com/CUSTOMERLIMITCUSTOMERSERVICEAMEND" xmlns:ns5="http://temenos.com/CUSTOMERLIMIT" xmlns:ns6="http://temenos.com/BRANCHLISTSUPERAPP" xmlns:ns7="http://temenos.com/GETPHONECUSTOMER" xmlns:ns8="http://temenos.com/CUSTOMERLIMITVIEW" xmlns:ns9="http://temenos.com/STANDINGORDERMANAGEORDERSUPERAPP" xmlns:ns10="http://temenos.com/STANDINGORDER" xmlns:ns11="http://temenos.com/PREMASTERCARDREGDETMCARDREGSUPERAPP" xmlns:ns12="http://temenos.com/PREMASTERCARDREGDET" xmlns:ns13="http://temenos.com/ATMCARDREGDETCARDREPLACESUPERAPP" xmlns:ns14="http://temenos.com/ATMCARDREGDET" xmlns:ns15="http://temenos.com/ACCTLOCKEDAMOUNTSSUPERAPP" xmlns:ns16="http://temenos.com/CUSTOMERLOOKUPSUPERAPP" xmlns:ns17="http://temenos.com/ACCTSTOLISTSUPERAPP" xmlns:ns18="http://temenos.com/FUNDSTRANSFERYABXDISBURSESUPERAPP" xmlns:ns19="http://temenos.com/FUNDSTRANSFER" xmlns:ns20="http://temenos.com/ACLOCKEDEVENTSRELEASELOCKSUPERAPP" xmlns:ns21="http://temenos.com/ACLOCKEDEVENTS" xmlns:ns22="http://temenos.com/ACCOUNTINFOSUPERAPPRESTRICT" xmlns:ns23="http://temenos.com/ACCTSTOLISTHISSUPERAPP" xmlns:ns24="http://temenos.com/SERVICELIMITCREATE" xmlns:ns25="http://temenos.com/SERVICELIMIT" xmlns:ns26="http://temenos.com/FUNDSTRANSFERVIEWDETAILSSUPERAPP" xmlns:ns27="http://temenos.com/FUNDSTRANSFERFTREVERSESUPERAPP" xmlns:ns28="http://temenos.com/ATMCARDREGDETCARDREQSUPERAPP" xmlns:ns29="http://temenos.com/GLOBALLIMITVIEWSUPERAPP" xmlns:ns30="http://temenos.com/EXCHANGERATESUPERAPP" xmlns:ns31="http://temenos.com/STANDINGORDERTXNLISTSUPERAPP" xmlns:ns32="http://temenos.com/TXNSTATUSSUPERAPP" xmlns:ns33="http://temenos.com/CARDTYPELISTSUPERAPP" xmlns:ns34="http://temenos.com/ACCTSTMTRGSUPERAPP" xmlns:ns35="http://temenos.com/ACLOCKEDEVENTSCREATELOCKSUPERAPP" xmlns:ns36="http://temenos.com/CUSTOMLIMITAMENDENQ" xmlns:ns37="http://temenos.com/CARDSTATUSSUPERAPP" xmlns:ns38="http://temenos.com/ACCOUNTENQUIRYSUPERAPP" xmlns:ns39="http://temenos.com/CBEMINISTMTENQ" xmlns:ns40="http://temenos.com/SUPERAPPUSERUNSUBSCRIBE" xmlns:ns41="http://temenos.com/SUPERAPPUSER" xmlns:ns42="http://temenos.com/FTBULKCREDITACSPLITPAYMENTACSUPERAPP" xmlns:ns43="http://temenos.com/FTBULKCREDITAC" xmlns:ns44="http://temenos.com/FUNDSTRANSFERFTTXNSUPERAPP" xmlns:ns45="http://temenos.com/CUSTOMERINFOSUPERAPP" xmlns:ns46="http://temenos.com/SUPERAPPUSERSTATUS" xmlns:ns47="http://temenos.com/NAMELOOKUPSUPERAPP" xmlns:ns48="http://temenos.com/ACCOUNTINFOSUPERAPP" xmlns:ns49="http://temenos.com/FUNDSTRANSFERBILLPAYSUPERAPP" xmlns:ns50="http://temenos.com/SUPERAPPUSERCREATE" xmlns:ns51="http://temenos.com/SERVICELISTVIEWSUPERAPP" xmlns:ns52="http://temenos.com/MMMONEYMARKETSUPERAPP" xmlns:ns53="http://temenos.com/MMMONEYMARKET" xmlns:ns54="http://temenos.com/FUNDSTRANSFERYABXREPAYMENTSUPERAPP" xmlns:ns55="http://temenos.com/CUSTOMERLIMITCUSTOMLIMIT" xmlns:ns56="http://temenos.com/CBESUPERAPP">
            <Status>
                <successIndicator>Success</successIndicator>
            </Status>
            <CUSTOMLIMITAMENDENQType>
                <ns36:gCUSTOMLIMITAMENDENQDetailType>
                    <ns36:mCUSTOMLIMITAMENDENQDetailType>
                        <ns36:CIF>1036095547</ns36:CIF>
                        <ns36:Channel>APP</ns36:Channel>
                        <ns36:ServiceCode>MPESA</ns36:ServiceCode>
                        <ns36:ServiceName></ns36:ServiceName>
                        <ns36:Limit>10000000</ns36:Limit>
                        <ns36:Count>5000</ns36:Count>
                    </ns36:mCUSTOMLIMITAMENDENQDetailType>
                    <ns36:mCUSTOMLIMITAMENDENQDetailType>
                        <ns36:CIF></ns36:CIF>
                        <ns36:Channel></ns36:Channel>
                        <ns36:ServiceCode>TELEBIRR</ns36:ServiceCode>
                        <ns36:ServiceName>Transfer To Telebirr</ns36:ServiceName>
                        <ns36:Limit>100000</ns36:Limit>
                        <ns36:Count>5</ns36:Count>
                    </ns36:mCUSTOMLIMITAMENDENQDetailType>
                    <ns36:mCUSTOMLIMITAMENDENQDetailType>
                        <ns36:CIF></ns36:CIF>
                        <ns36:Channel></ns36:Channel>
                        <ns36:ServiceCode>CBE</ns36:ServiceCode>
                        <ns36:ServiceName>Transfer To CBE</ns36:ServiceName>
                        <ns36:Limit>1000000000</ns36:Limit>
                        <ns36:Count>10000000</ns36:Count>
                    </ns36:mCUSTOMLIMITAMENDENQDetailType>
                    <ns36:mCUSTOMLIMITAMENDENQDetailType>
                        <ns36:CIF></ns36:CIF>
                        <ns36:Channel>USSD</ns36:Channel>
                        <ns36:ServiceCode>AAPARKING</ns36:ServiceCode>
                        <ns36:ServiceName>AATMA Parking Payment</ns36:ServiceName>
                        <ns36:Limit>100000</ns36:Limit>
                        <ns36:Count>5</ns36:Count>
                    </ns36:mCUSTOMLIMITAMENDENQDetailType>
                </ns36:gCUSTOMLIMITAMENDENQDetailType>
            </CUSTOMLIMITAMENDENQType>
        </ns56:ViewCustomedLimitResponse>
    </S:Body>
</S:Envelope>
*/

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
