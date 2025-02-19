-- +goose Up
CREATE TABLE movies(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL DEFAULT '',
    description TEXT NOT NULL DEFAULT '',
    release_year INT NOT NULL DEFAULT 0,
    director TEXT NOT NULL DEFAULT '',
    rating INT NOT NULL DEFAULT 0,
    is_watched BOOLEAN NOT NULL DEFAULT FALSE,
    trailer_url TEXT NOT NULL DEFAULT ''
);

-- +goose Down
DROP TABLE movies;