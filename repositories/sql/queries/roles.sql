-- name: GetRoles :many
SELECT * FROM roles;
--

-- name: GetRoleById :one
SELECT * FROM roles WHERE id = ?;
--

-- name: CreateRole :one
INSERT INTO roles(title, projects, genres, age_categories, types, users, roles)
VALUES (?, ?, ?, ?, ?, ?, ?)
RETURNING id;
--

-- name: UpdateType :exec
UPDATE roles 
SET title = ?,
    projects = ?,
    genres = ?,
    age_categories = ?,
    types = ?,
    users = ?,
    roles = ?
WHERE id = ?;
--

-- name: GetRolesOfUser :many
SELECT r.*
FROM roles AS r
JOIN users_roles AS ur
ON r.id = ur.role_id
WHERE ur.user_id = ?;
--

-- name: DeleteType :exec
DELETE FROM roles WHERE id = ?;
--