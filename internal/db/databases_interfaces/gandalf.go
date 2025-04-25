package db

import (
	gandalfmodels "backendv1/internal/models/gandalf"
	"context"
)

type GandalfDB interface {
	Database
	GetEntityDetails(ctx context.Context) interface{}
	GetLegacyHolders(ctx context.Context, entity *gandalfmodels.LegacyHoldersRequest) interface{}
	GetUserRegisteredEntity(ctx context.Context, user *gandalfmodels.UserRegisteredEntityRequest) interface{}
	CreateEntityDetails(ctx context.Context, entity *gandalfmodels.EntityDetailsCreateRequest) interface{}
}
