package processors

import (
	"fmt"
)

import "netcommander/messages"


type HttpRespProcessor struct {
	respMessage *messages.HttpRespMessage
}

func NewHttpRespProcessor(pRespMessage *messages.HttpRespMessage) *HttpRespProcessor {
	
	var theProcessor *HttpRespProcessor
	theProcessor = new(HttpRespProcessor)
	
	theProcessor.respMessage = pRespMessage
	
	return theProcessor
}

func analyseRespCode(pRespMessage *messages.HttpRespMessage) *messages.HttpRespMessage {
	
	var respCode int
	var respCodeType int
	
	respCode = pRespMessage.GetCode()
	respCodeType = respCode/100
	
	switch(respCodeType) {
		
		case 1:
			informationalStatus(respCode)
		
		case 2:
			return successStatus(respCode, pRespMessage)
			
		case 3:
			redirectionStatus(respCode, pRespMessage)
		
		case 4:
			clientErrStatus(respCode)
		
		case 5:
			serverErrStatus(respCode)
		
	}
	
	return nil
}

func informationalStatus(pRespCode int) {
	
	const majorCode = 100
	var minorCode int
	
	minorCode = pRespCode - majorCode
	
	switch(minorCode) {
		
		case 0:
			
		
		case 1:
		
		
	}
	
}

func successStatus(pRespCode int, pRespMessage *messages.HttpRespMessage) *messages.HttpRespMessage {
	
	const majorCode = 200
	var minorCode int
	
	minorCode = pRespCode - majorCode
	
	switch(minorCode) {
		
		case 0:
			return pRespMessage
		
		case 1:
		
		case 2:
		
		case 3:
		
		case 4:
		
		case 5:
		
	}
	
	return nil
}

func redirectionStatus(pRespCode int, pRespMessage *messages.HttpRespMessage) *messages.HttpRespMessage {
	
	const majorCode = 300
	var minorCode int
	
	minorCode = pRespCode - majorCode
	
	switch(minorCode) {
		
		case 0:
		
		case 1:
			
			newHttpRqstMessage := messages.NewHttpRqstMsg()
			httpRqstProessor := NewHttpRqstProcessor()
			
			fmt.Println(pRespMessage.GetHeader().GetValue("Location"))
			
			httpRqstProessor.Open("GET", pRespMessage.GetHeader().GetValue("Location"))
			httpRqstProessor.FlushUp(newHttpRqstMessage)
			
		case 2:
		
		case 3:
		
		case 4:
		
		case 5:
		
		case 6:
		
		case 7:
		
	}
	
	return nil
}

func clientErrStatus(pRespCode int) {
	
	const majorCode = 400
	var minorCode int
	
	minorCode = pRespCode - majorCode
	
	switch(minorCode) {
		
		case 0:
		
		case 2:
		
		case 3:
		
		case 4:
		
		case 5:
		
		case 6:
		
		case 8:
		
		case 9:
		
		case 10:
		
		case 11:
		
		case 13:
		
		case 14:
		
		case 15:
		
		case 17:
		
		case 26:
		
	}
	
}

func serverErrStatus(pRespCode int) {
	
	const majorCode = 500
	var minorCode int
	
	minorCode = pRespCode - majorCode
	
	switch(minorCode) {
		case 0:
		
		case 1:
		
		case 2:
		
		case 3:
		
		case 4:
		
		case 5:
		
	}
	
}

func (theRespProcessor *HttpRespProcessor) GetTheResult() *messages.HttpRespMessage {
	
	

	return analyseRespCode(theRespProcessor.respMessage)

}

