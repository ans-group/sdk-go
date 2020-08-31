//go:generate go run ../../gen/model_paginated/main.go -package sharedexchange -typename Domain -destination model_paginated_generated.go
//go:generate go run ../../gen/model_response/main.go -package sharedexchange -typename Domain -destination model_response_generated.go

package sharedexchange

import "github.com/ukfast/sdk-go/pkg/connection"

// Domain represents an Shared Exchange domain
type Domain struct {
	ID        int                 `json:"id"`
	Name      string              `json:"name"`
	Version   string              `json:"version"`
	CreatedAt connection.DateTime `json:"created_at"`
}
