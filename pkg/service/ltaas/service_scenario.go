package ltaas

import (
	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetScenarios retrieves a list of scenarios
func (s *Service) GetScenarios(parameters connection.APIRequestParameters) ([]Scenario, error) {
	var sites []Scenario

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetScenariosPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, site := range response.(*PaginatedScenario).Items {
			sites = append(sites, site)
		}
	}

	return sites, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetScenariosPaginated retrieves a paginated list of scenarios
func (s *Service) GetScenariosPaginated(parameters connection.APIRequestParameters) (*PaginatedScenario, error) {
	body, err := s.getScenariosPaginatedResponseBody(parameters)

	return NewPaginatedScenario(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetScenariosPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getScenariosPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetScenariosResponseBody, error) {
	body := &GetScenariosResponseBody{}

	response, err := s.connection.Get("/ltaas/v1/scenarios", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}
