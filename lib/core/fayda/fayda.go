package fayda

import "fmt"

type Params struct {
	Username         string
	Password         string
	NID              string
	CustomerMnemonic string
}

type FaydaParam struct {
	NID              string
	CustomerMnemonic string
}

func NewFayda(param Params) string {
	return fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"
xmlns:cbes="http://temenos.com/CBESUPERAPP">
    <soapenv:Header/>
    <soapenv:Body>
        <cbes:CustomerUniqueVerification>
            <WebRequestCommon>
                <company/>
                <password>%s</password>
                <userName>%s</userName>
            </WebRequestCommon>
            <CUSTOMERVERIFYENQSUPERAPPType>
                <enquiryInputCollection>
                    <columnName>NID</columnName>
                    <criteriaValue>%s</criteriaValue>
                    <operand>EQ</operand>
                </enquiryInputCollection>
                <enquiryInputCollection>
                    <columnName>CUS.MENMO</columnName>
                    <criteriaValue>%s</criteriaValue>
                    <operand>EQ</operand>
                </enquiryInputCollection>
            </CUSTOMERVERIFYENQSUPERAPPType>
        </cbes:CustomerUniqueVerification>
    </soapenv:Body>
</soapenv:Envelope>
	`, param.Password, param.Username, param.NID, param.CustomerMnemonic)
}
