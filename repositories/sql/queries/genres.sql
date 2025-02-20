-- name: GetGenres :many
SELECT * FROM genres;
--

-- name: GetGenreById :one
SELECT * FROM genres WHERE id = ?;
--

-- name: CreateGenre :one
INSERT INTO genres(title)
VALUES (?)
RETURNING id;
--

-- name: UpdateGenre :exec
UPDATE genres 
SET title = ? 
WHERE id = ?;
--

-- name: DeleteGenre :exec
DELETE FROM genres WHERE id = ?;
--

-- name: GetAllGenresOfProject :many
SELECT g.* FROM genres AS g
JOIN projects_genres AS mg
ON g.id = mg.genre_id
WHERE mg.project_id = ?;
--