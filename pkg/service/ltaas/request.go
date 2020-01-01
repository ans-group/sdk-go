package ltaas

import "github.com/ukfast/sdk-go/pkg/connection"

// CreateJobRequest represents a request to create a job
type CreateJobRequest struct {
	connection.APIRequestBodyDefaultValidator

	TestID             string              `json:"test_id" validate:"required"`
	ScheduledTimestamp connection.DateTime `json:"scheduled_timestamp"`
	RunNow             bool                `json:"run_now"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *CreateJobRequest) Validate() *connection.ValidationError {
	return c.APIRequestBodyDefaultValidator.Validate(c)
}
