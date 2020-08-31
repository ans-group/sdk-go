//go:generate go run ../../gen/model_paginated/main.go -package threatmonitoring -typename Agent -destination model_paginated_generated.go
//go:generate go run ../../gen/model_response/main.go -package threatmonitoring -typename Agent -destination model_response_generated.go

package threatmonitoring

import "github.com/ukfast/sdk-go/pkg/connection"

// Agent represents a threat monitoring agent
type Agent struct {
	ID                    int                 `json:"id"`
	Status                string              `json:"status"`
	ThreatResponseEnabled string              `json:"threat_response_enabled"`
	FieldName             string              `json:"friendly_name"`
	Platform              string              `json:"platform"`
	CreatedAt             connection.DateTime `json:"created_at"`
	UpdatedAt             connection.DateTime `json:"updated_at"`
}
