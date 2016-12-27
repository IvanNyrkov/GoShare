package middleware

import "net/http"

// loggedResponseWriter is a wrapper over the ResponseWriter interface that allows us to use its response status
type loggedResponseWriter struct {
	http.ResponseWriter
	status int
}

// WriteHeader wraps default ResponseWriter.WriteHeader with saving response status
func (l *loggedResponseWriter) WriteHeader(status int) {
	l.status = status
	l.ResponseWriter.WriteHeader(status)
}
