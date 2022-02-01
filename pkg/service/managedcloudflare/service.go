package managedcloudflare

import (
	"github.com/ukfast/sdk-go/pkg/connection"
)

// ManagedCloudflareService is an interface for managing Shared Exchange
type ManagedCloudflareService interface {
	// Account
	GetAccounts(parameters connection.APIRequestParameters) ([]Account, error)
	GetAccountsPaginated(parameters connection.APIRequestParameters) (*PaginatedAccount, error)
	GetAccount(accountID int) (Account, error)
	CreateAccount(req CreateAccountRequest) error
	CreateAccountMember(accountID string, req CreateAccountMemberRequest) error

	// Orchestration
	CreateOrchestration(req CreateOrchestrationRequest) error
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
