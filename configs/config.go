package configs

import (
	errors "github.com/pkg/errors"
	multierr "go.uber.org/multierr"
	"gopkg.in/yaml.v3"
	"io/fs"
	"os"
)

type Config struct {
	API API `yaml:"api"`
	DB  DB  `yaml:"db"`
}

type (
	API struct {
		Host string `yaml:"host"`
	}

	DB struct {
		Conn string `yaml:"conn"`
	}
)

func New(path string) (config Config, err error) {
	file, err := os.OpenFile(path, os.O_RDONLY, fs.ModePerm)
	if err != nil {
		return config, errors.Wrapf(err, "open config by path %s", path)
	}
	defer func(err error) {
		multierr.AppendInto(&err, file.Close())
	}(err)

	return config, errors.Wrap(
		yaml.NewDecoder(file).Decode(&config),
		"decode config information",
	)
}
