package web

import (
	"fmt"
	"net/http"

	"github.com/shiro8613/litebans-go/src/config"
	"github.com/shiro8613/litebans-go/src/database"
	"github.com/shiro8613/litebans-go/src/web/engine"
	"github.com/shiro8613/litebans-go/src/web/handler/pages"
	"github.com/shiro8613/litebans-go/src/web/middlewares/errorpages"
	"github.com/shiro8613/litebans-go/src/web/middlewares/logger"
)

func StartWebServer(configWb config.Config_Web, configSt config.Config_Settings, db database.DBConnection) error {

	fileServer := http.FileServer(http.Dir(fmt.Sprintf("%s/static/", "./views")))
	engine := engine.New(configWb.ViewsDirectory, configWb.Name, configSt)
	pages := pages.New(config.Config_Settings{}, engine, fileServer)
	errorm := errorpages.NewErrorPage(pages)
	loggerware := logger.Logger(errorm)

	return http.ListenAndServe(configWb.Listen, loggerware)
}