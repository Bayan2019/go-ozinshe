-- name: AddImage2Movie :exec
INSERT INTO images(id, project_id, href)
VALUES (?, ?, ?);
--

-- name: GetImage :one
SELECT * FROM images
WHERE id = ?;
--

-- name: GetImagesOfProject :many
SELECT * FROM images
WHERE project_id = ?;
--

-- name: GetImages :many
SELECT * FROM images;
--

-- name: UpdateImage :exec
UPDATE images
SET updated_at = CURRENT_DATE,
    project_id = ?,
    href = ?
WHERE id = ?;
--

-- name: DeleteImage :exec
DELETE FROM images WHERE id = ?;
--