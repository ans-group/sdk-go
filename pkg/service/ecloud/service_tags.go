package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetTagsV2 retrieves a list of tags
func (s *Service) GetTagsV2(parameters connection.APIRequestParameters) ([]TagV2, error) {
	return connection.InvokeRequestAll(s.GetTagsV2Paginated, parameters)
}

// GetTagsV2Paginated retrieves a paginated list of tags
func (s *Service) GetTagsV2Paginated(parameters connection.APIRequestParameters) (*connection.Paginated[TagV2], error) {
	body, err := s.getTagsV2PaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetTagsV2Paginated), err
}

func (s *Service) getTagsV2PaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]TagV2], error) {
	body := &connection.APIResponseBodyData[[]TagV2]{}

	response, err := s.connection.Get("/ecloud/v2/tags", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetTagV2 retrieves a single tag by id
func (s *Service) GetTagV2(tagID string) (TagV2, error) {
	body, err := s.getTagV2ResponseBody(tagID)

	return body.Data, err
}

func (s *Service) getTagV2ResponseBody(tagID string) (*connection.APIResponseBodyData[TagV2], error) {
	body := &connection.APIResponseBodyData[TagV2]{}

	if tagID == "" {
		return body, fmt.Errorf("ecloud: invalid tag id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/tags/%s", tagID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &TagV2NotFoundError{ID: tagID}
		}

		return nil
	})
}

// CreateTagV2 creates a new tag
func (s *Service) CreateTagV2(req CreateTagV2Request) (string, error) {
	body, err := s.createTagV2ResponseBody(req)

	return body.Data.ID, err
}

func (s *Service) createTagV2ResponseBody(req CreateTagV2Request) (*connection.APIResponseBodyData[TagV2], error) {
	body := &connection.APIResponseBodyData[TagV2]{}

	response, err := s.connection.Post("/ecloud/v2/tags", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchTagV2 patches a tag
func (s *Service) PatchTagV2(tagID string, req PatchTagV2Request) error {
	_, err := s.patchTagV2ResponseBody(tagID, req)

	return err
}

func (s *Service) patchTagV2ResponseBody(tagID string, req PatchTagV2Request) (*connection.APIResponseBody, error) {
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
			return &TagV2NotFoundError{ID: tagID}
		}

		return nil
	})
}

// DeleteTagV2 removes a tag
func (s *Service) DeleteTagV2(tagID string) error {
	_, err := s.deleteTagV2ResponseBody(tagID)

	return err
}

func (s *Service) deleteTagV2ResponseBody(tagID string) (*connection.APIResponseBody, error) {
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
			return &TagV2NotFoundError{ID: tagID}
		}

		return nil
	})
}
