package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetVPNSessions retrieves a list of VPN sessions
func (s *Service) GetVPNSessions(parameters connection.APIRequestParameters) ([]VPNSession, error) {
	return connection.InvokeRequestAll(s.GetVPNSessionsPaginated, parameters)
}

// GetVPNSessionsPaginated retrieves a paginated list of VPN sessions
func (s *Service) GetVPNSessionsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[VPNSession], error) {
	body, err := connection.Get[[]VPNSession](s.connection, "/ecloud/v2/vpn-sessions", parameters)
	return connection.NewPaginated(body, parameters, s.GetVPNSessionsPaginated), err
}

// GetVPNSession retrieves a single VPN session by id
func (s *Service) GetVPNSession(sessionID string) (VPNSession, error) {
	if sessionID == "" {
		return VPNSession{}, fmt.Errorf("invalid vpn session id")
	}
	body, err := connection.Get[VPNSession](s.connection, fmt.Sprintf("/ecloud/v2/vpn-sessions/%s", sessionID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&VPNSessionNotFoundError{ID: sessionID}))
	return body.Data, err
}

// CreateVPNSession creates a new VPN session
func (s *Service) CreateVPNSession(req CreateVPNSessionRequest) (TaskReference, error) {
	body, err := connection.Post[TaskReference](s.connection, "/ecloud/v2/vpn-sessions", &req)
	return body.Data, err
}

// PatchVPNSession patches a VPN session
func (s *Service) PatchVPNSession(sessionID string, req PatchVPNSessionRequest) (TaskReference, error) {
	if sessionID == "" {
		return TaskReference{}, fmt.Errorf("invalid vpn session id")
	}
	body, err := connection.Patch[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/vpn-sessions/%s", sessionID), &req, connection.NotFoundResponseHandler(&VPNSessionNotFoundError{ID: sessionID}))
	return body.Data, err
}

// DeleteVPNSession deletes a VPN session
func (s *Service) DeleteVPNSession(sessionID string) (string, error) {
	if sessionID == "" {
		return "", fmt.Errorf("invalid vpn session id")
	}
	body, err := connection.Delete[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/vpn-sessions/%s", sessionID), nil, connection.NotFoundResponseHandler(&VPNSessionNotFoundError{ID: sessionID}))
	return body.Data.TaskID, err
}

// GetVPNSessionTasks retrieves a list of VPN session tasks
func (s *Service) GetVPNSessionTasks(sessionID string, parameters connection.APIRequestParameters) ([]Task, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetVPNSessionTasksPaginated(sessionID, p)
	}, parameters)
}

// GetVPNSessionTasksPaginated retrieves a paginated list of VPN session tasks
func (s *Service) GetVPNSessionTasksPaginated(sessionID string, parameters connection.APIRequestParameters) (*connection.Paginated[Task], error) {
	if sessionID == "" {
		return nil, fmt.Errorf("invalid vpn session id")
	}
	body, err := connection.Get[[]Task](s.connection, fmt.Sprintf("/ecloud/v2/vpn-sessions/%s/tasks", sessionID), parameters, connection.NotFoundResponseHandler(&VPNSessionNotFoundError{ID: sessionID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetVPNSessionTasksPaginated(sessionID, p)
	}), err
}

// GetVPNSessionPreSharedKey retrieves a single VPN session by id
func (s *Service) GetVPNSessionPreSharedKey(sessionID string) (VPNSessionPreSharedKey, error) {
	if sessionID == "" {
		return VPNSessionPreSharedKey{}, fmt.Errorf("invalid vpn session id")
	}
	body, err := connection.Get[VPNSessionPreSharedKey](s.connection, fmt.Sprintf("/ecloud/v2/vpn-sessions/%s/pre-shared-key", sessionID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&VPNSessionNotFoundError{ID: sessionID}))
	return body.Data, err
}

// UpdateVPNSessionPreSharedKey retrieves a single VPN session by id
func (s *Service) UpdateVPNSessionPreSharedKey(sessionID string, req UpdateVPNSessionPreSharedKeyRequest) (TaskReference, error) {
	if sessionID == "" {
		return TaskReference{}, fmt.Errorf("invalid vpn session id")
	}
	body, err := connection.Put[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/vpn-sessions/%s/pre-shared-key", sessionID), &req, connection.NotFoundResponseHandler(&VPNSessionNotFoundError{ID: sessionID}))
	return body.Data, err
}
