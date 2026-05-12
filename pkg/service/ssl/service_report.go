package ssl

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetReport retrieves a single report by id
func (s *Service) GetReport(domainName string) (Report, error) {
	if domainName == "" {
		return Report{}, fmt.Errorf("invalid domain name")
	}
	body, err := connection.Get[Report](s.connection, fmt.Sprintf("/ssl/v1/reports/%s", domainName), connection.APIRequestParameters{})
	return body.Data, err
}
