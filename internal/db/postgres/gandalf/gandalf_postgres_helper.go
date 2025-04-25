package gandalfpg

import (
	gandalfmodels "backendv1/internal/models/gandalf"
	genricresponses "backendv1/internal/models/generic_responses"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/jackc/pgx/v5"
)

func (pdb *GandalfPostgres) Close() error {
	log.Println("Closing postgres database connection")
	pdb.Pool.Close()
	return nil
}

func (pdb *GandalfPostgres) Ping() error {
	return pdb.Pool.Ping(context.Background())
}

func (pdb *GandalfPostgres) marshalGetEntitiesDbResponse(rows pgx.Rows, response *gandalfmodels.EntityDetailsResponse) interface{} {
	defer rows.Close()
	for rows.Next() {
		var coreMembersJSON string
		entity := gandalfmodels.EntityMetaData{}

		err := rows.Scan(
			&entity.Name,
			&entity.Founder,
			&entity.CoFounder,
			&entity.ClubEmail,
			&entity.Description,
			&entity.YearEstablished,
			&entity.ClubLogoImageUrl,
			&entity.ClubBannerImageUrl,
			&entity.ClubWebsiteUrl,
			&entity.ClubTwitterUrl,
			&entity.ClubYoutubeUrl,
			&entity.ClubFacebookUrl,
			&entity.ClubLinkedinUrl,
			&entity.ClubInstagramUrl,
			&coreMembersJSON,
		)
		if err != nil {
			log.Println(err)
			return genricresponses.GenericInternalServerErrorResponse
		}

		var coreMembers []*gandalfmodels.CurrentCoreMembersMapper
		err = json.Unmarshal([]byte(coreMembersJSON), &coreMembers)
		if err != nil {
			log.Println(err)
			return genricresponses.GenericInternalServerErrorResponse

		}
		entity.CurrentCoreMembers = coreMembers
		response.Entites = append(response.Entites, &entity)
	}

	return nil
}

func (pdb *GandalfPostgres) insertEntities(ctx context.Context, query string, entity *gandalfmodels.EntityDetailsCreateRequest) interface{} {
	var entityId, entityName string
	err := pdb.Pool.QueryRow(
		ctx,
		query,
		entity.EntityId,
		entity.Entity.Name,
		entity.Entity.Description,
		entity.Entity.YearEstablished,
		entity.Entity.Founder,
		entity.Entity.ClubLogoImageUrl,
		entity.Entity.CoFounder,
		entity.Entity.ClubBannerImageUrl,
		entity.Entity.ClubWebsiteUrl,
		entity.Entity.ClubFacebookUrl,
		entity.Entity.ClubTwitterUrl,
		entity.Entity.ClubInstagramUrl,
		entity.Entity.ClubYoutubeUrl,
		entity.Entity.ClubLinkedinUrl,
		entity.Entity.ClubEmail,
	).Scan(&entityId, &entityName)

	if err != nil {
		log.Println(err)
		return genricresponses.GenericInternalServerErrorResponse
	}
	return []string{entityId, entityName}
}

func (pdb *GandalfPostgres) insertCurrentCoreMembers(ctx context.Context, query string, entity *gandalfmodels.EntityDetailsCreateRequest, entityName string) interface{} {
	var wg sync.WaitGroup
	var mu sync.Mutex

	var hasError error = nil
	for _, coreMember := range entity.Entity.CurrentCoreMembers {
		wg.Add(1)
		go func(coreMember *gandalfmodels.CurrentCoreMembersMapper) {
			defer wg.Done()
			_, err := pdb.Pool.Exec(ctx, query, *entity.EntityId, coreMember.Name, coreMember.Title, coreMember.ImageUrl)
			if err != nil {
				mu.Lock()
				hasError = err
				mu.Unlock()
			}
		}(coreMember)
	}
	wg.Wait()

	if hasError != nil {
		log.Println("Cannot update all of the core members for the entity: ", *entity.EntityId, hasError)
		return &gandalfmodels.EntityCreationResponse{
			StatusCode: http.StatusPartialContent,
			EntityId:   *entity.EntityId,
			EntityName: entityName,
			Message:    "Creation Partially successful",
		}
	}
	return &gandalfmodels.EntityCreationResponse{
		StatusCode: http.StatusCreated,
		EntityId:   *entity.EntityId,
		EntityName: entityName,
		Message:    "Entity created successfully",
	}
}
