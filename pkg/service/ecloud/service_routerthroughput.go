package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetRouterThroughputs retrieves a list of router throughputs
func (s *Service) GetRouterThroughputs(parameters connection.APIRequestParameters) ([]RouterThroughput, error) {
	return connection.InvokeRequestAll(s.GetRouterThroughputsPaginated, parameters)
}

// GetRouterThroughputsPaginated retrieves a paginated list of router throughputs
func (s *Service) GetRouterThroughputsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[RouterThroughput], error) {
	body, err := connection.Get[[]RouterThroughput](s.connection, "/ecloud/v2/router-throughputs", parameters)
	return connection.NewPaginated(body, parameters, s.GetRouterThroughputsPaginated), err
}

// GetRouterThroughput retrieves a single router throughput by id
func (s *Service) GetRouterThroughput(throughputID string) (RouterThroughput, error) {
	if throughputID == "" {
		return RouterThroughput{}, fmt.Errorf("invalid router throughput id")
	}
	body, err := connection.Get[RouterThroughput](s.connection, fmt.Sprintf("/ecloud/v2/router-throughputs/%s", throughputID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&RouterThroughputNotFoundError{ID: throughputID}))
	return body.Data, err
}
