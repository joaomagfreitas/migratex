package migrations

import "github.com/golang-migrate/migrate/v4"

func Force(m *migrate.Migrate, version int) error {
	return m.Force(version)
}
