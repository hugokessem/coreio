package fundtransfercheck

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type Params struct {
	Username string
	Password string
	FTNumber string
}

type FundTransferCheckParams struct {
	FTNumber string
}

func NewFundTransferCheck(param Params) string {
	return fmt.Sprintf(`<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cbes="http://temenos.com/CBESUPERAPP">
    <soapenv:Header/>
    <soapenv:Body>
        <cbes:TransferViewDetails>
            <WebRequestCommon>
                <company></company>
                <password>%s</password>
                <userName>%s</userName>
            </WebRequestCommon>
            <FUNDSTRANSFERVIEWDETAILSSUPERAPPType>
                <transactionId>%s</transactionId>
            </FUNDSTRANSFERVIEWDETAILSSUPERAPPType>
        </cbes:TransferViewDetails>
    </soapenv:Body>
</soapenv:Envelope>`, param.Password, param.Username, param.FTNumber)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	TransferViewDetailsResponse *TransferViewDetailsResponse `xml:"TransferViewDetailsResponse"`
}

type TransferViewDetailsResponse struct {
	Status *struct {
		SuccessIndicator string   `xml:"successIndicator"`
		TransactionId    string   `xml:"transactionId"`
		Application      string   `xml:"application"`
		Messages         []string `xml:"messages"`
	} `xml:"Status"`
	FundTransferType *FundTransferType `xml:"FUNDSTRANSFERType"`
}

type FundTransferType struct {
	TransactionType         string `xml:"TRANSACTIONTYPE"`
	DebitAccountNumber      string `xml:"DEBITACCTNO"`
	DebitCurrency           string `xml:"DEBITCURRENCY"`
	DebitAmount             string `xml:"DEBITAMOUNT"`
	DebitValuedate          string `xml:"DEBITVALUEDATE"`
	CreditAccountNumber     string `xml:"CREDITACCTNO"`
	CreditCurrency          string `xml:"CREDITCURRENCY"`
	CreditAmount            string `xml:"CREDITAMOUNT"`
	CreditValuedate         string `xml:"CREDITVALUEDATE"`
	ProcessingDate          string `xml:"PROCESSINGDATE"`
	ChargeCommissionDisplay struct {
		MultipleCommissionType []struct {
			CommissionType   string `xml:"COMMISSIONTYPE"`
			CommissionAmount string `xml:"COMMISSIONAMT"`
		} `xml:"mCOMMISSIONTYPE"`
	} `xml:"gCOMMISSIONTYPE"`
	CurrentRate              string        `xml:"CUSTOMERRATE"`
	CommissionCode           string        `xml:"COMMISSIONCODE"`
	ChargeCode               string        `xml:"CHARGECODE"`
	ProfitCenterCustomer     string        `xml:"PROFITCENTRECUST"`
	ReturnToDept             string        `xml:"RETURNTODEPT"`
	FedFunds                 string        `xml:"FEDFUNDS"`
	PositionType             string        `xml:"POSITIONTYPE"`
	DebitAmountWithCurrency  string        `xml:"AMOUNTDEBITED"`
	CreditAmountWithCurrency string        `xml:"AMOUNTCREDITED"`
	TotalChargeAmount        string        `xml:"TOTALCHARGEAMT"`
	CreditCompCode           string        `xml:"CREDITCOMPCODE"`
	DebitCompCode            string        `xml:"DEBITCOMPCODE"`
	LocAmtDebited            string        `xml:"LOCAMTDEBITED"`
	LocAmtCredited           string        `xml:"LOCAMTCREDITED"`
	LocalChargeAmount        string        `xml:"LOCALCHARGEAMT"`
	LocalTotalTaxAmount      string        `xml:"LOCTOTTAXAMT"`
	LocalPosChgsAmount       string        `xml:"LOCPOSCHGSAMT"`
	CustGroupLevel           string        `xml:"CUSTGROUPLEVEL"`
	DebitCustomerName        string        `xml:"SENDERNAME"`
	CreditCustomerName       string        `xml:"RECEIVERNAME"`
	DebitCustomerNumber      string        `xml:"DEBITCUSTOMER"`
	CreditCustomerNumber     string        `xml:"CREDITCUSTOMER"`
	DrAdvicerEqdYN           string        `xml:"DRADVICEREQDYN"`
	CrAdvicerEqdYN           string        `xml:"CRADVICEREQDYN"`
	ChargedCustomer          string        `xml:"CHARGEDCUSTOMER"`
	TotRecComm               string        `xml:"TOTRECCOMM"`
	TotRecCommLcl            string        `xml:"TOTRECCOMMLCL"`
	TotRecChg                string        `xml:"TOTRECCHG"`
	TotRecChgLcl             string        `xml:"TOTRECCHGLCL"`
	RateFixing               string        `xml:"RATEFIXING"`
	TotRecChgCrcCy           string        `xml:"TOTRECCHGCRCCY"`
	TotSndChgCrcCy           string        `xml:"TOTSNDCHGCRCCY"`
	AuthDate                 string        `xml:"AUTHDATE"`
	RoundType                string        `xml:"ROUNDTYPE"`
	PaymentDetail            PaymentDetail `xml:"gPAYMENTDETAILS"`
	GlobalTaxType            struct {
		TaxType []struct {
			TaxType   string `xml:"TAXTYPE"`
			TaxAmount string `xml:"TAXAMT"`
		} `xml:"mTAXTYPE"`
	} `xml:"gTAXTYPE"`
	StatementNos struct {
		StatementNo []string `xml:"STMTNOS"`
	} `xml:"gSTMTNOS"`
	CurrNo    string `xml:"CURRNO"`
	GInputter struct {
		Inputter string `xml:"INPUTTER"`
	} `xml:"gINPUTTER"`
	GDatetime struct {
		Datetime string `xml:"DATETIME"`
	} `xml:"gDATETIME"`
	Authoriser                  string `xml:"AUTHORISER"`
	CoCode                      string `xml:"COCODE"`
	DeptCode                    string `xml:"DEPTCODE"`
	Question                    string `xml:"QUESTION"`
	SecAnswer                   string `xml:"SECANSWER"`
	SecNumber                   string `xml:"SECNUMBER"`
	LmtssSendNo                 string `xml:"LMTSSENDNO"`
	DisasterReservedFund        string `xml:"DISASTERRESERVEFUND"`
	OriginalPaidAmount          string `xml:"ORIGPAIDAMT"`
	TotalCommisionWithComission string `xml:"TOTALCOMISSON"`
	DebitReference              string `xml:"DEBITTHEIRREF"`
	CreditReference             string `xml:"CREDITTHEIRREF"`
	ServiceCode                 string `xml:"SERVICECODE"`
}

