package account

import "fmt"

// ContactNotFoundError indicates a contact was not found
type ContactNotFoundError struct {
	ID int
}

func (e *ContactNotFoundError) Error() string {
	return fmt.Sprintf("Contact not found with ID [%d]", e.ID)
}

// InvoiceNotFoundError indicates an invoice was not found
type InvoiceNotFoundError struct {
	ID int
}

func (e *InvoiceNotFoundError) Error() string {
	return fmt.Sprintf("Invoice not found with ID [%d]", e.ID)
}

// InvoiceQueryNotFoundError indicates an invoice query was not found
type InvoiceQueryNotFoundError struct {
	ID int
}

func (e *InvoiceQueryNotFoundError) Error() string {
	return fmt.Sprintf("Invoice query not found with ID [%d]", e.ID)
}

// ClientNotFoundError indicates a client was not found
type ClientNotFoundError struct {
	ID int
}

func (e *ClientNotFoundError) Error() string {
	return fmt.Sprintf("Client not found with ID [%d]", e.ID)
}

type ApplicationNotFoundError struct {
	ID string
}

func (e *ApplicationNotFoundError) Error() string {
	return fmt.Sprintf("Application not found with ID [%s]", e.ID)
}
