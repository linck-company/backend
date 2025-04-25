package odysseyhandler

import (
	odysseymodels "backendv1/internal/models/odyssey"
	"encoding/json"
	"net/http"
)

type OdysseyHandler struct {
}

func (o *OdysseyHandler) GetReportDetails(w http.ResponseWriter, r *http.Request) {
	resp := odysseymodels.ReportDetailsResponse{
		StatusCode:    http.StatusOK,
		EventTitle:    "FAST WEDNESDAYS",
		EventNature:   "Competition",
		EventDate:     "12-03-2024",
		EventTime:     "4:00 pm",
		EventHost:     "Cubing Club",
		EventLocation: "S310",
		EventNumberAttendees: struct {
			ExternalAttendees int `json:"external_attendees"`
			FacultyAttendees  int `json:"faculty_attendees"`
			PhdAttendees      int `json:"phd_attendees"`
			PgAttendees       int `json:"pg_attendees"`
			UgAttendees       int `json:"ug_attendees"`
		}{
			ExternalAttendees: 50,
			FacultyAttendees:  10,
			PhdAttendees:      5,
			PgAttendees:       15,
			UgAttendees:       20,
		},
		EventOrganizerContactInfo:            "Cubing Club",
		EventSubjectArea:                     "Cubing",
		EventResourcePerson:                  "NA",
		EventAffiliation:                     "Cubing Club",
		EventResourcePersonProfile:           "NA",
		EventObjective:                       "The Objective of the event is to teach new solving techniques and allow members to practice different cube puzzles (3x3, Pyraminx, mirror cube etc.).\n Organize friendly competitions, time trials, or team-solving challenges to make it fun and engaging.",
		EventOutcome:                         "The outcomes of this event are that participants enhanced their solving techniques and speed, Participants learnt new algorithms, tricks, and strategies from each other. Participants experienced a sense of accomplishment through challenges or competitions.",
		EventParticipantFeedback:             "The event participants provided positive feedback, with many participants reporting that they enjoyed the event and learned new techniques.",
		EventFlyerImageUrl:                   "https://drive.usercontent.google.com/download?id=1J8BesUsIu6c3npyXsqtnM1WlocpvB15H&authuser=0",
		EventAttendanceImageUrl:              "https://drive.usercontent.google.com/download?id=1J8BesUsIu6c3npyXsqtnM1WlocpvB15H&authuser=0",
		EventParticipantsCertificateImageUrl: "https://drive.usercontent.google.com/download?id=1J8BesUsIu6c3npyXsqtnM1WlocpvB15H&authuser=0",
		EventWinnerDetails:                   "The winners of the event are the participants who performed the best in the event.",
		EventRecommendedActions:              "Participants should practice the techniques they learned in the event and continue to improve their solving skills.",
		EventExpenses:                        "The event expenses will be covered by the Cubing Club.",
		EventRevenue:                         "The event revenue will be used to cover the expenses of the event.",
		EventPictures: struct {
			Image1Url string `json:"event_picture_1"`
			Image2Url string `json:"event_picture_2"`
			Image3Url string `json:"event_picture_3"`
			Image4Url string `json:"event_picture_4"`
		}{
			Image1Url: "https://drive.usercontent.google.com/download?id=1J8BesUsIu6c3npyXsqtnM1WlocpvB15H&authuser=0",
			Image2Url: "https://drive.usercontent.google.com/download?id=1J8BesUsIu6c3npyXsqtnM1WlocpvB15H&authuser=0",
			Image3Url: "https://drive.usercontent.google.com/download?id=1J8BesUsIu6c3npyXsqtnM1WlocpvB15H&authuser=0",
			Image4Url: "https://drive.usercontent.google.com/download?id=1J8BesUsIu6c3npyXsqtnM1WlocpvB15H&authuser=0",
		},
	}
	w.WriteHeader(resp.StatusCode)
	json.NewEncoder(w).Encode(resp)
}
