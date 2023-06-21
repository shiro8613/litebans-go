package logger

import "net/http"

type loggingResponseWriter struct {
    http.ResponseWriter
    StatusCode int
}