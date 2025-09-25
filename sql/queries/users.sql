-- name: CreateUser :one
CREATE EXTENSION IF NOT EXISTS pgcrypto;
INSERT INTO users (
  name, email, created_at, api_key
) VALUES ($1, $2, $3,
  encode(gen_random_bytes(32), 'hex')
)
RETURNING *;

