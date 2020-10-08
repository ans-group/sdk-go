package loadbalancer

import "github.com/ukfast/sdk-go/pkg/connection"

// PaginatedGroup represents a paginated collection of Group
type PaginatedGroup struct {
	*connection.PaginatedBase
	Items []Group
}

// NewPaginatedGroup returns a pointer to an initialized PaginatedGroup struct
func NewPaginatedGroup(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Group) *PaginatedGroup {
	return &PaginatedGroup{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedConfiguration represents a paginated collection of Configuration
type PaginatedConfiguration struct {
	*connection.PaginatedBase
	Items []Configuration
}

// NewPaginatedConfiguration returns a pointer to an initialized PaginatedConfiguration struct
func NewPaginatedConfiguration(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Configuration) *PaginatedConfiguration {
	return &PaginatedConfiguration{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}
