-- +goose Up
CREATE TABLE images(
    id TEXT PRIMARY KEY,
    project_id INTEGER NOT NULL REFERENCES projects(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE images;