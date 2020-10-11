package util

import (
	"bytes"
	"encoding/json"
	"goweb/log"
	"io/ioutil"
	"net/http"
	"time"
)

const TIMEOUT time.Duration = 10 * time.Second

// 发送GET请求
// url:请求地址
// response:请求返回的内容
func Get(url string) (string, error) {
	client := http.Client{Timeout: TIMEOUT}
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	return string(body), nil
}

// 发送POST请求
// url:请求地址，data:POST请求提交的数据,contentType:请求体格式，如：application/json
// content:请求放回的内容
func Post(url string, data interface{}, contentType string) (string, error) {
	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest(`POST`, url, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Error("get http new request error")
		return "", err
	}

	req.Header.Add(`content-type`, contentType)
	defer req.Body.Close()

	client := &http.Client{Timeout: TIMEOUT}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result), nil
}

func Put(url string, data interface{}, contentType string) (string, error) {
	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest(`PUT`, url, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Error("get http new request error")
		return "", err
	}

	req.Header.Add(`content-type`, contentType)
	defer req.Body.Close()

	client := &http.Client{Timeout: TIMEOUT}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result), nil
}
