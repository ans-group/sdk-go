package cloudflare

import (
	"github.com/ukfast/sdk-go/pkg/connection"
)

// CloudflareService is an interface for managing Cloudflare services
type CloudflareService interface {
	// Account
	GetAccounts(parameters connection.APIRequestParameters) ([]Account, error)
	GetAccountsPaginated(parameters connection.APIRequestParameters) (*PaginatedAccount, error)
	GetAccount(accountID string) (Account, error)
	CreateAccount(req CreateAccountRequest) (string, error)
	CreateAccountMember(accountID string, req CreateAccountMemberRequest) error

	// Orchestration
	CreateOrchestration(req CreateOrchestrationRequest) error

	// Spend plan
	GetSpendPlans(parameters connection.APIRequestParameters) ([]SpendPlan, error)
	GetSpendPlansPaginated(parameters connection.APIRequestParameters) (*PaginatedSpendPlan, error)

	// Subscription
	GetSubscriptions(parameters connection.APIRequestParameters) ([]Subscription, error)
	GetSubscriptionsPaginated(parameters connection.APIRequestParameters) (*PaginatedSubscription, error)

	// Zone
	GetZones(parameters connection.APIRequestParameters) ([]Zone, error)
	GetZonesPaginated(parameters connection.APIRequestParameters) (*PaginatedZone, error)
	GetZone(zoneID string) (Zone, error)
	CreateZone(req CreateZoneRequest) (string, error)
	DeleteZone(zoneID string) error

	// Spend
	GetTotalSpendMonthToDate() (TotalSpend, error)
}

// Service implements CloudflareService for managing the Shared Exchange service
type Service struct {
	connection connection.Connection
}

// NewService returns a new instance of CloudflareService
func NewService(connection connection.Connection) *Service {
	return &Service{
		connection: connection,
	}
}
