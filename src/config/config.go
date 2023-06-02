package config

import (
	"fmt"
	"os"
)

func LoadAndCreate(path string) (Config, error) {

	config, err := load(path)

	if os.IsNotExist(err) {
		fmt.Println("FileNotFound")
		create(path)
		os.Exit(1)
	} else if err != nil {
		return Config{}, err
	}

	return config, err
}
