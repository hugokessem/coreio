package main

import (
	"context"
	"fmt"
	"log"

	"github.com/hugokessem/coreio/core"
	frauddetection "github.com/hugokessem/coreio/lib/core/fraud_detection"
)

type CoreAPI struct {
	coreInterface core.CBECoreAPIInterface
}

func InitCoreAPICalls(username, password, url, fraudAPIAuth, fraudAPIUrl, fraudAPIForwardHost string) CoreAPI {
	return CoreAPI{
		coreInterface: core.NewCBECoreAPI(core.CBECoreCredential{
			Username: username,
			Password: password,
			Url:      url,
			FraudAPICredential: core.FraudAPICredential{
				Authorization: fraudAPIAuth,
				Url:           fraudAPIUrl,
				ForwardHost:   fraudAPIForwardHost,
			},
		}),
	}
}

func (c *CoreAPI) FT(ft core.FundTransferParam) (*core.FundTransferResult, error) {
	result, err := c.coreInterface.FundTransfer(ft)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *CoreAPI) MiniStatementByDate(ctx context.Context, ft core.MiniStatementByDateRangeParam) (*core.MiniStatementByDateRangeResult, error) {
	result, err := c.coreInterface.MiniStatementByDateRange(ft)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func main() {
	// ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// defer cancel()

	calls := InitCoreAPICalls(
		"SUPERAPP",
		"123456",
		"https://devopscbe.eaglelionsystems.com/superapp/parser/proxy/CBESUPERAPP/services?target=http%3A%2F%2F10.1.15.195%3A8080&wsdl=null",
		"Basic YWRtaW46YWRtaW4=",
		"https://devapisuperapp.cbe.com.et/superapp/parser/proxy/scoringapi/digital-transactions/?target=https://ngdrfms.cbe.com.et",
		"nguat.cbe.com.et",
	)
	ft := core.FundTransferParam{
		DebitAccountNumber:  "1000517052152",
		CreditAccountNumber: "1000000006924",

		// CreditAccountNumber: "1000517052152",
		// DebitAccountNumber: "1000319950331", // usd account
		// DebitCurrency: "ETB",
		// DebitAmount:         "2.00",
		CreditAmount:    "12",
		CreditCurrency:  "ETB",
		TransactionID:   "TXN123458889",
		DebitReference:  "Payment",
		CreditReference: "Received payment",
		PaymentDetail:   "Fund transfer",
		ServiceCode:     "GLOBAL",
		Meta: frauddetection.FraudAPIPayload{
			TranasctionID:              "FT24330T1NSA3",
			AccountID:                  "1000517052152",
			CustomerName:               "YOHHANES TESHOME SHIFERAW",
			CustomerPhoneMobileSMS:     "+251911706628",
			BeneficiaryAccountID:       "1000000006924",
			BeneficiaryName:            "ABIY HAILEYESUS MENGISTU",
			AccountCategory:            "6502",
			AccountCurrency:            "ETB",
			TransactionConvertedAmount: "180",
			TransactionType:            "Mobile Transfer",
			SourceUser:                 "104723KIK",
			ChangeInPhoneEmail:         "Y",
			TransactionTimestamp:       "2025-10-28 09:27:20",
			ChangeInPIN:                "Y",
			ChangeInPassword:           "N",
			ChangeInDevice:             "N",
		},
	}

	// ft := core.MiniStatementByDateRangeParam{
	// 	AccountNumber: "1000184349713",
	// 	From:          "20200101",
	// 	To:            "20200105",
	// }

	// result, err := calls.MiniStatementByDate(ctx, ft)
	result, err := calls.FT(ft)
	if err != nil {
		log.Fatalf("%v", err)
		return
	}

	fmt.Println("result", result)
	if result.Success {
		fmt.Println("amount", result)
		fmt.Println("detai;", result.Detail)
		fmt.Println("TransctionID", result.Detail.TransactionID)
		fmt.Println("ft", result.Detail.FTNumber)
		return
	} else {
		// fmt.Println("error", result.Message)
		fmt.Println("error", result.Messages)
	}

	fmt.Println("failed to make ft")
}
