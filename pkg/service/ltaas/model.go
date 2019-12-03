//go:generate go run ../../gen/model_paginated_gen.go -package ltaas -typename Domain,Test,Job -destination model_paginated.go

package ltaas

import (
	"github.com/ukfast/sdk-go/pkg/connection"
)

type DomainVerificationMethod string

func (s DomainVerificationMethod) String() string {
	return string(s)
}

const (
	DomainVerificationMethodDNS        DomainVerificationMethod = "DNS"
	DomainVerificationMethodFileUpload DomainVerificationMethod = "File upload"
)

type DomainStatus string

func (s DomainStatus) String() string {
	return string(s)
}

const (
	DomainStatusVerified    DomainStatus = "Verified"
	DomainStatusNotVerified DomainStatus = "Not verified"
)

type TestProtocol string

func (s TestProtocol) String() string {
	return string(s)
}

const (
	TestProtocolHTTP  TestProtocol = "http"
	TestProtocolHTTPS TestProtocol = "https"
)

type TestRecurringType string

func (s TestRecurringType) String() string {
	return string(s)
}

const (
	TestRecurringTypeDaily  TestRecurringType = "Daily"
	TestRecurringTypeOneOff TestRecurringType = "One off"
	TestRecurringTypeWeekly TestRecurringType = "Weekly"
)

type JobStatus string

func (s JobStatus) String() string {
	return string(s)
}

const (
	JobStatusPending  JobStatus = "Pending"
	JobStatusRunning  JobStatus = "Running"
	JobStatusFailed   JobStatus = "Failed"
	JobStatusStopped  JobStatus = "Stopped"
	JobStatusComplete JobStatus = "Complete"
)

type JobFailType string

func (s JobFailType) String() string {
	return string(s)
}

const (
	JobFailTypeTest           JobFailType = "Test"
	JobFailTypeInfrastructure JobFailType = "Infrastructure"
)

// Domain represents an LTaaS domain
type Domain struct {
	ID                 string                   `json:"id"`
	Name               string                   `json:"name"`
	VerificationMethod DomainVerificationMethod `json:"verification_method"`
	VerifyHash         string                   `json:"verify_hash"`
	Status             DomainStatus             `json:"status"`
	CreatedAt          connection.DateTime      `json:"created_at"`
	UpdatedAt          connection.DateTime      `json:"updated_at"`
}

// Test represents an LTaaS test
type Test struct {
	ID             string              `json:"id"`
	Name           string              `json:"name"`
	ScriptID       string              `json:"script_id"`
	ScenarioID     string              `json:"scenario_id"`
	DomainID       string              `json:"domain_id"`
	Protocol       TestProtocol        `json:"protocol"`
	Path           string              `json:"path"`
	NumberOfUsers  int                 `json:"number_of_users"`
	Duration       string              `json:"duration"`
	RecurringType  TestRecurringType   `json:"recurring_type"`
	RecurringValue int                 `json:"recurring_value"`
	NextRun        connection.DateTime `json:"next_run"`
	CreatedAt      connection.DateTime `json:"created_at"`
	UpdatedAt      connection.DateTime `json:"updated_at"`
}

// Job represents an LTaaS job
type Job struct {
	ID                 string              `json:"id"`
	TestID             string              `json:"test_id"`
	DomainID           string              `json:"domain_id"`
	ScheduledTimestamp connection.DateTime `json:"scheduled_timestamp"`
	JobStartTimestamp  connection.DateTime `json:"job_start_timestamp"`
	JobEndTimestamp    connection.DateTime `json:"job_end_timestamp"`
	Status             JobStatus           `json:"status"`
	FailType           JobFailType         `json:"fail_type"`
	CreatedAt          connection.DateTime `json:"created_at"`
	UpdatedAt          connection.DateTime `json:"updated_at"`
}

type JobResultsAxis struct {
	X connection.DateTime `json:"x"`
	Y float64             `json:"y"`
}

// JobResults represents the results of an LTaaS job
type JobResults struct {
	VirtualUsers       []JobResultsAxis `json:"virtual_users"`
	SuccessfulRequests []JobResultsAxis `json:"successful_requests"`
	FailedRequests     []JobResultsAxis `json:"failed_requests"`
	Latency            []JobResultsAxis `json:"latency"`
}
