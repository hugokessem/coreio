// Package fundtransfer provides SOAP XML generators and parsers
// for core banking fund transfer requests.
package fundtransfer

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Params struct {
	Username            string
	Password            string
	DebitAccountNumber  string
	DebitCurrency       string
	CreditAccountNumber string
	CreditAmount        string
	CreditCurrency      string
	DebitReference      string
	CreditReference     string
	DebitAmount         string
	TransactionID       string
	PaymentDetail       string
	ServiceCode         string
}

type FundTransferParam struct {
	DebitAccountNumber  string
	DebitCurrency       string
	CreditAccountNumber string
	CreditAmount        string
	CreditCurrency      string
	DebitReference      string
	CreditReference     string
	DebitAmount         string
	TransactionID       string
	PaymentDetail       string
	ServiceCode         string
}

func NewFundTransfer(param Params) string {
	return fmt.Sprintf(
		`<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cbes="http://temenos.com/CBESUPERAPP" xmlns:fun="http://temenos.com/FUNDSTRANSFERFTTXNSUPERAPP">
		<soapenv:Header/>
		<soapenv:Body>
			<cbes:AccountTransfer>
				<WebRequestCommon>
					<company></company>
					<password>%s</password>
					<userName>%s</userName>
				</WebRequestCommon>
				<OfsFunction></OfsFunction>
				<FUNDSTRANSFERFTTXNSUPERAPPType id="">
					<fun:DEBITACCTNO>%s</fun:DEBITACCTNO>
					<fun:DEBITCURRENCY>%s</fun:DEBITCURRENCY>
					<fun:DEBITAMOUNT>%s</fun:DEBITAMOUNT>
					<fun:DEBITTHEIRREF>%s</fun:DEBITTHEIRREF>
					<fun:CREDITTHEIRREF>%s</fun:CREDITTHEIRREF>
					<fun:CREDITACCTNO>%s</fun:CREDITACCTNO>
					<fun:CREDITCURRENCY>%s</fun:CREDITCURRENCY>
					<fun:CREDITAMOUNT>%s</fun:CREDITAMOUNT>
					<fun:gPAYMENTDETAILS g="1">
						<fun:PAYMENTDETAILS>%s</fun:PAYMENTDETAILS>
					</fun:gPAYMENTDETAILS>
					<fun:ClientReference>%s</fun:ClientReference>
					<fun:ServiceCode>%s</fun:ServiceCode>
				</FUNDSTRANSFERFTTXNSUPERAPPType>
			</cbes:AccountTransfer>
		</soapenv:Body>
		</soapenv:Envelope>
`, param.Password, param.Username, param.DebitAccountNumber, param.DebitCurrency, param.DebitAmount, param.DebitReference, param.CreditReference, param.CreditAccountNumber, param.CreditCurrency, param.CreditAmount, param.PaymentDetail, param.TransactionID, param.ServiceCode)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	FundTransferResponse *FundTransferResponse `xml:"AccountTransferResponse"`
}

