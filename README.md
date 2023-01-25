# sdk-go

This is an SDK for interacting with ANS APIs from Go applications.

You should refer to the [Getting Started](https://developers.ukfast.io/getting-started) section of the API documentation before proceeding below

## Basic usage

First we'll instantiate a Client struct with an API key:

```go
c := client.NewClient(connection.NewAPIKeyCredentialsAPIConnection("myapikey"))
```

And away we go:

```go
service := c.SafeDNSService()
zone, err := service.GetZone("ans.co.uk")

fmt.Printf("Zone: %s", zone.Name)
```

## Services

Resources/models are separated into separate service packages, found within `pkg/service`.

## Config

The SDK has default implementation for managing config, which is utilised by several utilities such as the [CLI](https://github.com/ans-group/cli) and Terraform providers. This config can be defined both within config files and environment variables.

There is a default connection factory included (`pkg/connection/factory.go`) which utilises the config package, which can be used in your applications as below:

```go
conn, err := connection.NewDefaultConnectionFactory().NewConnection()
if err != nil {
    panic(err)
}

service := client.NewClient(conn).SafeDNSService()
```

### Configuration File

The configuration file is read from
`$HOME/.ans{.extension}` by default (extension being one of the `viper` supported formats such as `yml`, `yaml`, `json`, `toml` etc.)

### Schema

* `api_key`: (String) *Required* API key for authenticating with API
* `api_timeout_seconds`: (int) HTTP timeout for API requests. Default: `90`
* `api_uri`: (string) API URI. Default: `api.ukfast.io`
* `api_insecure`: (bool) Specifies to ignore API certificate validation checks

### Contexts

Contexts can be defined in the config file to allow for different sets of configuration to be defined:

```yaml
contexts:
  testcontext1:
    api_key: mykey1
  testcontext2:
    api_key: mykey2
current_context: testcontext1
```

### Environment variables

These variables match the naming of directives in the configuration file defined above, however are uppercased and prefixed with `ANS_`, such as `ANS_API_KEY`

### Precedence

Values defined in the configuration file take precedence over environment variables
