package gandalfhandlers

import (
	"backendv1/internal/cache/redisclient"
	gandalfrediscache "backendv1/internal/cache/redisclient/gandalf"
	"backendv1/internal/config"
	gandalfmodels "backendv1/internal/models/gandalf"
	genericresponses "backendv1/internal/models/generic_responses"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

func NewGandalfHandler(dsn string) *GandalfHandler {
	pdb := config.NewGandalfPostgres(dsn, "GANDALF_DB_MAX_CONNECTIONS", "GANDALF_DB_MIN_CONNECTIONS")
	return &GandalfHandler{
		gandalfDb:     pdb,
		redisc:        config.GetNewRedisClient(),
		jwt:           config.NewJWT(),
		EntityIdParam: "eid",
	}
}

func SendResponse(w http.ResponseWriter, rc *redisclient.RedisClient, response interface{}) {
	switch v := response.(type) {
	case genericresponses.GenericResponse:
		w.WriteHeader(v.StatusCode)
	case gandalfmodels.EntityDetailsResponse:
		w.WriteHeader(v.StatusCode)
		if v.StatusCode == http.StatusOK {
			go gandalfrediscache.PutEntityDetails(rc, v)
		}
	case *gandalfmodels.EntityCreationResponse:
		w.WriteHeader(v.StatusCode)
		if v.StatusCode == http.StatusCreated {
			gandalfrediscache.DeleteEntityDetails(rc)
		}
	case gandalfmodels.LegacyHoldersResponse:
		w.WriteHeader(v.StatusCode)
		// if v.StatusCode == http.StatusOK {
		// 	go gandalfrediscache.PutLegacyHolders(rc, v)
		// }
	}
	json.NewEncoder(w).Encode(response)
}

func (h *GandalfHandler) validatateQueryParams(queryParams url.Values) (bool, interface{}) {
	if len(queryParams) == 0 {
		response := genericresponses.GenericResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "No query parameters provided",
		}
		return false, response
	}
	if _, ok := queryParams[h.EntityIdParam]; !ok {
		response := genericresponses.GenericResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "No eid provided",
		}
		return false, response
	}
	return true, nil
}

func (h *GandalfHandler) setEntityIdFromName(entity *gandalfmodels.EntityDetailsCreateRequest) interface{} {
	if entity.Entity.Name == "" {
		return genericresponses.GenericBadRequestResponse
	}
	if entity.EntityId == nil {
		entity.EntityId = new(string)
	}
	*entity.EntityId = strings.ReplaceAll(strings.ToLower(entity.Entity.Name), " ", "_")
	return nil
}
