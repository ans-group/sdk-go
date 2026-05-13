package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) instanceRes() *resource.Resource[Instance, string] {
	return resource.NewStringResource[Instance](s.connection, "/ecloud/v2/instances", "instance", func(id string) error {
		return &InstanceNotFoundError{ID: id}
	})
}

// GetInstances retrieves a list of instances
func (s *Service) GetInstances(parameters connection.APIRequestParameters) ([]Instance, error) {
	return s.instanceRes().List(parameters)
}

// GetInstancesPaginated retrieves a paginated list of instances
func (s *Service) GetInstancesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Instance], error) {
	return s.instanceRes().ListPaginated(parameters)
}

// GetInstance retrieves a single instance by id
func (s *Service) GetInstance(instanceID string) (Instance, error) {
	return s.instanceRes().Get(instanceID)
}

// CreateInstance creates a new instance
func (s *Service) CreateInstance(req CreateInstanceRequest) (string, error) {
	data, err := s.instanceRes().Create(&req)
	return data.ID, err
}

// PatchInstance updates an instance
func (s *Service) PatchInstance(instanceID string, req PatchInstanceRequest) error {
	return s.instanceRes().Patch(instanceID, &req)
}

// DeleteInstance removes an instance
func (s *Service) DeleteInstance(instanceID string) error {
	return s.instanceRes().Delete(instanceID)
}

// LockInstance locks an instance from update/removal
func (s *Service) LockInstance(instanceID string) error {
	if instanceID == "" {
		return fmt.Errorf("invalid instance id")
	}
	return connection.PutRaw(s.connection, fmt.Sprintf("/ecloud/v2/instances/%s/lock", instanceID), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&InstanceNotFoundError{ID: instanceID}))
}

// UnlockInstance unlocks an instance
func (s *Service) UnlockInstance(instanceID string) error {
	if instanceID == "" {
		return fmt.Errorf("invalid instance id")
	}
	return connection.PutRaw(s.connection, fmt.Sprintf("/ecloud/v2/instances/%s/unlock", instanceID), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&InstanceNotFoundError{ID: instanceID}))
}

