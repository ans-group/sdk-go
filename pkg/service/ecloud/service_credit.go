package ecloud

import (
	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/account"
)

// GetCredits retrieves a list of credits
func (s *Service) GetCredits(parameters connection.APIRequestParameters) ([]account.Credit, error) {
	body, err := connection.Get[[]account.Credit](s.connection, "/ecloud/v1/credits", parameters)
	return body.Data, err
}
