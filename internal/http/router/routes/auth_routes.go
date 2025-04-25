package router

import (
	"backendv1/internal/http/handlers/authhandlers"
	"fmt"
	"net/http"
)

func InitAuthRoutes(mux *http.ServeMux, dsn string, apiPrefix string) {
	aho := authhandlers.NewAuthHandler(dsn)

	mux.HandleFunc(fmt.Sprintf("POST %s/auth/logout", apiPrefix), aho.AuthLogoutUser)
	mux.HandleFunc(fmt.Sprintf("POST %s/auth/login", apiPrefix), aho.AuthLoginValidator)
	mux.HandleFunc(fmt.Sprintf("POST %s/auth/validate", apiPrefix), aho.AuthValidateJWT)
	mux.HandleFunc(fmt.Sprintf("GET %s/auth/user/details", apiPrefix), aho.GetUserDetails)
	mux.HandleFunc(fmt.Sprintf("POST %s/auth/change_password", apiPrefix), aho.AuthChangePassword)
	mux.HandleFunc(fmt.Sprintf("POST %s/auth/update_mobile_number", apiPrefix), aho.UpdateMobileNumber)
}
