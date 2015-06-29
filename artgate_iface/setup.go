package main

import (
	"net"
	"strconv"
)

type Setup struct {
	IpAddress string	// Адрес IP
	IpMask 	  string	// Маска IP
	IpGw 	  string    // Шлюз IP
	Mac		  string    // MAC Адрес
	ArtnetInputs int    // Число входов ArtNet
}

func (s *Setup) UpdateIpAddr(ipAddr string) error {
	ipA := net.ParseIP(ipAddr)
	if ipA == nil {
		return errInvalidIP
	}
	s.IpAddress = ipAddr
	return nil
}

func (s *Setup) UpdateIpMask(ipMask string) error {
	ipM := net.ParseIP(ipMask)
	if ipM == nil {
		return errInvalidMask
	}
	s.IpMask = ipMask
	return nil
}

func (s *Setup) UpdateIpGateway(ipGw string) error {
	ipG := net.ParseIP(ipGw)
	if ipG == nil {
		return errInvalidGw
	}
	s.IpGw = ipGw
	return nil
}

func (s *Setup) UpdateMac(macs string) error {
	_,err := net.ParseMAC(macs)
	if err != nil {
		return errInvalidMAC
	}
	s.Mac = macs
	return nil
}

func (s *Setup) UpdateArtNetInputs(numArtnet string) error {
	i, err := strconv.Atoi(numArtnet)
	if err != nil{
		return errInvalidArtnetInputs
	}
	s.ArtnetInputs = i
	return nil
}

func NewSetup() *Setup {
	return &Setup{
		IpAddress: "10.101.0.245",
		IpMask: "255.0.0.0",
		IpGw: "10.0.0.1",
		Mac: "00:01:02:03:04:05",
		ArtnetInputs: 2,
	}
}

var globalSetup *Setup

func init(){
	globalSetup = NewSetup()
}

