# Companies

## Overview

Companies

### Available Operations

* [Enrich](#enrich) - Enrich a company profile
* [ListEmployees](#listemployees) - List company employees
* [ListJobs](#listjobs) - List company jobs
* [Search](#search) - Search Companies

## Enrich

Enrich a company profile

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
    req := operations.EnrichCompanyRequest{
        ID: "a05dfc2d-df7c-4c78-8a1b-a928fc816742",
    }

    res, err := s.Companies.Enrich(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Body != nil {
        // handle response
    }
}
```

## ListEmployees

List company employees

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
    req := operations.ListEmployeesRequest{
        ID: "cb739205-9293-496f-aa75-96eb10faaa23",
        Offset: sdk.String("corporis"),
        PerPage: sdk.String("explicabo"),
    }

    res, err := s.Companies.ListEmployees(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Body != nil {
        // handle response
    }
}
```

## ListJobs

List company jobs

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
    req := operations.ListJobsRequest{
        ID: "c5955907-aff1-4a3a-afa9-467739251aa5",
        Offset: sdk.String("odit"),
        PerPage: sdk.String("quo"),
    }

    res, err := s.Companies.ListJobs(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Body != nil {
        // handle response
    }
}
```

## Search

Search Companies

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
    req := operations.SearchCompaniesApplicationJSON{
        Filter: sdk.String("sequi"),
        Limit: sdk.String("tenetur"),
        Page: sdk.String("ipsam"),
        Query: sdk.String("id"),
    }

    res, err := s.Companies.Search(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Body != nil {
        // handle response
    }
}
```
