-- name: CreateUser :one
INSERT INTO users(name, email, password_hash)
VALUES (?, ?, ?)
RETURNING id;
--

-- name: GetUsers :many
SELECT * FROM users;
--

-- name: GetUsersOfRole :many
SELECT u.*
FROM users AS u
JOIN users_roles AS ur
ON u.id = ur.user_id
WHERE ur.role_id = ?;
--

-- name: GetUserById :one
SELECT * FROM users WHERE id = ?;
--

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = ?;
--

-- name: UpdateUser :exec
UPDATE users
SET updated_at = CURRENT_TIMESTAMP,
    name = ?,
    email = ?,
    date_of_birth = ?,
    phone = ?
WHERE id = ?;
--

-- name: ChangePassword :exec
UPDATE users
SET password_hash = ?
WHERE id = ?;
--

-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?;
--