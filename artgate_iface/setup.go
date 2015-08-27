package main

import (
	"net"
	"strconv"
	"fmt"
	"log"
)

type ArtIn struct {
	Universe uint16 // Вселенная
	Enabled bool
	Name    string // Имя Вселенной
}

type ArtOut struct {
	Universe uint16 // Вселенная
	Enabled bool
	Name string
}


type Setup struct {
	IpAddress string	// Адрес IP
	IpMask 	  string	// Маска IP
	IpGw 	  string    // Шлюз IP
	Mac		  string    // MAC Адрес
	ArtnetInputs int    // Число входов ArtNet
	ArtIns	map[int]ArtIn		// Входы ArtNet
	ArtnetOutputs int // Число выходов ArtNet
	ArtOuts map[int]ArtOut
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

	// Нужно сделать следующее
	// 1 Взять из старого конфига все порты
	// 2 Переписать их в новый конфиг
	// 3 Дописать чистые конфиги по входам
	numOldActualIns := s.ArtnetInputs - i
	if numOldActualIns < 0 { numOldActualIns = 0}
	idx := 0
	for idx= 0; idx < numOldActualIns; idx++{
	}
	for ;idx < i; idx++ {
		s.ArtIns[idx] = ArtIn{
			Enabled: false,
			Universe: 0,
			Name: fmt.Sprintf("tag%d",idx),
		}
	}
	for ;idx < s.ArtnetInputs; idx++ {
		delete(s.ArtIns, idx)
	}
	s.ArtnetInputs = i
	return nil
}

func (s *Setup) UpdateArtNetOutputs(numArtnet string) error {
	i, err := strconv.Atoi(numArtnet)
	if err != nil{
		return errInvalidArtnetOutputs
	}

	// Нужно сделать следующее
	// 1 Взять из старого конфига все порты
	// 2 Переписать их в новый конфиг
	// 3 Дописать чистые конфиги по входам
	numOldActualOuts := s.ArtnetOutputs - i
	if numOldActualOuts < 0 { numOldActualOuts = 0}
	idx := 0
	for idx= 0; idx < numOldActualOuts; idx++{
	}
	for ;idx < i; idx++ {
		s.ArtOuts[idx] = ArtOut{
			Enabled: false,
			Universe: 0,
			Name: fmt.Sprintf("tag%d",idx),
		}
	}
	for ;idx < s.ArtnetOutputs; idx++ {
		delete(s.ArtOuts, idx)
	}
	s.ArtnetOutputs = i
	return nil
}


func (s *Setup) EnableArtnetIn(idx int){
	vals := s.ArtIns[idx]
	vals.Enabled = true
	s.ArtIns[idx] = vals
}

func (s *Setup) DisableArtnetIn(idx int){
	vals := s.ArtIns[idx]
	vals.Enabled = false
	s.ArtIns[idx] = vals
}

func (s *Setup) EnableArtnetOut(idx int){
	vals := s.ArtOuts[idx]
	vals.Enabled = true
	s.ArtOuts[idx] = vals
}

func (s *Setup) DisableArtnetOut(idx int){
	vals := s.ArtOuts[idx]
	vals.Enabled = false
	s.ArtOuts[idx] = vals
}



func (s *Setup) UpdateArtNetInUniverse(idx int, v string) error {

	i, err := strconv.Atoi(v)
	if err != nil{
		return errInvalidUniverse
	}

	vals := s.ArtIns[idx]

	vals.Universe = uint16(i)
	log.Printf("Vals: %v", vals)
	s.ArtIns[idx] = vals
	log.Printf("Vals: %v", s.ArtIns[idx])
	return nil
}

func (s *Setup) UpdateArtNetOutUniverse(idx int, v string) error {

	i, err := strconv.Atoi(v)
	if err != nil{
		return errInvalidUniverse
	}

	vals := s.ArtOuts[idx]
	vals.Universe = uint16(i)
	s.ArtOuts[idx] = vals
	return nil
}


func NewSetup() *Setup {
	return &Setup{
		IpAddress: "10.101.0.245",
		IpMask: "255.0.0.0",
		IpGw: "10.0.0.1",
		Mac: "00:01:02:03:04:05",
		ArtnetInputs: 0,
		ArtIns: map[int]ArtIn{},
		ArtOuts: map[int]ArtOut{},
	}
}

var globalSetup *Setup

func init(){
	globalSetup = NewSetup()
}
