package odysseymodels

type ReportDetailsResponse struct {
	StatusCode           int    `json:"status_code"`
	EventTitle           string `json:"event_title"`
	EventNature          string `json:"event_nature"`
	EventDate            string `json:"event_date"`
	EventTime            string `json:"event_time"`
	EventHost            string `json:"event_host"`
	EventLocation        string `json:"event_location"`
	EventNumberAttendees struct {
		ExternalAttendees int `json:"external_attendees"`
		FacultyAttendees  int `json:"faculty_attendees"`
		PhdAttendees      int `json:"phd_attendees"`
		PgAttendees       int `json:"pg_attendees"`
		UgAttendees       int `json:"ug_attendees"`
	} `json:"event_number_attendees"`
	EventOrganizerContactInfo            string `json:"event_organizer_contact_info"`
	EventSubjectArea                     string `json:"event_subject_area"`
	EventResourcePerson                  string `json:"event_resource_person"`
	EventAffiliation                     string `json:"event_affiliation"`
	EventResourcePersonProfile           string `json:"event_resource_person_profile"`
	EventObjective                       string `json:"event_objective"`
	EventOutcome                         string `json:"event_outcome"`
	EventParticipantFeedback             string `json:"event_participant_feedback"`
	EventFlyerImageUrl                   string `json:"event_flyer_image_url"`
	EventAttendanceImageUrl              string `json:"event_attendance"`
	EventParticipantsCertificateImageUrl string `json:"event_certificate_image_url"`
	EventWinnerDetails                   string `json:"event_winner_details"`
	EventRecommendedActions              string `json:"event_recommended_actions"`
	EventExpenses                        string `json:"event_expenses"`
	EventRevenue                         string `json:"event_revenue"`
	EventPictures                        struct {
		Image1Url string `json:"event_picture_1"`
		Image2Url string `json:"event_picture_2"`
		Image3Url string `json:"event_picture_3"`
		Image4Url string `json:"event_picture_4"`
	} `json:"event_pictures"`
}
