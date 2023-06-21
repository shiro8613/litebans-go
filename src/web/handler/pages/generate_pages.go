package pages

import (
	"net/http"

	"github.com/shiro8613/litebans-go/src/web/engine"
)

func generatePages(d_map map[string]string, engin engine.Engine) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}