package storage

import (
	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) hostRes() *resource.Resource[Host, int] {
	return resource.NewIntResource[Host](s.connection, "/ukfast-storage/v1/hosts", "host",
		func(id int) error { return &HostNotFoundError{ID: id} })
}

// GetHosts retrieves a list of hosts
func (s *Service) GetHosts(parameters connection.APIRequestParameters) ([]Host, error) {
	return s.hostRes().List(parameters)
}

// GetHostsPaginated retrieves a paginated list of hosts
func (s *Service) GetHostsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Host], error) {
	return s.hostRes().ListPaginated(parameters)
}

// GetHost retrieves a single host by id
func (s *Service) GetHost(hostID int) (Host, error) {
	return s.hostRes().Get(hostID)
}
