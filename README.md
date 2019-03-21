# sdk-go

This is an SDK for interacting with UKFast APIs from Go applications.

You should refer to the [Getting Started](https://developers.ukfast.io/getting-started) section of the API documentation before proceeding below

### Basic usage

First we'll instantiate a Client struct with an API key:

```
ukfclient := client.NewClient(connection.NewAPIKeyCredentialsConnection("myapikey"))
```

And away we go:

```
service := ukfclient.SafeDNSService()
zone, err := service.GetZone("ukfast.co.uk")
...
```

### Services

Resources/models are separated into separate service packages, found within `pkg/service`.
There are currently 5 services available:

- Account
- DDoSX
- eCloud
- SafeDNS
- SSL
- Registrar
