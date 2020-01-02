package ltaas

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetLatestAgreement retrieves the latest agreement for given type
func (s *Service) GetLatestAgreement(agreementType AgreementType) (Agreement, error) {
	body, err := s.getLatestAgreementResponseBody(agreementType)

	return body.Data, err
}

func (s *Service) getLatestAgreementResponseBody(agreementType AgreementType) (*GetAgreementResponseBody, error) {
	body := &GetAgreementResponseBody{}

	if agreementType == "" {
		return body, fmt.Errorf("invalid agreement id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ltaas/v1/agreements/latest/%s", agreementType), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &AgreementNotFoundError{Type: agreementType}
		}

		return nil
	})
}
