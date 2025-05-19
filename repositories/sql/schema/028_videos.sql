-- +goose Up
ALTER TABLE videos ADD COLUMN href TEXT NOT NULL DEFAULT '';

-- +goose Down
ALTER TABLE videos DROP COLUMN href;