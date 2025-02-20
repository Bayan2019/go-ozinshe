-- +goose Up
CREATE TABLE age_categories(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE age_categories;