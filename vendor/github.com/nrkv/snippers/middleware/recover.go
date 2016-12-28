package middleware

import (
	"net/http"
	"log"
)

// Recover is a recovery middleware that catches panics and responds with StatusInternalServerError
func Recover(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		inner.ServeHTTP(w, r)
	})
}