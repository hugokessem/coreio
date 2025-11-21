package lockedamountft

import "fmt"

type Params struct {
	Username            string
	Password            string
	CreditCurrent       string
	CreditAccountNumber string
	CrediterReference   string
	DebitAmount         string
	DebitAccountNumber  string
	DebitCurrency       string
	DebiterReference    string
	ClientReference     string
	LockID              string
}

func NewLockedAmountFt(param Params) string {
	return fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"
xmlns:cbes="http://temenos.com/CBESUPERAPP"
xmlns:fun="http://temenos.com/FUNDSTRANSFERFTTXNSUPERAPP">
    <soapenv:Header/>
    <soapenv:Body>
        <cbes:AccountTransfer>
            <WebRequestCommon>
                <company/>
                <password>%s</password>
                <userName>%s</userName>
            </WebRequestCommon>
            <OfsFunction/>
            <FUNDSTRANSFERFTTXNSUPERAPPType id="">
                <fun:DEBITACCTNO>%s</fun:DEBITACCTNO>
                <fun:DEBITCURRENCY>%s</fun:DEBITCURRENCY>
                <fun:DEBITAMOUNT>%s</fun:DEBITAMOUNT>
                <fun:DEBITTHEIRREF>%s</fun:DEBITTHEIRREF>
                <fun:CREDITTHEIRREF>%s</fun:CREDITTHEIRREF>
                <fun:CREDITACCTNO>%s</fun:CREDITACCTNO>
                <fun:CREDITCURRENCY>%s</fun:CREDITCURRENCY>
                <fun:CREDITAMOUNT/>
                <fun:gPAYMENTDETAILS g="1">
                    <fun:PAYMENTDETAILS/>
                </fun:gPAYMENTDETAILS>
                <fun:COMMISSIONCODE/>
                <fun:CHARGECODE/>
                <fun:gCHARGETYPE g="1">
                    <fun:CHARGETYPE/>
                </fun:gCHARGETYPE>
                <fun:ClientReference>%s</fun:ClientReference>
                <fun:LockID>%s</fun:LockID>
            </FUNDSTRANSFERFTTXNSUPERAPPType>
        </cbes:AccountTransfer>
    </soapenv:Body>
</soapenv:Envelope>
	`, param.Password, param.Username, param.DebitAccountNumber, param.DebitCurrency, param.DebitAmount, param.DebiterReference, param.CrediterReference, param.CreditAccountNumber, param.CreditCurrent, param.ClientReference, param.LockID)
}
