-- name: CreateFeedFollow :one
INSERT INTO feed_follows (id_feeds_follow, create_at,  user_id, feed_id)
VALUES ($1, $2, $3, $4)
RETURNING *;
