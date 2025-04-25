package gandalfmodels

type EntityDetailsCreateRequest struct {
	EntityId *string        `json:"entity_id"`
	Entity   EntityMetaData `json:"entity"`
}

type LegacyHoldersRequest struct {
	EntityId string `json:"entity_id"`
}

type UserRegisteredEntityRequest struct {
	UserId string
}
