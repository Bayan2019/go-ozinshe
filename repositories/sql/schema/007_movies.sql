-- +goose Up
ALTER TABLE movies ADD COLUMN poster_url TEXT NOT NULL DEFAULT '';

-- +goose Down
ALTER TABLE movies DROP COLUMN poster_url;