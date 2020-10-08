//go:generate go run ../../gen/model_response/main.go -package loadbalancer -source model.go -destination model_response_generated.go
//go:generate go run ../../gen/model_paginated/main.go -package loadbalancer -source model.go -destination model_paginated_generated.go

package loadbalancer

// Group represents a group
// +genie:model_response
// +genie:model_paginated
type Group struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Configuration represents a configuration
// +genie:model_response
// +genie:model_paginated
type Configuration struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
