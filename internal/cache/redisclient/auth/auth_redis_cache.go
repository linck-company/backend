package authrediscache

import (
	"backendv1/internal/cache/redisclient"
	"backendv1/pkg/utils"
	"context"
	"errors"
	"net/http"
)

func CheckForJWT(w http.ResponseWriter, r *http.Request, rc *redisclient.RedisClient) error {
	_, err := rc.Get(r.Context(), r.Header.Get("Authorization"))
	return err
}

func RedisLoginReturnError(w http.ResponseWriter, r *http.Request, rc *redisclient.RedisClient) error {
	return errors.New("redis check is invalid for login")
}

func SetExpiredJwtInRedis(rc *redisclient.RedisClient, ctx context.Context, token string, value interface{}) error {
	return rc.Set(ctx, token, value, utils.RemainingTimeForExpiration(rc.JWT.GetExpirationEpoch(token)))
}
