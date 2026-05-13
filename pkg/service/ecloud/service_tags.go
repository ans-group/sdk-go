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
	body, err := connection.Get[[]Tag](s.connection, "/ecloud/v2/tags", parameters)
	return connection.NewPaginated(body, parameters, s.GetTagsPaginated), err
}

// GetTag retrieves a single tag by id
func (s *Service) GetTag(tagID string) (Tag, error) {
	if tagID == "" {
		return Tag{}, fmt.Errorf("ecloud: invalid tag id")
	}
	body, err := connection.Get[Tag](s.connection, fmt.Sprintf("/ecloud/v2/tags/%s", tagID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&TagNotFoundError{ID: tagID}))
	return body.Data, err
}

// CreateTag creates a new tag
func (s *Service) CreateTag(req CreateTagRequest) (string, error) {
	body, err := connection.Post[Tag](s.connection, "/ecloud/v2/tags", &req)
	return body.Data.ID, err
}

// PatchTag patches a tag
func (s *Service) PatchTag(tagID string, req PatchTagRequest) error {
	if tagID == "" {
		return fmt.Errorf("ecloud: invalid tag id")
	}
	return connection.PatchRaw(s.connection, fmt.Sprintf("/ecloud/v2/tags/%s", tagID), &req, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&TagNotFoundError{ID: tagID}))
}

// DeleteTag removes a tag
func (s *Service) DeleteTag(tagID string) error {
	if tagID == "" {
		return fmt.Errorf("ecloud: invalid tag id")
	}
	return connection.DeleteRaw(s.connection, fmt.Sprintf("/ecloud/v2/tags/%s", tagID), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&TagNotFoundError{ID: tagID}))
}
