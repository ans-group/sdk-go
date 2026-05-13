package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) vpnServiceRes() *resource.Resource[VPNService, string] {
	return resource.NewStringResource[VPNService](s.connection, "/ecloud/v2/vpn-services", "vpn service", func(id string) error {
		return &VPNServiceNotFoundError{ID: id}
	})
}

// GetVPNServices retrieves a list of VPN services
func (s *Service) GetVPNServices(parameters connection.APIRequestParameters) ([]VPNService, error) {
	return s.vpnServiceRes().List(parameters)
}

// GetVPNServicesPaginated retrieves a paginated list of VPN services
func (s *Service) GetVPNServicesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[VPNService], error) {
	return s.vpnServiceRes().ListPaginated(parameters)
}

// GetVPNService retrieves a single VPN service by id
func (s *Service) GetVPNService(serviceID string) (VPNService, error) {
	return s.vpnServiceRes().Get(serviceID)
}

// CreateVPNService creates a new VPN service
func (s *Service) CreateVPNService(req CreateVPNServiceRequest) (TaskReference, error) {
	body, err := connection.Post[TaskReference](s.connection, "/ecloud/v2/vpn-services", &req)
	return body.Data, err
}

// PatchVPNService patches a VPN service
func (s *Service) PatchVPNService(serviceID string, req PatchVPNServiceRequest) (TaskReference, error) {
	if serviceID == "" {
		return TaskReference{}, fmt.Errorf("invalid service id")
	}
	body, err := connection.Patch[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/vpn-services/%s", serviceID), &req, connection.NotFoundResponseHandler(&VPNServiceNotFoundError{ID: serviceID}))
	return body.Data, err
}

// DeleteVPNService deletes a VPN service
func (s *Service) DeleteVPNService(serviceID string) (string, error) {
	if serviceID == "" {
		return "", fmt.Errorf("invalid service id")
	}
	body, err := connection.Delete[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/vpn-services/%s", serviceID), nil, connection.NotFoundResponseHandler(&VPNServiceNotFoundError{ID: serviceID}))
	return body.Data.TaskID, err
}

func (s *Service) vpnServiceTasksRes() *resource.SubResourceList[Task, string] {
	return resource.NewStringSubResourceList[Task](s.connection,
		func(serviceID string) string { return fmt.Sprintf("/ecloud/v2/vpn-services/%s/tasks", serviceID) },
		"vpn service", "id", func(serviceID string) error { return &VPNServiceNotFoundError{ID: serviceID} })
}

// GetVPNServiceTasks retrieves a list of VPN service tasks
func (s *Service) GetVPNServiceTasks(serviceID string, parameters connection.APIRequestParameters) ([]Task, error) {
	return s.vpnServiceTasksRes().List(serviceID, parameters)
}

// GetVPNServiceTasksPaginated retrieves a paginated list of VPN service tasks
func (s *Service) GetVPNServiceTasksPaginated(serviceID string, parameters connection.APIRequestParameters) (*connection.Paginated[Task], error) {
	return s.vpnServiceTasksRes().ListPaginated(serviceID, parameters)
}
