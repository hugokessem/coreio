package accountlookup

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Params struct {
	DebitBankBIC         string
	CreditBankBIC        string
	CreditAccountNumber  string
	BizMessageIdentifier string
	MessageIdentifier    string
	CreditDateTime       string
	CreditDate           string
}

func NewAccountLookup(param Params) string {
	return fmt.Sprintf(
		`<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:mb="http://MB_IPS" xmlns:urn="urn:iso:std:iso:20022:tech:xsd:head.001.001.03" xmlns:urn1="urn:iso:std:iso:20022:tech:xsd:acmt.023.001.03">
    <soapenv:Header/>
    <soapenv:Body>
        <mb:AccountVerfication>
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
                    <urn:MsgDefIdr>acmt.023.001.03</urn:MsgDefIdr>
                    <urn:CreDt>%s</urn:CreDt>
                </urn:AppHdr>
                <urn1:Document>
                    <urn1:IdVrfctnReq>
                        <urn1:Assgnmt>
                            <urn1:MsgId>%s</urn1:MsgId>
                            <urn1:CreDtTm>%s</urn1:CreDtTm>
                            <urn1:Assgnr>
                                <urn1:Agt>
                                    <urn1:FinInstnId>
                                        <urn1:Othr>
                                            <urn1:Id>%s</urn1:Id>
                                        </urn1:Othr>
                                    </urn1:FinInstnId>
                                </urn1:Agt>
                            </urn1:Assgnr>
                            <urn1:Assgne>
                                <urn1:Agt>
                                    <urn1:FinInstnId>
                                        <urn1:Othr>
                                            <urn1:Id>%s</urn1:Id>
                                        </urn1:Othr>
                                    </urn1:FinInstnId>
                                </urn1:Agt>
                            </urn1:Assgne>
                        </urn1:Assgnmt>
                        <urn1:Vrfctn>
                            <urn1:Id>%s</urn1:Id>
                            <urn1:PtyAndAcctId>
                                <urn1:Acct>
                                    <urn1:Id>
                                        <urn1:Othr>
                                            <urn1:Id>%s</urn1:Id>
                                            <urn1:SchmeNm>
                                                <urn1:Prtry>ACCT</urn1:Prtry>
                                            </urn1:SchmeNm>
                                        </urn1:Othr>
                                    </urn1:Id>
                                </urn1:Acct>
                            </urn1:PtyAndAcctId>
                        </urn1:Vrfctn>
                    </urn1:IdVrfctnReq>
                </urn1:Document>
            </input1>
        </mb:AccountVerfication>
    </soapenv:Body>
</soapenv:Envelope>`, param.DebitBankBIC, "FP", param.BizMessageIdentifier, param.CreditDate, param.MessageIdentifier, param.CreditDateTime, param.DebitBankBIC, param.CreditBankBIC, param.MessageIdentifier, param.CreditAccountNumber)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	AccountVerficationResponse AccountVerficationResponse `xml:"AccountVerficationResponse"`
}

type AccountVerficationResponse struct {
	Output Output `xml:"output1"`
}

type Output struct {
	AppHeader *AppHeader `xml:"AppHdr"`
	Document  *Document  `xml:"Document"`
}
type AppHeader struct {
	From struct {
		FIID struct {
			FinInstnId struct {
				Other struct {
					Identifier string `xml:"Id"`
				} `xml:"Other"`
			} `xml:"FinInstnId"`
		} `xml:"FIId"`
	} `xml:"Fr"`
	To struct {
		FIID struct {
			FinInstnId struct {
				Other struct {
					Identifier string `xml:"Id"`
				} `xml:"Other"`
			} `xml:"FinInstnId"`
		} `xml:"FIId"`
	} `xml:"To"`
	BizMessageIdentifier string `xml:"BizMsgIdr"`
	CreditDate           string `xml:"CreDt"`
}
type Document struct {
	IdVrfctnRpt *struct {
		Assignment *struct {
			MessageIdentifier string `xml:"MsgId"`
			CreditDateTime    string `xml:"CreDtTm"`
			Assgnr            *struct {
				Agt *struct {
					FinInstnId struct {
						Other struct {
							Identifier string `xml:"Id"`
						} `xml:"Other"`
					} `xml:"FinInstnId"`
				} `xml:"Agt"`
			} `xml:"Assgnr"`
			Assgne *struct {
				Agt *struct {
					FinInstnId struct {
						Other struct {
							Identifier string `xml:"Id"`
						} `xml:"Other"`
					} `xml:"FinInstnId"`
				} `xml:"Agt"`
			} `xml:"Assgne"`
		} `xml:"Assgnmt"`
		OrgnlAssgnmt *struct {
			MessageId      string `xml:"MsgId"`
			CreditDateTime string `xml:"CreDtTm"`
		} `xml:"OrgnlAssgnmt"`
		Rpt *struct {
			OriginalIdentifier string `xml:"OrgnlId"`
			Verification       string `xml:"Vrfctn"`
			Reason             struct {
				Prtry string `xml:"Prtry"`
			} `xml:"Rsn"`
			OrgnlPtyAndAcctId *struct {
				Acct *struct {
					Id *struct {
						Othr *struct {
							Id      string `xml:"Id"`
							SchmeNm *struct {
								Prtry string `xml:"Prtry"`
							} `xml:"SchmeNm"`
						} `xml:"Othr"`
					} `xml:"Id"`
				} `xml:"Acct"`
			} `xml:"OrgnlPtyAndAcctId"`
			UpdtdPtyAndAcctId *struct {
				Pty *struct {
					Nm string `xml:"Nm"`
				} `xml:"Pty"`
				Acct *struct {
					Id *struct {
						Othr *struct {
							Id      string `xml:"Id"`
							SchmeNm *struct {
								Prtry string `xml:"Prtry"`
							} `xml:"SchmeNm"`
						} `xml:"Othr"`
					} `xml:"Id"`
				} `xml:"Acct"`
			} `xml:"UpdtdPtyAndAcctId"`
		} `xml:"Rpt"`
	} `xml:"IdVrfctnRpt"`
}

