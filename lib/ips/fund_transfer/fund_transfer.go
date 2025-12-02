package fundtransfer

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Params struct {
	DebitBankBIC              string
	CreditBankBIC             string
	BizMessageIdentifier      string
	CreditDate                string
	MessageIdentifier         string
	CreditDateTime            string
	EndToEndIdentifier        string
	TransactionIdentifier     string
	InterBankSettlementAmount string
	AccptanceDtatTime         string
	InstructedAmount          string
	DebitAccountNumber        string
	CreditAccountNumber       string
	CreditAccountHolderName   string
	Narative                  string
}

func NewFundTransfer(param Params) string {
	return fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:mb="http://MB_IPS" xmlns:urn="urn:iso:std:iso:20022:tech:xsd:head.001.001.03" xmlns:urn1="urn:iso:std:iso:20022:tech:xsd:pacs.008.001.10">
   <soapenv:Header/>
   <soapenv:Body>
      <mb:Payment>
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
                           <urn:Id>%s</urn:Id>
                        </urn:Othr>
                     </urn:FinInstnId>
                  </urn:FIId>
               </urn:To>
               <urn:BizMsgIdr>%s</urn:BizMsgIdr>
               <urn:MsgDefIdr>pacs.008.001.10</urn:MsgDefIdr>
               <urn:CreDt>%s</urn:CreDt>
            </urn:AppHdr>
            <urn1:Document>
               <urn1:FIToFICstmrCdtTrf>
                  <urn1:GrpHdr>
                     <urn1:MsgId>CBETETAA894480939908</urn1:MsgId>
                     <urn1:CreDtTm>%s</urn1:CreDtTm>
                     <urn1:NbOfTxs>1</urn1:NbOfTxs>
                     <urn1:SttlmInf>
                        <urn1:SttlmMtd>CLRG</urn1:SttlmMtd>
                        <urn1:ClrSys>
                           <urn1:Prtry>FP</urn1:Prtry>
                        </urn1:ClrSys>
                     </urn1:SttlmInf>
                     <urn1:PmtTpInf>
                        <urn1:LclInstrm>
                           <urn1:Prtry>CRTRM</urn1:Prtry>
                        </urn1:LclInstrm>
                     </urn1:PmtTpInf>
                     <urn1:InstgAgt>
                        <urn1:FinInstnId>
                           <urn1:Othr>
                              <urn1:Id>%s</urn1:Id>
                           </urn1:Othr>
                        </urn1:FinInstnId>
                     </urn1:InstgAgt>
                     <urn1:InstdAgt>
                        <urn1:FinInstnId>
                           <urn1:Othr>
                              <urn1:Id>%s</urn1:Id>
                           </urn1:Othr>
                        </urn1:FinInstnId>
                     </urn1:InstdAgt>
                  </urn1:GrpHdr>
                  <urn1:CdtTrfTxInf>
                     <urn1:PmtId>
                        <urn1:EndToEndId>%s</urn1:EndToEndId>
                        <urn1:TxId>%s</urn1:TxId>
                     </urn1:PmtId>
                     <urn1:IntrBkSttlmAmt Ccy="ETB">%s</urn1:IntrBkSttlmAmt>
                     <urn1:AccptncDtTm>%s</urn1:AccptncDtTm>
                     <urn1:InstdAmt Ccy="ETB">%s</urn1:InstdAmt>
                     <urn1:ChrgBr>SLEV</urn1:ChrgBr>
                     <urn1:Dbtr>
                        <urn1:Nm>MSCWT</urn1:Nm>
                        <urn1:PstlAdr>
                           <urn1:AdrLine>MOSCOW</urn1:AdrLine>
                        </urn1:PstlAdr>
                     </urn1:Dbtr>
                     <urn1:DbtrAcct>
                        <urn1:Id>
                           <urn1:Othr>
                              <urn1:Id>%s</urn1:Id>
                              <urn1:SchmeNm>
                                 <urn1:Prtry>ACCT</urn1:Prtry>
                              </urn1:SchmeNm>
                              <urn1:Issr>C</urn1:Issr>
                           </urn1:Othr>
                        </urn1:Id>
                     </urn1:DbtrAcct>
                     <urn1:DbtrAgt>
                        <urn1:FinInstnId>
                           <urn1:Othr>
                              <urn1:Id>%s</urn1:Id>
                              <urn1:Issr>ATM</urn1:Issr>
                           </urn1:Othr>
                        </urn1:FinInstnId>
                     </urn1:DbtrAgt>
                     <urn1:CdtrAgt>
                        <urn1:FinInstnId>
                           <urn1:Othr>
                              <urn1:Id>%s</urn1:Id>
                           </urn1:Othr>
                        </urn1:FinInstnId>
                     </urn1:CdtrAgt>
                     <urn1:Cdtr>
                        <urn1:Nm>%s</urn1:Nm>
                     </urn1:Cdtr>
                     <urn1:CdtrAcct>
                        <urn1:Id>
                           <urn1:Othr>
                              <urn1:Id>%s</urn1:Id>
                              <urn1:SchmeNm>
                                 <urn1:Prtry>ACCT</urn1:Prtry>
                              </urn1:SchmeNm>
                           </urn1:Othr>
                        </urn1:Id>
                     </urn1:CdtrAcct>
                     <urn1:RmtInf>
                        <urn1:Ustrd>%s</urn1:Ustrd>
                     </urn1:RmtInf>
                  </urn1:CdtTrfTxInf>
               </urn1:FIToFICstmrCdtTrf>
            </urn1:Document>
         </input1>
      </mb:Payment>
   </soapenv:Body>
</soapenv:Envelope>
	`, param.DebitBankBIC, param.CreditBankBIC, param.BizMessageIdentifier, param.CreditDate, param.CreditDateTime, param.DebitBankBIC, param.CreditBankBIC, param.EndToEndIdentifier, param.TransactionIdentifier, param.InterBankSettlementAmount, param.AccptanceDtatTime, param.InstructedAmount, param.DebitAccountNumber, param.DebitBankBIC, param.CreditBankBIC, param.CreditAccountHolderName, param.CreditAccountNumber, param.Narative)
}

type Envelop struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	PaymentResponse *PaymentResponse `xml:"PaymentResponse"`
}

type PaymentResponse struct {
	Output Output `xml:"output1"`
}

type Output struct {
	AppHdr   AppHeader `xml:"AppHdr"`
	Document Document  `xml:"Document"`
}
type AppHeader struct {
	From struct {
		FIID struct {
			FinInstnId struct {
				Other struct {
					Identifier string `xml:"Id"`
				} `xml:"Other"`
			} `xml:"FinInstnId"`
		} `xml:"FFID"`
	} `xml:"Fr"`
	To struct {
		FIID struct {
			FinInstnId struct {
				Other struct {
					Identifier string `xml:"Id"`
				} `xml:"Other"`
			} `xml:"FinInstnId"`
		} `xml:"FFID"`
	} `xml:"To"`
	BizMessageIdentifier string `xml:"BizMsgIdr"`
	CreditDate           string `xml:"CreDt"`
	Rltd                 *struct {
		Fr *struct {
			FIId *struct {
				FinInstnId *struct {
					Othr *struct {
						Id string `xml:"Id"`
					} `xml:"Othr"`
				} `xml:"FinInstnId"`
			} `xml:"FIId"`
		} `xml:"Fr"`
		To *struct {
			FIId *struct {
				FinInstnId *struct {
					Othr *struct {
						Id string `xml:"Id"`
					} `xml:"Othr"`
				} `xml:"FinInstnId"`
			} `xml:"FIId"`
		} `xml:"To"`
		BizMessageIdentifier string `xml:"BizMsgIdr"`
		CreditDate           string `xml:"CreDt"`
	} `xml:"Rltd"`
}
type Document struct {
	FIToFIPmtStsRpt *struct {
		GrpHdr *struct {
			MsgId          string `xml:"MsgId"`
			CreditDateTime string `xml:"CreDtTm"`
			InstgAgt       *struct {
				FinInstnId *struct {
					Othr *struct {
						Id string `xml:"Id"`
					} `xml:"Othr"`
				} `xml:"FinInstnId"`
				InstdAgt *struct {
					Othr *struct {
						Id string `xml:"Id"`
					} `xml:"Othr"`
				} `xml:"InstdAgt"`
			} `xml:"InstgAgt"`
		} `xml:"GrpHdr"`
		TxInfAndSts *struct {
			OrgnlGrpInf *struct {
				OriginalMessageIdentifier string `xml:"OrgnlMsgId"`
				OriginalCreditDateTime    string `xml:"OrgnlCreDtTm"`
			} `xml:"OrgnlGrpInf"`
			OriginalEndtoEndIdentifier    string `xml:"OrgnlEndToEndId"`
			OriginalTransactionIdentifier string `xml:"OrgnlTxId"`
			TransactionStatus             string `xml:"TxSts"`
			AccptanceDtatTime             string `xml:"AccptncDtTm"`
			OrgnlTransactionRefrence      *struct {
				InterBankSettlementAmount string `xml:"IntrBkSttlmAmt"`
				Amt                       *struct {
					InstructedAmount string `xml:"InstdAmt"`
				} `xml:"Amt"`
				PmtTpInf *struct {
					CtgyPurp *struct {
						Prtry string `xml:"Prtry"`
					} `xml:"CtgyPurp"`
				} `xml:"PmtTpInf"`
				RmtInf *struct {
					Ustrd string `xml:"Ustrd"`
				} `xml:"RmtInf"`
				Dbtr *struct {
					Pty *struct {
						Nm      string `xml:"Nm"`
						PstlAdr *struct {
							AdrLine string `xml:"AdrLine"`
						} `xml:"PstlAdr"`
					} `xml:"Pty"`
				} `xml:"Dbtr"`
				DbtrAcct *struct {
					Id *struct {
						Othr *struct {
							Id      string `xml:"Id"`
							SchmeNm *struct {
								Prtry string `xml:"Prtry"`
							} `xml:"SchmeNm"`
							Issr string `xml:"Issr"`
						} `xml:"Othr"`
					} `xml:"Id"`
				} `xml:"DbtrAcct"`
				Cdtr *struct {
					Pty *struct {
						Nm string `xml:"Nm"`
					} `xml:"Pty"`
				} `xml:"Cdtr"`
				CdtrAcct *struct {
					Id *struct {
						Othr *struct {
							Id      string `xml:"Id"`
							SchmeNm *struct {
								Prtry string `xml:"Prtry"`
							} `xml:"SchmeNm"`
						} `xml:"Othr"`
					} `xml:"Id"`
				} `xml:"CdtrAcct"`
			} `xml:"OrgnlTxRef"`
		} `xml:"TxInfAndSts"`
	} `xml:"FIToFIPmtStsRpt"`
}

type PaymentResponseDetail struct {
	CreditBankBIC                 string
	DebitBankBIC                  string
	CreditAccountNumber           string
	CreditAccountHolderName       string
	DebitAccountNumber            string
	DebitAccountHolderName        string
	CreditDate                    string
	CreditDateTime                string
	EndToEndIdentifier            string
	TransactionIdentifier         string
	InterBankSettlementAmount     string
	AccptanceDtatTime             string
	InstructedAmount              string
	Narative                      string
	OriginalTransactionIdentifier string
	OriginalEndtoEndIdentifier    string
	OriginalCreditDateTime        string
	OriginalMessageIdentifier     string
}

type FundTransferResult struct {
	Success  bool
	Detail   *PaymentResponseDetail
	Messages []string
}

func ParseFundTransferSOAP(xmlDate string) (*FundTransferResult, error) {
	var env Envelop

	if err := xml.Unmarshal([]byte(xmlDate), &env); err != nil {
		return nil, err
	}

	if env.Body.PaymentResponse != nil {
		resp := env.Body.PaymentResponse.Output
		if strings.ToLower(resp.Document.FIToFIPmtStsRpt.TxInfAndSts.TransactionStatus) != "acsc" {
			return &FundTransferResult{
				Success:  false,
				Messages: []string{"Failed to intiate transaction!"},
			}, nil
		}

		return &FundTransferResult{
			Success: true,
			Detail: &PaymentResponseDetail{
				OriginalTransactionIdentifier: resp.Document.FIToFIPmtStsRpt.TxInfAndSts.OriginalTransactionIdentifier,
				OriginalEndtoEndIdentifier:    resp.Document.FIToFIPmtStsRpt.TxInfAndSts.OriginalEndtoEndIdentifier,
				OriginalCreditDateTime:        resp.Document.FIToFIPmtStsRpt.TxInfAndSts.OrgnlGrpInf.OriginalCreditDateTime,
				OriginalMessageIdentifier:     resp.Document.FIToFIPmtStsRpt.TxInfAndSts.OrgnlGrpInf.OriginalMessageIdentifier,
				InterBankSettlementAmount:     resp.Document.FIToFIPmtStsRpt.TxInfAndSts.OrgnlTransactionRefrence.InterBankSettlementAmount,
				AccptanceDtatTime:             resp.Document.FIToFIPmtStsRpt.TxInfAndSts.AccptanceDtatTime,
				InstructedAmount:              resp.Document.FIToFIPmtStsRpt.TxInfAndSts.OrgnlTransactionRefrence.InterBankSettlementAmount,
				DebitAccountNumber:            resp.Document.FIToFIPmtStsRpt.TxInfAndSts.OrgnlTransactionRefrence.Dbtr.Pty.Nm,
				DebitAccountHolderName:        resp.Document.FIToFIPmtStsRpt.TxInfAndSts.OrgnlTransactionRefrence.Dbtr.Pty.Nm,
				CreditBankBIC:                 resp.AppHdr.To.FIID.FinInstnId.Other.Identifier,
				DebitBankBIC:                  resp.AppHdr.From.FIID.FinInstnId.Other.Identifier,
				CreditDate:                    resp.AppHdr.CreditDate,
				CreditDateTime:                resp.AppHdr.Rltd.CreditDate,
				EndToEndIdentifier:            resp.Document.FIToFIPmtStsRpt.TxInfAndSts.OriginalEndtoEndIdentifier,
				CreditAccountNumber:           resp.Document.FIToFIPmtStsRpt.TxInfAndSts.OrgnlTransactionRefrence.CdtrAcct.Id.Othr.Id,
				CreditAccountHolderName:       resp.Document.FIToFIPmtStsRpt.TxInfAndSts.OrgnlTransactionRefrence.Cdtr.Pty.Nm,
				Narative:                      resp.Document.FIToFIPmtStsRpt.TxInfAndSts.OrgnlTransactionRefrence.RmtInf.Ustrd,
			},
		}, nil

	}

	return &FundTransferResult{
		Success:  false,
		Messages: []string{"Invalid Response!"},
	}, nil
}
