package qiita

import (
	"testing"
)

func TestTagItems(t *testing.T) {
	c := NewClient()

	params := map[string]interface{}{}
	_, err := c.TagItems("Go", params)
	if err != nil {
		t.Error(err)
	}

	params = map[string]interface{}{}
	_, err = c.TagItems("", params)
	if err == nil {
		t.Error(err)
	}
}

func TestTags(t *testing.T) {
	c := NewClient()
	params := map[string]interface{}{}
	_, err := c.Tags(params)
	if err != nil {
		t.Error(err)
	}
}
