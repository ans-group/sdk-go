package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) monitoringGatewayRes() *resource.Resource[MonitoringGateway, string] {
	return resource.NewStringResource[MonitoringGateway](s.connection, "/ecloud/v2/monitoring-gateways", "monitoring gateway", func(id string) error {
		return &MonitoringGatewayNotFoundError{ID: id}
	})
}

// GetMonitoringGateways retrieves a list of monitoring gateways
func (s *Service) GetMonitoringGateways(parameters connection.APIRequestParameters) ([]MonitoringGateway, error) {
	return s.monitoringGatewayRes().List(parameters)
}

// GetMonitoringGatewaysPaginated retrieves a paginated list of monitoring gateways
func (s *Service) GetMonitoringGatewaysPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[MonitoringGateway], error) {
	return s.monitoringGatewayRes().ListPaginated(parameters)
}

// GetMonitoringGateway retrieves a single monitoring gateway by ID
func (s *Service) GetMonitoringGateway(gatewayID string) (MonitoringGateway, error) {
	return s.monitoringGatewayRes().Get(gatewayID)
}

// CreateMonitoringGateway creates a new monitoring gateway
func (s *Service) CreateMonitoringGateway(req CreateMonitoringGatewayRequest) (TaskReference, error) {
	body, err := connection.Post[TaskReference](s.connection, "/ecloud/v2/monitoring-gateways", &req)
	return body.Data, err
}

// PatchMonitoringGateway patches a monitoring gateway
func (s *Service) PatchMonitoringGateway(gatewayID string, req PatchMonitoringGatewayRequest) (TaskReference, error) {
	if gatewayID == "" {
		return TaskReference{}, fmt.Errorf("invalid gateway id")
	}
	body, err := connection.Patch[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/monitoring-gateways/%s", gatewayID), &req, connection.NotFoundResponseHandler(&MonitoringGatewayNotFoundError{ID: gatewayID}))
	return body.Data, err
}

// DeleteMonitoringGateway deletes a monitoring gateway
func (s *Service) DeleteMonitoringGateway(gatewayID string) (string, error) {
	if gatewayID == "" {
		return "", fmt.Errorf("invalid gateway id")
	}
	body, err := connection.Delete[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/monitoring-gateways/%s", gatewayID), nil, connection.NotFoundResponseHandler(&MonitoringGatewayNotFoundError{ID: gatewayID}))
	return body.Data.TaskID, err
}
