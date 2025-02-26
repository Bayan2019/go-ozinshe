-- +goose Up
ALTER TABLE projects ADD COLUMN cover TEXT REFERENCES images(id);

-- +goose Down
ALTER TABLE projects DROP COLUMN cover;