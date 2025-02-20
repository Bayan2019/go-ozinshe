-- +goose Up
CREATE TABLE roles(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL UNIQUE,
    projects INTEGER NOT NULL REFERENCES permissions(id) DEFAULT 2,
    genres INTEGER NOT NULL REFERENCES permissions(id) DEFAULT 2,
    age_categories INTEGER NOT NULL REFERENCES permissions(id) DEFAULT 2,
    types INTEGER NOT NULL REFERENCES permissions(id) DEFAULT 2,
    users INTEGER NOT NULL REFERENCES permissions(id) DEFAULT 2,
    roles INTEGER NOT NULL REFERENCES permissions(id) DEFAULT 2
);

-- +goose Down
DROP TABLE roles;