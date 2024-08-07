package pss

import "github.com/ans-group/sdk-go/pkg/connection"

// CreateRequestRequest represents a request to create a PSS request
type CreateRequestRequest struct {
	connection.APIRequestBodyDefaultValidator

	Author            Author          `json:"author" validate:"required"`
	Secure            bool            `json:"secure"`
	Priority          RequestPriority `json:"priority" validate:"required"`
	Subject           string          `json:"subject" validate:"required"`
	Details           string          `json:"details" validate:"required"`
	CC                []string        `json:"cc,omitempty"`
	RequestSMS        bool            `json:"request_sms"`
	CustomerReference string          `json:"customer_reference,omitempty"`
	Product           *Product        `json:"product,omitempty"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *CreateRequestRequest) Validate() *connection.ValidationError {
	return c.APIRequestBodyDefaultValidator.Validate(c)
}

// PatchRequestRequest represents a PSS Request patch request
type PatchRequestRequest struct {
	Secure     *bool           `json:"secure,omitempty"`
	Read       *bool           `json:"read,omitempty"`
	Priority   RequestPriority `json:"priority,omitempty"`
	Status     RequestStatus   `json:"status,omitempty"`
	RequestSMS *bool           `json:"request_sms,omitempty"`
	Archived   *bool           `json:"archived,omitempty"`
}

// CreateReplyRequest represents a request to create a PSS request reply
type CreateReplyRequest struct {
	connection.APIRequestBodyDefaultValidator

	Author      Author `json:"author" validate:"required"`
	Description string `json:"description" validate:"required"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *CreateReplyRequest) Validate() *connection.ValidationError {
	return c.APIRequestBodyDefaultValidator.Validate(c)
}

// CreateFeedbackRequest represents a request to create PSS request feedback
type CreateFeedbackRequest struct {
	ContactID         int    `json:"contact_id"`
	Score             int    `json:"score"`
	Comment           string `json:"comment"`
	SpeedResolved     int    `json:"speed_resolved"`
	Quality           int    `json:"quality"`
	NPSScore          int    `json:"nps_score"`
	ThirdPartyConsent bool   `json:"thirdparty_consent"`
}

// CreateChangeCaseRequest represents a request to create PSS change case
type CreateChangeCaseRequest struct {
	CaseType           CaseType         `json:"case_type"`
	Title              string           `json:"title"`
	Description        string           `json:"description"`
	IsSecurity         bool             `json:"is_security"`
	CustomerReference  string           `json:"customer_reference"`
	Reason             string           `json:"reason"`
	Impact             ChangeCaseImpact `json:"impact"`
	Risk               ChangeCaseRisk   `json:"risk"`
	CategoryID         string           `json:"category_id"`
	SupportedServiceID string           `json:"supported_service_id"`
}

// CreateIncidentCaseRequest represents a request to create PSS incident case
type CreateIncidentCaseRequest struct {
	CaseType           CaseType           `json:"case_type"`
	Title              string             `json:"title"`
	Description        string             `json:"description"`
	IsSecurity         bool               `json:"is_security"`
	Type               IncidentCaseType   `json:"type"`
	CustomerReference  string             `json:"customer_reference,omitempty"`
	ContactID          int                `json:"contact_id,omitempty"`
	CategoryID         string             `json:"category_id"`
	SupportedServiceID string             `json:"supported_service_id"`
	Impact             IncidentCaseImpact `json:"impact"`
}

// ApproveChangeCaseRequest represents a request to approve a PSS change case
type ApproveChangeCaseRequest struct {
	Reason    string `json:"reason,omitempty"`
	ContactID int    `json:"contact_id,omitempty"`
}

// CloseIncidentCaseRequest represents a request to close a PSS incident case
type CloseIncidentCaseRequest struct {
	Reason    string `json:"reason"`
	ContactID int    `json:"contact_id,omitempty"`
}

// CreateCaseUpdateRequest represents a request to create a PSS incident case update
type CreateCaseUpdateRequest struct {
	Description string `json:"description"`
	ContactID   int    `json:"contact_id,omitempty"`
}
