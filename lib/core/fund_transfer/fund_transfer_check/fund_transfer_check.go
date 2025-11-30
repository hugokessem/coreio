package fundtransfercheck

import (
	"encoding/xml"
	"errors"
	"fmt"
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
		SuccessIndicator string `xml:"successIndicator"`
		TransactionId    string `xml:"transactionId"`
		Application      string `xml:"application"`
		MessageId        string `xml:"messageId"`
	} `xml:"Status"`
	FundTransferType *FundTransferType `xml:"FUNDSTRANSFERType"`
}

type FundTransferType struct {
	TransactionType         string                    `xml:"TRANSACTIONTYPE"`
	DebitAccountNumber      string                    `xml:"DEBITACCTNO"`
	DebitCurrency           string                    `xml:"DEBITCURRENCY"`
	DebitAmount             string                    `xml:"DEBITAMOUNT"`
	DebitValuedate          string                    `xml:"DEBITVALUEDATE"`
	CreditAccountNumber     string                    `xml:"CREDITACCTNO"`
	CreditCurrency          string                    `xml:"CREDITCURRENCY"`
	CreditAmount            string                    `xml:"CREDITAMOUNT"`
	CreditValuedate         string                    `xml:"CREDITVALUEDATE"`
	ProcessingDate          string                    `xml:"PROCESSINGDATE"`
	ChargeCommissionDisplay []ChargeCommissionDisplay `xml:"gCOMMISSIONTYPE"`
	CommissionCode          string                    `xml:"COMMISSIONCODE"`
	ChargeCode              string                    `xml:"CHARGECODE"`
	ProfitCenterCustomer    string                    `xml:"PROFITCENTRECUST"`
	ReturnToDept            string                    `xml:"RETURNTODEPT"`
	FedFunds                string                    `xml:"FEDFUNDS"`
	PositionType            string                    `xml:"POSITIONTYPE"`
	AmountDebited           string                    `xml:"AMOUNTDEBITED"`
	AmountCredited          string                    `xml:"AMOUNTCREDITED"`
	TotalChargeAmount       string                    `xml:"TOTALCHARGEAMT"`
	CreditCompCode          string                    `xml:"CREDITCOMPCODE"`
	DebitCompCode           string                    `xml:"DEBITCOMPCODE"`
	LocAmtDebited           string                    `xml:"LOCAMTDEBITED"`
	LocAmtCredited          string                    `xml:"LOCAMTCREDITED"`
	LocalChargeAmount       string                    `xml:"LOCALCHARGEAMT"`
	LocalPosChgsAmount      string                    `xml:"LOCPOSCHGSAMT"`
	CustGroupLevel          string                    `xml:"CUSTGROUPLEVEL"`
	DebitCustomer           string                    `xml:"DEBITCUSTOMER"`
	CreditCustomer          string                    `xml:"CREDITCUSTOMER"`
	DrAdvicerEqdYN          string                    `xml:"DRADVICEREQDYN"`
	CrAdvicerEqdYN          string                    `xml:"CRADVICEREQDYN"`
	ChargedCustomer         string                    `xml:"CHARGEDCUSTOMER"`
	TotRecComm              string                    `xml:"TOTRECCOMM"`
	TotRecCommLcl           string                    `xml:"TOTRECCOMMLCL"`
	TotRecChg               string                    `xml:"TOTRECCHG"`
	TotRecChgLcl            string                    `xml:"TOTRECCHGLCL"`
	RateFixing              string                    `xml:"RATEFIXING"`
	TotRecChgCrcCy          string                    `xml:"TOTRECCHGCRCCY"`
	TotSndChgCrcCy          string                    `xml:"TOTSNDCHGCRCCY"`
	AuthDate                string                    `xml:"AUTHDATE"`
	RoundType               string                    `xml:"ROUNDTYPE"`
	StatementNos            struct {
		StatementNo []string `xml:"STMTNOS"`
	} `xml:"gSTMTNOS"`
	CurrNo    string `xml:"CURRNO"`
	GInputter struct {
		Inputter string `xml:"INPUTTER"`
	} `xml:"gINPUTTER"`
	GDatetime struct {
		Datetime string `xml:"DATETIME"`
	} `xml:"gDATETIME"`
	Authoriser  string `xml:"AUTHORISER"`
	CoCode      string `xml:"COCODE"`
	DeptCode    string `xml:"DEPTCODE"`
	Question    string `xml:"QUESTION"`
	SecAnswer   string `xml:"SECANSWER"`
	SecNumber   string `xml:"SECNUMBER"`
	LmtssSendNo string `xml:"LMTSSENDNO"`
}

type ChargeCommissionDisplay struct {
	CommisionType struct {
		ComissionType   string `xml:"COMMISSIONTYPE"`
		ComissionAmount string `xml:"COMMISSIONAMT"`
	} `xml:"mCOMMISSIONTYPE"`
}

