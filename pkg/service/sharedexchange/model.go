package sharedexchange

import "github.com/ukfast/sdk-go/pkg/connection"

// Domain represents an Shared Exchange domain
type Domain struct {
	ID        int                 `json:"id"`
	Name      string              `json:"name"`
	Version   string              `json:"version"`
	CreatedAt connection.DateTime `json:"created_at"`
}
