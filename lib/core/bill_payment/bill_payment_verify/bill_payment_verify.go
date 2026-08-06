package billpaymentverify

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type Param struct {
	Username            string
	Password            string
	DebitAccountNumber  string
	DebitCurrency       string
	DebitAmount         string
	DebitReference      string
	CreditReference     string
	CreditAccountNumber string
	CreditCurrency      string
	CreditAmount        string
	ServiceCode         string
	PaymentDetails      string
	ClientReference     string
	ServiceDescription  string
}

type BillPaymentVerifyParam struct {
	DebitAccountNumber  string
	DebitCurrency       string
	DebitAmount         string
	DebitReference      string
	CreditReference     string
	CreditAccountNumber string
	CreditCurrency      string
	CreditAmount        string
	ServiceCode         string
	PaymentDetails      string
	ClientReference     string
	ServiceDescription  string
}

func NewBillPaymentVerify(param Param) string {
	return fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cbes="http://temenos.com/CBESUPERAPP" xmlns:fun="http://temenos.com/FUNDSTRANSFERBILLPAYSUPERAPP">
    <soapenv:Header/>
    <soapenv:Body>
        <cbes:FTBillPayment_Validate>
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
                <fun:CREDITAMOUNT>%s</fun:CREDITAMOUNT>
                <fun:gPAYMENTDETAILS g="%s">
                    <fun:PAYMENTDETAILS>%s</fun:PAYMENTDETAILS>
                </fun:gPAYMENTDETAILS>
                <fun:ClientReference>%s</fun:ClientReference>
                <fun:ServiceDescription>%s</fun:ServiceDescription>
            </FUNDSTRANSFERBILLPAYSUPERAPPType>
        </cbes:FTBillPayment_Validate>
    </soapenv:Body>
</soapenv:Envelope>`, param.Password, param.Username, param.DebitAccountNumber, param.DebitCurrency, param.DebitAmount, param.DebitReference, param.CreditReference, param.CreditAccountNumber, param.CreditCurrency, param.CreditAmount, param.ServiceCode, param.PaymentDetails, param.ClientReference, param.ServiceDescription)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	FTBillPaymentValidateResponse *FTBillPaymentValidateResponse `xml:"FTBillPayment_ValidateResponse"`
}

type FTBillPaymentValidateResponse struct {
	Status           *Status             `xml:"Status"`
	FundTransferType *FundTransferDetail `xml:"FUNDSTRANSFERType"`
}

type Status struct {
	SuccessIndicator string   `xml:"successIndicator"`
	TransactionId    string   `xml:"transactionId"`
	MessageId        string   `xml:"messageId"`
	Application      string   `xml:"application"`
	Messages         []string `xml:"messages"`
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
	CreditValueDate      string   `xml:"CREDITVALUEDATE"`
	ProcessingDate       string   `xml:"PROCESSINGDATE"`
	AccountCategory      string   `xml:"SEPACATEG"`
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
	DebitAmountWithCurrency      string `xml:"AMOUNTDEBITED"`
	CreditAmountWithCurrency     string `xml:"AMOUNTCREDITED"`
	TotalChargeAmount            string `xml:"TOTALCHARGEAMT"`
	TotalTaxAmount               string `xml:"TOTALTAXAMOUNT"`
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
	RoundType                    string `xml:"ROUNDTYPE"`
	GlobalOverride               struct {
		Override []string `xml:"OVERRIDE"`
	} `xml:"gOVERRIDE"`
	CompanyCode                        string `xml:"COCODE"`
	TransactionID                      string `xml:"MTOREF"`
	DebitAccountHolderName             string `xml:"SENDERNAME"`
	ReceiverName                       string `xml:"RECEIVERNAME"`
	DebitAccountCurrentWorkingBalance  string `xml:"CEKCS"`
	CreditAccountCurrentWorkingBalance string `xml:"GPONU"`
	CreditPhoneNumber                  string `xml:"PHONENUM"`
	DebitPhoneNumber                   string `xml:"PHONE"`
	Segment                            string `xml:"SASEGMENT"`
	ServiceDescription                 string `xml:"SERVICEDESC"`
	DisasterReservedFund               string
	OriginalPaidAmount                 string
	TotalCommisionWithComission        string
}

type BillPaymentVerifyResult struct {
	Success  bool
	Detail   *FundTransferDetail
	Messages []string
}

func ParseBillPaymentVerifySOAP(xmlData string) (*BillPaymentVerifyResult, error) {
	var env Envelope
	if err := xml.Unmarshal([]byte(xmlData), &env); err != nil {
		return nil, err
	}

	if env.Body.FTBillPaymentValidateResponse == nil {
		return &BillPaymentVerifyResult{
			Success:  false,
			Messages: []string{"Invalid response type"},
		}, nil
	}

	resp := env.Body.FTBillPaymentValidateResponse
	if resp.Status == nil {
		return &BillPaymentVerifyResult{
			Success:  false,
			Messages: []string{"Missing Status!"},
		}, nil
	}

	if strings.ToLower(resp.Status.SuccessIndicator) != "success" {
		messages := resp.Status.Messages
		if len(messages) == 0 {
			if msg := strings.TrimSpace(resp.Status.MessageId); msg != "" {
				messages = []string{msg}
			}
		}
		return &BillPaymentVerifyResult{
			Success:  false,
			Messages: messages,
		}, nil
	}

	if resp.FundTransferType == nil {
		return &BillPaymentVerifyResult{
			Success:  true,
			Messages: []string{},
		}, nil
	}

	var totalComission float64
	var disasterRecoveryFund float64
	totalTaxAmount, _ := strconv.ParseFloat(resp.FundTransferType.LocalTotalTaxAmount, 64)
	debitCurrency := resp.FundTransferType.DebitCurrency
	creditCurrency := resp.FundTransferType.CreditCurrency
	var currency string
	if debitCurrency == creditCurrency {
		currency = debitCurrency
	} else {
		currency = creditCurrency
	}

	for _, v := range resp.FundTransferType.GlobalCommissionType.MultipleCommissionType {
		if v.CommissionType == "CBECOMSPDIS" {
			if v.CommissionAmount != "" {
				dr, _ := strconv.ParseFloat(strings.TrimPrefix(v.CommissionAmount, currency), 64)
				disasterRecoveryFund += dr
			}
		}

		if v.CommissionType == "CARDDRT" {
			if v.CommissionAmount != "" {
				dr, _ := strconv.ParseFloat(strings.TrimPrefix(v.CommissionAmount, currency), 64)
				disasterRecoveryFund += dr
			}
		}

		if v.CommissionType == "COMFTATM" {
			if v.CommissionAmount != "" {
				totalTaxAmount, _ = strconv.ParseFloat(strings.TrimPrefix(v.CommissionAmount, currency), 64)
			}
		}

		if v.CommissionType == "DRFWALLETSP" {
			if v.CommissionAmount != "" {
				dr, _ := strconv.ParseFloat(strings.TrimPrefix(v.CommissionAmount, currency), 64)
				disasterRecoveryFund += dr
			}
		}

		if v.CommissionType == "ECOMDRFSP" {
			if v.CommissionAmount != "" {
				dr, _ := strconv.ParseFloat(strings.TrimPrefix(v.CommissionAmount, currency), 64)
				disasterRecoveryFund += dr
			}
		}

		if v.CommissionType == "IPSDRFLATSP" {
			if v.CommissionAmount != "" {
				dr, _ := strconv.ParseFloat(strings.TrimPrefix(v.CommissionAmount, currency), 64)
				disasterRecoveryFund += dr
			}
		}

		if v.CommissionType == "IPSPDRPCSP" {
			if v.CommissionAmount != "" {
				dr, _ := strconv.ParseFloat(strings.TrimPrefix(v.CommissionAmount, currency), 64)
				disasterRecoveryFund += dr
			}
		}
	}

	totalChargedAmount, _ := strconv.ParseFloat(resp.FundTransferType.LocalChargeAmount, 64)
	totalComission = totalChargedAmount - totalTaxAmount - disasterRecoveryFund

	originalDebitAmountStr := strings.TrimSpace(strings.TrimPrefix(resp.FundTransferType.DebitAmountWithCurrency, debitCurrency))
	originalDebitAmountWithoutCurrency, err := strconv.ParseFloat(originalDebitAmountStr, 64)
	if err != nil && strings.TrimSpace(resp.FundTransferType.DebitAmount) != "" {
		originalDebitAmountWithoutCurrency, err = strconv.ParseFloat(strings.TrimSpace(resp.FundTransferType.DebitAmount), 64)
		if err != nil {
			originalDebitAmountWithoutCurrency = 0
		}
	}

	var originalTotalChargeAmountWithoutCurrency float64
	trimmedTotalChargeAmount := strings.TrimSpace(resp.FundTransferType.TotalChargeAmount)
	if trimmedTotalChargeAmount != "" {
		originalTotalChargeAmountWithoutCurrency, err = strconv.ParseFloat(strings.TrimSpace(strings.TrimPrefix(trimmedTotalChargeAmount, debitCurrency)), 64)
		if err != nil {
			originalTotalChargeAmountWithoutCurrency, err = strconv.ParseFloat(trimmedTotalChargeAmount, 64)
			if err != nil {
				originalTotalChargeAmountWithoutCurrency = 0
			}
		}
	}

	originalPaidAmount := originalDebitAmountWithoutCurrency - originalTotalChargeAmountWithoutCurrency
	originalPaidAmountWithCurrency := fmt.Sprintf("%s%.4f", debitCurrency, originalPaidAmount)
	totalServiceChargeWithCurrency := fmt.Sprintf("%s%.4f", debitCurrency, totalComission)
	totalTaxAmountWithCurrency := fmt.Sprintf("%s%.4f", currency, totalTaxAmount)
	disasterRecoveryFundWithCurrency := fmt.Sprintf("%s%.4f", currency, disasterRecoveryFund)

	return &BillPaymentVerifyResult{
		Success: true,
		Detail: &FundTransferDetail{
			FTNumber:                           resp.FundTransferType.FTNumber,
			TransactionType:                    resp.FundTransferType.TransactionType,
			DebitAccountNumber:                 resp.FundTransferType.DebitAccountNumber,
			CurrencyMarketDebit:                resp.FundTransferType.CurrencyMarketDebit,
			DebitCurrency:                      resp.FundTransferType.DebitCurrency,
			DebitAmount:                        resp.FundTransferType.DebitAmount,
			DebitValueDate:                     resp.FundTransferType.DebitValueDate,
			DebitReference:                     resp.FundTransferType.DebitReference,
			CreditReference:                    resp.FundTransferType.CreditReference,
			CreditAccountNumber:                resp.FundTransferType.CreditAccountNumber,
			CurrencyMarketCredit:               resp.FundTransferType.CurrencyMarketCredit,
			CreditCurrency:                     resp.FundTransferType.CreditCurrency,
			CreditValueDate:                    resp.FundTransferType.CreditValueDate,
			ProcessingDate:                     resp.FundTransferType.ProcessingDate,
			PaymentDetails:                     resp.FundTransferType.PaymentDetails,
			ChargeCommisionDisplay:             resp.FundTransferType.ChargeCommisionDisplay,
			CommissionCode:                     resp.FundTransferType.CommissionCode,
			GlobalCommissionType:               resp.FundTransferType.GlobalCommissionType,
			ChargeCode:                         resp.FundTransferType.ChargeCode,
			ProfitCentreCustomer:               resp.FundTransferType.ProfitCentreCustomer,
			ReturnToDept:                       resp.FundTransferType.ReturnToDept,
			FedFunds:                           resp.FundTransferType.FedFunds,
			PositionType:                       resp.FundTransferType.PositionType,
			GlobalTaxType:                      resp.FundTransferType.GlobalTaxType,
			DebitAmountWithCurrency:            resp.FundTransferType.DebitAmountWithCurrency,
			CreditAmountWithCurrency:           resp.FundTransferType.CreditAmountWithCurrency,
			TotalChargeAmount:                  resp.FundTransferType.TotalChargeAmount,
			TotalTaxAmount:                     fmt.Sprintf("%.4f", totalTaxAmount),
			CreditCompanyCode:                  resp.FundTransferType.CreditCompanyCode,
			DebitCompanyCode:                   resp.FundTransferType.DebitCompanyCode,
			LocalAmountDebited:                 resp.FundTransferType.LocalAmountDebited,
			LocalAmountCredited:                resp.FundTransferType.LocalAmountCredited,
			LocalTotalTaxAmount:                totalTaxAmountWithCurrency,
			LocalChargeAmount:                  resp.FundTransferType.LocalChargeAmount,
			LocalPositionChargesAmount:         resp.FundTransferType.LocalPositionChargesAmount,
			CustomerGroupLevel:                 resp.FundTransferType.CustomerGroupLevel,
			DebitCustomer:                      resp.FundTransferType.DebitCustomer,
			CreditCustomer:                     resp.FundTransferType.CreditCustomer,
			DebitAdviceRequired:                resp.FundTransferType.DebitAdviceRequired,
			CreditAdviceRequired:               resp.FundTransferType.CreditAdviceRequired,
			ChargedCustomer:                    resp.FundTransferType.ChargedCustomer,
			TotalReceivedCommission:            resp.FundTransferType.TotalReceivedCommission,
			TotalReceivedCommissionLocal:       resp.FundTransferType.TotalReceivedCommissionLocal,
			TotalReceivedCharge:                resp.FundTransferType.TotalReceivedCharge,
			TotalReceivedChargeLocal:           resp.FundTransferType.TotalReceivedChargeLocal,
			RateFixing:                         resp.FundTransferType.RateFixing,
			TotalReceivedChargeCurrency:        resp.FundTransferType.TotalReceivedChargeCurrency,
			TotalSentChargeCurrency:            resp.FundTransferType.TotalSentChargeCurrency,
			RoundType:                          resp.FundTransferType.RoundType,
			GlobalOverride:                     resp.FundTransferType.GlobalOverride,
			CompanyCode:                        resp.FundTransferType.CompanyCode,
			TransactionID:                      resp.FundTransferType.TransactionID,
			DebitAccountHolderName:             resp.FundTransferType.DebitAccountHolderName,
			ReceiverName:                       resp.FundTransferType.ReceiverName,
			DebitAccountCurrentWorkingBalance:  resp.FundTransferType.DebitAccountCurrentWorkingBalance,
			CreditAccountCurrentWorkingBalance: resp.FundTransferType.CreditAccountCurrentWorkingBalance,
			CreditPhoneNumber:                  resp.FundTransferType.CreditPhoneNumber,
			DebitPhoneNumber:                   resp.FundTransferType.DebitPhoneNumber,
			DisasterReservedFund:               disasterRecoveryFundWithCurrency,
			OriginalPaidAmount:                 originalPaidAmountWithCurrency,
			TotalCommisionWithComission:        totalServiceChargeWithCurrency,
			Segment:                            resp.FundTransferType.Segment,
			AccountCategory:                    resp.FundTransferType.AccountCategory,
			ServiceDescription:                 resp.FundTransferType.ServiceDescription,
		},
		Messages: resp.Status.Messages,
	}, nil
}
