package helper

import (
	"flag"
	"log"

	"github.com/0x4c6565/genie"
)

func Generate(generators ...genie.Generator) {
	var flagSource string
	var flagDestination string
	var flagPackage string
	flag.StringVar(&flagSource, "source", "", "Path to source file")
	flag.StringVar(&flagDestination, "destination", "", "Path to destination file")
	flag.StringVar(&flagPackage, "package", "", "Package name")
	flag.Parse()

	log.Printf("Generating %s from %s for package %s", flagDestination, flagSource, flagPackage)
	err := genie.NewGenie(
		genie.NewFileInput(flagSource),
		genie.NewFileOutput(flagDestination),
		flagPackage,
		generators...,
	).Generate()

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Finished generating %s", flagDestination)
}
