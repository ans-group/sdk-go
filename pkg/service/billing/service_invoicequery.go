package billing

import (
	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) invoiceQueryRes() *resource.Resource[InvoiceQuery, int] {
	return resource.NewIntResource[InvoiceQuery](s.connection, "/billing/v1/invoice-queries", "invoice query",
		func(id int) error { return &InvoiceQueryNotFoundError{ID: id} })
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
