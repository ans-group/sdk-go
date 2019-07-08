package pss

import (
	"github.com/ukfast/sdk-go/pkg/connection"
)

// PSSService is an interface for managing PSS
type PSSService interface {
	GetRequests(parameters connection.APIRequestParameters) ([]Request, error)
	GetRequestsPaginated(parameters connection.APIRequestParameters) (*PaginatedRequest, error)
	GetRequest(requestID int) (Request, error)
	GetRequestConversation(requestID int, parameters connection.APIRequestParameters) ([]Reply, error)
	GetRequestConversationPaginated(requestID int, parameters connection.APIRequestParameters) (*PaginatedReply, error)

	GetReply(replyID int) (Reply, error)
}

// Service implements PSSService for managing
// PSS certificates via the UKFast API
type Service struct {
	connection connection.Connection
}

// NewService returns a new instance of PSSService
func NewService(connection connection.Connection) *Service {
	return &Service{
		connection: connection,
	}
}
