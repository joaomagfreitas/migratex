package migrations

import (
	"github.com/golang-migrate/migrate/v4"
)

func Upgrade(m *migrate.Migrate) error {
	return m.Steps(1)
}

func UpgradeAll(m *migrate.Migrate) error {
	return m.Up()
}
