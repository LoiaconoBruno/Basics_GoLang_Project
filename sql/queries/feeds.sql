-- name: CreateFeed :one
INSERT INTO feeds(id_feeds, created_at, name, url, user_id)
VALUES($1, $2, $3, $4, $5)
RETURNING *;

--name: GetFeeds :many
SELECT * FROM feeds;
