package loadbalancer

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetTargets retrieves a list of targets
func (s *Service) GetTargets(parameters connection.APIRequestParameters) ([]Target, error) {
	var sites []Target

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetTargetsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, site := range response.(*PaginatedTarget).Items {
			sites = append(sites, site)
		}
	}

	return sites, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetTargetsPaginated retrieves a paginated list of targets
func (s *Service) GetTargetsPaginated(parameters connection.APIRequestParameters) (*PaginatedTarget, error) {
	body, err := s.getTargetsPaginatedResponseBody(parameters)

	return NewPaginatedTarget(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetTargetsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getTargetsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetTargetSliceResponseBody, error) {
	body := &GetTargetSliceResponseBody{}

	response, err := s.connection.Get("/loadbalancers/v2/targets", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetTarget retrieves a single target by id
func (s *Service) GetTarget(targetID string) (Target, error) {
	body, err := s.getTargetResponseBody(targetID)

	return body.Data, err
}

func (s *Service) getTargetResponseBody(targetID string) (*GetTargetResponseBody, error) {
	body := &GetTargetResponseBody{}

	if targetID == "" {
		return body, fmt.Errorf("invalid target id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/loadbalancers/v2/targets/%s", targetID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &TargetNotFoundError{ID: targetID}
		}

		return nil
	})
}

// CreateTarget creates a Target
func (s *Service) CreateTarget(req CreateTargetRequest) (string, error) {
	body, err := s.createTargetResponseBody(req)

	return body.Data.ID, err
}

func (s *Service) createTargetResponseBody(req CreateTargetRequest) (*GetTargetResponseBody, error) {
	body := &GetTargetResponseBody{}

	response, err := s.connection.Post("/loadbalancers/v2/targets", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchTarget patches a Target
func (s *Service) PatchTarget(targetID string, req PatchTargetRequest) error {
	_, err := s.patchTargetResponseBody(targetID, req)

	return err
}

func (s *Service) patchTargetResponseBody(targetID string, req PatchTargetRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if targetID == "" {
		return body, fmt.Errorf("invalid target id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/loadbalancers/v2/targets/%s", targetID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &TargetNotFoundError{ID: targetID}
		}

		return nil
	})
}

// DeleteTarget deletes a Target
func (s *Service) DeleteTarget(targetID string) error {
	_, err := s.deleteTargetResponseBody(targetID)

	return err
}

func (s *Service) deleteTargetResponseBody(targetID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if targetID == "" {
		return body, fmt.Errorf("invalid target id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/loadbalancers/v2/targets/%s", targetID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &TargetNotFoundError{ID: targetID}
		}

		return nil
	})
}
