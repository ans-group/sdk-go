package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetTags retrieves a list of tags
func (s *Service) GetTags(parameters connection.APIRequestParameters) ([]Tag, error) {
	return connection.InvokeRequestAll(s.GetTagsPaginated, parameters)
}

// GetTagsPaginated retrieves a paginated list of tags
func (s *Service) GetTagsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Tag], error) {
	body, err := s.getTagsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetTagsPaginated), err
}

func (s *Service) getTagsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Tag], error) {
	body := &connection.APIResponseBodyData[[]Tag]{}

	response, err := s.connection.Get("/ecloud/v2/tags", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetTag retrieves a single tag by id
func (s *Service) GetTag(tagID string) (Tag, error) {
	body, err := s.getTagResponseBody(tagID)

	return body.Data, err
}

func (s *Service) getTagResponseBody(tagID string) (*connection.APIResponseBodyData[Tag], error) {
	body := &connection.APIResponseBodyData[Tag]{}

	if tagID == "" {
		return body, fmt.Errorf("ecloud: invalid tag id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/tags/%s", tagID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &TagNotFoundError{ID: tagID}
		}

		return nil
	})
}

// CreateTag creates a new tag
func (s *Service) CreateTag(req CreateTagRequest) (string, error) {
	body, err := s.createTagResponseBody(req)

	return body.Data.ID, err
}

func (s *Service) createTagResponseBody(req CreateTagRequest) (*connection.APIResponseBodyData[Tag], error) {
	body := &connection.APIResponseBodyData[Tag]{}

	response, err := s.connection.Post("/ecloud/v2/tags", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchTag patches a tag
func (s *Service) PatchTag(tagID string, req PatchTagRequest) error {
	_, err := s.patchTagResponseBody(tagID, req)

	return err
}

func (s *Service) patchTagResponseBody(tagID string, req PatchTagRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if tagID == "" {
		return body, fmt.Errorf("ecloud: invalid tag id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ecloud/v2/tags/%s", tagID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &TagNotFoundError{ID: tagID}
		}

		return nil
	})
}

// DeleteTag removes a tag
func (s *Service) DeleteTag(tagID string) error {
	_, err := s.deleteTagResponseBody(tagID)

	return err
}

func (s *Service) deleteTagResponseBody(tagID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if tagID == "" {
		return body, fmt.Errorf("ecloud: invalid tag id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v2/tags/%s", tagID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &TagNotFoundError{ID: tagID}
		}

		return nil
	})
}
