package ddosx

import (
	"fmt"
	"io"
	"io/ioutil"
	"mime"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetDomains retrieves a list of domains
func (s *Service) GetDomains(parameters connection.APIRequestParameters) ([]Domain, error) {
	return connection.InvokeRequestAll(s.GetDomainsPaginated, parameters)
}

// GetDomainsPaginated retrieves a paginated list of domains
func (s *Service) GetDomainsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Domain], error) {
	body, err := s.getDomainsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetDomainsPaginated), err
}

func (s *Service) getDomainsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Domain], error) {
	body := &connection.APIResponseBodyData[[]Domain]{}

	response, err := s.connection.Get("/ddosx/v1/domains", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetDomain retrieves a single domain by name
func (s *Service) GetDomain(domainName string) (Domain, error) {
	body, err := s.getDomainResponseBody(domainName)

	return body.Data, err
}

func (s *Service) getDomainResponseBody(domainName string) (*connection.APIResponseBodyData[Domain], error) {
	body := &connection.APIResponseBodyData[Domain]{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ddosx/v1/domains/%s", domainName), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainNotFoundError{Name: domainName}
		}

		return nil
	})
}

// CreateDomain creates a new domain
func (s *Service) CreateDomain(req CreateDomainRequest) error {
	_, err := s.createDomainResponseBody(req)

	return err
}

func (s *Service) createDomainResponseBody(req CreateDomainRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	response, err := s.connection.Post("/ddosx/v1/domains", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// DeleteDomain removes a domain
func (s *Service) DeleteDomain(domainName string, req DeleteDomainRequest) error {
	_, err := s.deleteDomainResponseBody(domainName, req)

	return err
}

func (s *Service) deleteDomainResponseBody(domainName string, req DeleteDomainRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ddosx/v1/domains/%s", domainName), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainNotFoundError{Name: domainName}
		}

		return nil
	})
}

// DeployDomain deploys/commits changes to a domain
func (s *Service) DeployDomain(domainName string) error {
	_, err := s.deployDomainResponseBody(domainName)

	return err
}

func (s *Service) deployDomainResponseBody(domainName string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}

	response, err := s.connection.Post(fmt.Sprintf("/ddosx/v1/domains/%s/deploy", domainName), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainNotFoundError{Name: domainName}
		}

		return nil
	})
}

// GetDomainRecords retrieves a list of records
func (s *Service) GetDomainRecords(domainName string, parameters connection.APIRequestParameters) ([]Record, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Record], error) {
		return s.GetDomainRecordsPaginated(domainName, p)
	}, parameters)
}

// GetDomainRecordsPaginated retrieves a paginated list of domains
func (s *Service) GetDomainRecordsPaginated(domainName string, parameters connection.APIRequestParameters) (*connection.Paginated[Record], error) {
	body, err := s.getDomainRecordsPaginatedResponseBody(domainName, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Record], error) {
		return s.GetDomainRecordsPaginated(domainName, p)
	}), err
}

func (s *Service) getDomainRecordsPaginatedResponseBody(domainName string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Record], error) {
	body := &connection.APIResponseBodyData[[]Record]{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ddosx/v1/domains/%s/records", domainName), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainNotFoundError{Name: domainName}
		}

		return nil
	})
}

// GetDomainRecord retrieves a single domain record by ID
func (s *Service) GetDomainRecord(domainName string, recordID string) (Record, error) {
	body, err := s.getDomainRecordResponseBody(domainName, recordID)

	return body.Data, err
}

func (s *Service) getDomainRecordResponseBody(domainName string, recordID string) (*connection.APIResponseBodyData[Record], error) {
	body := &connection.APIResponseBodyData[Record]{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}
	if recordID == "" {
		return body, fmt.Errorf("invalid record ID")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ddosx/v1/domains/%s/records/%s", domainName, recordID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainRecordNotFoundError{DomainName: domainName, ID: recordID}
		}

		return nil
	})
}

// CreateDomainRecord creates a new record for a domain
func (s *Service) CreateDomainRecord(domainName string, req CreateRecordRequest) (string, error) {
	body, err := s.createDomainRecordResponseBody(domainName, req)

	return body.Data.ID, err
}

func (s *Service) createDomainRecordResponseBody(domainName string, req CreateRecordRequest) (*connection.APIResponseBodyData[Record], error) {
	body := &connection.APIResponseBodyData[Record]{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}

	response, err := s.connection.Post(fmt.Sprintf("/ddosx/v1/domains/%s/records", domainName), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainNotFoundError{Name: domainName}
		}

		return nil
	})
}

// PatchDomainRecord patches a single domain record by ID
func (s *Service) PatchDomainRecord(domainName string, recordID string, req PatchRecordRequest) error {
	_, err := s.patchDomainRecordResponseBody(domainName, recordID, req)

	return err
}

func (s *Service) patchDomainRecordResponseBody(domainName string, recordID string, req PatchRecordRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}
	if recordID == "" {
		return body, fmt.Errorf("invalid record ID")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ddosx/v1/domains/%s/records/%s", domainName, recordID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainRecordNotFoundError{ID: recordID}
		}

		return nil
	})
}

// DeleteDomainRecord deletes a single domain record by ID
func (s *Service) DeleteDomainRecord(domainName string, recordID string) error {
	_, err := s.deleteDomainRecordResponseBody(domainName, recordID)

	return err
}

