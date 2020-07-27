package main

import (
	"github.com/ukfast/sdk-go/pkg/gen/helper"
)

var modelTemplate = `
// Paginated{{.TypeName}} represents a paginated collection of {{.TypeName}}
type Paginated{{.TypeName}} struct {
	*connection.PaginatedBase

	Items []{{.TypeName}}
}

// NewPaginated{{.TypeName}} returns a pointer to an initialized Paginated{{.TypeName}} struct
func NewPaginated{{.TypeName}}(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []{{.TypeName}}) *Paginated{{.TypeName}} {
	return &Paginated{{.TypeName}}{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}
`

func main() {
	helper.TypeTemplate("github.com/ukfast/sdk-go/pkg/gen/model_paginated", modelTemplate)
}
