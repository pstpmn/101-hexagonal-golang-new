package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"strings"
)

type IRequest interface {
	Json(method string, url string, header map[string]string, mapBody map[string]interface{}) (map[string]interface{}, error)
	List(method string, url string, header map[string]string, mapBody map[string]interface{}) ([]string, error)
	Request(method string, url string, header map[string]string, mapBody map[string]interface{}) (request *gentleman.Request)
	SendRequest(request *gentleman.Request) (string, error)
}

type requests struct {
}

func NewRequests() IRequest {
	return &requests{}
}

func (r requests) Json(method string, url string, header map[string]string, mapBody map[string]interface{}) (map[string]interface{}, error) {
	req := r.Request(method, url, header, mapBody)

	// send http receipt json
	res, err := r.SendRequest(req)

	var body map[string]interface{}
	json.Unmarshal([]byte(res), &body)

	// validated body
	if len(body) == 0 {
		return nil, errors.New("invalid Information Not Json")
	}
	return body, err
}

func (r requests) List(method string, url string, header map[string]string, mapBody map[string]interface{}) ([]string, error) {
	req := r.Request(method, url, header, mapBody)

	// send http receipt json
	res, err := r.SendRequest(req)

	var body []string
	json.Unmarshal([]byte(res), &body)

	// validated body
	if len(body) == 0 {
		return nil, errors.New("invalid Information Not List")
	}
	return body, err
}

func (r requests) Request(method string, url string, header map[string]string, mapBody map[string]interface{}) (request *gentleman.Request) {
	methodUpperCase := strings.ToUpper(method)
	cli := gentleman.New()
	// case send body for method POST , PUT , PATCH
	if methodUpperCase != "DELETE" && methodUpperCase != "GET" {
		cli.Use(body.JSON(mapBody))
	}

	cli.URL(url)
	req := cli.Request()
	req.SetHeaders(header)
	req.Method(methodUpperCase)
	return req
}

func (r requests) SendRequest(request *gentleman.Request) (string, error) {
	res, err := request.Send()
	if err != nil {
		return "", errors.New(fmt.Sprintf("request error: %s\n", err))
	}
	if !res.Ok {
		return "", errors.New(fmt.Sprintf("invalid response: %d\n", res.StatusCode))
	}
	return res.String(), nil
}