type FundTransferDetail struct {
	XMLName              xml.Name `xml:"FUNDSTRANSFERType"`
	FTNumber             string   `xml:"id,attr"`
	TransactionType      string   `xml:"TRANSACTIONTYPE"`
	DebitAccountNumber   string   `xml:"DEBITACCTNO"`
	CurrencyMarketDebit  string   `xml:"CURRENCYMKTDR"`
	DebitCurrency        string   `xml:"DEBITCURRENCY"`
	DebitAmount          string   `xml:"DEBITAMOUNT"`
	DebitValueDate       string   `xml:"DEBITVALUEDATE"`
	DebitReference       string   `xml:"DEBITTHEIRREF"`
	CreditReference      string   `xml:"CREDITTHEIRREF"`
	CreditAccountNumber  string   `xml:"CREDITACCTNO"`
	CurrencyMarketCredit string   `xml:"CURRENCYMKTCR"`
	CreditCurrency       string   `xml:"CREDITCURRENCY"`
	CreditValidationDare string   `xml:"CREDITVALUEDATE"`
	ProcessingDate       string   `xml:"PROCESSINGDATE"`
	PaymentDetails       struct {
		PaymentDetail string `xml:"PAYMENTDETAILS"`
	} `xml:"gPAYMENTDETAILS"`
	ChargeCommisionDisplay string `xml:"CHARGECOMDISPLAY"`
	CommissionCode         string `xml:"COMMISSIONCODE"`
	GlobalCommissionType   struct {
		MultipleCommissionType []struct {
			CommissionType   string `xml:"COMMISSIONTYPE"`
			CommissionAmount string `xml:"COMMISSIONAMT"`
		} `xml:"mCOMMISSIONTYPE"`
	} `xml:"gCOMMISSIONTYPE"`
	ChargeCode           string `xml:"CHARGECODE"`
	ProfitCentreCustomer string `xml:"PROFITCENTRECUST"`
	ReturnToDept         string `xml:"RETURNTODEPT"`
	FedFunds             string `xml:"FEDFUNDS"`
	PositionType         string `xml:"POSITIONTYPE"`
	GlobalTaxType        struct {
		MultipleTaxType []struct {
			TaxType   string `xml:"TAXTYPE"`
			TaxAmount string `xml:"TAXAMT"`
		} `xml:"mTAXTYPE"`
	} `xml:"gTAXTYPE"`
	DebitAmountWithCurrency  string `xml:"AMOUNTDEBITED"`
	CreditAmountWithCurrency string `xml:"AMOUNTCREDITED"`
	TotalChargeAmount        string `xml:"TOTALCHARGEAMT"`
	TotalTaxAmount           string `xml:"TOTALTAXAMOUNT"`
	DeliveryOutRef           struct {
		MultipleDeliveryOutRef []string `xml:"DELIVERYOUTREF"`
	} `xml:"gDELIVERYOUTREF"`
	CreditCompanyCode            string `xml:"CREDITCOMPCODE"`
	DebitCompanyCode             string `xml:"DEBITCOMPCODE"`
	LocalAmountDebited           string `xml:"LOCAMTDEBITED"`
	LocalAmountCredited          string `xml:"LOCAMTCREDITED"`
	LocalTotalTaxAmount          string `xml:"LOCTOTTAXAMT"`
	LocalChargeAmount            string `xml:"LOCALCHARGEAMT"`
	LocalPositionChargesAmount   string `xml:"LOCPOSCHGSAMT"`
	CustomerGroupLevel           string `xml:"CUSTGROUPLEVEL"`
	DebitCustomer                string `xml:"DEBITCUSTOMER"`
	CreditCustomer               string `xml:"CREDITCUSTOMER"`
	DebitAdviceRequired          string `xml:"DRADVICEREQDYN"`
	CreditAdviceRequired         string `xml:"CRADVICEREQDYN"`
	ChargedCustomer              string `xml:"CHARGEDCUSTOMER"`
	TotalReceivedCommission      string `xml:"TOTRECCOMM"`
	TotalReceivedCommissionLocal string `xml:"TOTRECCOMMLCL"`
	TotalReceivedCharge          string `xml:"TOTRECCHG"`
	TotalReceivedChargeLocal     string `xml:"TOTRECCHGLCL"`
	RateFixing                   string `xml:"RATEFIXING"`
	TotalReceivedChargeCurrency  string `xml:"TOTRECCHGCRCCY"`
	TotalSentChargeCurrency      string `xml:"TOTSNDCHGCRCCY"`
	AuthDate                     string `xml:"AUTHDATE"`
	RoundType                    string `xml:"ROUNDTYPE"`
	GlobalStatementNumbers       struct {
		MultipleStatementNumbers []string `xml:"STMTNOS"`
	} `xml:"gSTMTNOS"`
	GlobalOverride struct {
		Override []string `xml:"OVERRIDE"`
	} `xml:"gOVERRIDE"`
	CurrentNumber  string `xml:"CURRNO"`
	GlobalInputter struct {
		Inputter string `xml:"INPUTTER"`
	} `xml:"gINPUTTER"`
	GlobalDateTime struct {
		DateTime string `xml:"DATETIME"`
	} `xml:"gDATETIME"`
	Authoriser                         string `xml:"AUTHORISER"`
	CompanyCode                        string `xml:"COCODE"`
	DepartmentCode                     string `xml:"DEPTCODE"`
	InputVersion                       string `xml:"LINPUTVERSION"`
	AuthVersion                        string `xml:"LAUTHVERSION"`
	TransactionID                      string `xml:"MTOREF"`
	DebitAccountHolderName             string `xml:"SENDERNAME"`
	ReceiverName                       string `xml:"RECEIVERNAME"`
	ServiceCode                        string `xml:"SERVICECODE"`
	DebitAccountCurrentWorkingBalance  string `xml:"CEKCS"`
	CreditAccountCurrentWorkingBalance string `xml:"GPONU"`
}

type FundTransferResponse struct {
	Status *struct {
		SuccessIndicator string   `xml:"successIndicator"`
		TransactionID    string   `xml:"transactionId"`
		Application      string   `xml:"application"`
		Messages         []string `xml:"messages"`
	} `xml:"Status"`
	FundTransferType *FundTransferDetail `xml:"FUNDSTRANSFERType"`
}

type FundTransferResult struct {
	Success  bool
	Detail   *FundTransferDetail
	Messages []string
}

func ParseFundTransferSOAP(xmlData string) (*FundTransferResult, error) {
	var env Envelope
	if err := xml.Unmarshal([]byte(xmlData), &env); err != nil {
		return nil, err
	}

	if env.Body.FundTransferResponse != nil {
		resp := env.Body.FundTransferResponse
		if resp.Status == nil {
			return &FundTransferResult{
				Success:  false,
				Messages: []string{"Missing Status"},
			}, nil
		}

		if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
			return &FundTransferResult{
				Success:  false,
				Messages: resp.Status.Messages,
			}, nil
		}

		if resp.FundTransferType == nil {
			return &FundTransferResult{
				Success:  true,
				Messages: []string{},
			}, nil
		}

		return &FundTransferResult{
			Success: true,
			Detail:  resp.FundTransferType,
		}, nil
	}

	return &FundTransferResult{
		Success:  false,
		Messages: []string{"Invalid response type"},
	}, nil
}
