package resty

import (
	"mathgpt/app/apiException"
	"sync"

	"github.com/go-resty/resty/v2"
)

var (
	client *resty.Client
	once   sync.Once
)

func initClient() {
	client = resty.New()
}

func GetClient() *resty.Client {
	once.Do(initClient)
	return client
}

func HttpSendPost(url string, req map[string]interface{}, headers map[string]string, resp interface{}) (*resty.Response, error) {
	client := GetClient()

	r, err := client.R().
		SetHeaders(headers).
		SetBody(req).
		SetResult(&resp).
		Post(url)
	if err != nil {
		return nil, apiException.RequestError
	}

	return r, nil
}

func HttpSendGet(url string, headers map[string]string, resp interface{}) (*resty.Response, error) {
	client := GetClient()

	r, err := client.R().
		SetHeaders(headers).
		SetResult(&resp).
		Get(url)
	if err != nil {
		return nil, apiException.RequestError
	}

	return r, nil
}
