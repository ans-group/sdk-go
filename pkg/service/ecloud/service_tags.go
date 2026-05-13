package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) tagRes() *resource.Resource[Tag, string] {
	return resource.NewStringResource[Tag](s.connection, "/ecloud/v2/tags", "tag", func(id string) error {
		return &TagNotFoundError{ID: id}
	})
}

// GetTags retrieves a list of tags
func (s *Service) GetTags(parameters connection.APIRequestParameters) ([]Tag, error) {
	return s.tagRes().List(parameters)
}

// GetTagsPaginated retrieves a paginated list of tags
func (s *Service) GetTagsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Tag], error) {
	return s.tagRes().ListPaginated(parameters)
}

// GetTag retrieves a single tag by id
func (s *Service) GetTag(tagID string) (Tag, error) {
	if tagID == "" {
		return Tag{}, fmt.Errorf("ecloud: invalid tag id")
	}
	return s.tagRes().Get(tagID)
}

// CreateTag creates a new tag
func (s *Service) CreateTag(req CreateTagRequest) (string, error) {
	data, err := s.tagRes().Create(&req)
	return data.ID, err
}

// PatchTag patches a tag
func (s *Service) PatchTag(tagID string, req PatchTagRequest) error {
	if tagID == "" {
		return fmt.Errorf("ecloud: invalid tag id")
	}
	return s.tagRes().Patch(tagID, &req)
}

// DeleteTag removes a tag
func (s *Service) DeleteTag(tagID string) error {
	if tagID == "" {
		return fmt.Errorf("ecloud: invalid tag id")
	}
	return s.tagRes().Delete(tagID)
}
