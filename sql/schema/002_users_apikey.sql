-- +goose Up
CREATE EXTENSION IF NOT EXISTS pgcrypto;
ALTER TABLE users ADD COLUMN api_key VARCHAR(64) UNIQUE NOT NULL DEFAULT (
  encode(gen_random_bytes(32), 'hex')
);

-- +goose Down
ALTER TABLE users DROP COLUMN api_key;
