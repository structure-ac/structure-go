# Account


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `AccountUsers`                                             | []*string*                                                 | :heavy_check_mark:                                         | List of the users in the account                           |
| `CreatedAt`                                                | *string*                                                   | :heavy_check_mark:                                         | Timestamp of the account creation date                     |
| `ID`                                                       | *float64*                                                  | :heavy_check_mark:                                         | The id of the account                                      |
| `Name`                                                     | *string*                                                   | :heavy_check_mark:                                         | The name of the account                                    |
| `OwnerID`                                                  | *float64*                                                  | :heavy_check_mark:                                         | The id of the account owner                                |
| `Personal`                                                 | *bool*                                                     | :heavy_check_mark:                                         | If the account is personal or belonging to another account |
| `UpdatedAt`                                                | *string*                                                   | :heavy_check_mark:                                         | Timestamp of the last account update date                  |