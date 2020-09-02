package ltaas

import "github.com/ukfast/sdk-go/pkg/connection"

// PaginatedDomain represents a paginated collection of Domain
type PaginatedDomain struct {
	*connection.PaginatedBase
	Items []Domain
}

// NewPaginatedDomain returns a pointer to an initialized PaginatedDomain struct
func NewPaginatedDomain(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Domain) *PaginatedDomain {
	return &PaginatedDomain{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedTest represents a paginated collection of Test
type PaginatedTest struct {
	*connection.PaginatedBase
	Items []Test
}

// NewPaginatedTest returns a pointer to an initialized PaginatedTest struct
func NewPaginatedTest(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Test) *PaginatedTest {
	return &PaginatedTest{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedJob represents a paginated collection of Job
type PaginatedJob struct {
	*connection.PaginatedBase
	Items []Job
}

// NewPaginatedJob returns a pointer to an initialized PaginatedJob struct
func NewPaginatedJob(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Job) *PaginatedJob {
	return &PaginatedJob{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedThreshold represents a paginated collection of Threshold
type PaginatedThreshold struct {
	*connection.PaginatedBase
	Items []Threshold
}

// NewPaginatedThreshold returns a pointer to an initialized PaginatedThreshold struct
func NewPaginatedThreshold(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Threshold) *PaginatedThreshold {
	return &PaginatedThreshold{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedScenario represents a paginated collection of Scenario
type PaginatedScenario struct {
	*connection.PaginatedBase
	Items []Scenario
}

// NewPaginatedScenario returns a pointer to an initialized PaginatedScenario struct
func NewPaginatedScenario(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Scenario) *PaginatedScenario {
	return &PaginatedScenario{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}
