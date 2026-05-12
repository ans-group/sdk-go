package account

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetInvoices retrieves a list of invoices
func (s *Service) GetInvoices(parameters connection.APIRequestParameters) ([]Invoice, error) {
	return connection.InvokeRequestAll(s.GetInvoicesPaginated, parameters)
}

// GetInvoicesPaginated retrieves a paginated list of invoices
func (s *Service) GetInvoicesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Invoice], error) {
	body, err := connection.Get[[]Invoice](s.connection, "/account/v1/invoices", parameters)
	return connection.NewPaginated(body, parameters, s.GetInvoicesPaginated), err
}

// GetInvoice retrieves a single invoice by id
func (s *Service) GetInvoice(invoiceID int) (Invoice, error) {
	if invoiceID < 1 {
		return Invoice{}, fmt.Errorf("invalid invoice id")
	}
	body, err := connection.Get[Invoice](s.connection, fmt.Sprintf("/account/v1/invoices/%d", invoiceID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&InvoiceNotFoundError{ID: invoiceID}))
	return body.Data, err
}

// GetInvoiceQueries retrieves a list of invoice queries
func (s *Service) GetInvoiceQueries(parameters connection.APIRequestParameters) ([]InvoiceQuery, error) {
	return connection.InvokeRequestAll(s.GetInvoiceQueriesPaginated, parameters)
}

// GetInvoiceQueriesPaginated retrieves a paginated list of invoice queries
func (s *Service) GetInvoiceQueriesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[InvoiceQuery], error) {
	body, err := connection.Get[[]InvoiceQuery](s.connection, "/account/v1/invoice-queries", parameters)
	return connection.NewPaginated(body, parameters, s.GetInvoiceQueriesPaginated), err
}

// GetInvoiceQuery retrieves a single invoice query by id
func (s *Service) GetInvoiceQuery(queryID int) (InvoiceQuery, error) {
	if queryID < 1 {
		return InvoiceQuery{}, fmt.Errorf("invalid invoice query id")
	}
	body, err := connection.Get[InvoiceQuery](s.connection, fmt.Sprintf("/account/v1/invoice-queries/%d", queryID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&InvoiceQueryNotFoundError{ID: queryID}))
	return body.Data, err
}

// CreateInvoiceQuery retrieves creates an InvoiceQuery
func (s *Service) CreateInvoiceQuery(req CreateInvoiceQueryRequest) (int, error) {
	body, err := connection.Post[InvoiceQuery](s.connection, "/account/v1/invoice-queries", &req)
	return body.Data.ID, err
}
