-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);
-- +goose Down
DROP TABLE users;
