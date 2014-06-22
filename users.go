package qiita

import (
	"encoding/json"
)

type User struct {
	Description string `json:"description"`
	Facebook    string `json:"facebook"`
	Teams       []struct {
		Name    string `json:"name"`
		URLName string `json:"url_name"`
	} `json:"teams"`
	Name            string `json:"name"`
	ProfileImageURL string `json:"profile_image_url"`
	URL             string `json:"url"`
	Organization    string `json:"organization"`
	Location        string `json:"location"`
	Github          string `json:"github"`
	Linkedin        string `json:"linkedin"`
	FollowingUsers  int64  `json:"following_users"`
	Items           int64  `json:"items"`
	URLName         string `json:"url_name"`
	WebsiteURL      string `json:"website_url"`
	Twitter         string `json:"twitter"`
	Followers       int64  `json:"followers"`
}

func (c *Client) UserItems(URLName string, params map[string]interface{}) (ret []Item, err error) {
	len := params["per_page"]
	switch len.(type) {
	case int:
		ret = make([]Item, len.(int))
	default:
		ret = make([]Item, 20)
	}
	var res []byte
	if URLName == "" {
		res, err = c.get("/items", params)
	} else {
		res, err = c.get("/users/"+URLName+"/items", params)

	}
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

func (c *Client) UserFollowingTags(URLName string, params map[string]interface{}) (ret []Tag, err error) {
	len := params["per_page"]
	switch len.(type) {
	case int:
		ret = make([]Tag, len.(int))
	default:
		ret = make([]Tag, 20)
	}
	res, err := c.get("/users/"+URLName+"/following_tags", params)
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

func (c *Client) UserFollowingUsers(URLName string, params map[string]interface{}) (ret []User, err error) {
	len := params["per_page"]
	switch len.(type) {
	case int:
		ret = make([]User, len.(int))
	default:
		ret = make([]User, 20)
	}
	res, err := c.get("/users/"+URLName+"/following_users", params)
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

func (c *Client) UserStocks(URLName string, params map[string]interface{}) (ret []Item, err error) {
	len := params["per_page"]
	switch len.(type) {
	case int:
		ret = make([]Item, len.(int))
	default:
		ret = make([]Item, 20)
	}
	var res []byte
	if URLName == "" {
		res, err = c.get("/stocks", params)
	} else {
		res, err = c.get("/users/"+URLName+"/stocks", params)

	}
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

func (c *Client) User(URLName string, params map[string]interface{}) (ret *User, err error) {
	ret = new(User)
	var res []byte
	if URLName == "" {
		res, err = c.get("/user", params)
	} else {
		res, err = c.get("/users/"+URLName, params)
	}
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
