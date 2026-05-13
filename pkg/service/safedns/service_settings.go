package safedns

import (
	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetSettings retrieves account settings
func (s *Service) GetSettings() (Settings, error) {
	body, err := connection.Get[Settings](s.connection, "/safedns/v1/settings", connection.APIRequestParameters{})
	return body.Data, err
}
