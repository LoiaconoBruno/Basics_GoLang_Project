-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
<<<<<<< HEAD
    id UUID PRIMARY KEY DEFAULT
=======
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
>>>>>>> schema
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    created TIMESTAMP NOT NULL DEFAULT now()
);
-- +goose Down
DROP TABLE users;
