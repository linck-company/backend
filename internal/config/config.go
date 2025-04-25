package config

import (
	"backendv1/internal/cache/redisclient"
	authpg "backendv1/internal/db/postgres/auth"
	gandalfpg "backendv1/internal/db/postgres/gandalf"
	hecatepg "backendv1/internal/db/postgres/hecate"
	"backendv1/internal/jwt"
	"backendv1/pkg/errcheck"
	"backendv1/pkg/utils"
	"context"
	"crypto/tls"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Config holds the application configuration settings.
type Config struct {
	HTTPServer struct {
		ADDR string
	}
	AuthDbUrl         string
	HecateDbUrl       string
	GandalfDbUrl      string
	ApiEndpointPrefix string
}

// LoadHTTPServerConfig loads the configuration from environment variables.
func LoadHTTPServerConfig() *Config {
	if err := utils.LoadEnv(); err != nil {
		log.Println("Error loading environment variables:", err)
	}

	return &Config{
		HTTPServer: struct{ ADDR string }{
			ADDR: utils.GetStringEnv("HTTP_SERVER_ADDR", "0.0.0.0:8081"),
		},
		AuthDbUrl:         utils.GetStringEnv("AUTH_DB_URL", ""),
		HecateDbUrl:       utils.GetStringEnv("HECATE_DB_URL", ""),
		GandalfDbUrl:      utils.GetStringEnv("GANDALF_DB_URL", ""),
		ApiEndpointPrefix: utils.GetStringEnv("API_URL_PREFIX", ""),
	}
}

func GetNewRedisClient() *redisclient.RedisClient {
	isRedisEnabled := utils.GetBoolEnv("REDIS_CLIENT_ENABLE", false)
	rc := &redisclient.RedisClient{
		Client: nil,
		Enable: isRedisEnabled,
		JWT:    NewJWT(),
	}
	if !isRedisEnabled {
		log.Println("Redis client not enabled")
		return rc
	}
	rc.Client = redis.NewClient(&redis.Options{
		Addr:     utils.GetStringEnv("REDIS_CLIENT_ADDR", "localhost:6379"),
		Password: utils.GetStringEnv("REDIS_CLIENT_PASSWORD", ""),
		Username: utils.GetStringEnv("REDIS_CLIENT_USERNAME", ""),
		DB:       utils.GetIntEnv("REDIS_CLIENT_DB", 0),
		TLSConfig: &tls.Config{ // Enable TLS
			InsecureSkipVerify: utils.GetBoolEnv("REDIS_TLS_CONFIG_INSECURE_SKIP_VERIFY", true), // Set to false in production for security
		},
		PoolSize:     utils.GetIntEnv("REDIS_CLIENT_POOL_SIZE", 1),
		MinIdleConns: utils.GetIntEnv("REDIS_CLIENT_IDLE_CONNECTIONS", 1),
	})

	err := rc.Ping()
	if err != nil {
		log.Println("Error connecting to redis:", err, "\nDisabling redis client")
		rc.Client = nil
		rc.Enable = false
		return rc
	}
	log.Println("Redis Client Connected")
	return rc
}

func NewJWT() *jwt.JWT {
	return &jwt.JWT{
		Secret:         []byte(utils.GetStringEnv("JWT_SECRET_KEY", "")),
		ExpirationHour: time.Duration(utils.GetIntEnv("JWT_EXPIRATION_HOUR", 3)),
		RememberMeDays: time.Duration(utils.GetIntEnv("JWT_REMEMBER_ME_DAYS", 30)),
	}
}

func NewAuthPostgres(dsn string, maxConn, minConn string) *authpg.AuthPostgres {
	config, err := pgxpool.ParseConfig(dsn)
	errcheck.FatalIfError(err, "failed to parse database config")

	config.MaxConns = int32(utils.GetIntEnv(maxConn, 1))
	config.MinConns = int32(utils.GetIntEnv(minConn, 1))
	config.MaxConnIdleTime = time.Duration(utils.GetIntEnv("AUTH_DB_CONNECTION_MAX_LIFETIME", 0))

	authPool, err := pgxpool.NewWithConfig(context.Background(), config)
	errcheck.FatalIfError(err, "Failed to connect to auth database")
	err = authPool.Ping(context.Background())
	errcheck.FatalIfError(err, "Connected to auth db but couldn't ping")
	log.Println("Connected to Auth PostgreSQL database")
	return &authpg.AuthPostgres{
		Jwt:            NewJWT(),
		Pool:           authPool,
		ActiveStatus:   utils.GetStringEnv("AUTH_DB_ACTIVE_STATUS", "ACTIVE"),
		InvitedStatus:  utils.GetStringEnv("AUTH_DB_INVITED_STATUS", "INVITED"),
		BlockedStatus:  utils.GetStringEnv("AUTH_DB_BLOCKED_STATUS", "BLOCKED"),
		InactiveStatus: utils.GetStringEnv("AUTH_DB_INACTIVE_STATUS", "INACTIVE"),
	}
}

func NewGandalfPostgres(dsn string, maxConn, minConn string) *gandalfpg.GandalfPostgres {
	config, err := pgxpool.ParseConfig(dsn)
	errcheck.FatalIfError(err, "failed to parse database config")
	config.MaxConns = int32(utils.GetIntEnv(maxConn, 1))
	config.MinConns = int32(utils.GetIntEnv(minConn, 1))
	config.MaxConnIdleTime = time.Duration(utils.GetIntEnv("GANDALF_DB_CONNECTION_MAX_LIFETIME", 0))
	gandalfPool, err := pgxpool.NewWithConfig(context.Background(), config)
	errcheck.FatalIfError(err, "Failed to connect to gandalf database")
	err = gandalfPool.Ping(context.Background())
	errcheck.FatalIfError(err, "Connected to gandalf db but couldn't ping")
	// var version string
	// if err := gandalfPool.QueryRow(context.Background(), "select version()").Scan(&version); err != nil {
	// fmt.Println(err)
	// }
	// fmt.Printf("version=%s\n", version)
	log.Println("Connected to Gandalf PostgreSQL database")
	return &gandalfpg.GandalfPostgres{
		Pool: gandalfPool,
	}
}

func NewHecatePostgres(dsn string, maxConn, minConn string) *hecatepg.HecatePostgres {
	config, err := pgxpool.ParseConfig(dsn)
	errcheck.FatalIfError(err, "failed to parse database config")
	config.MaxConns = int32(utils.GetIntEnv(maxConn, 1))
	config.MinConns = int32(utils.GetIntEnv(minConn, 1))
	config.MaxConnIdleTime = time.Duration(utils.GetIntEnv("HECATE_DB_CONNECTION_MAX_LIFETIME", 0))
	hecatePool, err := pgxpool.NewWithConfig(context.Background(), config)
	errcheck.FatalIfError(err, "Failed to connect to hecate database")
	err = hecatePool.Ping(context.Background())
	errcheck.FatalIfError(err, "Connected to hecate db but couldn't ping")
	log.Println("Connected to Hecate PostgreSQL database")
	return &hecatepg.HecatePostgres{
		Pool: hecatePool,
	}
}
