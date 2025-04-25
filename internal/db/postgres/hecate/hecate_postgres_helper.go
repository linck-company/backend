package hecatepg

import (
	genricresponses "backendv1/internal/models/generic_responses"
	hecatemodels "backendv1/internal/models/hecate"
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func (hdb *HecatePostgres) Close() error {
	log.Println("Closing postgres database connection")
	hdb.Pool.Close()
	return nil
}

func (hdb *HecatePostgres) Ping() error {
	return hdb.Pool.Ping(context.Background())
}

func (hdb *HecatePostgres) marshalGetEventsDbResponse(rows pgx.Rows, response *hecatemodels.GetEventsResponse) interface{} {
	defer rows.Close()
	for rows.Next() {
		event := hecatemodels.SingleEventResponse{}
		err := rows.Scan(
			&event.ID,
			&event.EventTitle,
			&event.EventNature,
			&event.EventDate,
			&event.EventStartTime,
			&event.EventEndTime,
			&event.EventHost,
			&event.EventDescription,
			&event.EventLocation,
			&event.EventOrganizerContactInfo,
			&event.EventSubjectArea,
			&event.EventResourcePerson,
			&event.EventAffiliation,
			&event.EventResourcePersonProfile,
			&event.EventObjective,
			&event.EventFlyerImageUrl,
			&event.ClubName,
			&event.ClubLogoImageUrl,
			&event.ClubId,
			&event.ClubOrganizerName,
			&event.TotalRegistered,
			&event.IsEventActive,
			&event.IsUserRegistered,
		)
		if err != nil {
			log.Println(err)
			return genricresponses.GenericInternalServerErrorResponse
		}
		response.Events = append(response.Events, event)
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
		return genricresponses.GenericInternalServerErrorResponse
	}
	return nil
}

func (hdb *HecatePostgres) insertEvent(ctx context.Context, query string, event *hecatemodels.CreateEventRequest) interface{} {
	var eventId int
	err := hdb.Pool.QueryRow(
		ctx,
		query,
		event.EventTitle,
		event.EventNature,
		event.EventDate,
		event.EventStartTime,
		event.EventEndTime,
		event.EventHost,
		event.EventVenue,
		event.EventOrganizerContactInfo,
		event.EventSubjectArea,
		event.EventResourcePerson,
		event.EventAffiliation,
		event.EventResourcePersonProfile,
		event.EventDescription,
		event.EventObjective,
		event.EventFlyerImageURL,
		event.ClubName,
		event.ClubId,
		event.ClubLogoImageUrl,
		event.ClubOrganizerName,
	).Scan(&eventId)
	if err != nil {
		log.Println(err)
		return genricresponses.GenericInternalServerErrorResponse
	}
	return eventId
}

func (hdb *HecatePostgres) marshalGetStudentRecordsDbResponse(rows pgx.Rows, response *hecatemodels.GetStudentRecordResponse) interface{} {
	defer rows.Close()
	for rows.Next() {
		event := hecatemodels.SingleStudentRecord{}
		err := rows.Scan(
			&event.EventTitle,
			&event.EventDate,
			&event.ClubName,
		)
		if err != nil {
			log.Println(err)
			return genricresponses.GenericInternalServerErrorResponse
		}
		response.Events = append(response.Events, event)
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
		return genricresponses.GenericInternalServerErrorResponse
	}
	return nil
}
