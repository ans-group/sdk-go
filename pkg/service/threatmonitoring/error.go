package threatmonitoring

import "fmt"

// AgentNotFoundError indicates an agent was not found
type AgentNotFoundError struct {
	ID string
}

func (e *AgentNotFoundError) Error() string {
	return fmt.Sprintf("Agent not found with ID [%s]", e.ID)
}

// AlertNotFoundError indicates an alert was not found
type AlertNotFoundError struct {
	ID string
}

func (e *AlertNotFoundError) Error() string {
	return fmt.Sprintf("Alert not found with ID [%s]", e.ID)
}
