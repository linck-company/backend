package authtokenvalidator

import (
	"backendv1/internal/cache/redisclient"
	authrediscache "backendv1/internal/cache/redisclient/auth"
	"backendv1/internal/config"
	authmodels "backendv1/internal/models/auth"
	"backendv1/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

// AuthTokenValidator - jwt token validator middleware service
func AuthTokenValidator(rc *redisclient.RedisClient) func(http.Handler) http.Handler {
	j := config.NewJWT()
	return func(next http.Handler) http.Handler {
		var skipValidationForLoginEndpoint = fmt.Sprintf("%s/auth/login", utils.GetStringEnv("API_URL_PREFIX", ""))
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.String() == skipValidationForLoginEndpoint {
				next.ServeHTTP(w, r)
				return
			}
			if authrediscache.CheckForJWT(w, r, rc) == nil {
				json.NewEncoder(w).Encode(authmodels.JwtInvalidResponse)
				return
			}
			c, err := j.ParseJWT(r.Header.Get("Authorization"))
			if err != nil || !j.IsJWTValid(c) {
				json.NewEncoder(w).Encode(authmodels.JwtInvalidResponse)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
