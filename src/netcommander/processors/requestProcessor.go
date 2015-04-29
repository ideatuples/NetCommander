package processors

import (
	"fmt"
	"strings"
)

import (
	"netcommander/network/utils"
	commanders "netcommander/network/commanders"
)

import "netcommander/messages"


type HttpRqstProcessor struct {
	theQueue *utils.TaskQueue
}

func NewHttpRqstProcessor() *HttpRqstProcessor {
	
	var theProcessor *HttpRqstProcessor
	
	theProcessor = new(HttpRqstProcessor)
	theProcessor.theQueue = utils.NewTaskQueue()
	
	return theProcessor
}

func (theProcessor *HttpRqstProcessor) Open(pMethod, pRawURLs string) bool {
	
	var fieldinisedURLs []string
	fieldinisedURLs = theProcessor.fieldiniseURLs(pRawURLs)
	theProcessor.queing(pMethod, fieldinisedURLs)
	
	return true
}

func (theProcessor *HttpRqstProcessor) fieldiniseURLs(pRawURLs string) []string {
	
	return strings.Fields(pRawURLs)
}

func (theProcessor *HttpRqstProcessor) queing(pMethod string, pFieldinisedURLs []string) bool {
	
	for i := range pFieldinisedURLs {
		
		var exURL *utils.ExURL
		exURL = utils.NewExURL(pFieldinisedURLs[i])
		
		theProcessor.theQueue.PushBack(pMethod, exURL)
	}
	
	return true
}

func (theProcessor *HttpRqstProcessor) FlushUp(pRqstMessage messages.IHttpRqstMessage) ([] *messages.HttpRespMessage) {
	
	
	queueLength := theProcessor.theQueue.GetLength()
	respResults := make([] *messages.HttpRespMessage, 0)
	
	for i := 0; i < queueLength; i++ {
		
		fmt.Println(i)
		
		respMessage, connErr := performRequest(theProcessor.theQueue.PullFront(), pRqstMessage)
		
		if(connErr != nil) {
			respResults = append(respResults, nil)
			continue
		}
		
		respResults = append(respResults, respMessage)
	}
	
	return respResults
}

func performRequest(pTask utils.Task, pRqstMessage messages.IHttpRqstMessage) (*messages.HttpRespMessage, error) {
	
	switch (pTask.GetMethod()) {
		
		case "GET": 
			respMessage, connErr := commanders.GET_HTTP(pTask.GetExURL(), pRqstMessage)
			
			return respMessage, connErr
			
		case "POST":
			respMessage, connErr:= commanders.GET_HTTP(pTask.GetExURL(), pRqstMessage)
			return respMessage, connErr
		
	}
	
	return nil, nil
}
