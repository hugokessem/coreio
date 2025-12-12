package statuscheck

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Params struct {
	DebitBankBIC                  string
	BizMessageIdentifier          string
	MessageIdentifier             string
	CreditDateTime                string
	CreditDate                    string
	OriginalTransactionIdentifier string
}

func NewStatusCheck(param Params) string {
	return fmt.Sprintf(
		`<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:mb="http://MB_IPS" xmlns:urn="urn:iso:std:iso:20022:tech:xsd:head.001.001.03" xmlns:urn1="urn:iso:std:iso:20022:tech:xsd:pacs.028.001.05">
    <soapenv:Header/>
    <soapenv:Body>
        <mb:PaymentStatus>
            <input1>
                <urn:AppHdr>
                    <urn:Fr>
                        <urn:FIId>
                            <urn:FinInstnId>
                                <urn:Othr>
                                    <urn:Id>%s</urn:Id>
                                </urn:Othr>
                            </urn:FinInstnId>
                        </urn:FIId>
                    </urn:Fr>
                    <urn:To>
                        <urn:FIId>
                            <urn:FinInstnId>
                                <urn:Othr>
                                    <urn:Id>FP</urn:Id>
                                </urn:Othr>
                            </urn:FinInstnId>
                        </urn:FIId>
                    </urn:To>
                    <urn:BizMsgIdr>%s</urn:BizMsgIdr>
                    <urn:MsgDefIdr>pacs.028.001.05</urn:MsgDefIdr>
                    <urn:CreDt>%s</urn:CreDt>
                </urn:AppHdr>
                <urn1:Document>
                    <urn1:FIToFIPmtStsReq>
                        <urn1:GrpHdr>
                            <urn1:MsgId>%s</urn1:MsgId>
                            <urn1:CreDtTm>%s</urn1:CreDtTm>
                        </urn1:GrpHdr>
                        <urn1:TxInf>
                            <urn1:OrgnlTxId>%s</urn1:OrgnlTxId>
                        </urn1:TxInf>
                    </urn1:FIToFIPmtStsReq>
                </urn1:Document>
            </input1>
        </mb:PaymentStatus>
    </soapenv:Body>
</soapenv:Envelope>`, param.DebitBankBIC, param.BizMessageIdentifier, param.CreditDate, param.MessageIdentifier, param.CreditDateTime, param.OriginalTransactionIdentifier)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	PaymentStatusResponse *PaymentStatusResponse `xml:"PaymentStatusResponse"`
	ErrorResponse         *ErrorResponse         `xml:"errorResponse"`
}

type PaymentStatusResponse struct {
	Output Output `xml:"output1"`
}

type Output struct {
	AppHdr   AppHeader `xml:"AppHdr"`
	Document Document  `xml:"Document"`
}

type AppHeader struct {
	From struct {
		FIId struct {
			FinInstnId struct {
				Othr struct {
					Id string `xml:"Id"`
				} `xml:"Othr"`
			} `xml:"FinInstnId"`
		} `xml:"FIId"`
	} `xml:"Fr"`
	To struct {
		FIId struct {
			FinInstnId struct {
				Othr struct {
					Id string `xml:"Id"`
				} `xml:"Othr"`
			} `xml:"FinInstnId"`
		} `xml:"FIId"`
	} `xml:"To"`
	BizMessageIdentifier string `xml:"BizMsgIdr"`
	MsgDefIdr            string `xml:"MsgDefIdr"`
	CreDt                string `xml:"CreDt"`
	Rltd                 *struct {
		Fr struct {
			FIId struct {
				FinInstnId struct {
					Othr struct {
						Id string `xml:"Id"`
					} `xml:"Othr"`
				} `xml:"FinInstnId"`
			} `xml:"FIId"`
		} `xml:"Fr"`
		To struct {
			FIId struct {
				FinInstnId struct {
					Othr struct {
						Id string `xml:"Id"`
					} `xml:"Othr"`
				} `xml:"FinInstnId"`
			} `xml:"FIId"`
		} `xml:"To"`
		BizMessageIdentifier string `xml:"BizMsgIdr"`
		MsgDefIdr            string `xml:"MsgDefIdr"`
		CreDt                string `xml:"CreDt"`
	} `xml:"Rltd"`
}

