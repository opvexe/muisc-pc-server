package util

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
)

type Data struct {
	Code int8        `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

/*
	自定义http请求
 */
func CustomHttpRequest(method, url string, p interface{}) (interface{}, error) {

	req := &fasthttp.Request{}

	req.SetRequestURI(url)

	// 设置 body
	bytes, err := json.Marshal(p)
	if err != nil {

		return nil, err
	}
	req.SetBody(bytes)

	req.Header.SetContentType("application/json")
	req.Header.SetMethod(method)

	resp := &fasthttp.Response{}

	client := &fasthttp.Client{}

	if err := client.Do(req, resp); err != nil {
		return nil, err
	}

	var param Data

	err=json.Unmarshal(resp.Body(),&param)

	if err!=nil {
		return nil,err
	}

	return param, nil
}

func CustomHttpGetRequest(url string) (interface{}, error) {

	statusCode, body, err := fasthttp.Get(nil, url)

	if err != nil {
		return nil, err
	}

	if statusCode != fasthttp.StatusOK {
		return nil, err
	}

	var param Data

	err = json.Unmarshal(body, &param)
	if err != nil {
		return nil, err
	}

	return param.Data, nil
}
