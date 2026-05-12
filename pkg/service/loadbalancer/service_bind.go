package loadbalancer

import (
	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetBinds retrieves a list of binds
func (s *Service) GetBinds(parameters connection.APIRequestParameters) ([]Bind, error) {
	return connection.InvokeRequestAll(s.GetBindsPaginated, parameters)
}

// GetBindsPaginated retrieves a paginated list of binds
func (s *Service) GetBindsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Bind], error) {
	body, err := connection.Get[[]Bind](s.connection, "/loadbalancers/v2/binds", parameters)
	return connection.NewPaginated(body, parameters, s.GetBindsPaginated), err
}
