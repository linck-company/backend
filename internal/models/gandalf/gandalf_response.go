package gandalfmodels

type EntityDetailsResponse struct {
	StatusCode int               `json:"status_code"`
	Entites    []*EntityMetaData `json:"entities"`
}

type EntityCreationResponse struct {
	StatusCode int    `json:"status_code"`
	EntityId   string `json:"entity_id"`
	EntityName string `json:"entity_name"`
	Message    string `json:"message"`
}

type UserRegisteredEntityResponse struct {
	StatusCode int    `json:"status_code"`
	EntityId   string `json:"entity_id"`
}

type LegacyHolder struct {
	Name     string `json:"name"`
	Title    string `json:"title"`
	ImageUrl string `json:"image_url"`
}

// LegacyHoldersResponse is the struct used for the HTTP response body
type LegacyHoldersResponse struct {
	Data       map[string][]LegacyHolder `json:"data"`
	StatusCode int                       `json:"status_code"`
}
