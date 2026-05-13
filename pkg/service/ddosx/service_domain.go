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
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Record], error) {
		return s.GetDomainRecordsPaginated(domainName, p)
	}, parameters)
}

// GetDomainRecordsPaginated retrieves a paginated list of domains
func (s *Service) GetDomainRecordsPaginated(domainName string, parameters connection.APIRequestParameters) (*connection.Paginated[Record], error) {
	if domainName == "" {
		return nil, fmt.Errorf("invalid domain name")
	}
	body, err := connection.Get[[]Record](s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/records", domainName), parameters, connection.NotFoundResponseHandler(&DomainNotFoundError{Name: domainName}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Record], error) {
		return s.GetDomainRecordsPaginated(domainName, p)
	}), err
}

// GetDomainRecord retrieves a single domain record by ID
func (s *Service) GetDomainRecord(domainName string, recordID string) (Record, error) {
	if domainName == "" {
		return Record{}, fmt.Errorf("invalid domain name")
	}
	if recordID == "" {
		return Record{}, fmt.Errorf("invalid record ID")
	}
	body, err := connection.Get[Record](s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/records/%s", domainName, recordID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&DomainRecordNotFoundError{DomainName: domainName, ID: recordID}))
	return body.Data, err
}

// CreateDomainRecord creates a new record for a domain
func (s *Service) CreateDomainRecord(domainName string, req CreateRecordRequest) (string, error) {
	if domainName == "" {
		return "", fmt.Errorf("invalid domain name")
	}
	body, err := connection.Post[Record](s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/records", domainName), &req, connection.NotFoundResponseHandler(&DomainNotFoundError{Name: domainName}))
	return body.Data.ID, err
}

// PatchDomainRecord patches a single domain record by ID
func (s *Service) PatchDomainRecord(domainName string, recordID string, req PatchRecordRequest) error {
	if domainName == "" {
		return fmt.Errorf("invalid domain name")
	}
	if recordID == "" {
		return fmt.Errorf("invalid record ID")
	}
	return connection.PatchRaw(s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/records/%s", domainName, recordID), &req, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&DomainRecordNotFoundError{ID: recordID}))
}

// DeleteDomainRecord deletes a single domain record by ID
func (s *Service) DeleteDomainRecord(domainName string, recordID string) error {
	if domainName == "" {
		return fmt.Errorf("invalid domain name")
	}
	if recordID == "" {
		return fmt.Errorf("invalid record ID")
	}
	return connection.DeleteRaw(s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/records/%s", domainName, recordID), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&DomainRecordNotFoundError{ID: recordID}))
}

// GetDomainProperties retrieves a list of domain properties
func (s *Service) GetDomainProperties(domainName string, parameters connection.APIRequestParameters) ([]DomainProperty, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[DomainProperty], error) {
		return s.GetDomainPropertiesPaginated(domainName, p)
	}, parameters)
}

// GetDomainPropertiesPaginated retrieves a paginated list of domain properties
func (s *Service) GetDomainPropertiesPaginated(domainName string, parameters connection.APIRequestParameters) (*connection.Paginated[DomainProperty], error) {
	if domainName == "" {
		return nil, fmt.Errorf("invalid domain name")
	}
	body, err := connection.Get[[]DomainProperty](s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/properties", domainName), parameters, connection.NotFoundResponseHandler(&DomainNotFoundError{Name: domainName}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[DomainProperty], error) {
		return s.GetDomainPropertiesPaginated(domainName, p)
	}), err
}

// GetDomainProperty retrieves a single domain property by ID
func (s *Service) GetDomainProperty(domainName string, propertyID string) (DomainProperty, error) {
	if domainName == "" {
		return DomainProperty{}, fmt.Errorf("invalid domain name")
	}
	if propertyID == "" {
		return DomainProperty{}, fmt.Errorf("invalid property ID")
	}
	body, err := connection.Get[DomainProperty](s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/properties/%s", domainName, propertyID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&DomainPropertyNotFoundError{ID: propertyID}))
	return body.Data, err
}

// PatchDomainProperty patches a single domain property by ID
func (s *Service) PatchDomainProperty(domainName string, propertyID string, req PatchDomainPropertyRequest) error {
	if domainName == "" {
		return fmt.Errorf("invalid domain name")
	}
	if propertyID == "" {
		return fmt.Errorf("invalid property ID")
	}
	return connection.PatchRaw(s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/properties/%s", domainName, propertyID), &req, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&DomainPropertyNotFoundError{ID: propertyID}))
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
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[WAFRuleSet], error) {
		return s.GetDomainWAFRuleSetsPaginated(domainName, p)
	}, parameters)
}

// GetDomainWAFRuleSetsPaginated retrieves a paginated list of domains
func (s *Service) GetDomainWAFRuleSetsPaginated(domainName string, parameters connection.APIRequestParameters) (*connection.Paginated[WAFRuleSet], error) {
	if domainName == "" {
		return nil, fmt.Errorf("invalid domain name")
	}
	body, err := connection.Get[[]WAFRuleSet](s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/waf/rulesets", domainName), parameters, connection.NotFoundResponseHandler(&DomainWAFNotFoundError{DomainName: domainName}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[WAFRuleSet], error) {
		return s.GetDomainWAFRuleSetsPaginated(domainName, p)
	}), err
}

