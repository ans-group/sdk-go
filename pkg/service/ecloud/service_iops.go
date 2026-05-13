package ecloud

import (
	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) iopsTierRes() *resource.Resource[IOPSTier, string] {
	return resource.NewStringResource[IOPSTier](s.connection, "/ecloud/v2/iops", "IOPS", func(id string) error {
		return &IOPSNotFoundError{ID: id}
	})
}

// GetIOPSs retrieves a list of iops
func (s *Service) GetIOPSTiers(parameters connection.APIRequestParameters) ([]IOPSTier, error) {
	return s.iopsTierRes().List(parameters)
}

// GetIOPSsPaginated retrieves a paginated list of iops
func (s *Service) GetIOPSTiersPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[IOPSTier], error) {
	return s.iopsTierRes().ListPaginated(parameters)
}

// GetIOPS retrieves a single IOPS by ID
func (s *Service) GetIOPSTier(iopsID string) (IOPSTier, error) {
	return s.iopsTierRes().Get(iopsID)
}
