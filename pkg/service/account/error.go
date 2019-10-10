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
