package gandalfrediscache

import (
	"backendv1/internal/cache/redisclient"
	gandalfmodels "backendv1/internal/models/gandalf"
	"context"
	"encoding/json"
	"log"
	"time"
)

const LEGACY_HOLDERS = "legacy_holders"
const ENTITY_DETAILS_KEY = "entity_details"

func PutEntityDetails(rc *redisclient.RedisClient, response gandalfmodels.EntityDetailsResponse) {
	data, err := json.Marshal(response)
	if err != nil {
		return
	}

	rc.Set(context.Background(), ENTITY_DETAILS_KEY, data, 0)
	log.Println("Entity data put in redis cache")
}

func DeleteEntityDetails(rc *redisclient.RedisClient) {
	rc.Del(context.Background(), ENTITY_DETAILS_KEY)
}

func PutLegacyHolders(rc *redisclient.RedisClient, response gandalfmodels.LegacyHoldersResponse) {
	data, err := json.Marshal(response)
	if err != nil {
		return
	}

	rc.Set(context.Background(), LEGACY_HOLDERS, data, 0)
	log.Println("Legacy holders data put in redis cache")
}

func DeleteLegacyHolders(rc *redisclient.RedisClient) {
	rc.Del(context.Background(), LEGACY_HOLDERS)
}

func GetEntityDetails(rc *redisclient.RedisClient) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	return rc.Get(ctx, ENTITY_DETAILS_KEY)
}
