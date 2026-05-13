package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) volumeGroupRes() *resource.Resource[VolumeGroup, string] {
	return resource.NewStringResource[VolumeGroup](s.connection, "/ecloud/v2/volume-groups", "volume group", func(id string) error {
		return &VolumeGroupNotFoundError{ID: id}
	})
}

// GetVolumeGroups retrieves a list of volume groups
func (s *Service) GetVolumeGroups(parameters connection.APIRequestParameters) ([]VolumeGroup, error) {
	return s.volumeGroupRes().List(parameters)
}

// GetVolumeGroupsPaginated retrieves a paginated list of volume groups
func (s *Service) GetVolumeGroupsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[VolumeGroup], error) {
	return s.volumeGroupRes().ListPaginated(parameters)
}

// GetVolumeGroup retrieves a single volumeGroup by id
func (s *Service) GetVolumeGroup(volumeGroupID string) (VolumeGroup, error) {
	return s.volumeGroupRes().Get(volumeGroupID)
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

func (s *Service) volumeGroupVolumeRes() *resource.SubResourceList[Volume, string] {
	return resource.NewStringSubResourceList[Volume](s.connection,
		func(volumeGroupID string) string {
			return fmt.Sprintf("/ecloud/v2/volume-groups/%s/volumes", volumeGroupID)
		},
		"volume group", "id", func(volumeGroupID string) error { return &VolumeGroupNotFoundError{ID: volumeGroupID} })
}

// GetVolumeGroupVolumes retrieves a list of VolumeGroup volumes
func (s *Service) GetVolumeGroupVolumes(volumeGroupID string, parameters connection.APIRequestParameters) ([]Volume, error) {
	return s.volumeGroupVolumeRes().List(volumeGroupID, parameters)
}

// GetVolumeGroupVolumesPaginated retrieves a paginated list of VolumeGroup volumes
func (s *Service) GetVolumeGroupVolumesPaginated(volumeGroupID string, parameters connection.APIRequestParameters) (*connection.Paginated[Volume], error) {
	return s.volumeGroupVolumeRes().ListPaginated(volumeGroupID, parameters)
}
