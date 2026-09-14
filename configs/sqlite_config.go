package configs

import (
	"fmt"
)

type sqliteConfig struct {
	Path string
}

func (cfg sqliteConfig) Conn() string {
	return fmt.Sprintf("file://%s", cfg.Path)
}

func Sqlite(p string) (sqliteConfig, error) {
	cfg := sqliteConfig{}
	err := unmarshal(p, &cfg)

	return cfg, err
}
