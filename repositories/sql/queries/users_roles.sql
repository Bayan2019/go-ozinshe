-- name: AddRole2User :exec
INSERT INTO users_roles(user_id, role_id)
VALUES (?, ?);
--

-- name: RemoveRoleFromUser :exec
DELETE FROM users_roles 
WHERE user_id = ? AND role_id = ?;
--