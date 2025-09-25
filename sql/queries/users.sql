-- name: CreateUser :one
INSERT INTO users (
  id, name, email, created_at, api_key
) VALUES ($1, $2, $3, NOW(), encode(gen_random_bytes(32), 'hex'))
RETURNING *;
