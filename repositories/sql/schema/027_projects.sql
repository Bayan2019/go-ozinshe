-- +goose Up
ALTER TABLE projects ADD COLUMN keywords TEXT NOT NULL DEFAULT '';

-- +goose Down
ALTER TABLE projects DROP COLUMN keywords;