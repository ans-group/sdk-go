package draas

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetIOPSTiers retrieves a list of solutions
func (s *Service) GetIOPSTiers(parameters connection.APIRequestParameters) ([]IOPSTier, error) {
	return connection.InvokeRequestAll(s.GetIOPSTiersPaginated, parameters)
}

// GetIOPSTiersPaginated retrieves a paginated list of solutions
func (s *Service) GetIOPSTiersPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[IOPSTier], error) {
	body, err := connection.Get[[]IOPSTier](s.connection, "/draas/v1/iops-tiers", parameters)
	return connection.NewPaginated(body, parameters, s.GetIOPSTiersPaginated), err
}

// GetIOPSTier retrieves a single solution by id
func (s *Service) GetIOPSTier(iopsTierID string) (IOPSTier, error) {
	if iopsTierID == "" {
		return IOPSTier{}, fmt.Errorf("invalid iops tier id")
	}
	body, err := connection.Get[IOPSTier](s.connection, fmt.Sprintf("/draas/v1/iops-tiers/%s", iopsTierID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&IOPSTierNotFoundError{ID: iopsTierID}))
	return body.Data, err
}
