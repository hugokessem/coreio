package upload

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUpload(t *testing.T) {
	tests := []struct {
		name   string
		param  Params
		expect []string
	}{
		{
			name: "Validate upload XML generation",
			param: Params{
				Username:         "SUPERAPP",
				Password:         "123456",
				Name:             "passport.pdf",
				DocumentType:     "NATIONAL.ID",
				CustomerNumber:   "1202724929",
				Source:           "SUPERAPP",
				UniqueIdentifier: "UID-001",
				Email:            "sample@gmail.com",
				PhoneNumber:      "+251911706628",
				FullName:         "JOHN DOE",
				MimeType:         "application/pdf",
				FileName:         "passport.pdf",
				Stream:           "base64encodedcontent",
			},
			expect: []string{
				`<wsse:Username>SUPERAPP</wsse:Username>`,
				`<wsse:Password Type="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordText">123456</wsse:Password>`,
				`<ns:repositoryId>CMTOS</ns:repositoryId>`,
				`<ns:createDocument>`,
				`<ns1:value>CBEAO_ACCOUNT_DOCS</ns1:value>`,
				`<ns1:value>passport.pdf</ns1:value>`,
				`<ns1:value>NATIONAL.ID</ns1:value>`,
				`<ns1:value>1202724929</ns1:value>`,
				`<ns1:value>SUPERAPP</ns1:value>`,
				`<ns1:value>UID-001</ns1:value>`,
				`<ns1:value>sample@gmail.com</ns1:value>`,
				`<ns1:value>+251911706628</ns1:value>`,
				`<ns1:value>JOHN DOE</ns1:value>`,
				`<ns:folderId>Weleta</ns:folderId>`,
				`<ns:mimeType>application/pdf</ns:mimeType>`,
				`<ns:filename>passport.pdf</ns:filename>`,
				`<ns:stream>base64encodedcontent</ns:stream>`,
			},
		},
		{
			name: "Validate upload XML with different credentials",
			param: Params{
				Username:         "TESTUSER",
				Password:         "PASSWORD123",
				Name:             "id-card.jpg",
				DocumentType:     "KEBELE.ID",
				CustomerNumber:   "1000000001",
				Source:           "MOBILE",
				UniqueIdentifier: "UID-999",
				Email:            "test@example.com",
				PhoneNumber:      "+251913323918",
				FullName:         "JANE DOE",
				MimeType:         "image/jpeg",
				FileName:         "id-card.jpg",
				Stream:           "abc123",
			},
			expect: []string{
				`<wsse:Username>TESTUSER</wsse:Username>`,
				`PASSWORD123</wsse:Password>`,
				`<ns1:value>id-card.jpg</ns1:value>`,
				`<ns1:value>KEBELE.ID</ns1:value>`,
				`<ns1:value>1000000001</ns1:value>`,
				`<ns1:value>UID-999</ns1:value>`,
				`<ns:mimeType>image/jpeg</ns:mimeType>`,
				`<ns:stream>abc123</ns:stream>`,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			xmlRequest := NewUpload(tc.param)

			assert.Contains(t, xmlRequest, "<soapenv:Envelope")
			assert.Contains(t, xmlRequest, "<soapenv:Body>")
			assert.Contains(t, xmlRequest, "<ns:createDocument>")
			assert.Contains(t, xmlRequest, `propertyDefinitionId="UniqueIdentifier"`)
			assert.Contains(t, xmlRequest, `<ns1:value>`+tc.param.UniqueIdentifier+`</ns1:value>`)

			for _, expectedStr := range tc.expect {
				assert.Contains(t, xmlRequest, expectedStr)
			}

			assert.NotEmpty(t, xmlRequest)
		})
	}
}

func TestParseUploadSOAP(t *testing.T) {
	tests := []struct {
		name             string
		xmlData          string
		expectedError    bool
		expectedObjectID string
	}{
		{
			name: "Parse successful response",
			xmlData: `<?xml version="1.0" encoding="UTF-8"?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns:createDocumentResponse xmlns:ns="http://docs.oasis-open.org/ns/cmis/messaging/200908/">
            <ns:objectId>workspace://SpacesStore/abc-123-def</ns:objectId>
            <ns:extension></ns:extension>
        </ns:createDocumentResponse>
    </S:Body>
</S:Envelope>`,
			expectedError:    false,
			expectedObjectID: "workspace://SpacesStore/abc-123-def",
		},
		{
			name: "Parse response with empty objectId",
			xmlData: `<?xml version="1.0" encoding="UTF-8"?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <ns:createDocumentResponse xmlns:ns="http://docs.oasis-open.org/ns/cmis/messaging/200908/">
            <ns:objectId></ns:objectId>
        </ns:createDocumentResponse>
    </S:Body>
</S:Envelope>`,
			expectedError:    false,
			expectedObjectID: "",
		},
		{
			name: "Parse invalid XML",
			xmlData: `<?xml version="1.0" encoding="UTF-8"?>
<InvalidXML>
    <Broken>
</InvalidXML>`,
			expectedError: true,
		},
		{
			name: "Parse response without createDocumentResponse",
			xmlData: `<?xml version="1.0" encoding="UTF-8"?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
    <S:Body>
        <OtherResponse></OtherResponse>
    </S:Body>
</S:Envelope>`,
			expectedError:    false,
			expectedObjectID: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ParseUploadSOAP(tc.xmlData)

			if tc.expectedError {
				assert.Error(t, err)
				assert.Nil(t, result)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, result)
			assert.Equal(t, tc.expectedObjectID, result.ObjectID)
		})
	}
}
