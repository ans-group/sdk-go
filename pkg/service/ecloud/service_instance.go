package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetInstances retrieves a list of instances
func (s *Service) GetInstances(parameters connection.APIRequestParameters) ([]Instance, error) {
	var sites []Instance

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetInstancesPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, site := range response.(*PaginatedInstance).Items {
			sites = append(sites, site)
		}
	}

	return sites, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetInstancesPaginated retrieves a paginated list of instances
func (s *Service) GetInstancesPaginated(parameters connection.APIRequestParameters) (*PaginatedInstance, error) {
	body, err := s.getInstancesPaginatedResponseBody(parameters)

	return NewPaginatedInstance(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetInstancesPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getInstancesPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetInstanceSliceResponseBody, error) {
	body := &GetInstanceSliceResponseBody{}

	response, err := s.connection.Get("/ecloud/v2/instances", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetInstance retrieves a single instance by id
func (s *Service) GetInstance(instanceID string) (Instance, error) {
	body, err := s.getInstanceResponseBody(instanceID)

	return body.Data, err
}

func (s *Service) getInstanceResponseBody(instanceID string) (*GetInstanceResponseBody, error) {
	body := &GetInstanceResponseBody{}

	if instanceID == "" {
		return body, fmt.Errorf("invalid instance id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/instances/%s", instanceID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &InstanceNotFoundError{ID: instanceID}
		}

		return nil
	})
}

// CreateInstance creates a new instance
func (s *Service) CreateInstance(req CreateInstanceRequest) (string, error) {
	body, err := s.createInstanceResponseBody(req)

	return body.Data.ID, err
}

func (s *Service) createInstanceResponseBody(req CreateInstanceRequest) (*GetInstanceResponseBody, error) {
	body := &GetInstanceResponseBody{}

	response, err := s.connection.Post("/ecloud/v2/instances", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchInstance updates an instance
func (s *Service) PatchInstance(instanceID string, req PatchInstanceRequest) error {
	_, err := s.patchInstanceResponseBody(instanceID, req)

	return err
}

func (s *Service) patchInstanceResponseBody(instanceID string, req PatchInstanceRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if instanceID == "" {
		return body, fmt.Errorf("invalid instance id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ecloud/v2/instances/%s", instanceID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &InstanceNotFoundError{ID: instanceID}
		}

		return nil
	})
}

// DeleteInstance removes an instance
func (s *Service) DeleteInstance(instanceID string) error {
	_, err := s.deleteInstanceResponseBody(instanceID)

	return err
}

func (s *Service) deleteInstanceResponseBody(instanceID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if instanceID == "" {
		return body, fmt.Errorf("invalid instance id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v2/instances/%s", instanceID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &InstanceNotFoundError{ID: instanceID}
		}

		return nil
	})
}

// LockInstance locks an instance from update/removal
func (s *Service) LockInstance(instanceID string) error {
	_, err := s.lockInstanceResponseBody(instanceID)

	return err
}

func (s *Service) lockInstanceResponseBody(instanceID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if instanceID == "" {
		return body, fmt.Errorf("invalid instance id")
	}

	response, err := s.connection.Put(fmt.Sprintf("/ecloud/v2/instances/%s/lock", instanceID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &InstanceNotFoundError{ID: instanceID}
		}

		return nil
	})
}

// UnlockInstance unlocks an instance
func (s *Service) UnlockInstance(instanceID string) error {
	_, err := s.unlockInstanceResponseBody(instanceID)

	return err
}

func (s *Service) unlockInstanceResponseBody(instanceID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if instanceID == "" {
		return body, fmt.Errorf("invalid instance id")
	}

	response, err := s.connection.Put(fmt.Sprintf("/ecloud/v2/instances/%s/unlock", instanceID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &InstanceNotFoundError{ID: instanceID}
		}

		return nil
	})
}

// PowerOnInstance powers on an instance
func (s *Service) PowerOnInstance(instanceID string) error {
	_, err := s.powerOnInstanceResponseBody(instanceID)

	return err
}

