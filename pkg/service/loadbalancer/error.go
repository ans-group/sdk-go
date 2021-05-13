package loadbalancer

import "fmt"

// TargetNotFoundError indicates a target was not found
type TargetNotFoundError struct {
	ID int
}

func (e *TargetNotFoundError) Error() string {
	return fmt.Sprintf("Target not found with ID [%d]", e.ID)
}

// ClusterNotFoundError indicates a cluster was not found
type ClusterNotFoundError struct {
	ID int
}

func (e *ClusterNotFoundError) Error() string {
	return fmt.Sprintf("Cluster not found with ID [%d]", e.ID)
}

// TargetGroupNotFoundError indicates a target group was not found
type TargetGroupNotFoundError struct {
	ID int
}

func (e *TargetGroupNotFoundError) Error() string {
	return fmt.Sprintf("Target group not found with ID [%d]", e.ID)
}
