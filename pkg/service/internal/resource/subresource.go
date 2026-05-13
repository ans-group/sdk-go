package resource

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// SubResourceList handles list-only operations for resources nested under a parent.
type SubResourceList[T any, ParentID comparable] struct {
	conn           connection.Connection
	basePath       func(ParentID) string
	name           string
	identifier     string
	validateParent func(ParentID) bool
	parentNotFound func(ParentID) error
}

// NewStringSubResourceList creates a SubResourceList with a string parent ID (valid when non-empty).
func NewStringSubResourceList[T any](conn connection.Connection, basePath func(string) string, name, identifier string, parentNotFound func(string) error) *SubResourceList[T, string] {
	return &SubResourceList[T, string]{
		conn: conn, basePath: basePath, name: name, identifier: identifier,
		validateParent: func(id string) bool { return id != "" },
		parentNotFound: parentNotFound,
	}
}

// NewIntSubResourceList creates a SubResourceList with an int parent ID (valid when > 0).
func NewIntSubResourceList[T any](conn connection.Connection, basePath func(int) string, name, identifier string, parentNotFound func(int) error) *SubResourceList[T, int] {
	return &SubResourceList[T, int]{
		conn: conn, basePath: basePath, name: name, identifier: identifier,
		validateParent: func(id int) bool { return id > 0 },
		parentNotFound: parentNotFound,
	}
}

// NewUncheckedStringSubResourceList creates a SubResourceList with no parent validation or not-found handler.
func NewUncheckedStringSubResourceList[T any](conn connection.Connection, basePath func(string) string) *SubResourceList[T, string] {
	return &SubResourceList[T, string]{conn: conn, basePath: basePath}
}

func (r *SubResourceList[T, ParentID]) List(parentID ParentID, params connection.APIRequestParameters) ([]T, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[T], error) {
		return r.ListPaginated(parentID, p)
	}, params)
}

func (r *SubResourceList[T, ParentID]) ListPaginated(parentID ParentID, params connection.APIRequestParameters) (*connection.Paginated[T], error) {
	if r.validateParent != nil && !r.validateParent(parentID) {
		return nil, fmt.Errorf("invalid %s %s", r.name, r.identifier)
	}
	var handlers []connection.ResponseHandler
	if r.parentNotFound != nil {
		handlers = append(handlers, connection.NotFoundResponseHandler(r.parentNotFound(parentID)))
	}
	body, err := connection.Get[[]T](r.conn, r.basePath(parentID), params, handlers...)
	return connection.NewPaginated(body, params, func(p connection.APIRequestParameters) (*connection.Paginated[T], error) {
		return r.ListPaginated(parentID, p)
	}), err
}

// SubResource handles CRUD operations for resources nested under a parent.
type SubResource[T any, ParentID comparable, ChildID comparable] struct {
	conn             connection.Connection
	basePath         func(ParentID) string
	parentName       string
	parentIdentifier string
	childName        string
	childIdentifier  string
	validateParent   func(ParentID) bool
	validateChild    func(ChildID) bool
	parentNotFound   func(ParentID) error
	childNotFound    func(ParentID, ChildID) error
}

// NewStringIntSubResource creates a SubResource with a string parent ID and int child ID.
func NewStringIntSubResource[T any](
	conn connection.Connection,
	basePath func(string) string,
	parentName, parentIdentifier string,
	parentNotFound func(string) error,
	childName, childIdentifier string,
	childNotFound func(string, int) error,
) *SubResource[T, string, int] {
	return &SubResource[T, string, int]{
		conn: conn, basePath: basePath,
		parentName: parentName, parentIdentifier: parentIdentifier,
		childName: childName, childIdentifier: childIdentifier,
		validateParent: func(id string) bool { return id != "" },
		validateChild:  func(id int) bool { return id > 0 },
		parentNotFound: parentNotFound,
		childNotFound:  childNotFound,
	}
}

// NewIntIntSubResource creates a SubResource with int parent and child IDs.
func NewIntIntSubResource[T any](
	conn connection.Connection,
	basePath func(int) string,
	parentName, parentIdentifier string,
	parentNotFound func(int) error,
	childName, childIdentifier string,
	childNotFound func(int, int) error,
) *SubResource[T, int, int] {
	return &SubResource[T, int, int]{
		conn: conn, basePath: basePath,
		parentName: parentName, parentIdentifier: parentIdentifier,
		childName: childName, childIdentifier: childIdentifier,
		validateParent: func(id int) bool { return id > 0 },
		validateChild:  func(id int) bool { return id > 0 },
		parentNotFound: parentNotFound,
		childNotFound:  childNotFound,
	}
}

// NewStringStringSubResource creates a SubResource with string parent and child IDs.
func NewStringStringSubResource[T any](
	conn connection.Connection,
	basePath func(string) string,
	parentName, parentIdentifier string,
	parentNotFound func(string) error,
	childName, childIdentifier string,
	childNotFound func(string, string) error,
) *SubResource[T, string, string] {
	return &SubResource[T, string, string]{
		conn: conn, basePath: basePath,
		parentName: parentName, parentIdentifier: parentIdentifier,
		childName: childName, childIdentifier: childIdentifier,
		validateParent: func(id string) bool { return id != "" },
		validateChild:  func(id string) bool { return id != "" },
		parentNotFound: parentNotFound,
		childNotFound:  childNotFound,
	}
}

