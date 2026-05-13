package billing

import (
	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) invoiceRes() *resource.Resource[Invoice, int] {
	return resource.NewIntResource[Invoice](s.connection, "/billing/v1/invoices", "invoice",
		func(id int) error { return &InvoiceNotFoundError{ID: id} })
}

// GetInvoices retrieves a list of invoices
func (s *Service) GetInvoices(parameters connection.APIRequestParameters) ([]Invoice, error) {
	return s.invoiceRes().List(parameters)
}

// GetInvoicesPaginated retrieves a paginated list of invoices
func (s *Service) GetInvoicesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Invoice], error) {
	return s.invoiceRes().ListPaginated(parameters)
}

// GetInvoice retrieves a single invoice by id
func (s *Service) GetInvoice(invoiceID int) (Invoice, error) {
	return s.invoiceRes().Get(invoiceID)
}
