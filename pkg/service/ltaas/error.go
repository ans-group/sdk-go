package ltaas

import "fmt"

// DomainNotFoundError indicates a virtual machine was not found
type DomainNotFoundError struct {
	ID int
}

func (e *DomainNotFoundError) Error() string {
	return fmt.Sprintf("domain not found with ID [%d]", e.ID)
}
