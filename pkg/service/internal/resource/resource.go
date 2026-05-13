package resource

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// Resource encapsulates standard CRUD operations for a single top-level API resource.
// T is the model type; ID is the identifier type (int or string).
type Resource[T any, ID comparable] struct {
	conn       connection.Connection
	basePath   string
	name       string
	identifier string // word after "invalid {name}" in validation errors, e.g. "id" or "name"
	validate   func(ID) bool
	notFound   func(ID) error
}

// NewIntResource constructs a Resource with int IDs, valid when > 0.
func NewIntResource[T any](conn connection.Connection, basePath string, name string, notFound func(int) error) *Resource[T, int] {
	return &Resource[T, int]{
		conn:       conn,
		basePath:   basePath,
		name:       name,
		identifier: "id",
		validate:   func(id int) bool { return id > 0 },
		notFound:   notFound,
	}
}

// NewStringResource constructs a Resource with string IDs, valid when non-empty.
func NewStringResource[T any](conn connection.Connection, basePath string, name string, notFound func(string) error) *Resource[T, string] {
	return &Resource[T, string]{
		conn:       conn,
		basePath:   basePath,
		name:       name,
		identifier: "id",
		validate:   func(id string) bool { return id != "" },
		notFound:   notFound,
	}
}

// NewStringResourceWithIdentifier constructs a Resource with string IDs where the identifier
// word in validation errors differs from the default "id" (e.g. "name" for domain/zone resources).
func NewStringResourceWithIdentifier[T any](conn connection.Connection, basePath string, name string, identifier string, notFound func(string) error) *Resource[T, string] {
	return &Resource[T, string]{
		conn:       conn,
		basePath:   basePath,
		name:       name,
		identifier: identifier,
		validate:   func(id string) bool { return id != "" },
		notFound:   notFound,
	}
}

// List retrieves all items, fetching all pages automatically.
func (r *Resource[T, ID]) List(params connection.APIRequestParameters) ([]T, error) {
	return connection.InvokeRequestAll(r.ListPaginated, params)
}

// ListPaginated retrieves a single page of items.
func (r *Resource[T, ID]) ListPaginated(params connection.APIRequestParameters) (*connection.Paginated[T], error) {
	body, err := connection.Get[[]T](r.conn, r.basePath, params)
	return connection.NewPaginated(body, params, r.ListPaginated), err
}

// Get retrieves a single item by ID.
func (r *Resource[T, ID]) Get(id ID) (T, error) {
	var zero T
	if !r.validate(id) {
		return zero, fmt.Errorf("invalid %s %s", r.name, r.identifier)
	}
	body, err := connection.Get[T](r.conn, fmt.Sprintf("%s/%v", r.basePath, id), connection.APIRequestParameters{},
		connection.NotFoundResponseHandler(r.notFound(id)))
	return body.Data, err
}

// Create posts a new item and returns the created model.
func (r *Resource[T, ID]) Create(req any) (T, error) {
	body, err := connection.Post[T](r.conn, r.basePath, req)
	return body.Data, err
}

// Patch updates an existing item by ID.
func (r *Resource[T, ID]) Patch(id ID, req any) error {
	if !r.validate(id) {
		return fmt.Errorf("invalid %s %s", r.name, r.identifier)
	}
	return connection.PatchRaw(r.conn, fmt.Sprintf("%s/%v", r.basePath, id), req, &connection.APIResponseBody{},
		connection.NotFoundResponseHandler(r.notFound(id)))
}

// Delete removes an item by ID.
func (r *Resource[T, ID]) Delete(id ID) error {
	if !r.validate(id) {
		return fmt.Errorf("invalid %s %s", r.name, r.identifier)
	}
	return connection.DeleteRaw(r.conn, fmt.Sprintf("%s/%v", r.basePath, id), nil, &connection.APIResponseBody{},
		connection.NotFoundResponseHandler(r.notFound(id)))
}
