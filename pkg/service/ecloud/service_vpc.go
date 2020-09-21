package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetVPCs retrieves a list of vpcs
func (s *Service) GetVPCs(parameters connection.APIRequestParameters) ([]VPC, error) {
	var sites []VPC

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetVPCsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, site := range response.(*PaginatedVPC).Items {
			sites = append(sites, site)
		}
	}

	return sites, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetVPCsPaginated retrieves a paginated list of vpcs
func (s *Service) GetVPCsPaginated(parameters connection.APIRequestParameters) (*PaginatedVPC, error) {
	body, err := s.getVPCsPaginatedResponseBody(parameters)

	return NewPaginatedVPC(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetVPCsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getVPCsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetVPCSliceResponseBody, error) {
	body := &GetVPCSliceResponseBody{}

	response, err := s.connection.Get("/ecloud/v2/vpcs", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetVPC retrieves a single vpc by id
func (s *Service) GetVPC(vpcID string) (VPC, error) {
	body, err := s.getVPCResponseBody(vpcID)

	return body.Data, err
}

func (s *Service) getVPCResponseBody(vpcID string) (*GetVPCResponseBody, error) {
	body := &GetVPCResponseBody{}

	if vpcID == "" {
		return body, fmt.Errorf("invalid vpc id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/vpcs/%s", vpcID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VPCNotFoundError{ID: vpcID}
		}

		return nil
	})
}

// CreateVPC creates a new VPC
func (s *Service) CreateVPC(req CreateVPCRequest) (string, error) {
	body, err := s.createVPCResponseBody(req)

	return body.Data.ID, err
}

func (s *Service) createVPCResponseBody(req CreateVPCRequest) (*GetVPCResponseBody, error) {
	body := &GetVPCResponseBody{}

	response, err := s.connection.Post("/ecloud/v2/vpcs", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchVPC patches a VPC
func (s *Service) PatchVPC(vpcID string, req PatchVPCRequest) error {
	_, err := s.patchVPCResponseBody(vpcID, req)

	return err
}

func (s *Service) patchVPCResponseBody(vpcID string, req PatchVPCRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if vpcID == "" {
		return body, fmt.Errorf("invalid vpc id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ecloud/v2/vpcs/%s", vpcID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VPCNotFoundError{ID: vpcID}
		}

		return nil
	})
}

// DeleteVPC deletes a VPC
func (s *Service) DeleteVPC(vpcID string) error {
	_, err := s.deleteVPCResponseBody(vpcID)

	return err
}

func (s *Service) deleteVPCResponseBody(vpcID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if vpcID == "" {
		return body, fmt.Errorf("invalid vpc id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v2/vpcs/%s", vpcID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VPCNotFoundError{ID: vpcID}
		}

		return nil
	})
}
