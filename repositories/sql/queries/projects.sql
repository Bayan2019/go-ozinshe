-- name: GetProjects :many
SELECT * FROM projects;
--

-- name: GetProjectById :one
SELECT * FROM projects WHERE id = ?;
--

-- name: GetProjectsOfGenre :many
SELECT p.* FROM projects AS p
JOIN projects_genres AS pg 
ON p.id = pg.project_id
WHERE pg.genre_id = ?;
--

-- name: GetProjectsOfAgeCategory :many
SELECT p.* FROM projects AS p
JOIN projects_age_categories AS pac 
ON p.id = pac.project_id
WHERE pac.age_category_id = ?;
--

-- name: GetProjectsOfType :many
SELECT * FROM projects 
WHERE type_id = ?;
--

-- name: GetProjectsOfSearch :many
SELECT * FROM projects 
WHERE LOWER(title) LIKE '%' + LOWER(?) + '%';
--

-- name: CreateProject :one
INSERT INTO projects(title, description, type_id, duration_in_mins, release_year, director, producer)
VALUES (?, ?, ?, ?, ?, ?, ?)
RETURNING id;
--

-- name: UpdateProject :exec
UPDATE projects
SET updated_at = CURRENT_TIMESTAMP,
    title = ?,
    description = ?,
    type_id = ?,
    duration_in_mins = ?,
    release_year = ?,
    director = ?,
    producer = ?
WHERE id = ?;
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