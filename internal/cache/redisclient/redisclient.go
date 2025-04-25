package redisclient

import (
	"backendv1/internal/jwt"
	"backendv1/pkg/errcheck"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisCheckFunc func(http.ResponseWriter, *http.Request, *RedisClient) error

type RedisClient struct {
	Enable bool
	Client *redis.Client
	JWT    *jwt.JWT
}

func (rc *RedisClient) Ping() error {
	return rc.Client.Ping(context.Background()).Err()
}

func (rc *RedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	s := rc.Client.Set(ctx, key, value, expiration)
	if s.Err() != nil {
		log.Print("Failed to set redis cache for key: ", key, " | error: ", s.Err(), "\n")
		return s.Err()
	}
	log.Print("Set redis cache for key: ", key, " | value: ", value)
	return nil
}

func (rc *RedisClient) Get(ctx context.Context, key string) (string, error) {
	return rc.Client.Get(ctx, key).Result()
}

func (rc *RedisClient) Del(ctx context.Context, key string) {
	rc.Client.Del(ctx, key)
	log.Print("Del redis cache for key: ", key)
}

func (rc *RedisClient) Expire(ctx context.Context, key string, expiration time.Duration) {
	rc.Client.Expire(ctx, key, expiration)
	log.Print("Added Expire to redis cache key: ", key)
}

func (rc *RedisClient) ExpireAt(ctx context.Context, key string, expirationTime time.Time) {
	rc.Client.ExpireAt(ctx, key, expirationTime)
	log.Print("Added ExpireAt to redis cache key: ", key)
}

func (rc *RedisClient) Close() {
	err := rc.Client.Close()
	errcheck.LogIfError(err, "Failed to close redis client")
	log.Print("Closed redis client connection")
}
