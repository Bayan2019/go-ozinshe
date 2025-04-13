-- name: GetProjects :many
SELECT * FROM projects;
--

-- name: GetProjectsOfGenrers :many
SELECT p.* FROM projects AS p
JOIN projects_genres AS pg 
ON p.id = pg.project_id
WHERE pg.genre_id IN (sqlc.slice('ids'));
--

-- -- name: PragmaCaseSensitiveOFF :exec
-- PRAGMA case_sensitive_like = OFF;
-- --

-- name: GetProjectsOfGenresAndSearch :many
SELECT p.* FROM projects AS p
JOIN projects_genres AS pg 
ON p.id = pg.project_id
WHERE pg.genre_id IN (sqlc.slice('ids')) 
    AND ((p.title LIKE '%'+@search+'%') 
        OR (p.description LIKE '%' + @search + '%')
        OR (p.keywords LIKE '%' + @search + '%'));
--

-- name: GetProjectsSearch :many
SELECT * FROM projects
WHERE ((title LIKE '%' + @search + '%') 
        OR (description LIKE '%' + @search + '%')
        OR (keywords LIKE '%' + @search + '%'));
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

-- name: CreateProject :one
INSERT INTO projects(title, description, type_id, duration_in_mins, release_year, director, producer, keywords)
VALUES (?, ?, ?, ?, ?, ?, ?, ?)
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
    producer = ?,
    keywords = ?
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