package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) volumeRes() *resource.Resource[Volume, string] {
	return resource.NewStringResource[Volume](s.connection, "/ecloud/v2/volumes", "volume", func(id string) error {
		return &VolumeNotFoundError{ID: id}
	})
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
func (s *Service) GetVolume(volumeID string) (Volume, error) {
	return s.volumeRes().Get(volumeID)
}

// CreateVolume creates a volume
func (s *Service) CreateVolume(req CreateVolumeRequest) (TaskReference, error) {
	body, err := connection.Post[TaskReference](s.connection, "/ecloud/v2/volumes", &req)
	return body.Data, err
}

// PatchVolume patches a Volume
func (s *Service) PatchVolume(volumeID string, req PatchVolumeRequest) (TaskReference, error) {
	if volumeID == "" {
		return TaskReference{}, fmt.Errorf("invalid volume id")
	}
	body, err := connection.Patch[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/volumes/%s", volumeID), &req, connection.NotFoundResponseHandler(&VolumeNotFoundError{ID: volumeID}))
	return body.Data, err
}

// DeleteVolume deletes a Volume
func (s *Service) DeleteVolume(volumeID string) (string, error) {
	if volumeID == "" {
		return "", fmt.Errorf("invalid volume id")
	}
	body, err := connection.Delete[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/volumes/%s", volumeID), nil, connection.NotFoundResponseHandler(&VolumeNotFoundError{ID: volumeID}))
	return body.Data.TaskID, err
}

func (s *Service) volumeInstanceRes() *resource.SubResourceList[Instance, string] {
	return resource.NewStringSubResourceList[Instance](s.connection,
		func(volumeID string) string { return fmt.Sprintf("/ecloud/v2/volumes/%s/instances", volumeID) },
		"volume", "id", func(volumeID string) error { return &VolumeNotFoundError{ID: volumeID} })
}

// GetVolumeInstances retrieves a list of volume instances
func (s *Service) GetVolumeInstances(volumeID string, parameters connection.APIRequestParameters) ([]Instance, error) {
	return s.volumeInstanceRes().List(volumeID, parameters)
}

// GetVolumeInstancesPaginated retrieves a paginated list of volume instances
func (s *Service) GetVolumeInstancesPaginated(volumeID string, parameters connection.APIRequestParameters) (*connection.Paginated[Instance], error) {
	return s.volumeInstanceRes().ListPaginated(volumeID, parameters)
}

func (s *Service) volumeTasksRes() *resource.SubResourceList[Task, string] {
	return resource.NewStringSubResourceList[Task](s.connection,
		func(volumeID string) string { return fmt.Sprintf("/ecloud/v2/volumes/%s/tasks", volumeID) },
		"volume", "id", func(volumeID string) error { return &VolumeNotFoundError{ID: volumeID} })
}

// GetVolumeTasks retrieves a list of Volume tasks
func (s *Service) GetVolumeTasks(volumeID string, parameters connection.APIRequestParameters) ([]Task, error) {
	return s.volumeTasksRes().List(volumeID, parameters)
}

// GetVolumeTasksPaginated retrieves a paginated list of Volume tasks
func (s *Service) GetVolumeTasksPaginated(volumeID string, parameters connection.APIRequestParameters) (*connection.Paginated[Task], error) {
	return s.volumeTasksRes().ListPaginated(volumeID, parameters)
}
