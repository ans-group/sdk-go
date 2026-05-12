package billing

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
	body, err := connection.Get[[]Invoice](s.connection, "/billing/v1/invoices", parameters)
	return connection.NewPaginated(body, parameters, s.GetInvoicesPaginated), err
}

// GetInvoice retrieves a single invoice by id
func (s *Service) GetInvoice(invoiceID int) (Invoice, error) {
	if invoiceID < 1 {
		return Invoice{}, fmt.Errorf("invalid invoice id")
	}
	body, err := connection.Get[Invoice](s.connection, fmt.Sprintf("/billing/v1/invoices/%d", invoiceID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&InvoiceNotFoundError{ID: invoiceID}))
	return body.Data, err
}
