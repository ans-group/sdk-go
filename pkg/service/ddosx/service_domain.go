package ddosx

import (
	"fmt"
	"io"
	"mime"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) domainRes() *resource.Resource[Domain, string] {
	return resource.NewStringResourceWithIdentifier[Domain](s.connection, "/ddosx/v1/domains", "domain", "name",
		func(id string) error { return &DomainNotFoundError{Name: id} })
}

func (s *Service) domainRecordRes() *resource.SubResource[Record, string, string] {
	return resource.NewStringStringSubResource[Record](s.connection,
		func(domainName string) string { return fmt.Sprintf("/ddosx/v1/domains/%s/records", domainName) },
		"domain", "name", func(domainName string) error { return &DomainNotFoundError{Name: domainName} },
		"record", "ID", func(domainName, recordID string) error {
			return &DomainRecordNotFoundError{DomainName: domainName, ID: recordID}
		})
}

func (s *Service) domainPropertyRes() *resource.SubResource[DomainProperty, string, string] {
	return resource.NewStringStringSubResource[DomainProperty](s.connection,
		func(domainName string) string { return fmt.Sprintf("/ddosx/v1/domains/%s/properties", domainName) },
		"domain", "name", func(domainName string) error { return &DomainNotFoundError{Name: domainName} },
		"property", "ID", func(_, propertyID string) error { return &DomainPropertyNotFoundError{ID: propertyID} })
}

func (s *Service) domainWAFRuleSetRes() *resource.SubResource[WAFRuleSet, string, string] {
	return resource.NewStringStringSubResource[WAFRuleSet](s.connection,
		func(domainName string) string { return fmt.Sprintf("/ddosx/v1/domains/%s/waf/rulesets", domainName) },
		"domain", "name", func(domainName string) error { return &DomainNotFoundError{Name: domainName} },
		"rule set", "ID", func(_, ruleSetID string) error { return &WAFRuleSetNotFoundError{ID: ruleSetID} })
}

func (s *Service) domainWAFRuleRes() *resource.SubResource[WAFRule, string, string] {
	return resource.NewStringStringSubResource[WAFRule](s.connection,
		func(domainName string) string { return fmt.Sprintf("/ddosx/v1/domains/%s/waf/rules", domainName) },
		"domain", "name", func(domainName string) error { return &DomainNotFoundError{Name: domainName} },
		"rule", "ID", func(_, ruleID string) error { return &WAFRuleNotFoundError{ID: ruleID} })
}

func (s *Service) domainWAFAdvancedRuleRes() *resource.SubResource[WAFAdvancedRule, string, string] {
	return resource.NewStringStringSubResource[WAFAdvancedRule](s.connection,
		func(domainName string) string {
			return fmt.Sprintf("/ddosx/v1/domains/%s/waf/advanced-rules", domainName)
		},
		"domain", "name", func(domainName string) error { return &DomainNotFoundError{Name: domainName} },
		"rule", "ID", func(_, ruleID string) error { return &WAFAdvancedRuleNotFoundError{ID: ruleID} })
}

func (s *Service) domainACLGeoIPRuleRes() *resource.SubResource[ACLGeoIPRule, string, string] {
	return resource.NewStringStringSubResource[ACLGeoIPRule](s.connection,
		func(domainName string) string { return fmt.Sprintf("/ddosx/v1/domains/%s/acls/geo-ips", domainName) },
		"domain", "name", func(domainName string) error { return &DomainNotFoundError{Name: domainName} },
		"rule", "ID", func(_, ruleID string) error { return &ACLGeoIPRuleNotFoundError{ID: ruleID} })
}

func (s *Service) domainACLIPRuleRes() *resource.SubResource[ACLIPRule, string, string] {
	return resource.NewStringStringSubResource[ACLIPRule](s.connection,
		func(domainName string) string { return fmt.Sprintf("/ddosx/v1/domains/%s/acls/ips", domainName) },
		"domain", "name", func(domainName string) error { return &DomainNotFoundError{Name: domainName} },
		"rule", "ID", func(_, ruleID string) error { return &ACLIPRuleNotFoundError{ID: ruleID} })
}