func (s *Service) deleteDomainRecordResponseBody(domainName string, recordID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}
	if recordID == "" {
		return body, fmt.Errorf("invalid record ID")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ddosx/v1/domains/%s/records/%s", domainName, recordID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainRecordNotFoundError{ID: recordID}
		}

		return nil
	})
}

// GetDomainProperties retrieves a list of domain properties
func (s *Service) GetDomainProperties(domainName string, parameters connection.APIRequestParameters) ([]DomainProperty, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[DomainProperty], error) {
		return s.GetDomainPropertiesPaginated(domainName, p)
	}, parameters)
}

// GetDomainPropertiesPaginated retrieves a paginated list of domain properties
func (s *Service) GetDomainPropertiesPaginated(domainName string, parameters connection.APIRequestParameters) (*connection.Paginated[DomainProperty], error) {
	body, err := s.getDomainPropertiesPaginatedResponseBody(domainName, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[DomainProperty], error) {
		return s.GetDomainPropertiesPaginated(domainName, p)
	}), err
}

func (s *Service) getDomainPropertiesPaginatedResponseBody(domainName string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]DomainProperty], error) {
	body := &connection.APIResponseBodyData[[]DomainProperty]{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ddosx/v1/domains/%s/properties", domainName), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainNotFoundError{Name: domainName}
		}

		return nil
	})
}

// GetDomainProperty retrieves a single domain property by ID
func (s *Service) GetDomainProperty(domainName string, propertyID string) (DomainProperty, error) {
	body, err := s.getDomainPropertyResponseBody(domainName, propertyID)

	return body.Data, err
}

func (s *Service) getDomainPropertyResponseBody(domainName string, propertyID string) (*connection.APIResponseBodyData[DomainProperty], error) {
	body := &connection.APIResponseBodyData[DomainProperty]{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}
	if propertyID == "" {
		return body, fmt.Errorf("invalid property ID")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ddosx/v1/domains/%s/properties/%s", domainName, propertyID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainPropertyNotFoundError{ID: propertyID}
		}

		return nil
	})
}

// PatchDomainProperty patches a single domain property by ID
func (s *Service) PatchDomainProperty(domainName string, propertyID string, req PatchDomainPropertyRequest) error {
	_, err := s.patchDomainPropertyResponseBody(domainName, propertyID, req)

	return err
}

func (s *Service) patchDomainPropertyResponseBody(domainName string, propertyID string, req PatchDomainPropertyRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}
	if propertyID == "" {
		return body, fmt.Errorf("invalid property ID")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ddosx/v1/domains/%s/properties/%s", domainName, propertyID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainPropertyNotFoundError{ID: propertyID}
		}

		return nil
	})
}

// GetDomainWAF retrieves the WAF configuration for a domain
func (s *Service) GetDomainWAF(domainName string) (WAF, error) {
	body, err := s.getDomainWAFResponseBody(domainName)

	return body.Data, err
}

func (s *Service) getDomainWAFResponseBody(domainName string) (*connection.APIResponseBodyData[WAF], error) {
	body := &connection.APIResponseBodyData[WAF]{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ddosx/v1/domains/%s/waf", domainName), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainWAFNotFoundError{DomainName: domainName}
		}

		return nil
	})
}

// CreateDomainWAF creates the WAF configuration for a domain
func (s *Service) CreateDomainWAF(domainName string, req CreateWAFRequest) error {
	_, err := s.createDomainWAFResponseBody(domainName, req)

	return err
}

func (s *Service) createDomainWAFResponseBody(domainName string, req CreateWAFRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}

	response, err := s.connection.Post(fmt.Sprintf("/ddosx/v1/domains/%s/waf", domainName), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainNotFoundError{Name: domainName}
		}

		return nil
	})
}

// PatchDomainWAF patches the WAF configuration for a domain
func (s *Service) PatchDomainWAF(domainName string, req PatchWAFRequest) error {
	_, err := s.patchDomainWAFResponseBody(domainName, req)

	return err
}

func (s *Service) patchDomainWAFResponseBody(domainName string, req PatchWAFRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ddosx/v1/domains/%s/waf", domainName), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainWAFNotFoundError{DomainName: domainName}
		}

		return nil
	})
}

// DeleteDomainWAF deletes the WAF configuration for a domain
func (s *Service) DeleteDomainWAF(domainName string) error {
	_, err := s.deleteDomainWAFResponseBody(domainName)

	return err
}

func (s *Service) deleteDomainWAFResponseBody(domainName string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ddosx/v1/domains/%s/waf", domainName), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainWAFNotFoundError{DomainName: domainName}
		}

		return nil
	})
}

// GetDomainWAFRuleSets retrieves a list of rulesets
func (s *Service) GetDomainWAFRuleSets(domainName string, parameters connection.APIRequestParameters) ([]WAFRuleSet, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[WAFRuleSet], error) {
		return s.GetDomainWAFRuleSetsPaginated(domainName, p)
	}, parameters)
}

// GetDomainWAFRuleSetsPaginated retrieves a paginated list of domains
func (s *Service) GetDomainWAFRuleSetsPaginated(domainName string, parameters connection.APIRequestParameters) (*connection.Paginated[WAFRuleSet], error) {
	body, err := s.getDomainWAFRuleSetsPaginatedResponseBody(domainName, parameters)
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[WAFRuleSet], error) {
		return s.GetDomainWAFRuleSetsPaginated(domainName, p)
	}), err
}

func (s *Service) getDomainWAFRuleSetsPaginatedResponseBody(domainName string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]WAFRuleSet], error) {
	body := &connection.APIResponseBodyData[[]WAFRuleSet]{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ddosx/v1/domains/%s/waf/rulesets", domainName), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainWAFNotFoundError{DomainName: domainName}
		}

		return nil
	})
}

