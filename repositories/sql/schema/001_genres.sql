-- +goose Up
CREATE TABLE genres(
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE genres;