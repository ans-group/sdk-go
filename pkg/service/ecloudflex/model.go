package ecloudflex

import "github.com/ukfast/sdk-go/pkg/connection"

// Project represents an eCloud Flex project
type Project struct {
	ID        int                 `json:"id"`
	Name      string              `json:"name"`
	CreatedAt connection.DateTime `json:"created_at"`
}