// GetDomainWAFRuleSet retrieves a waf advanced rule set for a domain
func (s *Service) GetDomainWAFRuleSet(domainName string, ruleSetID string) (WAFRuleSet, error) {
	if domainName == "" {
		return WAFRuleSet{}, fmt.Errorf("invalid domain name")
	}
	if ruleSetID == "" {
		return WAFRuleSet{}, fmt.Errorf("invalid rule set ID")
	}
	body, err := connection.Get[WAFRuleSet](s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/waf/rulesets/%s", domainName, ruleSetID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&WAFRuleSetNotFoundError{ID: ruleSetID}))
	return body.Data, err
}

// PatchDomainWAFRuleSet patches a waf advanced rule set for a domain
func (s *Service) PatchDomainWAFRuleSet(domainName string, ruleSetID string, req PatchWAFRuleSetRequest) error {
	if domainName == "" {
		return fmt.Errorf("invalid domain name")
	}
	if ruleSetID == "" {
		return fmt.Errorf("invalid rule set ID")
	}
	return connection.PatchRaw(s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/waf/rulesets/%s", domainName, ruleSetID), &req, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&WAFRuleSetNotFoundError{ID: ruleSetID}))
}

// GetDomainWAFRules retrieves a list of rules
func (s *Service) GetDomainWAFRules(domainName string, parameters connection.APIRequestParameters) ([]WAFRule, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[WAFRule], error) {
		return s.GetDomainWAFRulesPaginated(domainName, p)
	}, parameters)
}

// GetDomainWAFRulesPaginated retrieves a paginated list of domains
func (s *Service) GetDomainWAFRulesPaginated(domainName string, parameters connection.APIRequestParameters) (*connection.Paginated[WAFRule], error) {
	if domainName == "" {
		return nil, fmt.Errorf("invalid domain name")
	}
	body, err := connection.Get[[]WAFRule](s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/waf/rules", domainName), parameters, connection.NotFoundResponseHandler(&DomainWAFNotFoundError{DomainName: domainName}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[WAFRule], error) {
		return s.GetDomainWAFRulesPaginated(domainName, p)
	}), err
}

