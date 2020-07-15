# Welcome to Pace for Developers. 

This repository contains:

* [Pace CLI](#pace-cli)
* [Go client library](#go-client-library)

![](pace.dev-banner.jpg)

# Pace CLI

To install the Pace CLI:

```bash
go install github.com/pacedotdev/pace/cmd/pace
```

Test it and see the help:

```bash
$ pace help

Usage:
  pace <service>.<method> [args...]

* CardsService - CardsService is used to work with cards.
* CommentsService

  pace help <service>[.<method>] - print specific help
  pace list - list all services and methods
  pace templates - show copy and paste examples

Flags:
  -apikey string
        Pace API Key
  -debug
        prints debug information
  -host string

```

# Go client library

To use the Go client:

```bash
go get github.com/pacedotdev/pace
```

Then import it into your code:

```go
package main

import (
	"context"
	
	"github.com/pacedotdev/pace"
)

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "err: %s\n", err)
	}
}

func run(ctx context.Context) error {
	apikey := os.Getenv("PACE_API_KEY")
	client := pace.New(apikey)
	cardsService := pace.NewCardsService(client)
	
	resp, err := cardsService.GetCard(pace.CreateCardRequest{
		OrgID:  "your-org-id",
		CardID: "12",
	})
	if err != nil {
		return err
	}

	fmt.Println("Card 12:", resp.Card.Title)
}
```
