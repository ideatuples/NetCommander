package commanders

import (
	"fmt"
	"net"
	"io"; "bufio"
	"strings"; "strconv"
)

import (
	"netcommander/network/utils"
	"netcommander/network/commanders/transporter"
)

import "netcommander/messages"


func GET_HTTP(pExURL *utils.ExURL, pRqstMessage messages.IHttpRqstMessage) (*messages.HttpRespMessage, error){
	
	const commandName = "GET"
	
	var netConn net.Conn; var connErr error
	var respRawMessage []string; var recvErr error
	
	var respMessage *messages.HttpRespMessage
	
	
	defer func() {
		if(connErr == nil) {
			netConn.Close()
		}
	} ()
	
	if netConn, connErr = estHTTPConnection(pExURL); connErr != nil {
		return nil, connErr
	}
	
	rwBuffer := bufio.NewReadWriter(bufio.NewReader(netConn), bufio.NewWriter(netConn))
	rqstMessage := pRqstMessage.Generate(commandName, pExURL.GetHost(), pExURL.GetResource())
	
	if sendErr := sendRqstMessage(rwBuffer.Writer, rqstMessage); sendErr != nil {
		
		return nil, sendErr
		
	} else if respRawMessage, recvErr = recvRespMessage(rwBuffer.Reader); recvErr != nil {
		
		return nil, recvErr
	}
	
	respCode, respHeader, respBody := parseHTTPRespRawMsg(respRawMessage)
	respMessage = messages.NewHttpRespMsg(respCode, respHeader, respBody)
	
	return respMessage, nil
}

func sendRqstMessage(pSendBuffer *bufio.Writer, pRqstMsg string) error {
	
	pSendBuffer.WriteString(pRqstMsg)
	
	if flushErr := pSendBuffer.Flush(); flushErr != nil {
		return flushErr
	}
	
	return nil
}

func recvRespMessage(pRecvBuffer *bufio.Reader) ([]string, error) {
	
	respRawMessage := make([] string, 0)
	
	for {
		
		recvMsgLine, recvErr := pRecvBuffer.ReadString('\n')
		
		if(recvErr != nil) {
			
			if (recvErr == io.EOF) {
				
				break
				
			} else {
				
				return nil, recvErr
			}
			
		} else {
			
			respRawMessage = append(respRawMessage, recvMsgLine)
		}
			
	}
	
	return respRawMessage, nil
}

func estHTTPConnection(pExURL *utils.ExURL) (net.Conn, error) {
	
	dataTransporter := transporter.NewTransporter(pExURL)
	netConn, connErr := dataTransporter.GetConnHandler()
	
	return netConn, connErr
}

func parseHTTPRespRawMsg(pRespRawMsg []string) (int, messages.HttpRespFmt, []string){
	
	var delimiterLineNum int
	
	for i := 0; i < len(pRespRawMsg); i++ {
		
		if(pRespRawMsg[i] == "\r\n") {
			delimiterLineNum = i
			break
		}
	}
	
	respRawHeader := make([]string, delimiterLineNum)
	respRawBody := make([]string, len(pRespRawMsg)-1)
	
	copy(respRawHeader, pRespRawMsg[0:delimiterLineNum])
	copy(respRawBody, pRespRawMsg[delimiterLineNum+1:])
	
	respCode, respHeader := getFormmattedHTTPRespHeader(respRawHeader)
	respBody := trimHTTPBodyDelimiters(respRawBody)
	
	return respCode, respHeader, respBody
}

func getFormmattedHTTPRespHeader (pRespHeader []string) (int, messages.HttpRespFmt) {
	
	formattedRespHeader := messages.NewHttpRespFmt()
	
	httpCodeLine := strings.Fields(pRespHeader[0])
	httpCode, atoiErr := strconv.Atoi(httpCodeLine[1])
	
	if(atoiErr != nil) {
		fmt.Println(atoiErr)
	}
	
	for i := 1; i < len(pRespHeader); i++ {
		trimmedBodyLine := strings.TrimLeft(pRespHeader[i], "\r\n")
		trimmedBodyLine = strings.TrimSpace(trimmedBodyLine)
		
		splittedBodyLine := strings.SplitN(trimmedBodyLine, ":", 2)
		
		formattedRespHeader.PutValue(splittedBodyLine[0], splittedBodyLine[1])
	}
	return httpCode, formattedRespHeader
}

func trimHTTPBodyDelimiters (pRespBody []string) []string {
	
	for i := 0; i < len(pRespBody); i++ {
		pRespBody[i] = strings.TrimLeft(pRespBody[i], "\r\n")
	}
	
	return pRespBody
}
