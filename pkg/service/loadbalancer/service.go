package loadbalancer

import (
	"github.com/ukfast/sdk-go/pkg/connection"
)

// LoadBalancerService is an interface for managing the LoadBalancer service
type LoadBalancerService interface {
	GetGroups(parameters connection.APIRequestParameters) ([]Group, error)
	GetGroupsPaginated(parameters connection.APIRequestParameters) (*PaginatedGroup, error)
	GetGroup(groupID string) (Group, error)

	GetConfigurations(parameters connection.APIRequestParameters) ([]Configuration, error)
	GetConfigurationsPaginated(parameters connection.APIRequestParameters) (*PaginatedConfiguration, error)
	GetConfiguration(configurationID string) (Configuration, error)
	CreateConfiguration(req CreateConfigurationRequest) (string, error)
	PatchConfiguration(configurationID string, req PatchConfigurationRequest) error
	DeleteConfiguration(configurationID string) error
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