func (s *Service) domainCDNRuleRes() *resource.SubResource[CDNRule, string, string] {
	return resource.NewStringStringSubResource[CDNRule](s.connection,
		func(domainName string) string { return fmt.Sprintf("/ddosx/v1/domains/%s/cdn/rules", domainName) },
		"domain", "name",
		func(domainName string) error { return &DomainCDNConfigurationNotFoundError{DomainName: domainName} },
		"rule", "ID", func(_, ruleID string) error { return &CDNRuleNotFoundError{ID: ruleID} })
}

func (s *Service) domainHSTSRuleRes() *resource.SubResource[HSTSRule, string, string] {
	return resource.NewStringStringSubResource[HSTSRule](s.connection,
		func(domainName string) string { return fmt.Sprintf("/ddosx/v1/domains/%s/hsts/rules", domainName) },
		"domain", "name",
		func(domainName string) error { return &DomainHSTSConfigurationNotFoundError{DomainName: domainName} },
		"rule", "ID", func(_, ruleID string) error { return &HSTSRuleNotFoundError{ID: ruleID} })
}

// GetDomains retrieves a list of domains
func (s *Service) GetDomains(parameters connection.APIRequestParameters) ([]Domain, error) {
	return s.domainRes().List(parameters)
}

// GetDomainsPaginated retrieves a paginated list of domains
func (s *Service) GetDomainsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Domain], error) {
	return s.domainRes().ListPaginated(parameters)
}

// GetDomain retrieves a single domain by name
func (s *Service) GetDomain(domainName string) (Domain, error) {
	return s.domainRes().Get(domainName)
}

// CreateDomain creates a new domain
func (s *Service) CreateDomain(req CreateDomainRequest) error {
	_, err := connection.Post[struct{}](s.connection, "/ddosx/v1/domains", &req)
	return err
}

// DeleteDomain removes a domain
func (s *Service) DeleteDomain(domainName string, req DeleteDomainRequest) error {
	if domainName == "" {
		return fmt.Errorf("invalid domain name")
	}
	return connection.DeleteRaw(s.connection, fmt.Sprintf("/ddosx/v1/domains/%s", domainName), &req, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&DomainNotFoundError{Name: domainName}))
}

