package authhandlers

import (
	"backendv1/internal/cache/redisclient"
	authrediscache "backendv1/internal/cache/redisclient/auth"
	dbi "backendv1/internal/db/databases_interfaces"
	"backendv1/internal/jwt"
	authmodels "backendv1/internal/models/auth"
	genericresponses "backendv1/internal/models/generic_responses"
	"backendv1/pkg/errcheck"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type AuthHandler struct {
	authDb dbi.AuthDB
	redisc *redisclient.RedisClient
	jwt    *jwt.JWT
}

func (h *AuthHandler) AuthLoginValidator(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user authmodels.UserAuthLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println("Failed to decode json request", err)
		SendResponse(w, genericresponses.GenericBadRequestResponse)
		return
	}
	user.Ip = r.RemoteAddr
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	response := h.authDb.LoginUser(ctx, &user)
	SendResponse(w, response)
}

func (h *AuthHandler) AuthValidateJWT(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	token := r.Header.Get("Authorization")
	response := h.authDb.ValidateJWT(ctx, token)
	if response == authmodels.JwtInvalidResponse {
		go authrediscache.SetExpiredJwtInRedis(h.redisc, context.Background(), token, "0")
	}
	SendResponse(w, response)
}

func (h *AuthHandler) AuthLogoutUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	dbctx, dbcancel := context.WithTimeout(r.Context(), 10*time.Second)
	token := r.Header.Get("Authorization")
	response := h.authDb.LogoutUser(dbctx, &authmodels.UserAuthLogoutRequest{Token: token})
	defer dbcancel()
	rctx, rcancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer rcancel()
	err := authrediscache.SetExpiredJwtInRedis(h.redisc, rctx, token, "0")
	errcheck.LogIfError(err, "Failed to set expired jwt in redis for token: "+token)
	SendResponse(w, response)
}

func (h *AuthHandler) AuthChangePassword(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user authmodels.UserAuthChangePasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println("Failes to decode json request", err)
		SendResponse(w, genericresponses.GenericBadRequestResponse)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	c, err := h.jwt.ParseJWT(r.Header.Get("Authorization"))
	if err != nil {
		SendResponse(w, genericresponses.GenericInternalServerErrorResponse)
	} else if !h.jwt.IsJWTValid(c) {
		SendResponse(w, authmodels.JwtInvalidResponse)
	} else if user.NewPassword != user.ConfirmNewPassword {
		SendResponse(w, authmodels.ConfirmPasswordDoesNotMatchResponse)
		return
	}
	user.Username = h.jwt.GetUserName(c)
	response := h.authDb.ChangePassword(ctx, &user)
	SendResponse(w, response)
}

func (h *AuthHandler) UpdateMobileNumber(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user authmodels.UserAuthUpdateMobileNumberRequest
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println("Failes to decode json request", err)
		SendResponse(w, genericresponses.GenericBadRequestResponse)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	c, err := h.jwt.ParseJWT(r.Header.Get("Authorization"))
	if err != nil {
		SendResponse(w, genericresponses.GenericInternalServerErrorResponse)
	}
	user.Username = h.jwt.GetUserName(c)
	response := h.authDb.UpdateMobileNumber(ctx, &user)
	SendResponse(w, response)
}

func (h *AuthHandler) GetUserDetails(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	c, err := h.jwt.ParseJWT(r.Header.Get("Authorization"))
	if err != nil {
		SendResponse(w, genericresponses.GenericInternalServerErrorResponse)
	}
	var user authmodels.GetUserRequest
	user.UserId = h.jwt.GetUserId(c)
	response := h.authDb.GetUserDetails(ctx, &user)
	SendResponse(w, response)
}
