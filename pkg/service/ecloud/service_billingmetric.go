package ecloud

import (
	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) billingMetricRes() *resource.Resource[BillingMetric, string] {
	return resource.NewStringResource[BillingMetric](s.connection, "/ecloud/v2/billing-metrics", "metric", func(id string) error {
		return &BillingMetricNotFoundError{ID: id}
	})
}

// GetBillingMetrics retrieves a list of billing metrics
func (s *Service) GetBillingMetrics(parameters connection.APIRequestParameters) ([]BillingMetric, error) {
	return s.billingMetricRes().List(parameters)
}

// GetBillingMetricsPaginated retrieves a paginated list of billing metrics
func (s *Service) GetBillingMetricsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[BillingMetric], error) {
	return s.billingMetricRes().ListPaginated(parameters)
}

// GetBillingMetric retrieves a single billing metrics by id
func (s *Service) GetBillingMetric(metricID string) (BillingMetric, error) {
	return s.billingMetricRes().Get(metricID)
}