// GetDomainWAFRuleSet retrieves a waf advanced rule set for a domain
func (s *Service) GetDomainWAFRuleSet(domainName string, ruleSetID string) (WAFRuleSet, error) {
	body, err := s.getDomainWAFRuleSetResponseBody(domainName, ruleSetID)

	return body.Data, err
}

func (s *Service) getDomainWAFRuleSetResponseBody(domainName string, ruleSetID string) (*connection.APIResponseBodyData[WAFRuleSet], error) {
	body := &connection.APIResponseBodyData[WAFRuleSet]{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}
	if ruleSetID == "" {
		return body, fmt.Errorf("invalid rule set ID")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ddosx/v1/domains/%s/waf/rulesets/%s", domainName, ruleSetID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &WAFRuleSetNotFoundError{ID: ruleSetID}
		}

		return nil
	})
}

// PatchDomainWAFRuleSet patches a waf advanced rule set for a domain
func (s *Service) PatchDomainWAFRuleSet(domainName string, ruleSetID string, req PatchWAFRuleSetRequest) error {
	_, err := s.patchDomainWAFRuleSetResponseBody(domainName, ruleSetID, req)

	return err
}

func (s *Service) patchDomainWAFRuleSetResponseBody(domainName string, ruleSetID string, req PatchWAFRuleSetRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}
	if ruleSetID == "" {
		return body, fmt.Errorf("invalid rule set ID")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ddosx/v1/domains/%s/waf/rulesets/%s", domainName, ruleSetID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &WAFRuleSetNotFoundError{ID: ruleSetID}
		}

		return nil
	})
}

// GetDomainWAFRules retrieves a list of rules
func (s *Service) GetDomainWAFRules(domainName string, parameters connection.APIRequestParameters) ([]WAFRule, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[WAFRule], error) {
		return s.GetDomainWAFRulesPaginated(domainName, p)
	}, parameters)
}

// GetDomainWAFRulesPaginated retrieves a paginated list of domains
func (s *Service) GetDomainWAFRulesPaginated(domainName string, parameters connection.APIRequestParameters) (*connection.Paginated[WAFRule], error) {
	body, err := s.getDomainWAFRulesPaginatedResponseBody(domainName, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[WAFRule], error) {
		return s.GetDomainWAFRulesPaginated(domainName, p)
	}), err
}

func (s *Service) getDomainWAFRulesPaginatedResponseBody(domainName string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]WAFRule], error) {
	body := &connection.APIResponseBodyData[[]WAFRule]{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ddosx/v1/domains/%s/waf/rules", domainName), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainWAFNotFoundError{DomainName: domainName}
		}

		return nil
	})
}

// GetDomainWAFRule retrieves a waf rule for a domain
func (s *Service) GetDomainWAFRule(domainName string, ruleID string) (WAFRule, error) {
	body, err := s.getDomainWAFRuleResponseBody(domainName, ruleID)

	return body.Data, err
}

func (s *Service) getDomainWAFRuleResponseBody(domainName string, ruleID string) (*connection.APIResponseBodyData[WAFRule], error) {
	body := &connection.APIResponseBodyData[WAFRule]{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}
	if ruleID == "" {
		return body, fmt.Errorf("invalid rule ID")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ddosx/v1/domains/%s/waf/rules/%s", domainName, ruleID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &WAFRuleNotFoundError{ID: ruleID}
		}

		return nil
	})
}

// CreateDomainWAFRule creates a WAF rule
func (s *Service) CreateDomainWAFRule(domainName string, req CreateWAFRuleRequest) (string, error) {
	body, err := s.createDomainWAFRuleResponseBody(domainName, req)

	return body.Data.ID, err
}

func (s *Service) createDomainWAFRuleResponseBody(domainName string, req CreateWAFRuleRequest) (*connection.APIResponseBodyData[WAFRule], error) {
	body := &connection.APIResponseBodyData[WAFRule]{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}

	response, err := s.connection.Post(fmt.Sprintf("/ddosx/v1/domains/%s/waf/rules", domainName), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainNotFoundError{Name: domainName}
		}

		return nil
	})
}

// PatchDomainWAFRule patches a waf rule for a domain
func (s *Service) PatchDomainWAFRule(domainName string, ruleID string, req PatchWAFRuleRequest) error {
	_, err := s.patchDomainWAFRuleResponseBody(domainName, ruleID, req)

	return err
}

func (s *Service) patchDomainWAFRuleResponseBody(domainName string, ruleID string, req PatchWAFRuleRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}
	if ruleID == "" {
		return body, fmt.Errorf("invalid rule ID")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ddosx/v1/domains/%s/waf/rules/%s", domainName, ruleID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &WAFRuleNotFoundError{ID: ruleID}
		}

		return nil
	})
}

// DeleteDomainWAFRule deletes a waf rule for a domain
func (s *Service) DeleteDomainWAFRule(domainName string, ruleID string) error {
	_, err := s.deleteDomainWAFRuleResponseBody(domainName, ruleID)

	return err
}

func (s *Service) deleteDomainWAFRuleResponseBody(domainName string, ruleID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}
	if ruleID == "" {
		return body, fmt.Errorf("invalid rule ID")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ddosx/v1/domains/%s/waf/rules/%s", domainName, ruleID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &WAFRuleNotFoundError{ID: ruleID}
		}

		return nil
	})
}

