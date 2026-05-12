package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetIOPSs retrieves a list of iops
func (s *Service) GetIOPSTiers(parameters connection.APIRequestParameters) ([]IOPSTier, error) {
	return connection.InvokeRequestAll(s.GetIOPSTiersPaginated, parameters)
}

// GetIOPSsPaginated retrieves a paginated list of iops
func (s *Service) GetIOPSTiersPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[IOPSTier], error) {
	body, err := connection.Get[[]IOPSTier](s.connection, "/ecloud/v2/iops", parameters)
	return connection.NewPaginated(body, parameters, s.GetIOPSTiersPaginated), err
}

// GetIOPS retrieves a single IOPS by ID
func (s *Service) GetIOPSTier(iopsID string) (IOPSTier, error) {
	if iopsID == "" {
		return IOPSTier{}, fmt.Errorf("invalid IOPS id")
	}
	body, err := connection.Get[IOPSTier](s.connection, fmt.Sprintf("/ecloud/v2/iops/%s", iopsID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&IOPSNotFoundError{ID: iopsID}))
	return body.Data, err
}
