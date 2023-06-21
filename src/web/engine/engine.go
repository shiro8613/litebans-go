package engine

import "github.com/shiro8613/litebans-go/src/config"

func New(path string, name string, config config.Config_Settings) Engine{ 
	return Engine{
		path: path,
		Name: name,
		Settings: config,
	}
}