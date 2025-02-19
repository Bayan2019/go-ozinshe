-- +goose Up
CREATE TABLE genres(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE genres;