type Document struct {
	FIToFIPmtStsRpt struct {
		GrpHdr struct {
			MsgId    string `xml:"MsgId"`
			CreDtTm  string `xml:"CreDtTm"`
			InstgAgt struct {
				FinInstnId struct {
					Othr struct {
						Id string `xml:"Id"`
					} `xml:"Othr"`
				} `xml:"FinInstnId"`
			} `xml:"InstgAgt"`
			InstdAgt struct {
				FinInstnId struct {
					Othr struct {
						Id string `xml:"Id"`
					} `xml:"Othr"`
				} `xml:"FinInstnId"`
			} `xml:"InstdAgt"`
		} `xml:"GrpHdr"`
		TxInfAndSts struct {
			OrgnlGrpInf struct {
				OrgnlMsgId   string `xml:"OrgnlMsgId"`
				OrgnlMsgNmId string `xml:"OrgnlMsgNmId"`
				OrgnlCreDtTm string `xml:"OrgnlCreDtTm"`
			} `xml:"OrgnlGrpInf"`
			OrgnlEndToEndId string `xml:"OrgnlEndToEndId"`
			OrgnlTxId       string `xml:"OrgnlTxId"`
			TxSts           string `xml:"TxSts"`
			AccptncDtTm     string `xml:"AccptncDtTm"`
			OrgnlTxRef      struct {
				IntrBkSttlmAmt struct {
					Amount string `xml:",chardata"`
					Ccy    string `xml:"Ccy,attr"`
				} `xml:"IntrBkSttlmAmt"`
				Amt struct {
					InstdAmt struct {
						Amount string `xml:",chardata"`
						Ccy    string `xml:"Ccy,attr"`
					} `xml:"InstdAmt"`
				} `xml:"Amt"`
				RmtInf struct {
					Ustrd string `xml:"Ustrd"`
				} `xml:"RmtInf"`
				Dbtr struct {
					Pty struct {
						Nm      string `xml:"Nm"`
						PstlAdr *struct {
							AdrLine string `xml:"AdrLine"`
						} `xml:"PstlAdr"`
					} `xml:"Pty"`
				} `xml:"Dbtr"`
				DbtrAcct struct {
					Id struct {
						Othr struct {
							Id      string `xml:"Id"`
							SchmeNm struct {
								Prtry string `xml:"Prtry"`
							} `xml:"SchmeNm"`
						} `xml:"Othr"`
					} `xml:"Id"`
				} `xml:"DbtrAcct"`
				Cdtr struct {
					Pty struct {
						Nm string `xml:"Nm"`
					} `xml:"Pty"`
				} `xml:"Cdtr"`
				CdtrAcct struct {
					Id struct {
						Othr struct {
							Id      string `xml:"Id"`
							SchmeNm struct {
								Prtry string `xml:"Prtry"`
							} `xml:"SchmeNm"`
						} `xml:"Othr"`
					} `xml:"Id"`
				} `xml:"CdtrAcct"`
			} `xml:"OrgnlTxRef"`
		} `xml:"TxInfAndSts"`
	} `xml:"FIToFIPmtStsRpt"`
}

type PaymentStatusDetail struct {
	FromBankBIC                 string
	ToBankBIC                   string
	BizMessageIdentifier        string
	MessageDefinitionIdentifier string
	CreationDate                string
	RelatedFromBankBIC          string
	RelatedToBankBIC            string
	RelatedBizMessageIdentifier string
	RelatedMessageDefinitionId  string
	RelatedCreationDate         string
	MessageId                   string
	CreationDateTime            string
	InstructingAgent            string
	InstructedAgent             string
	OriginalMessageId           string
	OriginalMessageNameId       string
	OriginalCreationDateTime    string
	OriginalEndToEndId          string
	OriginalTransactionId       string
	TransactionStatus           string
	AcceptanceDateTime          string
	InterBankSettlementAmount   string
	InterBankSettlementCurrency string
	InstructedAmount            string
	InstructedAmountCurrency    string
	RemittanceInformation       string
	DebtorName                  string
	DebtorAddress               string
	DebtorAccountNumber         string
	CreditorName                string
	CreditorAccountNumber       string
}

