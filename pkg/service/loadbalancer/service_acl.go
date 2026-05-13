package loadbalancer

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetACLs retrieves a list of ACLs
// Currently, a target_group_id or listener_id filter must be provided for this to return data
func (s *Service) GetACLs(parameters connection.APIRequestParameters) ([]ACL, error) {
	return connection.InvokeRequestAll(s.GetACLsPaginated, parameters)
}

// GetACLsPaginated retrieves a paginated list of ACLs
// Currently, a target_group_id or listener_id filter must be provided for this to return data
func (s *Service) GetACLsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[ACL], error) {
	body, err := connection.Get[[]ACL](s.connection, "/loadbalancers/v2/acls", parameters)
	return connection.NewPaginated(body, parameters, s.GetACLsPaginated), err
}

// GetACL retrieves a single ACL by id
func (s *Service) GetACL(aclID int) (ACL, error) {
	if aclID < 1 {
		return ACL{}, fmt.Errorf("invalid acl id")
	}
	body, err := connection.Get[ACL](s.connection, fmt.Sprintf("/loadbalancers/v2/acls/%d", aclID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&ACLNotFoundError{ID: aclID}))
	return body.Data, err
}

// CreateACL creates an ACL
func (s *Service) CreateACL(req CreateACLRequest) (int, error) {
	body, err := connection.Post[ACL](s.connection, "/loadbalancers/v2/acls", &req)
	return body.Data.ID, err
}

// PatchACL patches an ACL
func (s *Service) PatchACL(aclID int, req PatchACLRequest) error {
	if aclID < 1 {
		return fmt.Errorf("invalid acl id")
	}
	return connection.PatchRaw(s.connection, fmt.Sprintf("/loadbalancers/v2/acls/%d", aclID), &req, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&ACLNotFoundError{ID: aclID}))
}

// DeleteACL deletes an ACL
func (s *Service) DeleteACL(aclID int) error {
	if aclID < 1 {
		return fmt.Errorf("invalid acl id")
	}
	return connection.DeleteRaw(s.connection, fmt.Sprintf("/loadbalancers/v2/acls/%d", aclID), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&ACLNotFoundError{ID: aclID}))
}
