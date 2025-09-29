-- +goose Up

CREATE TABLE feed_follows (
  id_feeds_follow UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  create_at TIMESTAMP NOT NULL DEFAULT now(),
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  feed_id UUID NOT NULL REFERENCES feeds(id_feeds) ON DELETE CASCADE,
  UNIQUE(user_id, feed_id)
);

-- +goose Down
DROP TABLE feed_follows;
