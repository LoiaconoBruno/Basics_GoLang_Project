-- +goose Up 
CREATE TABLE feeds(
    id_feeds UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    name TEXT NOT NULL,
    url TEXT NOT NULL UNIQUE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);
    

-- +goose Down
DROP TABLE feeds;
