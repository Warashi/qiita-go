package qiita

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestCheckError(t *testing.T) {
	testDATAs := map[string]error{
		"error": errors.New("error"),
		"":      nil,
	}
	for k, v := range testDATAs {
		testDATA, err := json.Marshal(map[string]interface{}{"error": k})
		if err != nil {
			t.Error(err)
		}
		err = checkError(testDATA)
		if err != v && err.Error() != v.Error() {
			t.Error(err, " != ", v)
		}
	}
}

func TestRequest(t *testing.T) {
	c := NewClient()
	_, err := c.request("", "", map[string]interface{}{})
	if err == nil {
		t.Error("method invalid, but not error")
	}
}

func TestRateLimit(t *testing.T) {
	c := NewClient()
	ret, err := c.RateLimit(map[string]interface{}{})
	if err != nil {
		t.Error(err)
	}

	if ret.Limit == 0 {
		t.Error("Limit is 0")
	}
}

func TestLogin(t *testing.T) {
	c := NewClient()
	err := c.Login(username, password)
	if err != nil {
		t.Error(err)
	}

	c = NewClient()
	err = c.Login(username, "")
	if err == nil {
		t.Fail()
	}

	c = NewClient()
	err = c.Login("", password)
	if err == nil {
		t.Fail()
	}

	c = NewClient()
	err = c.Login("", "")
	if err == nil {
		t.Fail()
	}
}
