# sdk-go

The Go SDK for the [ANS APIs](https://developers.ukfast.io/documentation).
See the API [getting started guide](https://developers.ukfast.io/getting-started)
for authentication and access to the APIs.

## Installation

The current module requires Go 1.26 or later.

```sh
go get github.com/ans-group/sdk-go
```

## Get started

Set `ANS_API_KEY` in your environment, then use a client to access a service.
Replace `example.com` with a SafeDNS zone you have access to.

```go
package main

import (
	"fmt"
	"log"

	"github.com/ans-group/sdk-go/pkg/client"
	"github.com/ans-group/sdk-go/pkg/config"
	"github.com/ans-group/sdk-go/pkg/connection"
)

func main() {
	if err := config.Init(""); err != nil {
		log.Fatal(err)
	}

	conn, err := connection.NewDefaultConnectionFactory().NewConnection()
	if err != nil {
		log.Fatal(err)
	}

	zone, err := client.NewClient(conn).SafeDNSService().GetZone("example.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(zone.Name)
}
```

The factory returns an error if no API key is configured. Keep API keys out
of source control and use a key with access to the services you need.

If you manage configuration yourself, you can skip `config.Init` and the
factory:

```go
conn := connection.NewAPIKeyCredentialsAPIConnection(apiKey)
c := client.NewClient(conn)
```

The direct constructor does not check whether the key is empty. It uses
HTTPS, `api.ukfast.io` and a 120-second HTTP timeout by default.
The snippets below assume `c` is a client created in this way or with the
factory. Add the imports used by each snippet to your program.

## Services

Use the accessors on `client.Client` to obtain a service. The SDK includes
account, billing, Cloudflare, DRaaS, eCloud, load balancer, PSS, registrar,
SafeDNS and SSL services. Models, request types and service interfaces are in
`pkg/service/<name>`. Available methods and return values differ by service.

## Listing and pagination

For list methods that provide both forms, `GetX` fetches all pages and returns
a slice. It starts at page 1 even if `Pagination.Page` is set. Use
`GetXPaginated` to request one page at a time:

```go
params := connection.APIRequestParameters{
	Pagination: connection.APIRequestPagination{PerPage: 100},
}

page, err := c.SafeDNSService().GetZonesPaginated(params)
if err != nil {
	log.Fatal(err)
}
for page != nil {
	for _, zone := range page.Items() {
		fmt.Println(zone.Name)
	}
	page, err = page.Next()
	if err != nil {
		log.Fatal(err)
	}
}
```

`Next` returns `nil, nil` on the last page. `Previous` does the same on the
first page. `First`, `Last`, `CurrentPage`, `TotalPages` and `Total` are also
available. Page navigation makes another API request, so check the error
before using the returned page.

Filters and sorting are passed through `APIRequestParameters`. For example:

```go
params := connection.NewAPIRequestParameters().
	WithFilter(connection.APIRequestFiltering{
		Property: "name",
		Operator: connection.EQOperator,
		Value:    []string{"example.com"},
	}).
	WithSorting(connection.APIRequestSorting{Property: "name"})
zones, err := c.SafeDNSService().GetZones(*params)
```

The API determines which properties, operators and sort fields are supported
for each endpoint. See the [API guide](https://developers.ukfast.io/api-guide)
and the service's API documentation.

## Requests and errors

Create and update methods generally take request structs from the service
package. For example, to create a SafeDNS zone:

```go
err := c.SafeDNSService().CreateZone(safedns.CreateZoneRequest{
	Name:        "example.com",
	Description: "Example zone",
})
```

This makes a write request. Use a zone name appropriate for your account.
Some request types validate required fields before sending a request. Optional
fields with `omitempty` are omitted when zero-valued; pointer fields can
distinguish an omitted value from an explicit zero or `false`. Check the
request type's JSON tags when updating a resource.

Always check errors before using returned values. Errors can come from local
validation, the HTTP client, response decoding or the API. Some methods return
service-specific not-found errors:

```go
zone, err := c.SafeDNSService().GetZone("example.com")
if err != nil {
	if notFound, ok := errors.AsType[*safedns.ZoneNotFoundError](err); ok {
		log.Printf("zone %s was not found", notFound.ZoneName)
	} else {
		log.Print(err)
	}
	return
}
fmt.Println(zone.Name)
```

Other errors may wrap an API response body with a message and error details.
Not every failure contains an API response, so retain a fallback for other
errors.

## Configuration

Call `config.Init("")` before using the default connection factory to enable
environment variable lookup and search for `$HOME/.ans` with a Viper-supported
extension, such as `.yml`, `.yaml`, `.json` or `.toml`. The default search
allows the file to be absent. Pass a path to `config.Init(path)` to use a
specific file; an error is returned if it cannot be read.

The factory reads these settings:

| Setting | Purpose |
| --- | --- |
| `api_key` | API key, required by the factory. |
| `api_uri` | API host, default `api.ukfast.io`. Supply a host, not a URL with a scheme. |
| `api_timeout_seconds` | HTTP timeout in seconds. A positive value overrides the 120-second default. |
| `api_insecure` | If true, disables TLS certificate verification. Leave this false except in controlled environments where you accept the risk. |
| `api_headers` | Map of additional HTTP headers. These are appended to the standard headers. |

Top-level settings can also be supplied as uppercase environment variables
prefixed with `ANS_`, such as `ANS_API_KEY` and `ANS_API_TIMEOUT_SECONDS`.
For the same key, environment variables take precedence over file values.

Contexts group settings in the config file:

```yaml
api_timeout_seconds: 30
current_context: production
contexts:
  production:
    api_key: replace-with-your-api-key
```

For each setting, the SDK first checks the selected context's key and falls
back to the top-level key if the context has no value for it. Consequently, a
context's `api_key` takes precedence over a top-level `ANS_API_KEY`. Protect
configuration files containing keys and avoid committing them.

## Operational notes

- Service methods do not accept `context.Context`. Set the HTTP client timeout
  on a direct connection, or use `api_timeout_seconds` with the factory.
- The SDK logs nothing unless a logger is registered with `logging.SetLogger`.
  Debug logging includes URLs, headers and response bodies; trace logging can
  include request bodies. Authorization headers are partially redacted, but
  logs may still contain sensitive data.
- For direct use of `connection.Connection`, handle errors and close response
  bodies. The service helpers deserialize and close bodies for you.

## Development

Run the tests and vet checks with:

```sh
go test ./...
go vet ./...
```
