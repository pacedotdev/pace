![](pace.dev-banner.jpg)

# Welcome to Pace for Developers. 

If you're new here, we recommend you take a few minutes to understand how Pace for Developers works:

* [Read about API keys in Pace](https://pace.dev/blog/2020/07/01/docs-api-keys.html)

### On this page

* [Pace CLI (command line interface)](#pace-cli)
* [Go client library](#go-client-library)
* [Pace JSON/HTTP API](#pace-jsonhttp-api)

# Pace CLI

The Pace command line interface provides programatic access to Pace.

```bash
pace help
pace list
pace templates
pace <service>.<method> -Flags=value
```

See the help:

```bash
$ pace help
```

Check the version:

```bash
$ pace version
```

* **What next?** To learn more, see the [Pace CLI documentation](https://pace.dev/app/docs/cli)

# Go client library

To use the Go client:

```bash
go get github.com/pacedotdev/pace
```

Then import it into your code:

* Use `pace.New(apikey)` to create a `pace.Client` with your API key ([Read more about API keys in Pace](https://pace.dev/blog/2020/07/01/docs-api-keys.html))
* Call the function to create the service you need (e.g. `pace.NewCardsService`) passing in the `pace.Client`
* Call the methods on that service to access the Pace API

Here is a complete example:

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
	secret := os.Getenv("PACE_API_SECRET")
	client := pace.New(apikey, secret)
	cardsService := pace.NewCardsService(client)
	
	resp, err := cardsService.GetCard(ctx, pace.CreateCardRequest{
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

You can make plain old HTTP calls with JSON payloads to interact with Pace.

* Make calls directly to `https://pace.dev/api`
* Set the `Content-Type` header to `application/json`
* Use `POST` method
* Set `X-API-KEY` header to your API key ([Read more about API keys in Pace](https://pace.dev/blog/2020/07/01/docs-api-keys.html))

```
POST https://pace.dev/api/CardsService.GetCard
Content-Type: application/json
X-API-KEY: your-api-key

{
	"OrgID": "your-org-id",
	"CardID": "12",
}
```

The response varies depending on which method you're calling, but there is always an `Error` string field which is empty (along with a `200` status code) if the operation was successful.
