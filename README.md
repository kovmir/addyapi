# addyapi

[addy.io][1] REST API client library.

# STATUS

Implemented all HTTP `GET` methods.

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
	c := addyapi.NewClient("<your_token>")

	// Get token name.
	details, err := c.GetAPITokenDetails()
	if err != nil {
		// Handle error...
	}
	fmt.Println(details.Name)

	// Get first 5 active aliases.
	aliases, _ := c.GetAliases(&addyapi.AliasParams{
		Filter:   map[string]string{"active": "true"},
		PageSize: 5,
	})
	if err != nil {
		// Handle error...
	}
	for _, val := range aliases.Data {
		fmt.Println(val.Email)
	}
}
```

[1]: https://addy.io/
