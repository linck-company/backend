package hecatehandlers

import (
	"backendv1/internal/cache/redisclient"
	dbi "backendv1/internal/db/databases_interfaces"
	"backendv1/internal/jwt"
	genricresponses "backendv1/internal/models/generic_responses"
	hecatemodels "backendv1/internal/models/hecate"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type HecateHandlers struct {
	hecateDb dbi.HecateDB
	redisc   *redisclient.RedisClient
	jwt      *jwt.JWT
}

func (h *HecateHandlers) GetEventDetails(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	c, err := h.jwt.ParseJWT(r.Header.Get("Authorization"))
	if err != nil {
		SendResponse(w, genricresponses.GenericInternalServerErrorResponse)
	}
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	username := h.jwt.GetUserName(c)
	response := h.hecateDb.GetEventDetails(ctx, username)
	SendResponse(w, response)
}

func (h *HecateHandlers) CreateEventDetails(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	var event hecatemodels.CreateEventRequest
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		log.Println("Failed to decode json request", err)
		SendResponse(w, genricresponses.GenericBadRequestResponse)
		return
	}
	response := h.hecateDb.CreateEvent(ctx, &event)
	SendResponse(w, response)
}

func (h *HecateHandlers) CloseEvent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
}

func (h *HecateHandlers) RegisterForEvent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	var event hecatemodels.RegisterForEventRequest
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		log.Println("Failed to decode json request", err)
		SendResponse(w, genricresponses.GenericBadRequestResponse)
		return
	}
	c, err := h.jwt.ParseJWT(r.Header.Get("Authorization"))
	if err != nil {
		SendResponse(w, genricresponses.GenericInternalServerErrorResponse)
	}
	userId := h.jwt.GetUserName(c)
	event.UserId = userId
	response := h.hecateDb.RegisterForEvent(ctx, &event)
	SendResponse(w, response)
}

func (h *HecateHandlers) UnRegisterForEvent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	var event hecatemodels.UnRegisterForEventRequest
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		log.Println("Failed to decode json request", err)
		SendResponse(w, genricresponses.GenericBadRequestResponse)
		return
	}
	c, err := h.jwt.ParseJWT(r.Header.Get("Authorization"))
	if err != nil {
		SendResponse(w, genricresponses.GenericInternalServerErrorResponse)
	}
	userId := h.jwt.GetUserName(c)
	event.UserId = userId
	response := h.hecateDb.UnRegisterForEvent(ctx, &event)
	SendResponse(w, response)
}

func (h *HecateHandlers) GetStudentRecords(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	var user hecatemodels.GetStudentRecordRequest
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println("Failed to decode json request", err)
		SendResponse(w, genricresponses.GenericBadRequestResponse)
		return
	}
	resonse := h.hecateDb.GetStudentRecord(ctx, &user)
	SendResponse(w, resonse)
}
