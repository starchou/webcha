package utils

import (
	"net/http"
	"net/url"
)

func DoPost(httpurl string, urls map[string]string) (*http.Response, error) {
	client := http.Client{}
	v := FormString(urls)
	resp, err := client.PostForm(httpurl, v)
	return resp, err
}

func DoGet(httpurl string, urls map[string]string) (*http.Response, error) {
	client := http.Client{}
	queryurl := QueryString(urls)
	resp, err := client.Get(httpurl + queryurl)
	return resp, err
}

func FormString(urls map[string]string) url.Values {
	value := make(url.Values)
	for k, v := range urls {
		value.Add(k, v)
	}
	return value
}
func QueryString(urls map[string]string) string {
	var temp string
	for k, v := range urls {
		if temp == "" {
			temp = k + "=" + v
		} else {
			temp += "&" + k + "=" + v
		}
	}
	return temp
}
