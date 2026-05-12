package storage

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetHosts retrieves a list of hosts
func (s *Service) GetHosts(parameters connection.APIRequestParameters) ([]Host, error) {
	return connection.InvokeRequestAll(s.GetHostsPaginated, parameters)
}

// GetHostsPaginated retrieves a paginated list of hosts
func (s *Service) GetHostsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Host], error) {
	body, err := connection.Get[[]Host](s.connection, "/ukfast-storage/v1/hosts", parameters)
	return connection.NewPaginated(body, parameters, s.GetHostsPaginated), err
}

// GetHost retrieves a single host by id
func (s *Service) GetHost(hostID int) (Host, error) {
	if hostID < 1 {
		return Host{}, fmt.Errorf("invalid host id")
	}
	body, err := connection.Get[Host](s.connection, fmt.Sprintf("/ukfast-storage/v1/hosts/%d", hostID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&HostNotFoundError{ID: hostID}))
	return body.Data, err
}
