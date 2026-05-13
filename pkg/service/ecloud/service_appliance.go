package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) applianceRes() *resource.Resource[Appliance, string] {
	return resource.NewStringResource[Appliance](s.connection, "/ecloud/v1/appliances", "appliance", func(id string) error {
		return &ApplianceNotFoundError{ID: id}
	})
}

// GetAppliances retrieves a list of appliances
func (s *Service) GetAppliances(parameters connection.APIRequestParameters) ([]Appliance, error) {
	return s.applianceRes().List(parameters)
}

// GetAppliancesPaginated retrieves a paginated list of appliances
func (s *Service) GetAppliancesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Appliance], error) {
	return s.applianceRes().ListPaginated(parameters)
}

// GetAppliance retrieves a single Appliance by ID
func (s *Service) GetAppliance(applianceID string) (Appliance, error) {
	return s.applianceRes().Get(applianceID)
}

// GetApplianceParameters retrieves a list of parameters
func (s *Service) GetApplianceParameters(applianceID string, parameters connection.APIRequestParameters) ([]ApplianceParameter, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[ApplianceParameter], error) {
		return s.GetApplianceParametersPaginated(applianceID, p)
	}, parameters)
}

// GetApplianceParametersPaginated retrieves a paginated list of domains
func (s *Service) GetApplianceParametersPaginated(applianceID string, parameters connection.APIRequestParameters) (*connection.Paginated[ApplianceParameter], error) {
	if applianceID == "" {
		return nil, fmt.Errorf("invalid appliance id")
	}
	body, err := connection.Get[[]ApplianceParameter](s.connection, fmt.Sprintf("/ecloud/v1/appliances/%s/parameters", applianceID), parameters, connection.NotFoundResponseHandler(&ApplianceNotFoundError{ID: applianceID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[ApplianceParameter], error) {
		return s.GetApplianceParametersPaginated(applianceID, p)
	}), err
}
