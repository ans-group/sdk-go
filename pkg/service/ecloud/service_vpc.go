package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) vpcRes() *resource.Resource[VPC, string] {
	return resource.NewStringResource[VPC](s.connection, "/ecloud/v2/vpcs", "vpc", func(id string) error {
		return &VPCNotFoundError{ID: id}
	})
}

// GetVPCs retrieves a list of vpcs
func (s *Service) GetVPCs(parameters connection.APIRequestParameters) ([]VPC, error) {
	return s.vpcRes().List(parameters)
}

// GetVPCsPaginated retrieves a paginated list of vpcs
func (s *Service) GetVPCsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[VPC], error) {
	return s.vpcRes().ListPaginated(parameters)
}

// GetVPC retrieves a single vpc by id
func (s *Service) GetVPC(vpcID string) (VPC, error) {
	return s.vpcRes().Get(vpcID)
}

// CreateVPC creates a new VPC
func (s *Service) CreateVPC(req CreateVPCRequest) (string, error) {
	data, err := s.vpcRes().Create(&req)
	return data.ID, err
}

// PatchVPC patches a VPC
func (s *Service) PatchVPC(vpcID string, req PatchVPCRequest) error {
	return s.vpcRes().Patch(vpcID, &req)
}

// DeleteVPC deletes a VPC
func (s *Service) DeleteVPC(vpcID string) error {
	return s.vpcRes().Delete(vpcID)
}

// DeployVPCDefaults deploys default resources for specified VPC
func (s *Service) DeployVPCDefaults(vpcID string) error {
	if vpcID == "" {
		return fmt.Errorf("invalid vpc id")
	}
	return connection.PostRaw(s.connection, fmt.Sprintf("/ecloud/v2/vpcs/%s/deploy-defaults", vpcID), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&VPCNotFoundError{ID: vpcID}))
}

// GetVPCVolumes retrieves a list of firewall rule volumes
func (s *Service) GetVPCVolumes(vpcID string, parameters connection.APIRequestParameters) ([]Volume, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Volume], error) {
		return s.GetVPCVolumesPaginated(vpcID, p)
	}, parameters)
}

// GetVPCVolumesPaginated retrieves a paginated list of firewall rule volumes
func (s *Service) GetVPCVolumesPaginated(vpcID string, parameters connection.APIRequestParameters) (*connection.Paginated[Volume], error) {
	if vpcID == "" {
		return nil, fmt.Errorf("invalid vpc id")
	}
	body, err := connection.Get[[]Volume](s.connection, fmt.Sprintf("/ecloud/v2/vpcs/%s/volumes", vpcID), parameters, connection.NotFoundResponseHandler(&VPCNotFoundError{ID: vpcID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Volume], error) {
		return s.GetVPCVolumesPaginated(vpcID, p)
	}), err
}

// GetVPCInstances retrieves a list of firewall rule instances
func (s *Service) GetVPCInstances(vpcID string, parameters connection.APIRequestParameters) ([]Instance, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Instance], error) {
		return s.GetVPCInstancesPaginated(vpcID, p)
	}, parameters)
}

// GetVPCInstancesPaginated retrieves a paginated list of firewall rule instances
func (s *Service) GetVPCInstancesPaginated(vpcID string, parameters connection.APIRequestParameters) (*connection.Paginated[Instance], error) {
	if vpcID == "" {
		return nil, fmt.Errorf("invalid vpc id")
	}
	body, err := connection.Get[[]Instance](s.connection, fmt.Sprintf("/ecloud/v2/vpcs/%s/instances", vpcID), parameters, connection.NotFoundResponseHandler(&VPCNotFoundError{ID: vpcID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Instance], error) {
		return s.GetVPCInstancesPaginated(vpcID, p)
	}), err
}

// GetVPCTasks retrieves a list of VPC tasks
func (s *Service) GetVPCTasks(vpcID string, parameters connection.APIRequestParameters) ([]Task, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetVPCTasksPaginated(vpcID, p)
	}, parameters)
}

// GetVPCTasksPaginated retrieves a paginated list of VPC tasks
func (s *Service) GetVPCTasksPaginated(vpcID string, parameters connection.APIRequestParameters) (*connection.Paginated[Task], error) {
	if vpcID == "" {
		return nil, fmt.Errorf("invalid vpc id")
	}
	body, err := connection.Get[[]Task](s.connection, fmt.Sprintf("/ecloud/v2/vpcs/%s/tasks", vpcID), parameters, connection.NotFoundResponseHandler(&VPCNotFoundError{ID: vpcID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetVPCTasksPaginated(vpcID, p)
	}), err
}
