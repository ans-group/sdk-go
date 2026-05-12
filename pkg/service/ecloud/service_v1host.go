package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetV1Hosts retrieves a list of v1 hosts
func (s *Service) GetV1Hosts(parameters connection.APIRequestParameters) ([]V1Host, error) {
	return connection.InvokeRequestAll(s.GetV1HostsPaginated, parameters)
}

// GetV1HostsPaginated retrieves a paginated list of v1 hosts
func (s *Service) GetV1HostsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[V1Host], error) {
	body, err := connection.Get[[]V1Host](s.connection, "/ecloud/v1/hosts", parameters)
	return connection.NewPaginated(body, parameters, s.GetV1HostsPaginated), err
}

// GetV1Host retrieves a single v1 host by ID
func (s *Service) GetV1Host(hostID int) (V1Host, error) {
	if hostID < 1 {
		return V1Host{}, fmt.Errorf("invalid host id")
	}
	body, err := connection.Get[V1Host](s.connection, fmt.Sprintf("/ecloud/v1/hosts/%d", hostID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&V1HostNotFoundError{ID: hostID}))
	return body.Data, err
}
