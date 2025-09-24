-- name: CreateUser :one
INSERT INTO users (
  name, email, created_at
) VALUES ($1, $2, $3)
RETURNING *;

