# Qalpuch API - Go SDK

Official Go SDK for the Qalpuch REST API.

This SDK provides an idiomatic Go interface for all the Qalpuch API resources, including Users, Files, Tasks, and Workers.

## Installation

To use this SDK in your project, you can add it to your `go.mod` file:

```bash
go get github.com/Q300Z/qalpuch_api/go_sdk_qalpuch_api
```

## Usage

Here is a basic example of how to use the client:

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Q300Z/qalpuch_api/go_sdk_qalpuch_api/pkg/clients"
)

func main() {
	// Create a new client
	client := clients.NewUserClient("https://api.qalpuch.cc", "YOUR_API_TOKEN")

	// Get the list of users
	users, err := client.Users.GetUsers(context.Background())
	if err != nil {
		log.Fatalf("Failed to get users: %v", err)
	}

	fmt.Println("Users:", users)
}
```

## Architecture

The SDK follows the Standard Go Project Layout.
- `pkg/`: Public modules exposed by the SDK (clients, models, errors).
- `pkg/services/`: Service interfaces defining the API contracts.
- `pkg/clients/`: Client implementations for each API resource.
- `pkg/models/`: Go structs representing the API entities.
- `examples/`: Usage examples.

For more details on the API, please refer to the `docs/api_reference.md` file.
