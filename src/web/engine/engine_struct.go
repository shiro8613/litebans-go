package engine

import (
	"os"

	"github.com/shiro8613/litebans-go/src/config"
)

type Engine struct {
	path		string
	Name		string
	Settings	config.Config_Settings
	Htmls		map[string]os.File
}

type engineConsts string 

const (
	common_filename engineConsts = "common.html"
)

type Mapper struct {
	Var			map[string]string
	Component 	map[string]string
	Content 	map[string]string
}