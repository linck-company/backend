package router

import (
	"backendv1/internal/config"
	router "backendv1/internal/http/router/routes"
	"net/http"
)

func InitRouter(serverConfig *config.Config) *http.ServeMux {
	mux := http.NewServeMux()
	router.InitAuthRoutes(mux, serverConfig.AuthDbUrl, serverConfig.ApiEndpointPrefix)
	router.InitGandalfRoutes(mux, serverConfig.GandalfDbUrl, serverConfig.ApiEndpointPrefix)
	router.InitOdysseyRoutes(mux, serverConfig.GandalfDbUrl, serverConfig.ApiEndpointPrefix)
	router.InitHecateRoutes(mux, serverConfig.HecateDbUrl, serverConfig.ApiEndpointPrefix)

	return mux
}
