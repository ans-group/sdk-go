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

// Alert represents a threat monitoring alert
// +genie:model_response
// +genie:model_paginated
type Alert struct {
	ID                string               `json:"id"`
	IP                connection.IPAddress `json:"ip"`
	AgentID           string               `json:"agent_id"`
	AgentFriendlyName string               `json:"agent_friendly_name"`
	Geolocation       struct {
		IP      connection.IPAddress `json:"ip"`
		Country string               `json:"country"`
		Lon     float32              `json:"lon"`
		Lat     float32              `json:"lat"`
	} `json:"geolocation"`
	Level       int                 `json:"level"`
	Groups      []string            `json:"groups"`
	Description string              `json:"description"`
	Link        string              `json:"link"`
	GDPR        []string            `json:"gdpr"`
	PCIDSS      []string            `json:"pci_dss"`
	FullLog     string              `json:"full_log"`
	Timestamp   connection.DateTime `json:"timestamp"`
	Syscheck    struct {
		Path              string   `json:"path"`
		ChangedAttributes []string `json:"changed_attributes"`
		Diff              string   `json:"diff"`
	} `json:"syscheck"`
	Audit struct {
		EffectiveUsername string `json:"effective_user_name"`
		ProcessName       string `json:"process_name"`
		ProcessID         string `json:"process_id"`
		User              struct {
			Name string `json:"name"`
		} `json:"user"`
		Group struct {
			Name string `json:"name"`
		} `json:"group"`
	} `json:"audit"`
}
