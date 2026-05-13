package account

import "github.com/ans-group/sdk-go/pkg/connection"

// GetCredits retrieves a list of credits
func (s *Service) GetCredits(parameters connection.APIRequestParameters) ([]Credit, error) {
	body, err := connection.Get[[]Credit](s.connection, "/account/v1/credits", parameters)
	return body.Data, err
}
