CREATE TABLE IF NOT EXISTS todo_list (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    title VARCHAR NOT NULL,
    description VARCHAR NOT NULL,
    date_created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    date_updated TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deadline TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS credentials (
    username varchar not null,
    password varchar not null
);