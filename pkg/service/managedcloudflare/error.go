package managedcloudflare

import "fmt"

// AccountNotFoundError indicates an account was not found
type AccountNotFoundError struct {
	ID string
}

func (e *AccountNotFoundError) Error() string {
	return fmt.Sprintf("Account not found with ID [%s]", e.ID)
}