// GetDomainWAFRule retrieves a waf rule for a domain
func (s *Service) GetDomainWAFRule(domainName string, ruleID string) (WAFRule, error) {
	if domainName == "" {
		return WAFRule{}, fmt.Errorf("invalid domain name")
	}
	if ruleID == "" {
		return WAFRule{}, fmt.Errorf("invalid rule ID")
	}
	body, err := connection.Get[WAFRule](s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/waf/rules/%s", domainName, ruleID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&WAFRuleNotFoundError{ID: ruleID}))
	return body.Data, err
}

// CreateDomainWAFRule creates a WAF rule
func (s *Service) CreateDomainWAFRule(domainName string, req CreateWAFRuleRequest) (string, error) {
	if domainName == "" {
		return "", fmt.Errorf("invalid domain name")
	}
	body, err := connection.Post[WAFRule](s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/waf/rules", domainName), &req, connection.NotFoundResponseHandler(&DomainNotFoundError{Name: domainName}))
	return body.Data.ID, err
}

// PatchDomainWAFRule patches a waf rule for a domain
func (s *Service) PatchDomainWAFRule(domainName string, ruleID string, req PatchWAFRuleRequest) error {
	if domainName == "" {
		return fmt.Errorf("invalid domain name")
	}
	if ruleID == "" {
		return fmt.Errorf("invalid rule ID")
	}
	return connection.PatchRaw(s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/waf/rules/%s", domainName, ruleID), &req, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&WAFRuleNotFoundError{ID: ruleID}))
}

// DeleteDomainWAFRule deletes a waf rule for a domain
func (s *Service) DeleteDomainWAFRule(domainName string, ruleID string) error {
	if domainName == "" {
		return fmt.Errorf("invalid domain name")
	}
	if ruleID == "" {
		return fmt.Errorf("invalid rule ID")
	}
	return connection.DeleteRaw(s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/waf/rules/%s", domainName, ruleID), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&WAFRuleNotFoundError{ID: ruleID}))
}

// GetDomainWAFAdvancedRules retrieves a list of rules
func (s *Service) GetDomainWAFAdvancedRules(domainName string, parameters connection.APIRequestParameters) ([]WAFAdvancedRule, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[WAFAdvancedRule], error) {
		return s.GetDomainWAFAdvancedRulesPaginated(domainName, p)
	}, parameters)
}

// GetDomainWAFAdvancedRulesPaginated retrieves a paginated list of domains
func (s *Service) GetDomainWAFAdvancedRulesPaginated(domainName string, parameters connection.APIRequestParameters) (*connection.Paginated[WAFAdvancedRule], error) {
	if domainName == "" {
		return nil, fmt.Errorf("invalid domain name")
	}
	body, err := connection.Get[[]WAFAdvancedRule](s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/waf/advanced-rules", domainName), parameters, connection.NotFoundResponseHandler(&DomainWAFNotFoundError{DomainName: domainName}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[WAFAdvancedRule], error) {
		return s.GetDomainWAFAdvancedRulesPaginated(domainName, p)
	}), err
}

// GetDomainWAFAdvancedRule retrieves a waf rule for a domain
func (s *Service) GetDomainWAFAdvancedRule(domainName string, ruleID string) (WAFAdvancedRule, error) {
	if domainName == "" {
		return WAFAdvancedRule{}, fmt.Errorf("invalid domain name")
	}
	if ruleID == "" {
		return WAFAdvancedRule{}, fmt.Errorf("invalid rule ID")
	}
	body, err := connection.Get[WAFAdvancedRule](s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/waf/advanced-rules/%s", domainName, ruleID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&WAFAdvancedRuleNotFoundError{ID: ruleID}))
	return body.Data, err
}

// CreateDomainWAFAdvancedRule creates a WAF rule
func (s *Service) CreateDomainWAFAdvancedRule(domainName string, req CreateWAFAdvancedRuleRequest) (string, error) {
	if domainName == "" {
		return "", fmt.Errorf("invalid domain name")
	}
	body, err := connection.Post[WAFAdvancedRule](s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/waf/advanced-rules", domainName), &req, connection.NotFoundResponseHandler(&DomainNotFoundError{Name: domainName}))
	return body.Data.ID, err
}

// PatchDomainWAFAdvancedRule patches a waf advanced rule for a domain
func (s *Service) PatchDomainWAFAdvancedRule(domainName string, ruleID string, req PatchWAFAdvancedRuleRequest) error {
	if domainName == "" {
		return fmt.Errorf("invalid domain name")
	}
	if ruleID == "" {
		return fmt.Errorf("invalid rule ID")
	}
	return connection.PatchRaw(s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/waf/advanced-rules/%s", domainName, ruleID), &req, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&WAFAdvancedRuleNotFoundError{ID: ruleID}))
}

// DeleteDomainWAFAdvancedRule deletees a waf advanced rule for a domain
func (s *Service) DeleteDomainWAFAdvancedRule(domainName string, ruleID string) error {
	if domainName == "" {
		return fmt.Errorf("invalid domain name")
	}
	if ruleID == "" {
		return fmt.Errorf("invalid rule ID")
	}
	return connection.DeleteRaw(s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/waf/advanced-rules/%s", domainName, ruleID), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&WAFAdvancedRuleNotFoundError{ID: ruleID}))
}

// GetDomainACLGeoIPRules retrieves a list of rules
func (s *Service) GetDomainACLGeoIPRules(domainName string, parameters connection.APIRequestParameters) ([]ACLGeoIPRule, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[ACLGeoIPRule], error) {
		return s.GetDomainACLGeoIPRulesPaginated(domainName, p)
	}, parameters)
}

// GetDomainACLGeoIPRulesPaginated retrieves a paginated list of domains
func (s *Service) GetDomainACLGeoIPRulesPaginated(domainName string, parameters connection.APIRequestParameters) (*connection.Paginated[ACLGeoIPRule], error) {
	if domainName == "" {
		return nil, fmt.Errorf("invalid domain name")
	}
	body, err := connection.Get[[]ACLGeoIPRule](s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/acls/geo-ips", domainName), parameters, connection.NotFoundResponseHandler(&DomainNotFoundError{Name: domainName}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[ACLGeoIPRule], error) {
		return s.GetDomainACLGeoIPRulesPaginated(domainName, p)
	}), err
}

// GetDomainACLGeoIPRule retrieves a single ACL GeoIP rule for a domain
func (s *Service) GetDomainACLGeoIPRule(domainName string, ruleID string) (ACLGeoIPRule, error) {
	if domainName == "" {
		return ACLGeoIPRule{}, fmt.Errorf("invalid domain name")
	}
	if ruleID == "" {
		return ACLGeoIPRule{}, fmt.Errorf("invalid rule ID")
	}
	body, err := connection.Get[ACLGeoIPRule](s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/acls/geo-ips/%s", domainName, ruleID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&ACLGeoIPRuleNotFoundError{ID: ruleID}))
	return body.Data, err
}

// CreateDomainACLGeoIPRule creates an ACL GeoIP rule
func (s *Service) CreateDomainACLGeoIPRule(domainName string, req CreateACLGeoIPRuleRequest) (string, error) {
	if domainName == "" {
		return "", fmt.Errorf("invalid domain name")
	}
	body, err := connection.Post[ACLGeoIPRule](s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/acls/geo-ips", domainName), &req, connection.NotFoundResponseHandler(&DomainNotFoundError{Name: domainName}))
	return body.Data.ID, err
}

// PatchDomainACLGeoIPRule patches an ACL GeoIP rule
func (s *Service) PatchDomainACLGeoIPRule(domainName string, ruleID string, req PatchACLGeoIPRuleRequest) error {
	if domainName == "" {
		return fmt.Errorf("invalid domain name")
	}
	if ruleID == "" {
		return fmt.Errorf("invalid rule ID")
	}
	_, err := connection.Patch[ACLGeoIPRule](s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/acls/geo-ips/%s", domainName, ruleID), &req, connection.NotFoundResponseHandler(&ACLGeoIPRuleNotFoundError{ID: ruleID}))
	return err
}

// DeleteDomainACLGeoIPRule deletes an ACL GeoIP rule
func (s *Service) DeleteDomainACLGeoIPRule(domainName string, ruleID string) error {
	if domainName == "" {
		return fmt.Errorf("invalid domain name")
	}
	if ruleID == "" {
		return fmt.Errorf("invalid rule ID")
	}
	return connection.DeleteRaw(s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/acls/geo-ips/%s", domainName, ruleID), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&ACLGeoIPRuleNotFoundError{ID: ruleID}))
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
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[ACLIPRule], error) {
		return s.GetDomainACLIPRulesPaginated(domainName, p)
	}, parameters)
}

// GetDomainACLIPRulesPaginated retrieves a paginated list of domains
func (s *Service) GetDomainACLIPRulesPaginated(domainName string, parameters connection.APIRequestParameters) (*connection.Paginated[ACLIPRule], error) {
	if domainName == "" {
		return nil, fmt.Errorf("invalid domain name")
	}
	body, err := connection.Get[[]ACLIPRule](s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/acls/ips", domainName), parameters, connection.NotFoundResponseHandler(&DomainWAFNotFoundError{DomainName: domainName}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[ACLIPRule], error) {
		return s.GetDomainACLIPRulesPaginated(domainName, p)
	}), err
}

// GetDomainACLIPRule retrieves a single ACL IP rule for a domain
func (s *Service) GetDomainACLIPRule(domainName string, ruleID string) (ACLIPRule, error) {
	if domainName == "" {
		return ACLIPRule{}, fmt.Errorf("invalid domain name")
	}
	if ruleID == "" {
		return ACLIPRule{}, fmt.Errorf("invalid rule ID")
	}
	body, err := connection.Get[ACLIPRule](s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/acls/ips/%s", domainName, ruleID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&ACLIPRuleNotFoundError{ID: ruleID}))
	return body.Data, err
}

// CreateDomainACLIPRule creates an ACL IP rule
func (s *Service) CreateDomainACLIPRule(domainName string, req CreateACLIPRuleRequest) (string, error) {
	if domainName == "" {
		return "", fmt.Errorf("invalid domain name")
	}
	body, err := connection.Post[ACLIPRule](s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/acls/ips", domainName), &req, connection.NotFoundResponseHandler(&DomainNotFoundError{Name: domainName}))
	return body.Data.ID, err
}

// PatchDomainACLIPRule patches an ACL IP rule
func (s *Service) PatchDomainACLIPRule(domainName string, ruleID string, req PatchACLIPRuleRequest) error {
	if domainName == "" {
		return fmt.Errorf("invalid domain name")
	}
	if ruleID == "" {
		return fmt.Errorf("invalid rule ID")
	}
	_, err := connection.Patch[ACLIPRule](s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/acls/ips/%s", domainName, ruleID), &req, connection.NotFoundResponseHandler(&ACLIPRuleNotFoundError{ID: ruleID}))
	return err
}

// DeleteDomainACLIPRule deletes an ACL IP rule
func (s *Service) DeleteDomainACLIPRule(domainName string, ruleID string) error {
	if domainName == "" {
		return fmt.Errorf("invalid domain name")
	}
	if ruleID == "" {
		return fmt.Errorf("invalid rule ID")
	}
	return connection.DeleteRaw(s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/acls/ips/%s", domainName, ruleID), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&ACLIPRuleNotFoundError{ID: ruleID}))
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
	if domainName == "" {
		return "", fmt.Errorf("invalid domain name")
	}
	body, err := connection.Post[CDNRule](s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/cdn/rules", domainName), &req, connection.NotFoundResponseHandler(&DomainCDNConfigurationNotFoundError{DomainName: domainName}))
	return body.Data.ID, err
}

// GetDomainCDNRules retrieves a list of rules
func (s *Service) GetDomainCDNRules(domainName string, parameters connection.APIRequestParameters) ([]CDNRule, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[CDNRule], error) {
		return s.GetDomainCDNRulesPaginated(domainName, p)
	}, parameters)
}

// GetDomainCDNRulesPaginated retrieves a paginated list of domains
func (s *Service) GetDomainCDNRulesPaginated(domainName string, parameters connection.APIRequestParameters) (*connection.Paginated[CDNRule], error) {
	if domainName == "" {
		return nil, fmt.Errorf("invalid domain name")
	}
	body, err := connection.Get[[]CDNRule](s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/cdn/rules", domainName), parameters, connection.NotFoundResponseHandler(&DomainCDNConfigurationNotFoundError{DomainName: domainName}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[CDNRule], error) {
		return s.GetDomainCDNRulesPaginated(domainName, p)
	}), err
}

// GetDomainCDNRule retrieves a CDN rule
func (s *Service) GetDomainCDNRule(domainName string, ruleID string) (CDNRule, error) {
	if domainName == "" {
		return CDNRule{}, fmt.Errorf("invalid domain name")
	}
	if ruleID == "" {
		return CDNRule{}, fmt.Errorf("invalid rule ID")
	}
	body, err := connection.Get[CDNRule](s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/cdn/rules/%s", domainName, ruleID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&CDNRuleNotFoundError{ID: ruleID}))
	return body.Data, err
}

// PatchDomainCDNRule patches a CDN rule
func (s *Service) PatchDomainCDNRule(domainName string, ruleID string, req PatchCDNRuleRequest) error {
	if domainName == "" {
		return fmt.Errorf("invalid domain name")
	}
	if ruleID == "" {
		return fmt.Errorf("invalid rule ID")
	}
	return connection.PatchRaw(s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/cdn/rules/%s", domainName, ruleID), &req, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&CDNRuleNotFoundError{ID: ruleID}))
}

// DeleteDomainCDNRule removes a CDN rule
func (s *Service) DeleteDomainCDNRule(domainName string, ruleID string) error {
	if domainName == "" {
		return fmt.Errorf("invalid domain name")
	}
	if ruleID == "" {
		return fmt.Errorf("invalid rule ID")
	}
	return connection.DeleteRaw(s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/cdn/rules/%s", domainName, ruleID), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&CDNRuleNotFoundError{ID: ruleID}))
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
	if domainName == "" {
		return "", fmt.Errorf("invalid domain name")
	}
	body, err := connection.Post[HSTSRule](s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/hsts/rules", domainName), &req, connection.NotFoundResponseHandler(&DomainHSTSConfigurationNotFoundError{DomainName: domainName}))
	return body.Data.ID, err
}

// GetDomainHSTSRules retrieves a list of rules
func (s *Service) GetDomainHSTSRules(domainName string, parameters connection.APIRequestParameters) ([]HSTSRule, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[HSTSRule], error) {
		return s.GetDomainHSTSRulesPaginated(domainName, p)
	}, parameters)
}

// GetDomainHSTSRulesPaginated retrieves a paginated list of domains
func (s *Service) GetDomainHSTSRulesPaginated(domainName string, parameters connection.APIRequestParameters) (*connection.Paginated[HSTSRule], error) {
	if domainName == "" {
		return nil, fmt.Errorf("invalid domain name")
	}
	body, err := connection.Get[[]HSTSRule](s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/hsts/rules", domainName), parameters, connection.NotFoundResponseHandler(&DomainHSTSConfigurationNotFoundError{DomainName: domainName}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[HSTSRule], error) {
		return s.GetDomainHSTSRulesPaginated(domainName, p)
	}), err
}

// GetDomainHSTSRule retrieves a HSTS rule
func (s *Service) GetDomainHSTSRule(domainName string, ruleID string) (HSTSRule, error) {
	if domainName == "" {
		return HSTSRule{}, fmt.Errorf("invalid domain name")
	}
	if ruleID == "" {
		return HSTSRule{}, fmt.Errorf("invalid rule ID")
	}
	body, err := connection.Get[HSTSRule](s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/hsts/rules/%s", domainName, ruleID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&HSTSRuleNotFoundError{ID: ruleID}))
	return body.Data, err
}

// PatchDomainHSTSRule patches a HSTS rule
func (s *Service) PatchDomainHSTSRule(domainName string, ruleID string, req PatchHSTSRuleRequest) error {
	if domainName == "" {
		return fmt.Errorf("invalid domain name")
	}
	if ruleID == "" {
		return fmt.Errorf("invalid rule ID")
	}
	return connection.PatchRaw(s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/hsts/rules/%s", domainName, ruleID), &req, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&HSTSRuleNotFoundError{ID: ruleID}))
}

// DeleteDomainHSTSRule removes a HSTS rule
func (s *Service) DeleteDomainHSTSRule(domainName string, ruleID string) error {
	if domainName == "" {
		return fmt.Errorf("invalid domain name")
	}
	if ruleID == "" {
		return fmt.Errorf("invalid rule ID")
	}
	return connection.DeleteRaw(s.connection, fmt.Sprintf("/ddosx/v1/domains/%s/hsts/rules/%s", domainName, ruleID), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&HSTSRuleNotFoundError{ID: ruleID}))
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