type ErrorResponse struct {
	Code        string `xml:"code"`
	Message     string `xml:"message"`
	Description string `xml:"description"`
}

type PaymentStatusResult struct {
	Success  bool
	Detail   *PaymentStatusDetail
	Messages []string
}

func ParsePaymentStatusSOAP(xmlData string) (*PaymentStatusResult, error) {
	// First try to parse as errorResponse directly (not wrapped in Envelope)
	// Check if the XML starts with errorResponse (not in Envelope)
	if strings.HasPrefix(strings.TrimSpace(xmlData), "<errorResponse") {
		var errorResp ErrorResponse
		if err := xml.Unmarshal([]byte(xmlData), &errorResp); err == nil && errorResp.Code != "" {
			errorMsg := "API Error"
			if errorResp.Code != "" {
				errorMsg += fmt.Sprintf(" [%s]", errorResp.Code)
			}
			if errorResp.Message != "" {
				errorMsg += fmt.Sprintf(": %s", errorResp.Message)
			}
			if errorResp.Description != "" {
				errorMsg += fmt.Sprintf(" - %s", errorResp.Description)
			}
			return &PaymentStatusResult{
				Success:  false,
				Messages: []string{errorMsg},
			}, nil
		}
	}

	// Try to parse as standard SOAP Envelope
	var env Envelope
	if err := xml.Unmarshal([]byte(xmlData), &env); err != nil {
		// If parsing fails, check if it's an errorResponse without proper structure
		if strings.Contains(xmlData, "errorResponse") {
			// Try to parse as direct errorResponse
			var errorResp ErrorResponse
			if err2 := xml.Unmarshal([]byte(xmlData), &errorResp); err2 == nil && errorResp.Code != "" {
				errorMsg := "API Error"
				if errorResp.Code != "" {
					errorMsg += fmt.Sprintf(" [%s]", errorResp.Code)
				}
				if errorResp.Message != "" {
					errorMsg += fmt.Sprintf(": %s", errorResp.Message)
				}
				if errorResp.Description != "" {
					errorMsg += fmt.Sprintf(" - %s", errorResp.Description)
				}
				return &PaymentStatusResult{
					Success:  false,
					Messages: []string{errorMsg},
				}, nil
			}
			return &PaymentStatusResult{
				Success:  false,
				Messages: []string{"API returned an error response"},
			}, nil
		}
		return nil, fmt.Errorf("failed to parse XML: %w", err)
	}

	// Check for error response in body (must check before PaymentStatusResponse)
	if env.Body.ErrorResponse != nil {
		errResp := env.Body.ErrorResponse
		errorMsg := "API Error"
		if errResp.Code != "" {
			errorMsg += fmt.Sprintf(" [%s]", errResp.Code)
		}
		if errResp.Message != "" {
			errorMsg += fmt.Sprintf(": %s", errResp.Message)
		}
		if errResp.Description != "" {
			errorMsg += fmt.Sprintf(" - %s", errResp.Description)
		}
		return &PaymentStatusResult{
			Success:  false,
			Messages: []string{errorMsg},
		}, nil
	}

	if env.Body.PaymentStatusResponse == nil {
		return &PaymentStatusResult{
			Success:  false,
			Messages: []string{"Invalid response type"},
		}, nil
	}

	output := env.Body.PaymentStatusResponse.Output

	// Check if Document and TxInfAndSts exist
	if output.Document.FIToFIPmtStsRpt.TxInfAndSts.OrgnlTxId == "" {
		return &PaymentStatusResult{
			Success:  false,
			Messages: []string{"Missing transaction information in response"},
		}, nil
	}

	txInfAndSts := output.Document.FIToFIPmtStsRpt.TxInfAndSts

	// Check transaction status - ACSC means accepted and settled
	// But we should still return the detail even if status is not ACSC
	transactionStatus := strings.ToUpper(txInfAndSts.TxSts)
	if transactionStatus == "" {
		return &PaymentStatusResult{
			Success:  false,
			Messages: []string{"Missing transaction status in response"},
		}, nil
	}

	detail := &PaymentStatusDetail{
		FromBankBIC:                 output.AppHdr.From.FIId.FinInstnId.Othr.Id,
		ToBankBIC:                   output.AppHdr.To.FIId.FinInstnId.Othr.Id,
		BizMessageIdentifier:        output.AppHdr.BizMessageIdentifier,
		MessageDefinitionIdentifier: output.AppHdr.MsgDefIdr,
		CreationDate:                output.AppHdr.CreDt,
		MessageId:                   output.Document.FIToFIPmtStsRpt.GrpHdr.MsgId,
		CreationDateTime:            output.Document.FIToFIPmtStsRpt.GrpHdr.CreDtTm,
		InstructingAgent:            output.Document.FIToFIPmtStsRpt.GrpHdr.InstgAgt.FinInstnId.Othr.Id,
		InstructedAgent:             output.Document.FIToFIPmtStsRpt.GrpHdr.InstdAgt.FinInstnId.Othr.Id,
		OriginalMessageId:           txInfAndSts.OrgnlGrpInf.OrgnlMsgId,
		OriginalMessageNameId:       txInfAndSts.OrgnlGrpInf.OrgnlMsgNmId,
		OriginalCreationDateTime:    txInfAndSts.OrgnlGrpInf.OrgnlCreDtTm,
		OriginalEndToEndId:          txInfAndSts.OrgnlEndToEndId,
		OriginalTransactionId:       txInfAndSts.OrgnlTxId,
		TransactionStatus:           txInfAndSts.TxSts,
		AcceptanceDateTime:          txInfAndSts.AccptncDtTm,
		InterBankSettlementAmount:   txInfAndSts.OrgnlTxRef.IntrBkSttlmAmt.Amount,
		InterBankSettlementCurrency: txInfAndSts.OrgnlTxRef.IntrBkSttlmAmt.Ccy,
		InstructedAmount:            txInfAndSts.OrgnlTxRef.Amt.InstdAmt.Amount,
		InstructedAmountCurrency:    txInfAndSts.OrgnlTxRef.Amt.InstdAmt.Ccy,
		RemittanceInformation:       txInfAndSts.OrgnlTxRef.RmtInf.Ustrd,
		DebtorName:                  txInfAndSts.OrgnlTxRef.Dbtr.Pty.Nm,
		DebtorAccountNumber:         txInfAndSts.OrgnlTxRef.DbtrAcct.Id.Othr.Id,
		CreditorName:                txInfAndSts.OrgnlTxRef.Cdtr.Pty.Nm,
		CreditorAccountNumber:       txInfAndSts.OrgnlTxRef.CdtrAcct.Id.Othr.Id,
	}

	// Extract related information if available
	if output.AppHdr.Rltd != nil {
		detail.RelatedFromBankBIC = output.AppHdr.Rltd.Fr.FIId.FinInstnId.Othr.Id
		detail.RelatedToBankBIC = output.AppHdr.Rltd.To.FIId.FinInstnId.Othr.Id
		detail.RelatedBizMessageIdentifier = output.AppHdr.Rltd.BizMessageIdentifier
		detail.RelatedMessageDefinitionId = output.AppHdr.Rltd.MsgDefIdr
		detail.RelatedCreationDate = output.AppHdr.Rltd.CreDt
	}

	// Extract debtor address if available
	if txInfAndSts.OrgnlTxRef.Dbtr.Pty.PstlAdr != nil {
		detail.DebtorAddress = txInfAndSts.OrgnlTxRef.Dbtr.Pty.PstlAdr.AdrLine
	}

	// Return success if status is ACSC, otherwise return with detail but Success=false
	success := transactionStatus == "ACSC"
	if !success {
		return &PaymentStatusResult{
			Success:  false,
			Detail:   detail,
			Messages: []string{fmt.Sprintf("Transaction status is not ACSC: %s", transactionStatus)},
		}, nil
	}

	return &PaymentStatusResult{
		Success: true,
		Detail:  detail,
	}, nil
}
