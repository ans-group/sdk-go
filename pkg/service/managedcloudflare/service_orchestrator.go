package managedcloudflare

import (
	"github.com/ukfast/sdk-go/pkg/connection"
)

// CreateOrchestration creates a new orchestration request
func (s *Service) CreateOrchestration(req CreateOrchestrationRequest) error {
	_, err := s.createOrchestrationResponseBody(req)

	return err
}

func (s *Service) createOrchestrationResponseBody(req CreateOrchestrationRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	response, err := s.connection.Post("/managed-cloudflare/v2/orchestrator", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}
