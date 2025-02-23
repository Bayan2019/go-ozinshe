-- +goose Up
INSERT INTO roles(title, projects, genres, age_categories, types, users, roles)
VALUES ('admïn', 3, 3, 3, 3, 3, 3),
    ('paydalanwşı', 2, 2, 2, 2, 1, 1);

-- +goose Down
DELETE FROM roles;