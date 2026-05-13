package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) imageRes() *resource.Resource[Image, string] {
	return resource.NewStringResource[Image](s.connection, "/ecloud/v2/images", "image", func(id string) error {
		return &ImageNotFoundError{ID: id}
	})
}

// GetImages retrieves a list of images
func (s *Service) GetImages(parameters connection.APIRequestParameters) ([]Image, error) {
	return s.imageRes().List(parameters)
}

// GetImagesPaginated retrieves a paginated list of images
func (s *Service) GetImagesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Image], error) {
	return s.imageRes().ListPaginated(parameters)
}

// GetImage retrieves a single Image by ID
func (s *Service) GetImage(imageID string) (Image, error) {
	return s.imageRes().Get(imageID)
}

// UpdateImage removes a single Image by ID
func (s *Service) UpdateImage(imageID string, req UpdateImageRequest) (TaskReference, error) {
	if imageID == "" {
		return TaskReference{}, fmt.Errorf("invalid image id")
	}
	body, err := connection.Patch[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/images/%s", imageID), &req, connection.NotFoundResponseHandler(&ImageNotFoundError{ID: imageID}))
	return body.Data, err
}

// DeleteImage removes a single Image by ID
func (s *Service) DeleteImage(imageID string) (string, error) {
	if imageID == "" {
		return "", fmt.Errorf("invalid image id")
	}
	body, err := connection.Delete[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/images/%s", imageID), nil, connection.NotFoundResponseHandler(&ImageNotFoundError{ID: imageID}))
	return body.Data.TaskID, err
}

func (s *Service) imageParameterRes() *resource.SubResourceList[ImageParameter, string] {
	return resource.NewStringSubResourceList[ImageParameter](s.connection,
		func(imageID string) string { return fmt.Sprintf("/ecloud/v2/images/%s/parameters", imageID) },
		"image", "id", func(imageID string) error { return &ImageNotFoundError{ID: imageID} })
}

// GetImageParameters retrieves a list of parameters
func (s *Service) GetImageParameters(imageID string, parameters connection.APIRequestParameters) ([]ImageParameter, error) {
	return s.imageParameterRes().List(imageID, parameters)
}

// GetImageParametersPaginated retrieves a paginated list of domains
func (s *Service) GetImageParametersPaginated(imageID string, parameters connection.APIRequestParameters) (*connection.Paginated[ImageParameter], error) {
	return s.imageParameterRes().ListPaginated(imageID, parameters)
}

func (s *Service) imageMetadataRes() *resource.SubResourceList[ImageMetadata, string] {
	return resource.NewStringSubResourceList[ImageMetadata](s.connection,
		func(imageID string) string { return fmt.Sprintf("/ecloud/v2/images/%s/metadata", imageID) },
		"image", "id", func(imageID string) error { return &ImageNotFoundError{ID: imageID} })
}

// GetImageMetadata retrieves a list of metadata
func (s *Service) GetImageMetadata(imageID string, parameters connection.APIRequestParameters) ([]ImageMetadata, error) {
	return s.imageMetadataRes().List(imageID, parameters)
}

// GetImageMetadataPaginated retrieves a paginated list of domains
func (s *Service) GetImageMetadataPaginated(imageID string, parameters connection.APIRequestParameters) (*connection.Paginated[ImageMetadata], error) {
	return s.imageMetadataRes().ListPaginated(imageID, parameters)
}
