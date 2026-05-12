package pss

import "github.com/ans-group/sdk-go/pkg/connection"

// GetCaseOptions retrieves a list of supported services
func (s *Service) GetSupportedServices(parameters connection.APIRequestParameters) ([]SupportedService, error) {
	return connection.InvokeRequestAll(s.GetSupportedServicesPaginated, parameters)
}

// GetSupportedServicesPaginated retrieves a paginated list of supported services
func (s *Service) GetSupportedServicesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[SupportedService], error) {
	body, err := connection.Get[[]SupportedService](s.connection, "/pss/v2/supported-services", parameters)
	return connection.NewPaginated(body, parameters, s.GetSupportedServicesPaginated), err
}
