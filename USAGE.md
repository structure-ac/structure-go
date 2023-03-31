<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "Structure"
    "Structure/pkg/models/shared"
    "Structure/pkg/models/operations"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    req := operations.EnrichCompanyRequest{
        CountryCode: "corrupti",
        Headquarters: "provident",
        ID: "distinctio",
        Name: "quibusdam",
    }

    ctx := context.Background()
    res, err := s.Companies.Enrich(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Body != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->