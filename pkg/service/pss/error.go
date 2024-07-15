package pss

import "fmt"

// RequestNotFoundError indicates a request was not found
type RequestNotFoundError struct {
	ID int
}

func (e *RequestNotFoundError) Error() string {
	return fmt.Sprintf("Request not found with id [%d]", e.ID)
}

// ReplyNotFoundError indicates a reply was not found
type ReplyNotFoundError struct {
	ID string
}

func (e *ReplyNotFoundError) Error() string {
	return fmt.Sprintf("Reply not found with id [%s]", e.ID)
}

// AttachmentNotFoundError indicates a attachment was not found
type AttachmentNotFoundError struct {
	Name string
}

func (e *AttachmentNotFoundError) Error() string {
	return fmt.Sprintf("Attachment not found with name [%s]", e.Name)
}

// RequestFeedbackNotFoundError indicates feedback for a request was not found
type RequestFeedbackNotFoundError struct {
	RequestID int
}

func (e *RequestFeedbackNotFoundError) Error() string {
	return fmt.Sprintf("Feedback not found for request [%d]", e.RequestID)
}

// RequestFeedbackNotFoundError indicates an incident case was not found
type IncidentCaseNotFoundError struct {
	ID string
}

func (e *IncidentCaseNotFoundError) Error() string {
	return fmt.Sprintf("Incident case not found for ID [%s]", e.ID)
}

// RequestFeedbackNotFoundError indicates a problem case was not found
type ProblemCaseNotFoundError struct {
	ID string
}

func (e *ProblemCaseNotFoundError) Error() string {
	return fmt.Sprintf("Problem case not found for ID [%s]", e.ID)
}

// RequestFeedbackNotFoundError indicates a change case was not found
type ChangeCaseNotFoundError struct {
	ID string
}

func (e *ChangeCaseNotFoundError) Error() string {
	return fmt.Sprintf("Change case not found for ID [%s]", e.ID)
}
