package ecloud

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
	body, err := connection.Get[[]Volume](s.connection, "/ecloud/v2/volumes", parameters)
	return connection.NewPaginated(body, parameters, s.GetVolumesPaginated), err
}

// GetVolume retrieves a single volume by id
func (s *Service) GetVolume(volumeID string) (Volume, error) {
	if volumeID == "" {
		return Volume{}, fmt.Errorf("invalid volume id")
	}
	body, err := connection.Get[Volume](s.connection, fmt.Sprintf("/ecloud/v2/volumes/%s", volumeID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&VolumeNotFoundError{ID: volumeID}))
	return body.Data, err
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

// GetVolumeInstances retrieves a list of volume instances
func (s *Service) GetVolumeInstances(volumeID string, parameters connection.APIRequestParameters) ([]Instance, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Instance], error) {
		return s.GetVolumeInstancesPaginated(volumeID, p)
	}, parameters)
}

// GetVolumeInstancesPaginated retrieves a paginated list of volume instances
func (s *Service) GetVolumeInstancesPaginated(volumeID string, parameters connection.APIRequestParameters) (*connection.Paginated[Instance], error) {
	if volumeID == "" {
		return nil, fmt.Errorf("invalid volume id")
	}
	body, err := connection.Get[[]Instance](s.connection, fmt.Sprintf("/ecloud/v2/volumes/%s/instances", volumeID), parameters, connection.NotFoundResponseHandler(&VolumeNotFoundError{ID: volumeID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Instance], error) {
		return s.GetVolumeInstancesPaginated(volumeID, p)
	}), err
}

// GetVolumeTasks retrieves a list of Volume tasks
func (s *Service) GetVolumeTasks(volumeID string, parameters connection.APIRequestParameters) ([]Task, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetVolumeTasksPaginated(volumeID, p)
	}, parameters)
}

// GetVolumeTasksPaginated retrieves a paginated list of Volume tasks
func (s *Service) GetVolumeTasksPaginated(volumeID string, parameters connection.APIRequestParameters) (*connection.Paginated[Task], error) {
	if volumeID == "" {
		return nil, fmt.Errorf("invalid volume id")
	}
	body, err := connection.Get[[]Task](s.connection, fmt.Sprintf("/ecloud/v2/volumes/%s/tasks", volumeID), parameters, connection.NotFoundResponseHandler(&VolumeNotFoundError{ID: volumeID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetVolumeTasksPaginated(volumeID, p)
	}), err
}
