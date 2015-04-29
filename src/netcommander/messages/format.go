package messages

type HttpRqstFmt struct {
	msgFields map[string] string 
}

type HttpRespFmt struct {
	msgFields map[string] string
}

/**
 * Below functions is for RequesteMessage
 */

func NewHttpRqstFmt() HttpRqstFmt{
	
	var httpRqstFmt HttpRqstFmt
	
	//rqstMessage = new(RqstMsgFormat)
	
	//reqMessage = new(RequestMessage)
	httpRqstFmt.msgFields = make(map[string] string)
	
	/**
	 * Request Header Fields
	 **/
	//reqMessage.msgFields["Host"]=""
	httpRqstFmt.msgFields["Accept"]=""
	httpRqstFmt.msgFields["Accept-Charset"]=""
	httpRqstFmt.msgFields["Accept-Encoding"]=""
	httpRqstFmt.msgFields["Accept-Language"]=""
	httpRqstFmt.msgFields["Authorization"]=""
	httpRqstFmt.msgFields["Content-Encoding"]=""
	httpRqstFmt.msgFields["Expect"]=""
	httpRqstFmt.msgFields["From"]=""
	httpRqstFmt.msgFields["If-match"]=""
	httpRqstFmt.msgFields["If-Modified-Since"]=""
	httpRqstFmt.msgFields["If-None-Match"]=""
	httpRqstFmt.msgFields["If-Range"]=""
	httpRqstFmt.msgFields["If-Unmodified-Since"]=""
	httpRqstFmt.msgFields["Max-Forwards"]=""
	httpRqstFmt.msgFields["Proxy-Authorization"]=""
	httpRqstFmt.msgFields["Range"]=""
	httpRqstFmt.msgFields["Referer"]=""
	httpRqstFmt.msgFields["TE"]=""
	httpRqstFmt.msgFields["User-Agent"]=""

	
	/**
	 * General Header Fields
	 **/
	httpRqstFmt.msgFields["Cache-Control"]=""
	httpRqstFmt.msgFields["Connection"]=""
	httpRqstFmt.msgFields["Date"]=""
	httpRqstFmt.msgFields["Pragma"]=""
	httpRqstFmt.msgFields["Trailer"]=""
	httpRqstFmt.msgFields["Transfer-Encoding"]=""
	httpRqstFmt.msgFields["Upgrade"]=""
	httpRqstFmt.msgFields["Via"]=""
	httpRqstFmt.msgFields["Warning"]=""
		
	/**
	 * Entity Header Fields
	 **/
	httpRqstFmt.msgFields["Allow"]=""
	httpRqstFmt.msgFields["Content-Encoding"]=""
	httpRqstFmt.msgFields["Content-Language"]=""
	httpRqstFmt.msgFields["Content-Length"]=""
	httpRqstFmt.msgFields["Content-Location"]=""
	httpRqstFmt.msgFields["Content-MD5"]=""
	httpRqstFmt.msgFields["Content-Range"]=""
	httpRqstFmt.msgFields["Content-Type"]=""
	httpRqstFmt.msgFields["Expires"]=""
	httpRqstFmt.msgFields["Last-Modified"]=""
	
	return httpRqstFmt
}

func (httpRqstFmt *HttpRqstFmt) PutValue(pFieldName string, pFieldValue string) {
	
	httpRqstFmt.msgFields[pFieldName] = pFieldValue
}

func (httpRqstFmt HttpRqstFmt) GetValue(pFieldName string) string {
	return httpRqstFmt.msgFields[pFieldName]
}


/**
 * Below functions is for ResponseMessage
 */

func NewHttpRespFmt() HttpRespFmt {
	/**
	 * General Header Fields
	 **/
	
	var httpRespFmt HttpRespFmt
	httpRespFmt.msgFields = make(map[string] string)
	
	httpRespFmt.msgFields["Cache-Control"]=""
	httpRespFmt.msgFields["Connection"]=""
	httpRespFmt.msgFields["Date"]=""
	httpRespFmt.msgFields["Pragma"]=""
	httpRespFmt.msgFields["Trailer"]=""
	httpRespFmt.msgFields["Transfer-Encoding"]=""
	httpRespFmt.msgFields["Upgrade"]=""
	httpRespFmt.msgFields["Via"]=""
	httpRespFmt.msgFields["Warning"]=""
	
	/**
	 * Response Header Fields
	 **/
	httpRespFmt.msgFields["Accept-Ranges"]=""
	httpRespFmt.msgFields["Age"]=""
	httpRespFmt.msgFields["ETag"]=""
	httpRespFmt.msgFields["Last-Modified"]=""
	httpRespFmt.msgFields["Location"]=""
	httpRespFmt.msgFields["Proxy-Authentication"]=""
	httpRespFmt.msgFields["Retry-After"]=""
	httpRespFmt.msgFields["Server"]=""
	httpRespFmt.msgFields["Vary"]=""
	httpRespFmt.msgFields["WWW-Authenticate"]=""
	
	/**
	 * Entity Header Fields
	 **/
	httpRespFmt.msgFields["Allow"]=""
	httpRespFmt.msgFields["Content-Encoding"]=""
	httpRespFmt.msgFields["Content-Language"]=""
	httpRespFmt.msgFields["Content-Length"]=""
	httpRespFmt.msgFields["Content-Location"]=""
	httpRespFmt.msgFields["Content-MD5"]=""
	httpRespFmt.msgFields["Content-Range"]=""
	httpRespFmt.msgFields["Content-Type"]=""
	httpRespFmt.msgFields["Expires"]=""
	httpRespFmt.msgFields["Last-Modified"]=""
	
	return httpRespFmt
}

func (httpRespFmt *HttpRespFmt) PutValue(pFieldName string, pFieldValue string) {
	
	httpRespFmt.msgFields[pFieldName] = pFieldValue
}

func (httpRespFmt HttpRespFmt) GetValue(pFieldName string) string{
	return httpRespFmt.msgFields[pFieldName]
}

