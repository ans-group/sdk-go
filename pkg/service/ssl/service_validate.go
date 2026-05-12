package ssl

import "github.com/ans-group/sdk-go/pkg/connection"

// ValidateCertificate validates a certificate
func (s *Service) ValidateCertificate(req ValidateRequest) (CertificateValidation, error) {
	body, err := connection.Post[CertificateValidation](s.connection, "/ssl/v1/validate", &req)
	return body.Data, err
}
