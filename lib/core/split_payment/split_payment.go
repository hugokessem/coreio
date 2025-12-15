package splitpayment

import "fmt"

type Param struct {
	Username        string
	Password        string
	DebitAccount    string
	DebitCurrency   string
	DebitReference  string
	CreditCurrency  string
	CreditAccount   string
	CreditAmount    string
	CreditReference string
}

func NewSplitPayment(param Param) string {
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
                <ftb:gCRACCOUNT g="1">
                    <ftb:mCRACCOUNT m="1">
                        <ftb:CRACCOUNT>%s</ftb:CRACCOUNT>
                        <ftb:CRAMOUNT>%s</ftb:CRAMOUNT>
                        <ftb:CRTHIERREF>%s</ftb:CRTHIERREF>
                    </ftb:mCRACCOUNT>
                    <ftb:mCRACCOUNT m="2">
                        <ftb:CRACCOUNT>1000000001148</ftb:CRACCOUNT>
                        <ftb:CRAMOUNT>60</ftb:CRAMOUNT>
                        <ftb:CRTHIERREF>BEN 22</ftb:CRTHIERREF>
                    </ftb:mCRACCOUNT>
                </ftb:gCRACCOUNT>
            </FTBULKCREDITACSPLITPAYMENTACSUPERAPPType>
        </cbes:SplitPaymentSuperApp>
    </soapenv:Body>
</soapenv:Envelope>
	`)
}
