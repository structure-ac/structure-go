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
            BearerAuth: shared.SchemeBearerAuth{
                Authorization: "Bearer YOUR_BEARER_TOKEN_HERE",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.Accounts.Accounts(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Body != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->