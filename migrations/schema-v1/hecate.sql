CREATE TABLE events (
    id SERIAL PRIMARY KEY,
    event_title TEXT,
    event_nature TEXT,
    event_date DATE,
    event_start_time TIME,
    event_end_time TIME,
    event_host TEXT,
    event_location TEXT,
    event_description TEXT,
    external_attendees INT,
    faculty_attendees INT,
    phd_attendees INT,
    pg_attendees INT,
    ug_attendees INT,
    event_organizer_contact_info TEXT,
    event_subject_area TEXT,
    event_resource_person TEXT,
    event_affiliation TEXT,
    event_resource_person_profile TEXT,
    event_objective TEXT,
    event_outcome TEXT,
    event_participant_feedback TEXT,
    event_flyer_image_url TEXT,
    event_attendance TEXT,
    event_certificate_image_url TEXT,
    event_winner_details TEXT,
    event_recommended_actions TEXT,
    event_expenses TEXT,
    event_revenue TEXT,
    event_picture_1 TEXT,
    event_picture_2 TEXT,
    event_picture_3 TEXT,
    event_picture_4 TEXT,
    is_active BOOLEAN DEFAULT true;
    club_name TEXT,
    club_id VARCHAR(255),
    club_logo_image_url TEXT,
    club_organizer_name TEXT
);

INSERT INTO Events (
    event_title, event_nature, event_date, event_start_time, event_end_time,
    event_host, event_location, external_attendees, faculty_attendees, phd_attendees,
    pg_attendees, ug_attendees, event_organizer_contact_info, event_subject_area,
    event_resource_person, event_affiliation, event_resource_person_profile, event_objective,
    event_outcome, event_participant_feedback, event_flyer_image_url, event_attendance,
    event_certificate_image_url, event_winner_details, event_recommended_actions,
    event_expenses, event_revenue, event_picture_1, event_picture_2, event_picture_3,
    event_picture_4, club_name, club_id, club_logo_image_url, club_organizer_name
) VALUES (
    'FAST WEDNESDAYS', 'Competition', '2024-03-12', '16:00:00', '18:00:00',
    'Cubing Club', 'S310', 50, 10, 5, 15, 20, 'Cubing Club', 'Cubing',
    'NA', 'Cubing Club', 'NA', 
    'The Objective of the event is to teach new solving techniques and allow members to practice different cube puzzles (3x3, Pyraminx, mirror cube etc.). Organize friendly competitions, time trials, or team-solving challenges to make it fun and engaging.',
    'The outcomes of this event are that participants enhanced their solving techniques and speed, Participants learnt new algorithms, tricks, and strategies from each other. Participants experienced a sense of accomplishment through challenges or competitions.',
    'The event participants provided positive feedback, with many participants reporting that they enjoyed the event and learned new techniques.',
    'https://drive.usercontent.google.com/download?id=1J8BesUsIu6c3npyXsqtnM1WlocpvB15H&authuser=0',
    'https://drive.usercontent.google.com/download?id=1J8BesUsIu6c3npyXsqtnM1WlocpvB15H&authuser=0',
    'https://drive.usercontent.google.com/download?id=1J8BesUsIu6c3npyXsqtnM1WlocpvB15H&authuser=0',
    'The winners of the event are the participants who performed the best in the event.',
    'Participants should practice the techniques they learned in the event and continue to improve their solving skills.',
    'The event expenses will be covered by the Cubing Club.',
    'The event revenue will be used to cover the expenses of the event.',
    'https://drive.usercontent.google.com/download?id=1J8BesUsIu6c3npyXsqtnM1WlocpvB15H&authuser=0',
    'https://drive.usercontent.google.com/download?id=1J8BesUsIu6c3npyXsqtnM1WlocpvB15H&authuser=0',
    'https://drive.usercontent.google.com/download?id=1J8BesUsIu6c3npyXsqtnM1WlocpvB15H&authuser=0',
    'https://drive.usercontent.google.com/download?id=1J8BesUsIu6c3npyXsqtnM1WlocpvB15H&authuser=0',
    'Cubing Club', 'cubing_club', 'https://drive.usercontent.google.com/download?id=1J8BesUsIu6c3npyXsqtnM1WlocpvB15H&authuser=0',
    'SELF DEVELOPMENT'
);

CREATE TABLE usereventregistrations(
    username VARCHAR(50) NOT NULL UNIQUE,
    event_id INT NOT NULL REFERENCES Events(id) ON DELETE CASCADE,
    registration_timestamp TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (username, event_id)
);

CREATE INDEX idx_user_event_registrations_event_id ON UserEventRegistrations(event_id);
CREATE INDEX idx_user_event_registrations_user_id ON UserEventRegistrations(username);
ALTER TABLE usereventregistrations DROP CONSTRAINT usereventregistrations_username_key;

CREATE TABLE events (
    event_id VARCHAR(26) PRIMARY KEY,
    club_id VARCHAR(26),
    club_name TEXT,
    club_organization TEXT,
    title TEXT NOT NULL,
    nature TEXT NOT NULL,  
    date DATE,
    time TIME,
    location TEXT,
    host TEXT, 
    subject_area TEXT,
    objective TEXT,
    outcome TEXT,
    expenses TEXT,
    revenue TEXT,
    recommended_actions TEXT,
    event_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    event_ended_at TIMESTAMP,
    is_active BOOLEAN DEFAULT true,
    is_completed BOOLEAN DEFAULT true
);
CREATE TABLE organizers (
    organizer_id UUID PRIMARY KEY,
    event_id VARCHAR(26),
    name TEXT,
    contact_info TEXT,
    role TEXT,
);

CREATE TABLE resource_persons (
    person_id UUID PRIMARY KEY,
    event_id VARCHAR(26),
    name TEXT,
    profile TEXT,
    affiliation TEXT,
);

CREATE TABLE event_winners (
    event_id VARCHAR(26),
    winner_id VARCHAR(26),
    PRIMARY KEY (event_id, winner_id),
    FOREIGN KEY (event_id) REFERENCES events(event_id) ON DELETE CASCADE
);

CREATE TABLE attendance (
    event_id VARCHAR(26),
    user_id VARCHAR(26),
    role TEXT CHECK (role IN ('external', 'faculty', 'phd', 'pg', 'ug')),
    PRIMARY KEY (event_id, user_id)
    FOREIGN KEY (event_id) REFERENCES events(event_id) ON DELETE CASCADE
);


CREATE TABLE media_assets (
    media_id UUID PRIMARY KEY,
    event_id VARCHAR(26),
    category TEXT CHECK (category IN ('flyer', 'certificate', 'geo_photo', 'attendance_sheet', 'feedback')),
    url TEXT NOT NULL,
    caption TEXT,
    FOREIGN KEY (event_id) REFERENCES events(event_id) ON DELETE CASCADE
);

CREATE TABLE checklist (
    item_id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE event_checklist (
    event_id VARCHAR(26),
    item_id INT,
    submitted BOOLEAN DEFAULT FALSE,
    PRIMARY KEY (event_id, item_id),
    FOREIGN KEY (event_id) REFERENCES events(event_id) ON DELETE CASCADE,
    FOREIGN KEY (item_id) REFERENCES checklist(item_id) ON DELETE CASCADE
);

CREATE TABLE approvals (
    approval_id UUID PRIMARY KEY,
    event_id VARCHAR(26),
    role TEXT, 
    name TEXT,
    department TEXT,
    signed BOOLEAN DEFAULT FALSE,
    signature_image_url TEXT,
    FOREIGN KEY (event_id) REFERENCES events(event_id) ON DELETE CASCADE
);
