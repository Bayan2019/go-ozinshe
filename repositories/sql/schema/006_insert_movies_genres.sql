-- +goose Up
INSERT INTO movies_genres(movie_id, genre_id)
VALUES (1, 1),
       (1, 2),
       (2, 1),
       (2, 3),
       (2, 4),
       (3, 1),
       (4, 1),
       (4, 5),
       (4, 6),
       (5, 7),
       (5, 1),
       (5, 6),
       (6, 7),
       (6, 8),
       (6, 1),
       (7, 1),
       (7, 2),
       (7, 9),
       (8, 10),
       (8, 11),
       (8, 5),
       (9, 5),
       (9, 4),
       (9, 1),
       (10, 12),
       (10, 7),
       (10, 1);

-- +goose Down
DELETE FROM movies_genres;