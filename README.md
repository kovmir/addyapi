# addyapi

[addy.io][1] RESTful API client library.

# INSTALL

Enter your project folder, and issue `go get`:

```bash
go get github.com/kovmir/addyapi
```

# USAGE

```go
package main

import (
	"fmt"

	"github.com/kovmir/addyapi"
)

func main() {
	c := addyapi.NewClient("YOUR_TOKEN")

	// Get token name.
	details, err := c.TokenGetAPIDetails()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Token name: %s\n", details.Name)

	// Get first 5 active aliases.
	aliases, err := c.AliasesGet(&addyapi.AliasesGetArgs{
		Filter:   map[string]string{"active": "true"},
		PageSize: 5,
	})
	if err != nil {
		panic(err)
	}
	for i, v := range aliases.Data {
		fmt.Printf("%d. %s\n", i, v.Email)
	}

	// Create a new UUID alias.
	alias, err := c.AliasNew(&addyapi.AliasNewArgs{
		Desc:   "addy client test",
		Domain: "mailer.me",
		Format: addyapi.AliasFmtUUID,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("alias %s created successfully\n", alias.Data.ID)
}
```

[1]: https://addy.io/
