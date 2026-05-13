package loadbalancer

import (
	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) deploymentRes() *resource.Resource[Deployment, int] {
	return resource.NewIntResource[Deployment](s.connection, "/loadbalancers/v2/deployments", "deployment",
		func(id int) error { return &DeploymentNotFoundError{ID: id} })
}

// GetDeployments retrieves a list of deployments
func (s *Service) GetDeployments(parameters connection.APIRequestParameters) ([]Deployment, error) {
	return s.deploymentRes().List(parameters)
}

// GetDeploymentsPaginated retrieves a paginated list of deployments
func (s *Service) GetDeploymentsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Deployment], error) {
	return s.deploymentRes().ListPaginated(parameters)
}

// GetDeployment retrieves a single deployment by id
func (s *Service) GetDeployment(deploymentID int) (Deployment, error) {
	return s.deploymentRes().Get(deploymentID)
}
