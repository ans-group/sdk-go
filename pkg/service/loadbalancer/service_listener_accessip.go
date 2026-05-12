package loadbalancer

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetListenerAccessIPs retrieves a list of access IPs
func (s *Service) GetListenerAccessIPs(listenerID int, parameters connection.APIRequestParameters) ([]AccessIP, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[AccessIP], error) {
		return s.GetListenerAccessIPsPaginated(listenerID, p)
	}, parameters)
}

// GetListenerAccessIPsPaginated retrieves a paginated list of access IPs
func (s *Service) GetListenerAccessIPsPaginated(listenerID int, parameters connection.APIRequestParameters) (*connection.Paginated[AccessIP], error) {
	if listenerID < 1 {
		return nil, fmt.Errorf("invalid listener id")
	}
	body, err := connection.Get[[]AccessIP](s.connection, fmt.Sprintf("/loadbalancers/v2/listeners/%d/access-ips", listenerID), parameters)
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[AccessIP], error) {
		return s.GetListenerAccessIPsPaginated(listenerID, p)
	}), err
}

// GetListenerAccessIP retrieves a single access IP by id
func (s *Service) GetListenerAccessIP(listenerID int, accessID int) (AccessIP, error) {
	if listenerID < 1 {
		return AccessIP{}, fmt.Errorf("invalid listener id")
	}
	if accessID < 1 {
		return AccessIP{}, fmt.Errorf("invalid access id")
	}
	body, err := connection.Get[AccessIP](s.connection, fmt.Sprintf("/loadbalancers/v2/listeners/%d/access-ips/%d", listenerID, accessID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&AccessIPNotFoundError{ID: listenerID}))
	return body.Data, err
}

// CreateListenerAccessIP creates an access IP
func (s *Service) CreateListenerAccessIP(listenerID int, req CreateAccessIPRequest) (int, error) {
	if listenerID < 1 {
		return 0, fmt.Errorf("invalid listener id")
	}
	body, err := connection.Post[AccessIP](s.connection, fmt.Sprintf("/loadbalancers/v2/listeners/%d/access-ips", listenerID), &req, connection.NotFoundResponseHandler(&AccessIPNotFoundError{ID: listenerID}))
	return body.Data.ID, err
}