// GetDomainWAFAdvancedRules retrieves a list of rules
func (s *Service) GetDomainWAFAdvancedRules(domainName string, parameters connection.APIRequestParameters) ([]WAFAdvancedRule, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[WAFAdvancedRule], error) {
		return s.GetDomainWAFAdvancedRulesPaginated(domainName, p)
	}, parameters)
}

// GetDomainWAFAdvancedRulesPaginated retrieves a paginated list of domains
func (s *Service) GetDomainWAFAdvancedRulesPaginated(domainName string, parameters connection.APIRequestParameters) (*connection.Paginated[WAFAdvancedRule], error) {
	body, err := s.getDomainWAFAdvancedRulesPaginatedResponseBody(domainName, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[WAFAdvancedRule], error) {
		return s.GetDomainWAFAdvancedRulesPaginated(domainName, p)
	}), err
}

func (s *Service) getDomainWAFAdvancedRulesPaginatedResponseBody(domainName string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]WAFAdvancedRule], error) {
	body := &connection.APIResponseBodyData[[]WAFAdvancedRule]{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ddosx/v1/domains/%s/waf/advanced-rules", domainName), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainWAFNotFoundError{DomainName: domainName}
		}

		return nil
	})
}

// GetDomainWAFAdvancedRule retrieves a waf rule for a domain
func (s *Service) GetDomainWAFAdvancedRule(domainName string, ruleID string) (WAFAdvancedRule, error) {
	body, err := s.getDomainWAFAdvancedRuleResponseBody(domainName, ruleID)

	return body.Data, err
}

func (s *Service) getDomainWAFAdvancedRuleResponseBody(domainName string, ruleID string) (*connection.APIResponseBodyData[WAFAdvancedRule], error) {
	body := &connection.APIResponseBodyData[WAFAdvancedRule]{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}
	if ruleID == "" {
		return body, fmt.Errorf("invalid rule ID")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ddosx/v1/domains/%s/waf/advanced-rules/%s", domainName, ruleID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &WAFAdvancedRuleNotFoundError{ID: ruleID}
		}

		return nil
	})
}

// CreateDomainWAFAdvancedRule creates a WAF rule
func (s *Service) CreateDomainWAFAdvancedRule(domainName string, req CreateWAFAdvancedRuleRequest) (string, error) {
	body, err := s.createDomainWAFAdvancedRuleResponseBody(domainName, req)

	return body.Data.ID, err
}

func (s *Service) createDomainWAFAdvancedRuleResponseBody(domainName string, req CreateWAFAdvancedRuleRequest) (*connection.APIResponseBodyData[WAFAdvancedRule], error) {
	body := &connection.APIResponseBodyData[WAFAdvancedRule]{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}

	response, err := s.connection.Post(fmt.Sprintf("/ddosx/v1/domains/%s/waf/advanced-rules", domainName), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainNotFoundError{Name: domainName}
		}

		return nil
	})
}

// PatchDomainWAFAdvancedRule patches a waf advanced rule for a domain
func (s *Service) PatchDomainWAFAdvancedRule(domainName string, ruleID string, req PatchWAFAdvancedRuleRequest) error {
	_, err := s.patchDomainWAFAdvancedRuleResponseBody(domainName, ruleID, req)

	return err
}

func (s *Service) patchDomainWAFAdvancedRuleResponseBody(domainName string, ruleID string, req PatchWAFAdvancedRuleRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}
	if ruleID == "" {
		return body, fmt.Errorf("invalid rule ID")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ddosx/v1/domains/%s/waf/advanced-rules/%s", domainName, ruleID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &WAFAdvancedRuleNotFoundError{ID: ruleID}
		}

		return nil
	})
}

// DeleteDomainWAFAdvancedRule deletees a waf advanced rule for a domain
func (s *Service) DeleteDomainWAFAdvancedRule(domainName string, ruleID string) error {
	_, err := s.deleteDomainWAFAdvancedRuleResponseBody(domainName, ruleID)

	return err
}

func (s *Service) deleteDomainWAFAdvancedRuleResponseBody(domainName string, ruleID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}
	if ruleID == "" {
		return body, fmt.Errorf("invalid rule ID")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ddosx/v1/domains/%s/waf/advanced-rules/%s", domainName, ruleID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &WAFAdvancedRuleNotFoundError{ID: ruleID}
		}

		return nil
	})
}

// GetDomainACLGeoIPRules retrieves a list of rules
func (s *Service) GetDomainACLGeoIPRules(domainName string, parameters connection.APIRequestParameters) ([]ACLGeoIPRule, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[ACLGeoIPRule], error) {
		return s.GetDomainACLGeoIPRulesPaginated(domainName, p)
	}, parameters)
}

// GetDomainACLGeoIPRulesPaginated retrieves a paginated list of domains
func (s *Service) GetDomainACLGeoIPRulesPaginated(domainName string, parameters connection.APIRequestParameters) (*connection.Paginated[ACLGeoIPRule], error) {
	body, err := s.getDomainACLGeoIPRulesPaginatedResponseBody(domainName, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[ACLGeoIPRule], error) {
		return s.GetDomainACLGeoIPRulesPaginated(domainName, p)
	}), err
}

func (s *Service) getDomainACLGeoIPRulesPaginatedResponseBody(domainName string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]ACLGeoIPRule], error) {
	body := &connection.APIResponseBodyData[[]ACLGeoIPRule]{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ddosx/v1/domains/%s/acls/geo-ips", domainName), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainNotFoundError{Name: domainName}
		}

		return nil
	})
}