// PowerOnInstance powers on an instance
func (s *Service) PowerOnInstance(instanceID string) (string, error) {
	if instanceID == "" {
		return "", fmt.Errorf("invalid instance id")
	}
	body, err := connection.Put[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/instances/%s/power-on", instanceID), nil, connection.NotFoundResponseHandler(&InstanceNotFoundError{ID: instanceID}))
	return body.Data.TaskID, err
}

// PowerOffInstance powers off an instance
func (s *Service) PowerOffInstance(instanceID string) (string, error) {
	if instanceID == "" {
		return "", fmt.Errorf("invalid instance id")
	}
	body, err := connection.Put[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/instances/%s/power-off", instanceID), nil, connection.NotFoundResponseHandler(&InstanceNotFoundError{ID: instanceID}))
	return body.Data.TaskID, err
}

// PowerResetInstance resets an instance
func (s *Service) PowerResetInstance(instanceID string) (string, error) {
	if instanceID == "" {
		return "", fmt.Errorf("invalid instance id")
	}
	body, err := connection.Put[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/instances/%s/power-reset", instanceID), nil, connection.NotFoundResponseHandler(&InstanceNotFoundError{ID: instanceID}))
	return body.Data.TaskID, err
}

// PowerShutdownInstance shuts down an instance
func (s *Service) PowerShutdownInstance(instanceID string) (string, error) {
	if instanceID == "" {
		return "", fmt.Errorf("invalid instance id")
	}
	body, err := connection.Put[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/instances/%s/power-shutdown", instanceID), nil, connection.NotFoundResponseHandler(&InstanceNotFoundError{ID: instanceID}))
	return body.Data.TaskID, err
}

// PowerRestartInstance restarts an instance
func (s *Service) PowerRestartInstance(instanceID string) (string, error) {
	if instanceID == "" {
		return "", fmt.Errorf("invalid instance id")
	}
	body, err := connection.Put[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/instances/%s/power-restart", instanceID), nil, connection.NotFoundResponseHandler(&InstanceNotFoundError{ID: instanceID}))
	return body.Data.TaskID, err
}

// MigrateInstance migrates an instance
func (s *Service) MigrateInstance(instanceID string, req MigrateInstanceRequest) (string, error) {
	if instanceID == "" {
		return "", fmt.Errorf("invalid instance id")
	}
	body, err := connection.Post[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/instances/%s/migrate", instanceID), &req, connection.NotFoundResponseHandler(&InstanceNotFoundError{ID: instanceID}))
	return body.Data.TaskID, err
}

func (s *Service) instanceVolumeRes() *resource.SubResourceList[Volume, string] {
	return resource.NewStringSubResourceList[Volume](s.connection,
		func(instanceID string) string { return fmt.Sprintf("/ecloud/v2/instances/%s/volumes", instanceID) },
		"instance", "id", func(instanceID string) error { return &InstanceNotFoundError{ID: instanceID} })
}

// GetInstanceVolumes retrieves a list of instance volumes
func (s *Service) GetInstanceVolumes(instanceID string, parameters connection.APIRequestParameters) ([]Volume, error) {
	return s.instanceVolumeRes().List(instanceID, parameters)
}

// GetInstanceVolumesPaginated retrieves a paginated list of instance volumes
func (s *Service) GetInstanceVolumesPaginated(instanceID string, parameters connection.APIRequestParameters) (*connection.Paginated[Volume], error) {
	return s.instanceVolumeRes().ListPaginated(instanceID, parameters)
}

func (s *Service) instanceCredentialRes() *resource.SubResourceList[Credential, string] {
	return resource.NewStringSubResourceList[Credential](s.connection,
		func(instanceID string) string { return fmt.Sprintf("/ecloud/v2/instances/%s/credentials", instanceID) },
		"instance", "id", func(instanceID string) error { return &InstanceNotFoundError{ID: instanceID} })
}

// GetInstanceCredentials retrieves a list of instance credentials
func (s *Service) GetInstanceCredentials(instanceID string, parameters connection.APIRequestParameters) ([]Credential, error) {
	return s.instanceCredentialRes().List(instanceID, parameters)
}

// GetInstanceCredentialsPaginated retrieves a paginated list of instance credentials
func (s *Service) GetInstanceCredentialsPaginated(instanceID string, parameters connection.APIRequestParameters) (*connection.Paginated[Credential], error) {
	return s.instanceCredentialRes().ListPaginated(instanceID, parameters)
}

func (s *Service) instanceNICRes() *resource.SubResourceList[NIC, string] {
	return resource.NewStringSubResourceList[NIC](s.connection,
		func(instanceID string) string { return fmt.Sprintf("/ecloud/v2/instances/%s/nics", instanceID) },
		"instance", "id", func(instanceID string) error { return &InstanceNotFoundError{ID: instanceID} })
}

// GetInstanceNICs retrieves a list of instance NICs
func (s *Service) GetInstanceNICs(instanceID string, parameters connection.APIRequestParameters) ([]NIC, error) {
	return s.instanceNICRes().List(instanceID, parameters)
}

// GetInstanceNICsPaginated retrieves a paginated list of instance NICs
func (s *Service) GetInstanceNICsPaginated(instanceID string, parameters connection.APIRequestParameters) (*connection.Paginated[NIC], error) {
	return s.instanceNICRes().ListPaginated(instanceID, parameters)
}

// CreateInstanceConsoleSession creates an instance console session
func (s *Service) CreateInstanceConsoleSession(instanceID string) (ConsoleSession, error) {
	if instanceID == "" {
		return ConsoleSession{}, fmt.Errorf("invalid instance id")
	}
	body, err := connection.Post[ConsoleSession](s.connection, fmt.Sprintf("/ecloud/v2/instances/%s/console-session", instanceID), nil, connection.NotFoundResponseHandler(&InstanceNotFoundError{ID: instanceID}))
	return body.Data, err
}

func (s *Service) instanceTasksRes() *resource.SubResourceList[Task, string] {
	return resource.NewStringSubResourceList[Task](s.connection,
		func(instanceID string) string { return fmt.Sprintf("/ecloud/v2/instances/%s/tasks", instanceID) },
		"instance", "id", func(instanceID string) error { return &InstanceNotFoundError{ID: instanceID} })
}

// GetInstanceTasks retrieves a list of Instance tasks
func (s *Service) GetInstanceTasks(instanceID string, parameters connection.APIRequestParameters) ([]Task, error) {
	return s.instanceTasksRes().List(instanceID, parameters)
}

// GetInstanceTasksPaginated retrieves a paginated list of Instance tasks
func (s *Service) GetInstanceTasksPaginated(instanceID string, parameters connection.APIRequestParameters) (*connection.Paginated[Task], error) {
	return s.instanceTasksRes().ListPaginated(instanceID, parameters)
}

// AttachInstanceVolume attaches a volume to an instance
func (s *Service) AttachInstanceVolume(instanceID string, req AttachDetachInstanceVolumeRequest) (string, error) {
	if instanceID == "" {
		return "", fmt.Errorf("invalid instance id")
	}
	body, err := connection.Post[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/instances/%s/volume-attach", instanceID), &req, connection.NotFoundResponseHandler(&InstanceNotFoundError{ID: instanceID}))
	return body.Data.TaskID, err
}

// DetachInstanceVolume detaches a volume from an instance
func (s *Service) DetachInstanceVolume(instanceID string, req AttachDetachInstanceVolumeRequest) (string, error) {
	if instanceID == "" {
		return "", fmt.Errorf("invalid instance id")
	}
	body, err := connection.Post[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/instances/%s/volume-detach", instanceID), &req, connection.NotFoundResponseHandler(&InstanceNotFoundError{ID: instanceID}))
	return body.Data.TaskID, err
}

func (s *Service) instanceFloatingIPRes() *resource.SubResourceList[FloatingIP, string] {
	return resource.NewStringSubResourceList[FloatingIP](s.connection,
		func(instanceID string) string { return fmt.Sprintf("/ecloud/v2/instances/%s/floating-ips", instanceID) },
		"instance", "id", func(instanceID string) error { return &InstanceNotFoundError{ID: instanceID} })
}

// GetInstanceFloatingIPs retrieves a list of instance fips
func (s *Service) GetInstanceFloatingIPs(instanceID string, parameters connection.APIRequestParameters) ([]FloatingIP, error) {
	return s.instanceFloatingIPRes().List(instanceID, parameters)
}

// GetInstanceFloatingIPsPaginated retrieves a paginated list of instance floating IPs
func (s *Service) GetInstanceFloatingIPsPaginated(instanceID string, parameters connection.APIRequestParameters) (*connection.Paginated[FloatingIP], error) {
	return s.instanceFloatingIPRes().ListPaginated(instanceID, parameters)
}

// CreateInstanceImage attaches a volume to an instance
func (s *Service) CreateInstanceImage(instanceID string, req CreateInstanceImageRequest) (TaskReference, error) {
	if instanceID == "" {
		return TaskReference{}, fmt.Errorf("invalid instance id")
	}
	body, err := connection.Post[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/instances/%s/create-image", instanceID), &req, connection.NotFoundResponseHandler(&InstanceNotFoundError{ID: instanceID}))
	return body.Data, err
}

// EncryptInstance encrypts an instance
func (s *Service) EncryptInstance(instanceID string) (string, error) {
	if instanceID == "" {
		return "", fmt.Errorf("invalid instance id")
	}
	body, err := connection.Put[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/instances/%s/encrypt", instanceID), nil, connection.NotFoundResponseHandler(&InstanceNotFoundError{ID: instanceID}))
	return body.Data.TaskID, err
}

// DecryptInstance decrypts an instance
func (s *Service) DecryptInstance(instanceID string) (string, error) {
	if instanceID == "" {
		return "", fmt.Errorf("invalid instance id")
	}
	body, err := connection.Put[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/instances/%s/decrypt", instanceID), nil, connection.NotFoundResponseHandler(&InstanceNotFoundError{ID: instanceID}))
	return body.Data.TaskID, err
}

func (s *Service) ExecuteInstanceScript(instanceID string, req ExecuteInstanceScriptRequest) (string, error) {
	if instanceID == "" {
		return "", fmt.Errorf("invalid instance id")
	}
	body, err := connection.Post[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/instances/%s/user-script", instanceID), &req, connection.NotFoundResponseHandler(&InstanceNotFoundError{ID: instanceID}))
	return body.Data.TaskID, err
}
