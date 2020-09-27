-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE users (
    id BINARY(32) PRIMARY KEY NOT NULL,
    email TEXT,
    username TEXT,
    photo_url TEXT,
    token TEXT
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE users;