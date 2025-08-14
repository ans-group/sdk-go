package account

import (
	"encoding/json"

	"github.com/ans-group/sdk-go/pkg/connection"
)

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

// Invoice represents a UKFast account invoice
type Invoice struct {
	ID    int             `json:"id"`
	Date  connection.Date `json:"date"`
	Paid  bool            `json:"paid"`
	Net   float32         `json:"net"`
	VAT   float32         `json:"vat"`
	Gross float32         `json:"gross"`
}

// InvoiceQuery represents a UKFast account invoice query
type InvoiceQuery struct {
	ID               int     `json:"id"`
	ContactID        int     `json:"contact_id"`
	Amount           float32 `json:"amount"`
	WhatWasExpected  string  `json:"what_was_expected"`
	WhatWasReceived  string  `json:"what_was_received"`
	ProposedSolution string  `json:"proposed_solution"`
	InvoiceIDs       []int   `json:"invoice_ids"`
}

// Client represents an account client
type Client struct {
	ID               int    `json:"id"`
	CompanyName      string `json:"company_name"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	EmailAddress     string `json:"email_address"`
	LimitedNumber    string `json:"limited_number"`
	VATNumber        string `json:"vat_number"`
	Address          string `json:"address"`
	Address1         string `json:"address1"`
	City             string `json:"city"`
	County           string `json:"county"`
	Country          string `json:"country"`
	Postcode         string `json:"postcode"`
	Phone            string `json:"phone"`
	Fax              string `json:"fax"`
	Mobile           string `json:"mobile"`
	Type             string `json:"type"`
	UserName         string `json:"user_name"`
	IDReference      string `json:"id_reference"`
	NominetContactID string `json:"nominet_contact_id"`
	CreatedDate      string `json:"created_date"`
}

// Application represents an API Application
type Application struct {
	ID          string              `json:"id"`
	Key         string              `json:"key"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	CreatedAt   connection.DateTime `json:"created_at"`
	CreatedBy   string              `json:"created_by"`
}

type CreateApplicationResponse struct {
	ID  string `json:"id"`
	Key string `json:"key"`
}

type ApplicationService struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ApplicationServiceMapping struct {
	Scopes []ApplicationServiceScope `json:"scopes"`
}

type ApplicationServiceScope struct {
	Service string   `json:"service"`
	Roles   []string `json:"roles"`
}

type ApplicationRestriction struct {
	IPRestrictionType string   `json:"ip_restriction_type"`
	IPRanges          []string `json:"ip_ranges"`
}

func (a *ApplicationRestriction) UnmarshalJSON(data []byte) error {
	// First, try to determine if the data is an array or object
	var raw json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	// If the data starts with '[', it's an array (empty restrictions)
	if len(raw) > 0 && raw[0] == '[' {
		// Empty restrictions case - initialize with empty values
		a.IPRestrictionType = ""
		a.IPRanges = nil
		return nil
	}

	// Otherwise, unmarshal as normal object
	type Alias ApplicationRestriction
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	}
	return json.Unmarshal(data, aux)
}