func (r *SubResource[T, ParentID, ChildID]) List(parentID ParentID, params connection.APIRequestParameters) ([]T, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[T], error) {
		return r.ListPaginated(parentID, p)
	}, params)
}

func (r *SubResource[T, ParentID, ChildID]) ListPaginated(parentID ParentID, params connection.APIRequestParameters) (*connection.Paginated[T], error) {
	if r.validateParent != nil && !r.validateParent(parentID) {
		return nil, fmt.Errorf("invalid %s %s", r.parentName, r.parentIdentifier)
	}
	var handlers []connection.ResponseHandler
	if r.parentNotFound != nil {
		handlers = append(handlers, connection.NotFoundResponseHandler(r.parentNotFound(parentID)))
	}
	body, err := connection.Get[[]T](r.conn, r.basePath(parentID), params, handlers...)
	return connection.NewPaginated(body, params, func(p connection.APIRequestParameters) (*connection.Paginated[T], error) {
		return r.ListPaginated(parentID, p)
	}), err
}

func (r *SubResource[T, ParentID, ChildID]) Get(parentID ParentID, childID ChildID) (T, error) {
	var zero T
	if r.validateParent != nil && !r.validateParent(parentID) {
		return zero, fmt.Errorf("invalid %s %s", r.parentName, r.parentIdentifier)
	}
	if r.validateChild != nil && !r.validateChild(childID) {
		return zero, fmt.Errorf("invalid %s %s", r.childName, r.childIdentifier)
	}
	body, err := connection.Get[T](r.conn, fmt.Sprintf("%s/%v", r.basePath(parentID), childID), connection.APIRequestParameters{},
		connection.NotFoundResponseHandler(r.childNotFound(parentID, childID)))
	return body.Data, err
}

func (r *SubResource[T, ParentID, ChildID]) Create(parentID ParentID, req any) (T, error) {
	var zero T
	if r.validateParent != nil && !r.validateParent(parentID) {
		return zero, fmt.Errorf("invalid %s %s", r.parentName, r.parentIdentifier)
	}
	var handlers []connection.ResponseHandler
	if r.parentNotFound != nil {
		handlers = append(handlers, connection.NotFoundResponseHandler(r.parentNotFound(parentID)))
	}
	body, err := connection.Post[T](r.conn, r.basePath(parentID), req, handlers...)
	return body.Data, err
}

// Update performs a PUT on the child resource and returns the updated object.
func (r *SubResource[T, ParentID, ChildID]) Update(parentID ParentID, childID ChildID, req any) (T, error) {
	var zero T
	if r.validateParent != nil && !r.validateParent(parentID) {
		return zero, fmt.Errorf("invalid %s %s", r.parentName, r.parentIdentifier)
	}
	if r.validateChild != nil && !r.validateChild(childID) {
		return zero, fmt.Errorf("invalid %s %s", r.childName, r.childIdentifier)
	}
	body, err := connection.Put[T](r.conn, fmt.Sprintf("%s/%v", r.basePath(parentID), childID), req,
		connection.NotFoundResponseHandler(r.childNotFound(parentID, childID)))
	return body.Data, err
}

// Patch performs a PATCH on the child resource (returns error only).
func (r *SubResource[T, ParentID, ChildID]) Patch(parentID ParentID, childID ChildID, req any) error {
	if r.validateParent != nil && !r.validateParent(parentID) {
		return fmt.Errorf("invalid %s %s", r.parentName, r.parentIdentifier)
	}
	if r.validateChild != nil && !r.validateChild(childID) {
		return fmt.Errorf("invalid %s %s", r.childName, r.childIdentifier)
	}
	return connection.PatchRaw(r.conn, fmt.Sprintf("%s/%v", r.basePath(parentID), childID), req, &connection.APIResponseBody{},
		connection.NotFoundResponseHandler(r.childNotFound(parentID, childID)))
}

// PatchReturning performs a PATCH on the child resource and returns the patched object.
func (r *SubResource[T, ParentID, ChildID]) PatchReturning(parentID ParentID, childID ChildID, req any) (T, error) {
	var zero T
	if r.validateParent != nil && !r.validateParent(parentID) {
		return zero, fmt.Errorf("invalid %s %s", r.parentName, r.parentIdentifier)
	}
	if r.validateChild != nil && !r.validateChild(childID) {
		return zero, fmt.Errorf("invalid %s %s", r.childName, r.childIdentifier)
	}
	body, err := connection.Patch[T](r.conn, fmt.Sprintf("%s/%v", r.basePath(parentID), childID), req,
		connection.NotFoundResponseHandler(r.childNotFound(parentID, childID)))
	return body.Data, err
}

func (r *SubResource[T, ParentID, ChildID]) Delete(parentID ParentID, childID ChildID) error {
	if r.validateParent != nil && !r.validateParent(parentID) {
		return fmt.Errorf("invalid %s %s", r.parentName, r.parentIdentifier)
	}
	if r.validateChild != nil && !r.validateChild(childID) {
		return fmt.Errorf("invalid %s %s", r.childName, r.childIdentifier)
	}
	return connection.DeleteRaw(r.conn, fmt.Sprintf("%s/%v", r.basePath(parentID), childID), nil, &connection.APIResponseBody{},
		connection.NotFoundResponseHandler(r.childNotFound(parentID, childID)))
}
