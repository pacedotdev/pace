# Welcome to Pace for Developers. 

Learn about how Pace for Developers works:

* [Read about API keys in Pace](https://pace.dev/blog/2020/07/01/docs-api-keys.html)

## On this page

* [Pace CLI (command line interface)](#pace-cli)
* [Go client library](#go-client-library)
* [Pace JSON/HTTP API](#pace-json-http-api)

![](pace.dev-banner.jpg)

# Pace CLI

To install the Pace command line interface:

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

# Pace JSON/HTTP API

You can make HTTP calls directly to the `https://pace.dev/api` endpoint by encoding the request fields as a JSON object in the body of a `POST` request:

* You will need an API key ([Read about API keys in Pace](https://pace.dev/blog/2020/07/01/docs-api-keys.html))

```
Content-Type: application/json
X-API-KEY: your-api-key
POST https://pace.dev/api/CardsService.GetCard
{
	"OrgID": "your-org-id",
	"CardID": "12",
}
```

