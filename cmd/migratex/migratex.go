package main

import (
	"flag"
	"log"
	"os"

	v4 "github.com/golang-migrate/migrate/v4"
	"github.com/joaomagfreitas/configx"
	"github.com/joaomagfreitas/migratex/migraters"
	"github.com/joaomagfreitas/migratex/migrations"
)

var down = flag.Bool("down", false, "rollsback to the previous version")
var all = flag.Bool("all", false, "runs all migrations")
var force = flag.Int("force", 0, "forces a version to reset dirty state")
var conn = flag.String("conn", "", "specifies the database connection string. alternative: use --config to point to a yaml config file")
var cfg = flag.String("config", "", "specifies the file path (absolute/relative) to a yaml config file that builds the database connection string")
var source = flag.String("source", "migrations", "specifies the file path (absolute/relative) of the directory that contains the migration files (*.sql)")

func init() {
	flag.Parse()
}

func main() {
	m := migrater()
	err := migration()(m)
	if err != nil {
		log.Fatal(err)
	}

	v, _, err := migrations.Version(m)
	if err != nil {
		log.Fatalf("migrated but failed to get current version, %v", err)
	}

	log.Printf("current version: %d", v)
}

func migrater() *v4.Migrate {
	st, err := os.Stat(*source)
	if err != nil {
		log.Fatal(err)
	}

	if !st.IsDir() {
		log.Fatalf("migrations source path specified (%s) is not a directory", *source)
	}

	conn := connectionString()
	fs := os.DirFS(*source)

	m, err := migraters.Sqlite(conn, fs)
	if err != nil {
		log.Fatal(err)
	}

	return m
}

func migration() func(m *v4.Migrate) error {
	var migration = migrations.Upgrade
	if *all {
		migration = migrations.UpgradeAll
	}

	if *down {
		migration = migrations.Downgrade
		if *all {
			migration = migrations.DowngradeAll
		}
	}

	if *force > 0 {
		migration = func(m *v4.Migrate) error {
			return migrations.Force(m, *force)
		}
	}

	return migration
}

func connectionString() string {
	if len(*conn) > 0 {
		return *conn
	}

	if len(*cfg) == 0 {
		log.Fatal("neither connection string (--conn) or config file (--config) specified.")
	}

	cfg, err := configx.Unmarshal[configx.Sqlite](*cfg)
	if err != nil {
		log.Fatal(err)
	}

	return cfg.Conn()
}
