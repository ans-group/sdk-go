package billing

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetPayments retrieves a list of payments
func (s *Service) GetPayments(parameters connection.APIRequestParameters) ([]Payment, error) {
	return connection.InvokeRequestAll(s.GetPaymentsPaginated, parameters)
}

// GetPaymentsPaginated retrieves a paginated list of payments
func (s *Service) GetPaymentsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Payment], error) {
	body, err := connection.Get[[]Payment](s.connection, "/billing/v1/payments", parameters)
	return connection.NewPaginated(body, parameters, s.GetPaymentsPaginated), err
}

// GetPayment retrieves a single payment by id
func (s *Service) GetPayment(paymentID int) (Payment, error) {
	if paymentID < 1 {
		return Payment{}, fmt.Errorf("invalid payment id")
	}
	body, err := connection.Get[Payment](s.connection, fmt.Sprintf("/billing/v1/payments/%d", paymentID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&PaymentNotFoundError{ID: paymentID}))
	return body.Data, err
}