type PaymentDetail struct {
	PaymentDetail string `xml:"PAYMENTDETAILS"`
}

type FundTransferCheckResult struct {
	Status   bool
	Detail   *FundTransferType
	Messages []string
}

func ParseFundTransferCheckSOAP(xmlData string) (*FundTransferCheckResult, error) {
	var env Envelope
	err := xml.Unmarshal([]byte(xmlData), &env)
	if err != nil {
		return nil, err
	}

	if env.Body.TransferViewDetailsResponse != nil {
		resp := env.Body.TransferViewDetailsResponse
		if resp.Status == nil {
			return &FundTransferCheckResult{
				Status:   false,
				Messages: []string{"Missing Status!"},
			}, nil
		}
		if resp.Status.SuccessIndicator != "Success" {
			return &FundTransferCheckResult{
				Status:   false,
				Messages: []string{"API return failur!"},
			}, nil
		}
		if resp.FundTransferType == nil {
			return &FundTransferCheckResult{
				Status:   true,
				Messages: resp.Status.Messages,
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
		for _, v := range resp.FundTransferType.ChargeCommissionDisplay.MultipleCommissionType {
			// REMAINIG
			if v.CommissionType == "CBECOMSPDIS" {
				if v.CommissionAmount != "" {
					dr, _ := strconv.ParseFloat(strings.TrimPrefix(v.CommissionAmount, currency), 64)
					disasterRecoveryFund += dr
				}
			}

			// VCARD, PMCARD, LOCAL CARD
			if v.CommissionType == "CARDDRT" {
				if v.CommissionAmount != "" {
					dr, _ := strconv.ParseFloat(strings.TrimPrefix(v.CommissionAmount, currency), 64)
					disasterRecoveryFund += dr
				}
			}

			// CARD
			if v.CommissionType == "COMFTATM" {
				if v.CommissionAmount != "" {
					totalTaxAmount, _ = strconv.ParseFloat(strings.TrimPrefix(v.CommissionAmount, currency), 64) // vat
				}
			}

			// TELEBURR, MPESA, EBIRR
			if v.CommissionType == "DRFWALLETSP" {
				if v.CommissionAmount != "" {
					dr, _ := strconv.ParseFloat(strings.TrimPrefix(v.CommissionAmount, currency), 64)
					disasterRecoveryFund += dr
				}
			}

			// ECOMMERCE
			if v.CommissionType == "ECOMDRFSP" {
				if v.CommissionAmount != "" {
					dr, _ := strconv.ParseFloat(strings.TrimPrefix(v.CommissionAmount, currency), 64)
					disasterRecoveryFund += dr
				}
			}

			// IPS
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

		return &FundTransferCheckResult{
			Status: true,
			Detail: &FundTransferType{
				TransactionType:             resp.FundTransferType.TransactionType,
				DebitAccountNumber:          resp.FundTransferType.DebitAccountNumber,
				DebitCurrency:               resp.FundTransferType.DebitCurrency,
				DebitAmount:                 resp.FundTransferType.DebitAmount,
				DebitValuedate:              resp.FundTransferType.DebitValuedate,
				CreditAccountNumber:         resp.FundTransferType.CreditAccountNumber,
				CreditCurrency:              resp.FundTransferType.CreditCurrency,
				CreditAmount:                resp.FundTransferType.CreditAmount,
				CreditValuedate:             resp.FundTransferType.CreditValuedate,
				ProcessingDate:              resp.FundTransferType.ProcessingDate,
				ChargeCommissionDisplay:     resp.FundTransferType.ChargeCommissionDisplay,
				CommissionCode:              resp.FundTransferType.CommissionCode,
				ChargeCode:                  resp.FundTransferType.ChargeCode,
				ProfitCenterCustomer:        resp.FundTransferType.ProfitCenterCustomer,
				ReturnToDept:                resp.FundTransferType.ReturnToDept,
				FedFunds:                    resp.FundTransferType.FedFunds,
				PositionType:                resp.FundTransferType.PositionType,
				DebitAmountWithCurrency:     resp.FundTransferType.DebitAmountWithCurrency,
				CreditAmountWithCurrency:    resp.FundTransferType.CreditAmountWithCurrency,
				TotalChargeAmount:           resp.FundTransferType.TotalChargeAmount,
				CreditCompCode:              fmt.Sprintf("%.4f", totalTaxAmount),
				DebitCompCode:               resp.FundTransferType.DebitCompCode,
				LocAmtDebited:               resp.FundTransferType.LocAmtDebited,
				LocAmtCredited:              resp.FundTransferType.LocAmtCredited,
				LocalChargeAmount:           resp.FundTransferType.LocalChargeAmount,
				LocalPosChgsAmount:          resp.FundTransferType.LocalPosChgsAmount,
				CustGroupLevel:              resp.FundTransferType.CustGroupLevel,
				DebitCustomerName:           resp.FundTransferType.DebitCustomerName,
				CreditCustomerName:          resp.FundTransferType.CreditCustomerName,
				DebitCustomerNumber:         resp.FundTransferType.DebitCustomerNumber,
				CreditCustomerNumber:        resp.FundTransferType.CreditCustomerNumber,
				DrAdvicerEqdYN:              resp.FundTransferType.DrAdvicerEqdYN,
				CrAdvicerEqdYN:              resp.FundTransferType.CrAdvicerEqdYN,
				ChargedCustomer:             resp.FundTransferType.ChargedCustomer,
				TotRecComm:                  resp.FundTransferType.TotRecComm,
				TotRecCommLcl:               resp.FundTransferType.TotRecCommLcl,
				TotRecChg:                   resp.FundTransferType.TotRecChg,
				TotRecChgLcl:                resp.FundTransferType.TotRecChgLcl,
				RateFixing:                  resp.FundTransferType.RateFixing,
				TotRecChgCrcCy:              resp.FundTransferType.TotRecChgCrcCy,
				TotSndChgCrcCy:              resp.FundTransferType.TotSndChgCrcCy,
				AuthDate:                    resp.FundTransferType.AuthDate,
				RoundType:                   resp.FundTransferType.RoundType,
				StatementNos:                resp.FundTransferType.StatementNos,
				CurrNo:                      resp.FundTransferType.CurrNo,
				GInputter:                   resp.FundTransferType.GInputter,
				GDatetime:                   resp.FundTransferType.GDatetime,
				Authoriser:                  resp.FundTransferType.Authoriser,
				CoCode:                      resp.FundTransferType.CoCode,
				DeptCode:                    resp.FundTransferType.DeptCode,
				Question:                    resp.FundTransferType.Question,
				SecAnswer:                   resp.FundTransferType.SecAnswer,
				SecNumber:                   resp.FundTransferType.SecNumber,
				LmtssSendNo:                 resp.FundTransferType.LmtssSendNo,
				GlobalTaxType:               resp.FundTransferType.GlobalTaxType,
				PaymentDetail:               resp.FundTransferType.PaymentDetail,
				CurrentRate:                 resp.FundTransferType.CurrentRate,
				LocalTotalTaxAmount:         totalTaxAmountWithCurrency,
				DisasterReservedFund:        disasterRecoveryFundWithCurrency,
				OriginalPaidAmount:          originalPaidAmountWithCurrency,
				TotalCommisionWithComission: totalServiceChargeWithCurrency,
				DebitReference:              resp.FundTransferType.DebitReference,
				CreditReference:             resp.FundTransferType.CreditReference,
				ServiceCode:                 resp.FundTransferType.ServiceCode,
			},
		}, nil
	}

	return &FundTransferCheckResult{
		Status:   false,
		Messages: []string{},
	}, nil
}
