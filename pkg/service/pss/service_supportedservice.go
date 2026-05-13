package pss

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) supportedServiceRes() *resource.Resource[SupportedService, string] {
	return resource.NewStringResource[SupportedService](s.connection, "/pss/v2/supported-services", "supported-service",
		func(id string) error { return fmt.Errorf("supported service %s not found", id) })
}

// GetCaseOptions retrieves a list of supported services
func (s *Service) GetSupportedServices(parameters connection.APIRequestParameters) ([]SupportedService, error) {
	return s.supportedServiceRes().List(parameters)
}

// GetSupportedServicesPaginated retrieves a paginated list of supported services
func (s *Service) GetSupportedServicesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[SupportedService], error) {
	return s.supportedServiceRes().ListPaginated(parameters)
}
