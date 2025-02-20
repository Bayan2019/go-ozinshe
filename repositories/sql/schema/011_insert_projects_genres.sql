-- +goose Up
INSERT INTO projects_genres(project_id, genre_id)
VALUES(1, 2),
       (1, 6),
       (1, 7),
       (2, 4),
       (2, 5),
       (2, 6),
       (2, 7),
       (3, 4),
       (3, 7),
       (4, 1),
       (4, 2);

-- +goose Down
DELETE FROM projects_genres;