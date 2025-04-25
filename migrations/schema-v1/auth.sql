CREATE TABLE component(
    id SERIAL PRIMARY KEY,
    component_name VARCHAR(50),
    code_name VARCHAR(50),
    enabled BOOLEAN DEFAULT true
);

CREATE TABLE action(
    id VARCHAR(26) PRIMARY KEY,
    code_name VARCHAR(50),
    display_name VARCHAR(50)
);

CREATE TABLE role(
    id VARCHAR(26) PRIMARY KEY,
    created_at INT DEFAULT EXTRACT(EPOCH FROM current_timestamp)::INT,
    deleted_at INT DEFAULT 0,
    name VARCHAR(50),
    description VARCHAR(250),
    is_active BOOLEAN DEFAULT true,
    is_default BOOLEAN DEFAULT false
);

CREATE TABLE role_permissions(
    role_id VARCHAR(26),
    action_id VARCHAR(26),
    component_id INT,
    CONSTRAINT fk_role_id FOREIGN KEY (role_id) REFERENCES role(id) ON DELETE CASCADE,
    CONSTRAINT fk_action_id FOREIGN KEY (action_id) REFERENCES action(id) ON DELETE CASCADE,
    CONSTRAINT fk_component_id FOREIGN KEY (component_id) REFERENCES component(id) ON DELETE CASCADE
);

CREATE TABLE auth_config(
    id VARCHAR(26) PRIMARY KEY,
    password_min_length SMALLINT NOT NULL DEFAULT 8,
    numeric_present BOOLEAN NOT NULL DEFAULT true,
    uppercase_present BOOLEAN NOT NULL DEFAULT true,
    lowercase_present BOOLEAN NOT NULL DEFAULT true,
    special_characters_present BOOLEAN NOT NULL DEFAULT true,
    password_reset_days SMALLINT,
    password_reset_alert_days SMALLINT,
    max_password_retries SMALLINT DEFAULT 5,
    max_active_session SMALLINT DEFAULT 2
);

CREATE TABLE users (
    id VARCHAR(26) PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    user_type VARCHAR(10) DEFAULT 'USER',
    password VARCHAR(200) NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100),
    email VARCHAR(100),
    country_code VARCHAR(5),
    contact_number VARCHAR(20),
    password_tries  INT DEFAULT 0,
    account_status VARCHAR(20) DEFAULT 'INVITED',
    change_password BOOLEAN DEFAULT true,
    created_at INT DEFAULT EXTRACT(EPOCH FROM current_timestamp)::INT,
    blocked_time INT,
    deleted_time INT DEFAULT 0,
    active_sessions INT DEFAULT 0,
    auth_config_id VARCHAR(26),
    last_password_updated INT,
    role_id VARCHAR(26) NOT NULL,
    CONSTRAINT fk_password_config FOREIGN KEY (auth_config_id) REFERENCES auth_config(id),
    CONSTRAINT fk_role_id FOREIGN KEY (role_id) REFERENCES role(id)
);

CREATE TABLE auth_tokens(
    id VARCHAR(26) PRIMARY KEY,
    token TEXT NOT NULL,
    user_id VARCHAR(26) NOT NULL,
    created_at INT DEFAULT EXTRACT(EPOCH FROM current_timestamp)::INT,
    deleted_at INT DEFAULT 0,
    is_expired BOOLEAN DEFAULT false,
    ip VARCHAR(50),
    is_active BOOLEAN DEFAULT true,
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id)
);