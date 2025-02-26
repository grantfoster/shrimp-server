package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5"
)

func main() {
	direction := flag.String("d", "up", "the direction we'll migrate in (up or down)")
	flag.Parse()
	db, err := sql.Open("pgx", "postgres:////shrimp_server?sslmode=disable")
	if err != nil {
		slog.Error("error opening psql connection", "err", err.Error())
		os.Exit(0)
	}
	driver, err := pgx.WithInstance(db, &pgx.Config{})
	if err != nil {
		slog.Error("error initializing psql driver", "err", err.Error())
		os.Exit(0)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"pgx", driver)
	if err != nil {
		slog.Error("error creating migration instance", "err", err.Error())
		os.Exit(0)
	}

	if *direction == "up" {
		err = m.Up() // or m.Steps(2) if you want to explicitly set the number of migrations to run
		if err != nil {
			slog.Error("error migrating up", "err", err.Error())
			os.Exit(0)
		}
	} else if *direction == "down" {
		err = m.Down()
		if err != nil {
			slog.Error("error migrating down", "err", err.Error())
			os.Exit(0)
		}
	} else {
		slog.Error("invalid flag option for d (direction), valid directions are \"up\" or \"down\"")
		os.Exit(0)
	}
	slog.Info("migration successful", "direction", *direction)
}
