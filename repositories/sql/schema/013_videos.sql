-- +goose Up
CREATE TABLE videos(
    id TEXT PRIMARY KEY,
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    project_id INTEGER NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    season INT NOT NULL DEFAULT 0,
    serie INT NOT NULL DEFAULT 0,
    UNIQUE(project_id, season, serie)
);

-- +goose Down
DROP TABLE videos;