// DeployDomain deploys/commits changes to a domain
func (s *Service) DeployDomain(domainName string) error {
	if domainName == "" {
		return fmt.Errorf("invalid domain name")
	}
	return connection.PostRaw(s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/deploy", domainName), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&DomainNotFoundError{Name: domainName}))
}

// GetDomainRecords retrieves a list of records
func (s *Service) GetDomainRecords(domainName string, parameters connection.APIRequestParameters) ([]Record, error) {
	return s.domainRecordRes().List(domainName, parameters)
}

// GetDomainRecordsPaginated retrieves a paginated list of domains
func (s *Service) GetDomainRecordsPaginated(domainName string, parameters connection.APIRequestParameters) (*connection.Paginated[Record], error) {
	return s.domainRecordRes().ListPaginated(domainName, parameters)
}

// GetDomainRecord retrieves a single domain record by ID
func (s *Service) GetDomainRecord(domainName string, recordID string) (Record, error) {
	return s.domainRecordRes().Get(domainName, recordID)
}

// CreateDomainRecord creates a new record for a domain
func (s *Service) CreateDomainRecord(domainName string, req CreateRecordRequest) (string, error) {
	record, err := s.domainRecordRes().Create(domainName, &req)
	return record.ID, err
}

// PatchDomainRecord patches a single domain record by ID
func (s *Service) PatchDomainRecord(domainName string, recordID string, req PatchRecordRequest) error {
	return s.domainRecordRes().Patch(domainName, recordID, &req)
}

// DeleteDomainRecord deletes a single domain record by ID
func (s *Service) DeleteDomainRecord(domainName string, recordID string) error {
	return s.domainRecordRes().Delete(domainName, recordID)
}

// GetDomainProperties retrieves a list of domain properties
func (s *Service) GetDomainProperties(domainName string, parameters connection.APIRequestParameters) ([]DomainProperty, error) {
	return s.domainPropertyRes().List(domainName, parameters)
}

// GetDomainPropertiesPaginated retrieves a paginated list of domain properties
func (s *Service) GetDomainPropertiesPaginated(domainName string, parameters connection.APIRequestParameters) (*connection.Paginated[DomainProperty], error) {
	return s.domainPropertyRes().ListPaginated(domainName, parameters)
}

// GetDomainProperty retrieves a single domain property by ID
func (s *Service) GetDomainProperty(domainName string, propertyID string) (DomainProperty, error) {
	return s.domainPropertyRes().Get(domainName, propertyID)
}

// PatchDomainProperty patches a single domain property by ID
func (s *Service) PatchDomainProperty(domainName string, propertyID string, req PatchDomainPropertyRequest) error {
	return s.domainPropertyRes().Patch(domainName, propertyID, &req)
}

// GetDomainWAF retrieves the WAF configuration for a domain
func (s *Service) GetDomainWAF(domainName string) (WAF, error) {
	if domainName == "" {
		return WAF{}, fmt.Errorf("invalid domain name")
	}
	body, err := connection.Get[WAF](s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/waf", domainName), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&DomainWAFNotFoundError{DomainName: domainName}))
	return body.Data, err
}

// CreateDomainWAF creates the WAF configuration for a domain
func (s *Service) CreateDomainWAF(domainName string, req CreateWAFRequest) error {
	if domainName == "" {
		return fmt.Errorf("invalid domain name")
	}
	return connection.PostRaw(s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/waf", domainName), &req, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&DomainNotFoundError{Name: domainName}))
}

// PatchDomainWAF patches the WAF configuration for a domain
func (s *Service) PatchDomainWAF(domainName string, req PatchWAFRequest) error {
	if domainName == "" {
		return fmt.Errorf("invalid domain name")
	}
	return connection.PatchRaw(s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/waf", domainName), &req, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&DomainWAFNotFoundError{DomainName: domainName}))
}

// DeleteDomainWAF deletes the WAF configuration for a domain
func (s *Service) DeleteDomainWAF(domainName string) error {
	if domainName == "" {
		return fmt.Errorf("invalid domain name")
	}
	return connection.DeleteRaw(s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/waf", domainName), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&DomainWAFNotFoundError{DomainName: domainName}))
}

// GetDomainWAFRuleSets retrieves a list of rulesets
func (s *Service) GetDomainWAFRuleSets(domainName string, parameters connection.APIRequestParameters) ([]WAFRuleSet, error) {
	return s.domainWAFRuleSetRes().List(domainName, parameters)
}

// GetDomainWAFRuleSetsPaginated retrieves a paginated list of domains
func (s *Service) GetDomainWAFRuleSetsPaginated(domainName string, parameters connection.APIRequestParameters) (*connection.Paginated[WAFRuleSet], error) {
	return s.domainWAFRuleSetRes().ListPaginated(domainName, parameters)
}

// GetDomainWAFRuleSet retrieves a waf advanced rule set for a domain
func (s *Service) GetDomainWAFRuleSet(domainName string, ruleSetID string) (WAFRuleSet, error) {
	return s.domainWAFRuleSetRes().Get(domainName, ruleSetID)
}

// PatchDomainWAFRuleSet patches a waf advanced rule set for a domain
func (s *Service) PatchDomainWAFRuleSet(domainName string, ruleSetID string, req PatchWAFRuleSetRequest) error {
	return s.domainWAFRuleSetRes().Patch(domainName, ruleSetID, &req)
}

// GetDomainWAFRules retrieves a list of rules
func (s *Service) GetDomainWAFRules(domainName string, parameters connection.APIRequestParameters) ([]WAFRule, error) {
	return s.domainWAFRuleRes().List(domainName, parameters)
}

// GetDomainWAFRulesPaginated retrieves a paginated list of domains
func (s *Service) GetDomainWAFRulesPaginated(domainName string, parameters connection.APIRequestParameters) (*connection.Paginated[WAFRule], error) {
	return s.domainWAFRuleRes().ListPaginated(domainName, parameters)
}

// GetDomainWAFRule retrieves a waf rule for a domain
func (s *Service) GetDomainWAFRule(domainName string, ruleID string) (WAFRule, error) {
	return s.domainWAFRuleRes().Get(domainName, ruleID)
}

// CreateDomainWAFRule creates a WAF rule
func (s *Service) CreateDomainWAFRule(domainName string, req CreateWAFRuleRequest) (string, error) {
	rule, err := s.domainWAFRuleRes().Create(domainName, &req)
	return rule.ID, err
}

// PatchDomainWAFRule patches a waf rule for a domain
func (s *Service) PatchDomainWAFRule(domainName string, ruleID string, req PatchWAFRuleRequest) error {
	return s.domainWAFRuleRes().Patch(domainName, ruleID, &req)
}

// DeleteDomainWAFRule deletes a waf rule for a domain
func (s *Service) DeleteDomainWAFRule(domainName string, ruleID string) error {
	return s.domainWAFRuleRes().Delete(domainName, ruleID)
}

// GetDomainWAFAdvancedRules retrieves a list of rules
func (s *Service) GetDomainWAFAdvancedRules(domainName string, parameters connection.APIRequestParameters) ([]WAFAdvancedRule, error) {
	return s.domainWAFAdvancedRuleRes().List(domainName, parameters)
}

// GetDomainWAFAdvancedRulesPaginated retrieves a paginated list of domains
func (s *Service) GetDomainWAFAdvancedRulesPaginated(domainName string, parameters connection.APIRequestParameters) (*connection.Paginated[WAFAdvancedRule], error) {
	return s.domainWAFAdvancedRuleRes().ListPaginated(domainName, parameters)
}

// GetDomainWAFAdvancedRule retrieves a waf rule for a domain
func (s *Service) GetDomainWAFAdvancedRule(domainName string, ruleID string) (WAFAdvancedRule, error) {
	return s.domainWAFAdvancedRuleRes().Get(domainName, ruleID)
}

// CreateDomainWAFAdvancedRule creates a WAF rule
func (s *Service) CreateDomainWAFAdvancedRule(domainName string, req CreateWAFAdvancedRuleRequest) (string, error) {
	rule, err := s.domainWAFAdvancedRuleRes().Create(domainName, &req)
	return rule.ID, err
}

// PatchDomainWAFAdvancedRule patches a waf advanced rule for a domain
func (s *Service) PatchDomainWAFAdvancedRule(domainName string, ruleID string, req PatchWAFAdvancedRuleRequest) error {
	return s.domainWAFAdvancedRuleRes().Patch(domainName, ruleID, &req)
}

// DeleteDomainWAFAdvancedRule deletees a waf advanced rule for a domain
func (s *Service) DeleteDomainWAFAdvancedRule(domainName string, ruleID string) error {
	return s.domainWAFAdvancedRuleRes().Delete(domainName, ruleID)
}

// GetDomainACLGeoIPRules retrieves a list of rules
func (s *Service) GetDomainACLGeoIPRules(domainName string, parameters connection.APIRequestParameters) ([]ACLGeoIPRule, error) {
	return s.domainACLGeoIPRuleRes().List(domainName, parameters)
}

// GetDomainACLGeoIPRulesPaginated retrieves a paginated list of domains
func (s *Service) GetDomainACLGeoIPRulesPaginated(domainName string, parameters connection.APIRequestParameters) (*connection.Paginated[ACLGeoIPRule], error) {
	return s.domainACLGeoIPRuleRes().ListPaginated(domainName, parameters)
}

// GetDomainACLGeoIPRule retrieves a single ACL GeoIP rule for a domain
func (s *Service) GetDomainACLGeoIPRule(domainName string, ruleID string) (ACLGeoIPRule, error) {
	return s.domainACLGeoIPRuleRes().Get(domainName, ruleID)
}

// CreateDomainACLGeoIPRule creates an ACL GeoIP rule
func (s *Service) CreateDomainACLGeoIPRule(domainName string, req CreateACLGeoIPRuleRequest) (string, error) {
	rule, err := s.domainACLGeoIPRuleRes().Create(domainName, &req)
	return rule.ID, err
}

// PatchDomainACLGeoIPRule patches an ACL GeoIP rule
func (s *Service) PatchDomainACLGeoIPRule(domainName string, ruleID string, req PatchACLGeoIPRuleRequest) error {
	return s.domainACLGeoIPRuleRes().Patch(domainName, ruleID, &req)
}

// DeleteDomainACLGeoIPRule deletes an ACL GeoIP rule
func (s *Service) DeleteDomainACLGeoIPRule(domainName string, ruleID string) error {
	return s.domainACLGeoIPRuleRes().Delete(domainName, ruleID)
}

// GetDomainACLGeoIPRulesMode retrieves the mode for ACL GeoIP rules
func (s *Service) GetDomainACLGeoIPRulesMode(domainName string) (ACLGeoIPRulesMode, error) {
	if domainName == "" {
		return "", fmt.Errorf("invalid domain name")
	}
	body, err := connection.Get[GetACLGeoIPRulesModeResponseBodyData](s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/acls/geo-ips/mode", domainName), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&DomainNotFoundError{Name: domainName}))
	return body.Data.Mode, err
}

// PatchDomainACLGeoIPRulesMode patches the mode for ACL GeoIP rules
func (s *Service) PatchDomainACLGeoIPRulesMode(domainName string, req PatchACLGeoIPRulesModeRequest) error {
	if domainName == "" {
		return fmt.Errorf("invalid domain name")
	}
	return connection.PatchRaw(s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/acls/geo-ips/mode", domainName), &req, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&DomainNotFoundError{Name: domainName}))
}

// GetDomainACLIPRules retrieves a list of rules
func (s *Service) GetDomainACLIPRules(domainName string, parameters connection.APIRequestParameters) ([]ACLIPRule, error) {
	return s.domainACLIPRuleRes().List(domainName, parameters)
}

// GetDomainACLIPRulesPaginated retrieves a paginated list of domains
func (s *Service) GetDomainACLIPRulesPaginated(domainName string, parameters connection.APIRequestParameters) (*connection.Paginated[ACLIPRule], error) {
	return s.domainACLIPRuleRes().ListPaginated(domainName, parameters)
}

// GetDomainACLIPRule retrieves a single ACL IP rule for a domain
func (s *Service) GetDomainACLIPRule(domainName string, ruleID string) (ACLIPRule, error) {
	return s.domainACLIPRuleRes().Get(domainName, ruleID)
}

// CreateDomainACLIPRule creates an ACL IP rule
func (s *Service) CreateDomainACLIPRule(domainName string, req CreateACLIPRuleRequest) (string, error) {
	rule, err := s.domainACLIPRuleRes().Create(domainName, &req)
	return rule.ID, err
}

// PatchDomainACLIPRule patches an ACL IP rule
func (s *Service) PatchDomainACLIPRule(domainName string, ruleID string, req PatchACLIPRuleRequest) error {
	return s.domainACLIPRuleRes().Patch(domainName, ruleID, &req)
}

// DeleteDomainACLIPRule deletes an ACL IP rule
func (s *Service) DeleteDomainACLIPRule(domainName string, ruleID string) error {
	return s.domainACLIPRuleRes().Delete(domainName, ruleID)
}

// DownloadDomainVerificationFile downloads the verification file for a domain, returning
// the file contents, file name and an error
func (s *Service) DownloadDomainVerificationFile(domainName string) (content string, filename string, err error) {
	stream, filename, err := s.DownloadDomainVerificationFileStream(domainName)
	if err != nil {
		return "", "", err
	}

	defer stream.Close() //nolint:errcheck

	bodyBytes, err := io.ReadAll(stream)
	if err != nil {
		return "", "", err
	}

	return string(bodyBytes), filename, nil
}

// DownloadDomainVerificationFileStream downloads the verification file for a domain, returning
// a stream of the file contents, file name and an error
func (s *Service) DownloadDomainVerificationFileStream(domainName string) (contentStream io.ReadCloser, filename string, err error) {
	response, err := s.downloadDomainVerificationFileResponse(domainName)
	if err != nil {
		return nil, "", err
	}

	_, params, err := mime.ParseMediaType(response.Header.Get("Content-Disposition"))
	if err != nil {
		return nil, "", err
	}

	return response.Body, params["filename"], nil
}

func (s *Service) downloadDomainVerificationFileResponse(domainName string) (*connection.APIResponse, error) {
	response := &connection.APIResponse{}

	if domainName == "" {
		return response, fmt.Errorf("invalid domain name")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ddosx/v1/domains/%s/verify/file-upload", domainName), connection.APIRequestParameters{})
	if err != nil {
		return response, err
	}

	if response.StatusCode == 404 {
		return response, &DomainNotFoundError{Name: domainName}
	}

	return response, response.HandleResponse(nil)
}

// VerifyDomainDNS verifies a domain via DNS method
func (s *Service) VerifyDomainDNS(domainName string) error {
	if domainName == "" {
		return fmt.Errorf("invalid domain name")
	}
	return connection.PostRaw(s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/verify/dns", domainName), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&DomainNotFoundError{Name: domainName}), connection.StatusCodeResponseHandler(400, &DomainAlreadyVerifiedError{Name: domainName}))
}

// VerifyDomainFileUpload verifies a domain via file-upload method
func (s *Service) VerifyDomainFileUpload(domainName string) error {
	if domainName == "" {
		return fmt.Errorf("invalid domain name")
	}
	return connection.PostRaw(s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/verify/file-upload", domainName), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&DomainNotFoundError{Name: domainName}), connection.StatusCodeResponseHandler(400, &DomainAlreadyVerifiedError{Name: domainName}))
}

// AddDomainCDNConfiguration adds CDN configuration to a domain
func (s *Service) AddDomainCDNConfiguration(domainName string) error {
	if domainName == "" {
		return fmt.Errorf("invalid domain name")
	}
	return connection.PostRaw(s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/cdn", domainName), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&DomainNotFoundError{Name: domainName}))
}

// DeleteDomainCDNConfiguration removes CDN configuration from a domain
func (s *Service) DeleteDomainCDNConfiguration(domainName string) error {
	if domainName == "" {
		return fmt.Errorf("invalid domain name")
	}
	return connection.DeleteRaw(s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/cdn", domainName), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&DomainCDNConfigurationNotFoundError{DomainName: domainName}))
}

// CreateDomainCDNRule creates a CDN rule
func (s *Service) CreateDomainCDNRule(domainName string, req CreateCDNRuleRequest) (string, error) {
	rule, err := s.domainCDNRuleRes().Create(domainName, &req)
	return rule.ID, err
}

// GetDomainCDNRules retrieves a list of rules
func (s *Service) GetDomainCDNRules(domainName string, parameters connection.APIRequestParameters) ([]CDNRule, error) {
	return s.domainCDNRuleRes().List(domainName, parameters)
}

// GetDomainCDNRulesPaginated retrieves a paginated list of domains
func (s *Service) GetDomainCDNRulesPaginated(domainName string, parameters connection.APIRequestParameters) (*connection.Paginated[CDNRule], error) {
	return s.domainCDNRuleRes().ListPaginated(domainName, parameters)
}

// GetDomainCDNRule retrieves a CDN rule
func (s *Service) GetDomainCDNRule(domainName string, ruleID string) (CDNRule, error) {
	return s.domainCDNRuleRes().Get(domainName, ruleID)
}

// PatchDomainCDNRule patches a CDN rule
func (s *Service) PatchDomainCDNRule(domainName string, ruleID string, req PatchCDNRuleRequest) error {
	return s.domainCDNRuleRes().Patch(domainName, ruleID, &req)
}

// DeleteDomainCDNRule removes a CDN rule
func (s *Service) DeleteDomainCDNRule(domainName string, ruleID string) error {
	return s.domainCDNRuleRes().Delete(domainName, ruleID)
}

// PurgeDomainCDN purges cached content
func (s *Service) PurgeDomainCDN(domainName string, req PurgeCDNRequest) error {
	if domainName == "" {
		return fmt.Errorf("invalid domain name")
	}
	return connection.PostRaw(s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/cdn/purge", domainName), &req, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&DomainCDNConfigurationNotFoundError{DomainName: domainName}))
}

// GetDomainHSTSConfiguration retrieves the HSTS configuration for a domain
func (s *Service) GetDomainHSTSConfiguration(domainName string) (HSTSConfiguration, error) {
	if domainName == "" {
		return HSTSConfiguration{}, fmt.Errorf("invalid domain name")
	}
	body, err := connection.Get[HSTSConfiguration](s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/hsts", domainName), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&DomainHSTSConfigurationNotFoundError{DomainName: domainName}))
	return body.Data, err
}

// AddDomainHSTSConfiguration adds HSTS headers to a domain
func (s *Service) AddDomainHSTSConfiguration(domainName string) error {
	if domainName == "" {
		return fmt.Errorf("invalid domain name")
	}
	return connection.PostRaw(s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/hsts", domainName), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&DomainNotFoundError{Name: domainName}))
}

// DeleteDomainHSTSConfiguration removes HSTS headers to a domain
func (s *Service) DeleteDomainHSTSConfiguration(domainName string) error {
	if domainName == "" {
		return fmt.Errorf("invalid domain name")
	}
	return connection.DeleteRaw(s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/hsts", domainName), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&DomainHSTSConfigurationNotFoundError{DomainName: domainName}))
}

// CreateDomainHSTSRule creates a HSTS rule
func (s *Service) CreateDomainHSTSRule(domainName string, req CreateHSTSRuleRequest) (string, error) {
	rule, err := s.domainHSTSRuleRes().Create(domainName, &req)
	return rule.ID, err
}

// GetDomainHSTSRules retrieves a list of rules
func (s *Service) GetDomainHSTSRules(domainName string, parameters connection.APIRequestParameters) ([]HSTSRule, error) {
	return s.domainHSTSRuleRes().List(domainName, parameters)
}

// GetDomainHSTSRulesPaginated retrieves a paginated list of domains
func (s *Service) GetDomainHSTSRulesPaginated(domainName string, parameters connection.APIRequestParameters) (*connection.Paginated[HSTSRule], error) {
	return s.domainHSTSRuleRes().ListPaginated(domainName, parameters)
}

// GetDomainHSTSRule retrieves a HSTS rule
func (s *Service) GetDomainHSTSRule(domainName string, ruleID string) (HSTSRule, error) {
	return s.domainHSTSRuleRes().Get(domainName, ruleID)
}

// PatchDomainHSTSRule patches a HSTS rule
func (s *Service) PatchDomainHSTSRule(domainName string, ruleID string, req PatchHSTSRuleRequest) error {
	return s.domainHSTSRuleRes().Patch(domainName, ruleID, &req)
}

// DeleteDomainHSTSRule removes a HSTS rule
func (s *Service) DeleteDomainHSTSRule(domainName string, ruleID string) error {
	return s.domainHSTSRuleRes().Delete(domainName, ruleID)
}

// ActivateDomainDNSRouting activates DNS routing for a domain
func (s *Service) ActivateDomainDNSRouting(domainName string) error {
	if domainName == "" {
		return fmt.Errorf("invalid domain name")
	}
	return connection.PostRaw(s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/dns/active", domainName), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&DomainNotFoundError{Name: domainName}))
}

// DeactivateDomainDNSRouting deactivates DNS routing for a domain
func (s *Service) DeactivateDomainDNSRouting(domainName string) error {
	if domainName == "" {
		return fmt.Errorf("invalid domain name")
	}
	return connection.DeleteRaw(s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/dns/active", domainName), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&DomainNotFoundError{Name: domainName}))
}
