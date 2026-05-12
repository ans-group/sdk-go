package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetBillingMetrics retrieves a list of billing metrics
func (s *Service) GetBillingMetrics(parameters connection.APIRequestParameters) ([]BillingMetric, error) {
	return connection.InvokeRequestAll(s.GetBillingMetricsPaginated, parameters)
}

// GetBillingMetricsPaginated retrieves a paginated list of billing metrics
func (s *Service) GetBillingMetricsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[BillingMetric], error) {
	body, err := connection.Get[[]BillingMetric](s.connection, "/ecloud/v2/billing-metrics", parameters)
	return connection.NewPaginated(body, parameters, s.GetBillingMetricsPaginated), err
}

// GetBillingMetric retrieves a single billing metrics by id
func (s *Service) GetBillingMetric(metricID string) (BillingMetric, error) {
	if metricID == "" {
		return BillingMetric{}, fmt.Errorf("invalid metric id")
	}
	body, err := connection.Get[BillingMetric](s.connection, fmt.Sprintf("/ecloud/v2/billing-metrics/%s", metricID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&BillingMetricNotFoundError{ID: metricID}))
	return body.Data, err
}
