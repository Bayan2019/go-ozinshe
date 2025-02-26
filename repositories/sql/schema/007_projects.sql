-- +goose Up
CREATE TABLE projects(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    title TEXT NOT NULL DEFAULT '',
    description TEXT NOT NULL DEFAULT '',
    type_id INTEGER NOT NULL REFERENCES types(id),
    duration_in_mins INTEGER NOT NULL DEFAULT 0,
    release_year INT NOT NULL DEFAULT 0,
    director TEXT NOT NULL DEFAULT '',
    producer TEXT NOT NULL DEFAULT ''
    -- cover TEXT REFERENCES images(id)
);

-- +goose Down
DROP TABLE projects;