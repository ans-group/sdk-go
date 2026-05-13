package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetHostSpecs retrieves a list of host specs
func (s *Service) GetHostSpecs(parameters connection.APIRequestParameters) ([]HostSpec, error) {
	return connection.InvokeRequestAll(s.GetHostSpecsPaginated, parameters)
}

// GetHostSpecsPaginated retrieves a paginated list of host specs
func (s *Service) GetHostSpecsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[HostSpec], error) {
	body, err := connection.Get[[]HostSpec](s.connection, "/ecloud/v2/host-specs", parameters)
	return connection.NewPaginated(body, parameters, s.GetHostSpecsPaginated), err
}

// GetHostSpec retrieves a single host spec by id
func (s *Service) GetHostSpec(specID string) (HostSpec, error) {
	if specID == "" {
		return HostSpec{}, fmt.Errorf("invalid spec id")
	}
	body, err := connection.Get[HostSpec](s.connection, fmt.Sprintf("/ecloud/v2/host-specs/%s", specID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&HostSpecNotFoundError{ID: specID}))
	return body.Data, err
}
