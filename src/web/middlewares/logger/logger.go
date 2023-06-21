package logger

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//aaa.bbb.ccc.ddd - - [day/month/year HH:MM:SS] "Method Path Protocol" Status -
		//100.10.0.1 - - [01/02/2023 12:13:14] "GET / HTTP/1.1" 200 -

		nl := NewLoggingResponseWriter(w)
		next.ServeHTTP(nl, r)

		time := time.Now().Format("02/01/2006T15:04:05")
		addr := r.RemoteAddr
		log := fmt.Sprintf(`[LOG] %s - - [%s] "%s %s %s" %d -`, time, addr, r.Method, r.URL.Path, r.Proto, nl.StatusCode)
		fmt.Println(log)
	})
}


