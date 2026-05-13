package safedns

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) templateRes() *resource.Resource[Template, int] {
	return resource.NewIntResource[Template](s.connection, "/safedns/v1/templates", "template",
		func(id int) error { return &TemplateNotFoundError{TemplateID: id} })
}

func (s *Service) templateRecordRes() *resource.SubResource[Record, int, int] {
	return resource.NewIntIntSubResource[Record](
		s.connection,
		func(templateID int) string { return fmt.Sprintf("/safedns/v1/templates/%d/records", templateID) },
		"template", "id", func(templateID int) error { return &TemplateNotFoundError{TemplateID: templateID} },
		"record", "id", func(templateID int, recordID int) error {
			return &TemplateRecordNotFoundError{TemplateID: templateID, RecordID: recordID}
		},
	)
}

// GetTemplates retrieves a list of templates
func (s *Service) GetTemplates(parameters connection.APIRequestParameters) ([]Template, error) {
	return s.templateRes().List(parameters)
}

// GetTemplatesPaginated retrieves a paginated list of templates
func (s *Service) GetTemplatesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Template], error) {
	return s.templateRes().ListPaginated(parameters)
}

// GetTemplate retrieves a single template by ID
func (s *Service) GetTemplate(templateID int) (Template, error) {
	return s.templateRes().Get(templateID)
}

// CreateTemplate creates a new SafeDNS template
func (s *Service) CreateTemplate(req CreateTemplateRequest) (int, error) {
	data, err := s.templateRes().Create(&req)
	return data.ID, err
}

// PatchTemplate patches a SafeDNS template
func (s *Service) PatchTemplate(templateID int, patch PatchTemplateRequest) (int, error) {
	if templateID < 1 {
		return 0, fmt.Errorf("invalid template id")
	}
	body, err := connection.Patch[Template](s.connection, fmt.Sprintf("/safedns/v1/templates/%d", templateID), &patch, connection.NotFoundResponseHandler(&TemplateNotFoundError{TemplateID: templateID}))
	return body.Data.ID, err
}

// DeleteTemplate removes a SafeDNS template
func (s *Service) DeleteTemplate(templateID int) error {
	return s.templateRes().Delete(templateID)
}

// GetTemplateRecords retrieves a list of records
func (s *Service) GetTemplateRecords(templateID int, parameters connection.APIRequestParameters) ([]Record, error) {
	return s.templateRecordRes().List(templateID, parameters)
}

// GetTemplateRecordsPaginated retrieves a paginated list of templates
func (s *Service) GetTemplateRecordsPaginated(templateID int, parameters connection.APIRequestParameters) (*connection.Paginated[Record], error) {
	return s.templateRecordRes().ListPaginated(templateID, parameters)
}

// GetTemplateRecord retrieves a single zone record by ID
func (s *Service) GetTemplateRecord(templateID int, recordID int) (Record, error) {
	return s.templateRecordRes().Get(templateID, recordID)
}

// CreateTemplateRecord creates a new SafeDNS template record
func (s *Service) CreateTemplateRecord(templateID int, req CreateRecordRequest) (int, error) {
	data, err := s.templateRecordRes().Create(templateID, &req)
	return data.ID, err
}

// PatchTemplateRecord patches a SafeDNS template record
func (s *Service) PatchTemplateRecord(templateID int, recordID int, patch PatchRecordRequest) (int, error) {
	data, err := s.templateRecordRes().PatchReturning(templateID, recordID, &patch)
	return data.ID, err
}

// DeleteTemplateRecord removes a SafeDNS template record
func (s *Service) DeleteTemplateRecord(templateID int, recordID int) error {
	return s.templateRecordRes().Delete(templateID, recordID)
}
