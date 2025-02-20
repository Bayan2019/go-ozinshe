-- name: AddVideo2Movie :exec
INSERT INTO videos(project_id)
VALUES (?);
--

-- name: AddVideo2Series :exec
INSERT INTO videos(project_id, season, serie)
VALUES (?, ?, ?);
--

-- name: GetVideo :one
SELECT * FROM videos
WHERE project_id = ?,
    season = ?,
    serie = ?;
--

-- name: GetVideosOfProject :many
SELECT * FROM videos
WHERE project_id = ?;
--

-- name: GetVideos :many
SELECT * FROM videos;
--

-- name: UpdateVideo :exec
UPDATE videos
SET updated_at = CURRENT_DATE,
    project_id = ?,
    season = ?,
    serie = ?
WHERE id = ?;
--

-- name: DeleteVideo :exec
DELETE FROM videos WHERE id = ?;
--