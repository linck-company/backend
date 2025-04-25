package db

import (
	hecatemodels "backendv1/internal/models/hecate"
	"context"
)

type HecateDB interface {
	Database
	GetEventDetails(ctx context.Context, username string) interface{}
	CreateEvent(ctx context.Context, event *hecatemodels.CreateEventRequest) interface{}
	RegisterForEvent(ctx context.Context, event *hecatemodels.RegisterForEventRequest) interface{}
	UnRegisterForEvent(ctx context.Context, event *hecatemodels.UnRegisterForEventRequest) interface{}
	GetStudentRecord(ctx context.Context, user *hecatemodels.GetStudentRecordRequest) interface{}

	CloseEvent(ctx context.Context, event *hecatemodels.CloseEventRequest) interface{}
}
