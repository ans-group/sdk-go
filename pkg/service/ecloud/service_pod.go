package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetPods retrieves a list of pods
func (s *Service) GetPods(parameters connection.APIRequestParameters) ([]Pod, error) {
	return connection.InvokeRequestAll(s.GetPodsPaginated, parameters)
}

// GetPodsPaginated retrieves a paginated list of pods
func (s *Service) GetPodsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Pod], error) {
	body, err := connection.Get[[]Pod](s.connection, "/ecloud/v1/pods", parameters)
	return connection.NewPaginated(body, parameters, s.GetPodsPaginated), err
}

// GetPod retrieves a single pod by ID
func (s *Service) GetPod(podID int) (Pod, error) {
	if podID < 1 {
		return Pod{}, fmt.Errorf("invalid pod id")
	}
	body, err := connection.Get[Pod](s.connection, fmt.Sprintf("/ecloud/v1/pods/%d", podID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&PodNotFoundError{ID: podID}))
	return body.Data, err
}

// GetPodTemplates retrieves a list of templates
func (s *Service) GetPodTemplates(podID int, parameters connection.APIRequestParameters) ([]Template, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Template], error) {
		return s.GetPodTemplatesPaginated(podID, p)
	}, parameters)
}

// GetPodTemplatesPaginated retrieves a paginated list of domains
func (s *Service) GetPodTemplatesPaginated(podID int, parameters connection.APIRequestParameters) (*connection.Paginated[Template], error) {
	if podID < 1 {
		return nil, fmt.Errorf("invalid pod id")
	}
	body, err := connection.Get[[]Template](s.connection, fmt.Sprintf("/ecloud/v1/pods/%d/templates", podID), parameters)
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Template], error) {
		return s.GetPodTemplatesPaginated(podID, p)
	}), err
}

// GetPodTemplate retrieves a single pod template by name
func (s *Service) GetPodTemplate(podID int, templateName string) (Template, error) {
	if podID < 1 {
		return Template{}, fmt.Errorf("invalid pod id")
	}
	if templateName == "" {
		return Template{}, fmt.Errorf("invalid template name")
	}
	body, err := connection.Get[Template](s.connection, fmt.Sprintf("/ecloud/v1/pods/%d/templates/%s", podID, templateName), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&TemplateNotFoundError{Name: templateName}))
	return body.Data, err
}

// RenamePodTemplate renames a pod template
func (s *Service) RenamePodTemplate(podID int, templateName string, req RenameTemplateRequest) error {
	if podID < 1 {
		return fmt.Errorf("invalid pod id")
	}
	if templateName == "" {
		return fmt.Errorf("invalid template name")
	}
	return connection.PostRaw(s.connection, fmt.Sprintf("/ecloud/v1/pods/%d/templates/%s/move", podID, templateName), &req, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&TemplateNotFoundError{Name: templateName}))
}

// DeletePodTemplate removes a pod template
func (s *Service) DeletePodTemplate(podID int, templateName string) error {
	if podID < 1 {
		return fmt.Errorf("invalid pod id")
	}
	if templateName == "" {
		return fmt.Errorf("invalid template name")
	}
	return connection.DeleteRaw(s.connection, fmt.Sprintf("/ecloud/v1/pods/%d/templates/%s", podID, templateName), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&TemplateNotFoundError{Name: templateName}))
}

// GetPodAppliances retrieves a list of appliances
func (s *Service) GetPodAppliances(podID int, parameters connection.APIRequestParameters) ([]Appliance, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Appliance], error) {
		return s.GetPodAppliancesPaginated(podID, p)
	}, parameters)
}

// GetPodAppliancesPaginated retrieves a paginated list of domains
func (s *Service) GetPodAppliancesPaginated(podID int, parameters connection.APIRequestParameters) (*connection.Paginated[Appliance], error) {
	if podID < 1 {
		return nil, fmt.Errorf("invalid pod id")
	}
	body, err := connection.Get[[]Appliance](s.connection, fmt.Sprintf("/ecloud/v1/pods/%d/appliances", podID), parameters)
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Appliance], error) {
		return s.GetPodAppliancesPaginated(podID, p)
	}), err
}

// PodConsoleAvailable removes a pod template
func (s *Service) PodConsoleAvailable(podID int) (bool, error) {
	if podID < 1 {
		return false, fmt.Errorf("invalid pod id")
	}
	resp, err := s.connection.Get(fmt.Sprintf("/ecloud/v1/pods/%d/console-available", podID), connection.APIRequestParameters{})
	if err != nil {
		return false, err
	}
	if resp.StatusCode == 200 {
		return true, nil
	}
	return false, nil
}
