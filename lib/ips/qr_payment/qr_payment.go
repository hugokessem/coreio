package qrpayment

import "fmt"

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
	`, param.DebitBankBIC, param.CreditBankBIC, param.BizMessageIdentifier, param.CreditDate, param.MessageIdentifier, param.CreditDateTime, param.DebitBankBIC, param.CreditBankBIC, param.EndToEndIdentifier, param.TransactionIdentifier, param.InterBankSettlementAmount, param.AccptanceDtatTime, param.InstructedAmount, param.DebiterInformation.Name, param.DebiterInformation.StreetName, param.DebiterInformation.BuildingNumber, param.DebiterInformation.PostalCode, param.DebiterInformation.TownName, param.DebiterInformation.Country, param.DebiterInformation.StreetName, param.DebitAccountNumber, param.DebitBankBIC, param.CreditBankBIC, param.CreditInformation.Name, param.CreditInformation.StreetName, param.CreditInformation.BuildingNumber, param.CreditInformation.PostalCode, param.CreditInformation.TownName, param.CreditInformation.Country, param.CreditInformation.AddressLine, param.CreditAccountNumber, param.Narative)
}
