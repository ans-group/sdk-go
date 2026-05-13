package storage

import (
	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) solutionRes() *resource.Resource[Solution, int] {
	return resource.NewIntResource[Solution](s.connection, "/ukfast-storage/v1/solutions", "solution",
		func(id int) error { return &SolutionNotFoundError{ID: id} })
}

// GetSolutions retrieves a list of solutions
func (s *Service) GetSolutions(parameters connection.APIRequestParameters) ([]Solution, error) {
	return s.solutionRes().List(parameters)
}

// GetSolutionsPaginated retrieves a paginated list of solutions
func (s *Service) GetSolutionsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Solution], error) {
	return s.solutionRes().ListPaginated(parameters)
}

// GetSolution retrieves a single solution by id
func (s *Service) GetSolution(solutionID int) (Solution, error) {
	return s.solutionRes().Get(solutionID)
}