func (s *Service) powerOnInstanceResponseBody(instanceID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if instanceID == "" {
		return body, fmt.Errorf("invalid instance id")
	}

	response, err := s.connection.Put(fmt.Sprintf("/ecloud/v2/instances/%s/power-on", instanceID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &InstanceNotFoundError{ID: instanceID}
		}

		return nil
	})
}

// PowerOffInstance powers off an instance
func (s *Service) PowerOffInstance(instanceID string) error {
	_, err := s.powerOffInstanceResponseBody(instanceID)

	return err
}

func (s *Service) powerOffInstanceResponseBody(instanceID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if instanceID == "" {
		return body, fmt.Errorf("invalid instance id")
	}

	response, err := s.connection.Put(fmt.Sprintf("/ecloud/v2/instances/%s/power-off", instanceID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &InstanceNotFoundError{ID: instanceID}
		}

		return nil
	})
}

// PowerResetInstance resets an instance
func (s *Service) PowerResetInstance(instanceID string) error {
	_, err := s.powerResetInstanceResponseBody(instanceID)

	return err
}

func (s *Service) powerResetInstanceResponseBody(instanceID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if instanceID == "" {
		return body, fmt.Errorf("invalid instance id")
	}

	response, err := s.connection.Put(fmt.Sprintf("/ecloud/v2/instances/%s/power-reset", instanceID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &InstanceNotFoundError{ID: instanceID}
		}

		return nil
	})
}

// PowerShutdownInstance shuts down an instance
func (s *Service) PowerShutdownInstance(instanceID string) error {
	_, err := s.powerShutdownInstanceResponseBody(instanceID)

	return err
}

func (s *Service) powerShutdownInstanceResponseBody(instanceID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if instanceID == "" {
		return body, fmt.Errorf("invalid instance id")
	}

	response, err := s.connection.Put(fmt.Sprintf("/ecloud/v2/instances/%s/power-shutdown", instanceID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &InstanceNotFoundError{ID: instanceID}
		}

		return nil
	})
}

// PowerRestartInstance restarts an instance
func (s *Service) PowerRestartInstance(instanceID string) error {
	_, err := s.powerRestartInstanceResponseBody(instanceID)

	return err
}

func (s *Service) powerRestartInstanceResponseBody(instanceID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if instanceID == "" {
		return body, fmt.Errorf("invalid instance id")
	}

	response, err := s.connection.Put(fmt.Sprintf("/ecloud/v2/instances/%s/power-restart", instanceID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &InstanceNotFoundError{ID: instanceID}
		}

		return nil
	})
}

// GetInstanceVolumes retrieves a list of instance volumes
func (s *Service) GetInstanceVolumes(instanceID string, parameters connection.APIRequestParameters) ([]Volume, error) {
	var volumes []Volume

	return volumes, connection.InvokeRequestAll(
		func(p connection.APIRequestParameters) (connection.Paginated, error) {
			return s.GetInstanceVolumesPaginated(instanceID, p)
		},
		func(response connection.Paginated) {
			for _, volume := range response.(*PaginatedVolume).Items {
				volumes = append(volumes, volume)
			}
		},
		parameters,
	)
}

// GetInstanceVolumesPaginated retrieves a paginated list of instance volumes
func (s *Service) GetInstanceVolumesPaginated(instanceID string, parameters connection.APIRequestParameters) (*PaginatedVolume, error) {
	body, err := s.getInstanceVolumesPaginatedResponseBody(instanceID, parameters)

	return NewPaginatedVolume(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetInstanceVolumesPaginated(instanceID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getInstanceVolumesPaginatedResponseBody(instanceID string, parameters connection.APIRequestParameters) (*GetVolumeSliceResponseBody, error) {
	body := &GetVolumeSliceResponseBody{}

	if instanceID == "" {
		return body, fmt.Errorf("invalid instance id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/instances/%s/volumes", instanceID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &InstanceNotFoundError{ID: instanceID}
		}

		return nil
	})
}
