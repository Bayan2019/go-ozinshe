-- name: AddGenre2Project :exec
INSERT INTO projects_genres(project_id, genre_id)
VALUES (?, ?);
--

-- name: DeleteGenresOfProject :exec
DELETE FROM projects_genres WHERE project_id = ?;
--