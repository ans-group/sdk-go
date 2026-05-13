package ecloud

import (
	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) routerThroughputRes() *resource.Resource[RouterThroughput, string] {
	return resource.NewStringResource[RouterThroughput](s.connection, "/ecloud/v2/router-throughputs", "router throughput", func(id string) error {
		return &RouterThroughputNotFoundError{ID: id}
	})
}

// GetRouterThroughputs retrieves a list of router throughputs
func (s *Service) GetRouterThroughputs(parameters connection.APIRequestParameters) ([]RouterThroughput, error) {
	return s.routerThroughputRes().List(parameters)
}

// GetRouterThroughputsPaginated retrieves a paginated list of router throughputs
func (s *Service) GetRouterThroughputsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[RouterThroughput], error) {
	return s.routerThroughputRes().ListPaginated(parameters)
}

// GetRouterThroughput retrieves a single router throughput by id
func (s *Service) GetRouterThroughput(throughputID string) (RouterThroughput, error) {
	return s.routerThroughputRes().Get(throughputID)
}
