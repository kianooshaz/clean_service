package config

import (
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
	"os"
)

func ReadFile(cfg *Config, cfgFilePath string) error {
	f, err := os.Open(cfgFilePath)
	if err != nil {
		return err
	}

	defer func() {
		cErr := f.Close()
		if cErr != nil {
			err = cErr
		}
	}()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		return err
	}
	return nil
}

func ReadEnv(cfg *Config) error {
	err := envconfig.Process("", cfg)
	if err != nil {
		return err
	}

	return nil
}
