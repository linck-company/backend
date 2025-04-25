package gandalfhandlers

import (
	"backendv1/internal/cache/redisclient"
	dbi "backendv1/internal/db/databases_interfaces"
	"backendv1/internal/jwt"
	gandalfmodels "backendv1/internal/models/gandalf"
	genricresponses "backendv1/internal/models/generic_responses"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type GandalfHandler struct {
	gandalfDb     dbi.GandalfDB
	redisc        *redisclient.RedisClient
	jwt           *jwt.JWT
	EntityIdParam string
}

func (h *GandalfHandler) GetEntityDetails(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()
	response := h.gandalfDb.GetEntityDetails(ctx)
	SendResponse(w, h.redisc, response)
}

func (h *GandalfHandler) GetLegacyHolders(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var entity gandalfmodels.LegacyHoldersRequest
	if err := json.NewDecoder(r.Body).Decode(&entity); err != nil {
		log.Println("Failed to decode json request", err)
		SendResponse(w, h.redisc, genricresponses.GenericInternalServerErrorResponse)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()
	response := h.gandalfDb.GetLegacyHolders(ctx, &entity)
	SendResponse(w, h.redisc, response)
}

func (h *GandalfHandler) CreateEntityDetails(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var entity gandalfmodels.EntityDetailsCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&entity); err != nil {
		log.Println("Failed to decode json request", err)
		SendResponse(w, h.redisc, genricresponses.GenericInternalServerErrorResponse)
		return
	}

	if status := h.setEntityIdFromName(&entity); status != nil {
		SendResponse(w, h.redisc, status)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*20)
	defer cancel()
	response := h.gandalfDb.CreateEntityDetails(ctx, &entity)
	SendResponse(w, h.redisc, response)
}

func (h *GandalfHandler) GetUserRegisteredEntity(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	c, errcancel := h.jwt.ParseJWT(r.Header.Get("Authorization"))
	if errcancel != nil {
		SendResponse(w, h.redisc, genricresponses.GenericInternalServerErrorResponse)
	}
	var user gandalfmodels.UserRegisteredEntityRequest
	user.UserId = h.jwt.GetUserId(c)
	response := h.gandalfDb.GetUserRegisteredEntity(ctx, &user)
	SendResponse(w, h.redisc, response)
}
