package storage

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetVolumes retrieves a list of volumes
func (s *Service) GetVolumes(parameters connection.APIRequestParameters) ([]Volume, error) {
	return connection.InvokeRequestAll(s.GetVolumesPaginated, parameters)
}

// GetVolumesPaginated retrieves a paginated list of volumes
func (s *Service) GetVolumesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Volume], error) {
	body, err := connection.Get[[]Volume](s.connection, "/ukfast-storage/v1/volumes", parameters)
	return connection.NewPaginated(body, parameters, s.GetVolumesPaginated), err
}

// GetVolume retrieves a single volume by id
func (s *Service) GetVolume(volumeID int) (Volume, error) {
	if volumeID < 1 {
		return Volume{}, fmt.Errorf("invalid volume id")
	}
	body, err := connection.Get[Volume](s.connection, fmt.Sprintf("/ukfast-storage/v1/volumes/%d", volumeID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&VolumeNotFoundError{ID: volumeID}))
	return body.Data, err
}
