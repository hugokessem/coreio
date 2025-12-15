package main

import (
	"context"
	"fmt"
	"log"

	"github.com/hugokessem/coreio/core"
)

type CoreAPI struct {
	coreInterface core.CBECoreAPIInterface
}

func InitCoreAPICalls(username, password, url string) CoreAPI {
	return CoreAPI{
		coreInterface: core.NewCBECoreAPI(core.CBECoreCredential{
			Username: username,
			Password: password,
			Url:      url,
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

	calls := InitCoreAPICalls("SUPERAPP", "123456", "https://devopscbe.eaglelionsystems.com/superapp/parser/proxy/CBESUPERAPP/services?target=http%3A%2F%2F10.1.15.195%3A8080&wsdl=null")
	ft := core.FundTransferParam{
		// DebitAccountNumber:  "1000382499388",
		// CreditAccountNumber: "1000000006924",

		CreditAccountNumber: "1000382499388",
		DebitAccountNumber:  "1000000006924",
		DebitCurrency:       "ETB",
		CreditCurrency:      "ETB",
		DebitAmount:         "260.00",
		TransactionID:       "TXN12345689",
		DebitReference:      "Payment",
		CreditReference:     "Received payment",
		PaymentDetail:       "Fund transfer",
	}

	result, err := calls.FT(ft)
	// ft := core.MiniStatementByDateRangeParam{
	// 	AccountNumber: "1000184349713",
	// 	From:          "20200101",
	// 	To:            "20200105",
	// }

	// result, err := calls.MiniStatementByDate(ctx, ft)
	if err != nil {
		log.Fatalf("%v", err)
		return
	}

	fmt.Println("result", result)
	if result.Success {
		fmt.Println("amount", result)
		fmt.Println("TransctionID", result.Detail.TransactionID)
		fmt.Println("ft", result.Detail.FTNumber)
		return
	} else {
		// fmt.Println("error", result.Message)
		fmt.Println("error", result.Messages)
	}

	fmt.Println("failed to make ft")
}
