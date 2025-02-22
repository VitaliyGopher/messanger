CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(60) NOT NULL,
    email VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE verification_codes (
    email VARCHAR(50) NOT NULL,
    code INTEGER NOT NULL,
    time_expire INTEGER NOT NULL
);

INSERT INTO users (username, email) VALUES ('GoGopher', 'qweqwe@edu.ystu.ru');