package managedcloudflare

import "github.com/ukfast/sdk-go/pkg/connection"

// PaginatedAccount represents a paginated collection of Account
type PaginatedAccount struct {
	*connection.PaginatedBase
	Items []Account
}

// NewPaginatedAccount returns a pointer to an initialized PaginatedAccount struct
func NewPaginatedAccount(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Account) *PaginatedAccount {
	return &PaginatedAccount{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedAccountMember represents a paginated collection of AccountMember
type PaginatedAccountMember struct {
	*connection.PaginatedBase
	Items []AccountMember
}

// NewPaginatedAccountMember returns a pointer to an initialized PaginatedAccountMember struct
func NewPaginatedAccountMember(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []AccountMember) *PaginatedAccountMember {
	return &PaginatedAccountMember{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedSpendPlan represents a paginated collection of SpendPlan
type PaginatedSpendPlan struct {
	*connection.PaginatedBase
	Items []SpendPlan
}

// NewPaginatedSpendPlan returns a pointer to an initialized PaginatedSpendPlan struct
func NewPaginatedSpendPlan(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []SpendPlan) *PaginatedSpendPlan {
	return &PaginatedSpendPlan{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedSubscription represents a paginated collection of Subscription
type PaginatedSubscription struct {
	*connection.PaginatedBase
	Items []Subscription
}

// NewPaginatedSubscription returns a pointer to an initialized PaginatedSubscription struct
func NewPaginatedSubscription(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Subscription) *PaginatedSubscription {
	return &PaginatedSubscription{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedZone represents a paginated collection of Zone
type PaginatedZone struct {
	*connection.PaginatedBase
	Items []Zone
}

// NewPaginatedZone returns a pointer to an initialized PaginatedZone struct
func NewPaginatedZone(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Zone) *PaginatedZone {
	return &PaginatedZone{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}
