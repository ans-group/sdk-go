package ddosx

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetSSLs retrieves a list of ssls
func (s *Service) GetSSLs(parameters connection.APIRequestParameters) ([]SSL, error) {
	return connection.InvokeRequestAll(s.GetSSLsPaginated, parameters)
}

// GetSSLsPaginated retrieves a paginated list of ssls
func (s *Service) GetSSLsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[SSL], error) {
	body, err := connection.Get[[]SSL](s.connection, "/ddosx/v1/ssls", parameters)
	return connection.NewPaginated(body, parameters, s.GetSSLsPaginated), err
}

// GetSSL retrieves a single ssl by id
func (s *Service) GetSSL(sslID string) (SSL, error) {
	if sslID == "" {
		return SSL{}, fmt.Errorf("invalid ssl id")
	}
	body, err := connection.Get[SSL](s.connection, fmt.Sprintf("/ddosx/v1/ssls/%s", sslID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&SSLNotFoundError{ID: sslID}))
	return body.Data, err
}

// CreateSSL retrieves creates an SSL
func (s *Service) CreateSSL(req CreateSSLRequest) (string, error) {
	body, err := connection.Post[SSL](s.connection, "/ddosx/v1/ssls", &req)
	return body.Data.ID, err
}

// PatchSSL retrieves patches an SSL
func (s *Service) PatchSSL(sslID string, req PatchSSLRequest) (string, error) {
	if sslID == "" {
		return "", fmt.Errorf("invalid ssl id")
	}
	body, err := connection.Patch[SSL](s.connection, fmt.Sprintf("/ddosx/v1/ssls/%s", sslID), &req, connection.NotFoundResponseHandler(&SSLNotFoundError{ID: sslID}))
	return body.Data.ID, err
}

// DeleteSSL deletes patches an SSL
func (s *Service) DeleteSSL(sslID string) error {
	if sslID == "" {
		return fmt.Errorf("invalid ssl id")
	}
	return connection.DeleteRaw(s.connection, fmt.Sprintf("/ddosx/v1/ssls/%s", sslID), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&SSLNotFoundError{ID: sslID}))
}

// GetSSLContent retrieves a single ssl by id
func (s *Service) GetSSLContent(sslID string) (SSLContent, error) {
	if sslID == "" {
		return SSLContent{}, fmt.Errorf("invalid ssl id")
	}
	body, err := connection.Get[SSLContent](s.connection, fmt.Sprintf("/ddosx/v1/ssls/%s/certificates", sslID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&SSLNotFoundError{ID: sslID}))
	return body.Data, err
}

// GetSSLPrivateKey retrieves a single ssl by id
func (s *Service) GetSSLPrivateKey(sslID string) (SSLPrivateKey, error) {
	if sslID == "" {
		return SSLPrivateKey{}, fmt.Errorf("invalid ssl id")
	}
	body, err := connection.Get[SSLPrivateKey](s.connection, fmt.Sprintf("/ddosx/v1/ssls/%s/private-key", sslID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&SSLNotFoundError{ID: sslID}))
	return body.Data, err
}
