package pages

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/shiro8613/litebans-go/src/config"
	"github.com/shiro8613/litebans-go/src/web/engine"
)

func New(configSt config.Config_Settings, engine engine.Engine, fileServer http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/static/") {
			fileServer.ServeHTTP(w, r)
		} else {
			var d_map_pattern string

			switch len(strings.Split(r.URL.Path, "/")) {
			case 1: d_map_pattern = "/"
			case 2: d_map_pattern = "/{page_name}"
			case 3: d_map_pattern = "/{page_name}/{page}"
			case 4: d_map_pattern = "/{page_name}/{page}/{id}"
			}

			d_map := urlDecoder(r.URL.Path, d_map_pattern)
			fmt.Println(strings.Split(r.URL.Path,"/"), len(strings.Split(r.URL.Path,"/")))
			if len(d_map) == 0 && d_map_pattern != "/" {
				w.WriteHeader(http.StatusNotFound)
			} else {
				generatePages(d_map, engine).ServeHTTP(w, r)
			}
		}
	})
}

/* dmap

/
/ban/1
/ban/1/123456
/history/1
/hisotry/1/aaaaaaaaaa13424aa
*/