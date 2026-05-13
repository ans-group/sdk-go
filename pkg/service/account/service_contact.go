package account

import (
	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) contactRes() *resource.Resource[Contact, int] {
	return resource.NewIntResource[Contact](s.connection, "/account/v1/contacts", "contact",
		func(id int) error { return &ContactNotFoundError{ID: id} })
}

// GetContacts retrieves a list of contacts
func (s *Service) GetContacts(parameters connection.APIRequestParameters) ([]Contact, error) {
	return s.contactRes().List(parameters)
}

// GetContactsPaginated retrieves a paginated list of contacts
func (s *Service) GetContactsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Contact], error) {
	return s.contactRes().ListPaginated(parameters)
}

// GetContact retrieves a single contact by id
func (s *Service) GetContact(contactID int) (Contact, error) {
	return s.contactRes().Get(contactID)
}
