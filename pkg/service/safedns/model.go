package safedns

import (
	"strconv"
	"time"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// RecordTTL represents the record TTL time in seconds
type RecordTTL int

// Time returns the record TTL time
func (r RecordTTL) Time() time.Time {
	return time.Now().Add(r.Duration())
}

// Duration returns the record TTL duration (seconds)
func (r RecordTTL) Duration() time.Duration {
	return (time.Second * time.Duration(int(r)))
}

func (r RecordTTL) String() string {
	return strconv.Itoa(int(r))
}

type RecordType string

func (s RecordType) String() string {
	return string(s)
}

const (
	RecordTypeA     RecordType = "A"
	RecordTypeAAAA  RecordType = "AAAA"
	RecordTypeCAA   RecordType = "CAA"
	RecordTypeCNAME RecordType = "CNAME"
	RecordTypeMX    RecordType = "MX"
	RecordTypeSPF   RecordType = "SPF"
	RecordTypeSRV   RecordType = "SRV"
	RecordTypeTXT   RecordType = "TXT"
	RecordTypeNS    RecordType = "NS"
	RecordTypeSOA   RecordType = "SOA"
	RecordTypeAXFR  RecordType = "AXFR"
)

// Zone represents a SafeDNS zone
type Zone struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type PaginatedZones struct {
	*connection.PaginatedBase

	Zones []Zone
}

func NewPaginatedZones(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, zones []Zone) *PaginatedZones {
	return &PaginatedZones{
		Zones:         zones,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// Record represents a SafeDNS record
type Record struct {
	connection.APIRequestBodyDefaultValidator

	ID         int                 `json:"id"`
	TemplateID int                 `json:"template_id"`
	Name       string              `json:"name" validate:"required"`
	Type       RecordType          `json:"type"`
	Content    string              `json:"content" validate:"required"`
	UpdatedAt  connection.DateTime `json:"updated_at"`
	TTL        RecordTTL           `json:"ttl"`
	Priority   int                 `json:"priority"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *Record) Validate() *connection.ValidationError {
	return c.APIRequestBodyDefaultValidator.Validate(c)
}

type PaginatedRecords struct {
	*connection.PaginatedBase

	Records []Record
}

func NewPaginatedRecords(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, records []Record) *PaginatedRecords {
	return &PaginatedRecords{
		Records:       records,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// Note represents a SafeDNS note
type Note struct {
	ID        int                  `json:"id"`
	ContactID int                  `json:"contact_id"`
	Notes     string               `json:"notes"`
	CreatedAt connection.DateTime  `json:"created_at"`
	IP        connection.IPAddress `json:"ip"`
}

type PaginatedNotes struct {
	*connection.PaginatedBase

	Notes []Note
}

func NewPaginatedNotes(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, notes []Note) *PaginatedNotes {
	return &PaginatedNotes{
		Notes:         notes,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// Template represents a SafeDNS template
type Template struct {
	connection.APIRequestBodyDefaultValidator

	ID        int             `json:"id"`
	Name      string          `json:"name" validate:"required"`
	Default   bool            `json:"default"`
	CreatedAt connection.Date `json:"created_at"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *Template) Validate() *connection.ValidationError {
	return c.APIRequestBodyDefaultValidator.Validate(c)
}

type PaginatedTemplates struct {
	*connection.PaginatedBase

	Templates []Template
}

func NewPaginatedTemplates(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, templates []Template) *PaginatedTemplates {
	return &PaginatedTemplates{
		Templates:     templates,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}
