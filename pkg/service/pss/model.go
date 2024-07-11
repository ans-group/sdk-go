package pss

import "github.com/ans-group/sdk-go/pkg/connection"

type AuthorType string

const (
	AuthorTypeClient  AuthorType = "Client"
	AuthorTypeAuto    AuthorType = "Auto"
	AuthorTypeSupport AuthorType = "Support"
)

var AuthorTypeEnum connection.Enum[AuthorType] = []AuthorType{AuthorTypeClient, AuthorTypeAuto, AuthorTypeSupport}

type RequestPriority string

const (
	RequestPriorityNormal   RequestPriority = "Normal"
	RequestPriorityHigh     RequestPriority = "High"
	RequestPriorityCritical RequestPriority = "Critical"
)

var RequestPriorityEnum connection.Enum[RequestPriority] = []RequestPriority{RequestPriorityNormal, RequestPriorityHigh, RequestPriorityCritical}

type RequestStatus string

const (
	RequestStatusCompleted                RequestStatus = "Completed"
	RequestStatusAwaitingCustomerResponse RequestStatus = "Awaiting Customer Response"
	RequestStatusRepliedAndCompleted      RequestStatus = "Replied and Completed"
	RequestStatusSubmitted                RequestStatus = "Submitted"
)

var RequestStatusEnum connection.Enum[RequestStatus] = []RequestStatus{
	RequestStatusCompleted,
	RequestStatusAwaitingCustomerResponse,
	RequestStatusRepliedAndCompleted,
	RequestStatusSubmitted,
}

// Request represents a PSS request
type Request struct {
	ID                int                 `json:"id"`
	Author            Author              `json:"author"`
	Type              string              `json:"type"`
	Secure            bool                `json:"secure"`
	Subject           string              `json:"subject"`
	CreatedAt         connection.DateTime `json:"created_at"`
	Priority          RequestPriority     `json:"priority"`
	Archived          bool                `json:"archived"`
	Status            RequestStatus       `json:"status"`
	RequestSMS        bool                `json:"request_sms"`
	Version           int                 `json:"version"`
	CustomerReference string              `json:"customer_reference"`
	Product           Product             `json:"product"`
	LastRepliedAt     connection.DateTime `json:"last_replied_at"`
	CC                []string            `json:"cc"`
	UnreadReplies     int                 `json:"unread_replies"`
	ContactMethod     string              `json:"contact_method"`
}

// Author represents a PSS request author
type Author struct {
	ID   int        `json:"id"`
	Name string     `json:"name"`
	Type AuthorType `json:"type"`
}

// Reply represents a PSS reply
type Reply struct {
	ID          string              `json:"id"`
	RequestID   int                 `json:"request_id"`
	Author      Author              `json:"author"`
	Description string              `json:"description"`
	Attachments []Attachment        `json:"attachments"`
	Read        bool                `json:"read"`
	CreatedAt   connection.DateTime `json:"created_at"`
}

// Attachment represents a PSS attachment
type Attachment struct {
	Name string `json:"name"`
}

// Product represents a product to which the request applies to
type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// Feedback represents PSS feedback
type Feedback struct {
	ID               int                 `json:"id"`
	ContactID        int                 `json:"contact_id"`
	Score            int                 `json:"score"`
	Comment          string              `json:"comment"`
	SpeedResolved    int                 `json:"speed_resolved"`
	Quality          int                 `json:"quality"`
	NPSScore         int                 `json:"nps_score"`
	ThirdPartConsent bool                `json:"thirdparty_consent"`
	CreatedAt        connection.DateTime `json:"created_at"`
}

// CaseCategory represents a PSS case category
type CaseCategory struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// CaseOption represents a PSS case option
type CaseOption struct {
	Option string `json:"option"`
}

type CaseType string

const (
	CaseTypeAllow CaseType = "allow"
	CaseTypeDeny  CaseType = "deny"
)

var CaseTypeEnum connection.Enum[CaseType] = []CaseType{
	CaseTypeAllow,
	CaseTypeDeny,
}

func test() {
	_, _ = CaseTypeEnum.Parse("allow")

}

// Case represents a PSS case
type Case struct {
	ID                    string   `json:"id"`
	CaseType              string   `json:"case_type"`
	Title                 string   `json:"title"`
	Description           string   `json:"description"`
	IsSecurity            bool     `json:"is_security"`
	Type                  CaseType `json:"type"`
	Priority              string   `json:"priority"`
	Squad                 string   `json:"squad"`
	Status                string   `json:"status"`
	CustomerReference     string   `json:"customer_reference"`
	Source                string   `json:"source"`
	Contact               string   `json:"contact"`
	CriteriaForResolution string   `json:"criteria_for_resolution"`
	ServiceName           string   `json:"service_name"`
	Owner                 string   `json:"owner"`
	ResponseTarget        string   `json:"response_target"`
	ResolutionTarget      string   `json:"resolution_target"`
	CreatedAt             string   `json:"created_at"`
	UpdatedAt             string   `json:"updated_at"`
}
