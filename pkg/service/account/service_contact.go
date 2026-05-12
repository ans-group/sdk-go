package account

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetContacts retrieves a list of contacts
func (s *Service) GetContacts(parameters connection.APIRequestParameters) ([]Contact, error) {
	return connection.InvokeRequestAll(s.GetContactsPaginated, parameters)
}

// GetContactsPaginated retrieves a paginated list of contacts
func (s *Service) GetContactsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Contact], error) {
	body, err := connection.Get[[]Contact](s.connection, "/account/v1/contacts", parameters)
	return connection.NewPaginated(body, parameters, s.GetContactsPaginated), err
}

// GetContact retrieves a single contact by id
func (s *Service) GetContact(contactID int) (Contact, error) {
	if contactID < 1 {
		return Contact{}, fmt.Errorf("invalid contact id")
	}
	body, err := connection.Get[Contact](s.connection, fmt.Sprintf("/account/v1/contacts/%d", contactID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&ContactNotFoundError{ID: contactID}))
	return body.Data, err
}
