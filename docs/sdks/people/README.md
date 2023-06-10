# People

## Overview

People

### Available Operations

* [Enrich](#enrich) - Enrich a person profile
* [Search](#search) - Search People

## Enrich

Enrich a person profile

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
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.People.Enrich(ctx, operations.EnrichPersonRequest{
        ID: "d019da1f-fe78-4f09-bb00-74f15471b5e6",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Body != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.EnrichPersonRequest](../../models/operations/enrichpersonrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.EnrichPersonResponse](../../models/operations/enrichpersonresponse.md), error**


## Search

Search People

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
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.People.Search(ctx, operations.SearchPeopleApplicationJSON{
        Filter: sdk.String("repudiandae"),
        Limit: sdk.String("quae"),
        Page: sdk.String("ipsum"),
        Query: sdk.String("quidem"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Body != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.SearchPeopleApplicationJSON](../../models/operations/searchpeopleapplicationjson.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.SearchPeopleResponse](../../models/operations/searchpeopleresponse.md), error**

