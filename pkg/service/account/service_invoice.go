package account

import (
	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) invoiceRes() *resource.Resource[Invoice, int] {
	return resource.NewIntResource[Invoice](s.connection, "/account/v1/invoices", "invoice",
		func(id int) error { return &InvoiceNotFoundError{ID: id} })
}

func (s *Service) invoiceQueryRes() *resource.Resource[InvoiceQuery, int] {
	return resource.NewIntResource[InvoiceQuery](s.connection, "/account/v1/invoice-queries", "invoice query",
		func(id int) error { return &InvoiceQueryNotFoundError{ID: id} })
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

// GetInvoiceQueries retrieves a list of invoice queries
func (s *Service) GetInvoiceQueries(parameters connection.APIRequestParameters) ([]InvoiceQuery, error) {
	return s.invoiceQueryRes().List(parameters)
}

// GetInvoiceQueriesPaginated retrieves a paginated list of invoice queries
func (s *Service) GetInvoiceQueriesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[InvoiceQuery], error) {
	return s.invoiceQueryRes().ListPaginated(parameters)
}

// GetInvoiceQuery retrieves a single invoice query by id
func (s *Service) GetInvoiceQuery(queryID int) (InvoiceQuery, error) {
	return s.invoiceQueryRes().Get(queryID)
}

// CreateInvoiceQuery retrieves creates an InvoiceQuery
func (s *Service) CreateInvoiceQuery(req CreateInvoiceQueryRequest) (int, error) {
	data, err := s.invoiceQueryRes().Create(&req)
	return data.ID, err
}
