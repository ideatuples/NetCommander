package messages

import (
	"fmt"
	"bytes";"strings"
)

type IHttpRqstMessage interface {
	
	AddExtraValue(pFieldKey string, pFieldValue string) bool
	Generate(pCommand string, pHostField string, pResourceField string) string
	
	AddBodyContents(pBodyContents string) bool
}

type HttpRqstMessage struct {
	
	rqstHeader HttpRqstFmt
	rqstBody string
}

func NewHttpRqstMsg() *HttpRqstMessage {
	
	var rqstMsg *HttpRqstMessage
	rqstMsg = new(HttpRqstMessage)
	
	rqstMsg.rqstHeader = NewHttpRqstFmt()
	rqstMsg.putDefaults()
	
	return rqstMsg
}


func (theMessage HttpRqstMessage) putDefaults() bool {
	
	//defaultMsg.reqMessage.PutValue("Accept", "text/html")
	//defaultMsg.reqMessage.PutValue("Accept-Language", "en-GB")
	
	//defaultMsg.reqMessage.PutValue("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2272.118 Safari/537.36")
	
	return true
}

func (theMessage HttpRqstMessage) AddExtraValue(pFieldKey string, pFieldValue string) bool {
	
	theMessage.rqstHeader.PutValue(pFieldKey, pFieldValue)
	
	return true
}

func (theMessage *HttpRqstMessage) AddBodyContents(pBodyContents string) bool {
	theMessage.rqstBody = pBodyContents
	
	return true
}

func (theMessage HttpRqstMessage) Generate(pCommand string, pHostValue string, pResourceField string) string {
	
	var headline string
	var hostLine string
	
	var criticalLines string
	
	var msgBuffer *bytes.Buffer
	
	defer func() {
		msgBuffer.Reset()
		
	} ()
	
	headline = fmt.Sprintf("%s %s %s\r\n", pCommand, pResourceField, "HTTP/1.1")
	hostLine = fmt.Sprintf("%s: %s\r\n", "Host", pHostValue)
	
	criticalLines = strings.Join([]string {headline, hostLine}, "")
	
	msgBuffer = bytes.NewBufferString(criticalLines)
	
	
	for key, val := range theMessage.rqstHeader.msgFields {
		
		if(val == ""){
			continue
		}
		
		msgBuffer.WriteString(fmt.Sprintf("%s: %s\r\n", key, val))
		
	}
	msgBuffer.WriteString("\r\n")
	
	if(theMessage.rqstBody != "") {
		msgBuffer.WriteString(theMessage.rqstBody)
	}
	
	return msgBuffer.String()
}

/***
 *
 *
 *
 ***/

type HttpRespMessage struct {
	
	respCode int
	respHeader HttpRespFmt
	respBody []string
	
}

func NewHttpRespMsg(pCode int, pHeader HttpRespFmt, pBody []string) *HttpRespMessage {
	
	var respMsg *HttpRespMessage
	respMsg = new(HttpRespMessage)
	
	respMsg.respCode = pCode
	respMsg.respHeader = pHeader
	respMsg.respBody = pBody
	
	return respMsg
}

func (respMsg *HttpRespMessage) GetHeader() HttpRespFmt {
	
	return respMsg.respHeader
}

func (respMsg *HttpRespMessage) GetCode() int {
	
	return respMsg.respCode
}

func (respMsg *HttpRespMessage) GetBody() []string {
	
	return respMsg.respBody
}


