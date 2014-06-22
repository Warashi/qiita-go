package qiita

import (
	"encoding/json"
)

type Tag struct {
	Name          string `json:"name"`
	URLName       string `json:"url_name"`
	IconURL       string `json:"icon_url"`
	ItemCount     int64  `json:"item_count"`
	FollowerCount int64  `json:"follower_count"`
	Following     bool   `json:"following"`
}

func (c *Client) TagItems(URLName string, params map[string]interface{}) (ret []Item, err error) {
	len := params["per_page"]
	switch len.(type) {
	case int:
		if len.(int) == 0 {
			ret = make([]Item, 20)
			break
		}
		ret = make([]Item, len.(int))
	default:
		ret = make([]Item, 20)
	}
	res, err := c.get("/tags/"+URLName+"/items", params)
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

func (c *Client) Tags(params map[string]interface{}) (ret []Tag, err error) {
	len := params["per_page"]
	switch len.(type) {
	case int:
		if len.(int) == 0 {
			ret = make([]Tag, 20)
			break
		}
		ret = make([]Tag, len.(int))
	default:
		ret = make([]Tag, 20)
	}
	res, err := c.get("/tags", params)
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
