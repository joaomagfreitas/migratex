package configs

import (
	"os"

	"github.com/goccy/go-yaml"
)

func unmarshal(p string, cfg interface{}) error {
	var bs []byte

	bs, err := os.ReadFile(p)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(bs, &cfg)
}
