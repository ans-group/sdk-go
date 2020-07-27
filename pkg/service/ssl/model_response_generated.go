// Code generated by github.com/ukfast/sdk-go/pkg/gen/response. DO NOT EDIT.
package ssl

import "github.com/ukfast/sdk-go/pkg/connection"

// GetCertificateSliceResponseBody represents an API response body containing []Certificate data
type GetCertificateSliceResponseBody struct {
	connection.APIResponseBody

	Data []Certificate `json:"data"`
}

// GetCertificateResponseBody represents an API response body containing Certificate data
type GetCertificateResponseBody struct {
	connection.APIResponseBody

	Data Certificate `json:"data"`
}

// GetCertificateContentSliceResponseBody represents an API response body containing []CertificateContent data
type GetCertificateContentSliceResponseBody struct {
	connection.APIResponseBody

	Data []CertificateContent `json:"data"`
}

// GetCertificateContentResponseBody represents an API response body containing CertificateContent data
type GetCertificateContentResponseBody struct {
	connection.APIResponseBody

	Data CertificateContent `json:"data"`
}

// GetCertificatePrivateKeySliceResponseBody represents an API response body containing []CertificatePrivateKey data
type GetCertificatePrivateKeySliceResponseBody struct {
	connection.APIResponseBody

	Data []CertificatePrivateKey `json:"data"`
}

// GetCertificatePrivateKeyResponseBody represents an API response body containing CertificatePrivateKey data
type GetCertificatePrivateKeyResponseBody struct {
	connection.APIResponseBody

	Data CertificatePrivateKey `json:"data"`
}

// GetCertificateValidationSliceResponseBody represents an API response body containing []CertificateValidation data
type GetCertificateValidationSliceResponseBody struct {
	connection.APIResponseBody

	Data []CertificateValidation `json:"data"`
}

// GetCertificateValidationResponseBody represents an API response body containing CertificateValidation data
type GetCertificateValidationResponseBody struct {
	connection.APIResponseBody

	Data CertificateValidation `json:"data"`
}