package main

import (
	"backendv1/internal/config"
	"backendv1/internal/http/router"
	"backendv1/middleware"
	authtokenvalidator "backendv1/middleware/auth_token_validator"
	"backendv1/middleware/cors"
	apilogger "backendv1/middleware/logger"
	rediscache "backendv1/middleware/redis_cache"
	"backendv1/pkg/utils"
	"log"
	"net/http"
)

func main() {
	// Load the server configuration
	utils.AsciiText()
	serverConfig := config.LoadHTTPServerConfig()
	redisClient := config.GetNewRedisClient()
	//Initialize the PostgreSQL connection

	//authpg, err := db.AuthPostgres(serverConfig.AuthDbUrl)
	//if err != nil {
	//	log.Fatal("Error connecting to database:", err)
	//}
	//authpg.InitAuthDbSchema()
	//defer authpg.Close()
	// Initialize the calls database interface
	mux := router.InitRouter(serverConfig)

	server := &http.Server{
		Addr: serverConfig.HTTPServer.ADDR,
		Handler: middleware.ChainMiddlewares(
			cors.CORSMiddleware,
			apilogger.APIRequestLogger,
			authtokenvalidator.AuthTokenValidator(redisClient),
			rediscache.RedisValidator(redisClient))(mux),

		// Handler: middleware.ChainMiddlewares(apilogger.APIRequestLogger)(mux),
	}
	log.Printf("Starting server on %s...\n", serverConfig.HTTPServer.ADDR)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
