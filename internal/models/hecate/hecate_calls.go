package hecatemodels

type CreateEventRequest struct {
	EventTitle                 string `json:"event_title"`
	EventNature                string `json:"event_nature"`
	EventDate                  string `json:"event_date"`
	EventStartTime             string `json:"event_start_time"`
	EventEndTime               string `json:"event_end_time"`
	EventHost                  string `json:"event_host"`
	EventVenue                 string `json:"event_venue"`
	EventOrganizerContactInfo  string `json:"event_organizer_contact_info"`
	EventSubjectArea           string `json:"event_subject_area"`
	EventResourcePerson        string `json:"event_resource_person"`
	EventAffiliation           string `json:"event_affiliation"`
	EventResourcePersonProfile string `json:"event_resource_person_profile"`
	EventDescription           string `json:"event_description"`
	EventObjective             string `json:"event_objective"`
	EventFlyerImageURL         string `json:"event_flyer_image_url"`
	ClubName                   string `json:"club_name"`
	ClubLogoImageUrl           string `json:"club_logo_image_url"`
	ClubId                     string `json:"club_id"`
	ClubOrganizerName          string `json:"club_organizer_name"`
}

type RegisterForEventRequest struct {
	EventId int `json:"event_id"`
	UserId  string
}

type UnRegisterForEventRequest struct {
	EventId int `json:"event_id"`
	UserId  string
}

type CloseEventRequest struct {
}

type GetStudentRecordRequest struct {
	UserName string `json:"username"`
}

var _ = `{
	"status_code": 200,
	"event_title": "FAST WEDNESDAYS",
	"event_nature": "Competition",
	"event_date": "12-03-2024",
	"event_start_time": "16:00:00",
	"event_end_time": "18:00:00",
	"event_host": "Cubing Club",
	"event_location": "S310",
	"external_attendees": 50,
	"faculty_attendees": 10,
	"phd_attendees": 5,
	"pg_attendees": 15,
	"ug_attendees": 20,
	"event_organizer_contact_info": "Cubing Club",
	"event_subject_area": "Cubing",
	"event_resource_person": "NA",
	"event_affiliation": "Cubing Club",
	"event_resource_person_profile": "NA",
	"event_objective": "The Objective of the event is to teach new solving techniques and allow members to practice different cube puzzles (3x3, Pyraminx, mirror cube etc.). Organize friendly competitions, time trials, or team-solving challenges to make it fun and engaging.",
	"event_outcome": "The outcomes of this event are that participants enhanced their solving techniques and speed, Participants learnt new algorithms, tricks, and strategies from each other. Participants experienced a sense of accomplishment through challenges or competitions.",
	"event_participant_feedback": "The event participants provided positive feedback, with many participants reporting that they enjoyed the event and learned new techniques.",
	"event_flyer_image_url": "https://drive.usercontent.google.com/download?id=1J8BesUsIu6c3npyXsqtnM1WlocpvB15H&authuser=0",
	"event_attendance": "https://drive.usercontent.google.com/download?id=1J8BesUsIu6c3npyXsqtnM1WlocpvB15H&authuser=0",
	"event_certificate_image_url": "https://drive.usercontent.google.com/download?id=1J8BesUsIu6c3npyXsqtnM1WlocpvB15H&authuser=0",
	"event_winner_details": "The winners of the event are the participants who performed the best in the event.",
"event_recommended_actions": "Participants should practice the techniques they learned in the event and continue to improve their solving skills.",
	"event_expenses": "The event expenses will be covered by the Cubing Club.",
	"event_revenue": "The event revenue will be used to cover the expenses of the event.",
	"event_picture_1": "https://drive.usercontent.google.com/download?id=1J8BesUsIu6c3npyXsqtnM1WlocpvB15H&authuser=0",
	"event_picture_2": "https://drive.usercontent.google.com/download?id=1J8BesUsIu6c3npyXsqtnM1WlocpvB15H&authuser=0",
	"event_picture_3": "https://drive.usercontent.google.com/download?id=1J8BesUsIu6c3npyXsqtnM1WlocpvB15H&authuser=0",
	"event_picture_4": "https://drive.usercontent.google.com/download?id=1J8BesUsIu6c3npyXsqtnM1WlocpvB15H&authuser=0"
	
}`
