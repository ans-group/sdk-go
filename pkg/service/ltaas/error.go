package ltaas

import "fmt"

// DomainNotFoundError indicates a domain was not found
type DomainNotFoundError struct {
	ID string
}

func (e *DomainNotFoundError) Error() string {
	return fmt.Sprintf("domain not found with ID [%s]", e.ID)
}

// TestNotFoundError indicates a test was not found
type TestNotFoundError struct {
	ID string
}

func (e *TestNotFoundError) Error() string {
	return fmt.Sprintf("test not found with ID [%s]", e.ID)
}

// JobNotFoundError indicates a job was not found
type JobNotFoundError struct {
	ID string
}

func (e *JobNotFoundError) Error() string {
	return fmt.Sprintf("job not found with ID [%s]", e.ID)
}

// ThresholdNotFoundError indicates a threshold was not found
type ThresholdNotFoundError struct {
	ID string
}

func (e *ThresholdNotFoundError) Error() string {
	return fmt.Sprintf("threshold not found with ID [%s]", e.ID)
}
