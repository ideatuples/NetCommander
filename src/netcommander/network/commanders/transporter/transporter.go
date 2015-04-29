package transporter

import (
	
	"net"
	"errors"
)

import "netcommander/network/utils"

type Transporter struct {
	
	connHandler net.Conn
	
	networkType string
	urlScheme string
	targetHost string
	
}

func NewTransporter(pExURL *utils.ExURL) *Transporter {
	
	var transporter *Transporter
	
	transporter= new(Transporter)
	
	transporter.targetHost = pExURL.GetHost()
	transporter.urlScheme = pExURL.GetScheme()
	transporter.networkType = "tcp"
	
	return transporter
}

func NewTransporterEx(pExURL *utils.ExURL, pNetworkType string) *Transporter {
	
	var transporter *Transporter
	
	transporter= new(Transporter)
	transporter.targetHost = pExURL.GetHost()
	transporter.networkType = pNetworkType
	
	return transporter
}

func (transporter Transporter) GetConnHandler() (net.Conn, error) {
	
	if(transporter.urlScheme == "https") {
		
		return nil, errors.New("Not yet implemented")		
	} 
	
	return net.Dial(transporter.networkType, transporter.targetHost)
}

