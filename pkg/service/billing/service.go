package billing

import (
	"github.com/ukfast/sdk-go/pkg/connection"
)

// BillingService is an interface for managing billing
type BillingService interface {
}

// Service implements BillingService for managing
// Billing certificates via the UKFast API
type Service struct {
	connection connection.Connection
}

// NewService returns a new instance of BillingService
func NewService(connection connection.Connection) *Service {
	return &Service{
		connection: connection,
	}
}
