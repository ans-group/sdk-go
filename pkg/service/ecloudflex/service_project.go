package ecloudflex

import (
	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) projectRes() *resource.Resource[Project, int] {
	return resource.NewIntResource[Project](s.connection, "/ecloud-flex/v1/projects", "project",
		func(id int) error { return &ProjectNotFoundError{ID: id} })
}

// GetProjects retrieves a list of projects
func (s *Service) GetProjects(parameters connection.APIRequestParameters) ([]Project, error) {
	return s.projectRes().List(parameters)
}

// GetProjectsPaginated retrieves a paginated list of projects
func (s *Service) GetProjectsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Project], error) {
	return s.projectRes().ListPaginated(parameters)
}

// GetProject retrieves a single project by id
func (s *Service) GetProject(projectID int) (Project, error) {
	return s.projectRes().Get(projectID)
}
