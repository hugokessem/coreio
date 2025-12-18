package qrpayment

import (
	"encoding/xml"
	"fmt"
)

type DebiterInformation struct {
	Name           string
	StreetName     string
	BuildingNumber string
	PostalCode     string
	TownName       string
	Country        string
}

type CrediterInformation struct {
	Name           string
	StreetName     string
	BuildingNumber string
	PostalCode     string
	City           string
	TownName       string
	Country        string
	AddressLine    string
}

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
	DebiterInformation        DebiterInformation
	CreditInformation         CrediterInformation
}

func NewQrPayment(param Params) string {
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
                     <urn1:MsgId>%s</urn1:MsgId>
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
                        <urn1:CtgyPurp>
                           <urn1:Prtry>C2BSQR</urn1:Prtry>
                        </urn1:CtgyPurp>
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
                     <urn1:UltmtDbtr>
                        <urn1:Nm>%s</urn1:Nm>
                        <urn1:PstlAdr>
                           <urn1:StrtNm>%s</urn1:StrtNm>
                           <urn1:BldgNb>%s</urn1:BldgNb>
                           <urn1:PstCd>%s</urn1:PstCd>
                           <urn1:TwnNm>%s</urn1:TwnNm>
                           <urn1:Ctry>%s</urn1:Ctry>
                        </urn1:PstlAdr>
                     </urn1:UltmtDbtr>
                     <urn1:Dbtr>
                        <urn1:Nm>MSCWT</urn1:Nm>
                        <urn1:PstlAdr>
                           <urn1:StrtNm>%s</urn1:StrtNm>
                        </urn1:PstlAdr>
                        <urn1:Id>
                           <urn1:PrvtId>
                              <urn1:Othr>
                                 <urn1:Id>MOBN</urn1:Id>
                                 <urn1:SchmeNm>
                                    <urn1:Prtry>LPNB</urn1:Prtry>
                                 </urn1:SchmeNm>
                              </urn1:Othr>
                           </urn1:PrvtId>
                        </urn1:Id>
                        <urn1:CtctDtls></urn1:CtctDtls>
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
                        <urn1:PstlAdr>
                           <urn1:StrtNm>%s</urn1:StrtNm>
                           <urn1:BldgNb>%s</urn1:BldgNb>
                           <urn1:PstCd>%s</urn1:PstCd>
                           <urn1:TwnNm>%s</urn1:TwnNm>
                           <urn1:Ctry>%s</urn1:Ctry>
                           <urn1:AdrLine>%s</urn1:AdrLine>
                        </urn1:PstlAdr>
                        <urn1:CtryOfRes>ET</urn1:CtryOfRes>
                        <urn1:CtctDtls>
                           <urn1:Nm>%s</urn1:Nm>
                           <urn1:Dept>%s</urn1:Dept>
                           <urn1:Othr>
                              <urn1:ChanlTp>QRCP</urn1:ChanlTp>
                              <urn1:Id>QRCPS</urn1:Id>
                           </urn1:Othr>
                        </urn1:CtctDtls>
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
                     <urn1:UltmtCdtr>
                        <urn1:Id>
                           <urn1:PrvtId>
                              <urn1:Othr>
                                 <urn1:Id>MOBN</urn1:Id>
                                 <urn1:SchmeNm>
                                    <urn1:Prtry>MOBN</urn1:Prtry>
                                 </urn1:SchmeNm>
                              </urn1:Othr>
                           </urn1:PrvtId>
                        </urn1:Id>
                     </urn1:UltmtCdtr>
                     <urn1:Purp>
                        <urn1:Prtry>GOVTP</urn1:Prtry>
                     </urn1:Purp>
                     <urn1:Tax>
                        <urn1:Cdtr>
                           <urn1:TaxId>123456789</urn1:TaxId>
                        </urn1:Cdtr>
                     </urn1:Tax>
                     <urn1:RmtInf>
                        <urn1:Ustrd>%s</urn1:Ustrd>
                        <urn1:Strd>
                           <urn1:RfrdDocInf>
                              <urn1:Tp>
                                 <urn1:CdOrPrtry>
                                    <urn1:Prtry>%s</urn1:Prtry>
                                 </urn1:CdOrPrtry>
                              </urn1:Tp>
                           </urn1:RfrdDocInf>
                        </urn1:Strd>
                     </urn1:RmtInf>
                  </urn1:CdtTrfTxInf>
               </urn1:FIToFICstmrCdtTrf>
            </urn1:Document>
         </input1>
      </mb:Payment>
   </soapenv:Body>
</soapenv:Envelope>
	`, param.DebitBankBIC, param.CreditBankBIC, param.BizMessageIdentifier, param.CreditDate, param.MessageIdentifier, param.CreditDateTime, param.DebitBankBIC, param.CreditBankBIC, param.EndToEndIdentifier, param.TransactionIdentifier, param.InterBankSettlementAmount, param.AccptanceDtatTime, param.InstructedAmount, param.DebiterInformation.Name, param.DebiterInformation.StreetName, param.DebiterInformation.BuildingNumber, param.DebiterInformation.PostalCode, param.DebiterInformation.TownName, param.DebiterInformation.Country, param.DebiterInformation.StreetName, param.DebitAccountNumber, param.DebitBankBIC, param.CreditBankBIC, param.CreditInformation.Name, param.CreditInformation.StreetName, param.CreditInformation.BuildingNumber, param.CreditInformation.PostalCode, param.CreditInformation.TownName, param.CreditInformation.Country, param.CreditInformation.AddressLine, param.CreditAccountHolderName, "", param.CreditAccountNumber, param.Narative, param.Narative)
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
		MessageDefinitionId  string `xml:"MsgDefIdr"`
		CreditDate           string `xml:"CreDt"`
	} `xml:"Rltd"`
}

type Document struct {
	FIToFIPmtStsRpt *struct {
		GrpHdr *struct {
			MsgId    string `xml:"MsgId"`
			CreDtTm  string `xml:"CreDtTm"`
			InstgAgt *struct {
				FinInstnId *struct {
					Othr *struct {
						Id string `xml:"Id"`
					} `xml:"Othr"`
				} `xml:"FinInstnId"`
			} `xml:"InstgAgt"`
			InstdAgt *struct {
				FinInstnId *struct {
					Othr *struct {
						Id string `xml:"Id"`
					} `xml:"Othr"`
				} `xml:"FinInstnId"`
			} `xml:"InstdAgt"`
		} `xml:"GrpHdr"`
		OrgnlGrpInfAndSts *struct {
			OriginalMessageIdentifier string `xml:"OrgnlMsgId"`
			OriginalMessageNameId     string `xml:"OrgnlMsgNmId"`
			OriginalCreditDateTime    string `xml:"OrgnlCreDtTm"`
		} `xml:"OrgnlGrpInfAndSts"`
		TxInfAndSts *struct {
			OriginalEndToEndIdentifier    string `xml:"OrgnlEndToEndId"`
			OriginalTransactionIdentifier string `xml:"OrgnlTxId"`
			TransactionStatus             string `xml:"TxSts"`
			AcceptanceDateTime            string `xml:"AccptncDtTm"`
			StsRsnInf                     *struct {
				Orgtr *struct {
					Id *struct {
						OrgId *struct {
							Othr *struct {
								Id string `xml:"Id"`
							} `xml:"Othr"`
						} `xml:"OrgId"`
					} `xml:"Id"`
				} `xml:"Orgtr"`
				Reason struct {
					Proprietary string `xml:"Prtry"`
				} `xml:"Rsn"`
				AdditionalInformation string `xml:"AddtlInf"`
			} `xml:"StsRsnInf"`
			OrgnlTxRef *struct {
				InterBankSettlementAmount struct {
					Amount string `xml:",chardata"`
					Ccy    string `xml:"Ccy,attr"`
				} `xml:"IntrBkSttlmAmt"`
				Amt *struct {
					InstructedAmount struct {
						Amount string `xml:",chardata"`
						Ccy    string `xml:"Ccy,attr"`
					} `xml:"InstdAmt"`
				} `xml:"Amt"`
				PmtTpInf *struct {
					CtgyPurp *struct {
						Prtry string `xml:"Prtry"`
					} `xml:"CtgyPurp"`
				} `xml:"PmtTpInf"`
				RmtInf *struct {
					Ustrd string `xml:"Ustrd"`
					Strd  *struct {
						RfrdDocInf *struct {
							Tp *struct {
								CdOrPrtry *struct {
									Prtry string `xml:"Prtry"`
								} `xml:"CdOrPrtry"`
							} `xml:"Tp"`
						} `xml:"RfrdDocInf"`
					} `xml:"Strd"`
				} `xml:"RmtInf"`
				Dbtr *struct {
					Pty *struct {
						Name    string `xml:"Nm"`
						PstlAdr *struct {
							StreetName string `xml:"StrtNm"`
						} `xml:"PstlAdr"`
						Id *struct {
							PrvtId *struct {
								Othr *struct {
									Id      string `xml:"Id"`
									SchmeNm *struct {
										Prtry string `xml:"Prtry"`
									} `xml:"SchmeNm"`
								} `xml:"Othr"`
							} `xml:"PrvtId"`
						} `xml:"Id"`
						CtctDtls string `xml:"CtctDtls"`
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
						Name    string `xml:"Nm"`
						PstlAdr *struct {
							StreetName     string `xml:"StrtNm"`
							BuildingNumber string `xml:"BldgNb"`
							PostalCode     string `xml:"PstCd"`
							TownName       string `xml:"TwnNm"`
							Country        string `xml:"Ctry"`
							AddressLine    string `xml:"AdrLine"`
						} `xml:"PstlAdr"`
						CtryOfRes string `xml:"CtryOfRes"`
						CtctDtls  *struct {
							Name       string `xml:"Nm"`
							Department string `xml:"Dept"`
							Othr       *struct {
								ChannelType string `xml:"ChanlTp"`
								Identifier  string `xml:"Id"`
							} `xml:"Othr"`
						} `xml:"CtctDtls"`
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
				UltmtCdtr *struct {
					Pty *struct {
						Id *struct {
							PrvtId *struct {
								Othr *struct {
									Id      string `xml:"Id"`
									SchmeNm *struct {
										Prtry string `xml:"Prtry"`
									} `xml:"SchmeNm"`
								} `xml:"Othr"`
							} `xml:"PrvtId"`
						} `xml:"Id"`
					} `xml:"Pty"`
				} `xml:"UltmtCdtr"`
				Purp *struct {
					Prtry string `xml:"Prtry"`
				} `xml:"Purp"`
			} `xml:"OrgnlTxRef"`
		} `xml:"TxInfAndSts"`
	} `xml:"FIToFIPmtStsRpt"`
}

type QrPaymentResult struct {
	Success  bool
	Detail   *QrPaymentDetail
	Messages []string
}

type QrPaymentDetail struct {
	FromBankBIC                   string
	ToBankBIC                     string
	BizMessageIdentifier          string
	CreditDate                    string
	RelatedFromBankBIC            string
	RelatedToBankBIC              string
	RelatedBizMessageIdentifier   string
	RelatedMessageDefinitionId    string
	RelatedCreditDate             string
	MessageId                     string
	CreationDateTime              string
	InstructingAgent              string
	InstructedAgent               string
	OriginalMessageIdentifier     string
	OriginalMessageNameId         string
	OriginalCreditDateTime        string
	OriginalEndToEndIdentifier    string
	OriginalTransactionIdentifier string
	TransactionStatus             string
	AcceptanceDateTime            string
	StatusReason                  string
	StatusReasonAdditionalInfo    string
	InterBankSettlementAmount     string
	InterBankSettlementCurrency   string
	InstructedAmount              string
	InstructedAmountCurrency      string
	PaymentTypeCategoryPurpose    string
	RemittanceInformation         string
	RemittanceReference           string
	DebtorName                    string
	DebtorStreetName              string
	DebtorAccountNumber           string
	CreditorName                  string
	CreditorStreetName            string
	CreditorBuildingNumber        string
	CreditorPostalCode            string
	CreditorTownName              string
	CreditorCountry               string
	CreditorAddressLine           string
	CreditorCountryOfResidence    string
	CreditorContactName           string
	CreditorContactDepartment     string
	CreditorAccountNumber         string
	Purpose                       string
}

func ParseQrPaymentSOAP(xmlData string) (*QrPaymentResult, error) {
	var env Envelop

	if err := xml.Unmarshal([]byte(xmlData), &env); err != nil {
		return nil, err
	}

	if env.Body.PaymentResponse == nil {
		return &QrPaymentResult{
			Success:  false,
			Messages: []string{"Invalid Response: Missing PaymentResponse"},
		}, nil
	}

	resp := env.Body.PaymentResponse.Output

	if resp.Document.FIToFIPmtStsRpt == nil || resp.Document.FIToFIPmtStsRpt.TxInfAndSts == nil {
		return &QrPaymentResult{
			Success:  false,
			Messages: []string{"Invalid Response: Missing transaction data"},
		}, nil
	}

	txInf := resp.Document.FIToFIPmtStsRpt.TxInfAndSts
	grpHdr := resp.Document.FIToFIPmtStsRpt.GrpHdr
	orgnlGrpInf := resp.Document.FIToFIPmtStsRpt.OrgnlGrpInfAndSts
	orgnlTxRef := txInf.OrgnlTxRef

	// Check transaction status
	if txInf.TransactionStatus != "ACSC" {
		statusMsg := "Transaction status is not ACSC"
		if txInf.StsRsnInf != nil {
			if txInf.StsRsnInf.AdditionalInformation != "" {
				statusMsg = txInf.StsRsnInf.AdditionalInformation
			} else if txInf.StsRsnInf.Reason.Proprietary != "" {
				statusMsg = txInf.StsRsnInf.Reason.Proprietary
			}
		}

		detail := &QrPaymentDetail{}
		if grpHdr != nil {
			if grpHdr.InstgAgt != nil {
				detail.InstructingAgent = grpHdr.InstgAgt.FinInstnId.Othr.Id
			}
			if grpHdr.InstdAgt != nil {
				detail.InstructedAgent = grpHdr.InstdAgt.FinInstnId.Othr.Id
			}
			detail.MessageId = grpHdr.MsgId
			detail.CreationDateTime = grpHdr.CreDtTm
		}

		if orgnlGrpInf != nil {
			detail.OriginalMessageIdentifier = orgnlGrpInf.OriginalMessageIdentifier
			detail.OriginalMessageNameId = orgnlGrpInf.OriginalMessageNameId
			detail.OriginalCreditDateTime = orgnlGrpInf.OriginalCreditDateTime
		}

		detail.OriginalEndToEndIdentifier = txInf.OriginalEndToEndIdentifier
		detail.OriginalTransactionIdentifier = txInf.OriginalTransactionIdentifier
		detail.TransactionStatus = txInf.TransactionStatus
		detail.AcceptanceDateTime = txInf.AcceptanceDateTime

		if txInf.StsRsnInf != nil {
			detail.StatusReason = txInf.StsRsnInf.Reason.Proprietary
			detail.StatusReasonAdditionalInfo = txInf.StsRsnInf.AdditionalInformation
		}

		return &QrPaymentResult{
			Success:  false,
			Detail:   detail,
			Messages: []string{statusMsg},
		}, nil
	}

	// Build detail for successful transaction
	detail := &QrPaymentDetail{
		FromBankBIC:                   resp.AppHdr.From.FIId.FinInstnId.Othr.Id,
		ToBankBIC:                     resp.AppHdr.To.FIId.FinInstnId.Othr.Id,
		BizMessageIdentifier:          resp.AppHdr.BizMessageIdentifier,
		CreditDate:                    resp.AppHdr.CreditDate,
		OriginalEndToEndIdentifier:    txInf.OriginalEndToEndIdentifier,
		OriginalTransactionIdentifier: txInf.OriginalTransactionIdentifier,
		TransactionStatus:             txInf.TransactionStatus,
		AcceptanceDateTime:            txInf.AcceptanceDateTime,
	}

	if resp.AppHdr.Rltd != nil {
		detail.RelatedFromBankBIC = resp.AppHdr.Rltd.Fr.FIId.FinInstnId.Othr.Id
		detail.RelatedToBankBIC = resp.AppHdr.Rltd.To.FIId.FinInstnId.Othr.Id
		detail.RelatedBizMessageIdentifier = resp.AppHdr.Rltd.BizMessageIdentifier
		detail.RelatedMessageDefinitionId = resp.AppHdr.Rltd.MessageDefinitionId
		detail.RelatedCreditDate = resp.AppHdr.Rltd.CreditDate
	}

	if grpHdr != nil {
		detail.MessageId = grpHdr.MsgId
		detail.CreationDateTime = grpHdr.CreDtTm
		if grpHdr.InstgAgt != nil {
			detail.InstructingAgent = grpHdr.InstgAgt.FinInstnId.Othr.Id
		}
		if grpHdr.InstdAgt != nil {
			detail.InstructedAgent = grpHdr.InstdAgt.FinInstnId.Othr.Id
		}
	}

	if orgnlGrpInf != nil {
		detail.OriginalMessageIdentifier = orgnlGrpInf.OriginalMessageIdentifier
		detail.OriginalMessageNameId = orgnlGrpInf.OriginalMessageNameId
		detail.OriginalCreditDateTime = orgnlGrpInf.OriginalCreditDateTime
	}

	if orgnlTxRef != nil {
		detail.InterBankSettlementAmount = orgnlTxRef.InterBankSettlementAmount.Amount
		detail.InterBankSettlementCurrency = orgnlTxRef.InterBankSettlementAmount.Ccy
		if orgnlTxRef.Amt != nil {
			detail.InstructedAmount = orgnlTxRef.Amt.InstructedAmount.Amount
			detail.InstructedAmountCurrency = orgnlTxRef.Amt.InstructedAmount.Ccy
		}
		if orgnlTxRef.PmtTpInf != nil && orgnlTxRef.PmtTpInf.CtgyPurp != nil {
			detail.PaymentTypeCategoryPurpose = orgnlTxRef.PmtTpInf.CtgyPurp.Prtry
		}
		if orgnlTxRef.RmtInf != nil {
			detail.RemittanceInformation = orgnlTxRef.RmtInf.Ustrd
			if orgnlTxRef.RmtInf.Strd != nil && orgnlTxRef.RmtInf.Strd.RfrdDocInf != nil {
				detail.RemittanceReference = orgnlTxRef.RmtInf.Strd.RfrdDocInf.Tp.CdOrPrtry.Prtry
			}
		}
		if orgnlTxRef.Dbtr != nil && orgnlTxRef.Dbtr.Pty != nil {
			detail.DebtorName = orgnlTxRef.Dbtr.Pty.Name
			if orgnlTxRef.Dbtr.Pty.PstlAdr != nil {
				detail.DebtorStreetName = orgnlTxRef.Dbtr.Pty.PstlAdr.StreetName
			}
		}
		if orgnlTxRef.DbtrAcct != nil && orgnlTxRef.DbtrAcct.Id != nil {
			detail.DebtorAccountNumber = orgnlTxRef.DbtrAcct.Id.Othr.Id
		}
		if orgnlTxRef.Cdtr != nil && orgnlTxRef.Cdtr.Pty != nil {
			detail.CreditorName = orgnlTxRef.Cdtr.Pty.Name
			if orgnlTxRef.Cdtr.Pty.PstlAdr != nil {
				detail.CreditorStreetName = orgnlTxRef.Cdtr.Pty.PstlAdr.StreetName
				detail.CreditorBuildingNumber = orgnlTxRef.Cdtr.Pty.PstlAdr.BuildingNumber
				detail.CreditorPostalCode = orgnlTxRef.Cdtr.Pty.PstlAdr.PostalCode
				detail.CreditorTownName = orgnlTxRef.Cdtr.Pty.PstlAdr.TownName
				detail.CreditorCountry = orgnlTxRef.Cdtr.Pty.PstlAdr.Country
				detail.CreditorAddressLine = orgnlTxRef.Cdtr.Pty.PstlAdr.AddressLine
			}
			detail.CreditorCountryOfResidence = orgnlTxRef.Cdtr.Pty.CtryOfRes
			if orgnlTxRef.Cdtr.Pty.CtctDtls != nil {
				detail.CreditorContactName = orgnlTxRef.Cdtr.Pty.CtctDtls.Name
				detail.CreditorContactDepartment = orgnlTxRef.Cdtr.Pty.CtctDtls.Department
			}
		}
		if orgnlTxRef.CdtrAcct != nil && orgnlTxRef.CdtrAcct.Id != nil {
			detail.CreditorAccountNumber = orgnlTxRef.CdtrAcct.Id.Othr.Id
		}
		if orgnlTxRef.Purp != nil {
			detail.Purpose = orgnlTxRef.Purp.Prtry
		}
	}

	return &QrPaymentResult{
		Success: true,
		Detail:  detail,
	}, nil
}
