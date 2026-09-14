package migraters

import (
	"database/sql"
	"io/fs"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "modernc.org/sqlite"
)

func Sqlite(conn string, fs fs.FS) (*migrate.Migrate, error) {
	db, err := sql.Open("sqlite", conn)
	if err != nil {
		return nil, err
	}

	dd, err := sqlite.WithInstance(db, &sqlite.Config{})
	if err != nil {
		return nil, err
	}

	sd, err := iofs.New(fs, ".")
	if err != nil {
		return nil, err
	}

	return migrate.NewWithInstance("iofs", sd, "sqlite", dd)
}
