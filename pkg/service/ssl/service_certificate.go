package ssl

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) certificateRes() *resource.Resource[Certificate, int] {
	return resource.NewIntResource[Certificate](s.connection, "/ssl/v1/certificates", "certificate",
		func(id int) error { return &CertificateNotFoundError{ID: id} })
}

// GetCertificates retrieves a list of certificates
func (s *Service) GetCertificates(parameters connection.APIRequestParameters) ([]Certificate, error) {
	return s.certificateRes().List(parameters)
}

// GetCertificatesPaginated retrieves a paginated list of certificates
func (s *Service) GetCertificatesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Certificate], error) {
	return s.certificateRes().ListPaginated(parameters)
}

// GetCertificate retrieves a single certificate by id
func (s *Service) GetCertificate(certificateID int) (Certificate, error) {
	return s.certificateRes().Get(certificateID)
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
