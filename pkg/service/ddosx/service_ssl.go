package ddosx

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) sslRes() *resource.Resource[SSL, string] {
	return resource.NewStringResource[SSL](s.connection, "/ddosx/v1/ssls", "ssl",
		func(id string) error { return &SSLNotFoundError{ID: id} })
}

// GetSSLs retrieves a list of ssls
func (s *Service) GetSSLs(parameters connection.APIRequestParameters) ([]SSL, error) {
	return s.sslRes().List(parameters)
}

// GetSSLsPaginated retrieves a paginated list of ssls
func (s *Service) GetSSLsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[SSL], error) {
	return s.sslRes().ListPaginated(parameters)
}

// GetSSL retrieves a single ssl by id
func (s *Service) GetSSL(sslID string) (SSL, error) {
	return s.sslRes().Get(sslID)
}

// CreateSSL retrieves creates an SSL
func (s *Service) CreateSSL(req CreateSSLRequest) (string, error) {
	data, err := s.sslRes().Create(&req)
	return data.ID, err
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
	return s.sslRes().Delete(sslID)
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
