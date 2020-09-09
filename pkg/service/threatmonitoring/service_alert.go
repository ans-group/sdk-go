package threatmonitoring

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetAlerts retrieves a list of alerts
func (s *Service) GetAlerts(parameters connection.APIRequestParameters) ([]Alert, error) {
	var alerts []Alert

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetAlertsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, alert := range response.(*PaginatedAlert).Items {
			alerts = append(alerts, alert)
		}
	}

	return alerts, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetAlertsPaginated retrieves a paginated list of alerts
func (s *Service) GetAlertsPaginated(parameters connection.APIRequestParameters) (*PaginatedAlert, error) {
	body, err := s.getAlertsPaginatedResponseBody(parameters)

	return NewPaginatedAlert(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetAlertsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getAlertsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetAlertSliceResponseBody, error) {
	body := &GetAlertSliceResponseBody{}

	response, err := s.connection.Get("/threat-monitoring/v1/alerts", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetAlert retrieves a single alert by id
func (s *Service) GetAlert(alertID string) (Alert, error) {
	body, err := s.getAlertResponseBody(alertID)

	return body.Data, err
}

func (s *Service) getAlertResponseBody(alertID string) (*GetAlertResponseBody, error) {
	body := &GetAlertResponseBody{}

	if alertID == "" {
		return body, fmt.Errorf("invalid alert id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/threat-monitoring/v1/alerts/%s", alertID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &AlertNotFoundError{ID: alertID}
		}

		return nil
	})
}
