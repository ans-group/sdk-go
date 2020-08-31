package threatmonitoring

import "fmt"

// AgentNotFoundError indicates an agent was not found
type AgentNotFoundError struct {
	ID int
}

func (e *AgentNotFoundError) Error() string {
	return fmt.Sprintf("Agent not found with ID [%d]", e.ID)
}
