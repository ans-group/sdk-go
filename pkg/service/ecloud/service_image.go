package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetImages retrieves a list of images
func (s *Service) GetImages(parameters connection.APIRequestParameters) ([]Image, error) {
	return connection.InvokeRequestAll(s.GetImagesPaginated, parameters)
}

// GetImagesPaginated retrieves a paginated list of images
func (s *Service) GetImagesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Image], error) {
	body, err := connection.Get[[]Image](s.connection, "/ecloud/v2/images", parameters)
	return connection.NewPaginated(body, parameters, s.GetImagesPaginated), err
}

// GetImage retrieves a single Image by ID
func (s *Service) GetImage(imageID string) (Image, error) {
	if imageID == "" {
		return Image{}, fmt.Errorf("invalid image id")
	}
	body, err := connection.Get[Image](s.connection, fmt.Sprintf("/ecloud/v2/images/%s", imageID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&ImageNotFoundError{ID: imageID}))
	return body.Data, err
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

// GetImageParameters retrieves a list of parameters
func (s *Service) GetImageParameters(imageID string, parameters connection.APIRequestParameters) ([]ImageParameter, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[ImageParameter], error) {
		return s.GetImageParametersPaginated(imageID, p)
	}, parameters)
}

// GetImageParametersPaginated retrieves a paginated list of domains
func (s *Service) GetImageParametersPaginated(imageID string, parameters connection.APIRequestParameters) (*connection.Paginated[ImageParameter], error) {
	if imageID == "" {
		return nil, fmt.Errorf("invalid image id")
	}
	body, err := connection.Get[[]ImageParameter](s.connection, fmt.Sprintf("/ecloud/v2/images/%s/parameters", imageID), parameters, connection.NotFoundResponseHandler(&ImageNotFoundError{ID: imageID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[ImageParameter], error) {
		return s.GetImageParametersPaginated(imageID, p)
	}), err
}

// GetImageMetadata retrieves a list of metadata
func (s *Service) GetImageMetadata(imageID string, parameters connection.APIRequestParameters) ([]ImageMetadata, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[ImageMetadata], error) {
		return s.GetImageMetadataPaginated(imageID, p)
	}, parameters)
}

// GetImageMetadataPaginated retrieves a paginated list of domains
func (s *Service) GetImageMetadataPaginated(imageID string, parameters connection.APIRequestParameters) (*connection.Paginated[ImageMetadata], error) {
	if imageID == "" {
		return nil, fmt.Errorf("invalid image id")
	}
	body, err := connection.Get[[]ImageMetadata](s.connection, fmt.Sprintf("/ecloud/v2/images/%s/metadata", imageID), parameters, connection.NotFoundResponseHandler(&ImageNotFoundError{ID: imageID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[ImageMetadata], error) {
		return s.GetImageMetadataPaginated(imageID, p)
	}), err
}
