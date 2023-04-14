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
        CountryCode: "MA",
        Headquarters: "provident",
        ID: "bd9d8d69-a674-4e0f-867c-c8796ed151a0",
        Name: "Estelle Will",
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