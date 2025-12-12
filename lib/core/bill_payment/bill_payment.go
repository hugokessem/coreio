package billpayment

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
	DebitAmount         string
	DebitReference      string
	CrediterReference   string
	CreditAccountNumber string
	CreditCurrency      string
	PaymentDetail       string
	ClientReference     string
}

type BillPaymentParams struct {
	DebitAccountNumber  string
	DebitCurrency       string
	DebitAmount         string
	DebitReference      string
	CrediterReference   string
	CreditAccountNumber string
	CreditCurrency      string
	PaymentDetail       string
	ClientReference     string
}

func NewBillPayment(param Params) string {
	return fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cbes="http://temenos.com/CBESUPERAPP" xmlns:fun="http://temenos.com/FUNDSTRANSFERBILLPAYSUPERAPP">
    <soapenv:Header/>
    <soapenv:Body>
        <cbes:FTBillPayment>
            <WebRequestCommon>
                <company/>
                <password>%s</password>
                <userName>%s</userName>
            </WebRequestCommon>
            <OfsFunction/>
            <FUNDSTRANSFERBILLPAYSUPERAPPType id="">
                <fun:DEBITACCTNO>%s</fun:DEBITACCTNO>
                <fun:DEBITCURRENCY>%s</fun:DEBITCURRENCY>
                <fun:DEBITAMOUNT>%s</fun:DEBITAMOUNT>
                <fun:DEBITTHEIRREF>%s</fun:DEBITTHEIRREF>
                <fun:CREDITTHEIRREF>%s</fun:CREDITTHEIRREF>
                <fun:CREDITACCTNO>%s</fun:CREDITACCTNO>
                <fun:CREDITCURRENCY>%s</fun:CREDITCURRENCY>
                <fun:CREDITAMOUNT/>
                <fun:gPAYMENTDETAILS g="1">
                    <fun:PAYMENTDETAILS>%s</fun:PAYMENTDETAILS>
                </fun:gPAYMENTDETAILS>
                <fun:gCOMMISSIONTYPE g="1">
                    <fun:COMMISSIONTYPE/>
                </fun:gCOMMISSIONTYPE>
                <fun:CHARGECODE/>
                <fun:gCHARGETYPE g="1">
                    <fun:CHARGETYPE></fun:CHARGETYPE>
                </fun:gCHARGETYPE>
                <fun:ClientReference>%s</fun:ClientReference>
            </FUNDSTRANSFERBILLPAYSUPERAPPType>
        </cbes:FTBillPayment>
    </soapenv:Body>
</soapenv:Envelope>`, param.Password, param.Username, param.DebitAccountNumber, param.DebitCurrency, param.DebitAmount, param.DebitReference, param.CrediterReference, param.CreditAccountNumber, param.CreditCurrency, param.PaymentDetail, param.ClientReference)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	FTBillPaymentResponse *FTBillPaymentResponse `xml:"FTBillPaymentResponse"`
}

type FTBillPaymentResponse struct {
	Status *struct {
		TransactionId    string `xml:"transactionId"`
		MessageId        string `xml:"messageId"`
		SuccessIndicator string `xml:"successIndicator"`
		Application      string `xml:"application"`
	} `xml:"Status"`
	BillPaymentDetail *BillPaymentDetail `xml:"FUNDSTRANSFERType"`
}

type BillPaymentDetail struct {
	TransactionId       string `xml:"id,attr"`
	TransactionType     string `xml:"TRANSACTIONTYPE"`
	DebitAccountNumber  string `xml:"DEBITACCTNO"`
	DebitCurrency       string `xml:"DEBITCURRENCY"`
	DebitAmount         string `xml:"DEBITAMOUNT"`
	DebitReference      string `xml:"DEBITTHEIRREF"`
	CrediterReference   string `xml:"CREDITTHEIRREF"`
	CreditAccountNumber string `xml:"CREDITACCTNO"`
	CreditCurrency      string `xml:"CREDITCURRENCY"`
	GlobalPaymentDetail struct {
		PaymentDetail string `xml:"PAYMENTDETAILS"`
	} `xml:"gPAYMENTDETAILS"`
	ChargeComDisplay    string `xml:"CHARGECOMDISPLAY"`
	CommissionCode      string `xml:"COMMISSIONCODE"`
	CommissionType      string `xml:"COMMISSIONTYPE"`
	ProfitCenterCust    string `xml:"PROFITCENTRECUST"`
	ReturnToDept        string `xml:"RETURNTODEPT"`
	FedFunds            string `xml:"FEDFUNDS"`
	PositionType        string `xml:"POSITIONTYPE"`
	AmountDebited       string `xml:"AMOUNTDEBITED"`
	AmountCredited      string `xml:"AMOUNTCREDITED"`
	LocalAmountDebited  string `xml:"LOCAMTDEBITED"`
	LocalAmountCredited string `xml:"LOCAMTCREDITED"`
	CustGroupLevel      string `xml:"CUSTGROUPLEVEL"`
	DebitCustomer       string `xml:"DEBITCUSTOMER"`
	CreditCustomer      string `xml:"CREDITCUSTOMER"`
	DrAdviceReqd        string `xml:"DRADVICEREQDYN"`
	CrAdviceReqd        string `xml:"CRADVICEREQDYN"`
	ChargedCustomer     string `xml:"CHARGEDCUSTOMER"`
	TotalRecComm        string `xml:"TOTRECCOMM"`
	TotalRecCommLcl     string `xml:"TOTRECCOMMLCL"`
	TotalRecCommCur     string `xml:"TOTRECCOMMCUR"`
	TotalRecCommLclCur  string `xml:"TOTRECCOMMLCLCUR"`
	ChargeCode          string `xml:"CHARGECODE"`
	ClientReference     string `xml:"ClientReference"`
	DeliveryInRef       string `xml:"DELIVERYINREF"`
	DeliveryOutRef      string `xml:"DELIVERYOUTREF"`
	ChargeType          string `xml:"CHARGETYPE"`
	ChargeAmount        string `xml:"CHARGEAMOUNT"`
	ChargeCurrency      string `xml:"CHARGECURRENCY"`
	ChargeDate          string `xml:"CHARGEDATE"`
	ChargeTime          string `xml:"CHARGETIME"`
	ChargeStatus        string `xml:"CHARGESTATUS"`
	ChargeDescription   string `xml:"CHARGEDESCRIPTION"`
}

type BillPaymentResult struct {
	Status  bool
	Detail  *BillPaymentDetail
	Message string
}

func ParseBillPaymentSOAP(xmlData string) (*BillPaymentResult, error) {
	var env Envelope
	if err := xml.Unmarshal([]byte(xmlData), &env); err != nil {
		return nil, err
	}

	if env.Body.FTBillPaymentResponse != nil {
		resp := env.Body.FTBillPaymentResponse
		if resp.Status == nil {
			return &BillPaymentResult{
				Status:  false,
				Message: "Missing Status",
			}, nil
		}
		if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
			return &BillPaymentResult{
				Status:  false,
				Message: "API returned failure",
			}, nil
		}
		if resp.BillPaymentDetail == nil {
			return &BillPaymentResult{
				Status:  false,
				Message: "Missing Bill Payment Detail",
			}, nil
		}

		return &BillPaymentResult{
			Status: true,
			Detail: resp.BillPaymentDetail,
		}, nil
	}

	return &BillPaymentResult{
		Status:  false,
		Message: "Invalid response type",
	}, nil
}
