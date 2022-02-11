package managedcloudflare

import (
	"github.com/ukfast/sdk-go/pkg/connection"
)

// ManagedCloudflareService is an interface for managing Shared Exchange
type ManagedCloudflareService interface {
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
	CreateZone(req CreateZoneRequest) error
	DeleteZone(zoneID string) error
}

// Service implements ManagedCloudflareService for managing the Shared Exchange service
type Service struct {
	connection connection.Connection
}

// NewService returns a new instance of ManagedCloudflareService
func NewService(connection connection.Connection) *Service {
	return &Service{
		connection: connection,
	}
}
