package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetAppliances retrieves a list of appliances
func (s *Service) GetAppliances(parameters connection.APIRequestParameters) ([]Appliance, error) {
	return connection.InvokeRequestAll(s.GetAppliancesPaginated, parameters)
}

// GetAppliancesPaginated retrieves a paginated list of appliances
func (s *Service) GetAppliancesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Appliance], error) {
	body, err := connection.Get[[]Appliance](s.connection, "/ecloud/v1/appliances", parameters)
	return connection.NewPaginated(body, parameters, s.GetAppliancesPaginated), err
}

// GetAppliance retrieves a single Appliance by ID
func (s *Service) GetAppliance(applianceID string) (Appliance, error) {
	if applianceID == "" {
		return Appliance{}, fmt.Errorf("invalid appliance id")
	}
	body, err := connection.Get[Appliance](s.connection, fmt.Sprintf("/ecloud/v1/appliances/%s", applianceID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&ApplianceNotFoundError{ID: applianceID}))
	return body.Data, err
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
