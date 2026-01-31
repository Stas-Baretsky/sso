package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_"github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var storargePath, migrationsPath, migrationsTable string

	flag.StringVar(&storargePath, "storage-path", "", "path to storage")
	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations")
	flag.StringVar(&migrationsTable, "migrations", "", "name of migration table")
	flag.Parse()

	m, err := migrate.New(
		"file://"+migrationsPath,
		fmt.Sprintf("sqlite3://%s?x-migrations-table=%s", storargePath, migrationsTable),
	)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Printf("ni migrations to apply")

			return
		}
		panic(err)
	}
	fmt.Println("migrations applied successfully")
}
