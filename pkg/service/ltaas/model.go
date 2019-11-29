//go:generate go run ../../gen/model_paginated_gen.go -package ltaas -typename Domain -destination model_paginated.go

package ltaas

import (
	"github.com/ukfast/sdk-go/pkg/connection"
)

type DomainVerificationMethod string

func (s DomainVerificationMethod) String() string {
	return string(s)
}

const (
	DomainVerificationMethodDNS        DomainVerificationMethod = "DNS"
	DomainVerificationMethodFileUpload DomainVerificationMethod = "File upload"
)

type DomainStatus string

func (s DomainStatus) String() string {
	return string(s)
}

const (
	DomainStatusVerified    DomainStatus = "Verified"
	DomainStatusNotVerified DomainStatus = "Not verified"
)

// Domain represents an LTaaS domain
type Domain struct {
	ID                 string                   `json:"id"`
	Name               string                   `json:"name"`
	VerificationMethod DomainVerificationMethod `json:"verification_method"`
	VerifyHash         string                   `json:"verify_hash"`
	Status             DomainStatus             `json:"status"`
	CreatedAt          connection.DateTime      `json:"created_at"`
	UpdatedAt          connection.DateTime      `json:"updated_at"`
}
