package loadbalancer

import "fmt"

// TargetNotFoundError indicates a target was not found
type TargetNotFoundError struct {
	ID string
}

func (e *TargetNotFoundError) Error() string {
	return fmt.Sprintf("Target not found with ID [%s]", e.ID)
}

// ClusterNotFoundError indicates a cluster was not found
type ClusterNotFoundError struct {
	ID string
}

func (e *ClusterNotFoundError) Error() string {
	return fmt.Sprintf("Cluster not found with ID [%s]", e.ID)
}
