package qiita

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type Client struct {
	URLName string `json:"url_name"`
	Token   string
}

type RateLimit struct {
	Remaining int
	Limit     int
}

const (
	rootURL  = "https://qiita.com"
	basePath = "/api/v1"
)

func NewClient() *Client {
	return &Client{}
}

func NewClientWithToken(Token string) *Client {
	return &Client{
		Token: Token,
	}
}

func checkError(res []byte) error {
	var e interface{}
	if len(res) == 0 {
		return nil
	}
	err := json.Unmarshal(res, &e)
	if err != nil {
		return err
	}
	switch val := e.(type) {
	case map[string]interface{}:
		if v, ok := val["error"]; ok && v.(string) != "" {
			return errors.New(v.(string))
		}
	}
	return nil
}

func (c *Client) request(method, path string, params map[string]interface{}) ([]byte, error) {
	if c.Token != "" {
		params["token"] = c.Token
	}

	switch method {
	case "GET", "DELETE":
		parameters := make([]string, 0, len(params))
		for k, v := range params {
			switch v.(type) {
			case string:
				parameters = append(parameters, k+"="+v.(string))
			default:
				return nil, errors.New("method is " + method + ", but params contain non string elements")
			}
		}
		url := rootURL + basePath + path + "?" + strings.Join(parameters, "&")

		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			return nil, err
		}
		hc := &http.Client{}
		res, err := hc.Do(req)
		if err != nil {
			return nil, err
		}
		if res.StatusCode >= 400 {
			return nil, errors.New(res.Status)
		}
		return ioutil.ReadAll(res.Body)

	case "POST", "PUT":
		url := rootURL + basePath + path
		body, err := json.Marshal(params)
		if err != nil {
			return nil, err
		}
		req, err := http.NewRequest(method, url, bytes.NewReader(body))
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/json")
		hc := &http.Client{}
		res, err := hc.Do(req)
		if err != nil {
			return nil, err
		}
		if res.StatusCode >= 400 {
			return nil, errors.New(res.Status)
		}
		return ioutil.ReadAll(res.Body)

	}
	return nil, errors.New("invalid method")
}

func (c *Client) get(path string, params map[string]interface{}) ([]byte, error) {
	return c.request("GET", path, params)
}
func (c *Client) delete(path string, params map[string]interface{}) ([]byte, error) {
	return c.request("DELETE", path, params)
}
func (c *Client) post(path string, params map[string]interface{}) ([]byte, error) {
	return c.request("POST", path, params)
}
func (c *Client) put(path string, params map[string]interface{}) ([]byte, error) {
	return c.request("PUT", path, params)
}

func (c *Client) RateLimit() (ret *RateLimit, err error) {
	ret = new(RateLimit)
	res, err := c.get("/rate_limit", map[string]interface{}{})
	if err != nil {
		return
	}

	err = checkError(res)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, ret)
	if err != nil {
		return
	}
	return
}

func (c *Client) Login(username, password string) error {
	res, err := c.post("/auth", map[string]interface{}{"url_name": username, "password": password})
	if err != nil {
		return err
	}

	err = checkError(res)
	if err != nil {
		return err
	}

	err = json.Unmarshal(res, c)
	if err != nil {
		return err
	}

	return nil
}
