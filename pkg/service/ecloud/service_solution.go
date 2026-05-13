package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) solutionRes() *resource.Resource[Solution, int] {
	return resource.NewIntResource[Solution](s.connection, "/ecloud/v1/solutions", "solution", func(id int) error {
		return &SolutionNotFoundError{ID: id}
	})
}

// GetSolutions retrieves a list of solutions
func (s *Service) GetSolutions(parameters connection.APIRequestParameters) ([]Solution, error) {
	return s.solutionRes().List(parameters)
}

// GetSolutionsPaginated retrieves a paginated list of solutions
func (s *Service) GetSolutionsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Solution], error) {
	return s.solutionRes().ListPaginated(parameters)
}

// GetSolution retrieves a single Solution by ID
func (s *Service) GetSolution(solutionID int) (Solution, error) {
	return s.solutionRes().Get(solutionID)
}

// PatchSolution patches an eCloud solution
func (s *Service) PatchSolution(solutionID int, patch PatchSolutionRequest) (int, error) {
	if solutionID < 1 {
		return 0, fmt.Errorf("invalid solution id")
	}
	body, err := connection.Patch[Solution](s.connection, fmt.Sprintf("/ecloud/v1/solutions/%d", solutionID), &patch, connection.NotFoundResponseHandler(&SolutionNotFoundError{ID: solutionID}))
	return body.Data.ID, err
}

// GetSolutionVirtualMachines retrieves a list of vms
func (s *Service) GetSolutionVirtualMachines(solutionID int, parameters connection.APIRequestParameters) ([]VirtualMachine, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[VirtualMachine], error) {
		return s.GetSolutionVirtualMachinesPaginated(solutionID, p)
	}, parameters)
}

// GetSolutionVirtualMachinesPaginated retrieves a paginated list of domains
func (s *Service) GetSolutionVirtualMachinesPaginated(solutionID int, parameters connection.APIRequestParameters) (*connection.Paginated[VirtualMachine], error) {
	if solutionID < 1 {
		return nil, fmt.Errorf("invalid solution id")
	}
	body, err := connection.Get[[]VirtualMachine](s.connection, fmt.Sprintf("/ecloud/v1/solutions/%d/vms", solutionID), parameters, connection.NotFoundResponseHandler(&SolutionNotFoundError{ID: solutionID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[VirtualMachine], error) {
		return s.GetSolutionVirtualMachinesPaginated(solutionID, p)
	}), err
}

// GetSolutionSites retrieves a list of sites
func (s *Service) GetSolutionSites(solutionID int, parameters connection.APIRequestParameters) ([]Site, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Site], error) {
		return s.GetSolutionSitesPaginated(solutionID, p)
	}, parameters)
}

// GetSolutionSitesPaginated retrieves a paginated list of domains
func (s *Service) GetSolutionSitesPaginated(solutionID int, parameters connection.APIRequestParameters) (*connection.Paginated[Site], error) {
	if solutionID < 1 {
		return nil, fmt.Errorf("invalid solution id")
	}
	body, err := connection.Get[[]Site](s.connection, fmt.Sprintf("/ecloud/v1/solutions/%d/sites", solutionID), parameters, connection.NotFoundResponseHandler(&SolutionNotFoundError{ID: solutionID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Site], error) {
		return s.GetSolutionSitesPaginated(solutionID, p)
	}), err
}

// GetSolutionDatastores retrieves a list of datastores
func (s *Service) GetSolutionDatastores(solutionID int, parameters connection.APIRequestParameters) ([]Datastore, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Datastore], error) {
		return s.GetSolutionDatastoresPaginated(solutionID, p)
	}, parameters)
}

// GetSolutionDatastoresPaginated retrieves a paginated list of domains
func (s *Service) GetSolutionDatastoresPaginated(solutionID int, parameters connection.APIRequestParameters) (*connection.Paginated[Datastore], error) {
	if solutionID < 1 {
		return nil, fmt.Errorf("invalid solution id")
	}
	body, err := connection.Get[[]Datastore](s.connection, fmt.Sprintf("/ecloud/v1/solutions/%d/datastores", solutionID), parameters, connection.NotFoundResponseHandler(&SolutionNotFoundError{ID: solutionID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Datastore], error) {
		return s.GetSolutionDatastoresPaginated(solutionID, p)
	}), err
}

