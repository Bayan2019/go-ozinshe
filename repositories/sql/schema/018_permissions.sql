-- +goose Up
CREATE TABLE permissions(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE permissions;