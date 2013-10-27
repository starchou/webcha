package utils

import (
	"encoding/xml"
	"math/big"
)

var _ = big.MaxBase // Avoid potential unused-import error

type ArrayOfString struct {
	String []*string `xml:"http://WebXml.com.cn/ string"`
}

type getDatabaseInfo struct {
	XMLName xml.Name `xml:"http://WebXml.com.cn/ getDatabaseInfo"`
}

type getDatabaseInfoResponse struct {
	XMLName xml.Name `xml:"http://WebXml.com.cn/ getDatabaseInfoResponse"`
}

type GetMobileCodeInfo struct {
	XMLName    xml.Name `xml:"http://WebXml.com.cn/ getMobileCodeInfo"`
	MobileCode string   `xml:"mobileCode"`
	UserID     string   `xml:"userID"`
}

type GetMobileCodeInfoResponse struct {
	XMLName xml.Name `xml:"http://WebXml.com.cn/ getMobileCodeInfoResponse"`
	Value   string   `xml:"getMobileCodeInfoResult"`
}
