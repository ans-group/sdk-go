package billing

import (
	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) paymentRes() *resource.Resource[Payment, int] {
	return resource.NewIntResource[Payment](s.connection, "/billing/v1/payments", "payment",
		func(id int) error { return &PaymentNotFoundError{ID: id} })
}

// GetPayments retrieves a list of payments
func (s *Service) GetPayments(parameters connection.APIRequestParameters) ([]Payment, error) {
	return s.paymentRes().List(parameters)
}

// GetPaymentsPaginated retrieves a paginated list of payments
func (s *Service) GetPaymentsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Payment], error) {
	return s.paymentRes().ListPaginated(parameters)
}

// GetPayment retrieves a single payment by id
func (s *Service) GetPayment(paymentID int) (Payment, error) {
	return s.paymentRes().Get(paymentID)
}
