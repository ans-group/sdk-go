package account

import "github.com/ukfast/sdk-go/pkg/connection"

type ContactType string

func (t ContactType) String() string {
	return string(t)
}

const (
	ContactTypePrimaryContact ContactType = "Primary Contact"
	ContactTypeAccounts       ContactType = "Accounts"
	ContactTypeTechnical      ContactType = "Technical"
	ContactTypeThirdParty     ContactType = "Third Party"
	ContactTypeOther          ContactType = "Other"
)

// Contact represents a UKFast account contact
type Contact struct {
	ID        int         `json:"id"`
	Type      ContactType `json:"type"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
}

// Details represents a UKFast account details
type Details struct {
	CompanyRegistrationNumber string `json:"company_registration_number"`
	VATIdentificationNumber   string `json:"vat_identification_number"`
	PrimaryContactID          int    `json:"primary_contact_id"`
}

// Credit represents a UKFast account credit
type Credit struct {
	Type      string `json:"type"`
	Total     int    `json:"total"`
	Remaining int    `json:"remaining"`
}

// PaginatedContacts represents a paginated collection of contacts
type PaginatedContacts struct {
	*connection.PaginatedBase

	Contacts []Contact
}

// NewPaginatedContacts returns a pointer to an initialized PaginatedContacts struct
func NewPaginatedContacts(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, contacts []Contact) *PaginatedContacts {
	return &PaginatedContacts{
		Contacts:      contacts,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}