// GetDomainACLGeoIPRule retrieves a single ACL GeoIP rule for a domain
func (s *Service) GetDomainACLGeoIPRule(domainName string, ruleID string) (ACLGeoIPRule, error) {
	body, err := s.getDomainACLGeoIPRuleResponseBody(domainName, ruleID)

	return body.Data, err
}

func (s *Service) getDomainACLGeoIPRuleResponseBody(domainName string, ruleID string) (*connection.APIResponseBodyData[ACLGeoIPRule], error) {
	body := &connection.APIResponseBodyData[ACLGeoIPRule]{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}
	if ruleID == "" {
		return body, fmt.Errorf("invalid rule ID")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ddosx/v1/domains/%s/acls/geo-ips/%s", domainName, ruleID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ACLGeoIPRuleNotFoundError{ID: ruleID}
		}

		return nil
	})
}

// CreateDomainACLGeoIPRule creates an ACL GeoIP rule
func (s *Service) CreateDomainACLGeoIPRule(domainName string, req CreateACLGeoIPRuleRequest) (string, error) {
	body, err := s.createDomainACLGeoIPRuleResponseBody(domainName, req)

	return body.Data.ID, err
}

func (s *Service) createDomainACLGeoIPRuleResponseBody(domainName string, req CreateACLGeoIPRuleRequest) (*connection.APIResponseBodyData[ACLGeoIPRule], error) {
	body := &connection.APIResponseBodyData[ACLGeoIPRule]{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}

	response, err := s.connection.Post(fmt.Sprintf("/ddosx/v1/domains/%s/acls/geo-ips", domainName), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainNotFoundError{Name: domainName}
		}

		return nil
	})
}

// PatchDomainACLGeoIPRule patches an ACL GeoIP rule
func (s *Service) PatchDomainACLGeoIPRule(domainName string, ruleID string, req PatchACLGeoIPRuleRequest) error {
	_, err := s.patchDomainACLGeoIPRuleResponseBody(domainName, ruleID, req)

	return err
}

func (s *Service) patchDomainACLGeoIPRuleResponseBody(domainName string, ruleID string, req PatchACLGeoIPRuleRequest) (*connection.APIResponseBodyData[ACLGeoIPRule], error) {
	body := &connection.APIResponseBodyData[ACLGeoIPRule]{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}
	if ruleID == "" {
		return body, fmt.Errorf("invalid rule ID")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ddosx/v1/domains/%s/acls/geo-ips/%s", domainName, ruleID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ACLGeoIPRuleNotFoundError{ID: ruleID}
		}

		return nil
	})
}

// DeleteDomainACLGeoIPRule deletes an ACL GeoIP rule
func (s *Service) DeleteDomainACLGeoIPRule(domainName string, ruleID string) error {
	_, err := s.deleteDomainACLGeoIPRuleResponseBody(domainName, ruleID)

	return err
}

func (s *Service) deleteDomainACLGeoIPRuleResponseBody(domainName string, ruleID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}
	if ruleID == "" {
		return body, fmt.Errorf("invalid rule ID")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ddosx/v1/domains/%s/acls/geo-ips/%s", domainName, ruleID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ACLGeoIPRuleNotFoundError{ID: ruleID}
		}

		return nil
	})
}

// GetDomainACLGeoIPRulesMode retrieves the mode for ACL GeoIP rules
func (s *Service) GetDomainACLGeoIPRulesMode(domainName string) (ACLGeoIPRulesMode, error) {
	body, err := s.getDomainACLGeoIPRulesModeResponseBody(domainName)

	return body.Data.Mode, err
}

func (s *Service) getDomainACLGeoIPRulesModeResponseBody(domainName string) (*connection.APIResponseBodyData[GetACLGeoIPRulesModeResponseBodyData], error) {
	body := &connection.APIResponseBodyData[GetACLGeoIPRulesModeResponseBodyData]{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ddosx/v1/domains/%s/acls/geo-ips/mode", domainName), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainNotFoundError{Name: domainName}
		}

		return nil
	})
}

// PatchDomainACLGeoIPRulesMode patches the mode for ACL GeoIP rules
func (s *Service) PatchDomainACLGeoIPRulesMode(domainName string, req PatchACLGeoIPRulesModeRequest) error {
	_, err := s.patchDomainACLGeoIPRulesModeResponseBody(domainName, req)

	return err
}

func (s *Service) patchDomainACLGeoIPRulesModeResponseBody(domainName string, req PatchACLGeoIPRulesModeRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ddosx/v1/domains/%s/acls/geo-ips/mode", domainName), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainNotFoundError{Name: domainName}
		}

		return nil
	})
}

// GetDomainACLIPRules retrieves a list of rules
func (s *Service) GetDomainACLIPRules(domainName string, parameters connection.APIRequestParameters) ([]ACLIPRule, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[ACLIPRule], error) {
		return s.GetDomainACLIPRulesPaginated(domainName, p)
	}, parameters)
}

// GetDomainACLIPRulesPaginated retrieves a paginated list of domains
func (s *Service) GetDomainACLIPRulesPaginated(domainName string, parameters connection.APIRequestParameters) (*connection.Paginated[ACLIPRule], error) {
	body, err := s.getDomainACLIPRulesPaginatedResponseBody(domainName, parameters)
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[ACLIPRule], error) {
		return s.GetDomainACLIPRulesPaginated(domainName, p)
	}), err
}

