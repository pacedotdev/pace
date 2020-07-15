# Pace CLI (and Go client library)

## Pace CLI

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

## Go client

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
	client := pace.New(os.Getenv("PACE_API_KEY"))
	cardsService := pace.NewCardsService(client)
	resp, err := cardsService.GetCard(pace.CreateCardRequest{
		OrgID: "your-org-id",
		CardID: "12",
	})
	if err != nil {
		return err
	}
	fmt.Println("Card 12:", resp.Card.Title)
}
```
