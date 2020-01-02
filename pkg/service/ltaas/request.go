package ltaas

import (
	"github.com/ukfast/sdk-go/pkg/connection"
)

// CreateJobRequest represents a request to create a job
type CreateJobRequest struct {
	connection.APIRequestBodyDefaultValidator

	TestID             string              `json:"test_id" validate:"required"`
	ScheduledTimestamp connection.DateTime `json:"scheduled_timestamp,omitempty"`
	RunNow             bool                `json:"run_now"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *CreateJobRequest) Validate() *connection.ValidationError {
	return c.APIRequestBodyDefaultValidator.Validate(c)
}

// CreateDomainRequest represents a request to create a domain
type CreateDomainRequest struct {
	connection.APIRequestBodyDefaultValidator

	Name               string                   `json:"name" validate:"required"`
	VerificationMethod DomainVerificationMethod `json:"verification_method" validate:"required"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *CreateDomainRequest) Validate() *connection.ValidationError {
	return c.APIRequestBodyDefaultValidator.Validate(c)
}

// CreateTestAuthorisation represents the test authorisation payload
type CreateTestAuthorisation struct {
	AgreementVersion string `json:"agreement_version" validate:"required"`
	Name             string `json:"name" validate:"required"`
	Position         string `json:"position" validate:"required"`
	Company          string `json:"company" validate:"required"`
}

// CreateTestThreshold represents the test threshold payload
type CreateTestThreshold struct {
	ThresholdID string `json:"threshold_id"`
	Values      string `json:"values"`
	Warn        int    `json:"warn"`
	Fail        int    `json:"fail"`
}

// CreateTestRequest represents a request to create a test
type CreateTestRequest struct {
	connection.APIRequestBodyDefaultValidator

	DomainID      string                  `json:"domain_id" validate:"required"`
	Name          string                  `json:"name" validate:"required"`
	ScenarioID    string                  `json:"scenario_id"`
	ScriptID      string                  `json:"script_id"`
	Protocol      TestProtocol            `json:"protocol"`
	Path          string                  `json:"path"`
	NumberOfUsers int                     `json:"number_of_users"`
	Duration      TestDuration            `json:"duration"`
	Authorisation CreateTestAuthorisation `json:"authorisation"`
	Thresholds    []CreateTestThreshold   `json:"thresholds"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *CreateTestRequest) Validate() *connection.ValidationError {
	return c.APIRequestBodyDefaultValidator.Validate(c)
}

// CreateTestJobRequest represents a request to create a job for a test
type CreateTestJobRequest struct {
	connection.APIRequestBodyDefaultValidator

	ScheduledTimestamp connection.DateTime `json:"scheduled_timestamp,omitempty"`
	RunNow             bool                `json:"run_now"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *CreateTestJobRequest) Validate() *connection.ValidationError {
	return c.APIRequestBodyDefaultValidator.Validate(c)
}
