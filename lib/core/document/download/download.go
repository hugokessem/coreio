package download

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type Params struct {
	Username   string
	Password   string
	DocumentID string
}

func NewDownload(param Params) string {
	return fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns="http://docs.oasis-open.org/ns/cmis/messaging/200908/">
	<soapenv:Header>
		<wsse:Security xmlns:wsse="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" soapenv:mustUnderstand="1" xmlns:wsu="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd">
			<wsse:UsernameToken>
				<wsse:Username>%s</wsse:Username>
				<wsse:Password Type="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordText">%s</wsse:Password>
			</wsse:UsernameToken>
		</wsse:Security>
	</soapenv:Header>
   <soapenv:Body>
      <ns:getContentStream>
	   <!--Pass Document repository,Pass "CMTOS CREDIT" as fixed Value-->
         <ns:repositoryId>CMTOS</ns:repositoryId>
		  <!--Pass Document document GUID to be downloaded-->
         <ns:objectId>%s</ns:objectId>
      </ns:getContentStream>
   </soapenv:Body>
</soapenv:Envelope>
`, param.Username, param.Password, param.DocumentID)
}

type Envelope struct {
	Body Body `xml:"Body"`
}

type Body struct {
	GetContentStreamResponse *ContentStream `xml:"getContentStreamResponse"`
}

type ContentStream struct {
	Length   string `xml:"length"`
	MimeType string `xml:"mimeType"`
	Filename string `xml:"filename"`
	Stream   string `xml:"stream"`
}

func ParseDownloadSOAP(xmlData string) (*ContentStream, error) {
	var envelope Envelope
	err := xml.Unmarshal([]byte(xmlData), &envelope)
	if err != nil {
		return nil, err
	}

	if envelope.Body.GetContentStreamResponse == nil {
		return nil, errors.New("getContentStreamResponse is nil")
	}

	response := envelope.Body.GetContentStreamResponse

	return response, nil
}
