package apilogger

import (
	"log"
	"net/http"
)

// APIRequestLogger - middleware logger service to log api calls
func APIRequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		go func() {
			log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		}()
		next.ServeHTTP(w, r)
	})
}
