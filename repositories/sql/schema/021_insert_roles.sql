-- +goose Up
INSERT INTO roles(title, projects, genres, age_categories, types, users, roles)
VALUES ('админ', 3, 3, 3, 3, 3, 3);

-- +goose Down
DELETE FROM roles;