// GetSolutionHosts retrieves a list of hosts
func (s *Service) GetSolutionHosts(solutionID int, parameters connection.APIRequestParameters) ([]V1Host, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[V1Host], error) {
		return s.GetSolutionHostsPaginated(solutionID, p)
	}, parameters)
}

// GetSolutionHostsPaginated retrieves a paginated list of domains
func (s *Service) GetSolutionHostsPaginated(solutionID int, parameters connection.APIRequestParameters) (*connection.Paginated[V1Host], error) {
	if solutionID < 1 {
		return nil, fmt.Errorf("invalid solution id")
	}
	body, err := connection.Get[[]V1Host](s.connection, fmt.Sprintf("/ecloud/v1/solutions/%d/hosts", solutionID), parameters, connection.NotFoundResponseHandler(&SolutionNotFoundError{ID: solutionID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[V1Host], error) {
		return s.GetSolutionHostsPaginated(solutionID, p)
	}), err
}

// GetSolutionNetworks retrieves a list of networks
func (s *Service) GetSolutionNetworks(solutionID int, parameters connection.APIRequestParameters) ([]V1Network, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[V1Network], error) {
		return s.GetSolutionNetworksPaginated(solutionID, p)
	}, parameters)
}

// GetSolutionNetworksPaginated retrieves a paginated list of domains
func (s *Service) GetSolutionNetworksPaginated(solutionID int, parameters connection.APIRequestParameters) (*connection.Paginated[V1Network], error) {
	if solutionID < 1 {
		return nil, fmt.Errorf("invalid solution id")
	}
	body, err := connection.Get[[]V1Network](s.connection, fmt.Sprintf("/ecloud/v1/solutions/%d/networks", solutionID), parameters, connection.NotFoundResponseHandler(&SolutionNotFoundError{ID: solutionID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[V1Network], error) {
		return s.GetSolutionNetworksPaginated(solutionID, p)
	}), err
}

// GetSolutionFirewalls retrieves a list of firewalls
func (s *Service) GetSolutionFirewalls(solutionID int, parameters connection.APIRequestParameters) ([]Firewall, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Firewall], error) {
		return s.GetSolutionFirewallsPaginated(solutionID, p)
	}, parameters)
}

// GetSolutionFirewallsPaginated retrieves a paginated list of domains
func (s *Service) GetSolutionFirewallsPaginated(solutionID int, parameters connection.APIRequestParameters) (*connection.Paginated[Firewall], error) {
	if solutionID < 1 {
		return nil, fmt.Errorf("invalid solution id")
	}
	body, err := connection.Get[[]Firewall](s.connection, fmt.Sprintf("/ecloud/v1/solutions/%d/firewalls", solutionID), parameters, connection.NotFoundResponseHandler(&SolutionNotFoundError{ID: solutionID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Firewall], error) {
		return s.GetSolutionFirewallsPaginated(solutionID, p)
	}), err
}

// GetSolutionTemplates retrieves a list of templates
func (s *Service) GetSolutionTemplates(solutionID int, parameters connection.APIRequestParameters) ([]Template, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Template], error) {
		return s.GetSolutionTemplatesPaginated(solutionID, p)
	}, parameters)
}

// GetSolutionTemplatesPaginated retrieves a paginated list of domains
func (s *Service) GetSolutionTemplatesPaginated(solutionID int, parameters connection.APIRequestParameters) (*connection.Paginated[Template], error) {
	if solutionID < 1 {
		return nil, fmt.Errorf("invalid solution id")
	}
	body, err := connection.Get[[]Template](s.connection, fmt.Sprintf("/ecloud/v1/solutions/%d/templates", solutionID), parameters, connection.NotFoundResponseHandler(&SolutionNotFoundError{ID: solutionID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Template], error) {
		return s.GetSolutionTemplatesPaginated(solutionID, p)
	}), err
}

// GetSolutionTemplate retrieves a single solution template by name
func (s *Service) GetSolutionTemplate(solutionID int, templateName string) (Template, error) {
	if solutionID < 1 {
		return Template{}, fmt.Errorf("invalid solution id")
	}
	if templateName == "" {
		return Template{}, fmt.Errorf("invalid template name")
	}
	body, err := connection.Get[Template](s.connection, fmt.Sprintf("/ecloud/v1/solutions/%d/templates/%s", solutionID, templateName), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&TemplateNotFoundError{Name: templateName}))
	return body.Data, err
}

