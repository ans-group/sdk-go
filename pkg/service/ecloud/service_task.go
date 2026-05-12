package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetTasks retrieves a list of tasks
func (s *Service) GetTasks(parameters connection.APIRequestParameters) ([]Task, error) {
	return connection.InvokeRequestAll(s.GetTasksPaginated, parameters)
}

// GetTasksPaginated retrieves a paginated list of tasks
func (s *Service) GetTasksPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Task], error) {
	body, err := connection.Get[[]Task](s.connection, "/ecloud/v2/tasks", parameters)
	return connection.NewPaginated(body, parameters, s.GetTasksPaginated), err
}

// GetTask retrieves a single task by id
func (s *Service) GetTask(taskID string) (Task, error) {
	if taskID == "" {
		return Task{}, fmt.Errorf("invalid task id")
	}
	body, err := connection.Get[Task](s.connection, fmt.Sprintf("/ecloud/v2/tasks/%s", taskID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&TaskNotFoundError{ID: taskID}))
	return body.Data, err
}
