package loadbalancer

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetConfigurations retrieves a list of configurations
func (s *Service) GetConfigurations(parameters connection.APIRequestParameters) ([]Configuration, error) {
	var sites []Configuration

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetConfigurationsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, site := range response.(*PaginatedConfiguration).Items {
			sites = append(sites, site)
		}
	}

	return sites, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetConfigurationsPaginated retrieves a paginated list of configurations
func (s *Service) GetConfigurationsPaginated(parameters connection.APIRequestParameters) (*PaginatedConfiguration, error) {
	body, err := s.getConfigurationsPaginatedResponseBody(parameters)

	return NewPaginatedConfiguration(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetConfigurationsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getConfigurationsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetConfigurationSliceResponseBody, error) {
	body := &GetConfigurationSliceResponseBody{}

	response, err := s.connection.Get("/loadbalancers/v2/configurations", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetConfiguration retrieves a single configuration by id
func (s *Service) GetConfiguration(configurationID string) (Configuration, error) {
	body, err := s.getConfigurationResponseBody(configurationID)

	return body.Data, err
}

func (s *Service) getConfigurationResponseBody(configurationID string) (*GetConfigurationResponseBody, error) {
	body := &GetConfigurationResponseBody{}

	if configurationID == "" {
		return body, fmt.Errorf("invalid configuration id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/loadbalancers/v2/configurations/%s", configurationID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ConfigurationNotFoundError{ID: configurationID}
		}

		return nil
	})
}

// CreateConfiguration creates a new Configuration
func (s *Service) CreateConfiguration(req CreateConfigurationRequest) (string, error) {
	body, err := s.createConfigurationResponseBody(req)

	return body.Data.ID, err
}

func (s *Service) createConfigurationResponseBody(req CreateConfigurationRequest) (*GetConfigurationResponseBody, error) {
	body := &GetConfigurationResponseBody{}

	response, err := s.connection.Post("/loadbalancers/v2/configurations", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchConfiguration patches a Configuration
func (s *Service) PatchConfiguration(configurationID string, req PatchConfigurationRequest) error {
	_, err := s.patchConfigurationResponseBody(configurationID, req)

	return err
}

func (s *Service) patchConfigurationResponseBody(configurationID string, req PatchConfigurationRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if configurationID == "" {
		return body, fmt.Errorf("invalid configuration id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/loadbalancers/v2/configurations/%s", configurationID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ConfigurationNotFoundError{ID: configurationID}
		}

		return nil
	})
}

// DeleteConfiguration deletes a Configuration
func (s *Service) DeleteConfiguration(configurationID string) error {
	_, err := s.deleteConfigurationResponseBody(configurationID)

	return err
}

func (s *Service) deleteConfigurationResponseBody(configurationID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if configurationID == "" {
		return body, fmt.Errorf("invalid configuration id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/loadbalancers/v2/configurations/%s", configurationID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ConfigurationNotFoundError{ID: configurationID}
		}

		return nil
	})
}
