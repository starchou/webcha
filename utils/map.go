package utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

//http://api.map.baidu.com/geocoder/v2/?ak=7a416a8357b78365225a2357e718556d&location=39.983424,116.322987&output=json&pois=0

const (
	baseurl  = "http://api.map.baidu.com/geocoder/v2/?"
	baiduKey = "7a416a8357b78365225a2357e718556d"
)

type BaiduMap struct {
	Status int
	Result Results
}
type Results struct {
	Location          Locations
	Formatted_address string
	Business          string
	AddressComponent  Address
	Pois              []Poises
	CityCode          int
}
type Locations struct {
	Lng float64
	Lat float64
}
type Address struct {
	City          string
	District      string
	Province      string
	Street        string
	Street_number string
}
type Poises struct {
	Addr     string
	Cp       string
	Distance string
	Name     string
	PoiType  string
	Point    Points
	Tel      string
	Uid      string
	Zip      string
}
type Points struct {
	X float64
	Y float64
}

func GetMap(location string) (*BaiduMap, error) {
	urls := make(map[string]string)
	urls["ak"] = baiduKey
	urls["output"] = "json"
	urls["location"] = location
	urls["pois"] = "1"
	resp, err := DoGet(baseurl, urls)
	if err != nil {
		println(err.Error())
	}
	baiduMap, err := JsonDecoder(resp.Body)
	return baiduMap, err
}

func JsonDecoder(body io.Reader) (*BaiduMap, error) {
	bodydata, err := ioutil.ReadAll(body)
	println(string(bodydata))
	var b *BaiduMap
	err = json.Unmarshal(bodydata, &b)
	if err != nil {
		return nil, err
	} else {
		return b, nil
	}
}
