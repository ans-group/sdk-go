package main

import (
	"fmt"

	"github.com/0x4c6565/genie"
	"github.com/dave/jennifer/jen"
	"github.com/ukfast/sdk-go/pkg/gen/helper"
)

type ModelPaginatedGenerator struct {
	*genie.BaseGenerator
}

func NewModelPaginatedGenerator() *ModelPaginatedGenerator {
	return &ModelPaginatedGenerator{
		BaseGenerator: genie.NewBaseGenerator("model_paginated"),
	}
}

func (g *ModelPaginatedGenerator) Generate(marker genie.Marker, typeName string, f *jen.File) error {
	f.ImportName("github.com/ukfast/sdk-go/pkg/connection", "connection")

	f.Comment(fmt.Sprintf("Paginated%s represents a paginated collection of %[1]s", typeName))
	f.Type().Id("Paginated"+typeName).Struct(
		jen.Op("*").Qual("github.com/ukfast/sdk-go/pkg/connection", "PaginatedBase"),
		jen.Id("Items").Index().Id(typeName),
	)

	f.Comment(fmt.Sprintf("NewPaginated%s returns a pointer to an initialized Paginated%[1]s struct", typeName))
	f.Func().Id("NewPaginated"+typeName).Params(
		jen.Id("getFunc").Qual("github.com/ukfast/sdk-go/pkg/connection", "PaginatedGetFunc"),
		jen.Id("parameters").Qual("github.com/ukfast/sdk-go/pkg/connection", "APIRequestParameters"),
		jen.Id("pagination").Qual("github.com/ukfast/sdk-go/pkg/connection", "APIResponseMetadataPagination"),
		jen.Id("items").Index().Id(typeName),
	).Op("*").Id("Paginated" + typeName).Block(
		jen.Return(
			jen.Op("&").Id("Paginated" + typeName).Values(jen.Dict{
				jen.Id("Items"): jen.Id("items"),
				jen.Id("PaginatedBase"): jen.Qual("github.com/ukfast/sdk-go/pkg/connection", "NewPaginatedBase").Call(
					jen.Id("parameters"),
					jen.Id("pagination"),
					jen.Id("getFunc"),
				),
			}),
		),
	)

	return nil
}

func main() {
	helper.Generate(NewModelPaginatedGenerator())
}
