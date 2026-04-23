package fundtrasferverify

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type Params struct {
	Username            string
	Password            string
	DebitAccountNumber  string
	DebitCurrency       string
	CreditCurrency      string
	DebitAmount         string
	DebitReference      string
	CreditReference     string
	CreditAccountNumber string
	CreditAmount        string
	PaymentDetails      string
	ClientReference     string
	ServiceCode         string
	CustomerSegment     string
	ChannelType         string
}

type FundTransferVerifyParams struct {
	Username            string
	Password            string
	DebitAccountNumber  string
	DebitCurrency       string
	DebitAmount         string
	DebitReference      string
	CreditReference     string
	CreditAmount        string
	CreditCurrency      string
	CreditAccountNumber string
	PaymentDetails      string
	ClientReference     string
	ServiceCode         string
	CustomerSegment     string
	ChannelType         string
}

func NewFundTransferVerify(params Params) string {
	var details []string
	if params.CreditCurrency == "" || params.DebitCurrency == "" {
		return "Both CreditCurrency and DebitCurrency are Requried!"
	}

	if params.CreditAmount != "" && params.DebitAmount != "" {
		return "Both CreditAmount and DebitAmount cannot be provided together!"
	}

	if params.CreditCurrency == params.DebitCurrency || params.DebitAmount != "" {
		details = append(details, fmt.Sprintf(`
			<fun:DEBITACCTNO>%s</fun:DEBITACCTNO>
            <fun:DEBITCURRENCY>%s</fun:DEBITCURRENCY>
			<fun:DEBITAMOUNT>%s</fun:DEBITAMOUNT>
			<fun:DEBITTHEIRREF>%s</fun:DEBITTHEIRREF>
			<fun:CREDITTHEIRREF>%s</fun:CREDITTHEIRREF>
			<fun:CREDITACCTNO>%s</fun:CREDITACCTNO>
			<fun:CREDITCURRENCY>%s</fun:CREDITCURRENCY>
			`, params.DebitAccountNumber, params.DebitCurrency, params.DebitAmount, params.DebitReference, params.CreditReference, params.CreditAccountNumber, params.CreditCurrency))

	} else {
		details = append(details, fmt.Sprintf(`
			<fun:DEBITACCTNO>%s</fun:DEBITACCTNO>
            <fun:DEBITCURRENCY>%s</fun:DEBITCURRENCY>
			<fun:DEBITTHEIRREF>%s</fun:DEBITTHEIRREF>
			<fun:CREDITTHEIRREF>%s</fun:CREDITTHEIRREF>
			<fun:CREDITACCTNO>%s</fun:CREDITACCTNO>
			<fun:CREDITCURRENCY>%s</fun:CREDITCURRENCY>
			<fun:CREDITAMOUNT>%s</fun:CREDITAMOUNT>
			`, params.DebitAccountNumber, params.DebitCurrency, params.DebitReference, params.CreditReference, params.CreditAccountNumber, params.CreditCurrency, params.CreditAmount))
	}

	return fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cbes="http://temenos.com/CBESUPERAPP" xmlns:fun="http://temenos.com/FUNDSTRANSFERFTTXNSUPERAPP">
    <soapenv:Header/>
    <soapenv:Body>
        <cbes:AccountTransfer_Validate>
            <WebRequestCommon>
                <company/>
                <password>%s</password>
                <userName>%s</userName>
            </WebRequestCommon>
            <OfsFunction/>
            <FUNDSTRANSFERFTTXNSUPERAPPType id="">
				%s
                <fun:gPAYMENTDETAILS g="1">
                    <fun:PAYMENTDETAILS>%s</fun:PAYMENTDETAILS>
                </fun:gPAYMENTDETAILS>
                <fun:ClientReference>%s</fun:ClientReference>
                <fun:ServiceCode>%s</fun:ServiceCode>
                <fun:CustomerRole>%s</fun:CustomerRole>
                <fun:ChannelType>%s</fun:ChannelType>
            </FUNDSTRANSFERFTTXNSUPERAPPType>
        </cbes:AccountTransfer_Validate>
    </soapenv:Body>
</soapenv:Envelope>
	`, params.Password, params.Username, strings.Join(details, "\n"), params.PaymentDetails, params.ClientReference, params.ServiceCode, params.CustomerSegment, params.ChannelType)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	AccountTransferValidateResponse *AccountTransferValidateResponse `xml:"AccountTransfer_ValidateResponse"`
}

type AccountTransferValidateResponse struct {
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
	TreasuryRate                 string `xml:"TREASURYRATE"`
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
	CustomerRole                       string `xml:"CUSTOMERROLE"`
	TransactionChannel                 string `xml:"TXNCHANNEL"`
	BudgetType                         string `xml:"BUDGETTYPE"`
	CreditPhoneNumner                  string `xml:"PHONENUM"`
	DebitPhoneNumner                   string `xml:"PHONE"`
	DisasterReservedFund               string `xml:"DISASTERRESERVEFUND"`
	OriginalPaidAmount                 string `xml:"ORIGPAIDAMT"`
	TotalCommisionWithComission        string `xml:"TOTALCOMISSON"`
}

type FundTransferVerifyResult struct {
	Success  bool
	Detail   *FundTransferDetail
	Messages []string
}

func ParseFundTransferVerifySOAP(xmlData string) (*FundTransferVerifyResult, error) {
	var env Envelope
	if err := xml.Unmarshal([]byte(xmlData), &env); err != nil {
		return nil, err
	}

	if env.Body.AccountTransferValidateResponse != nil {
		resp := env.Body.AccountTransferValidateResponse
		if resp.Status == nil {
			return &FundTransferVerifyResult{
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
			return &FundTransferVerifyResult{
				Success:  false,
				Messages: messages,
			}, nil
		}

		if resp.FundTransferType == nil {
			return &FundTransferVerifyResult{
				Success:  true,
				Messages: []string{},
			}, nil
		}

		var totalComission float64
		totalTaxAmount, _ := strconv.ParseFloat(resp.FundTransferType.LocalTotalTaxAmount, 64)
		debitCurrency := resp.FundTransferType.DebitCurrency
		creditCurrency := resp.FundTransferType.CreditCurrency
		var currency string
		if debitCurrency == creditCurrency {
			currency = debitCurrency
		} else {
			currency = creditCurrency
		}
		dr := fmt.Sprintf("%s0", currency)
		for _, v := range resp.FundTransferType.GlobalCommissionType.MultipleCommissionType {
			if v.CommissionType == "CBECOMSPDIS" {
				if v.CommissionAmount != "" {
					dr = v.CommissionAmount
				}
			}

			if v.CommissionType == "CARDDRT" {
				if v.CommissionAmount != "" {
					dr = v.CommissionAmount
				}
			}

			if v.CommissionType == "COMFTATM" {
				if v.CommissionAmount != "" {
					totalTaxAmount, _ = strconv.ParseFloat(v.CommissionAmount, 64)
				}
			}

		}
		amount, _ := strconv.ParseFloat(strings.TrimPrefix(dr, debitCurrency), 64)
		totalChargedAmount, _ := strconv.ParseFloat(resp.FundTransferType.LocalChargeAmount, 64)
		totalComission = totalChargedAmount - totalTaxAmount - amount

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
		totalTaxAmountWithCurrency := fmt.Sprintf("%s%.4f", debitCurrency, totalTaxAmount)

		return &FundTransferVerifyResult{
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
				CreditValidationDare:               resp.FundTransferType.CreditValidationDare,
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
				TotalTaxAmount:                     fmt.Sprintf("%s", totalTaxAmount),
				DeliveryOutRef:                     resp.FundTransferType.DeliveryOutRef,
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
				AuthDate:                           resp.FundTransferType.AuthDate,
				RoundType:                          resp.FundTransferType.RoundType,
				GlobalStatementNumbers:             resp.FundTransferType.GlobalStatementNumbers,
				GlobalOverride:                     resp.FundTransferType.GlobalOverride,
				CurrentNumber:                      resp.FundTransferType.CurrentNumber,
				GlobalInputter:                     resp.FundTransferType.GlobalInputter,
				GlobalDateTime:                     resp.FundTransferType.GlobalDateTime,
				Authoriser:                         resp.FundTransferType.Authoriser,
				CompanyCode:                        resp.FundTransferType.CompanyCode,
				DepartmentCode:                     resp.FundTransferType.DepartmentCode,
				InputVersion:                       resp.FundTransferType.InputVersion,
				AuthVersion:                        resp.FundTransferType.AuthVersion,
				TransactionID:                      resp.FundTransferType.TransactionID,
				DebitAccountHolderName:             resp.FundTransferType.DebitAccountHolderName,
				ReceiverName:                       resp.FundTransferType.ReceiverName,
				ServiceCode:                        resp.FundTransferType.ServiceCode,
				DebitAccountCurrentWorkingBalance:  resp.FundTransferType.DebitAccountCurrentWorkingBalance,
				CreditAccountCurrentWorkingBalance: resp.FundTransferType.CreditAccountCurrentWorkingBalance,
				CustomerRole:                       resp.FundTransferType.CustomerRole,
				TransactionChannel:                 resp.FundTransferType.TransactionChannel,
				BudgetType:                         resp.FundTransferType.BudgetType,
				CreditPhoneNumner:                  resp.FundTransferType.CreditPhoneNumner,
				DebitPhoneNumner:                   resp.FundTransferType.DebitPhoneNumner,
				TreasuryRate:                       resp.FundTransferType.TreasuryRate,
				DisasterReservedFund:               dr,
				OriginalPaidAmount:                 originalPaidAmountWithCurrency,
				TotalCommisionWithComission:        totalServiceChargeWithCurrency,
			},
			Messages: resp.Status.Messages,
		}, nil
	}

	return &FundTransferVerifyResult{
		Success:  false,
		Messages: []string{"Invalid response type"},
	}, nil
}
