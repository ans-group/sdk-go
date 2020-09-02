package main

import (
	"fmt"

	"github.com/0x4c6565/genie"
	"github.com/dave/jennifer/jen"
	"github.com/ukfast/sdk-go/pkg/gen/helper"
)

type ModelResponseGenerator struct {
	*genie.BaseGenerator
}

func NewModelResponseGenerator() *ModelResponseGenerator {
	return &ModelResponseGenerator{
		BaseGenerator: genie.NewBaseGenerator("model_response"),
	}
}

func (g *ModelResponseGenerator) Generate(marker genie.Marker, typeName string, f *jen.File) error {
	f.ImportName("github.com/ukfast/sdk-go/pkg/connection", "connection")

	f.Comment(fmt.Sprintf("Get%sSliceResponseBody represents an API response body containing []%[1]s data", typeName))
	f.Type().Id("Get"+typeName+"SliceResponseBody").Struct(
		jen.Qual("github.com/ukfast/sdk-go/pkg/connection", "APIResponseBody"),
		jen.Id("Data").Index().Id(typeName).Tag(map[string]string{"json": "data"}),
	)

	f.Comment(fmt.Sprintf("Get%sResponseBody represents an API response body containing %[1]s data", typeName))
	f.Type().Id("Get"+typeName+"ResponseBody").Struct(
		jen.Qual("github.com/ukfast/sdk-go/pkg/connection", "APIResponseBody"),
		jen.Id("Data").Id(typeName).Tag(map[string]string{"json": "data"}),
	)
	return nil
}

func main() {
	helper.Generate(NewModelResponseGenerator())
}
