package ssl

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetRecommendations retrieves SSL recommendations for a domain
func (s *Service) GetRecommendations(domainName string) (Recommendations, error) {
	if domainName == "" {
		return Recommendations{}, fmt.Errorf("invalid domain name")
	}
	body, err := connection.Get[Recommendations](s.connection, fmt.Sprintf("/ssl/v1/recommendations/%s", domainName), connection.APIRequestParameters{})
	return body.Data, err
}
