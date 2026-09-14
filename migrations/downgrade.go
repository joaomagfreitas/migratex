package migrations

import (
	"github.com/golang-migrate/migrate/v4"
)

func Downgrade(m *migrate.Migrate) error {
	return m.Steps(-1)
}

func DowngradeAll(m *migrate.Migrate) error {
	return m.Down()
}
