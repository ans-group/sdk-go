package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetVolumes retrieves a list of volumes
func (s *Service) GetVolumes(parameters connection.APIRequestParameters) ([]Volume, error) {
	var volumes []Volume

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetVolumesPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, volume := range response.(*PaginatedVolume).Items {
			volumes = append(volumes, volume)
		}
	}

	return volumes, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetVolumesPaginated retrieves a paginated list of volumes
func (s *Service) GetVolumesPaginated(parameters connection.APIRequestParameters) (*PaginatedVolume, error) {
	body, err := s.getVolumesPaginatedResponseBody(parameters)

	return NewPaginatedVolume(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetVolumesPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getVolumesPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetVolumeSliceResponseBody, error) {
	body := &GetVolumeSliceResponseBody{}

	response, err := s.connection.Get("/ecloud/v2/volumes", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetVolume retrieves a single volume by id
func (s *Service) GetVolume(volumeID string) (Volume, error) {
	body, err := s.getVolumeResponseBody(volumeID)

	return body.Data, err
}

func (s *Service) getVolumeResponseBody(volumeID string) (*GetVolumeResponseBody, error) {
	body := &GetVolumeResponseBody{}

	if volumeID == "" {
		return body, fmt.Errorf("invalid volume id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/volumes/%s", volumeID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VolumeNotFoundError{ID: volumeID}
		}

		return nil
	})
}

// CreateVolume creates a volume
func (s *Service) CreateVolume(req CreateVolumeRequest) (string, error) {
	body, err := s.createVolumeResponseBody(req)

	return body.Data.ID, err
}

func (s *Service) createVolumeResponseBody(req CreateVolumeRequest) (*GetVolumeResponseBody, error) {
	body := &GetVolumeResponseBody{}

	response, err := s.connection.Post("/ecloud/v2/volumes", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchVolume patches a Volume
func (s *Service) PatchVolume(volumeID string, req PatchVolumeRequest) error {
	_, err := s.patchVolumeResponseBody(volumeID, req)

	return err
}

func (s *Service) patchVolumeResponseBody(volumeID string, req PatchVolumeRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if volumeID == "" {
		return body, fmt.Errorf("invalid volume id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ecloud/v2/volumes/%s", volumeID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VolumeNotFoundError{ID: volumeID}
		}

		return nil
	})
}

// DeleteVolume deletes a Volume
func (s *Service) DeleteVolume(volumeID string) error {
	_, err := s.deleteVolumeResponseBody(volumeID)

	return err
}

func (s *Service) deleteVolumeResponseBody(volumeID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if volumeID == "" {
		return body, fmt.Errorf("invalid volume id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v2/volumes/%s", volumeID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VolumeNotFoundError{ID: volumeID}
		}

		return nil
	})
}

// GetVolumeInstances retrieves a list of volume instances
func (s *Service) GetVolumeInstances(volumeID string, parameters connection.APIRequestParameters) ([]Instance, error) {
	var instances []Instance

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetVolumeInstancesPaginated(volumeID, p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, instance := range response.(*PaginatedInstance).Items {
			instances = append(instances, instance)
		}
	}

	return instances, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetVolumeInstancesPaginated retrieves a paginated list of volume instances
func (s *Service) GetVolumeInstancesPaginated(volumeID string, parameters connection.APIRequestParameters) (*PaginatedInstance, error) {
	body, err := s.getVolumeInstancesPaginatedResponseBody(volumeID, parameters)

	return NewPaginatedInstance(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetVolumeInstancesPaginated(volumeID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getVolumeInstancesPaginatedResponseBody(volumeID string, parameters connection.APIRequestParameters) (*GetInstanceSliceResponseBody, error) {
	body := &GetInstanceSliceResponseBody{}

	if volumeID == "" {
		return body, fmt.Errorf("invalid volume id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/volumes/%s/instances", volumeID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VolumeNotFoundError{ID: volumeID}
		}

		return nil
	})
}

func (s *Service) AttachVolume(volumeID string, req AttachVolumeRequest) (string, error) {
	body, err := s.attachVolumeResponseBody(volumeID, req)

	return body.Data.TaskID, err
}

func (s *Service) attachVolumeResponseBody(volumeID string, req AttachVolumeRequest) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

	if volumeID == "" {
		return body, fmt.Errorf("invalid volume id")
	}

	response, err := s.connection.Post(fmt.Sprintf("/ecloud/v2/volumes/%s/attach", volumeID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VolumeNotFoundError{ID: volumeID}
		}

		return nil
	})
}

func (s *Service) DetachVolume(volumeID string, req DetachVolumeRequest) (string, error) {
	body, err := s.detachVolumeResponseBody(volumeID, req)

	return body.Data.TaskID, err
}

func (s *Service) detachVolumeResponseBody(volumeID string, req DetachVolumeRequest) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

	if volumeID == "" {
		return body, fmt.Errorf("invalid volume id")
	}

	response, err := s.connection.Post(fmt.Sprintf("/ecloud/v2/volumes/%s/detach", volumeID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VolumeNotFoundError{ID: volumeID}
		}

		return nil
	})
}

// GetVolumeTasks retrieves a list of Volume tasks
func (s *Service) GetVolumeTasks(volumeID string, parameters connection.APIRequestParameters) ([]Task, error) {
	var tasks []Task

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetVolumeTasksPaginated(volumeID, p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, task := range response.(*PaginatedTask).Items {
			tasks = append(tasks, task)
		}
	}

	return tasks, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetVolumeTasksPaginated retrieves a paginated list of Volume tasks
func (s *Service) GetVolumeTasksPaginated(volumeID string, parameters connection.APIRequestParameters) (*PaginatedTask, error) {
	body, err := s.getVolumeTasksPaginatedResponseBody(volumeID, parameters)

	return NewPaginatedTask(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetVolumeTasksPaginated(volumeID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getVolumeTasksPaginatedResponseBody(volumeID string, parameters connection.APIRequestParameters) (*GetTaskSliceResponseBody, error) {
	body := &GetTaskSliceResponseBody{}

	if volumeID == "" {
		return body, fmt.Errorf("invalid volume id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/volumes/%s/tasks", volumeID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VolumeNotFoundError{ID: volumeID}
		}

		return nil
	})
}
