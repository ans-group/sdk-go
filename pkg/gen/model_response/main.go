package main

import (
	"github.com/ukfast/sdk-go/pkg/gen/helper"
)

var modelTemplate = `
// Get{{.TypeName}}ArrayResponseBody represents an API response body containing []{{.TypeName}} data
type Get{{.TypeName}}ArrayResponseBody struct {
	connection.APIResponseBody

	Data []{{.TypeName}} ` + "`json:\"data\"`" + `
}

// Get{{.TypeName}}ResponseBody represents an API response body containing {{.TypeName}} data
type Get{{.TypeName}}ResponseBody struct {
	connection.APIResponseBody

	Data {{.TypeName}} ` + "`json:\"data\"`" + `
}
`

func main() {
	helper.TypeTemplate("github.com/ukfast/sdk-go/pkg/gen/response", modelTemplate)
}
