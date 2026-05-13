package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetVirtualMachines retrieves a list of vms
func (s *Service) GetVirtualMachines(parameters connection.APIRequestParameters) ([]VirtualMachine, error) {
	return connection.InvokeRequestAll(s.GetVirtualMachinesPaginated, parameters)
}

// GetVirtualMachinesPaginated retrieves a paginated list of vms
func (s *Service) GetVirtualMachinesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[VirtualMachine], error) {
	body, err := connection.Get[[]VirtualMachine](s.connection, "/ecloud/v1/vms", parameters)
	return connection.NewPaginated(body, parameters, s.GetVirtualMachinesPaginated), err
}

// GetVirtualMachine retrieves a single virtual machine by ID
func (s *Service) GetVirtualMachine(vmID int) (VirtualMachine, error) {
	if vmID < 1 {
		return VirtualMachine{}, fmt.Errorf("invalid virtual machine id")
	}
	body, err := connection.Get[VirtualMachine](s.connection, fmt.Sprintf("/ecloud/v1/vms/%d", vmID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&VirtualMachineNotFoundError{ID: vmID}))
	return body.Data, err
}

// DeleteVirtualMachine removes a virtual machine
func (s *Service) DeleteVirtualMachine(vmID int) error {
	if vmID < 1 {
		return fmt.Errorf("invalid virtual machine id")
	}
	return connection.DeleteRaw(s.connection, fmt.Sprintf("/ecloud/v1/vms/%d", vmID), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&VirtualMachineNotFoundError{ID: vmID}))
}

// CreateVirtualMachine creates a new virtual machine
func (s *Service) CreateVirtualMachine(req CreateVirtualMachineRequest) (int, error) {
	body, err := connection.Post[VirtualMachine](s.connection, "/ecloud/v1/vms", &req)
	return body.Data.ID, err
}

// PatchVirtualMachine patches an eCloud virtual machine
func (s *Service) PatchVirtualMachine(vmID int, patch PatchVirtualMachineRequest) error {
	if vmID < 1 {
		return fmt.Errorf("invalid virtual machine id")
	}
	return connection.PatchRaw(s.connection, fmt.Sprintf("/ecloud/v1/vms/%d", vmID), &patch, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&VirtualMachineNotFoundError{ID: vmID}))
}

// CloneVirtualMachine clones a virtual machine
func (s *Service) CloneVirtualMachine(vmID int, req CloneVirtualMachineRequest) (int, error) {
	if vmID < 1 {
		return 0, fmt.Errorf("invalid virtual machine id")
	}
	body, err := connection.Post[VirtualMachine](s.connection, fmt.Sprintf("/ecloud/v1/vms/%d/clone", vmID), &req, connection.NotFoundResponseHandler(&VirtualMachineNotFoundError{ID: vmID}))
	return body.Data.ID, err
}

// PowerOnVirtualMachine powers on a virtual machine
func (s *Service) PowerOnVirtualMachine(vmID int) error {
	if vmID < 1 {
		return fmt.Errorf("invalid virtual machine id")
	}
	return connection.PutRaw(s.connection, fmt.Sprintf("/ecloud/v1/vms/%d/power-on", vmID), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&VirtualMachineNotFoundError{ID: vmID}))
}

// PowerOffVirtualMachine powers off a virtual machine
func (s *Service) PowerOffVirtualMachine(vmID int) error {
	if vmID < 1 {
		return fmt.Errorf("invalid virtual machine id")
	}
	return connection.PutRaw(s.connection, fmt.Sprintf("/ecloud/v1/vms/%d/power-off", vmID), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&VirtualMachineNotFoundError{ID: vmID}))
}

// PowerResetVirtualMachine resets a virtual machine (hard power off)
func (s *Service) PowerResetVirtualMachine(vmID int) error {
	if vmID < 1 {
		return fmt.Errorf("invalid virtual machine id")
	}
	return connection.PutRaw(s.connection, fmt.Sprintf("/ecloud/v1/vms/%d/power-reset", vmID), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&VirtualMachineNotFoundError{ID: vmID}))
}

// PowerShutdownVirtualMachine shuts down a virtual machine
func (s *Service) PowerShutdownVirtualMachine(vmID int) error {
	if vmID < 1 {
		return fmt.Errorf("invalid virtual machine id")
	}
	return connection.PutRaw(s.connection, fmt.Sprintf("/ecloud/v1/vms/%d/power-shutdown", vmID), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&VirtualMachineNotFoundError{ID: vmID}))
}

