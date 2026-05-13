package ecloud

import (
	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) taskRes() *resource.Resource[Task, string] {
	return resource.NewStringResource[Task](s.connection, "/ecloud/v2/tasks", "task", func(id string) error {
		return &TaskNotFoundError{ID: id}
	})
}

// GetTasks retrieves a list of tasks
func (s *Service) GetTasks(parameters connection.APIRequestParameters) ([]Task, error) {
	return s.taskRes().List(parameters)
}

// GetTasksPaginated retrieves a paginated list of tasks
func (s *Service) GetTasksPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Task], error) {
	return s.taskRes().ListPaginated(parameters)
}

// GetTask retrieves a single task by id
func (s *Service) GetTask(taskID string) (Task, error) {
	return s.taskRes().Get(taskID)
}
