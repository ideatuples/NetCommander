package utils

import (
	"fmt"
	"strings"
	"net/url"
)

type ExURL struct {
	
	preURL *url.URL
	
	//targetScheme string
	targetDomain string
	//targetHost string
	targetPort string
	targetResource string
	targetHostname string
	
	targetRawURL string
}

func NewExURL(pRawURL string) *ExURL {
	
	var exURL *ExURL
	exURL = new(ExURL)
	
	var rawURL string
	var preURL *url.URL
	
	var parseErr error
	
	rawURL = checkoutScheme(pRawURL)
	exURL.targetRawURL = rawURL
	
	if preURL, parseErr = url.Parse(rawURL); parseErr != nil {
		fmt.Println(parseErr)
	}
	
	exURL.preURL = preURL
	exURL.parseHostname(preURL)
	
	return exURL
}

func checkoutScheme(pRawURL string) string{
	
	var checkoutResult string
	
	containedOK := strings.Contains(pRawURL, "://")
	
	if(containedOK != true) {
		checkoutResult = strings.Join([]string {"http", pRawURL}, "://")
		return checkoutResult
	}
	
	return pRawURL
}

func (exURL *ExURL) parseHostname(pPreURL *url.URL) {
	
	unparsedHost := pPreURL.Host
	
	if (strings.Contains(unparsedHost, ":")) {
		splittedName := strings.Split(unparsedHost, ":")
		exURL.targetDomain = splittedName[0]
		exURL.targetPort = splittedName[1]
		
	} else {
		
		exURL.targetDomain = unparsedHost
		
		switch(pPreURL.Scheme) {
			
			case "http":
				exURL.targetPort = "80"
				
			case "ftp":
				exURL.targetPort = "21"
				
			case "ssh":
				exURL.targetPort = "22"
			
			case "sftp":
				exURL.targetPort = "22"
				
			case "telnet":
				exURL.targetPort = "23"
		}
	}
	
}


func (exURL *ExURL) GetScheme() string {
	
	return exURL.preURL.Scheme
	
}

func (exURL *ExURL) GetResource() string {
	
	if(exURL.preURL.Path == ""){
		return "/"
	}
	
	return exURL.preURL.Path
	
}

func (exURL *ExURL) GetDomain() string {
	
	return exURL.targetDomain
}

func (exURL *ExURL) GetPort() string {
	return exURL.targetPort
}

func (exURL *ExURL) GetHost() string {
	
	return fmt.Sprintf("%s:%s", exURL.targetDomain, exURL.targetPort)
}

func (exURL *ExURL) GetRawURL() string {
	
	return exURL.targetRawURL
}