func (s *Service) getDomainACLIPRulesPaginatedResponseBody(domainName string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]ACLIPRule], error) {
	body := &connection.APIResponseBodyData[[]ACLIPRule]{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ddosx/v1/domains/%s/acls/ips", domainName), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainWAFNotFoundError{DomainName: domainName}
		}

		return nil
	})
}

// GetDomainACLIPRule retrieves a single ACL IP rule for a domain
func (s *Service) GetDomainACLIPRule(domainName string, ruleID string) (ACLIPRule, error) {
	body, err := s.getDomainACLIPRuleResponseBody(domainName, ruleID)

	return body.Data, err
}

func (s *Service) getDomainACLIPRuleResponseBody(domainName string, ruleID string) (*connection.APIResponseBodyData[ACLIPRule], error) {
	body := &connection.APIResponseBodyData[ACLIPRule]{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}
	if ruleID == "" {
		return body, fmt.Errorf("invalid rule ID")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ddosx/v1/domains/%s/acls/ips/%s", domainName, ruleID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ACLIPRuleNotFoundError{ID: ruleID}
		}

		return nil
	})
}

// CreateDomainACLIPRule creates an ACL IP rule
func (s *Service) CreateDomainACLIPRule(domainName string, req CreateACLIPRuleRequest) (string, error) {
	body, err := s.createDomainACLIPRuleResponseBody(domainName, req)

	return body.Data.ID, err
}

func (s *Service) createDomainACLIPRuleResponseBody(domainName string, req CreateACLIPRuleRequest) (*connection.APIResponseBodyData[ACLIPRule], error) {
	body := &connection.APIResponseBodyData[ACLIPRule]{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}

	response, err := s.connection.Post(fmt.Sprintf("/ddosx/v1/domains/%s/acls/ips", domainName), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainNotFoundError{Name: domainName}
		}

		return nil
	})
}

// PatchDomainACLIPRule patches an ACL IP rule
func (s *Service) PatchDomainACLIPRule(domainName string, ruleID string, req PatchACLIPRuleRequest) error {
	_, err := s.patchDomainACLIPRuleResponseBody(domainName, ruleID, req)

	return err
}

func (s *Service) patchDomainACLIPRuleResponseBody(domainName string, ruleID string, req PatchACLIPRuleRequest) (*connection.APIResponseBodyData[ACLIPRule], error) {
	body := &connection.APIResponseBodyData[ACLIPRule]{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}
	if ruleID == "" {
		return body, fmt.Errorf("invalid rule ID")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ddosx/v1/domains/%s/acls/ips/%s", domainName, ruleID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ACLIPRuleNotFoundError{ID: ruleID}
		}

		return nil
	})
}

// DeleteDomainACLIPRule deletes an ACL IP rule
func (s *Service) DeleteDomainACLIPRule(domainName string, ruleID string) error {
	_, err := s.deleteDomainACLIPRuleResponseBody(domainName, ruleID)

	return err
}

func (s *Service) deleteDomainACLIPRuleResponseBody(domainName string, ruleID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}
	if ruleID == "" {
		return body, fmt.Errorf("invalid rule ID")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ddosx/v1/domains/%s/acls/ips/%s", domainName, ruleID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ACLIPRuleNotFoundError{ID: ruleID}
		}

		return nil
	})
}

// DownloadDomainVerificationFile downloads the verification file for a domain, returning
// the file contents, file name and an error
func (s *Service) DownloadDomainVerificationFile(domainName string) (content string, filename string, err error) {
	stream, filename, err := s.DownloadDomainVerificationFileStream(domainName)
	if err != nil {
		return "", "", err
	}

	defer stream.Close()

	bodyBytes, err := ioutil.ReadAll(stream)
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
	_, err := s.verifyDomainDNSResponseBody(domainName)

	return err
}

func (s *Service) verifyDomainDNSResponseBody(domainName string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}

	response, err := s.connection.Post(fmt.Sprintf("/ddosx/v1/domains/%s/verify/dns", domainName), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainNotFoundError{Name: domainName}
		}
		if response.StatusCode == 400 {
			return &DomainAlreadyVerifiedError{Name: domainName}
		}

		return nil
	})
}

// VerifyDomainFileUpload verifies a domain via file-upload method
func (s *Service) VerifyDomainFileUpload(domainName string) error {
	_, err := s.verifyDomainFileUploadResponseBody(domainName)

	return err
}

func (s *Service) verifyDomainFileUploadResponseBody(domainName string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}

	response, err := s.connection.Post(fmt.Sprintf("/ddosx/v1/domains/%s/verify/file-upload", domainName), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainNotFoundError{Name: domainName}
		}
		if response.StatusCode == 400 {
			return &DomainAlreadyVerifiedError{Name: domainName}
		}

		return nil
	})
}

// AddDomainCDNConfiguration adds CDN configuration to a domain
func (s *Service) AddDomainCDNConfiguration(domainName string) error {
	_, err := s.addDomainCDNConfigurationResponseBody(domainName)

	return err
}

func (s *Service) addDomainCDNConfigurationResponseBody(domainName string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}

	response, err := s.connection.Post(fmt.Sprintf("/ddosx/v1/domains/%s/cdn", domainName), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainNotFoundError{Name: domainName}
		}

		return nil
	})
}

// DeleteDomainCDNConfiguration removes CDN configuration from a domain
func (s *Service) DeleteDomainCDNConfiguration(domainName string) error {
	_, err := s.deleteDomainCDNConfigurationResponseBody(domainName)

	return err
}

func (s *Service) deleteDomainCDNConfigurationResponseBody(domainName string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ddosx/v1/domains/%s/cdn", domainName), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainCDNConfigurationNotFoundError{DomainName: domainName}
		}

		return nil
	})
}

