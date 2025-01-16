CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(60) NOT NULL,
    phone VARCHAR(12) NOT NULL
);

CREATE TABLE sms_codes (
    user_id INTEGER NOT NULL,
    code INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (user_id)
);

INSERT INTO users (username, phone) VALUES ('GoGopher', '+79993337788');