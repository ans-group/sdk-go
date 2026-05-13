package billing

import "github.com/ans-group/sdk-go/pkg/connection"

// GetDirectDebit retrieves direct debit details
func (s *Service) GetDirectDebit() (DirectDebit, error) {
	body, err := connection.Get[DirectDebit](s.connection, "/billing/v1/direct-debit", connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&DirectDebitNotFoundError{}))
	return body.Data, err
}
