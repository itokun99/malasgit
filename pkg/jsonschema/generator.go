//go:build ignore

package main

import (
	"fmt"

	"github.com/itokun99/malasgit/pkg/jsonschema"
)

func main() {
	fmt.Printf("Generating jsonschema in %s...\n", jsonschema.GetSchemaDir())
	schema := jsonschema.GenerateSchema()
	jsonschema.GenerateConfigDocs(schema)
}
