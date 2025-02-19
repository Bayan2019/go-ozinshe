-- name: CreateUser :one
INSERT INTO users(name, email, password_hash)
VALUES (?, ?, ?)
RETURNING id;
--

-- name: GetUsers :many
SELECT * FROM users;
--

-- name: GetUserById :one
SELECT * FROM users WHERE id = ?;
--

-- name: UpdateUser :exec
UPDATE users
SET name = ?,
    email = ?
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

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = ?;
--