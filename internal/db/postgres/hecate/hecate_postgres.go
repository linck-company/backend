package hecatepg

import (
	genricresponses "backendv1/internal/models/generic_responses"
	hecatemodels "backendv1/internal/models/hecate"
	"context"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type HecatePostgres struct {
	Pool *pgxpool.Pool
}

func (hdb *HecatePostgres) GetEventDetails(ctx context.Context, username string) interface{} {
	query := hdb.getEventsQuery()
	rows, err := hdb.Pool.Query(ctx, query)
	if err != nil {
		log.Println(err)
		return genricresponses.GenericInternalServerErrorResponse
	}
	var response hecatemodels.GetEventsResponse
	marshal := hdb.marshalGetEventsDbResponse(rows, &response)
	if err = rows.Err(); err != nil || marshal != nil {
		log.Println(err)
		return genricresponses.GenericInternalServerErrorResponse
	}
	response.StatusCode = http.StatusOK
	return response
}

func (hdb *HecatePostgres) CreateEvent(ctx context.Context, event *hecatemodels.CreateEventRequest) interface{} {
	query := hdb.createEventQuery()
	status := hdb.insertEvent(ctx, query, event)
	switch v := status.(type) {
	case genricresponses.GenericResponse:
		return v
	case int:
		return &hecatemodels.CreateEventResponse{
			StatusCode: http.StatusOK,
			EventId:    v,
		}
	}
	return genricresponses.GenericInternalServerErrorResponse
}

func (hdb *HecatePostgres) RegisterForEvent(ctx context.Context, event *hecatemodels.RegisterForEventRequest) interface{} {
	query := `INSERT INTO usereventregistrations (event_id, username) VALUES ($1, $2)`
	_, err := hdb.Pool.Exec(ctx, query, event.EventId, event.UserId)
	if err != nil {
		log.Println("SQL", err)
		return genricresponses.GenericInternalServerErrorResponse
	}
	return &hecatemodels.RegisterForEventResponse{}
}

func (hdb *HecatePostgres) UnRegisterForEvent(ctx context.Context, event *hecatemodels.UnRegisterForEventRequest) interface{} {
	query := `DELETE FROM usereventregistrations WHERE event_id = $1 AND username = $2`
	_, err := hdb.Pool.Exec(ctx, query, event.EventId, event.UserId)
	if err != nil {
		log.Println("SQL", err)
		return genricresponses.GenericInternalServerErrorResponse
	}
	return &hecatemodels.UnRegisterForEventResponse{}
}

func (hdb *HecatePostgres) GetStudentRecord(ctx context.Context, user *hecatemodels.GetStudentRecordRequest) interface{} {
	query := hdb.getStudentRecordQuery()
	rows, err := hdb.Pool.Query(ctx, query, user.UserName)
	if err != nil {
		log.Println(err)
		return genricresponses.GenericInternalServerErrorResponse
	}
	var response hecatemodels.GetStudentRecordResponse
	marshal := hdb.marshalGetStudentRecordsDbResponse(rows, &response)
	if err = rows.Err(); err != nil || marshal != nil {
		log.Println(err)
		return genricresponses.GenericInternalServerErrorResponse
	}
	response.StatusCode = http.StatusOK
	return response
}

func (hdb *HecatePostgres) CloseEvent(ctx context.Context, event *hecatemodels.CloseEventRequest) interface{} {
	return nil
}
