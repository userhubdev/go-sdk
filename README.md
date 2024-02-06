# UserHub Go SDK

Stability: alpha

## Usage

```go
package main

import (
	"context"
	"fmt"

	"github.com/userhubdev/go-sdk/adminapi"
)

func main() {
	adminApi, err := adminapi.New("sk_123...")
	if err != nil {
		panic(err)
	}

	res, err := adminApi.Users().List(context.TODO(), &adminapi.UserListInput{
		PageSize: 5,
	})
	if err != nil {
		panic(err)
	}

	for _, user := range res.Users {
		fmt.Println(user.Id, user.DisplayName)
	}
}
```