// RenameSolutionTemplate renames a solution template
func (s *Service) RenameSolutionTemplate(solutionID int, templateName string, req RenameTemplateRequest) error {
	if solutionID < 1 {
		return fmt.Errorf("invalid solution id")
	}
	if templateName == "" {
		return fmt.Errorf("invalid template name")
	}
	return connection.PostRaw(s.connection, fmt.Sprintf("/ecloud/v1/solutions/%d/templates/%s/move", solutionID, templateName), &req, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&TemplateNotFoundError{Name: templateName}))
}

// DeleteSolutionTemplate removes a solution template
func (s *Service) DeleteSolutionTemplate(solutionID int, templateName string) error {
	if solutionID < 1 {
		return fmt.Errorf("invalid solution id")
	}
	if templateName == "" {
		return fmt.Errorf("invalid template name")
	}
	return connection.DeleteRaw(s.connection, fmt.Sprintf("/ecloud/v1/solutions/%d/templates/%s", solutionID, templateName), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&TemplateNotFoundError{Name: templateName}))
}

// GetSolutionTags retrieves a list of tags
func (s *Service) GetSolutionTags(solutionID int, parameters connection.APIRequestParameters) ([]TagV1, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[TagV1], error) {
		return s.GetSolutionTagsPaginated(solutionID, p)
	}, parameters)
}

// GetSolutionTagsPaginated retrieves a paginated list of v1 solution tags
func (s *Service) GetSolutionTagsPaginated(solutionID int, parameters connection.APIRequestParameters) (*connection.Paginated[TagV1], error) {
	if solutionID < 1 {
		return nil, fmt.Errorf("invalid solution id")
	}
	body, err := connection.Get[[]TagV1](s.connection, fmt.Sprintf("/ecloud/v1/solutions/%d/tags", solutionID), parameters, connection.NotFoundResponseHandler(&SolutionNotFoundError{ID: solutionID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[TagV1], error) {
		return s.GetSolutionTagsPaginated(solutionID, p)
	}), err
}

// GetSolutionTag retrieves a single solution v1 tag by key
func (s *Service) GetSolutionTag(solutionID int, tagKey string) (TagV1, error) {
	if solutionID < 1 {
		return TagV1{}, fmt.Errorf("invalid solution id")
	}
	if tagKey == "" {
		return TagV1{}, fmt.Errorf("invalid tag key")
	}
	body, err := connection.Get[TagV1](s.connection, fmt.Sprintf("/ecloud/v1/solutions/%d/tags/%s", solutionID, tagKey), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&TagV1NotFoundError{Key: tagKey}))
	return body.Data, err
}

// CreateSolutionTag creates a new solution v1 tag
func (s *Service) CreateSolutionTag(solutionID int, req CreateTagV1Request) error {
	if solutionID < 1 {
		return fmt.Errorf("invalid solution id")
	}
	return connection.PostRaw(s.connection, fmt.Sprintf("/ecloud/v1/solutions/%d/tags", solutionID), &req, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&SolutionNotFoundError{ID: solutionID}))
}

// PatchSolutionTag patches an eCloud solution v1 tag
func (s *Service) PatchSolutionTag(solutionID int, tagKey string, patch PatchTagV1Request) error {
	if solutionID < 1 {
		return fmt.Errorf("invalid solution id")
	}
	if tagKey == "" {
		return fmt.Errorf("invalid tag key")
	}
	return connection.PatchRaw(s.connection, fmt.Sprintf("/ecloud/v1/solutions/%d/tags/%s", solutionID, tagKey), &patch, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&TagV1NotFoundError{Key: tagKey}))
}

// DeleteSolutionTag removes a solution v1 tag
func (s *Service) DeleteSolutionTag(solutionID int, tagKey string) error {
	if solutionID < 1 {
		return fmt.Errorf("invalid solution id")
	}
	if tagKey == "" {
		return fmt.Errorf("invalid tag key")
	}
	return connection.DeleteRaw(s.connection, fmt.Sprintf("/ecloud/v1/solutions/%d/tags/%s", solutionID, tagKey), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&TagV1NotFoundError{Key: tagKey}))
}
