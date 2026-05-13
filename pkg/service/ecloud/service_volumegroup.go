package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetVolumeGroups retrieves a list of volume groups
func (s *Service) GetVolumeGroups(parameters connection.APIRequestParameters) ([]VolumeGroup, error) {
	return connection.InvokeRequestAll(s.GetVolumeGroupsPaginated, parameters)
}

// GetVolumeGroupsPaginated retrieves a paginated list of volume groups
func (s *Service) GetVolumeGroupsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[VolumeGroup], error) {
	body, err := connection.Get[[]VolumeGroup](s.connection, "/ecloud/v2/volume-groups", parameters)
	return connection.NewPaginated(body, parameters, s.GetVolumeGroupsPaginated), err
}

// GetVolumeGroup retrieves a single volumeGroup by id
func (s *Service) GetVolumeGroup(volumeGroupID string) (VolumeGroup, error) {
	if volumeGroupID == "" {
		return VolumeGroup{}, fmt.Errorf("invalid volume group id")
	}
	body, err := connection.Get[VolumeGroup](s.connection, fmt.Sprintf("/ecloud/v2/volume-groups/%s", volumeGroupID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&VolumeGroupNotFoundError{ID: volumeGroupID}))
	return body.Data, err
}

// CreateVolumeGroup creates a volumeGroup
func (s *Service) CreateVolumeGroup(req CreateVolumeGroupRequest) (TaskReference, error) {
	body, err := connection.Post[TaskReference](s.connection, "/ecloud/v2/volume-groups", &req)
	return body.Data, err
}

// PatchVolumeGroup patches a volumeGroup
func (s *Service) PatchVolumeGroup(volumeGroupID string, req PatchVolumeGroupRequest) (TaskReference, error) {
	if volumeGroupID == "" {
		return TaskReference{}, fmt.Errorf("invalid volume group id")
	}
	body, err := connection.Patch[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/volume-groups/%s", volumeGroupID), &req, connection.NotFoundResponseHandler(&VolumeGroupNotFoundError{ID: volumeGroupID}))
	return body.Data, err
}

// DeleteVolumeGroup deletes a volumeGroup
func (s *Service) DeleteVolumeGroup(volumeGroupID string) (string, error) {
	if volumeGroupID == "" {
		return "", fmt.Errorf("invalid volume group id")
	}
	body, err := connection.Delete[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/volume-groups/%s", volumeGroupID), nil, connection.NotFoundResponseHandler(&VolumeGroupNotFoundError{ID: volumeGroupID}))
	return body.Data.TaskID, err
}

// GetVolumeGroupVolumes retrieves a list of VolumeGroup volumes
func (s *Service) GetVolumeGroupVolumes(volumeGroupID string, parameters connection.APIRequestParameters) ([]Volume, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Volume], error) {
		return s.GetVolumeGroupVolumesPaginated(volumeGroupID, p)
	}, parameters)
}

// GetVolumeGroupVolumesPaginated retrieves a paginated list of VolumeGroup volumes
func (s *Service) GetVolumeGroupVolumesPaginated(volumeGroupID string, parameters connection.APIRequestParameters) (*connection.Paginated[Volume], error) {
	if volumeGroupID == "" {
		return nil, fmt.Errorf("invalid volume group id")
	}
	body, err := connection.Get[[]Volume](s.connection, fmt.Sprintf("/ecloud/v2/volume-groups/%s/volumes", volumeGroupID), parameters, connection.NotFoundResponseHandler(&VolumeGroupNotFoundError{ID: volumeGroupID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Volume], error) {
		return s.GetVolumeGroupVolumesPaginated(volumeGroupID, p)
	}), err
}
