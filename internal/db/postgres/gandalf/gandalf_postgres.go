package gandalfpg

import (
	gandalfmodels "backendv1/internal/models/gandalf"
	genricresponses "backendv1/internal/models/generic_responses"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type GandalfPostgres struct {
	Pool *pgxpool.Pool
}

func (pdb *GandalfPostgres) GetEntityDetails(ctx context.Context) interface{} {
	query := pdb.fetchGetEntityDetailsQuery()

	rows, err := pdb.Pool.Query(ctx, query)
	if err != nil {
		log.Println(err)
		return genricresponses.GenericInternalServerErrorResponse
	}

	var response gandalfmodels.EntityDetailsResponse
	marshal := pdb.marshalGetEntitiesDbResponse(rows, &response)
	if err = rows.Err(); err != nil || marshal != nil {
		log.Println(err)
		return genricresponses.GenericInternalServerErrorResponse
	}
	response.StatusCode = http.StatusOK
	return response
}

func (pdb *GandalfPostgres) GetLegacyHolders(ctx context.Context, entity *gandalfmodels.LegacyHoldersRequest) interface{} {
	query := pdb.fetchGetLegacyHoldersQuery()
	var jsonData json.RawMessage
	err := pdb.Pool.QueryRow(ctx, query, entity.EntityId).Scan(&jsonData)
	response := gandalfmodels.LegacyHoldersResponse{
		Data:       make(map[string][]gandalfmodels.LegacyHolder),
		StatusCode: http.StatusOK,
	}
	if err != nil {
		if err == sql.ErrNoRows {
			return response
		}
		fmt.Println("Error in GetLegacyHolders: ", err)
		return genricresponses.GenericInternalServerErrorResponse
	}

	if len(jsonData) == 0 || string(jsonData) == "null" {
		return response
	}

	err = json.Unmarshal(jsonData, &response.Data)
	if err != nil {
		fmt.Println("Error in GetLegacyHolders Marshal: ", err)
		return genricresponses.GenericInternalServerErrorResponse
	}
	return response
}

func (pdb *GandalfPostgres) CreateEntityDetails(ctx context.Context, entity *gandalfmodels.EntityDetailsCreateRequest) interface{} {
	query := pdb.fetchPostEntityDetailsQuery()
	status := pdb.insertEntities(ctx, query, entity)

	switch v := status.(type) {
	case genricresponses.GenericResponse:
		return v
	case []string:
		query = pdb.fetchPostCurrentCoreMembersQuery()
		if v[0] != *entity.EntityId || v[1] != entity.Entity.Name {
			return genricresponses.GenericInternalServerErrorResponse
		}
		return pdb.insertCurrentCoreMembers(ctx, query, entity, v[1])
	}

	return genricresponses.GenericInternalServerErrorResponse
}

func (pdb *GandalfPostgres) GetUserRegisteredEntity(ctx context.Context, user *gandalfmodels.UserRegisteredEntityRequest) interface{} {
	query := `SELECT entities_id FROM user_registered_entities WHERE user_id = $1`
	var response gandalfmodels.UserRegisteredEntityResponse

	err := pdb.Pool.QueryRow(ctx, query, user.UserId).Scan(&response.EntityId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return genricresponses.GenericInternalServerErrorResponse
		}
		log.Printf("Database query error in GetUserRegisteredEntity for user %s: %v", user.UserId, err)
		return genricresponses.GenericInternalServerErrorResponse
	}
	response.StatusCode = http.StatusOK
	return response
}
