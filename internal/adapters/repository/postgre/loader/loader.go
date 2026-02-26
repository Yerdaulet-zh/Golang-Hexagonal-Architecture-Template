// Package main serves as a schema loader for Atlas.
// It uses the Atlas GORM provider to extract the database schema from
// GORM models (persistency structs) and output the equivalent SQL
// statements to stdout, enabling automated database migrations.
package main

import (
	"fmt"
	"io"
	"os"

	"ariga.io/atlas-provider-gorm/gormschema"
	persistency "gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/adapters/repository/postgre/persistency/notification"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	config := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}
	stmts, err := gormschema.New("postgres", gormschema.WithConfig(config)).Load(
		&persistency.EmailNotification{},
	)

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}
	io.WriteString(os.Stdout, stmts)
}
