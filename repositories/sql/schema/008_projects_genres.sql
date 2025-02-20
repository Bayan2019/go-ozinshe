-- +goose Up
CREATE TABLE projects_genres(
    project_id INTEGER NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    genre_id INTEGER NOT NULL REFERENCES genres(id) ON DELETE CASCADE,
    UNIQUE(project_id, genre_id)
);

-- +goose Down
DROP TABLE projects_genres;