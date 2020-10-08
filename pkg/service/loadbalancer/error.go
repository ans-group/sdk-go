package loadbalancer

import "fmt"

// GroupNotFoundError indicates a group was not found
type GroupNotFoundError struct {
	ID string
}

func (e *GroupNotFoundError) Error() string {
	return fmt.Sprintf("Group not found with ID [%s]", e.ID)
}

// ConfigurationNotFoundError indicates a configuration was not found
type ConfigurationNotFoundError struct {
	ID string
}

func (e *ConfigurationNotFoundError) Error() string {
	return fmt.Sprintf("Configuration not found with ID [%s]", e.ID)
}