// CreateDomainCDNRule creates a CDN rule
func (s *Service) CreateDomainCDNRule(domainName string, req CreateCDNRuleRequest) (string, error) {
	body, err := s.createDomainCDNRuleResponseBody(domainName, req)

	return body.Data.ID, err
}

func (s *Service) createDomainCDNRuleResponseBody(domainName string, req CreateCDNRuleRequest) (*connection.APIResponseBodyData[CDNRule], error) {
	body := &connection.APIResponseBodyData[CDNRule]{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}

	response, err := s.connection.Post(fmt.Sprintf("/ddosx/v1/domains/%s/cdn/rules", domainName), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainCDNConfigurationNotFoundError{DomainName: domainName}
		}

		return nil
	})
}

// GetDomainCDNRules retrieves a list of rules
func (s *Service) GetDomainCDNRules(domainName string, parameters connection.APIRequestParameters) ([]CDNRule, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[CDNRule], error) {
		return s.GetDomainCDNRulesPaginated(domainName, p)
	}, parameters)
}

// GetDomainCDNRulesPaginated retrieves a paginated list of domains
func (s *Service) GetDomainCDNRulesPaginated(domainName string, parameters connection.APIRequestParameters) (*connection.Paginated[CDNRule], error) {
	body, err := s.getDomainCDNRulesPaginatedResponseBody(domainName, parameters)
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[CDNRule], error) {
		return s.GetDomainCDNRulesPaginated(domainName, p)
	}), err
}

func (s *Service) getDomainCDNRulesPaginatedResponseBody(domainName string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]CDNRule], error) {
	body := &connection.APIResponseBodyData[[]CDNRule]{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ddosx/v1/domains/%s/cdn/rules", domainName), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainCDNConfigurationNotFoundError{DomainName: domainName}
		}

		return nil
	})
}

// GetDomainCDNRule retrieves a CDN rule
func (s *Service) GetDomainCDNRule(domainName string, ruleID string) (CDNRule, error) {
	body, err := s.getDomainCDNRuleResponseBody(domainName, ruleID)

	return body.Data, err
}

func (s *Service) getDomainCDNRuleResponseBody(domainName string, ruleID string) (*connection.APIResponseBodyData[CDNRule], error) {
	body := &connection.APIResponseBodyData[CDNRule]{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}
	if ruleID == "" {
		return body, fmt.Errorf("invalid rule ID")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ddosx/v1/domains/%s/cdn/rules/%s", domainName, ruleID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &CDNRuleNotFoundError{ID: ruleID}
		}

		return nil
	})
}

// PatchDomainCDNRule patches a CDN rule
func (s *Service) PatchDomainCDNRule(domainName string, ruleID string, req PatchCDNRuleRequest) error {
	_, err := s.patchDomainCDNRuleResponseBody(domainName, ruleID, req)

	return err
}

func (s *Service) patchDomainCDNRuleResponseBody(domainName string, ruleID string, req PatchCDNRuleRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}
	if ruleID == "" {
		return body, fmt.Errorf("invalid rule ID")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ddosx/v1/domains/%s/cdn/rules/%s", domainName, ruleID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &CDNRuleNotFoundError{ID: ruleID}
		}

		return nil
	})
}

// DeleteDomainCDNRule removes a CDN rule
func (s *Service) DeleteDomainCDNRule(domainName string, ruleID string) error {
	_, err := s.deleteDomainCDNRuleResponseBody(domainName, ruleID)

	return err
}

func (s *Service) deleteDomainCDNRuleResponseBody(domainName string, ruleID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}
	if ruleID == "" {
		return body, fmt.Errorf("invalid rule ID")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ddosx/v1/domains/%s/cdn/rules/%s", domainName, ruleID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &CDNRuleNotFoundError{ID: ruleID}
		}

		return nil
	})
}

// PurgeDomainCDN purges cached content
func (s *Service) PurgeDomainCDN(domainName string, req PurgeCDNRequest) error {
	_, err := s.purgeDomainCDNRuleResponseBody(domainName, req)

	return err
}

func (s *Service) purgeDomainCDNRuleResponseBody(domainName string, req PurgeCDNRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}

	response, err := s.connection.Post(fmt.Sprintf("/ddosx/v1/domains/%s/cdn/purge", domainName), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainCDNConfigurationNotFoundError{DomainName: domainName}
		}

		return nil
	})
}

// GetDomainHSTSConfiguration retrieves the HSTS configuration for a domain
func (s *Service) GetDomainHSTSConfiguration(domainName string) (HSTSConfiguration, error) {
	body, err := s.getDomainHSTSConfigurationResponseBody(domainName)

	return body.Data, err
}

func (s *Service) getDomainHSTSConfigurationResponseBody(domainName string) (*connection.APIResponseBodyData[HSTSConfiguration], error) {
	body := &connection.APIResponseBodyData[HSTSConfiguration]{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ddosx/v1/domains/%s/hsts", domainName), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainHSTSConfigurationNotFoundError{DomainName: domainName}
		}

		return nil
	})
}

// AddDomainHSTSConfiguration adds HSTS headers to a domain
func (s *Service) AddDomainHSTSConfiguration(domainName string) error {
	_, err := s.addDomainHSTSConfigurationResponseBody(domainName)

	return err
}

