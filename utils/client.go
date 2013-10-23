package utils

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type SoapEnvelope struct {
	XMLName xml.Name `xml:"soap12:Envelope"`
	Xsi     string   `xml:"xmlns:xsi,attr"`
	Xsd     string   `xml:"xmlns:xsd,attr"`
	Soap12  string   `xml:"xmlns:soap12,attr"`
	Body    SoapBody `xml:"soap12:Body"`
}

type SoapBody struct {
	Body interface{}
}

func CreateSoapEnvelope() *SoapEnvelope {
	retval := &SoapEnvelope{}
	retval.Xsi = "http://www.w3.org/2001/XMLSchema-instance"
	retval.Xsd = "http://www.w3.org/2001/XMLSchema"
	retval.Soap12 = "http://www.w3.org/2003/05/soap-envelope"
	return retval
}

func GetDataString(mobileCode string, userID string) (string, error) {
	buffer := &bytes.Buffer{}
	requestEnvelope := CreateSoapEnvelope()
	info := GetMobileCodeInfo{}

	if mobileCode != "" {
		info.MobileCode = mobileCode
		info.UserID = userID
	}
	requestEnvelope.Body.Body = info
	encoder := xml.NewEncoder(buffer)
	err := encoder.Encode(requestEnvelope)
	if err != nil {
		println("Error encoding document:", err.Error())
		return "", err
	}
	fmt.Println("data", string(buffer.Bytes()))
	// FIXME: encoding
	client := http.Client{}
	req, err := http.NewRequest("POST", "http://webservice.webxml.com.cn/WebServices/MobileCodeWS.asmx", buffer)
	if err != nil {
		println("Error creating HTTP request:", err.Error())
		return "", err
	}

	req.Header.Add("Content-Type", "application/soap+xml; charset=utf-8")
	req.Header.Add("Content-Length", string(buffer.Len()))
	req.Host = "webservice.webxml.com.cn"
	resp, err := client.Do(req)
	if err != nil {
		println("Error POSTing HTTP request:", err.Error())
		return "", err
	}
	if resp.StatusCode != 200 {
		println("Error:", resp.Status)
		return "", err
	}
	// FIXME: check Content-Type
	// FIXME: encoding

	// responseEnvelope := SoapEnvelope{}
	bodyElement, err := DecodeResponseBody(resp.Body)
	if err != nil {
		println("Error decoding body:", err.Error())
		return "", err
	}

	if bodyElement == nil {
		println("Result is nil")
		return "", err
	}
	fmt.Printf("string: %#v\n", bodyElement.Value)
	return bodyElement.Value, nil
}

func DecodeResponseBody(body io.Reader) (*GetMobileCodeInfoResponse, error) {
	decoder := xml.NewDecoder(body)
	nextElementIsBody := false
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		switch startElement := token.(type) {
		case xml.StartElement:
			if nextElementIsBody {
				responseBody := GetMobileCodeInfoResponse{}
				err = decoder.DecodeElement(&responseBody, &startElement)
				if err != nil {
					return nil, err
				}
				return &responseBody, nil
			}
			if startElement.Name.Local == "Body" {
				nextElementIsBody = true
			}
		}
	}

	return nil, errors.New("Did not find SOAP body element")
}
