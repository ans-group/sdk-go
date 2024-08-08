package account

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetApplications retrieves a list of applications
func (s *Service) GetApplications(parameters connection.APIRequestParameters) ([]Application, error) {
	return connection.InvokeRequestAll(s.GetApplicationsPaginated, parameters)
}

func (s *Service) GetApplicationsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Application], error) {
	body, err := s.getApplicationsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetApplicationsPaginated), err
}

func (s *Service) getApplicationsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Application], error) {
	body := &connection.APIResponseBodyData[[]Application]{}

	response, err := s.connection.Get("/account/v1/applications", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetApplication retrieves a single application by id
func (s *Service) GetApplication(appID string) (Application, error) {
	body, err := s.getApplicationResponseBody(appID)

	return body.Data, err
}

func (s *Service) getApplicationResponseBody(appID string) (*connection.APIResponseBodyData[Application], error) {
	body := &connection.APIResponseBodyData[Application]{}

	if appID == "" {
		return body, fmt.Errorf("invalid application id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/account/v1/applications/%s", appID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ApplicationNotFoundError{ID: appID}
		}

		return nil
	})
}

// GetApplications retrieves a list of applications
func (s *Service) GetServices(parameters connection.APIRequestParameters) ([]ApplicationService, error) {
	return connection.InvokeRequestAll(s.GetServicesPaginated, parameters)
}

func (s *Service) GetServicesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[ApplicationService], error) {
	body, err := s.getServicesPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetServicesPaginated), err
}

func (s *Service) getServicesPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]ApplicationService], error) {
	body := &connection.APIResponseBodyData[[]ApplicationService]{}

	response, err := s.connection.Get("/account/v1/services", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

func (s *Service) CreateApplication(req CreateApplicationRequest) (string, error) {
	body, err := s.createApplicationResponseBody(req)

	return body.Data.ID, err
}

func (s *Service) createApplicationResponseBody(req CreateApplicationRequest) (*connection.APIResponseBodyData[Application], error) {
	body := &connection.APIResponseBodyData[Application]{}

	response, err := s.connection.Post("/account/v1/applications", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetApplicationServices retrieves the services and roles of an application by id
func (s *Service) GetApplicationServices(appID string) (ApplicationServiceMapping, error) {
	body, err := s.getApplicationServicesResponseBody(appID)

	return body.Data, err
}

func (s *Service) getApplicationServicesResponseBody(appID string) (*connection.APIResponseBodyData[ApplicationServiceMapping], error) {
	body := &connection.APIResponseBodyData[ApplicationServiceMapping]{}

	if appID == "" {
		return body, fmt.Errorf("invalid application id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/account/v1/applications/%s/services", appID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ApplicationNotFoundError{ID: appID}
		}

		return nil
	})
}

func (s *Service) SetApplicationServices(appID string, req SetServiceRequest) error {
	_, err := s.setApplicationServicesResponseBody(appID, req)

	return err
}

func (s *Service) setApplicationServicesResponseBody(appID string, req SetServiceRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if appID == "" {
		return body, fmt.Errorf("invalid application id")
	}

	response, err := s.connection.Put(fmt.Sprintf("/account/v1/applications/%s/services", appID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ApplicationNotFoundError{ID: appID}
		}

		return nil
	})
}

// DeleteApplication removes an application
func (s *Service) DeleteApplication(appID string) error {
	_, err := s.deleteApplicationResponseBody(appID)

	return err
}

func (s *Service) deleteApplicationResponseBody(appID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if appID == "" {
		return body, fmt.Errorf("invalid application id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/account/v1/applications/%s", appID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ApplicationNotFoundError{ID: appID}
		}

		return nil
	})
}

// GetApplicationRestrictions retrieves the IP restrictions of an application by id
func (s *Service) GetApplicationRestrictions(appID string) (ApplicationRestriction, error) {
	body, err := s.getApplicationRestrictionsResponseBody(appID)

	return body.Data, err
}

func (s *Service) getApplicationRestrictionsResponseBody(appID string) (*connection.APIResponseBodyData[ApplicationRestriction], error) {
	body := &connection.APIResponseBodyData[ApplicationRestriction]{}

	if appID == "" {
		return body, fmt.Errorf("invalid application id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/account/v1/applications/%s/ip-restrictions", appID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ApplicationNotFoundError{ID: appID}
		}

		return nil
	})
}

func (s *Service) SetApplicationRestrictions(appID string, req SetRestrictionRequest) error {
	_, err := s.setApplicationRestrictionsResponseBody(appID, req)

	return err
}

func (s *Service) setApplicationRestrictionsResponseBody(appID string, req SetRestrictionRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if appID == "" {
		return body, fmt.Errorf("invalid application id")
	}

	response, err := s.connection.Put(fmt.Sprintf("/account/v1/applications/%s/ip-restrictions", appID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ApplicationNotFoundError{ID: appID}
		}

		return nil
	})
}
