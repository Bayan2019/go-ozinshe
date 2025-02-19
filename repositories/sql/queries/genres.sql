-- name: GetGenres :many
SELECT * FROM genres;

-- name: GetGenreById :one
SELECT * FROM genres WHERE id = ?;

-- name: CreateGenre :one
INSERT INTO genres(title)
VALUES (?)
RETURNING *;

-- name: UpdateGenre :one
UPDATE genres 
SET title = ? 
WHERE id = ?
RETURNING *;

-- name: DeleteGenre :exec
DELETE FROM genres WHERE id = ?;

-- name: GetAllGenresOfMovie :many
SELECT g.* FROM genres AS g
JOIN movies_genres AS mg
ON g.id = mg.genre_id
WHERE mg.movie_id = ?;