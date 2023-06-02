package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

func load(path string) (Config, error) {

	b, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var config Config

	err = yaml.Unmarshal(b, config)
	if err != nil {
		return Config{}, err
	}

	return config, nil

}