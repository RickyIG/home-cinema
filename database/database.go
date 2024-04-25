package database

import (
	"database/sql"
	"embed"
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
)

var (
	DbConnection *sql.DB
)

// cara ke-1: embed, note: perlu import "embed"

//go:embed sql_migrations/*.sql
var dbMigrations embed.FS

func DbMigrate(dbParam *sql.DB) {
	// cara ke-1: embed, note: perlu import "embed"
	migrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: dbMigrations,
		Root:       "sql_migrations",
	}

	// cara ke-2: packr, note: perlu import "github.com/gobuffalo/packr"
	// migrations := &migrate.PackrMigrationSource{
	// 	Box: packr.NewBox("./sql_migrations"),
	// }

	n, errs := migrate.Exec(dbParam, "postgres", migrations, migrate.Up)
	if errs != nil {
		panic(errs)
	}

	DbConnection = dbParam

	fmt.Println("Applied ", n, " migrations!")
}
