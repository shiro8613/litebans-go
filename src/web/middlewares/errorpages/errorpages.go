package errorpages

import (
	"net/http"

	"github.com/shiro8613/litebans-go/src/web/middlewares/logger"
)

func NewErrorPage(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		nl := logger.NewLoggingResponseWriter(w)
		next.ServeHTTP(nl, r)

		switch nl.StatusCode {
		case http.StatusNotFound:
			nl.Write([]byte("ねーよばーか"))
			break
		case http.StatusBadGateway:
			nl.Write([]byte("げーとうぇいがだめだーめ"))
			break
		case http.StatusInternalServerError:
			nl.Write([]byte("さーばーちゃんはご機嫌斜めな用で..."))
			break
		default:
			break
		}
	})
}