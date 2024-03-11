# addyapi

[Addy][1] [RESTful API][2] client library.

[![Go Reference](https://pkg.go.dev/badge/github.com/kovmir/addyapi.svg)](https://pkg.go.dev/github.com/kovmir/addyapi)

# INSTALL

Enter your project folder, and issue `go get`:

```bash
go get github.com/kovmir/addyapi
```

# USAGE

Go to your [account settings][2] to issue a new token, and then:

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

# DOCUMENTATION

The entire codebase resides within [`client.go`](client.go), the rest of the
files define methods and JSONs from [Addy API reference][2]. Each method is
more or less self-descriptive and has a URL pointing to the upstream reference.

[1]: https://addy.io/
[2]: https://app.addy.io/settings/api
[3]: https://app.addy.io/docs/#
