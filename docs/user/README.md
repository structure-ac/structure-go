# User

## Overview

User

### Available Operations

* [Login](#login) - Login user
* [Me](#me) - Show current user

## Login

Login user

### Example Usage

```go
package main

import(
	"context"
	"log"
	"Structure"
	"Structure/pkg/models/operations"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.User.Login(ctx, operations.LoginApplicationJSON{
        Email: "Kenny50@yahoo.com",
        Password: "rem",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Body != nil {
        // handle response
    }
}
```

## Me

Show current user

### Example Usage

```go
package main

import(
	"context"
	"log"
	"Structure"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.User.Me(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Body != nil {
        // handle response
    }
}
```
