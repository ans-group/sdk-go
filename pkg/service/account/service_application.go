package account

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) applicationRes() *resource.Resource[Application, string] {
	return resource.NewStringResource[Application](s.connection, "/account/v1/applications", "application",
		func(id string) error { return &ApplicationNotFoundError{ID: id} })
}

func (s *Service) serviceRes() *resource.Resource[ApplicationService, string] {
	return resource.NewStringResource[ApplicationService](s.connection, "/account/v1/services", "service",
		func(id string) error { return fmt.Errorf("service not found: %s", id) })
}

// GetApplications retrieves a list of applications
func (s *Service) GetApplications(parameters connection.APIRequestParameters) ([]Application, error) {
	return s.applicationRes().List(parameters)
}

func (s *Service) GetApplicationsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Application], error) {
	return s.applicationRes().ListPaginated(parameters)
}

// GetApplication retrieves a single application by id
func (s *Service) GetApplication(appID string) (Application, error) {
	return s.applicationRes().Get(appID)
}

// GetApplications retrieves a list of applications
func (s *Service) GetServices(parameters connection.APIRequestParameters) ([]ApplicationService, error) {
	return s.serviceRes().List(parameters)
}

func (s *Service) GetServicesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[ApplicationService], error) {
	return s.serviceRes().ListPaginated(parameters)
}

func (s *Service) CreateApplication(req CreateApplicationRequest) (CreateApplicationResponse, error) {
	body, err := connection.Post[CreateApplicationResponse](s.connection, "/account/v1/applications", &req)
	return body.Data, err
}

func (s *Service) UpdateApplication(appID string, req UpdateApplicationRequest) error {
	if appID == "" {
		return fmt.Errorf("invalid application id")
	}
	_, err := connection.Patch[Application](s.connection, fmt.Sprintf("/account/v1/applications/%s", appID), &req, connection.NotFoundResponseHandler(&ApplicationNotFoundError{ID: appID}))
	return err
}

// GetApplicationServices retrieves the services and roles of an application by id
func (s *Service) GetApplicationServices(appID string) (ApplicationServiceMapping, error) {
	if appID == "" {
		return ApplicationServiceMapping{}, fmt.Errorf("invalid application id")
	}
	body, err := connection.Get[ApplicationServiceMapping](s.connection, fmt.Sprintf("/account/v1/applications/%s/services", appID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&ApplicationNotFoundError{ID: appID}))
	return body.Data, err
}

func (s *Service) SetApplicationServices(appID string, req SetServiceRequest) error {
	if appID == "" {
		return fmt.Errorf("invalid application id")
	}
	_, err := connection.Put[interface{}](s.connection, fmt.Sprintf("/account/v1/applications/%s/services", appID), &req, connection.NotFoundResponseHandler(&ApplicationNotFoundError{ID: appID}))
	return err
}

func (s *Service) DeleteApplicationServices(appID string) error {
	if appID == "" {
		return fmt.Errorf("invalid application id")
	}
	_, err := connection.Put[interface{}](s.connection, fmt.Sprintf("/account/v1/applications/%s/services", appID), SetServiceRequest{Scopes: []ApplicationServiceScope{}}, connection.NotFoundResponseHandler(&ApplicationNotFoundError{ID: appID}))
	return err
}

// DeleteApplication removes an application
func (s *Service) DeleteApplication(appID string) error {
	if appID == "" {
		return fmt.Errorf("invalid application id")
	}
	_, err := connection.Delete[interface{}](s.connection, fmt.Sprintf("/account/v1/applications/%s", appID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&ApplicationNotFoundError{ID: appID}))
	return err
}

// GetApplicationRestrictions retrieves the IP restrictions of an application by id
func (s *Service) GetApplicationRestrictions(appID string) (ApplicationRestriction, error) {
	if appID == "" {
		return ApplicationRestriction{}, fmt.Errorf("invalid application id")
	}
	body, err := connection.Get[ApplicationRestriction](s.connection, fmt.Sprintf("/account/v1/applications/%s/ip-restrictions", appID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&ApplicationNotFoundError{ID: appID}))
	return body.Data, err
}

func (s *Service) SetApplicationRestrictions(appID string, req SetRestrictionRequest) error {
	if appID == "" {
		return fmt.Errorf("invalid application id")
	}
	_, err := connection.Put[interface{}](s.connection, fmt.Sprintf("/account/v1/applications/%s/ip-restrictions", appID), &req, connection.NotFoundResponseHandler(&ApplicationNotFoundError{ID: appID}))
	return err
}

func (s *Service) DeleteApplicationRestrictions(appID string) error {
	if appID == "" {
		return fmt.Errorf("invalid application id")
	}
	_, err := connection.Put[interface{}](s.connection, fmt.Sprintf("/account/v1/applications/%s/ip-restrictions", appID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&ApplicationNotFoundError{ID: appID}))
	return err
}
