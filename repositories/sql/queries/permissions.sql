-- name: GetPermissions :many
SELECT * FROM permissions;
--

-- name: GetPermissionById :one
SELECT * FROM permissions WHERE id = ?;
--

-- name: CreatePermission :one
INSERT INTO permissions(title)
VALUES (?)
RETURNING id;
--

-- name: UpdatePermission :exec
UPDATE permissions 
SET title = ? 
WHERE id = ?;
--

-- name: DeletePermission :exec
DELETE FROM permissions WHERE id = ?;
--