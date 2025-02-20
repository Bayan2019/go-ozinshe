-- +goose Up
CREATE TABLE types(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE types;