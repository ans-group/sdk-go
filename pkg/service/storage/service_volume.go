package storage

import (
	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) volumeRes() *resource.Resource[Volume, int] {
	return resource.NewIntResource[Volume](s.connection, "/ukfast-storage/v1/volumes", "volume",
		func(id int) error { return &VolumeNotFoundError{ID: id} })
}

// GetVolumes retrieves a list of volumes
func (s *Service) GetVolumes(parameters connection.APIRequestParameters) ([]Volume, error) {
	return s.volumeRes().List(parameters)
}

// GetVolumesPaginated retrieves a paginated list of volumes
func (s *Service) GetVolumesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Volume], error) {
	return s.volumeRes().ListPaginated(parameters)
}

// GetVolume retrieves a single volume by id
func (s *Service) GetVolume(volumeID int) (Volume, error) {
	return s.volumeRes().Get(volumeID)
}
