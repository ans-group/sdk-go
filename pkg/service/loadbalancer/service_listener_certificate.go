package loadbalancer

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) listenerCertRes() *resource.SubResource[Certificate, int, int] {
	return resource.NewIntIntSubResource[Certificate](s.connection,
		func(listenerID int) string { return fmt.Sprintf("/loadbalancers/v2/listeners/%d/certs", listenerID) },
		"listener", "id", func(listenerID int) error { return &CertificateNotFoundError{ID: listenerID} },
		"certificate", "id", func(listenerID, _ int) error { return &CertificateNotFoundError{ID: listenerID} })
}

// GetListenerCertificates retrieves a list of certificates
func (s *Service) GetListenerCertificates(listenerID int, parameters connection.APIRequestParameters) ([]Certificate, error) {
	return s.listenerCertRes().List(listenerID, parameters)
}

// GetListenerCertificatesPaginated retrieves a paginated list of certificates
func (s *Service) GetListenerCertificatesPaginated(listenerID int, parameters connection.APIRequestParameters) (*connection.Paginated[Certificate], error) {
	return s.listenerCertRes().ListPaginated(listenerID, parameters)
}

// GetListenerCertificate retrieves a single certificate by id
func (s *Service) GetListenerCertificate(listenerID int, certificateID int) (Certificate, error) {
	return s.listenerCertRes().Get(listenerID, certificateID)
}

// CreateListenerCertificate creates an certificate
func (s *Service) CreateListenerCertificate(listenerID int, req CreateCertificateRequest) (int, error) {
	cert, err := s.listenerCertRes().Create(listenerID, &req)
	return cert.ID, err
}

// PatchListenerCertificate patches an certificate
func (s *Service) PatchListenerCertificate(listenerID int, certificateID int, req PatchCertificateRequest) error {
	return s.listenerCertRes().Patch(listenerID, certificateID, &req)
}

// DeleteListenerCertificate deletes a certificate
func (s *Service) DeleteListenerCertificate(listenerID int, certificateID int) error {
	return s.listenerCertRes().Delete(listenerID, certificateID)
}
