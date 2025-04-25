package rediscache

import (
	"backendv1/internal/cache/redisclient"
	gandalfrediscache "backendv1/internal/cache/redisclient/gandalf"
	"log"
	"net/http"
	"strings"
)

func RedisValidator(rc *redisclient.RedisClient) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !rc.Enable {
				next.ServeHTTP(w, r)
				return
			}

			if strings.Contains(r.URL.Path, "gandalf/entity/details") {
				log.Println("Cache trying")
				response, err := gandalfrediscache.GetEntityDetails(rc)
				if err == nil {
					log.Println("Sending result from cache")
					w.WriteHeader(http.StatusOK)
					w.Header().Set("Content-Type", "application/json")
					w.Write([]byte(response))
					return
				}
				next.ServeHTTP(w, r)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