type FundTransferCheckResult struct {
	Status  bool
	Detail  *FundTransferType
	Message string
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
			return nil, errors.New("missing status")
		}
		if resp.Status.SuccessIndicator != "Success" {
			return nil, errors.New(resp.Status.MessageId)
		}
		if resp.FundTransferType == nil {
			return nil, errors.New("missing fund transfer type")
		}
		return &FundTransferCheckResult{
			Status: true,
			Detail: &FundTransferType{
				TransactionType:         resp.FundTransferType.TransactionType,
				DebitAccountNumber:      resp.FundTransferType.DebitAccountNumber,
				DebitCurrency:           resp.FundTransferType.DebitCurrency,
				DebitAmount:             resp.FundTransferType.DebitAmount,
				DebitValuedate:          resp.FundTransferType.DebitValuedate,
				CreditAccountNumber:     resp.FundTransferType.CreditAccountNumber,
				CreditCurrency:          resp.FundTransferType.CreditCurrency,
				CreditAmount:            resp.FundTransferType.CreditAmount,
				CreditValuedate:         resp.FundTransferType.CreditValuedate,
				ProcessingDate:          resp.FundTransferType.ProcessingDate,
				ChargeCommissionDisplay: resp.FundTransferType.ChargeCommissionDisplay,
				CommissionCode:          resp.FundTransferType.CommissionCode,
				ChargeCode:              resp.FundTransferType.ChargeCode,
				ProfitCenterCustomer:    resp.FundTransferType.ProfitCenterCustomer,
				ReturnToDept:            resp.FundTransferType.ReturnToDept,
				FedFunds:                resp.FundTransferType.FedFunds,
				PositionType:            resp.FundTransferType.PositionType,
				AmountDebited:           resp.FundTransferType.AmountDebited,
				AmountCredited:          resp.FundTransferType.AmountCredited,
				TotalChargeAmount:       resp.FundTransferType.TotalChargeAmount,
				CreditCompCode:          resp.FundTransferType.CreditCompCode,
				DebitCompCode:           resp.FundTransferType.DebitCompCode,
				LocAmtDebited:           resp.FundTransferType.LocAmtDebited,
				LocAmtCredited:          resp.FundTransferType.LocAmtCredited,
				LocalChargeAmount:       resp.FundTransferType.LocalChargeAmount,
				LocalPosChgsAmount:      resp.FundTransferType.LocalPosChgsAmount,
				CustGroupLevel:          resp.FundTransferType.CustGroupLevel,
				DebitCustomer:           resp.FundTransferType.DebitCustomer,
				CreditCustomer:          resp.FundTransferType.CreditCustomer,
				DrAdvicerEqdYN:          resp.FundTransferType.DrAdvicerEqdYN,
				CrAdvicerEqdYN:          resp.FundTransferType.CrAdvicerEqdYN,
				ChargedCustomer:         resp.FundTransferType.ChargedCustomer,
				TotRecComm:              resp.FundTransferType.TotRecComm,
				TotRecCommLcl:           resp.FundTransferType.TotRecCommLcl,
				TotRecChg:               resp.FundTransferType.TotRecChg,
				TotRecChgLcl:            resp.FundTransferType.TotRecChgLcl,
				RateFixing:              resp.FundTransferType.RateFixing,
				TotRecChgCrcCy:          resp.FundTransferType.TotRecChgCrcCy,
				TotSndChgCrcCy:          resp.FundTransferType.TotSndChgCrcCy,
				AuthDate:                resp.FundTransferType.AuthDate,
				RoundType:               resp.FundTransferType.RoundType,
				StatementNos:            resp.FundTransferType.StatementNos,
				CurrNo:                  resp.FundTransferType.CurrNo,
				GInputter:               resp.FundTransferType.GInputter,
				GDatetime:               resp.FundTransferType.GDatetime,
				Authoriser:              resp.FundTransferType.Authoriser,
				CoCode:                  resp.FundTransferType.CoCode,
				DeptCode:                resp.FundTransferType.DeptCode,
				Question:                resp.FundTransferType.Question,
				SecAnswer:               resp.FundTransferType.SecAnswer,
				SecNumber:               resp.FundTransferType.SecNumber,
				LmtssSendNo:             resp.FundTransferType.LmtssSendNo,
			},
			Message: resp.Status.MessageId,
		}, nil
	}

	// Handle case where TransferViewDetailsResponse is nil
	message := ""
	if env.Body.TransferViewDetailsResponse != nil && env.Body.TransferViewDetailsResponse.Status != nil {
		message = env.Body.TransferViewDetailsResponse.Status.MessageId
	}

	return &FundTransferCheckResult{
		Status:  false,
		Message: message,
	}, nil
}
