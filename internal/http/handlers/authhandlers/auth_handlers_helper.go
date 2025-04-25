package authhandlers

import (
	"backendv1/internal/config"
	authmodels "backendv1/internal/models/auth"
	genericresponses "backendv1/internal/models/generic_responses"
	"encoding/json"
	"net/http"
)

func NewAuthHandler(dsn string) *AuthHandler {
	adb := config.NewAuthPostgres(dsn, "AUTH_DB_MAX_CONNECTIONS", "AUTH_DB_MIN_CONNECTIONS")
	return &AuthHandler{
		authDb: adb,
		redisc: config.GetNewRedisClient(),
		jwt:    config.NewJWT(),
	}
}

func SendResponse(w http.ResponseWriter, response interface{}) {
	switch v := response.(type) {
	case genericresponses.GenericResponse:
		w.WriteHeader(v.StatusCode)
	case authmodels.UserAuthLoginResponse:
		w.WriteHeader(v.StatusCode)
	case authmodels.UserAuthChangePasswordResponse:
		w.WriteHeader(v.StatusCode)
	case authmodels.GetUserDetailsResponse:
		w.WriteHeader(v.StatusCode)
	}
	json.NewEncoder(w).Encode(response)
}
