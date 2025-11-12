package accountlookup

import "fmt"

/*
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:mb="http://MB_IPS" xmlns:urn="urn:iso:std:iso:20022:tech:xsd:head.001.001.03" xmlns:urn1="urn:iso:std:iso:20022:tech:xsd:acmt.023.001.03">

	<soapenv:Header/>
	<soapenv:Body>
	   <mb:AccountVerfication>
	      <input1>
	         <urn:AppHdr>
	            <urn:Fr>
	               <urn:FIId>
	                  <urn:FinInstnId>
	                     <urn:Othr>
	                        <urn:Id>CBETETAA</urn:Id>
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
	            <urn:BizMsgIdr>CBETETAA843572771</urn:BizMsgIdr>
	            <urn:MsgDefIdr>acmt.023.001.03</urn:MsgDefIdr>
	            <urn:CreDt>2023-06-24T00:00:00.000Z</urn:CreDt>
	         </urn:AppHdr>
	         <urn1:Document>
	            <urn1:IdVrfctnReq>
	               <urn1:Assgnmt>
	                  <urn1:MsgId>CBETETAA843572771</urn1:MsgId>
	                  <urn1:CreDtTm>2023-06-24T00:00:00.000+03:00</urn1:CreDtTm>
	                  <urn1:Assgnr>
	                     <urn1:Agt>
	                        <urn1:FinInstnId>
	                           <urn1:Othr>
	                              <urn1:Id>CBETETAA</urn1:Id>
	                           </urn1:Othr>
	                        </urn1:FinInstnId>
	                     </urn1:Agt>
	                  </urn1:Assgnr>
	                  <urn1:Assgne>
	                     <urn1:Agt>
	                        <urn1:FinInstnId>
	                           <urn1:Othr>
	                              <urn1:Id>ETSETAA</urn1:Id>
	                           </urn1:Othr>
	                        </urn1:FinInstnId>
	                     </urn1:Agt>
	                  </urn1:Assgne>
	               </urn1:Assgnmt>
	               <urn1:Vrfctn>
	                  <urn1:Id>CBETETAA843572771</urn1:Id>
	                  <urn1:PtyAndAcctId>
	                     <urn1:Acct>
	                        <urn1:Id>
	                           <urn1:Othr>
	                              <urn1:Id>1234567890</urn1:Id>
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

</soapenv:Envelope>
*/
type Params struct {
	AccountNumber               string
	DebitBankBIC                string
	CreaditBankBIC              string
	CreditAccountNumber         string
	BizMessageIdentifier        string
	MessageDefinitionIdentifier string
	MessageIdentifier           string
	CreditDateTime              string
	CreditDate                  string
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
                     <urn1:CreDtTm>2023-06-24T00:00:00.000+03:00</urn1:CreDtTm>
                     <urn1:Assgnr>
                        <urn1:Agt>
                           <urn1:FinInstnId>
                              <urn1:Othr>
                                 <urn1:Id>CBETETAA</urn1:Id>
                              </urn1:Othr>
                           </urn1:FinInstnId>
                        </urn1:Agt>
                     </urn1:Assgnr>
                     <urn1:Assgne>
                        <urn1:Agt>
                           <urn1:FinInstnId>
                              <urn1:Othr>
                                 <urn1:Id>ETSETAA</urn1:Id>
                              </urn1:Othr>
                           </urn1:FinInstnId>
                        </urn1:Agt>
                     </urn1:Assgne>
                  </urn1:Assgnmt>
                  <urn1:Vrfctn>
                     <urn1:Id>CBETETAA843572771</urn1:Id>
                     <urn1:PtyAndAcctId>
                        <urn1:Acct>
                           <urn1:Id>
                              <urn1:Othr>
                                 <urn1:Id>1234567890</urn1:Id>
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
</soapenv:Envelope>`,
		param.DebitBankBIC,
		param.CreaditBankBIC,
		param.BizMessageIdentifier,
		param.CreditDate,
		param.MessageIdentifier,
		param.CreditDateTime,
	)
}
