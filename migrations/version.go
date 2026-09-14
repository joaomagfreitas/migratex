package migrations

import (
	"github.com/golang-migrate/migrate/v4"
)

func Version(m *migrate.Migrate) (uint, bool, error) {
	return m.Version()
}
