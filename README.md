# UserHub Go SDK

Stability: alpha

### Requirements

 * Go 1.20 or later

### Getting Started

Install SDK

```sh
go get -u github.com/userhubdev/go-sdk
```

Example

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

See the `examples` directory for more examples.
