//go:generate go run ../../gen/model_paginated/main.go -package threatmonitoring -typename Agent -destination model_paginated_generated.go
//go:generate go run ../../gen/model_response/main.go -package threatmonitoring -typename Agent -destination model_response_generated.go

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

// YesNo represents a value containing either 'Yes' or 'No'
type YesNo string

func (e YesNo) String() string {
	return string(e)
}

const (
	Yes YesNo = "Yes"
	No  YesNo = "No"
)

var YesNoEnum connection.EnumSlice = []connection.Enum{
	Yes,
	No,
}

// ParseYesNo attempts to parse a YesNo from string
func ParseYesNo(s string) (YesNo, error) {
	e, err := connection.ParseEnum(s, YesNoEnum)
	if err != nil {
		return "", err
	}

	return e.(YesNo), err
}

// Agent represents a threat monitoring agent
type Agent struct {
	ID                    int                 `json:"id"`
	Status                AgentStatus         `json:"status"`
	ThreatResponseEnabled YesNo               `json:"threat_response_enabled"`
	FieldName             string              `json:"friendly_name"`
	Platform              string              `json:"platform"`
	CreatedAt             connection.DateTime `json:"created_at"`
	UpdatedAt             connection.DateTime `json:"updated_at"`
}
