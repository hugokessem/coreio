package splitpayment

import "fmt"

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
