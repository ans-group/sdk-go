package cloudflare

import (
	"github.com/ans-group/sdk-go/pkg/connection"
)

// CreateOrchestration creates a new orchestration request
func (s *Service) CreateOrchestration(req CreateOrchestrationRequest) error {
	_, err := connection.Post[struct{}](s.connection, "/cloudflare/v1/orchestrator", &req)
	return err
}
