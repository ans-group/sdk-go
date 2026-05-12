package storage

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetSolutions retrieves a list of solutions
func (s *Service) GetSolutions(parameters connection.APIRequestParameters) ([]Solution, error) {
	return connection.InvokeRequestAll(s.GetSolutionsPaginated, parameters)
}

// GetSolutionsPaginated retrieves a paginated list of solutions
func (s *Service) GetSolutionsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Solution], error) {
	body, err := connection.Get[[]Solution](s.connection, "/ukfast-storage/v1/solutions", parameters)
	return connection.NewPaginated(body, parameters, s.GetSolutionsPaginated), err
}

// GetSolution retrieves a single solution by id
func (s *Service) GetSolution(solutionID int) (Solution, error) {
	if solutionID < 1 {
		return Solution{}, fmt.Errorf("invalid solution id")
	}
	body, err := connection.Get[Solution](s.connection, fmt.Sprintf("/ukfast-storage/v1/solutions/%d", solutionID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&SolutionNotFoundError{ID: solutionID}))
	return body.Data, err
}
