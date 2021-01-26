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
