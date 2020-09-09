//go:generate go run ../../gen/model_response/main.go -package threatmonitoring -source model.go -destination model_response_generated.go
//go:generate go run ../../gen/model_paginated/main.go -package threatmonitoring -source model.go -destination model_paginated_generated.go

package threatmonitoring

import "github.com/ukfast/sdk-go/pkg/connection"

type AgentStatus string

func (s AgentStatus) String() string {
	return string(s)
}

const (
	AgentStatusPending    AgentStatus = "Pending"
	AgentStatusInstalling AgentStatus = "Installing"
	AgentStatusFailed     AgentStatus = "Failed"
	AgentStatusCompleted  AgentStatus = "Completed"
	AgentStatusUnknown    AgentStatus = "Unknown"
	AgentStatusRemoved    AgentStatus = "Removed"
)

// Agent represents a threat monitoring agent
// +genie:model_response
// +genie:model_paginated
type Agent struct {
	ID                    string              `json:"id"`
	Status                AgentStatus         `json:"status"`
	ThreatResponseEnabled bool                `json:"threat_response_enabled"`
	FieldName             string              `json:"friendly_name"`
	Platform              string              `json:"platform"`
	CreatedAt             connection.DateTime `json:"created_at"`
	UpdatedAt             connection.DateTime `json:"updated_at"`
}
