package hecatemodels

import "time"

type GetEventsResponse struct {
	StatusCode int                   `json:"status_code"`
	Events     []SingleEventResponse `json:"events"`
}

type SingleEventResponse struct {
	ID                         int       `json:"id" db:"id"`
	EventTitle                 *string   `json:"event_title" db:"event_title"`
	EventNature                *string   `json:"event_nature" db:"event_nature"`
	EventDate                  time.Time `json:"event_date" db:"event_date"`
	EventStartTime             time.Time `json:"event_start_time" db:"event_start_time"`
	EventEndTime               time.Time `json:"event_end_time" db:"event_end_time"`
	EventHost                  *string   `json:"event_host" db:"event_host"`
	EventDescription           *string   `json:"event_description" db:"event_description"`
	EventLocation              *string   `json:"event_location" db:"event_location"`
	EventOrganizerContactInfo  *string   `json:"event_organizer_contact_info" db:"event_organizer_contact_info"`
	EventSubjectArea           *string   `json:"event_subject_area" db:"event_subject_area"`
	EventResourcePerson        *string   `json:"event_resource_person" db:"event_resource_person"`
	EventAffiliation           *string   `json:"event_affiliation" db:"event_affiliation"`
	EventResourcePersonProfile *string   `json:"event_resource_person_profile" db:"event_resource_person_profile"`
	EventObjective             *string   `json:"event_objective" db:"event_objective"`
	EventFlyerImageUrl         *string   `json:"event_flyer_image_url" db:"event_flyer_image_url"`
	IsUserRegistered           bool      `json:"is_user_registered" db:"is_user_registered"`
	IsEventActive              bool      `json:"is_event_active" db:"is_active"`
	TotalRegistered            int       `json:"registered_count" db:"registered_count"`
	ClubName                   *string   `json:"club_name" db:"club_name"`
	ClubLogoImageUrl           *string   `json:"club_logo_image_url" db:"club_logo"`
	ClubId                     *string   `json:"club_id" db:"club_id"`
	ClubOrganizerName          *string   `json:"club_organization" db:"club_organizer_name"`
}

type CreateEventResponse struct {
	StatusCode int `json:"status_code"`
	EventId    int `json:"event_id"`
}

type RegisterForEventResponse struct {
}

type UnRegisterForEventResponse struct {
}

type GetStudentRecordResponse struct {
	StatusCode int                   `json:"status_code"`
	Events     []SingleStudentRecord `json:"events"`
}

type SingleStudentRecord struct {
	ClubName   string    `json:"club_name"`
	EventTitle string    `json:"event_title"`
	EventDate  time.Time `json:"event_date"`
}
