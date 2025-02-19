-- +goose Up
CREATE TABLE watchlist(
    movie_id BIGINT NOT NULL REFERENCES movies(id) ON DELETE CASCADE,
    added_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE watchlist;