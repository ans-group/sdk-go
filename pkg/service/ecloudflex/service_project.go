package ecloudflex

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetProjects retrieves a list of projects
func (s *Service) GetProjects(parameters connection.APIRequestParameters) ([]Project, error) {
	return connection.InvokeRequestAll(s.GetProjectsPaginated, parameters)
}

// GetProjectsPaginated retrieves a paginated list of projects
func (s *Service) GetProjectsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Project], error) {
	body, err := connection.Get[[]Project](s.connection, "/ecloud-flex/v1/projects", parameters)
	return connection.NewPaginated(body, parameters, s.GetProjectsPaginated), err
}

// GetProject retrieves a single project by id
func (s *Service) GetProject(projectID int) (Project, error) {
	if projectID < 1 {
		return Project{}, fmt.Errorf("invalid project id")
	}
	body, err := connection.Get[Project](s.connection, fmt.Sprintf("/ecloud-flex/v1/projects/%d", projectID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&ProjectNotFoundError{ID: projectID}))
	return body.Data, err
}
