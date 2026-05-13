package ecloud

import (
	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) hostSpecRes() *resource.Resource[HostSpec, string] {
	return resource.NewStringResource[HostSpec](s.connection, "/ecloud/v2/host-specs", "spec", func(id string) error {
		return &HostSpecNotFoundError{ID: id}
	})
}

// GetHostSpecs retrieves a list of host specs
func (s *Service) GetHostSpecs(parameters connection.APIRequestParameters) ([]HostSpec, error) {
	return s.hostSpecRes().List(parameters)
}

// GetHostSpecsPaginated retrieves a paginated list of host specs
func (s *Service) GetHostSpecsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[HostSpec], error) {
	return s.hostSpecRes().ListPaginated(parameters)
}

// GetHostSpec retrieves a single host spec by id
func (s *Service) GetHostSpec(specID string) (HostSpec, error) {
	return s.hostSpecRes().Get(specID)
}
