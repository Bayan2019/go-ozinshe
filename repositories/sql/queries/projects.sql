-- name: GetProjects :many
SELECT * FROM projects;
--

-- name: GetProjectById :one
SELECT * FROM projects WHERE id = ?;
--

-- name: CreateProject :one
INSERT INTO projects(title, description, type_id, duration_in_mins, release_year, director, producer)
VALUES (?, ?, ?, ?, ?, ?, ?)
RETURNING id;
--

-- name: GetWatchlistProjects :many
SELECT p.*
FROM projects AS p
JOIN watchlist AS w
ON p.id = w.project_id
WHERE w.user_id = ?
ORDER BY added_at;
--

-- name: UpdateProjects :exec
UPDATE projects 
SET updated_at = CURRENT_TIMESTAMP,
    title = ?, 
    description = ?, 
    type_id = ?,
    duration_in_mins = ?,
    release_year = ?, 
    director = ?,
    producer = ?
WHERE id = ?
--

-- name: SetCover :exec
UPDATE projects
SET updated_at = CURRENT_TIMESTAMP,
    cover = ?
WHERE id = ?;
--

-- name: DeleteProject :exec
DELETE FROM projects WHERE id = ?;
--