package draas

import (
	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) iopsTierRes() *resource.Resource[IOPSTier, string] {
	return resource.NewStringResource[IOPSTier](s.connection, "/draas/v1/iops-tiers", "iops tier",
		func(id string) error { return &IOPSTierNotFoundError{ID: id} })
}

// GetIOPSTiers retrieves a list of solutions
func (s *Service) GetIOPSTiers(parameters connection.APIRequestParameters) ([]IOPSTier, error) {
	return s.iopsTierRes().List(parameters)
}

// GetIOPSTiersPaginated retrieves a paginated list of solutions
func (s *Service) GetIOPSTiersPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[IOPSTier], error) {
	return s.iopsTierRes().ListPaginated(parameters)
}

// GetIOPSTier retrieves a single solution by id
func (s *Service) GetIOPSTier(iopsTierID string) (IOPSTier, error) {
	return s.iopsTierRes().Get(iopsTierID)
}
