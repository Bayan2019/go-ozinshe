-- name: GetProjects :many
SELECT * FROM projects;
--

-- name: GetProjectById :one
SELECT * FROM projects WHERE id = ?;
--

-- name: CreateProject :one
INSERT INTO projects(title, description, type_id, duration_in_mins, release_year, director, producer)
VALUES (?, ?, ?, ?, ?, ?, ?)
RETURNING *;
--

-- name: UpdateProjects :one
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
RETURNING *;
--

-- name: DeleteMovie :exec
DELETE FROM projects WHERE id = ?;
--