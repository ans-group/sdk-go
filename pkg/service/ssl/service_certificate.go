package ssl

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetCertificates retrieves a list of certificates
func (s *Service) GetCertificates(parameters connection.APIRequestParameters) ([]Certificate, error) {
	return connection.InvokeRequestAll(s.GetCertificatesPaginated, parameters)
}

// GetCertificatesPaginated retrieves a paginated list of certificates
func (s *Service) GetCertificatesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Certificate], error) {
	body, err := connection.Get[[]Certificate](s.connection, "/ssl/v1/certificates", parameters)
	return connection.NewPaginated(body, parameters, s.GetCertificatesPaginated), err
}

// GetCertificate retrieves a single certificate by id
func (s *Service) GetCertificate(certificateID int) (Certificate, error) {
	if certificateID < 1 {
		return Certificate{}, fmt.Errorf("invalid certificate id")
	}
	body, err := connection.Get[Certificate](s.connection, fmt.Sprintf("/ssl/v1/certificates/%d", certificateID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&CertificateNotFoundError{ID: certificateID}))
	return body.Data, err
}

// GetCertificateContent retrieves the content of an SSL certificate
func (s *Service) GetCertificateContent(certificateID int) (CertificateContent, error) {
	if certificateID < 1 {
		return CertificateContent{}, fmt.Errorf("invalid certificate id")
	}
	body, err := connection.Get[CertificateContent](s.connection, fmt.Sprintf("/ssl/v1/certificates/%d/download", certificateID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&CertificateNotFoundError{ID: certificateID}))
	return body.Data, err
}

// GetCertificatePrivateKey retrieves an SSL certificate private key
func (s *Service) GetCertificatePrivateKey(certificateID int) (CertificatePrivateKey, error) {
	if certificateID < 1 {
		return CertificatePrivateKey{}, fmt.Errorf("invalid certificate id")
	}
	body, err := connection.Get[CertificatePrivateKey](s.connection, fmt.Sprintf("/ssl/v1/certificates/%d/private-key", certificateID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&CertificateNotFoundError{ID: certificateID}))
	return body.Data, err
}
