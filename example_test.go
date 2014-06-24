package qiita

import (
	"fmt"
)

func ExampleRateLimit() {
	c := NewClient()
	ret, _ := c.RateLimit()
	fmt.Println(ret.Remaining)
}

func ExampleLogin() {
	c := NewClient()
	_ = c.Login("username", "password")
}

func ExamplePostItem() {
	c := NewClientWithToken("token")
	param := map[string]interface{}{
		"title": "Title",
		"tags": []map[string]interface{}{
			{"name": "Tag1", "versions": []string{"1.1", "1.2"}},
			{"name": "Tag2", "versions": []string{"1.3", "1.4"}},
		},
		"body":    "body",
		"private": false,
		"gist":    true,
		"tweet":   true,
	}
	_, _ = c.PostItem(param)
}
