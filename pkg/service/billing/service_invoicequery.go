package billing

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetInvoiceQueries retrieves a list of invoice queries
func (s *Service) GetInvoiceQueries(parameters connection.APIRequestParameters) ([]InvoiceQuery, error) {
	return connection.InvokeRequestAll(s.GetInvoiceQueriesPaginated, parameters)
}

// GetInvoiceQueriesPaginated retrieves a paginated list of invoice queries
func (s *Service) GetInvoiceQueriesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[InvoiceQuery], error) {
	body, err := connection.Get[[]InvoiceQuery](s.connection, "/billing/v1/invoice-queries", parameters)
	return connection.NewPaginated(body, parameters, s.GetInvoiceQueriesPaginated), err
}

// GetInvoiceQuery retrieves a single invoice query by id
func (s *Service) GetInvoiceQuery(queryID int) (InvoiceQuery, error) {
	if queryID < 1 {
		return InvoiceQuery{}, fmt.Errorf("invalid invoice query id")
	}
	body, err := connection.Get[InvoiceQuery](s.connection, fmt.Sprintf("/billing/v1/invoice-queries/%d", queryID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&InvoiceQueryNotFoundError{ID: queryID}))
	return body.Data, err
}

// CreateInvoiceQuery retrieves creates an InvoiceQuery
func (s *Service) CreateInvoiceQuery(req CreateInvoiceQueryRequest) (int, error) {
	body, err := connection.Post[InvoiceQuery](s.connection, "/billing/v1/invoice-queries", &req)
	return body.Data.ID, err
}
