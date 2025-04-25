CREATE TABLE entities (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    club_organization TEXT,
    year_established INT NOT NULL,
    founder VARCHAR(255) NOT NULL,
    co_founder VARCHAR(255),
    club_logo_image_url VARCHAR(255),
    club_banner_image_url VARCHAR(255),
    club_website_url VARCHAR(255),
    club_facebook_url VARCHAR(255),
    club_twitter_url VARCHAR(255),
    club_instagram_url VARCHAR(255),
    club_youtube_url VARCHAR(255),
    club_linkedin_url VARCHAR(255),
    club_email VARCHAR(255)
);

CREATE TABLE present_core_members (
    id SERIAL PRIMARY KEY,
    eid VARCHAR(255) NOT NULL,
    name VARCHAR(255),
    title VARCHAR(255),
    image_url VARCHAR(255),
    FOREIGN KEY (eid) REFERENCES entities(id) ON DELETE CASCADE
);

CREATE TABLE entity_core_members(
    id SERIAL PRIMARY KEY,
    eid VARCHAR(255) NOT NULL,
    name VARCHAR(255),
    title VARCHAR(255),
    image_url VARCHAR(255),
    FOREIGN KEY (eid) REFERENCES entities(id) ON DELETE CASCADE
);

CREATE TABLE legacy_holders (
    legacy_id SERIAL PRIMARY KEY,
    eid VARCHAR(255) NOT NULL,
    year_start VARCHAR(7) NOT NULL,
    year_end VARCHAR(7) NOT NULL,
    name VARCHAR(100) NOT NULL,
    title VARCHAR(100),
    image_url VARCHAR(255),
    FOREIGN KEY (eid) REFERENCES entities(id) ON DELETE CASCADE
);

CREATE TABLE user_registered_entities(
    id SERIAL PRIMARY KEY,
    entities_id VARCHAR(255) NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    FOREIGN KEY (entities_id) REFERENCES entities(id) ON DELETE CASCADE
);