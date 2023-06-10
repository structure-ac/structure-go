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
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Companies.Enrich(ctx, operations.EnrichCompanyRequest{
        ID: "a05dfc2d-df7c-4c78-8a1b-a928fc816742",
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.EnrichCompanyRequest](../../models/operations/enrichcompanyrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.EnrichCompanyResponse](../../models/operations/enrichcompanyresponse.md), error**


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
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Companies.ListEmployees(ctx, operations.ListEmployeesRequest{
        ID: "cb739205-9293-496f-aa75-96eb10faaa23",
        Offset: sdk.String("corporis"),
        PerPage: sdk.String("explicabo"),
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListEmployeesRequest](../../models/operations/listemployeesrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.ListEmployeesResponse](../../models/operations/listemployeesresponse.md), error**


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
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Companies.ListJobs(ctx, operations.ListJobsRequest{
        ID: "c5955907-aff1-4a3a-afa9-467739251aa5",
        Offset: sdk.String("odit"),
        PerPage: sdk.String("quo"),
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

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.ListJobsRequest](../../models/operations/listjobsrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.ListJobsResponse](../../models/operations/listjobsresponse.md), error**


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
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Companies.Search(ctx, operations.SearchCompaniesApplicationJSON{
        Filter: sdk.String("sequi"),
        Limit: sdk.String("tenetur"),
        Page: sdk.String("ipsam"),
        Query: sdk.String("id"),
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.SearchCompaniesApplicationJSON](../../models/operations/searchcompaniesapplicationjson.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.SearchCompaniesResponse](../../models/operations/searchcompaniesresponse.md), error**

