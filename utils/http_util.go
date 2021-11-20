package utils

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/valyala/fasthttp"
)

var HttpUtil = new(HttpClient)

type HttpClient struct {
}

func (*HttpClient) HttpPost(url string, body interface{}) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(req)
	}()
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")

	req.SetRequestURI(url)
	requestBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req.SetBody(requestBody)

	if err := fasthttp.Do(req, resp); err != nil {
		return nil, err
	}
	return resp.Body(), nil
}

func (*HttpClient) HttpGet(url string, input map[string]interface{}) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(req)
	}()
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("GET")
	req.SetRequestURI(buildGetUrl(url, input))

	if err := fasthttp.Do(req, resp); err != nil {
		return nil, err
	}
	return resp.Body(), nil
}

func buildGetUrl(url string, input map[string]interface{}) string {
	params := ""
	for k, v := range input {
		params = fmt.Sprintf("%s%s=%v&", params, k, v)
	}
	paramsStr := strings.TrimRight(params, "&")
	return fmt.Sprintf("%s?%s", url, paramsStr)
}
