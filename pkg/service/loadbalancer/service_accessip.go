package loadbalancer

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetAccessIP retrieves a single access IP by id
func (s *Service) GetAccessIP(accessID int) (AccessIP, error) {
	if accessID < 1 {
		return AccessIP{}, fmt.Errorf("invalid access id")
	}
	body, err := connection.Get[AccessIP](s.connection, fmt.Sprintf("/loadbalancers/v2/access-ips/%d", accessID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&AccessIPNotFoundError{ID: accessID}))
	return body.Data, err
}

// PatchAccessIP patches an access IP
func (s *Service) PatchAccessIP(accessID int, req PatchAccessIPRequest) error {
	if accessID < 1 {
		return fmt.Errorf("invalid access id")
	}
	return connection.PatchRaw(s.connection, fmt.Sprintf("/loadbalancers/v2/access-ips/%d", accessID), &req, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&AccessIPNotFoundError{ID: accessID}))
}

// DeleteAccessIP deletes an access IP
func (s *Service) DeleteAccessIP(accessID int) error {
	if accessID < 1 {
		return fmt.Errorf("invalid access id")
	}
	return connection.DeleteRaw(s.connection, fmt.Sprintf("/loadbalancers/v2/access-ips/%d", accessID), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&AccessIPNotFoundError{ID: accessID}))
}
