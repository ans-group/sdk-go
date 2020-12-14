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

// PaginatedTargetGroup represents a paginated collection of TargetGroup
type PaginatedTargetGroup struct {
	*connection.PaginatedBase
	Items []TargetGroup
}

// NewPaginatedTargetGroup returns a pointer to an initialized PaginatedTargetGroup struct
func NewPaginatedTargetGroup(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []TargetGroup) *PaginatedTargetGroup {
	return &PaginatedTargetGroup{
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

// PaginatedVIP represents a paginated collection of VIP
type PaginatedVIP struct {
	*connection.PaginatedBase
	Items []VIP
}

// NewPaginatedVIP returns a pointer to an initialized PaginatedVIP struct
func NewPaginatedVIP(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []VIP) *PaginatedVIP {
	return &PaginatedVIP{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedListener represents a paginated collection of Listener
type PaginatedListener struct {
	*connection.PaginatedBase
	Items []Listener
}

// NewPaginatedListener returns a pointer to an initialized PaginatedListener struct
func NewPaginatedListener(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Listener) *PaginatedListener {
	return &PaginatedListener{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedErrorPage represents a paginated collection of ErrorPage
type PaginatedErrorPage struct {
	*connection.PaginatedBase
	Items []ErrorPage
}

// NewPaginatedErrorPage returns a pointer to an initialized PaginatedErrorPage struct
func NewPaginatedErrorPage(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []ErrorPage) *PaginatedErrorPage {
	return &PaginatedErrorPage{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedAccess represents a paginated collection of Access
type PaginatedAccess struct {
	*connection.PaginatedBase
	Items []Access
}

// NewPaginatedAccess returns a pointer to an initialized PaginatedAccess struct
func NewPaginatedAccess(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Access) *PaginatedAccess {
	return &PaginatedAccess{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedBind represents a paginated collection of Bind
type PaginatedBind struct {
	*connection.PaginatedBase
	Items []Bind
}

// NewPaginatedBind returns a pointer to an initialized PaginatedBind struct
func NewPaginatedBind(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Bind) *PaginatedBind {
	return &PaginatedBind{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedListenerCertificate represents a paginated collection of ListenerCertificate
type PaginatedListenerCertificate struct {
	*connection.PaginatedBase
	Items []ListenerCertificate
}

// NewPaginatedListenerCertificate returns a pointer to an initialized PaginatedListenerCertificate struct
func NewPaginatedListenerCertificate(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []ListenerCertificate) *PaginatedListenerCertificate {
	return &PaginatedListenerCertificate{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedListenerErrorPage represents a paginated collection of ListenerErrorPage
type PaginatedListenerErrorPage struct {
	*connection.PaginatedBase
	Items []ListenerErrorPage
}

// NewPaginatedListenerErrorPage returns a pointer to an initialized PaginatedListenerErrorPage struct
func NewPaginatedListenerErrorPage(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []ListenerErrorPage) *PaginatedListenerErrorPage {
	return &PaginatedListenerErrorPage{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedSSL represents a paginated collection of SSL
type PaginatedSSL struct {
	*connection.PaginatedBase
	Items []SSL
}

// NewPaginatedSSL returns a pointer to an initialized PaginatedSSL struct
func NewPaginatedSSL(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []SSL) *PaginatedSSL {
	return &PaginatedSSL{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedCustomOption represents a paginated collection of CustomOption
type PaginatedCustomOption struct {
	*connection.PaginatedBase
	Items []CustomOption
}

// NewPaginatedCustomOption returns a pointer to an initialized PaginatedCustomOption struct
func NewPaginatedCustomOption(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []CustomOption) *PaginatedCustomOption {
	return &PaginatedCustomOption{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}
