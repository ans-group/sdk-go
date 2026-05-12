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
	body, err := connection.Get[[]Application](s.connection, "/account/v1/applications", parameters)
	return connection.NewPaginated(body, parameters, s.GetApplicationsPaginated), err
}

// GetApplication retrieves a single application by id
func (s *Service) GetApplication(appID string) (Application, error) {
	if appID == "" {
		return Application{}, fmt.Errorf("invalid application id")
	}
	body, err := connection.Get[Application](s.connection, fmt.Sprintf("/account/v1/applications/%s", appID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&ApplicationNotFoundError{ID: appID}))
	return body.Data, err
}

// GetApplications retrieves a list of applications
func (s *Service) GetServices(parameters connection.APIRequestParameters) ([]ApplicationService, error) {
	return connection.InvokeRequestAll(s.GetServicesPaginated, parameters)
}

func (s *Service) GetServicesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[ApplicationService], error) {
	body, err := connection.Get[[]ApplicationService](s.connection, "/account/v1/services", parameters)
	return connection.NewPaginated(body, parameters, s.GetServicesPaginated), err
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
