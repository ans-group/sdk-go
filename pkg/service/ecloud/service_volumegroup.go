package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetVolumeGroups retrieves a list of volume groups
func (s *Service) GetVolumeGroups(parameters connection.APIRequestParameters) ([]VolumeGroup, error) {
	var volGroups []VolumeGroup

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetVolumeGroupsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, volumeGroup := range response.(*PaginatedVolumeGroup).Items {
			volGroups = append(volGroups, volumeGroup)
		}
	}

	return volGroups, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetVolumeGroupsPaginated retrieves a paginated list of volume groups
func (s *Service) GetVolumeGroupsPaginated(parameters connection.APIRequestParameters) (*PaginatedVolumeGroup, error) {
	body, err := s.getVolumeGroupsPaginatedResponseBody(parameters)

	return NewPaginatedVolumeGroup(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetVolumeGroupsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getVolumeGroupsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetVolumeGroupSliceResponseBody, error) {
	body := &GetVolumeGroupSliceResponseBody{}

	response, err := s.connection.Get("/ecloud/v2/volume-groups", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetVolumeGroup retrieves a single volumeGroup by id
func (s *Service) GetVolumeGroup(volumeGroupID string) (VolumeGroup, error) {
	body, err := s.getVolumeGroupResponseBody(volumeGroupID)

	return body.Data, err
}

func (s *Service) getVolumeGroupResponseBody(volumeGroupID string) (*GetVolumeGroupResponseBody, error) {
	body := &GetVolumeGroupResponseBody{}

	if volumeGroupID == "" {
		return body, fmt.Errorf("invalid volume group id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/volume-groups/%s", volumeGroupID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VolumeGroupNotFoundError{ID: volumeGroupID}
		}

		return nil
	})
}

// CreateVolumeGroup creates a volumeGroup
func (s *Service) CreateVolumeGroup(req CreateVolumeGroupRequest) (TaskReference, error) {
	body, err := s.createVolumeGroupResponseBody(req)

	return body.Data, err
}

func (s *Service) createVolumeGroupResponseBody(req CreateVolumeGroupRequest) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

	response, err := s.connection.Post("/ecloud/v2/volume-groups", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchVolumeGroup patches a volumeGroup
func (s *Service) PatchVolumeGroup(volumeGroupID string, req PatchVolumeGroupRequest) (TaskReference, error) {
	body, err := s.patchVolumeGroupResponseBody(volumeGroupID, req)

	return body.Data, err
}

func (s *Service) patchVolumeGroupResponseBody(volumeGroupID string, req PatchVolumeGroupRequest) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

	if volumeGroupID == "" {
		return body, fmt.Errorf("invalid volume group id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ecloud/v2/volume-groups/%s", volumeGroupID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VolumeGroupNotFoundError{ID: volumeGroupID}
		}

		return nil
	})
}

// DeleteVolumeGroup deletes a volumeGroup
func (s *Service) DeleteVolumeGroup(volumeGroupID string) (string, error) {
	body, err := s.deleteVolumeGroupResponseBody(volumeGroupID)

	return body.Data.TaskID, err
}

func (s *Service) deleteVolumeGroupResponseBody(volumeGroupID string) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

	if volumeGroupID == "" {
		return body, fmt.Errorf("invalid volume group id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v2/volume-groups/%s", volumeGroupID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VolumeGroupNotFoundError{ID: volumeGroupID}
		}

		return nil
	})
}

// GetVolumeGroupTasks retrieves a list of VolumeGroup tasks
func (s *Service) GetVolumeGroupTasks(volumeGroupID string, parameters connection.APIRequestParameters) ([]Task, error) {
	var tasks []Task

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetVolumeGroupTasksPaginated(volumeGroupID, p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, task := range response.(*PaginatedTask).Items {
			tasks = append(tasks, task)
		}
	}

	return tasks, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetVolumeGroupTasksPaginated retrieves a paginated list of VolumeGroup tasks
func (s *Service) GetVolumeGroupTasksPaginated(volumeGroupID string, parameters connection.APIRequestParameters) (*PaginatedTask, error) {
	body, err := s.getVolumeGroupTasksPaginatedResponseBody(volumeGroupID, parameters)

	return NewPaginatedTask(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetVolumeGroupTasksPaginated(volumeGroupID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getVolumeGroupTasksPaginatedResponseBody(volumeGroupID string, parameters connection.APIRequestParameters) (*GetTaskSliceResponseBody, error) {
	body := &GetTaskSliceResponseBody{}

	if volumeGroupID == "" {
		return body, fmt.Errorf("invalid volume group id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/volume-groups/%s/tasks", volumeGroupID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VolumeGroupNotFoundError{ID: volumeGroupID}
		}

		return nil
	})
}
