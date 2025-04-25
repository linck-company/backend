package hecatepg

func (hdb *HecatePostgres) getEventsQuery() string {
	query := `
		SELECT
		    e.id,
		    e.event_title,
		    e.event_nature,
		    e.event_date,
		    e.event_start_time,
		    e.event_end_time,
		    e.event_host,
		    e.event_description,
		    e.event_location,
		    e.event_organizer_contact_info,
		    e.event_subject_area, 
		    e.event_resource_person,
		    e.event_affiliation,
		    e.event_resource_person_profile,
		    e.event_objective,
		    e.event_flyer_image_url,
		    e.club_name,
		    e.club_logo_image_url,
		    e.club_id,
		    e.club_organizer_name,
		    COUNT(ur.username) AS registered_count,
			e.is_active,
		    CASE
		        WHEN COUNT(ur.username) > 0 THEN TRUE
		        ELSE FALSE
		    END AS is_user_registered
		FROM
		    events e
		LEFT JOIN
		    usereventregistrations ur ON e.id = ur.event_id
		GROUP BY
		    e.id, e.event_title, e.event_nature, e.event_date, e.event_start_time,
		    e.event_end_time, e.event_host, e.event_description, e.event_location,
		    e.event_organizer_contact_info, e.event_subject_area, e.event_resource_person,
		    e.event_affiliation, e.event_resource_person_profile, e.event_objective,
		    e.event_flyer_image_url, e.club_name, e.club_logo_image_url, e.club_id,
		    e.club_organizer_name;
	`
	return query
}

func (hdb *HecatePostgres) createEventQuery() string {
	query := `
		INSERT INTO events (
			event_title,
			event_nature,
			event_date,
			event_start_time,
			event_end_time,
			event_host,
			event_location,
			event_organizer_contact_info,
			event_subject_area,
			event_resource_person,
			event_affiliation,
			event_resource_person_profile,
			event_description,
			event_objective,
			event_flyer_image_url,
			club_name,
			club_id,
			club_logo_image_url,
			club_organizer_name
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
			$11, $12, $13, $14, $15, $16, $17, $18, $19
		) RETURNING id;
	`
	return query
}

func (hdb *HecatePostgres) getStudentRecordQuery() string {
	query := `
		SELECT 
			e.event_title,
			e.event_date,
			e.club_name
		FROM 
			usereventregistrations ur
		JOIN 
			events e ON ur.event_id = e.id
		WHERE 
			ur.username = $1;
	`
	return query
}
