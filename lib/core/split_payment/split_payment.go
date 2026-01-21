package splitpayment

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Param struct {
	Username                 string
	Password                 string
	DebitCurrency            string
	DebitAccount             string
	DebitReference           string
	CreditCurrency           string
	CreditAccountInformation []CreditAccountInformation
}

type CreditAccountInformation struct {
	CreditAccount   string
	CreditAmount    string
	CreditReference string
}

type SplitPaymentParam struct {
	DebitCurrency            string
	DebitAccount             string
	DebitReference           string
	CreditCurrency           string
	CreditAccountInformation []CreditAccountInformation
}

func NewSplitPayment(param Param) string {
	creditInformation := make([]string, 0, len(param.CreditAccountInformation))
	for index, value := range param.CreditAccountInformation {
		temp := fmt.Sprintf(`
			<ftb:mCRACCOUNT m="%d">
				<ftb:CRACCOUNT>%s</ftb:CRACCOUNT>
				<ftb:CRAMOUNT>%s</ftb:CRAMOUNT>
				<ftb:CRTHIERREF>%s</ftb:CRTHIERREF>
			</ftb:mCRACCOUNT>
		`, index, value.CreditAccount, value.CreditAmount, value.CreditReference)
		creditInformation = append(creditInformation, temp)
	}

	return fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"
xmlns:cbes="http://temenos.com/CBESUPERAPP"
xmlns:ftb="http://temenos.com/FTBULKCREDITACSPLITPAYMENTACSUPERAPP">
    <soapenv:Header/>
    <soapenv:Body>
        <cbes:SplitPaymentSuperApp>
            <WebRequestCommon>
                <company/>
                <password>%s</password>
                <userName>%s</userName>
            </WebRequestCommon>
            <OfsFunction></OfsFunction>
            <FTBULKCREDITACSPLITPAYMENTACSUPERAPPType id="">
                <ftb:DRACCOUNT>%s</ftb:DRACCOUNT>
                <ftb:DRCURRENCY>%s</ftb:DRCURRENCY>
                <ftb:DRAMOUNT>%s</ftb:DRAMOUNT>
                <ftb:DRTHIERREF>%s</ftb:DRTHIERREF>
                <ftb:CRCURRENCY>%s</ftb:CRCURRENCY>
                <ftb:gCRACCOUNT g="1">%s</ftb:gCRACCOUNT>
            </FTBULKCREDITACSPLITPAYMENTACSUPERAPPType>
        </cbes:SplitPaymentSuperApp>
    </soapenv:Body>
</soapenv:Envelope>
	`, param.Password, param.Username, param.DebitAccount, param.DebitCurrency, param.DebitAccount, param.DebitReference, param.CreditCurrency, creditInformation)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	SplitPaymentResponse *SplitPaymentResponse `xml:"SplitPaymentSuperAppResponse"`
}

type SplitPaymentResponse struct {
	Status             *SplitPaymentStatus `xml:"Status"`
	SplitPaymentDetail *SplitPaymentDetail `xml:"FTBULKCREDITACType"`
}

type SplitPaymentStatus struct {
	TransactionID    string   `xml:"transactionId"`
	MessageID        string   `xml:"messageId"`
	SuccessIndicator string   `xml:"successIndicator"`
	Application      string   `xml:"application"`
	Messages         []string `xml:"messages"`
}

type CreditAccountDetail struct {
	Account        string `xml:"CRACCOUNT"`
	Amount         string `xml:"CRAMOUNT"`
	Reference      string `xml:"CRTHIERREF"`
	OfsGeneratedID string `xml:"OFSGENID"`
	OfsErrorYN     string `xml:"OFSERRYN"`
}

type SplitPaymentDetail struct {
	XMLName         xml.Name `xml:"FTBULKCREDITACType"`
	ID              string   `xml:"id,attr"`
	TransactionType string   `xml:"TRANSACTIONTYPE"`
	DebitAccount    string   `xml:"DRACCOUNT"`
	DebitCurrency   string   `xml:"DRCURRENCY"`
	DebitAmount     string   `xml:"DRAMOUNT"`
	DebitReference  string   `xml:"DRTHIERREF"`
	CreditCurrency  string   `xml:"CRCURRENCY"`
	CreditAccounts  struct {
		Items []CreditAccountDetail `xml:"mCRACCOUNT"`
	} `xml:"gCRACCOUNT"`
	ProcessingDate string `xml:"PROCESSINGDATE"`
	OrderingBanks  struct {
		Banks []string `xml:"ORDERINGBK"`
	} `xml:"gORDERINGBK"`
	ProfitCentreCust string `xml:"PROFITCENTRECUST"`
	Overrides        struct {
		Items []string `xml:"OVERRIDE"`
	} `xml:"gOVERRIDE"`
	CurrentNumber string `xml:"CURRNO"`
	Inputter      struct {
		Value string `xml:"INPUTTER"`
	} `xml:"gINPUTTER"`
	DateTime struct {
		Value string `xml:"DATETIME"`
	} `xml:"gDATETIME"`
	Authoriser     string `xml:"AUTHORISER"`
	CompanyCode    string `xml:"COCODE"`
	DepartmentCode string `xml:"DEPTCODE"`
}

type SplitPaymentResult struct {
	Success  bool
	Detail   *SplitPaymentDetail
	Messages []string
}

func ParseSplitPaymentSOAP(xmlData string) (*SplitPaymentResult, error) {
	var env Envelope
	if err := xml.Unmarshal([]byte(xmlData), &env); err != nil {
		return nil, err
	}

	if env.Body.SplitPaymentResponse != nil {
		resp := env.Body.SplitPaymentResponse
		if resp.Status == nil {
			return &SplitPaymentResult{
				Success:  false,
				Messages: []string{"Missing Status"},
			}, nil
		}

		if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
			return &SplitPaymentResult{
				Success:  false,
				Messages: resp.Status.Messages,
			}, nil
		}

		return &SplitPaymentResult{
			Success:  true,
			Detail:   resp.SplitPaymentDetail,
			Messages: resp.Status.Messages,
		}, nil
	}

	return &SplitPaymentResult{
		Success:  false,
		Messages: []string{"Invalid response type"},
	}, nil
}
