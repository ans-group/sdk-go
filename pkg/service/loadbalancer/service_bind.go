package loadbalancer

import (
	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) bindRes() *resource.Resource[Bind, int] {
	return resource.NewIntResource[Bind](s.connection, "/loadbalancers/v2/binds", "bind",
		func(id int) error { return &BindNotFoundError{ID: id} })
}

// GetBinds retrieves a list of binds
func (s *Service) GetBinds(parameters connection.APIRequestParameters) ([]Bind, error) {
	return s.bindRes().List(parameters)
}

// GetBindsPaginated retrieves a paginated list of binds
func (s *Service) GetBindsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Bind], error) {
	return s.bindRes().ListPaginated(parameters)
}
