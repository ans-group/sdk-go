package account

import "github.com/ans-group/sdk-go/pkg/connection"

// GetDetails retrieves account details
func (s *Service) GetDetails() (Details, error) {
	body, err := connection.Get[Details](s.connection, "/account/v1/details", connection.APIRequestParameters{})
	return body.Data, err
}
