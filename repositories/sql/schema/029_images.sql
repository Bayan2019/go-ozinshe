-- +goose Up
ALTER TABLE images ADD COLUMN href TEXT NOT NULL DEFAULT '';

-- +goose Down
ALTER TABLE images DROP COLUMN href;