-- name: AddVideo2Movie :exec
INSERT INTO videos(id, project_id, href)
VALUES (?, ?, ?);
--

-- name: AddVideo2Series :exec
INSERT INTO videos(id, project_id, season, serie, href)
VALUES (?, ?, ?, ?, ?);
--

-- name: GetVideo :one
SELECT * FROM videos
WHERE project_id = ? AND
    season = ? AND
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
    serie = ?,
    href = ?
WHERE id = ?;
--

-- name: DeleteVideo :exec
DELETE FROM videos WHERE id = ?;
--