type AccountVerficationDetail struct {
	CreditBankBIC           string
	OriginalIdentifier      string
	CreditAccountNumber     string
	CreditAccountHolderName string
}

type AccountLookupResult struct {
	Success  bool
	Detail   *AccountVerficationDetail
	Messages []string
}

func ParseAccountLookupSOAP(xmlData string) (*AccountLookupResult, error) {
	var env Envelope
	if err := xml.Unmarshal([]byte(xmlData), &env); err != nil {
		return nil, err
	}

	if env.Body.AccountVerficationResponse.Output.AppHeader != nil && env.Body.AccountVerficationResponse.Output.Document != nil {
		resp := env.Body.AccountVerficationResponse.Output

		// Check for nil pointers before accessing nested fields
		if resp.Document.IdVrfctnRpt == nil || resp.Document.IdVrfctnRpt.Rpt == nil {
			return &AccountLookupResult{
				Success:  false,
				Messages: []string{"Invalid Response: Missing report data"},
			}, nil
		}

		if strings.ToLower(resp.Document.IdVrfctnRpt.Rpt.Verification) != "true" {
			return &AccountLookupResult{
				Success:  false,
				Messages: []string{"Account Not Found!"},
			}, nil
		}

		// Safely extract account details with nil checks
		detail := &AccountVerficationDetail{
			OriginalIdentifier: resp.Document.IdVrfctnRpt.Rpt.OriginalIdentifier,
		}

		// Extract CreditBankBIC from AppHeader (To field contains the destination/credit bank)
		// CreditBankBIC should be the bank where the credit account is located
		if resp.AppHeader.To.FIID.FinInstnId.Other.Identifier != "" {
			detail.CreditBankBIC = resp.AppHeader.To.FIID.FinInstnId.Other.Identifier
		} else if resp.AppHeader.From.FIID.FinInstnId.Other.Identifier != "" {
			// Fallback to From field if To is empty
			detail.CreditBankBIC = resp.AppHeader.From.FIID.FinInstnId.Other.Identifier
		}

		// Extract account holder name
		if resp.Document.IdVrfctnRpt.Rpt.UpdtdPtyAndAcctId != nil &&
			resp.Document.IdVrfctnRpt.Rpt.UpdtdPtyAndAcctId.Pty != nil {
			detail.CreditAccountHolderName = resp.Document.IdVrfctnRpt.Rpt.UpdtdPtyAndAcctId.Pty.Nm
		}

		// Extract account number
		if resp.Document.IdVrfctnRpt.Rpt.OrgnlPtyAndAcctId != nil &&
			resp.Document.IdVrfctnRpt.Rpt.OrgnlPtyAndAcctId.Acct != nil &&
			resp.Document.IdVrfctnRpt.Rpt.OrgnlPtyAndAcctId.Acct.Id != nil &&
			resp.Document.IdVrfctnRpt.Rpt.OrgnlPtyAndAcctId.Acct.Id.Othr != nil {
			detail.CreditAccountNumber = resp.Document.IdVrfctnRpt.Rpt.OrgnlPtyAndAcctId.Acct.Id.Othr.Id
		}

		return &AccountLookupResult{
			Success: true,
			Detail:  detail,
		}, nil
	}

	return &AccountLookupResult{
		Success:  false,
		Messages: []string{"Invalid Response!"},
	}, nil

}
