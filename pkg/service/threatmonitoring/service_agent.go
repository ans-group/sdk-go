package threatmonitoring

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetAgents retrieves a list of agents
func (s *Service) GetAgents(parameters connection.APIRequestParameters) ([]Agent, error) {
	var agents []Agent

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetAgentsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, agent := range response.(*PaginatedAgent).Items {
			agents = append(agents, agent)
		}
	}

	return agents, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetAgentsPaginated retrieves a paginated list of agents
func (s *Service) GetAgentsPaginated(parameters connection.APIRequestParameters) (*PaginatedAgent, error) {
	body, err := s.getAgentsPaginatedResponseBody(parameters)

	return NewPaginatedAgent(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetAgentsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getAgentsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetAgentSliceResponseBody, error) {
	body := &GetAgentSliceResponseBody{}

	response, err := s.connection.Get("/threat-monitoring/v1/agents", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetAgent retrieves a single agent by id
func (s *Service) GetAgent(agentID string) (Agent, error) {
	body, err := s.getAgentResponseBody(agentID)

	return body.Data, err
}

func (s *Service) getAgentResponseBody(agentID string) (*GetAgentResponseBody, error) {
	body := &GetAgentResponseBody{}

	if agentID == "" {
		return body, fmt.Errorf("invalid agent id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/threat-monitoring/v1/agents/%s", agentID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &AgentNotFoundError{ID: agentID}
		}

		return nil
	})
}
