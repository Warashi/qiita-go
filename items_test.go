package qiita

import (
	"testing"
)

func TestPostItem(t *testing.T) {
	c := NewClientWithToken(token)

	params := map[string]interface{}{}
	_, err := c.PostItem(params)
	if err == nil {
		t.Fail()
	}

	item, err := c.PostItem(map[string]interface{}{
		"title": "テスト",
		"tags": []map[string]string{{
			"name": "test",
		}},
		"body":    "テスト投稿",
		"private": true,
	})
	if err != nil {
		t.Error(err)
	}
	if item.Title != "テスト" {
		t.Fail()
	}
}

func TestUpdateItem(t *testing.T) {
	c := NewClientWithToken(token)

	uuid := ""
	params := map[string]interface{}{}
	_, err := c.UpdateItem(uuid, params)
	if err == nil {
		t.Fail()
	}

	params = map[string]interface{}{}
	_, err = c.UpdateItem(uuidForTest, params)
	if err == nil {
		t.Fail()
	}
}

func TestDeleteItem(t *testing.T) {
	c := NewClientWithToken(token)

	uuid := ""
	err := c.DeleteItem(uuid)
	if err == nil {
		t.Fail()
	}
}

func TestItem(t *testing.T) {
	c := NewClient()

	uuid := ""
	_, err := c.Item(uuid)
	if err == nil {
		t.Fail()
	}

	uuid = "b18a6f79ba42ad132764"
	ret, err := c.Item(uuid)
	if err != nil {
		t.Error(err)
	}
	if ret.Title != "テスト" {
		t.Fail()
	}
}

func TestItems(t *testing.T) {
	c := NewClient()

	params := map[string]interface{}{}
	_, err := c.Items(params)
	if err != nil {
		t.Error(err)
	}
}

func TestSearchItems(t *testing.T) {
	c := NewClient()

	query := ""
	params := map[string]interface{}{}
	_, err := c.SearchItems(query, params)
	if err == nil {
		t.Fail()
	}

	query = "Go"
	params = map[string]interface{}{}
	_, err = c.SearchItems(query, params)
	if err != nil {
		t.Error(err)
	}
}

func TestStockItem(t *testing.T) {
	c := NewClientWithToken(token)

	uuid := ""
	err := c.StockItem(uuid)
	if err == nil {
		t.Fail()
	}

	err = c.StockItem(uuidForTest)
	if err != nil {
		t.Error(err)
	}

}

func TestUnStockItem(t *testing.T) {
	c := NewClientWithToken(token)

	uuid := ""
	err := c.UnStockItem(uuid)
	if err == nil {
		t.Fail()
	}

	err = c.UnStockItem(uuidForTest)
	if err != nil {
		t.Error(err)
	}

}
