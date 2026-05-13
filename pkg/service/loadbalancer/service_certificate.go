package loadbalancer

import (
	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) certificateRes() *resource.Resource[Certificate, int] {
	return resource.NewIntResource[Certificate](s.connection, "/loadbalancers/v2/certs", "certificate",
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
