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
