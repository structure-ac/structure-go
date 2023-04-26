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
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.EnrichPersonRequest{
        ID: "d019da1f-fe78-4f09-bb00-74f15471b5e6",
    }

    res, err := s.People.Enrich(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Body != nil {
        // handle response
    }
}
```

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
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.SearchPeopleApplicationJSON{
        Filter: sdk.String("repudiandae"),
        Limit: sdk.String("quae"),
        Page: sdk.String("ipsum"),
        Query: sdk.String("quidem"),
    }

    res, err := s.People.Search(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Body != nil {
        // handle response
    }
}
```
