package loadbalancer

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetListenerCertificates retrieves a list of certificates
func (s *Service) GetListenerCertificates(listenerID int, parameters connection.APIRequestParameters) ([]Certificate, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Certificate], error) {
		return s.GetListenerCertificatesPaginated(listenerID, p)
	}, parameters)
}

// GetListenerCertificatesPaginated retrieves a paginated list of certificates
func (s *Service) GetListenerCertificatesPaginated(listenerID int, parameters connection.APIRequestParameters) (*connection.Paginated[Certificate], error) {
	if listenerID < 1 {
		return nil, fmt.Errorf("invalid listener id")
	}
	body, err := connection.Get[[]Certificate](s.connection, fmt.Sprintf("/loadbalancers/v2/listeners/%d/certs", listenerID), parameters)
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Certificate], error) {
		return s.GetListenerCertificatesPaginated(listenerID, p)
	}), err
}

// GetListenerCertificate retrieves a single certificate by id
func (s *Service) GetListenerCertificate(listenerID int, certificateID int) (Certificate, error) {
	if listenerID < 1 {
		return Certificate{}, fmt.Errorf("invalid listener id")
	}
	if certificateID < 1 {
		return Certificate{}, fmt.Errorf("invalid certificate id")
	}
	body, err := connection.Get[Certificate](s.connection, fmt.Sprintf("/loadbalancers/v2/listeners/%d/certs/%d", listenerID, certificateID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&CertificateNotFoundError{ID: listenerID}))
	return body.Data, err
}

// CreateListenerCertificate creates an certificate
func (s *Service) CreateListenerCertificate(listenerID int, req CreateCertificateRequest) (int, error) {
	if listenerID < 1 {
		return 0, fmt.Errorf("invalid listener id")
	}
	body, err := connection.Post[Certificate](s.connection, fmt.Sprintf("/loadbalancers/v2/listeners/%d/certs", listenerID), &req, connection.NotFoundResponseHandler(&CertificateNotFoundError{ID: listenerID}))
	return body.Data.ID, err
}

// PatchListenerCertificate patches an certificate
func (s *Service) PatchListenerCertificate(listenerID int, certificateID int, req PatchCertificateRequest) error {
	if listenerID < 1 {
		return fmt.Errorf("invalid listener id")
	}
	if certificateID < 1 {
		return fmt.Errorf("invalid certificate id")
	}
	return connection.PatchRaw(s.connection, fmt.Sprintf("/loadbalancers/v2/listeners/%d/certs/%d", listenerID, certificateID), &req, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&CertificateNotFoundError{ID: listenerID}))
}

// DeleteListenerCertificate deletes a certificate
func (s *Service) DeleteListenerCertificate(listenerID int, certificateID int) error {
	if listenerID < 1 {
		return fmt.Errorf("invalid listener id")
	}
	if certificateID < 1 {
		return fmt.Errorf("invalid certificate id")
	}
	return connection.DeleteRaw(s.connection, fmt.Sprintf("/loadbalancers/v2/listeners/%d/certs/%d", listenerID, certificateID), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&CertificateNotFoundError{ID: listenerID}))
}
