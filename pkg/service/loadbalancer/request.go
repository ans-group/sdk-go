package loadbalancer

// CreateConfigurationRequest represents a request to create a configuration
type CreateConfigurationRequest struct {
	Name *string `json:"name,omitempty"`
}

// PatchConfigurationRequest represents a request to patch a configuration
type PatchConfigurationRequest struct {
	Name *string `json:"name,omitempty"`
}
