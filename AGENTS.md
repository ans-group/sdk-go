# sdk-go

Go SDK for the ANS (formerly UKFast) REST APIs. Module `github.com/ans-group/sdk-go`. Consumed by the ANS [CLI](https://github.com/ans-group/cli) and Terraform providers, so exported API changes break downstream users.

## Modern Go

Invoke the `use-modern-go` skill before writing or editing any Go. The module's Go version is in `go.mod`; use syntax and stdlib that version offers.

## Layout

- `pkg/client`: `Client` interface and `NewClient(conn)`; one accessor per service (`SafeDNSService()`, ...). A new service needs an accessor here.
- `pkg/connection`: HTTP `Connection` interface, request/response handling, `APIRequestParameters` (filtering, sorting, pagination), `Paginated[T]`, credentials, config-driven `factory.go`.
- `pkg/config`: viper-based config (`$HOME/.ans`, env vars, contexts).
- `pkg/logging`, `pkg/ptr`: small helpers.
- `pkg/service/<name>`: one package per API. Conventional files: `service.go` (interface plus constructor), `service_<resource>.go` (+ `_test.go`), `model.go`, `request.go`, `error.go` (typed not-found errors).
- `pkg/service/internal/resource`: generic `Resource[T, ID]` CRUD helper with identifier-aware validation and not-found errors. Build new resource methods on it.
- `test/mocks`: generated gomock `Connection`. Regenerate with `go generate ./pkg/connection/...` (see the `go:generate` line in `pkg/connection/connection.go`).

## Commands

- Test: `go test ./...`
- Coverage: `cover.ps1`
- Format/vet: `gofmt` and `go vet ./...`

## Conventions

- Every service method has a test using the mocked `Connection` (gomock + testify), covering success, API error, and invalid-identifier cases.
- List methods come in pairs: `GetX` (all pages) and `GetXPaginated`.
- Mirror the API's field names and types from the OpenAPI spec in models; confirm against the spec before adding or changing a model, request, or endpoint.

## API reference

Docs index: <https://developers.ukfast.io/documentation>. Fetch the OpenAPI YAML for the service you are changing; it is the source of truth for paths, fields, and enums.

| Package | OpenAPI spec |
|---|---|
| `account` | https://developers.ukfast.io/api/documentation/account |
| `billing` | https://developers.ukfast.io/api/documentation/billing |
| `ddosx` | https://developers.ukfast.io/api/documentation/ddosx |
| `draas` | https://developers.ukfast.io/api/documentation/draas/2 (v1: `/1`) |
| `ecloud` | https://developers.ukfast.io/api/documentation/ecloud/2 (v1: `/1`) |
| `loadbalancer` | https://developers.ukfast.io/api/documentation/load-balancers |
| `cloudflare` | https://developers.ukfast.io/api/documentation/managed-cloudflare |
| `pss` | https://developers.ukfast.io/api/documentation/pss |
| `registrar` | https://developers.ukfast.io/api/documentation/domains/2 (v1: `/1`) |
| `safedns` | https://developers.ukfast.io/api/documentation/safedns |
| `ssl` | https://developers.ukfast.io/api/documentation/ssl |

Published specs without an SDK package: `datacentres`, `monitoring`, `threat-monitoring` (same URL pattern). `ecloudflex` and `sharedexchange` are out of service and slated for removal; leave them untouched and add nothing to them.

General API behaviour (auth, filtering, sorting, pagination, errors): <https://developers.ukfast.io/api-guide>.
