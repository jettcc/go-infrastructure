package infra

import (
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"

	"net/http"
	"strconv"
)

// 常用常量
const (
	CONTENT_TYPE      = "Content-Type"
	CONTENT_TYPE_JSON = "application/json"
	AUTHORIZATION     = "Authorization"
	BEARER            = "Bearer"
	STATUS_CODE       = "StatusCode"
	STATUS            = "Status"
)

type HttpClient struct {
	client *http.Client
}
type Response struct {
	Headers http.Header
	Data    []byte
}

func NewHttpClient() *HttpClient {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: tr,
	}

	return &HttpClient{client: client}
}

func (httpClient *HttpClient) Post(url string, header map[string]string, body io.Reader) (*Response, error) {
	return httpClient.Call(http.MethodPost, url, header, body)
}

func (httpClient *HttpClient) Get(url string, header map[string]string) (*Response, error) {
	return httpClient.Call(http.MethodGet, url, header, nil)
}

func (httpClient *HttpClient) Call(method, url string, header map[string]string, body io.Reader) (*Response, error) {
	req, err := http.NewRequest(method, url, body)

	if err != nil {
		return nil, err
	}

	if len(header) != 0 {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}

	resp, err := httpClient.client.Do(req)
	if err != nil {
		return nil, err
	}

	respHeader := resp.Header

	defer resp.Body.Close()

	respCode := resp.StatusCode
	respHeader.Set(STATUS_CODE, strconv.Itoa(respCode))
	respHeader.Set(STATUS, resp.Status)

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read from response'body error. response code %v. response content: %v",
			respCode, string(content))
	}

	return &Response{Headers: respHeader, Data: content}, nil
}
