package qiita

import (
	"encoding/json"
)

type Item struct {
	ID        int64  `json:"id"`
	UpdatedAt string `json:"updated_at"`
	Tags      []struct {
		Name     string   `json:"name"`
		URLName  string   `json:"url_name"`
		IconURL  string   `json:"icon_url"`
		Versions []string `json:"versions"`
	} `json:"tags"`
	StockUsers   []string `json:"stock_users"`
	CommentCount int64    `json:"comment_count"`
	URL          string   `json:"url"`
	GistURL      string   `json:"gist_url"`
	Tweet        bool     `json:"tweet"`
	Title        string   `json:"title"`
	StockCount   int64    `json:"stock_count"`
	Comments     []struct {
		ID   int64  `json:"id"`
		Uuid string `json:"uuid"`
		User struct {
			Name            string `json:"name"`
			URLName         string `json:"url_name"`
			ProfileImageURL string `json:"profile_image_url"`
		} `json:"user"`
		Body string `json:"body"`
	} `json:"comments"`
	Uuid string `json:"uuid"`
	User struct {
		Name            string `json:"name"`
		URLName         string `json:"url_name"`
		ProfileImageURL string `json:"profile_image_url"`
	} `json:"user"`
	UpdatedAtInWords string `json:"updated_at_in_words"`
	Private          bool   `json:"private"`
	Stocked          bool   `json:"stocked"`
	Body             string `json:"body"`
	CreatedAt        string `json:"created_at"`
	CreatedAtInWords string `json:"created_at_in_words"`
}

func (c *Client) PostItem(params map[string]interface{}) (ret *Item, err error) {
	ret = new(Item)
	res, err := c.post("/items", params)
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

func (c *Client) UpdateItem(uuid string, params map[string]interface{}) (ret *Item, err error) {
	ret = new(Item)
	res, err := c.put("/items/"+uuid, map[string]interface{}{})
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

func (c *Client) DeleteItem(uuid string) error {
	res, err := c.delete("/items/"+uuid, map[string]interface{}{})
	if err != nil {
		return err
	}

	err = checkError(res)
	return err
}

func (c *Client) Item(uuid string) (ret *Item, err error) {
	ret = new(Item)
	res, err := c.get("/items/"+uuid, map[string]interface{}{})
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

func (c *Client) Items(params map[string]interface{}) (ret []Item, err error) {
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
	res, err := c.get("/items", params)
	if err != nil {
		return
	}

	err = checkError(res)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &ret)
	if err != nil {
		return
	}
	return
}

func (c *Client) SearchItems(query string, params map[string]interface{}) (ret []Item, err error) {
	params["q"] = query
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
	res, err := c.get("/search", params)
	if err != nil {
		return
	}

	err = checkError(res)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &ret)
	if err != nil {
		return
	}
	return
}

func (c *Client) StockItem(uuid string) error {
	res, err := c.put("/items/"+uuid+"/stock", map[string]interface{}{})
	if err != nil {
		return err
	}

	err = checkError(res)
	return err
}

func (c *Client) UnStockItem(uuid string) error {
	res, err := c.delete("/items/"+uuid+"/stock", map[string]interface{}{})
	if err != nil {
		return err
	}

	err = checkError(res)
	return err
}
