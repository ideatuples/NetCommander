package main

import (
	//"fmt"
)

import "netcommander/messages"
import "netcommander/processors"
//import "netcommander/network/commanders"
//import "netcommander/network/commanders/transporter"
//import "netcommander/network/utils"

func main() {
	
	theRqstMsg := messages.NewHttpRqstMsg()
	//theURL := utils.NewExURL("http://www.daum.net")
	
	rqstProcessor := processors.NewHttpRqstProcessor()
	rqstProcessor.Open("GET", "http://www.daum.net")
	rqstProcessor.Open("GET", "http://www.aaaaaaaaaaaaaaaaaaaaa.net")
	rqstProcessor.Open("GET", "https://www.yahoo.com")
	rqstProcessor.FlushUp(theRqstMsg)
	
	
}	

