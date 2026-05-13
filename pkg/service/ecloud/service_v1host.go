package ecloud

import (
	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) v1HostRes() *resource.Resource[V1Host, int] {
	return resource.NewIntResource[V1Host](s.connection, "/ecloud/v1/hosts", "host", func(id int) error {
		return &V1HostNotFoundError{ID: id}
	})
}

// GetV1Hosts retrieves a list of v1 hosts
func (s *Service) GetV1Hosts(parameters connection.APIRequestParameters) ([]V1Host, error) {
	return s.v1HostRes().List(parameters)
}

// GetV1HostsPaginated retrieves a paginated list of v1 hosts
func (s *Service) GetV1HostsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[V1Host], error) {
	return s.v1HostRes().ListPaginated(parameters)
}

// GetV1Host retrieves a single v1 host by ID
func (s *Service) GetV1Host(hostID int) (V1Host, error) {
	return s.v1HostRes().Get(hostID)
}
