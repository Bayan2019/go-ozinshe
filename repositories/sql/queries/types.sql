-- name: GetTypes :many
SELECT * FROM types;
--

-- name: GetTypeById :one
SELECT * FROM types WHERE id = ?;
--

-- name: CreateType :one
INSERT INTO types(title)
VALUES (?)
RETURNING id;
--

-- name: UpdateType :exec
UPDATE types 
SET title = ? 
WHERE id = ?;
--

-- name: DeleteType :exec
DELETE FROM types WHERE id = ?;
--