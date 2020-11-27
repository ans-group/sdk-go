package loadbalancer

import "github.com/ukfast/sdk-go/pkg/connection"

// PaginatedTarget represents a paginated collection of Target
type PaginatedTarget struct {
	*connection.PaginatedBase
	Items []Target
}

// NewPaginatedTarget returns a pointer to an initialized PaginatedTarget struct
func NewPaginatedTarget(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Target) *PaginatedTarget {
	return &PaginatedTarget{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedCluster represents a paginated collection of Cluster
type PaginatedCluster struct {
	*connection.PaginatedBase
	Items []Cluster
}

// NewPaginatedCluster returns a pointer to an initialized PaginatedCluster struct
func NewPaginatedCluster(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Cluster) *PaginatedCluster {
	return &PaginatedCluster{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}
