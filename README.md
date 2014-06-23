# Qiita

Go wrapper for Qiita API v1.

## Installation

```sh
go get -u github.com/Warashi/qiita-go
```
and
```go
import "github.com/Warashi/qiita-go"
```


## Usage

### Get user's items
```go
c := qiita.NewClient()
UserName := "saveji"
params := map[string]interface{}{}
items,err := c.UserItems(UserName,params)
```

### Get tag's items
```go
c := qiita.NewClient()
tag := "Go"
params := map[string]interface{}{}
items,err := c.TagItems(tag,params)
```

### Get a specified item with comments and raw markdown content
```go
c := qiita.NewClient()
uuid = "1234567890abcdefg"
item,err := c.Item(uuid)
```


## Authenticated requests

### Login with "username & password" or "token"
```go
c := qiita.NewClient()
err := c.Login("UserName","Password")
```
or
```go
c := qiita.NewClientWithToken("Token")
```

### Get my items
```go
params := map[string]interface{}{}
items,err := c.MyItems(params)
```

### Post/Update/Delete an item
post
```go
params := map[string]interface{}{
			"title": "Hello",
			"tags": []map[string]interface{}{{
				"name": "Tag",
                "versions": []string{"1.1","1.2"}
			}},
			"body":    "markdown text",
			"private": false,
}
item,err := c.PostItem(param)
```
update
```go
params := map[string]interface{}{
			"title": "modified",
}
item,err := c.UpdateItem("uuid",param)
```
delete
```go
err := c.DeleteItem("uuid")
```
