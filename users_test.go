package qiita

import (
	"testing"
)

func TestUserItems(t *testing.T) {
	c := NewClient()

	username := ""
	params := map[string]interface{}{}
	_, err := c.UserItems(username, params)
	if err == nil {
		t.Fail()
	}

	username = "saveji"
	params = map[string]interface{}{}
	_, err = c.UserItems(username, params)
	if err != nil {
		t.Error(err)
	}
}

func TestMyItems(t *testing.T) {
	c := NewClientWithToken(token)

	params := map[string]interface{}{}
	_, err := c.MyItems(params)
	if err != nil {
		t.Error(err)
	}
}

func TestUserFollowingTags(t *testing.T) {
	c := NewClient()

	username := ""
	params := map[string]interface{}{}
	_, err := c.UserFollowingTags(username, params)
	if err == nil {
		t.Fail()
	}

	username = "saveji"
	params = map[string]interface{}{}
	_, err = c.UserFollowingTags(username, params)
	if err != nil {
		t.Error(err)
	}
}

func TestUserFollowingUsers(t *testing.T) {
	c := NewClient()

	username := ""
	params := map[string]interface{}{}
	_, err := c.UserFollowingUsers(username, params)
	if err == nil {
		t.Fail()
	}

	username = "saveji"
	params = map[string]interface{}{}
	_, err = c.UserFollowingUsers(username, params)
	if err != nil {
		t.Error(err)
	}
}

func TestUserStocks(t *testing.T) {
	c := NewClient()

	username := ""
	params := map[string]interface{}{}
	_, err := c.UserStocks(username, params)
	if err == nil {
		t.Fail()
	}

	username = "saveji"
	params = map[string]interface{}{}
	_, err = c.UserStocks(username, params)
	if err != nil {
		t.Error(err)
	}
}

func TestMyStocks(t *testing.T) {
	c := NewClientWithToken(token)

	params := map[string]interface{}{}
	_, err := c.MyStocks(params)
	if err != nil {
		t.Error(err)
	}
}

func TestUser(t *testing.T) {
	c := NewClient()

	username := ""
	params := map[string]interface{}{}
	_, err := c.User(username, params)
	if err == nil {
		t.Fail()
	}

	username = "saveji"
	params = map[string]interface{}{}
	_, err = c.User(username, params)
	if err != nil {
		t.Error(err)
	}
}

func TestMe(t *testing.T) {
	c := NewClientWithToken(token)

	params := map[string]interface{}{}
	_, err := c.Me(params)
	if err != nil {
		t.Error(err)
	}
}
