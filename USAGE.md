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

    ctx := context.Background()    
    req := operations.EnrichCompanyRequest{
        ID: "89bd9d8d-69a6-474e-8f46-7cc8796ed151",
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
<!-- End SDK Example Usage -->