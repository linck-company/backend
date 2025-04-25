package hecatehandlers

import (
	"backendv1/internal/config"
	genricresponses "backendv1/internal/models/generic_responses"
	hecatemodels "backendv1/internal/models/hecate"
	"encoding/json"
	"net/http"
)

func NewHecateHandlers(dsn string) *HecateHandlers {
	return &HecateHandlers{
		hecateDb: config.NewHecatePostgres(dsn, "HECATE_DB_MAX_CONNECTIONS", "HECATE_DB_MIN_CONNECTIONS"),
		redisc:   config.GetNewRedisClient(),
		jwt:      config.NewJWT(),
	}
}

func SendResponse(w http.ResponseWriter, response interface{}) {
	switch v := response.(type) {
	case genricresponses.GenericResponse:
		w.WriteHeader(v.StatusCode)
	case hecatemodels.GetEventsResponse:
		w.WriteHeader(v.StatusCode)
	case hecatemodels.CreateEventResponse:
		w.WriteHeader(v.StatusCode)
	case hecatemodels.RegisterForEventResponse:
		w.WriteHeader(http.StatusOK)
	case hecatemodels.UnRegisterForEventResponse:
		w.WriteHeader(http.StatusOK)
	case hecatemodels.GetStudentRecordResponse:
		w.WriteHeader(v.StatusCode)
	}
	json.NewEncoder(w).Encode(response)
}
