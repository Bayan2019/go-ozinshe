-- +goose Up
INSERT INTO projects_age_categories(project_id, age_category_id)
VALUES (1, 2),
       (1, 3),
       (2, 2),
       (2, 3),
       (2, 4),
       (3, 1),
       (3, 2),
       (4, 3),
       (4, 4),
       (4, 5),
       (4, 6),
       (4, 7);

-- +goose Down
DELETE FROM projects_genres;