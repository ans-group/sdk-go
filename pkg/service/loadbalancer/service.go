package loadbalancer

import (
	"github.com/ukfast/sdk-go/pkg/connection"
)

// LoadBalancerService is an interface for managing the LoadBalancer service
type LoadBalancerService interface {
	GetTargets(parameters connection.APIRequestParameters) ([]Target, error)
	GetTargetsPaginated(parameters connection.APIRequestParameters) (*PaginatedTarget, error)
	GetTarget(targetID string) (Target, error)
	CreateTarget(req CreateTargetRequest) (string, error)
	PatchTarget(targetID string, req PatchTargetRequest) error
	DeleteTarget(targetID string) error

	GetClusters(parameters connection.APIRequestParameters) ([]Cluster, error)
	GetClustersPaginated(parameters connection.APIRequestParameters) (*PaginatedCluster, error)
	GetCluster(configurationID string) (Cluster, error)
	PatchCluster(configurationID string, req PatchClusterRequest) error
	DeleteCluster(configurationID string) error
}

// Service implements LoadBalancerService for managing
// LoadBalancer certificates via the UKFast API
type Service struct {
	connection connection.Connection
}

// NewService returns a new instance of LoadBalancerService
func NewService(connection connection.Connection) *Service {
	return &Service{
		connection: connection,
	}
}