// PowerRestartVirtualMachine resets a virtual machine (graceful power off)
func (s *Service) PowerRestartVirtualMachine(vmID int) error {
	if vmID < 1 {
		return fmt.Errorf("invalid virtual machine id")
	}
	return connection.PutRaw(s.connection, fmt.Sprintf("/ecloud/v1/vms/%d/power-restart", vmID), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&VirtualMachineNotFoundError{ID: vmID}))
}

// CreateVirtualMachineTemplate creates a virtual machine template
func (s *Service) CreateVirtualMachineTemplate(vmID int, req CreateVirtualMachineTemplateRequest) error {
	if vmID < 1 {
		return fmt.Errorf("invalid virtual machine id")
	}
	return connection.PostRaw(s.connection, fmt.Sprintf("/ecloud/v1/vms/%d/clone-to-template", vmID), &req, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&VirtualMachineNotFoundError{ID: vmID}))
}

// GetVirtualMachineTags retrieves a list of tags
func (s *Service) GetVirtualMachineTags(vmID int, parameters connection.APIRequestParameters) ([]TagV1, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[TagV1], error) {
		return s.GetVirtualMachineTagsPaginated(vmID, p)
	}, parameters)
}

// GetVirtualMachineTagsPaginated retrieves a paginated list of v1 virtual machine tags
func (s *Service) GetVirtualMachineTagsPaginated(vmID int, parameters connection.APIRequestParameters) (*connection.Paginated[TagV1], error) {
	if vmID < 1 {
		return nil, fmt.Errorf("invalid virtual machine id")
	}
	body, err := connection.Get[[]TagV1](s.connection, fmt.Sprintf("/ecloud/v1/vms/%d/tags", vmID), parameters, connection.NotFoundResponseHandler(&VirtualMachineNotFoundError{ID: vmID}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[TagV1], error) {
		return s.GetVirtualMachineTagsPaginated(vmID, p)
	}), err
}

// GetVirtualMachineTag retrieves a single virtual machine v1 tag by key
func (s *Service) GetVirtualMachineTag(vmID int, tagKey string) (TagV1, error) {
	if vmID < 1 {
		return TagV1{}, fmt.Errorf("invalid virtual machine id")
	}
	if tagKey == "" {
		return TagV1{}, fmt.Errorf("invalid tag key")
	}
	body, err := connection.Get[TagV1](s.connection, fmt.Sprintf("/ecloud/v1/vms/%d/tags/%s", vmID, tagKey), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&TagV1NotFoundError{Key: tagKey}))
	return body.Data, err
}

// CreateVirtualMachineTag creates a new virtual machine v1 tag
func (s *Service) CreateVirtualMachineTag(vmID int, req CreateTagV1Request) error {
	if vmID < 1 {
		return fmt.Errorf("invalid virtual machine id")
	}
	return connection.PostRaw(s.connection, fmt.Sprintf("/ecloud/v1/vms/%d/tags", vmID), &req, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&VirtualMachineNotFoundError{ID: vmID}))
}

// PatchVirtualMachineTag patches an eCloud virtual machine v1 tag
func (s *Service) PatchVirtualMachineTag(vmID int, tagKey string, patch PatchTagV1Request) error {
	if vmID < 1 {
		return fmt.Errorf("invalid virtual machine id")
	}
	if tagKey == "" {
		return fmt.Errorf("invalid tag key")
	}
	return connection.PatchRaw(s.connection, fmt.Sprintf("/ecloud/v1/vms/%d/tags/%s", vmID, tagKey), &patch, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&TagV1NotFoundError{Key: tagKey}))
}

// DeleteVirtualMachineTag removes a virtual machine v1 tag
func (s *Service) DeleteVirtualMachineTag(vmID int, tagKey string) error {
	if vmID < 1 {
		return fmt.Errorf("invalid virtual machine id")
	}
	if tagKey == "" {
		return fmt.Errorf("invalid tag key")
	}
	return connection.DeleteRaw(s.connection, fmt.Sprintf("/ecloud/v1/vms/%d/tags/%s", vmID, tagKey), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&TagV1NotFoundError{Key: tagKey}))
}

// CreateVirtualMachineConsoleSession creates a virtual machine console session
func (s *Service) CreateVirtualMachineConsoleSession(vmID int) (ConsoleSession, error) {
	if vmID < 1 {
		return ConsoleSession{}, fmt.Errorf("invalid virtual machine id")
	}
	body, err := connection.Put[ConsoleSession](s.connection, fmt.Sprintf("/ecloud/v1/vms/%d/console-session", vmID), nil, connection.NotFoundResponseHandler(&VirtualMachineNotFoundError{ID: vmID}))
	return body.Data, err
}
