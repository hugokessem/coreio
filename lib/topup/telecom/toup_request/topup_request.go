package touprequest

import "fmt"

type Params struct {
	OriginatorConversationID string
	Password                 string
	Timestamp                string
}

func NewTopupRequest(param Params) string {
	return fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:com="http://cps.huawei.com/synccpsinterface/common" xmlns:cus="http://cps.huawei.com/cpsinterface/customizedrequest" xmlns:api="http://cps.huawei.com/synccpsinterface/api_requestmgr" xmlns:req="http://cps.huawei.com/synccpsinterface/request">
   <soapenv:Header/>
   <soapenv:Body>
      <api:Request>
         <req:Header>
            <req:Version>1.0</req:Version>
            <req:CommandID>InitTrans_PrepaidAirtimeOrganization</req:CommandID>
            <req:OriginatorConversationID>df9e719889908977f3e7</req:OriginatorConversationID>
            <req:Caller>
               <req:CallerType>2</req:CallerType>
               <req:ThirdPartyID>TestCallerID</req:ThirdPartyID>
               <req:Password>4HXImxHXpXgB082OtAK7S0Vr+o/Hzx1w0Y0bB2y8ixs=</req:Password>
            </req:Caller>
            <req:KeyOwner>1</req:KeyOwner>
            <req:Timestamp></req:Timestamp>
         </req:Header>
         <req:Body>
            <req:Identity>
               <req:Initiator>
                  <req:IdentifierType>12</req:IdentifierType>
                  <req:Identifier>232323</req:Identifier>
                  <req:SecurityCredential>ur4OB6Ket0eDLWtmiXrq9mtshXcQ05LkcT6ruka9Z0Y=</req:SecurityCredential>
                  <req:ShortCode>232323</req:ShortCode>
               </req:Initiator>
               <req:ReceiverParty>
                  <req:IdentifierType>4</req:IdentifierType>
                  <req:Identifier>111222</req:Identifier>
               </req:ReceiverParty>
            </req:Identity>
            <req:TransactionRequest>
               <req:Parameters>
                  <req:Parameter>
                     <com:Key>RechargedMSISDN</com:Key>
                     <com:Value>251900000012</com:Value>
                  </req:Parameter>
                  <req:Amount>100</req:Amount>
                  <req:Currency>ETB</req:Currency>
               </req:Parameters>
            </req:TransactionRequest>
            <req:ReferenceData>             
               <req:ReferenceItem>
                  <com:Key>Bank Transaction Number</com:Key>
                  <com:Value>CL85105O6QN</com:Value>
               </req:ReferenceItem>
            </req:ReferenceData>
         </req:Body>
      </api:Request>
   </soapenv:Body>
</soapenv:Envelope>
	`)
}
