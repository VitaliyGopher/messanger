CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(60) NOT NULL,
    phone VARCHAR(20) NOT NULL UNIQUE
);

CREATE TABLE sms_codes (
    phone VARCHAR(20) NOT NULL,
    code INTEGER NOT NULL,
    time_expire INTEGER NOT NULL
);

INSERT INTO users (username, phone) VALUES ('GoGopher', '+79993337788');