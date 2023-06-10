# Accounts

## Overview

Accounts

### Available Operations

* [ListUsers](#listusers) - Show current user accounts

## ListUsers

Show current user accounts

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
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Accounts.ListUsers(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Body != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListUsersResponse](../../models/operations/listusersresponse.md), error**

