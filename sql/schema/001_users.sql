-- +goose Up
CREATE TABLE users (
    id UUID Primary Key,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    email TEXT NOT NULL UNIQUE
);
-- +goose Down
DROP TABLE users;