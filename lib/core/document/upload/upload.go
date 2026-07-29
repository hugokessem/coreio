package upload

import (
	"encoding/xml"
	"fmt"
)

type Params struct {
	Username         string
	Password         string
	Name             string
	DocumentType     string
	CustomerNumber   string
	Source           string
	UniqueIdentifier string
	Email            string
	PhoneNumber      string
	FullName         string
	MimeType         string
	FileName         string
	Stream           string
}

type UploadParams struct {
	Name             string
	DocumentType     string
	CustomerNumber   string
	Source           string
	UniqueIdentifier string
	Email            string
	PhoneNumber      string
	FullName         string
	FolderID         string
	MimeType         string
	FileName         string
	Stream           string
}

func NewUpload(param Params) string {
	return fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns="http://docs.oasis-open.org/ns/cmis/messaging/200908/" xmlns:ns1="http://docs.oasis-open.org/ns/cmis/core/200908/">
	<soapenv:Header>
		<wsse:Security xmlns:wsse="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" soapenv:mustUnderstand="1" xmlns:wsu="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd">
			<wsse:UsernameToken>
				<wsse:Username>%s</wsse:Username>
				<wsse:Password Type="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordText">%s</wsse:Password>
			</wsse:UsernameToken>
		</wsse:Security>
	</soapenv:Header>
	<soapenv:Body>
		<ns:createDocument>
			<ns:repositoryId>CMTOS</ns:repositoryId>
			<ns:properties>
				<ns1:propertyString propertyDefinitionId="cmis:objectTypeId">
					<ns1:value>CBEAO_ACCOUNT_DOCS</ns1:value>
				</ns1:propertyString>
				<ns1:propertyString propertyDefinitionId="cmis:name">
					<ns1:value>%s</ns1:value>
				</ns1:propertyString>
				<ns1:propertyString propertyDefinitionId="CBEAO_DocumentType">
					<ns1:value>%s</ns1:value>
				</ns1:propertyString>
				<ns1:propertyString propertyDefinitionId="CBEAO_CaseReferenceNumber">
					<ns1:value>%s</ns1:value>
				</ns1:propertyString>
				<ns1:propertyString propertyDefinitionId="Source">
					<ns1:value>%s</ns1:value>
				</ns1:propertyString>
				<ns1:propertyString propertyDefinitionId="UniqueIdentifier">
					<ns1:value>%s</ns1:value>
				</ns1:propertyString>
				<ns1:propertyString propertyDefinitionId="Field1">
					<ns1:value>%s</ns1:value>
				</ns1:propertyString>
				<ns1:propertyString propertyDefinitionId="Field2">
					<ns1:value>%s</ns1:value>
				</ns1:propertyString>
			    <ns1:propertyString propertyDefinitionId="Field3">
					<ns1:value>%s</ns1:value>
				</ns1:propertyString>
		</ns:properties>
		<ns:folderId>Weleta</ns:folderId>
		<ns:contentStream>
			<ns:mimeType>%s</ns:mimeType>
			<ns:filename>%s</ns:filename>
			<ns:stream>%s</ns:stream>
		</ns:contentStream>
		<ns:extension>
         </ns:extension>
	</ns:createDocument>
</soapenv:Body>
</soapenv:Envelope>
	`, param.Username,
		param.Password,
		param.Name,
		param.DocumentType,
		param.CustomerNumber,
		param.Source,
		param.UniqueIdentifier,
		param.Email,
		param.PhoneNumber,
		param.FullName,
		param.MimeType,
		param.FileName,
		param.Stream,
	)
}

type Envelope struct {
	Body Body `xml:"Body"`
}

type Body struct {
	CreateDocumentResponse CreateDocumentResponse `xml:"createDocumentResponse"`
}

type CreateDocumentResponse struct {
	ObjectID  string `xml:"objectId"`
	Extension string `xml:"extension"`
}

func ParseUploadSOAP(xmlData string) (*CreateDocumentResponse, error) {
	var envelope Envelope
	err := xml.Unmarshal([]byte(xmlData), &envelope)
	if err != nil {
		return nil, err
	}

	return &envelope.Body.CreateDocumentResponse, nil
}