func (s *Service) addDomainHSTSConfigurationResponseBody(domainName string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}

	response, err := s.connection.Post(fmt.Sprintf("/ddosx/v1/domains/%s/hsts", domainName), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainNotFoundError{Name: domainName}
		}

		return nil
	})
}

// DeleteDomainHSTSConfiguration removes HSTS headers to a domain
func (s *Service) DeleteDomainHSTSConfiguration(domainName string) error {
	_, err := s.deleteDomainHSTSConfigurationResponseBody(domainName)

	return err
}

func (s *Service) deleteDomainHSTSConfigurationResponseBody(domainName string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ddosx/v1/domains/%s/hsts", domainName), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainHSTSConfigurationNotFoundError{DomainName: domainName}
		}

		return nil
	})
}

// CreateDomainHSTSRule creates a HSTS rule
func (s *Service) CreateDomainHSTSRule(domainName string, req CreateHSTSRuleRequest) (string, error) {
	body, err := s.createDomainHSTSRuleResponseBody(domainName, req)

	return body.Data.ID, err
}

func (s *Service) createDomainHSTSRuleResponseBody(domainName string, req CreateHSTSRuleRequest) (*connection.APIResponseBodyData[HSTSRule], error) {
	body := &connection.APIResponseBodyData[HSTSRule]{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}

	response, err := s.connection.Post(fmt.Sprintf("/ddosx/v1/domains/%s/hsts/rules", domainName), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainHSTSConfigurationNotFoundError{DomainName: domainName}
		}

		return nil
	})
}

// GetDomainHSTSRules retrieves a list of rules
func (s *Service) GetDomainHSTSRules(domainName string, parameters connection.APIRequestParameters) ([]HSTSRule, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[HSTSRule], error) {
		return s.GetDomainHSTSRulesPaginated(domainName, p)
	}, parameters)
}

// GetDomainHSTSRulesPaginated retrieves a paginated list of domains
func (s *Service) GetDomainHSTSRulesPaginated(domainName string, parameters connection.APIRequestParameters) (*connection.Paginated[HSTSRule], error) {
	body, err := s.getDomainHSTSRulesPaginatedResponseBody(domainName, parameters)
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[HSTSRule], error) {
		return s.GetDomainHSTSRulesPaginated(domainName, p)
	}), err
}

func (s *Service) getDomainHSTSRulesPaginatedResponseBody(domainName string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]HSTSRule], error) {
	body := &connection.APIResponseBodyData[[]HSTSRule]{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ddosx/v1/domains/%s/hsts/rules", domainName), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainHSTSConfigurationNotFoundError{DomainName: domainName}
		}

		return nil
	})
}

// GetDomainHSTSRule retrieves a HSTS rule
func (s *Service) GetDomainHSTSRule(domainName string, ruleID string) (HSTSRule, error) {
	body, err := s.getDomainHSTSRuleResponseBody(domainName, ruleID)

	return body.Data, err
}

func (s *Service) getDomainHSTSRuleResponseBody(domainName string, ruleID string) (*connection.APIResponseBodyData[HSTSRule], error) {
	body := &connection.APIResponseBodyData[HSTSRule]{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}
	if ruleID == "" {
		return body, fmt.Errorf("invalid rule ID")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ddosx/v1/domains/%s/hsts/rules/%s", domainName, ruleID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &HSTSRuleNotFoundError{ID: ruleID}
		}

		return nil
	})
}

// PatchDomainHSTSRule patches a HSTS rule
func (s *Service) PatchDomainHSTSRule(domainName string, ruleID string, req PatchHSTSRuleRequest) error {
	_, err := s.patchDomainHSTSRuleResponseBody(domainName, ruleID, req)

	return err
}

func (s *Service) patchDomainHSTSRuleResponseBody(domainName string, ruleID string, req PatchHSTSRuleRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}
	if ruleID == "" {
		return body, fmt.Errorf("invalid rule ID")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ddosx/v1/domains/%s/hsts/rules/%s", domainName, ruleID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &HSTSRuleNotFoundError{ID: ruleID}
		}

		return nil
	})
}

// DeleteDomainHSTSRule removes a HSTS rule
func (s *Service) DeleteDomainHSTSRule(domainName string, ruleID string) error {
	_, err := s.deleteDomainHSTSRuleResponseBody(domainName, ruleID)

	return err
}

func (s *Service) deleteDomainHSTSRuleResponseBody(domainName string, ruleID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}
	if ruleID == "" {
		return body, fmt.Errorf("invalid rule ID")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ddosx/v1/domains/%s/hsts/rules/%s", domainName, ruleID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &HSTSRuleNotFoundError{ID: ruleID}
		}

		return nil
	})
}

// ActivateDomainDNSRouting activates DNS routing for a domain
func (s *Service) ActivateDomainDNSRouting(domainName string) error {
	_, err := s.activateDomainDNSRoutingResponseBody(domainName)

	return err
}

func (s *Service) activateDomainDNSRoutingResponseBody(domainName string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}

	response, err := s.connection.Post(fmt.Sprintf("/ddosx/v1/domains/%s/dns/active", domainName), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainNotFoundError{Name: domainName}
		}

		return nil
	})
}

// DeactivateDomainDNSRouting deactivates DNS routing for a domain
func (s *Service) DeactivateDomainDNSRouting(domainName string) error {
	_, err := s.deactivateDomainDNSRoutingResponseBody(domainName)

	return err
}

func (s *Service) deactivateDomainDNSRoutingResponseBody(domainName string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ddosx/v1/domains/%s/dns/active", domainName), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DomainNotFoundError{Name: domainName}
		}

		return nil
	})
}
